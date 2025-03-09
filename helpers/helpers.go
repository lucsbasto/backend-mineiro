package helpers

import "fmt"

func Coalesce(newValue, defaultValue int) int {
	fmt.Println(" Coalesce newValue: ", newValue, "defaultValue: ", defaultValue)
	if newValue != 0 {
		return newValue
	}
	return defaultValue
}

func CoalesceFloat64(newValue, defaultValue float64) float64 {
	fmt.Println(" CoalesceFloat64 newValue: ", newValue, "defaultValue: ", defaultValue)

	if newValue != 0 {
		return newValue
	}
	return defaultValue
}
