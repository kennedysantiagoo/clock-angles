package helper

import "strconv"

func GetParams(vars map[string]string) (string, string) {
	hourParam := vars["hour"]
	minuteParam := vars["minute"]
	return hourParam, minuteParam
}

func ConvertParamsToInt(hour, minute string) (int, int) {
	intHour, _ := strconv.Atoi(hour)
	var intMinute int

	if minute != "" {
		intMinute, _ = strconv.Atoi(minute)
		return intHour, intMinute
	}

	intMinute = 0
	return intHour, intMinute
}

func NegativeToPositive(number int) int {
	if number < 0 {
		return -number
	}
	return number
}
