package models

import "time"

type Attendee struct {
	ID         int       `boil:"id" json:"id" toml:"id" yaml:"id"`
	CustomerID int       `boil:"customer_id" json:"customer_id" toml:"customer_id" yaml:"customer_id"`
	TicketID   int       `boil:"ticket_id" json:"ticket_id" toml:"ticket_id" yaml:"ticket_id"`
	Email      string    `boil:"email" json:"email" toml:"email" yaml:"email"`
	Status     string    `boil:"status" json:"status" toml:"status" yaml:"status"`
	CreatedAt  time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt  time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
	DeletedAt  time.Time `boil:"deleted_at" json:"deleted_at" toml:"deleted_at" yaml:"deleted_at"`
	EventID    int       `boil:"event_id" json:"event_id" toml:"event_id" yaml:"event_id"`
	Customer   *Customer `json:"customer"`
	Ticket     *Ticket   `json:"ticket"`
	Event      *Event    `json:"event"`
}