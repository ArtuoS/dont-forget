package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Item struct {
	Guid        uuid.UUID
	Name        string
	Description string
	StartDate   time.Time
	EndDate     time.Time
}

func (i *Item) IsValid() error {
	if i.Name == "" {
		return errors.New("name is required")
	}
	if i.Description == "" {
		return errors.New("description is required")
	}
	return nil
}

func (i *Item) GenerateGuid() {
	i.Guid = uuid.New()
}
