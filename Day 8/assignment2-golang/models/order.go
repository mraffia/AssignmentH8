package models

import "time"

type Order struct {
	ID           uint   `gorm:"primaryKey"`
	CustomerName string `gorm:"not null;unique;type:varchar(191)"`
	OrderedAt    time.Time
	Items        []Item
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
