package ptr

func OrElse[T any](ptr *T, orElse T) T {
	if ptr == nil {
		return orElse
	}
	return *ptr
}

func OrElseF[T any](ptr *T, orElse func() T) T {
	if ptr == nil {
		return orElse()
	}
	return *ptr
}

func ToPtr[T any](value T) *T {
	return &value
}
