package usecase

import (
	"github.com/wsasouza/fcx-desafio-02/internal/domain"
)

type GetEventUseCase struct {
	repo domain.EventRepository
}

func NewGetEventUseCase(repo domain.EventRepository) *GetEventUseCase {
	return &GetEventUseCase{repo: repo}
}

func (uc *GetEventUseCase) Execute(id int) (domain.Event, error) {
	return uc.repo.GetEventByID(id)
}
