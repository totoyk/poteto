package typecast

func Bool2Bptr(b bool) *bool {
	return &b
}

func Boolptr2B(b *bool) bool {
	if b == nil {
		return false
	}
	return *b
}
