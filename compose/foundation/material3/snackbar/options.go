package snackbar

import "time"

type SnackbarOptions struct {
	Duration time.Duration
}

type SnackbarOption func(*SnackbarOptions)

func WithDuration(duration time.Duration) SnackbarOption {
	return func(o *SnackbarOptions) {
		o.Duration = duration
	}
}

func DefaultOptions() SnackbarOptions {
	return SnackbarOptions{
		Duration: 4 * time.Second,
	}
}
