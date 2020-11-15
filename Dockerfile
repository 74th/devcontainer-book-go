# -----------------------------
# 依存ライブラリをもつイメージ
FROM golang:1.15-buster AS base

# go.mod から モジュールのダウンロード
WORKDIR /tmp/gomod
ADD go.mod .
ADD go.sum .
RUN go mod download

# -----------------------------
# Dev Container
FROM base AS devcontainer

# VS Code のスクリプトを使う
ARG USERNAME=vscode
ARG USER_UID=1000
ARG USER_GID=$USER_UID
RUN curl -o /tmp/common-debian.sh \
  -L https://raw.githubusercontent.com/microsoft/vscode-dev-containers/master/script-library/common-debian.sh \
  && /bin/bash /tmp/common-debian.sh \
    "false" \
    "${USERNAME}" \
    "${USER_UID}" \
    "${USER_GID}" \
    "false" \
  && rm /tmp/common-debian.sh
RUN curl -o /tmp/go-debian.sh \
  -L https://raw.githubusercontent.com/microsoft/vscode-dev-containers/master/script-library/go-debian.sh \
  && /bin/bash /tmp/go-debian.sh \
    "${GOLANG_VERSION}" \
    "/usr/loca/go" \
    "${GOPATH}" \
    "${USERNAME}" \
    "true" \
  && rm /tmp/go-debian.sh
# baseではrootユーザでgo mod downloadしたため
# vscodeユーザでアクセス可能にする
RUN chown -R $USERNAME:$USERNAME /go

# CIとも併用するツールをいれる
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.30.0

USER $USERNAME
CMD ["sleep", "infinity"]

# -----------------------------
# Dev Container
FROM base AS builder

WORKDIR /workspace
ADD . .
RUN go build -o server cmd/server/main.go

# -----------------------------
# Dev Container
FROM debian:buster AS app

COPY --from=builder /workspace/server .
EXPOSE 8080

CMD ["./server"]