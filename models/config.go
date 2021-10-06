package models

type ServerConfig struct {
	Name        string `env:"DB_NAME"`
	Port        string `env:"DB_PORT"`
	Host        string `env:"DB_HOST"`
	Username    string `env:"DB_USERNAME"`
	Password    string `env:"DB_PASSWORD"`
	ServiceHost string `env:"SERVICE_HOST"`
	ServiceName string `env:"SERVICE_NAME"`
	ServicePort string `env:"PORT"`
	UrlJsonUser string `env:"URL_JSON_USER"`
}
