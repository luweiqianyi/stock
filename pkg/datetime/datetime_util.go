package datetime

import "time"

func RFC3339ToDateTimeFormat(in string) (string, error) {
	t, err := time.Parse(time.RFC3339, in)
	if err != nil {
		return "", err
	}
	return t.Format(time.DateTime), nil
}
