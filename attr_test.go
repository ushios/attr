package attr

import (
	"reflect"
	"sort"
	"testing"
)

func TestAttributesKeys(t *testing.T) {
	test := func(a Attributes, ek []Keyer) {
		list := a.Keys()

		sort.Sort(Keyers(ek))
		sort.Sort(Keyers(list))

		if !reflect.DeepEqual(list, ek) {
			t.Errorf("Attribute keys expected (%v) but (%v)", ek, list)
		}
	}

	a1 := Attributes{}
	a1[NewKey("hoge")] = "hoge"
	a1[NewKey("fuga")] = "fuga"

	test(a1, []Keyer{NewKey("hoge"), NewKey("fuga")})
}