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
# INSTALL_ZSH=${1:-"true"}
# USERNAME=${2:-"automatic"}
# USER_UID=${3:-"automatic"}
# USER_GID=${4:-"automatic"}
# UPGRADE_PACKAGES=${5:-"true"}
# INSTALL_OH_MYS=${6:-"true"}
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

# TARGET_GO_VERSION=${1:-"latest"}
# TARGET_GOROOT=${2:-"/usr/local/go"}
# TARGET_GOPATH=${3:-"/go"}
# USERNAME=${4:-"automatic"}
# UPDATE_RC=${5:-"true"}
# INSTALL_GO_TOOLS=${6:-"true"}
RUN curl -o /tmp/go-debian.sh \
  -L https://raw.githubusercontent.com/microsoft/vscode-dev-containers/master/script-library/go-debian.sh \
  && /bin/bash /tmp/go-debian.sh \
    "${GOLANG_VERSION}" \
    "/usr/loca/go" \
    "${GOPATH}" \
    "${USERNAME}" \
    "false" \
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
# Builder
FROM base AS builder

WORKDIR /workspace
ADD . .
RUN go build -o server cmd/server/main.go

# -----------------------------
# Application
FROM debian:buster AS app

COPY --from=builder /workspace/server .
EXPOSE 8080

CMD ["./server"]