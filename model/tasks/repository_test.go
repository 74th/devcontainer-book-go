package tasks

import (
	"testing"
)

func TestNewRepository(t *testing.T) {
	rep := NewRepository()
	if len(rep.(*repository).tasks) != 2 {
		t.Errorf("初期化時点で2つのタスクが格納されていること %d", len(rep.(*repository).tasks))
	}
}
func TestAdd(t *testing.T) {
	rep := NewRepository()
	rep.Add(Task{
		Text: "new task",
	})

	if len(rep.(*repository).tasks) != 3 {
		t.Errorf("タスクが追加されていること %d", len(rep.(*repository).tasks))
	} else {
		addedTask := rep.(*repository).tasks[2]

		if addedTask.Text != "new task" {
			t.Errorf("追加したタスクが末尾に追加されていること %s", addedTask.Text)
		}

		if addedTask.ID <= 2 {
			t.Errorf("タスクに新しいIDが振られること %d", addedTask.ID)
		}

		for i, task := range rep.(*repository).tasks {
			if i != 2 {
				if addedTask.ID == task.ID {
					t.Errorf("既存のタスクとは異なるIDが振られていること %d == %d", addedTask.ID, task.ID)
				}
				if addedTask.Text == task.Text {
					t.Errorf("既存のタスクとを上書きしていないこと %s == %s", addedTask.Text, task.Text)
				}
			}
		}

	}
}

func TestDone(t *testing.T) {
	rep := NewRepository()
	rep.Add(Task{
		Text: "3rd task",
	})

	l := rep.List()
	if len(l) != 3 {
		t.Errorf("完了にする前のタスクの数のチェック")
	}

	doneID := rep.(*repository).tasks[1].ID
	rep.Done(doneID)

	l = rep.List()
	if len(l) != 2 {
		t.Errorf("完了したタスクは取り除かれていること")
	}
	for _, task := range l {
		if task.ID == doneID {
			t.Errorf("完了したタスクは取り除かれていること %d", task.ID)
		}
	}
}
