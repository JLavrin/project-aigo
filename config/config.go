package config

import "time"

type DatabaseType struct {
	Type     string
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
	TimeZone string
}

type ServerType struct {
	Port         string
	RunMode      string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var Server = ServerType{
	Port:         "8080",
	RunMode:      "debug",
	ReadTimeout:  60 * time.Second,
	WriteTimeout: 60 * time.Second,
}

var Database = DatabaseType{
	Type:     "postgres",
	Host:     "localhost",
	Port:     "5432",
	User:     "postgres",
	Password: "postgres",
	Name:     "aigo",
	SSLMode:  "disable",
	TimeZone: "Europe/Berlin",
}
