package utils

// ToValue returns the memory address of the string.
func ToValue[T any](s T) *T {
	return &s
}

// FromValue returns the string from the memory address.
func FromValue[T any](s *T) T {
	return *s
}
