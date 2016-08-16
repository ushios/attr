package attr

import (
	"reflect"
	"testing"
)

// SampleKey key
type SampleKey struct {
	name string
}

// Name return name
func (s SampleKey) Name() string {
	return s.name
}

type SampleKeyHolder struct {
	KeyList []string
}

// Keys return key list
func (s SampleKeyHolder) Keys() KeyerList {
	ks := []Keyer{}

	for _, k := range s.KeyList {
		ks = append(ks, SampleKey{name: k})
	}

	return ks
}

func TestKeyExists(t *testing.T) {
	test := func(keys []string, checkStr string, e bool) {
		kh := SampleKeyHolder{KeyList: keys}

		res := KeyExists(kh, SampleKey{name: checkStr})
		if res != e {
			t.Errorf("(%s) key exists (%s) result expected (%t) but (%t)", keys, checkStr, e, res)
		}
	}

	test([]string{"a", "b", "c"}, "a", true)
	test([]string{"a", "b", "c"}, "d", false)
	test([]string{"A", "b", "c"}, "a", false)
}

func TestKeyDiff(t *testing.T) {
	test := func(base, target []string, e map[Keyer]bool) {
		res := KeyDiff(
			SampleKeyHolder{KeyList: base},
			SampleKeyHolder{KeyList: target},
		)

		if !reflect.DeepEqual(res, e) {
			t.Errorf("(%s) diff (%s) expected (%v) but (%v)", base, target, e, res)
		}

	}

	test([]string{"a", "b", "c"}, []string{"a", "b"}, map[Keyer]bool{
		SampleKey{name: "a"}: true,
		SampleKey{name: "b"}: true,
		SampleKey{name: "c"}: false,
	})

	test([]string{"a", "b", "c"}, []string{"a", "b", "c", "d"}, map[Keyer]bool{
		SampleKey{name: "a"}: true,
		SampleKey{name: "b"}: true,
		SampleKey{name: "c"}: true,
	})
}

func TestKeyMerge(t *testing.T) {
	test := func(base, target []string, e []Keyer) {
		b := SampleKeyHolder{KeyList: base}
		tgt := SampleKeyHolder{KeyList: target}

		merged := KeyMerge(b, tgt)

		if !reflect.DeepEqual(e, merged) {
			t.Errorf("merge (%s) and (%s) expected (%s) but (%s)", base, target, e, merged)
		}
	}

	test([]string{"a", "b", "c"}, []string{"b", "d"},
		[]Keyer{
			SampleKey{name: "a"},
			SampleKey{name: "b"},
			SampleKey{name: "c"},
			SampleKey{name: "d"},
		},
	)
}
