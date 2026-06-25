package category

import "time"

type Category struct {
	Id        uint32
	Name      string
	Slug      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
