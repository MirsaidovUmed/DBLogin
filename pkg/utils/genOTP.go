package utils

import (
	"fmt"
	"math/rand"
)

func GenOTP(n int) (res string) {

	for i := 0; i < n; i++ {
		res += fmt.Sprint(rand.Intn(9))
	}

	return
}
