package config

import (
	"log"
	"os"
	"strconv"
	"time"
)

type Config struct {
	Server struct {
		Port     string        // .env: SERVER_PORT
		KillTime time.Duration // .env: SERVER_KILL_TIME
		// Time allotted for program to shutdown gracefully
	}
	Database struct {
		Name     string // .env: DB_NAME
		Host     string // .env: DB_HOST
		Port     string // .env: DB_PORT
		User     string // .env: DB_USER
		Password string // .env: DB_PASSWORD
	}
	Session struct {
		Lifespan time.Duration //.env: SESSION_LIFESPAN
	}
}

func New() *Config {
	return &Config{}
}

func (cfg *Config) Init() {
	// Server Configuration
	srvPort := os.Getenv("SERVER_PORT")
	if srvPort == "" {
		log.Printf("missing server port at 'SERVER_PORT'; defaulting to ':8080'")
		srvPort = ":8080"
	} else {
		srvPort = ":" + srvPort
	}

	var srvKillTime time.Duration
	srvKillStr := os.Getenv("SERVER_KILL_TIME")
	if srvKillStr == "" {
		log.Printf("missing server kill time at 'SERVER_KILL_TIME'; defaulting to '10' seconds")
		srvKillTime = time.Duration(10)
	} else {
		srvKillInt, err := strconv.Atoi(srvKillStr)
		if err != nil {
			log.Printf("invalid kill time at 'SERVER_KILL_TIME'; defaulting to '10' seconds")
			srvKillTime = time.Duration(10)
		} else {
			srvKillTime = time.Duration(srvKillInt)
		}
	}

	// Database Configuration
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		log.Printf("missing database name at 'DB_NAME'; defaulting to 'database'")
		dbName = "database"
	}

	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		log.Printf("missing database host at 'DB_HOST'; defaulting to 'database_host'")
		dbHost = "database_host"
	}

	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		log.Printf("missing database port at 'DB_PORT'; defaulting to ':5432'")
		dbPort = ":5432"
	}

	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		log.Printf("missing database user at 'DB_USER'; defaulting to 'database_user'")
		dbUser = "database_user"
	}

	dbPass := os.Getenv("DB_PASSWORD")
	if dbPass == "" {
		log.Printf("missing database password at 'DB_PASSWORD'; defaulting to 'database_password'")
		dbPass = "database_password"
	}

	// Session Configuration
	var sessLife time.Duration
	sessLifeStr := os.Getenv("SESSION_LIFESPAN")
	if sessLifeStr == "" {
		log.Printf("missing session lifespan at 'SESSION_LIFESPAN'; defaulting to 43200 seconds (12 hours)")
		sessLife = time.Duration(43200) * time.Second
	} else {
		var err error
		sessLife, err = time.ParseDuration(sessLifeStr)
		if err != nil {
			log.Printf("invalid session life at 'SESSION_LIFESPAN'; defaulting to 43200 seconds (12 hours)")
			sessLife = time.Duration(43200) * time.Second
		}
	}

	cfg.Server.Port = srvPort
	cfg.Server.KillTime = srvKillTime

	cfg.Database.Name = dbName
	cfg.Database.Host = dbHost
	cfg.Database.Port = dbPort
	cfg.Database.User = dbUser
	cfg.Database.Password = dbPass

	cfg.Session.Lifespan = sessLife
}
