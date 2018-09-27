package model

import "time"

type Model struct {
	ID        uint64     `form:"id" gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `binding:"-" form:"-" json:"created_at"`
	UpdatedAt time.Time  `binding:"-" form:"-" json:"updated_at"`
	DeletedAt *time.Time `binding:"-" form:"-" json:"-"`
}
