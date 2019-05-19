package tasks

// Repository タスクリポジトリ
type Repository interface {
	Add(Task) int
	List() []*Task
	Done(id int) error
}
