package main

type Config struct {
	Default string
	Drivers map[string]DriverConfig
}

type DriverConfig struct {
	Database string
	UserName string
	Password string
	Port     int
}
