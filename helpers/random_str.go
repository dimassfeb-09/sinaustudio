package helpers

import "math/rand"

func RandStr(n int) string {

	var charset = []byte("123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]byte, n)
	for i := range b {
		// randomly select 1 character from given charset
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
