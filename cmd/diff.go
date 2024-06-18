/*
Copyright © 2024 ValentinTT
*/
package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

// diffCmd represents the diff command
var diffCmd = &cobra.Command{
	Use:   "diff",
	Short: "Calculate days, weeks, months and years between two dates",
	Long: `Calculate the difference between two dates in: 
• years, months and days
• weeks and days
• days
• hours

For example
> gone diff '5 10 2020' '17-6-2024'

• start date: Mon Oct 5 2020 (05/10/2020)
• end date: Mon Jun 17 2024 (17/06/2024)

Difference:
• 3 years 8 months 13 days
• 193 weeks 4 days
• 1351 days
• 32424 hours

◘ you can use ' ', '.', '-' or '/' as separators: '5 10-2020' equals '05/10.2020'

FLAGS
◘ layout (l) It also support many date formats:
> l="dmy" // dd-mm-yyyy (default)
	>If you ommit the year in the 'dmy' layout it takes the current year
> l="mdy" // mm-dd-yyyy
> l="ymd" // yyyy-mm-dd
[!] Both dates must be in the same format.

◘ today (t) takes only one date and the current date`,

	Args: cobra.RangeArgs(1, 2),
	RunE: diffTwoDates,
}

func init() {
	rootCmd.AddCommand(diffCmd)
	diffCmd.Flags().StringP("layout", "l", "dmy", "dates format/layout")
	diffCmd.Flags().BoolP("today", "t", false, "use current date as second date")
	customUsageTemplate := `Usage:
  gone diff "date1" ["date2"] [flags]

Flags:
  -h, --help            help for diff
  -l, --layout string   dates format/layout (default "dmy")
  -t, --today           use current date as second date
`
	diffCmd.SetUsageTemplate(customUsageTemplate)
}

var layouts = map[string]string{
	"dmy": "02-01-2006", // dd-mm-yyyy
	"mdy": "01-02-2006", // mm-dd-yyyy
	"ymd": "2006-01-02", // yyyy-mm-dd
}

func diffTwoDates(cmd *cobra.Command, args []string) error {
	layout_flag, _ := cmd.Flags().GetString("layout")
	today_flag, _ := cmd.Flags().GetBool("today")

	if !isValidLayoutFlag(layout_flag) { // error
		return errors.New("invalid layout flag")
	}
	if today_flag && len(args) != 1 { // error way two many args
		return errors.New("way to many args")
	} else if !today_flag && len(args) != 2 { //error missing date
		return errors.New("msissing date (or missing --today flag)")
	}

	date1 := &Date{}
	date1.parseDate(args[0], layout_flag)
	date2 := &Date{}
	if today_flag {
		date2.parseToday()
	} else {
		date2.parseDate(args[1], layout_flag)
	}

	fmt.Printf("\n• start date: %v (%v)\n", date1.date.Format("Mon Jan 2 2006"), date1.date.Format("02/01/2006"))
	fmt.Printf("• end date: %v (%v)", date2.date.Format("Mon Jan 2 2006"), date2.date.Format("02/01/2006"))
	if today_flag {
		fmt.Printf(" [today flag]")
	}
	fmt.Printf("\n\n")

	years, months, days := calculateYears(date1.date, date2.date)
	fmt.Println("Difference:")
	fmt.Printf("•")
	if years != 0 {
		fmt.Printf(" %d year", years)
		if years > 1 {
			fmt.Printf("s")
		}
	}
	if months != 0 {
		fmt.Printf(" %d month", months)
		if months > 1 {
			fmt.Printf("s")
		}
	}
	if days != 0 {
		fmt.Printf(" %d day", days)
		if days > 1 {
			fmt.Printf("s")
		}
	}
	fmt.Printf("\n")

	days = calculateDays(date1.date, date2.date)
	weeks := int(float64(days / 7))
	if weeks != 0 {
		fmt.Printf("• %d week", weeks)
		if weeks > 1 {
			fmt.Printf("s")
		}

		daysForWeek := weeks % 7
		if daysForWeek != 0 {
			fmt.Printf(" %d day", daysForWeek)
			if daysForWeek > 1 {
				fmt.Printf("s")
			}
		}
		fmt.Printf("\n")
	}

	fmt.Printf("• %d day", days)
	if days > 1 {
		fmt.Printf("s")
	}
	fmt.Printf("\n")

	hours := calculateHours(date1.date, date2.date)
	fmt.Printf("• %d hour", hours)
	if hours > 1 {
		fmt.Printf("s")
	}
	fmt.Printf("\n")

	return nil
}

func isValidLayoutFlag(layout_flag string) bool {
	_, ok := layouts[layout_flag]
	return ok
}
