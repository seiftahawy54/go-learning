package models

import "time"

type Events struct {
	title       string
	description string
	eventDate   time.Time
	users []string
}