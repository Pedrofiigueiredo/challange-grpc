package domain

import "time"

// Base contem dados padrao de uma entidade (id e datas)
type Base struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
