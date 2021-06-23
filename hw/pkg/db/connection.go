package db

import (
	"database/sql"
	"math"
	"strconv"
	"time"

	"github.com/karolhrdina/misc/hw/pkg/env"
	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/drivers/sqlboiler-psql/driver"
)

type User struct {
	Name     string
	Password string
}

// UserFromEnv returns user from <prefix>_NAME and <prefix>_PASS variables
func UserFromEnv(prefix string) *User {
	defaultUser := "postgres"
	defaultPass := ""
	name := env.GetVar(prefix+"_USER", defaultUser)
	pass := env.GetVar(prefix+"_PASS", defaultPass)
	return &User{Name: name, Password: pass}
}

// Connection encodes SQL connection data.
type Connection struct {
	User     User
	DBName   string
	Host     string
	Port     int
	SSLMode  string
	MaxConns int
}

// ConnectionFromEnv returns Connection from env variables
func ConnectionFromEnv(prefix string) (Connection, error) {
	pgPort := prefix + "_PORT"

	port, err := strconv.Atoi(env.GetVar(pgPort, "5432"))
	if err != nil {
		return Connection{}, errors.Errorf("error parsing %s as integer", pgPort)
	}

	pgMaxConns := prefix + "_MAX_CONNS"
	maxConns, err := strconv.Atoi(env.GetVar(pgMaxConns, "10"))
	if err != nil {
		return Connection{}, errors.Errorf("error parsing %s as integer", pgMaxConns)
	}
	c := Connection{
		User:     *UserFromEnv(prefix),
		DBName:   env.GetVar(prefix+"_NAME", ""),
		Host:     env.GetVar(prefix+"_HOST", "localhost"),
		Port:     port,
		SSLMode:  env.GetVar(prefix+"_SSL_MODE", "disable"),
		MaxConns: maxConns,
	}
	return c, nil
}

func (c *Connection) String() string {
	connstring := driver.PSQLBuildQueryString(
		c.User.Name,
		c.User.Password,
		c.DBName,
		c.Host,
		c.Port,
		c.SSLMode,
	)
	return connstring
}

// Connect tries to connect to postgres until it succeeds.
func (c *Connection) Connect() (*sql.DB, error) {
	timeout := 1 * time.Second
	var db *sql.DB
	var err error

	// This is my workaround around not being able to setup the docker
	// compose file properly so that `port-domain` would wait for postgres.
	// I tried a few things (waitscript) that I googled, but none of them
	// seemed to work :shrug:. This was much faster than more googling ;)
	for {
		db, err = sql.Open("postgres", c.String())
		if err != nil {
			return nil, err
		}
		// Check if DB is actually working.
		err = db.Ping()
		if err == nil {
			break
		}
		time.Sleep(timeout)
		timeout = timeout + (time.Duration(math.Log(float64(timeout))) * time.Second)
	}

	// Set maximum open connections to some sane default.
	// Use PG_MAX_CONNS to change it.
	db.SetMaxOpenConns(c.MaxConns)
	return db, nil
}
