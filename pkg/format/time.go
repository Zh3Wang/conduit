package format

import "time"

func ConvertTime(t int64) string {
	return time.Unix(t, 0).Format("2006-01-02T15:04:05.000Z07:00")
}
