package sentinel

type Sentinel interface {
	IsSpecified() bool
}

// Generic TakeOrElse for any type with IsSpecified() method
func TakeOrElse[T interface{ Sentinel }](a, b, unspecified T) T {
	if a.IsSpecified() {
		return a
	}
	return b
}
