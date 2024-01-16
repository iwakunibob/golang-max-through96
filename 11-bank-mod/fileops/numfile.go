package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func GetNumFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)
	if err != nil { // Error handler
		errMsg := fmt.Sprintf("Error = Failed to find %s file", fileName)
		return 0, errors.New(errMsg)
	}
	numText := string(data)
	num, err := strconv.ParseFloat(numText, 64)
	if err != nil { // Error handler
		return 0, errors.New("Error = Failed to parse strored balance value")
	}
	return num, nil
}

func WriteNumToFile(num float64, fileName string) {
	numText := fmt.Sprint(num)
	os.WriteFile(fileName, []byte(numText), 0644)
}
