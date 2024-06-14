package conversion

import (
	"errors"
	"strconv"
)


func StringsToFloats(strings []string) ([]float64, error) {

	var floats []float64

	for _, stringValue := range strings {
		floatVal, err := strconv.ParseFloat(stringValue, 64)

		if err != nil {
			return nil, errors.New("failed ot convert string to float")
		}

		floats = append(floats, floatVal)

	}

	return floats, nil

}