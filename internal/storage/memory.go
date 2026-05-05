package storage

import (
	"errors"
	"sync"

	"portfolio-go/internal/models"
)

type Storage struct {
	services      map[int]models.Service
	messages      map[int]models.Message
	nextMessageID int
	mu            sync.Mutex
}

func NewStorage() *Storage {
	s := &Storage{
		services:      make(map[int]models.Service),
		messages:      make(map[int]models.Message),
		nextMessageID: 1,
	}

	s.services[1] = models.Service{
		ID:          1,
		Title:       "Фирменный стиль",
		Image:       "Rectangle 15.png",
		Price:       "от 15 000 ₽",
		Duration:    "7–10 дней",
		Description: "Логотип, палитра, шрифты и базовый визуальный стиль бренда.",
	}

	s.services[2] = models.Service{
		ID:          2,
		Title:       "Лендинг",
		Image:       "Rectangle 15-1.png",
		Price:       "от 18 000 ₽",
		Duration:    "10–14 дней",
		Description: "Дизайн одностраничного сайта для презентации услуги, продукта или личного бренда.",
	}

	s.services[3] = models.Service{
		ID:          3,
		Title:       "UI/UX дизайн",
		Image:       "Rectangle 15-2.png",
		Price:       "от 20 000 ₽",
		Duration:    "14–21 день",
		Description: "Проектирование интерфейса, структуры экранов и визуальной системы продукта.",
	}

	s.services[4] = models.Service{
		ID:          4,
		Title:       "Оформление соцсетей",
		Image:       "Rectangle 15-3.png",
		Price:       "от 8 000 ₽",
		Duration:    "5–7 дней",
		Description: "Шаблоны постов, обложки, баннеры и единый визуальный стиль аккаунта.",
	}

	s.services[5] = models.Service{
		ID:          5,
		Title:       "Презентация",
		Image:       "Rectangle 15-4.png",
		Price:       "от 10 000 ₽",
		Duration:    "5–10 дней",
		Description: "Структура и дизайн презентации для защиты, продажи идеи или проекта.",
	}

	s.services[6] = models.Service{
		ID:          6,
		Title:       "Полиграфия",
		Image:       "Rectangle 15-5.png",
		Price:       "от 6 000 ₽",
		Duration:    "3–7 дней",
		Description: "Плакаты, листовки, визитки и другие печатные материалы.",
	}

	return s
}

func (s *Storage) GetServices() []models.Service {
	s.mu.Lock()
	defer s.mu.Unlock()

	services := make([]models.Service, 0, len(s.services))

	for i := 1; i <= len(s.services); i++ {
		service, ok := s.services[i]
		if ok {
			services = append(services, service)
		}
	}

	return services
}

func (s *Storage) GetServiceByID(id int) (models.Service, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	service, ok := s.services[id]
	if !ok {
		return models.Service{}, errors.New("service not found")
	}

	return service, nil
}

func (s *Storage) AddMessage(message models.Message) models.Message {
	s.mu.Lock()
	defer s.mu.Unlock()

	message.ID = s.nextMessageID
	s.messages[message.ID] = message
	s.nextMessageID++

	return message
}

func (s *Storage) GetMessages() []models.Message {
	s.mu.Lock()
	defer s.mu.Unlock()

	messages := make([]models.Message, 0, len(s.messages))

	for _, message := range s.messages {
		messages = append(messages, message)
	}

	return messages
}