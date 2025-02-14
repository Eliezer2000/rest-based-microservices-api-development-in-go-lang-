package domain

import (
	"fmt"
	"time"
)

type TimeService struct {
}

func (ts *TimeService) GetCurrentTime(tz string) (string, error) {
	if tz == "" {
		tz = "UTC"
	}
	loc, err := time.LoadLocation(tz)
	if err != nil {
		return "", fmt.Errorf("invalid timezone: %s", tz)
	}
	currentTime := time.Now().In(loc)
	return currentTime.Format("2006-01-02 15:04:05 -0700 MST"), nil
}