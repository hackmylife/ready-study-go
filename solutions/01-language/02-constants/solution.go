//go:build ignore

package koan

const MinutesPerHour = 60
const HoursPerDay = 24

func MinutesInDays(days int) int { return days * HoursPerDay * MinutesPerHour }
