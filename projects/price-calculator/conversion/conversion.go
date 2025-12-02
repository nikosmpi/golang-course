package conversion

import (
	"strconv"
)

func StringsToFloats(strings []string) ([]float64, error) {
	floats := make([]float64, len(strings))
	for stringIndex, stringValue := range strings {
		floatPrice, err := strconv.ParseFloat(stringValue, 64)
		if err != nil {
			return nil, err
		}
		floats[stringIndex] = floatPrice
	}
	return floats, nil
}
