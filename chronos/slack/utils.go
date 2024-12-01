package slack

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func calculate(dob string) (int, error) {
	parts := strings.Split(dob, "-")
	if len(parts) != 3 {
		return 0, fmt.Errorf("invalid date format")
	}

	year, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, fmt.Errorf("invalid year")
	}

	month, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, fmt.Errorf("invalid month")
	}

	day, err := strconv.Atoi(parts[2])
	if err != nil {
		return 0, fmt.Errorf("invalid day")
	}

	dobTime := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	now := time.Now()
	age := now.Year() - dobTime.Year()
	if now.Before(dobTime.AddDate(age, 0, 0)) {
		age--
	}

	return age, nil
}
