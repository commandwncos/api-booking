package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(pass string) ([]byte, error) {
	return bcrypt.GenerateFromPassword(
		[]byte(pass),
		bcrypt.DefaultCost,
	)
}

func CheckPassword(pass, hashPass string) bool {
	return bcrypt.CompareHashAndPassword(
		[]byte(hashPass),
		[]byte(pass),
	) == nil
}
