package utils

func ContainsAll(tags1, tags2 []string) bool {

	set := make(map[string]struct{})
	for _, tag := range tags1 {
		set[tag] = struct{}{}
	}

	for _, tag := range tags2 {
		if _, exists := set[tag]; !exists {
			return false
		}
	}
	return true
}
