// Filename: cmd/web/data.go
package main

import "github.com/ACuellarbz/3162/internal/models"

type templateData struct {
	Schedule *models.BusSchedule
	Flash    string
}
