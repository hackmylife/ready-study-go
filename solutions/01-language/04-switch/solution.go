//go:build ignore

package koan

func DayKind(day string) string {
	switch day {
	case "Sat", "Sun":
		return "weekend"
	case "Mon", "Tue", "Wed", "Thu", "Fri":
		return "weekday"
	default:
		return "invalid"
	}
}
