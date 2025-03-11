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

func (r *Repository) GetAll() ([]*Vacancy, error) {
	query := `SELECT * FROM vacancies`
	rows, err := r.DbPool.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	if vacancies, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByName[Vacancy]); err != nil {
		return nil, err
	} else {
		return vacancies, nil
	}
}

func (r *Repository) AddVacancy(form VacancyCreateForm) error {
	query := `INSERT INTO vacancies (email, role, company, salary, type, location, created_at) VALUES (@email, @role, @company, @salary, @type, @location, @created_at)`
	args := pgx.NamedArgs{
		"email":      form.Email,
		"role":       form.Role,
		"company":    form.Company,
		"salary":     form.Salary,
		"type":       form.Type,
		"location":   form.Location,
		"created_at": time.Now(),
	}
	_, err := r.DbPool.Exec(context.Background(), query, args)
	if err != nil {
		r.Logger.Error().Err(err).Msg("Ошибка записи в базу данных")
		return err
	}
	return nil
}
