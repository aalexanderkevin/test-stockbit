package repository

import (
	model "case2/models"
	"database/sql"
)

type Repository interface {
	Insert(model.Logger) error
}

type logRepository struct {
	DB *sql.DB
}

func NewLogRepository(db *sql.DB) Repository {
	return &logRepository{
		DB: db,
	}
}

func (l *logRepository) Insert(log model.Logger) error {
	_, err := l.DB.Exec("INSERT INTO log (timestamp, request, response) VALUES (?, ?, ?)",
		log.Timestamp,
		log.Request,
		log.Response,
	)
	if err != nil {
		return err
	}

	return nil
}
