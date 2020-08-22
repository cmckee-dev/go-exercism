// Package leap handles functionality around Leap Years.
package leap

// IsLeapYear returns a bool if the given year is a leap year.
func IsLeapYear(year int) bool {
	if (year%4 == 0) && (year%100 != 0 || year%400 == 0) {
		return true
	}

	return false
}
