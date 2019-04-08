package db

import (
	"database/sql"
	"fmt"

	"github.com/pkg/errors"

	// import MySQL driver.
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"

	// import migrate source file.
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/gotopia/more/config"
	"github.com/volatiletech/sqlboiler/boil"
)

var database *sql.DB

func init() {
	name := config.DB.Name()
	if name != "mysql" {
		err := errors.New("only mysql is supported now")
		panic(err.Error())
	}
	dataSource := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true",
		config.DB.Username(),
		config.DB.Password(),
		config.DB.Address(),
		config.DB.Port(),
		config.DB.Database(),
	)
	db, err := sql.Open(name, dataSource)
	if db == nil || err != nil {
		err = errors.Wrap(err, "failed to open database")
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		err = errors.Wrap(err, "failed to connect database")
		panic(err)
	}
	if config.DB.Migrate.Enable() {
		driver, err := mysql.WithInstance(db, &mysql.Config{
			MigrationsTable: config.DB.Migrate.Table(),
		})
		if err != nil {
			err = errors.Wrap(err, "failed to get db driver")
			panic(err)
		}
		m, err := migrate.NewWithDatabaseInstance(config.DB.Migrate.Source(), name, driver)
		if err != nil {
			err = errors.Wrap(err, "failed to init migrate instance")
			panic(err)
		}
		if err := m.Up(); err != nil {
			err = errors.Wrap(err, "failed to migrate")
			panic(err)
		}
	}
	database = db
	boil.DebugMode = config.Development()
}

// DB returns a global database.
func DB() *sql.DB {
	return database
}
