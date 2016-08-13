package attr

import "testing"

func TestIntIsEmpty(t *testing.T) {
	test := func(v int64, e bool) {
		res := IntIsEmpty(v)

		if res != e {
			t.Errorf("(%d) expected empty (%t) but (%t)", v, e, res)
		}
	}

	test(0, true)
	test(1, false)
	test(-1, false)
}

func TestStringIsEmpty(t *testing.T) {
	test := func(v string, e bool) {
		res := StringIsEmpty(v)

		if res != e {
			t.Errorf("(%s) expected empty (%t) but (%t)", v, e, res)
		}
	}

	test("", true)
	test("hoge", false)
}

func TestFloatIsEmpty(t *testing.T) {
	test := func(v float64, e bool) {
		res := FloatIsEmpty(v)

		if res != e {
			t.Errorf("(%f) expected empty (%t) but (%t)", v, e, res)
		}
	}

	test(0.0, true)
	test(0.1, false)
	test(0.0000001, false)
}

func TestSliceIsEmpty(t *testing.T) {
	test := func(v interface{}, e bool) {
		res := SliceIsEmpty(v)

		if res != e {
			t.Errorf("(%v) expected empty (%t) but (%t)", v, e, res)
		}
	}

	test([]string{}, true)
}
