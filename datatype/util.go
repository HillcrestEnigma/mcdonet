package datatype

import (
	"errors"
	"math"
	"strings"
)

func PackIntoLongArray[T Number](bpe uint8, array []T) []int64 {
	var shift uint8 = 0
	var long uint64 = 0

	longArray := make([]int64, int(math.Ceil(float64(len(array))/math.Floor(64.0/float64(bpe)))))
	longArrayIndex := 0

	for _, val := range array {
		long |= uint64(val) << shift
		shift += bpe
		if shift+bpe > 64 {
			longArray[longArrayIndex] = int64(long)

			shift = 0
			long = 0
			longArrayIndex++
		}
	}

	if shift != 0 {
		longArray[longArrayIndex] = int64(long)
	}
	return longArray
}

func ParseIdentifier(identifier string) (namespace, value string, err error) {
	parts := strings.SplitN(identifier, ":", 2)
	switch len(parts) {
		case 1:
			return "minecraft", parts[0], nil
		case 2:
			return parts[0], parts[1], nil
		default:
			return "", "", errors.New("invalid identifier")
	}
}
