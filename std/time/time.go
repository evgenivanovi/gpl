package time

import "time"

/* __________________________________________________ */

func NowPlus(value time.Duration) time.Time {
	return time.Now().Add(value)
}

func NowMinus(value time.Duration) time.Time {
	return time.Now().Add(-value)
}

/* __________________________________________________ */
