package slice

// ContainsString returns if value in values.
func ContainsString(value string, values []string) bool {
	for _, s := range values {
		if s == value {
			return true
		}
	}

	return false
}