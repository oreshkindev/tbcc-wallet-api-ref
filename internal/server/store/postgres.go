package store

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/oresdev/tbcc-wallet-api-v3/internal/server/conf"
	"github.com/sirupsen/logrus"
)

// СreateDB ...
func СreateDB(c *conf.Config) (*pgxpool.Pool, error) {

	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		c.Database.Host,
		c.Database.User,
		c.Database.Password,
		c.Database.Name,
		c.Database.Port,
	)
	conf, _ := pgxpool.ParseConfig(connStr)

	// Set max pool connections
	conf.MaxConns = c.Database.MaxConns

	conn, err := pgxpool.ConnectConfig(context.Background(), conf)
	if err != nil {
		logrus.Fatalf("opening connection: %v", err)
	}
	logrus.Info("Connected:")

	return conn, nil
}
