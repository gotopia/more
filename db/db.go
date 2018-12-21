package db

import (
	"database/sql"
	"fmt"

	"github.com/pkg/errors"

	// import MySQL driver.
	_ "github.com/go-sql-driver/mysql"
	"github.com/gotopia/more/config"
	"github.com/volatiletech/sqlboiler/boil"
)

var database *sql.DB

func init() {
	driver := config.DB.Driver()
	if driver != "mysql" {
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
	db, err := sql.Open(driver, dataSource)
	if db == nil || err != nil {
		err = errors.Wrap(err, "failed to open database")
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		err = errors.Wrap(err, "failed to connect database")
		panic(err)
	}
	database = db
	boil.DebugMode = config.Development()
}

// DB returns a global database.
func DB() *sql.DB {
	return database
}
