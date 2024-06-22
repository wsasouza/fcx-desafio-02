package usecase

import (
	"github.com/wsasouza/fcx-desafio-02/internal/domain"
)

type ListEventsUseCase struct {
	repo domain.EventRepository
}

func NewListEventsUseCase(repo domain.EventRepository) *ListEventsUseCase {
	return &ListEventsUseCase{repo: repo}
}

func (uc *ListEventsUseCase) Execute() []domain.Event {
	return uc.repo.ListEvents()
}
