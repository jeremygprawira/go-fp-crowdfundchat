package util

import (
	"errors"
	"strconv"
)

/*
func repeatingPIN(inputtedString string) (string, error) {
	convertedString, err := strconv.Atoi(inputtedString)
	if err != nil {
		return inputtedString, errors.New("failed to convert string into integer")
	}

	if convertedString % 1111 == 0 {
		errors.New("pin cannot be repeated.")
	}

	repeatingPINChecker := strconv.Itoa(convertedString)

	return repeatingPINChecker, nil
}*/

func RepeatingPINChecker(stringPIN string) (string, error) {
	integerPIN, err := strconv.Atoi(stringPIN)

	if err != nil {
		return stringPIN, err
	}

	if integerPIN % 111111 == 0 {
		return stringPIN, errors.New("PIN cannot be 111111")
	}

	return strconv.Itoa(integerPIN), nil
}
/*
// Online Go compiler to run Golang program online
// Print "Hello World!" message

package main
import (
	"strconv"
	"fmt"
)

func main() {
    var num = "1111"
    conNumb, _ := strconv.Atoi(num)
    
    if conNumb % 1111 == 0 {
        fmt.Println("nt")
        return
    }
    
    fmt.Println(num)
}
*/