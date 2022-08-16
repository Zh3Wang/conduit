package encrypt

import "golang.org/x/crypto/bcrypt"

func Hash(ori string) string {
	b, err := bcrypt.GenerateFromPassword([]byte(ori), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func Verify(hashed, input string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(input))
}
