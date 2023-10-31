package memorystore

import "app/entity"

type Task struct {
	Tasks []entity.Task
}

func NewTaskStore() *Task {
	return &Task{
		Tasks: make([]entity.Task, 0),
	}
}
func (t *Task) CreateNewTask(task entity.Task) (entity.Task, error) {
	task.ID = len(t.Tasks) + 1
	t.Tasks = append(t.Tasks, task)
	return task, nil
}
func (t *Task) ListUserTasks(userID int) ([]entity.Task, error) {
	var userTasks []entity.Task
	for _, task := range t.Tasks {
		if task.UserID == userID {
			userTasks = append(userTasks, task)
		}
	}
	return userTasks, nil
}
