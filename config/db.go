package config

func init() {
	config.SetDefault("db.name", "mysql")
	config.SetDefault("db.address", "localhost")
	config.SetDefault("db.port", 3306)
	config.SetDefault("db.username", "root")
	config.SetDefault("db.collation", "utf8_general_ci")
	config.SetDefault("db.loc", "Local")
	config.SetDefault("db.parse_time", true)
	config.SetDefault("db.migrate.enabled", false)
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

// Name returns the name of db driver.
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

// Database returns the address of database.
func (d *db) Database() string {
	return config.GetString("db.database")
}

// Username returns the username of db.
func (d *db) Username() string {
	return config.GetString("db.username")
}

// password returns the password of db.
func (d *db) Password() string {
	return config.GetString("db.password")
}

// Collation returns the collation of db.
func (d *db) Collation() string {
	return config.GetString("db.collation")
}

// Loc returns the loc of db.
func (d *db) Loc() string {
	return config.GetString("db.loc")
}

// ParseTime returns the parse_time of db.
func (d *db) ParseTime() bool {
	return config.GetBool("db.parse_time")
}

// Enabled checks whether the migrate is enabled.
func (m *migrate) Enabled() bool {
	return config.GetBool("db.migrate.enabled")
}

// Source returns the source of migrate.
func (m *migrate) Source() string {
	return config.GetString("db.migrate.source")
}

// Table returns the table of migrate.
func (m *migrate) Table() string {
	return config.GetString("db.migrate.table")
}
