package taskService

import "gorm.io/gorm"

type TaskRepository interface {
	// CreateTask - Передаем в функцию task типа Task из orm.go
	// возвращаем созданный Task и ошибку
	CreateTask(task Task) (Task, error)
	// GetAllTasks - Возвращаем массив из всех задач в БД и ошибку
	GetAllTasks() ([]Task, error)
	// UpdateTaskByID - Передаем id и Task, возвращаем обновленный Task
	// и ошибку
	UpdateTaskByID(id uint, task Task) (Task, error)
	// DeleteTaskByID - Передаем id для удаления, возвращаем только ошибку
	DeleteTaskByID(id uint) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *taskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) CreateTask(task Task) (Task, error) {
	result := r.db.Create(&task)
	if result.Error != nil {
		return Task{}, result.Error
	}
	return task, nil
}

func (r *taskRepository) GetAllTasks() ([]Task, error) {
	// Тут мы создаем массив данных и типом данных Task
	var tasks []Task
	err := r.db.Find(&tasks).Error //заполняем массив, получаем ошибку и массив
	return tasks, err
}

func (r *taskRepository) UpdateTaskByID(id uint, task Task) (Task, error) {
	//Ищем запись по id
	var existingTask Task
	result := r.db.First(&existingTask, id)
	if result.Error != nil {
		return Task{}, result.Error //Возвращаем ошибку, если запись не найдена
	}
	// Обновляем поля записи
	existingTask.Task = task.Task
	existingTask.IsDone = task.IsDone
	// Сохраняем изменения в базе данных
	result = r.db.Save(&existingTask)
	if result.Error != nil {
		return Task{}, result.Error // Возвращаем ошибку, если не получилось обновить
	}
	return existingTask, nil // Возвращаем обновленную задачу
}

func (r *taskRepository) DeleteTaskByID(id uint) error {
	// Ищем запись по id
	var task Task
	result := r.db.First(&task, id)
	if result.Error != nil {
		return result.Error // Возвращаем ошибку, если запись не найдена
	}
	// Удаляем найденную запись
	result = r.db.Delete(&task)
	return result.Error // Возвращаем ошибку, если не получилось удалить, иначе nil
}
