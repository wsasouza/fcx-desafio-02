package domain

type SpotStatus string

const (
	SpotStatusAvailable SpotStatus = "available"
	SpotStatusReserved  SpotStatus = "reserved"
)

type Spot struct {
	ID      int        `json:"id"`
	Name    string     `json:"name"`
	Status  SpotStatus `json:"status"`
	EventID int        `json:"event_id"`
}
