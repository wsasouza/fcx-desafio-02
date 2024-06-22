package usecase

import (
	"fmt"
	"strings"

	"github.com/wsasouza/fcx-desafio-02/internal/domain"
)

type ReserveSpotUseCase struct {
	repo domain.EventRepository
}

type ReserveSpotInputDTO struct {
	Spots []string `json:"spots"`
}

func NewReserveSpotUseCase(repo domain.EventRepository) *ReserveSpotUseCase {
	return &ReserveSpotUseCase{repo: repo}
}

func contains(spotNames []string, name string) bool {
	for _, sn := range spotNames {
		if sn == name {
			return true
		}
	}
	return false
}

func (uc *ReserveSpotUseCase) Execute(eventID int, spotNames []string) error {
	spots := uc.repo.FindSpotsByEventID(eventID)

	var unavailable []string

	for _, spot := range spots {
		if spot.Status == domain.SpotStatusReserved {
			if contains(spotNames, spot.Name) {
				unavailable = append(unavailable, spot.Name)
			}
		}
	}

	if len(unavailable) > 0 {
		return fmt.Errorf("some spots are unavailable: %s", strings.Join(unavailable, ", "))
	}

	uc.repo.ReserveSpot(eventID, spotNames)

	return nil
}
