package validator

import "time"

type Time interface {
	IsTimeEmpty(t time.Time) bool
	IsTimeAfter(t time.Time, t2 time.Time) bool
	IsTimeBefore(t time.Time, t2 time.Time) bool
	IsTimeEqual(t time.Time, t2 time.Time) bool
	IsTimeAfterNow(t time.Time) bool
	IsTimeBeforeNow(t time.Time) bool
	IsTimeEqualNow(t time.Time) bool
	IsTimeAfterNowUTC(t time.Time) bool
	IsTimeBeforeNowUTC(t time.Time) bool
	IsTimeEqualNowUTC(t time.Time) bool
}

func IsTimeEmpty(t time.Time) bool {
	return t.IsZero()
}

func IsTimeAfter(t time.Time, t2 time.Time) bool {
	return t.After(t2)
}

func IsTimeBefore(t time.Time, t2 time.Time) bool {
	return t.Before(t2)
}

func IsTimeEqual(t time.Time, t2 time.Time) bool {
	return t.Equal(t2)
}

func IsTimeAfterNow(t time.Time) bool {
	return t.After(time.Now())
}

func IsTimeBeforeNow(t time.Time) bool {
	return t.Before(time.Now())
}

func IsTimeEqualNow(t time.Time) bool {
	return t.Equal(time.Now())
}

func IsTimeAfterNowUTC(t time.Time) bool {
	return t.After(time.Now().UTC())
}

func IsTimeBeforeNowUTC(t time.Time) bool {
	return t.Before(time.Now().UTC())
}

func IsTimeEqualNowUTC(t time.Time) bool {
	return t.Equal(time.Now().UTC())
}
