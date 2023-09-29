package bcrypt

import "golang.org/x/crypto/bcrypt"

func Hash(secret string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(secret), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

