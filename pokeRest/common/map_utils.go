package common

func Mapkeys(m map[string]struct{}) []string {
	keys := []string{}
	for k, _ := range m {
		keys = append(keys, k)
	}
	return keys
}
