package domain

type EventRepository interface {
	ListEvents() []Event
	GetEventByID(id int) (Event, error)
	FindSpotsByEventID(eventID int) []Spot
	ReserveSpot(eventID int, spotNames []string)
}
