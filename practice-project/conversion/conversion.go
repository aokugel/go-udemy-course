package conversion

import "strconv"

func StringsToFloats(strings []string) ([]float64, error) {
	floats := make([]float64, len(strings))

	for i, val := range strings {
		float, err := strconv.ParseFloat(val, 64)
		if err != nil {
			return []float64{}, err
		}
		floats[i] = float
	}
	return floats, nil
}
