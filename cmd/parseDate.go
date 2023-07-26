package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/GLAD-DEV/dn/constants"
)

func parseDate(dateStr string) (string, error) {
	dateStr = strings.ReplaceAll(dateStr, ".", "-")

	// Is date YYYY-MM-DD?
	date, err := time.Parse("2006-01-02", dateStr)
	if err == nil {
		return date.Format(constants.DateFormat), nil
	}

	// Is date DD-MM-YYYY?
	date, err = time.Parse("02-01-2006", dateStr)
	if err == nil {
		return date.Format(constants.DateFormat), nil
	}

	return "", fmt.Errorf("the passed date '%s' is incorrectly formatted: %w", dateStr, err)
}
