package sourcegraphoperator

import "time"

// LifecycleDuration returns the converted lifecycle duration from given minutes.
// It returns the default duration (60 minutes) if the given minutes is
// non-positive.
func LifecycleDuration(minutes int) time.Duration {
	if minutes <= 0 {
		return 60 * time.Minute
	}
	return time.Duration(minutes) * time.Minute
}
