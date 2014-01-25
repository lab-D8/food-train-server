package error

const SUCCESS = 0     // success, no error
const DBERROR = 1     // failure, db error
const ALREADY = 2     // failure, already in DB
const NOTFOUND = 4    // failure, not found in DB
const WEFUCKEDUP = 8  // failure, something bad on our end
const BADTOKEN = 16   // failure, the token is expired
const NOTFRIENDS = 32 // failure, they aren't friends

var errorNames = map[int]string{
	0:  "Success",
	1:  "Error: DB failed",
	2:  "Error: Request didn't go through",
	4:  "Error: Request did not go through",
	8:  "Error: We fucked up T_T",
	16: "Error: Invalid request",
	32: "Error: These users aren't friends",
}

func GetDescription(code int) string {
	return errorNames[code]
}
