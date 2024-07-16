package internal

import "github.com/sosodev/duration"

func StringToDuration(s string) (*duration.Duration, error) {
	if s == "" {
		return nil, nil
	}

	dur, err := duration.Parse(s)
	if err != nil {
		return nil, err
	}
	return dur, nil
}

func DurationToString(dura *duration.Duration) string {
	if dura == nil {
		return ""
	}
	return dura.String()

}
