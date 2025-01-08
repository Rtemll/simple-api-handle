package taskService

// "Дублируем" методы из репозитория, вызываем из сервиса методы репозитория для разделения бизнес логики и работы с базой данных
type TaskService struct {
	repo TaskRepository
}

func NewService(repo TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

//создаем CreateTask

func (s *TaskService) CreateTask(task Task) (Task, error) {
	return s.repo.CreateTask(task)
}

// создаем GetAllTasks
func (s *TaskService) GetAllTasks() ([]Task, error) {
	return s.repo.GetAllTasks()
}
func (s *TaskService) UpdateTask(id int, task Task) (Task, error) {
	uintID := uint(id)
	return s.repo.UpdateTaskByID(uintID, task)
}
func (s *TaskService) DeleteTask(id int) error {
	uintID := uint(id)
	return s.repo.DeleteTaskByID(uintID)
}
