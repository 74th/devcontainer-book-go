package tasks

import "testing"

func TestNewRepository(t *testing.T) {
	rep := NewRepository()
	if len(rep.tasks) != 2 {
		t.Errorf("unexpected size %d", len(rep.tasks))
	}
}
func TestAdd(t *testing.T) {
}
