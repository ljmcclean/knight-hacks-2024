package auth

import (
	"context"
	"log"
	"knight-hacks-2024/services"

	"golang.org/x/crypto/bcrypt"
)

// TODO: Verify email address
func RegisterProfile(ctx context.Context, ps services.ProfileService, name, email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("error hashing user password: %s", err)
		return err
	}
	profile := &services.Profile{
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
	}

	ps.PostProfile(ctx, profile)

	return nil
}
