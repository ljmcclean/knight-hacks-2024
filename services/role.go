package services

type Role struct {
	ID          int
	Name        string
	Description string
	Skills      []*Skill
}
