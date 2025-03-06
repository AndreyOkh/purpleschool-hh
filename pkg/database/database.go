package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
	"hh/config"
)

func CreateDBPool(conf *config.DatabaseConfig, log *zerolog.Logger) *pgxpool.Pool {
	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", conf.Host, conf.Port, conf.User, conf.Password, conf.Database)
	dbPool, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		log.Fatal().Err(err).Msg("Ошибка создания подключения к БД")
	}
	err = dbPool.Ping(context.Background())
	if err != nil {
		log.Fatal().Err(err).Msg("Ошибка подключения к БД")
	}
	log.Info().Msg("Подключение к БД установлено")
	return dbPool
}
