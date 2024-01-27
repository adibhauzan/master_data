package domain

import "time"

type Department struct {
	ID        int       `gorm:"column:id;primaryKey"`
	Name      string    `gorm:"column:name;not null;size:256"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}
