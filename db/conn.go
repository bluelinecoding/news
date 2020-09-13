package db

import (
	"context"
	"database/sql"
	"errors"
	"os"
	"strings"
	"sync"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	log "github.com/sirupsen/logrus"
)

var (
	defaultConnString = "host=localhost user=postgres dbname=postgres sslmode=disable password=password"
	database          *gorm.DB
	lock              = &sync.Mutex{}
)

func init() {
	_, err := GetDB(context.Background())
	if err != nil {
		log.Error(err)
	}
}

func connect(connString string) (*gorm.DB, error) {
	db, err := gorm.Open("postgres", connString)
	if err != nil {
		if db.DB() != nil {
			db.DB().Close()
		}

		return nil, err
	}

	if os.Getenv("DEBUG_SQL") == "true" {
		db.LogMode(true)
	}

	err = migrateDB(db.DB())
	if err != nil && err != migrate.ErrNoChange {
		if db.DB() != nil {
			db.Close()
		}

		return nil, err
	}

	db.SetLogger(log.StandardLogger())

	return db, nil
}

func GetDB(ctx context.Context) (*gorm.DB, error) {
	lock.Lock()
	defer lock.Unlock()

	if database != nil {
		return database, nil
	}

	db, err := connect(defaultConnString)
	if err != nil {
		return nil, err
	}

	database = db

	return database, err
}

func ResetDB() error {
	db, err := GetDB(context.Background())
	if err != nil {
		return err
	}

	var tables []string
	err = db.Table("pg_tables").
		Where("schemaname = 'public' and tablename != 'schema_migrations'").
		Pluck("tablename", &tables).Error
	if err != nil {
		return err
	}

	err = db.Exec("TRUNCATE TABLE " + strings.Join(tables, ",") + " CASCADE").Error
	if err != nil {
		return err
	}

	return nil
}

func migrateDB(db *sql.DB) error {
	migration, err := migrator(db)
	if err != nil {
		return err
	}

	return migration.Up()
}

func migrator(db *sql.DB) (*migrate.Migrate, error) {
	dir := os.Getenv("MIGRATIONS_DIR")
	if dir == "" {
		return nil, errors.New("MIGRATIONS_DIR env var is not set")
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return nil, err
	}

	return migrate.NewWithDatabaseInstance("file://"+dir, "postgres", driver)
}
