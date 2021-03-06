package attr

// Keyer is attr key
type Keyer interface {
	Name() string
}

// KeyerList is Keyer list
type KeyerList []Keyer

// Len for sort
func (k KeyerList) Len() int {
	return len(k)
}

// Swap for sort
func (k KeyerList) Swap(i, j int) {
	k[i], k[j] = k[j], k[i]
}

// Less for sort
func (k KeyerList) Less(i, j int) bool {
	return k[i].Name() < k[j].Name()
}

// KeyHolder has keys
type KeyHolder interface {
	Keys() KeyerList
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

// KeyCommon get common key
func KeyCommon(base, target KeyHolder) []Keyer {
	res := []Keyer{}

	for _, k := range target.Keys() {
		if KeyExists(base, k) {
			res = append(res, k)
		}
	}

	return res
}

// Key is attribute key
type Key struct {
	name string
}

// NewKey create Key object
func NewKey(name string) Key {
	return Key{
		name: name,
	}
}

// Name return name
func (k Key) Name() string {
	return k.name
}
