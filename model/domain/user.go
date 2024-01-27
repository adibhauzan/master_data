package domain

import "time"

type User struct {
	ID        int       `gorm:"column:id;primaryKey"`
	Name      string    `gorm:"column:name;not null;size:256"`
	Email     string    `gorm:"column:email;not null;unique"`
	Password  string    `gorm:"column:password;not null"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}
