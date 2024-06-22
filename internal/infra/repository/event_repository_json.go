package repository

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/wsasouza/fcx-desafio-02/internal/domain"
)

type Data struct {
	Events []domain.Event `json:"events"`
	Spots  []domain.Spot  `json:"spots"`
}

type JsonEventRepository struct {
	data Data
}

func (er JsonEventRepository) ListEvents() []domain.Event {
	return er.data.Events
}

func (er JsonEventRepository) GetEventByID(id int) (domain.Event, error) {
	for _, event := range er.data.Events {
		if event.ID == id {
			return event, nil
		}
	}
	return domain.Event{}, errors.New("cant find event with this id")
}

func (er JsonEventRepository) FindSpotsByEventID(eventID int) []domain.Spot {
	var spots []domain.Spot
	for _, spot := range er.data.Spots {
		if spot.EventID == eventID {
			spots = append(spots, spot)
		}
	}

	return spots
}

func contains(array []string, str string) bool {
	for _, item := range array {
		if item == str {
			return true
		}
	}
	return false
}

func (er JsonEventRepository) ReserveSpot(eventID int, spotNames []string) {
	for i, spot := range er.data.Spots {
		if spot.EventID == eventID && contains(spotNames, spot.Name) {
			er.data.Spots[i].Status = domain.SpotStatusReserved
		}
	}
}

func NewJsonEventRepository(filePath string) (domain.EventRepository, error) {
	repo := JsonEventRepository{}

	data, err := os.ReadFile(filePath)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &repo.data)

	if err != nil {
		return nil, err
	}

	return repo, nil
}
