package geardate

import (
	"time"
)

// built-in format layout
const (
	DefaultLayout      = "2006-01-02 15:04:05"
	ReverseLayout      = "15:04:05 01-02-2006"
	SlashLayout        = "2006/01/02 15:04:05"
	ZhLayout           = "2006年01月02日 15时04分05秒"
	OnlyYMDLayout      = "2006-01-02"
	OnlyYMDSlashLayout = "2006/01/02"
	OnlyHMSLayout      = "15:04:05"
	OnlyYMDZhLayout    = "2006年01月02日"
	OnlyHMSZhLayout    = "15时04分05秒"
)

// Format formats a unix timestamp according the layout
func Format(t int64, layout string) string {
	tm := time.Unix(t, 0)
	return tm.Format(layout)
}

// IsLeap weather the year is leap
func IsLeap(year int) bool {
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		return true
	}
	return false
}
