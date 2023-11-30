package util

import "time"

func ToDate(dateString string) *time.Time {
	parsedDate, _ := time.Parse(dateString, dateString)
	return &parsedDate
}
