package repository

import (
	"github.com/74th/vscode-book-golang/model/tasks"
	"fmt"
)

// repository タスク	リポジトリの実装
type instance struct{
	tasks []tasks.Task
}

// New リポジトリの作成
func New() tasks.Repository{
	s := new(instance)
	s.tasks = make([]tasks.Task, 2, 20)
	s.tasks[0] = tasks.Task{
		ID:   1,
		Text: "task1",
		Done: false,
	}
	s.tasks[1] = tasks.Task{
		ID:   2,
		Text: "task2",
		Done: false,
	}
	return s
}

// Add タスクの追加
func (s *instance) Add(task tasks.Task) int {
	task.ID = len(s.tasks) + 1
	s.tasks = append(s.tasks, task)
	return task.ID
}

// List 未完了のタスクの一覧
func (s *instance) List() []*tasks.Task {
	result := []*tasks.Task{}
	for i, task := range s.tasks {
		if !task.Done {
			// taskは一時変数のインスタンスのため、
			// リポジトリのtasksのインスタンスを返す
			result = append(result, &s.tasks[i])
		}
	}
	return result
}

// Done タスクを完了にする
func (s *instance) Done(id int) error {
	for i, task := range s.tasks {
		if task.ID == id {
			s.tasks[i].Done = true
			return nil
		}
	}
	return fmt.Errorf("Not found id:%d", id)
}
