// Filename: internal/models/ticket_info
package models

import (
	"context"
	"database/sql"
	"time"
)

type TicketInfo struct {
	TicketInfoID    int64
	ScheduleID      int64
	TicketPrice     int64 //Need to Ask Question on How to Reference
	NumberofTickets int64
}

type TicketModel struct {
	DB *sql.DB
}

// Code to access the database
func (m *TicketModel) Get() (*TicketInfo, error) {
	var t TicketInfo

	statement := `
				SELECT id, fname, lname, email, address, phone_number, password
				FROM ticket_info
				LIMIT 1
				`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := m.DB.QueryRowContext(ctx, statement).Scan(&t.TicketInfoID, &t.ScheduleID, &t.TicketPrice, &t.NumberofTickets)
	if err != nil {
		return nil, err
	}
	return &t, err
}
