package input;

import (
	"github.com/nyaruka/phonenumbers"
    "math/rand"
    "strconv"
    "time"
)

func IsValidPhone(phone string) bool {
	num, err := phonenumbers.Parse(phone, "")
	if err != nil {
		return false
	}
	if !phonenumbers.IsValidNumber(num) {
		return false
	}
	return true
}

func RandomStringWithInteger(length int, num int) string {
    const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
    seed := rand.NewSource(time.Now().UnixNano())
    randGen := rand.New(seed)

    // Generate the random position
    pos := randGen.Intn(length + 1)

    b := make([]byte, length)
    for i := range b {
        b[i] = charset[randGen.Intn(len(charset))]
    }

    // Convert the random string to a slice of bytes
    randomStr := string(b)

    // Insert the integer at the random position
    numStr := strconv.Itoa(num)
    modifiedStr := randomStr[:pos] + numStr + randomStr[pos:]

    return modifiedStr
}