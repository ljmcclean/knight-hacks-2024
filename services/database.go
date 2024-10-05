package services

type Database interface {
	ProfileService
	SessionService
	Close()
}
