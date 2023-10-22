package pg

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/evgenivanovi/gpl/std"
	"github.com/evgenivanovi/gpl/std/conv"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	_ "github.com/spf13/viper/remote"
)

/* __________________________________________________ */

//goland:noinspection GoNameStartsWithPackageName
const (
	PGHostVar     string = "PGHOST"
	PGPortVar     string = "PGPORT"
	PGDatabaseVar string = "PGDATABASE"
	PGUserVar     string = "PGUSER"
	PGPasswordVar string = "PGPASSWORD"
	PGSSLModeVar  string = "PGSSLMODE"
)

//goland:noinspection GoNameStartsWithPackageName
const (
	PGDefaultHost     string = "postgres"
	PGDefaultPort     uint16 = 5432
	PGDefaultDatabase string = "postgres"
	PGDefaultUser     string = "postgres"
	PGDefaultPassword string = "password"
	PGDefaultSSLMode  string = "disable"
)

//goland:noinspection GoNameStartsWithPackageName
const (
	PGCDefaultTimeout = time.Second * 5
)

/* __________________________________________________ */

func init() {
	if err := initDefaultEnv(); err != nil {
		panic(err)
	}
}

func initDefaultEnv() error {

	if len(os.Getenv(PGHostVar)) == 0 {
		if err := os.Setenv(PGHostVar, PGDefaultHost); err != nil {
			return errors.WithStack(err)
		}
	}

	if len(os.Getenv(PGPortVar)) == 0 {
		port := conv.Uint16ToString(PGDefaultPort)
		if err := os.Setenv(PGPortVar, port); err != nil {
			return errors.WithStack(err)
		}
	}

	if len(os.Getenv(PGDatabaseVar)) == 0 {
		if err := os.Setenv(PGDatabaseVar, PGDefaultDatabase); err != nil {
			return errors.WithStack(err)
		}
	}

	if len(os.Getenv(PGUserVar)) == 0 {
		if err := os.Setenv(PGUserVar, PGDefaultUser); err != nil {
			return errors.WithStack(err)
		}
	}

	if len(os.Getenv(PGPasswordVar)) == 0 {
		if err := os.Setenv(PGPasswordVar, PGDefaultPassword); err != nil {
			return errors.WithStack(err)
		}
	}

	if len(os.Getenv(PGSSLModeVar)) == 0 {
		if err := os.Setenv(PGSSLModeVar, PGDefaultSSLMode); err != nil {
			return errors.WithStack(err)
		}
	}

	return nil

}

/* __________________________________________________ */

type ConnectionSettingOp func(*ConnectionSettings)

func (o ConnectionSettingOp) Join(op ConnectionSettingOp) ConnectionSettingOp {
	return func(opts *ConnectionSettings) {
		o(opts)
		op(opts)
	}
}

func (o ConnectionSettingOp) And(ops ...ConnectionSettingOp) ConnectionSettingOp {
	return func(opts *ConnectionSettings) {
		o(opts)
		for _, fn := range ops {
			fn(opts)
		}
	}
}

func WithTimeout(duration time.Duration) ConnectionSettingOp {
	return func(opts *ConnectionSettings) {
		opts.Timeout = duration
	}
}

/* __________________________________________________ */

type ConnectionSettings struct {
	Timeout time.Duration
}

func NewConnectionSettings(opts ...ConnectionSettingOp) *ConnectionSettings {
	settings := connectionSettings()
	for _, op := range opts {
		op(settings)
	}
	return settings
}

func connectionSettings() *ConnectionSettings {
	return &ConnectionSettings{
		Timeout: PGCDefaultTimeout,
	}
}

/* __________________________________________________ */

type Settings struct {
	Host     string
	Port     uint16
	Database string
	User     string
	Password string
	SSLMode  string
}

func NewSettings() *Settings {
	return &Settings{
		Host:     PGDefaultHost,
		Port:     PGDefaultPort,
		Database: PGDefaultDatabase,
		User:     PGDefaultUser,
		Password: PGDefaultPassword,
		SSLMode:  PGDefaultSSLMode,
	}
}

func (s Settings) toDSN() string {

	var args []string

	if len(s.Host) > 0 {
		args = append(args, fmt.Sprintf("host=%s", s.Host))
	}

	if s.Port > 0 {
		args = append(args, fmt.Sprintf("port=%d", s.Port))
	}

	if len(s.Database) > 0 {
		args = append(args, fmt.Sprintf("dbname=%s", s.Database))
	}

	if len(s.User) > 0 {
		args = append(args, fmt.Sprintf("user=%s", s.User))
	}

	if len(s.Password) > 0 {
		args = append(args, fmt.Sprintf("password=%s", s.Password))
	}

	if len(s.SSLMode) > 0 {
		args = append(args, fmt.Sprintf("sslmode=%s", s.SSLMode))
	}

	return strings.Join(args, std.Space)

}

/* __________________________________________________ */

type Datasource struct {
	Pool *pgxpool.Pool
	DB   *sql.DB
}

func NewDatasource(
	ctx context.Context,
	dsn string,
	connectionSettings ConnectionSettings,
) (*Datasource, error) {

	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(ctx, connectionSettings.Timeout)
	defer cancel()

	if err = pool.Ping(ctx); err != nil {
		return nil, err
	}

	db, err := sql.Open(
		"pgx",
		stdlib.RegisterConnConfig(pool.Config().ConnConfig),
	)
	if err != nil {
		return nil, err
	}

	return &Datasource{
		Pool: pool,
		DB:   db,
	}, nil

}

/* __________________________________________________ */
