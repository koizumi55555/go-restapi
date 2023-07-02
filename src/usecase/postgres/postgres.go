package postgres

import (
	"gorm.io/gorm"
)

type Postgres struct {
	Conn *gorm.DB
}

func New(conn *gorm.DB) Postgres {
	return Postgres{
		Conn: conn,
	}
}
