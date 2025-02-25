package goose

import (
	"log"

	_ "github.com/lib/pq"
	"github.com/pressly/goose"
)

const UpCommand = "up"
const DownCommand = "down"

func MigrateUp(dir string, driver string, dsn string, args ...string) {
	Migrate(dir, driver, dsn, UpCommand, args...)
}

func MigrateDown(dir string, driver string, dsn string, args ...string) {
	Migrate(dir, driver, dsn, DownCommand, args...)
}

func Migrate(dir string, driver string, dsn string, command string, args ...string) {

	db, err := goose.OpenDBWithDriver(driver, dsn)
	if err != nil {
		log.Fatalf("goose: failed to open DB: %v\n", err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("goose: failed to close DB: %v\n", err)
		}
	}()

	if gooseError := goose.Run(command, db, dir, args...); gooseError != nil {
		log.Fatalf("goose %v: %v", command, gooseError)
	}

}
