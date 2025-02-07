package utils

import "time"

func ConvertDate(date string) (time.Time, error) {
	if date == "" {
		return time.Now(), nil
	}

	t, err := time.Parse(time.RFC3339, date)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}
