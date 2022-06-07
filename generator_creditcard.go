package go_generator_creditcard

import (
	crypto_rand "crypto/rand"
	"encoding/binary"
	"math"
	math_rand "math/rand"
	"strconv"
)

func CreditCardNumber(prefixList []string, length, howMany int) []int64 {

	var result []int64
	for i := 0; i < howMany; i++ {

		// 随机选择一个
		var randomArrayIndex = randomInt(0, len(prefixList)-1)
		var ccnumber = prefixList[randomArrayIndex]
		result = append(result, completedNumber(ccnumber, length))
	}

	return result
}

func completedNumber(prefix string, length int) int64 {
	var ccnumber = prefix

	// generate digits
	for i := len(prefix); i < (length - 1); i++ {
		randomNumber := randomInt(0, 9)
		ccnumber += strconv.Itoa(randomNumber)
	}

	// reverse number and convert to int
	reversedCCnumberString := reverse(ccnumber)
	var reversedCCnumber []int
	for i := 0; i < len(reversedCCnumberString); i++ {
		number, _ := strconv.Atoi(string(reversedCCnumberString[i]))
		reversedCCnumber = append(reversedCCnumber, number)
	}

	// calculate sum
	var sum = 0
	var pos = 0
	if len(reversedCCnumber) > 0 {
		for {
			if pos < length-1 {
				odd := reversedCCnumber[pos] * 2
				if odd > 9 {
					odd -= 9
				}

				sum += odd

				if pos != (length - 2) {
					sum += reversedCCnumber[pos+1]
				}
				pos += 2
			} else {
				break
			}

		}
	}

	checkdigit := (int(math.Floor(float64(sum)/10)+1)*10 - sum) % 10
	ccnumber += strconv.Itoa(checkdigit)

	result2, _ := strconv.Atoi(ccnumber)
	return int64(result2)
}

func randomInt(start int, end int) int {
	var b [8]byte
	_, err := crypto_rand.Read(b[:])
	if err != nil {
		panic("cannot seed math/rand package with cryptographically secure random number generator")
	}
	math_rand.Seed(int64(binary.LittleEndian.Uint64(b[:])))

	random := math_rand.Intn(end - start)
	random = start + random
	return random
}

func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
