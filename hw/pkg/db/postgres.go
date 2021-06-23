package db

import (
    "database/sql"

    // Postgres driver.
    _ "github.com/jackc/pgx/v4/stdlib"
)

// ProvidePG returns sql.DB connection pool for Postgres from env vars
//  * PG_PORT
//  * PG_MAX_CONNS
//  * PG_USER
//  * PG_PASS
//  * PG_NAME
//  * PG_HOST
//  * PG_SSL_MODE.
func ProvidePG() (*sql.DB, error) {
    conn, err := ConnectionFromEnv("PG")
    if err != nil {
        return nil, err
    }
    return conn.Connect()
}
