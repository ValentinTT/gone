package cmd

import (
	"strconv"
	"testing"
	"time"
)

// Helper function to create a Date struct and call parseDate
func testParseDate(t *testing.T, input, layoutFlag string, expectedDay, expectedMonth, expectedYear string, expectError bool) {
	date := &Date{}
	err := date.parseDate(input, layoutFlag)

	if expectError {
		if err == nil {
			t.Errorf("Expected error but got nil for input: %s", input)
		}
	} else {
		if err != nil {
			t.Errorf("Did not expect error but got %v for input: %s", err, input)
		}
		if date.day != expectedDay || date.month != expectedMonth || date.year != expectedYear {
			t.Errorf("Expected date %s-%s-%s, but got %s-%s-%s",
				expectedDay, expectedMonth, expectedYear,
				date.day, date.month, date.year)
		}
	}
}

func TestParseDate(t *testing.T) {currentYear := time.Now().Year()

	// Test case for a valid date using the "dmy" layout.
	// The date "01-02-2023" should parse correctly as day "01", month "02", year "2023".
	testParseDate(t, "01-02-2023", "dmy", "01", "02", "2023", false)
	
	// Test case for a date missing the year using the "dmy" layout.
	// The date "01-02" should default to the current year. The day is "01", month is "02", and the year is the current year.
	testParseDate(t, "01-02", "dmy", "01", "02", strconv.Itoa(currentYear), false)
	
	// Test case for a layout mismatch error.
	// The date "01-2023-02" does not match the "dmy" layout and should return an error.
	testParseDate(t, "01-2023-02", "dmy", "", "", "", true)
	
	// Test case for a date with only day and month using the "dmy" layout.
	// The date "01-02" should parse as day "01", month "02", and default to the current year.
	testParseDate(t, "01-02", "dmy", "01", "02", strconv.Itoa(currentYear), false)
	
	// Test case for a valid date using the "dmy" layout (duplicate of the first test).
	// The date "01-02-2023" should parse as day "01", month "02", year "2023".
	testParseDate(t, "01-02-2023", "dmy", "01", "02", "2023", false)
	
	// Test case for a two-digit year using the "dmy" layout.
	// The date "02-01-23" should parse as day "02", month "01", and year "2023" (assumes 21st century).
	testParseDate(t, "02-01-23", "dmy", "02", "01", "2023", false)
	
	// Test case for a date missing the year using the "dmy" layout.
	// The date "15-04" should parse as day "15", month "04", and default to the current year.
	testParseDate(t, "15-04", "dmy", "15", "04", strconv.Itoa(currentYear), false)
	
	// Test case for a valid date using the "mdy" layout.
	// The date "02-01-2023" should parse as month "02", day "01", and year "2023".
	testParseDate(t, "02-01-2023", "mdy", "01", "02", "2023", false)
	
	// Test case for a two-digit year using the "mdy" layout.
	// The date "12-31-22" should parse as month "12", day "31", and year "2022".
	testParseDate(t, "12-31-22", "mdy", "31", "12", "2022", false)
	
	// Test case for a date missing the year using the "mdy" layout.
	// The date "08-25" should parse as month "08", day "25", and default to the current year.
	testParseDate(t, "08-25", "mdy", "25", "08", strconv.Itoa(currentYear), false)
	
	// Test case for a valid date using the "ymd" layout.
	// The date "2023-02-01" should parse as year "2023", month "02", and day "01".
	testParseDate(t, "2023-02-01", "ymd", "01", "02", "2023", false)
	
	// Test case for a date with two-digit year using the "ymd" layout.
	// The date "23-03-04" should parse as year "2023", month "03", and day "04".
	testParseDate(t, "23-03-04", "ymd", "04", "03", "2023", false)
	
	// Test case for a date with a two-digit year using the "ymd" layout.
	// The date "22-05-09" should parse as year "2022", month "05", and day "09".
	testParseDate(t, "22-05-09", "ymd", "09", "05", "2022", false)
	
}
