package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/wsasouza/fcx-desafio-02/internal/usecase"
)

type EventsHandler struct {
	listEventsUseCase  *usecase.ListEventsUseCase
	getEventUseCase    *usecase.GetEventUseCase
	listSpotsUseCase   *usecase.ListSpotsUseCase
	reserveSpotUseCase *usecase.ReserveSpotUseCase
}

func NewEventsHandler(
	listEventsUseCase *usecase.ListEventsUseCase,
	getEventUseCase *usecase.GetEventUseCase,
	listSpotsUseCase *usecase.ListSpotsUseCase,
	reserveSpotUseCase *usecase.ReserveSpotUseCase,
) *EventsHandler {
	return &EventsHandler{
		listEventsUseCase:  listEventsUseCase,
		getEventUseCase:    getEventUseCase,
		listSpotsUseCase:   listSpotsUseCase,
		reserveSpotUseCase: reserveSpotUseCase,
	}
}

func (handler *EventsHandler) ListEvents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(handler.listEventsUseCase.Execute())
}

func (handler *EventsHandler) GetEvent(w http.ResponseWriter, r *http.Request) {
	eventID, err := strconv.Atoi(r.PathValue("eventID"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	event, err := handler.getEventUseCase.Execute(eventID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(event)
}

func (handler *EventsHandler) ListSpotEvent(w http.ResponseWriter, r *http.Request) {
	eventID, err := strconv.Atoi(r.PathValue("eventID"))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(handler.listSpotsUseCase.Execute(eventID))
}

func (handler *EventsHandler) ReserveSpot(w http.ResponseWriter, r *http.Request) {
	eventID, err := strconv.Atoi(r.PathValue("eventID"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	var input usecase.ReserveSpotInputDTO
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := handler.reserveSpotUseCase.Execute(eventID, input.Spots); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
