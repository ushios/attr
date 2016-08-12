package attr

// KeyHolder has keys
type KeyHolder interface {
	Keys() []string
}

// KeyValuer return valuer from key
type KeyValuer interface {
	KeyValue(string) interface{}
}

// KeyEquals key check function
func KeyEquals(str1, str2 string) bool {
	return (str1 == str2)
}

// KeyExists check the KeyHolder have key or not
func KeyExists(kh KeyHolder, key string) bool {
	for _, k := range kh.Keys() {
		if KeyEquals(k, key) {
			return true
		}
	}

	return false
}

// KeyDiff Check key diff
func KeyDiff(base, target KeyHolder) map[string]bool {
	res := map[string]bool{}
	baseKeys := base.Keys()

	for _, k := range baseKeys {
		res[k] = KeyExists(target, k)

	}

	return res
}

// KeyMerge merge keys
// func KeyMerge(base, target KeyHolder) []string {
//
// }
