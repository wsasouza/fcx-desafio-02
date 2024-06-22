package usecase

import (
	"github.com/wsasouza/fcx-desafio-02/internal/domain"
)

type ListSpotsUseCase struct {
	repo domain.EventRepository
}

func NewListSpotsUseCase(repo domain.EventRepository) *ListSpotsUseCase {
	return &ListSpotsUseCase{repo: repo}
}

func (uc *ListSpotsUseCase) Execute(eventID int) []domain.Spot {
	return uc.repo.FindSpotsByEventID(eventID)
}
