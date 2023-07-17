package postgres

type Config struct {
	Database string `config:"DATABASE_NAME"`
	User     string `config:"DATABASE_USERNAME"`
	Password string `config:"DATABASE_PASSWORD"`
	Host     string `config:"DATABASE_HOST"`
	Port     int    `config:"DATABASE_PORT"`
}
