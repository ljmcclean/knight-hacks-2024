package services

type Database interface {
	ProfileService
	ProjectService
	SessionService
	Close()
}
