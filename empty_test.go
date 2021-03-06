package attr

import "testing"

func TestIsEmpty(t *testing.T) {
	test := func(v interface{}, e bool) {
		res := IsEmpty(v)

		if res != e {
			t.Errorf("(%v) expected empty (%t) but (%t)", v, e, res)
		}
	}

	test(0, true)
	test(1, false)
	test(-1, false)
	test(int64(0), true)
	test(int(1), false)
	test("", true)
	test("hoge", false)
	test(123456789123456789, false)
	test(0.0, true)
	test(0.1, false)
	test(0.0000001, false)
	test([]string{}, true)
	test([]string{"a"}, false)
	test([]int{1}, false)
	test(nil, true)
}

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
	test([]string{"a"}, false)
	test([]int{1}, false)
	test(nil, true)

	var s []int
	test(s, true)
}

func TestInvalidIsEmpty(t *testing.T) {
	test := func(v interface{}, e bool) {
		res := InvalidIsEmpty(v)

		if res != e {
			t.Errorf("(%v) expected empty (%t) but (%t)", v, e, res)
		}
	}

	test(nil, true)
	test(1, false)
}
