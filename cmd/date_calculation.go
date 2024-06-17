package cmd

import "time"

func calculateHours(start, end time.Time) int {
	// Swap if start date is after end date
	if start.After(end) {
		start, end = end, start
	}
	duration := end.Sub(start)
	return int(duration.Hours())
}

func calculateDays(start, end time.Time) int {
	// Swap if start date is after end date
	if start.After(end) {
		start, end = end, start
	}
	duration := end.Sub(start)
	return int((duration.Hours()) / 24)
}

// Function to calculate years, months and days difference between two dates
func calculateYears(start, end time.Time) (years, months, days int) {
	// Swap if start date is after end date
	if start.After(end) {
		start, end = end, start
	}

	years = end.Year() - start.Year()
	months = int(end.Month() - start.Month())
	days = end.Day() - start.Day() + 1
	if start.Day() > end.Day() {
		months--
		days += daysIn(start.Year(), start.Month())
	}
	if months <= 0 {
		years--
		months += 12
	}

	return years, months, days
}

// Function to get number of days in a month
func daysIn(year int, month time.Month) int {
	return time.Date(year, month, 0, 0, 0, 0, 0, time.UTC).Day()
}
