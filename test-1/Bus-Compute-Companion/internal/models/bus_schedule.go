// Filename: internal/models/bus_schedule.go
package models

import (
	"context"
	"database/sql"
	"fmt"
	"log"
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

func (m *BusScheduleModel) Insert(schedule_id string, company string, beginning string, destination string) (int64, error) {
	var id int64

	statement :=
		`
			INSERT INTO bus_schedule(id, company_id, beginning_location_id, destination_location_id)
			VALUES($1, $2, $3, $4)
			RETURNING ID
		`
	//timeout for DB connection
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := m.DB.QueryRowContext(ctx, statement, schedule_id, company, beginning, destination).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (m *BusScheduleModel) Delete(schedule_id string) error {
	//Statement that will delete record
	statement := `
	DELETE FROM bus_schedule
	WHERE id = $1
	`
	//timeout for DB connection
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	_, err := m.DB.ExecContext(ctx, statement, schedule_id)

	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

// Code to access the database
func (m *BusScheduleModel) Get() (*BusSchedule, error) {
	var b BusSchedule

	statement := `
				SELECT id, company_id, beginning_location_id, destination_location_id
				FROM bus_schedule
				`
	//timeout for DB connection
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := m.DB.QueryRowContext(ctx, statement).Scan(&b.ScheduleID, &b.CompanyID, &b.BeginningID, &b.DestinationID)
	if err != nil {
		return nil, err
	}
	return &b, nil
}

// Credit: Interpretation of Hipolito's Read Function
func (m *BusScheduleModel) SearchRecord(schedule_id string) ([]*BusSchedule, error) {
	//SQL statement
	statement :=
		`
			SELECT id, company_id, beginning_location_id, destination_location_id
			FROM bus_schedule
			WHERE id = $1
		`
	data, err := m.DB.Query(statement, schedule_id)
	if err != nil {
		return nil, err
	}
	defer data.Close()
	log.Println(data)

	routes := []*BusSchedule{}

	data.Next()

	route := &BusSchedule{}
	err = data.Scan(&route.ScheduleID, &route.CompanyID, &route.BeginningID, &route.DestinationID)

	routes = append(routes, route)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	if err = data.Err(); err != nil {
		return nil, err
	}
	return routes, nil
}
