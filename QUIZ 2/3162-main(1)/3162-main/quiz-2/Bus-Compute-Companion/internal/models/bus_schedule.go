// Filename: internal/models/bus_schedule.go
package models

import (
	"context"
	"database/sql"
	"time"
)

type BusSchedule struct {
	ScheduleID       int64
	CompanyID        int64
	BeginningID      int64
	DestinationID    int64
	BusDepartureTime time.Time
	BusArrivalTime   time.Time
}

type BusScheduleModel struct {
	DB *sql.DB
}

// Code to access the database
func (m *BusScheduleModel) Get() (*BusSchedule, error) {
	var b BusSchedule

	statement := `
				SELECT id, company_id, beginning_location_id, destination_location_id, bus_departutre_time, bus_arrival_time
				FROM bus_schedule
				LIMIT 1
				`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := m.DB.QueryRowContext(ctx, statement).Scan(&b.ScheduleID, &b.CompanyID, &b.BeginningID, &b.DestinationID, &b.BusArrivalTime, b.BusDepartureTime)
	if err != nil {
		return nil, err
	}
	return &b, err
}
