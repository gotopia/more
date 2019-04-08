package config

func init() {
	config.SetDefault("db.name", "mysql")
	config.SetDefault("db.address", "localhost")
	config.SetDefault("db.port", 3306)
	config.SetDefault("db.username", "root")
	config.SetDefault("db.migrate.enable", true)
	config.SetDefault("db.migrate.source", "file://db/migrate")
	config.SetDefault("db.migrate.table", "schema_migrations")
}

type migrate struct {
}

type db struct {
	Migrate migrate
}

// DB returns the collection of the db config.
var DB = &db{}

// Driver returns the driver of db.
func (d *db) Name() string {
	return config.GetString("db.name")
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

func (m *migrate) Enable() bool {
	return config.GetBool("db.migrate.enable")
}

func (m *migrate) Source() string {
	return config.GetString("db.migrate.source")
}

func (m *migrate) Table() string {
	return config.GetString("db.migrate.table")
}
