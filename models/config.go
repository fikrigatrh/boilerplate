package models

import "gorm.io/gorm"

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
	UsernameSU  string `env:"USERNAME_SU"`
	PassSu      string `env:"PASSWORD_SU"`
	JSONPathFile   string `env:"JSON_PATHFILE,required"`
	PostgresConfig PostgresConfig
}

type Auth struct {
	gorm.Model
	AuthUUID string `gorm:"size:255;not null;" json:"auth_uuid"`
	Username string `gorm:"size:15;not null;" json:"username"`
	RoleId   string `json:"role_id"`
}

type TokenStruct struct {
	Token string `json:"token"`
}

type PostgresConfig struct {
	Name     string `env:"NAME_POSTGRES,required"`
	Host     string `env:"HOST_POSTGRES,required"`
	Port     string `env:"PORT_POSTGRES,required"`
	User     string `env:"USER_POSTGRES"`
	Password string `env:"PASS_POSTGRES"`
}