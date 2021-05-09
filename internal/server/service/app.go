package service

import (
	"context"
	"encoding/json"

	"github.com/jackc/pgx/v4/pgxpool"
)

// DbGetUpdates ...
func DbGetUpdates(conn *pgxpool.Pool, ctx context.Context) (data []byte, err error) {

	if err := conn.QueryRow(ctx, `SELECT app_update_get_rows()`).Scan(&data); err != nil {
		return nil, err
	}

	return data, nil
}

// DbCreateUpdate ...
func DbCreateUpdate(v int, u string, f bool, cs string, cl string, conn *pgxpool.Pool, ctx context.Context) (data []byte, err error) {

	if err := conn.QueryRow(ctx, `SELECT app_update_create_row($1, $2, $3, $4, $5)`, v, u, f, cs, cl).Scan(&data); err != nil {
		return nil, err
	}

	return data, nil
}

// DbGetConfig ...
func DbGetConfig(conn *pgxpool.Pool, ctx context.Context) (data []byte, err error) {

	if err := conn.QueryRow(ctx, `SELECT app_config_get_rows()`).Scan(&data); err != nil {
		return nil, err
	}

	return data, nil
}

// DbCreateConfig ...
func DbCreateConfig(cg string, v json.RawMessage, conn *pgxpool.Pool, ctx context.Context) (data []byte, err error) {

	if err := conn.QueryRow(ctx, `SELECT app_config_create_row($1, $2)`, cg, v).Scan(&data); err != nil {
		return nil, err
	}

	return data, nil
}

// DbCountVersion ...
func DbCountVersion(v int, conn *pgxpool.Pool, ctx context.Context) {

	conn.QueryRow(ctx, `SELECT app_counter_update_row($1)`, v)

}
