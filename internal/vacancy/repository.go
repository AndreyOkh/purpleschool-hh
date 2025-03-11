package vacancy

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
	"time"
)

type Repository struct {
	DbPool *pgxpool.Pool
	Logger *zerolog.Logger
}

func NewRepository(db *pgxpool.Pool, logger *zerolog.Logger) *Repository {
	return &Repository{
		DbPool: db,
		Logger: logger,
	}
}

func (r *Repository) AddVacancy(form VacancyCreateForm) error {
	query := `INSERT INTO vacancies (email, role, company, salary, type, location, createdat) VALUES (@email, @role, @company, @salary, @type, @location, @createdat)`
	args := pgx.NamedArgs{
		"email":     form.Email,
		"role":      form.Role,
		"company":   form.Company,
		"salary":    form.Salary,
		"type":      form.Type,
		"location":  form.Location,
		"createdat": time.Now(),
	}
	_, err := r.DbPool.Exec(context.Background(), query, args)
	if err != nil {
		r.Logger.Error().Err(err).Msg("Ошибка записи в базу данных")
		return err
	}
	return nil
}
