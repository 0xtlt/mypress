package utils

// Includes func
func Includes(eq string, strings []string) bool {
	check := false

	for _, str := range strings {
		if str == eq {
			check = true
		}
	}

	return check
}
