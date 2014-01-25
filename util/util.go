package util

import (
	"crypto/sha512"
	"encoding/hex"
)

func GenerateSHA512Hash(str string) (returnStr string) {
	h := sha512.New()
	h.Write([]byte(str))
	returnStr = hex.EncodeToString(h.Sum(nil))
	return returnStr
}


func HashPassword(password string) (hashed_password string){
	HASH_AMOUNT := 50
	hashed_password = password
	for i := 0; i < HASH_AMOUNT; i++ {
		hashed_password = GenerateSHA512Hash(hashed_password)
	}

	return hashed_password
}
