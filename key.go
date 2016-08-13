package attr

// Keyer is attr key
type Keyer interface {
	Name() string
}

// KeyHolder has keys
type KeyHolder interface {
	Keys() []Keyer
}

// KeyEquals key check function
func KeyEquals(str1, str2 Keyer) bool {
	return (str1.Name() == str2.Name())
}

// KeyExists check the KeyHolder have key or not
func KeyExists(kh KeyHolder, key Keyer) bool {
	for _, k := range kh.Keys() {
		if KeyEquals(k, key) {
			return true
		}
	}

	return false
}

// KeyDiff Check key diff
func KeyDiff(base, target KeyHolder) map[Keyer]bool {
	res := map[Keyer]bool{}
	baseKeys := base.Keys()

	for _, k := range baseKeys {
		res[k] = KeyExists(target, k)

	}

	return res
}

// KeyMerge merge keys
func KeyMerge(base, target KeyHolder) []Keyer {
	res := base.Keys()

	for _, k := range target.Keys() {
		if !KeyExists(base, k) {
			res = append(res, k)
		}
	}

	return res
}
