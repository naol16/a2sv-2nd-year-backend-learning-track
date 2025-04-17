package data
import(
	"time"
     "errors"
	"taskmanager/model"
)

var tasks = []model.Task{
	{ID: "1", Title: "Task 1", Description: "First task", DueDate: time.Now(), Status: "Pending"},
	{ID: "2", Title: "Task 2", Description: "Second task", DueDate: time.Now().AddDate(0, 0, 1), Status: "In Progress"},
	{ID: "3", Title: "Task 3", Description: "Third task", DueDate: time.Now().AddDate(0, 0, 2), Status: "Completed"},
}


func GetAllTasks() []model.Task {
	return tasks
}

func GetTaskByID(id string) (*model.Task, error) {
	for _, t := range tasks {
		if t.ID == id {
			return &t, nil
		}
	}
	return nil, errors.New("task not found")
}

func CreateTask(task model.Task) model.Task {
	tasks = append(tasks, task)
	return task
}

func UpdateTask(id string, updated model.Task) (*model.Task, error) {
	for i, t := range tasks {
		if t.ID == id {
			updated.ID = id
			tasks[i] = updated
			return &updated, nil
		}
	}
	return nil, errors.New("task not found")
}

func DeleteTask(id string) error {
	for i, t := range tasks {
		if t.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return nil
		}
	}
	return errors.New("task not found")
}

