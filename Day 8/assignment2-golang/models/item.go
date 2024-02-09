package models

import "time"

type Item struct {
	ID          uint   `gorm:"primaryKey"`
	ItemCode    string `gorm:"not null;unique;type:varchar(191)"`
	Description string `gorm:"varchar(191)"`
	Quantity    int    `gorm:"integer(11)"`
	OrderID     uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
