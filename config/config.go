package config

import "time"

type DatabaseType struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

type ServerType struct {
	Port         string
	RunMode      string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var Server = ServerType{
	Port:         ":8080",
	RunMode:      "debug",
	ReadTimeout:  60 * time.Second,
	WriteTimeout: 60 * time.Second,
}

var Database = DatabaseType{
	Host:     "localhost",
	Port:     "5432",
	User:     "postgres",
	Password: "password",
	Name:     "postgres",
}
