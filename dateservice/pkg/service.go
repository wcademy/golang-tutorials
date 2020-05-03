package pkg

import (
	"context"
	"time"
)

type Service interface {
	Status(ctx context.Context) (string, error)
	Get(ctx context.Context) (string, error)
	Validate(ctx context.Context, date string) (bool, error)
}

type dateService struct{}

func NewService() Service {
	return dateService{}
}

func (dateService) Status(ctx context.Context) (string, error) {
	return "ok", nil
}

func (dateService) Get(ctx context.Context) (string, error) {
	now := time.Now()
	return now.Format("02.01.2006"), nil
}

func (dateService) Validate(ctx context.Context, date string) (bool, error) {
	_, err := time.Parse("02.01.2006", date)
	if err != nil {
		return false, err
	}

	return true, nil
}
