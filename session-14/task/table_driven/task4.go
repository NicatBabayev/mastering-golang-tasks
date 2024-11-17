package table_driven

import "fmt"

func ConvertToF(temperature float64) float64 {
	return temperature*1.8 + 32
}

func ConvertToC(temperature float64) float64 {
	return (temperature - 32) / 1.8
}

func ConvertTemperature(temperature float64, scale string) (float64, error) {
	if scale == "C" {
		if temperature < -100 || temperature > 100 {
			return 0, fmt.Errorf("Temperature should be in the range of 0 and 100")
		} else {
			return ConvertToF(temperature), nil
		}
	} else if scale == "F" {
		if temperature < -212 || temperature > 212 {
			return 0, fmt.Errorf("Temperature should be in the range of 32 and 212")
		} else {
			return ConvertToC(temperature), nil
		}
	} else {
		return 0, fmt.Errorf("Please select correct scale. It can be either C or F.")
	}
}

func RunTask4() {
	fmt.Println(ConvertTemperature(22, "C"))
}
