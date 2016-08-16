package attr

import (
	"reflect"
	"sort"
	"testing"
)

func TestAttributesKeys(t *testing.T) {
	test := func(a KeyHolder, ek []Keyer) {
		list := a.Keys()

		sort.Sort(KeyerList(ek))
		sort.Sort(list)

		if !reflect.DeepEqual(list, ek) {
			t.Errorf("Attribute keys expected (%v) but (%v)", ek, list)
		}
	}

	a1 := Attributes{}
	a1[NewKey("hoge")] = "hoge"
	a1[NewKey("fuga")] = "fuga"

	test(&a1, []Keyer{NewKey("hoge"), NewKey("fuga")})
}
