package goal

type TaskStatus string

const (
	StatusTodo       TaskStatus = "todo"
	StatusInProgress TaskStatus = "in_progress"
	StatusDone       TaskStatus = "done"
)

type Task struct {
	Title  string
	Status TaskStatus
}

type Goal struct {
	Title string
	Tasks []Task
}

func New(title string, tasks []string) Goal {
	goalTasks := make([]Task, 0, len(tasks))
	for _, task := range tasks {
		goalTasks = append(goalTasks, Task{Title: task, Status: StatusTodo})
	}

	return Goal{Title: title, Tasks: goalTasks}
}
