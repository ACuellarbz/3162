// Filename: internal/models/mileage_of_trip
package models

import (
	"context"
	"database/sql"
	"time"
)

type MileageofTrip struct {
	MileageID      int64
	BeginningID       int64
	DestinationID       int64
	TotalMiles       int64
}

type MileageModel struct {
	DB *sql.DB
}

// Code to access the database
func (m *MileageModel) Get() (*MileageofTrip, error) {
	var o MileageofTrip

	
	statement := `
				SELECT id, beginning_location_id, destination_location_id, total_miles
				FROM mileage_of_trip
				ORDER BY RANDOM()
				LMIT 1
				`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := m.DB.QueryRowContext(ctx, statement).Scan(&o.MileageID, &o.BeginningID, &o.DestinationID, &o.TotalMiles)
	if err != nil {
		return nil, err
	}
	return &o, err
}
