// Filename: internal/models/bus_schedule.go
package models

import (
	"context"
	"database/sql"
	"time"
)

type BusSchedule struct {
	ScheduleID    int64
	CompanyID     int64
	BeginningID   int64
	DestinationID int64
}

type BusScheduleModel struct {
	DB *sql.DB
}

// Code to access the database
func (m *BusScheduleModel) Get() (*BusSchedule, error) {
	var b BusSchedule

	statement := `
				SELECT id, company_id, beginning_location_id, destination_location_id
				FROM bus_schedule
				`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := m.DB.QueryRowContext(ctx, statement).Scan(&b.ScheduleID, &b.CompanyID, &b.BeginningID, &b.DestinationID)
	if err != nil {
		return nil, err
	}
	return &b, nil
}
 func (m *BusScheduleModel) Insert(schedule_id string,company string, beginning string, destination string) (int64, error){
	var id int64

	statement := 
		`
			INSERT INTO bus_schedule(id, company_id, beginning_location_id, destination_location_id)
			VALUES($1, $2, $3, $4)
			RETURNING ID
		`
	ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
	defer cancel()
	err := m.DB.QueryRowContext(ctx, statement, schedule_id, company, beginning, destination).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id,nil
}