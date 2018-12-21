package config

func init() {
	config.SetDefault("db.driver", "mysql")
	config.SetDefault("db.address", "localhost")
	config.SetDefault("db.port", 3306)
	config.SetDefault("db.username", "root")
}

type db struct{}

// DB returns the collection of the db config.
var DB = &db{}

// Driver returns the driver of db.
func (d *db) Driver() string {
	return config.GetString("db.driver")
}

// Address returns the address of db.
func (d *db) Address() string {
	return config.GetString("db.address")
}

// Port returns the port of db.
func (d *db) Port() int {
	return config.GetInt("db.port")
}

// DB returns the address of db.
func (d *db) Database() string {
	return config.GetString("db.database")
}

// DB returns the username of db.
func (d *db) Username() string {
	return config.GetString("db.username")
}

// DB returns the password of db.
func (d *db) Password() string {
	return config.GetString("db.password")
}
