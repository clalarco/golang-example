package names

import (
	"errors"
	"math/rand"
)

func GetNames(length int) ([]string, error) {
	allNames := []string{
		"Chris",
		"Katie",
		"Jenny",
		"Cata",
		"Ema",
		"Fran",
		"Perico",
		"Pagüencio",
	}

	if length < 1 {
		return nil, errors.New("length must be greater than 1")
	}

	returnedNames := []string{}
	for i := 0; i < length; i++ {
		returnedNames = append(returnedNames, allNames[rand.Intn(len(allNames))])
	}
	return returnedNames, nil

}
