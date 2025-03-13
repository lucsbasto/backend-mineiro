package helpers

func Coalesce(newValue, defaultValue int) int {
	if newValue != 0 {
		return newValue
	}
	return defaultValue
}

func CoalesceFloat64(newValue, defaultValue float64) float64 {
	if newValue != 0 {
		return newValue
	}
	return defaultValue
}
