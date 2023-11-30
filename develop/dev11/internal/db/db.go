package db

import (
	"dev11/internal/models"
)

type Storage interface {
	GetEvent(models.Event, string) (map[models.Date][]string, error)
	CreateEvent(models.Event) error
	UpdateEvent([] models.Event) error
	DeleteEvent(models.Event) error
}
