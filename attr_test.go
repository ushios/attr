package attr

import (
	"reflect"
	"testing"
)

type SampleKeyHolder struct {
	KeyList []string
}

// Keys return key list
func (s SampleKeyHolder) Keys() []string {
	return s.KeyList
}

func TestKeyExists(t *testing.T) {
	test := func(keys []string, checkStr string, e bool) {
		kh := SampleKeyHolder{KeyList: keys}

		res := KeyExists(kh, checkStr)
		if res != e {
			t.Errorf("(%s) key exists (%s) result expected (%t) but (%t)", keys, checkStr, e, res)
		}
	}

	test([]string{"a", "b", "c"}, "a", true)
	test([]string{"a", "b", "c"}, "d", false)
	test([]string{"A", "b", "c"}, "a", false)
}

func TestKeyDiff(t *testing.T) {
	test := func(base, target []string, e map[string]bool) {
		res := KeyDiff(
			SampleKeyHolder{KeyList: base},
			SampleKeyHolder{KeyList: target},
		)

		if !reflect.DeepEqual(res, e) {
			t.Errorf("(%s) diff (%s) expected (%v) but (%v)", base, target, e, res)
		}

	}

	test([]string{"a", "b", "c"}, []string{"a", "b"}, map[string]bool{
		"a": true,
		"b": true,
		"c": false,
	})

	test([]string{"a", "b", "c"}, []string{"a", "b", "c", "d"}, map[string]bool{
		"a": true,
		"b": true,
		"c": true,
	})
}

func TestKeyMerge(t *testing.T) {
	test := func(base, target, e []string) {
		b := SampleKeyHolder{KeyList: base}
		tgt := SampleKeyHolder{KeyList: target}

		merged := KeyMerge(b, tgt)

		if !reflect.DeepEqual(e, merged) {
			t.Errorf("merge (%s) and (%s) expected (%s) but (%s)", base, target, e, merged)
		}
	}

	test([]string{"a", "b", "c"}, []string{"b", "d"}, []string{"a", "b", "c", "d"})
}
