package service

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/oresdev/tbcc-wallet-api-v3/internal/server/util"
)

// DbGetUsers ...
func DbGetUsers(conn *pgxpool.Pool, ctx context.Context) (data []byte, err error) {

	if err := conn.QueryRow(ctx, `SELECT users_get_rows()`).Scan(&data); err != nil {
		return nil, err
	}

	return data, nil
}

// DbGetUserByID ...
func DbGetUserByID(id util.UUID, conn *pgxpool.Pool, ctx context.Context) (data []byte, err error) {

	if err := conn.QueryRow(ctx, `SELECT users_get_by_uuid($1)`, id).Scan(&data); err != nil {
		return nil, err
	}

	return data, nil
}

// DbGetUserExt ...
func DbGetUserExt(uuid string, conn *pgxpool.Pool, ctx context.Context) (data []byte, err error) {

	if err := conn.QueryRow(ctx, `SELECT users_get_extended_by_uuid($1)`, uuid).Scan(&data); err != nil {
		return nil, err
	}

	return data, nil
}

// DbUpdateUser ...
func DbUpdateUser(uuid string, address string, conn *pgxpool.Pool, ctx context.Context) (data []byte, err error) {

	if err := conn.QueryRow(ctx, `SELECT users_update_by_uuid($1, $2)`, uuid, address).Scan(&data); err != nil {
		return nil, err
	}

	return data, nil
}

// DbCreateUser ...
func DbCreateUser(useraddress []string, accounttype string, smartcard bool, conn *pgxpool.Pool, ctx context.Context) (data []byte, err error) {

	if err := conn.QueryRow(ctx, `SELECT users_create_row($1, $2, $3)`, useraddress, accounttype, smartcard).Scan(&data); err != nil {
		return nil, err
	}

	return data, nil
}

// DbMigrateUser ...
func DbMigrateUser(addresses []string, conn *pgxpool.Pool, ctx context.Context) (data []byte, err error) {

	if err := conn.QueryRow(ctx, `SELECT public.c_migrate($1)`, addresses).Scan(&data); err != nil {
		return nil, err
	}

	return data, nil
}

// DbUpdateVpnKey ...
func DbUpdateVpnKey(uuid string, txhash string, conn *pgxpool.Pool, ctx context.Context) (data []byte, err error) {

	if err := conn.QueryRow(ctx, `SELECT vpn_keys_update_by_uuid($1, $2)`, uuid, txhash).Scan(&data); err != nil {
		return nil, err
	}

	return data, nil
}
