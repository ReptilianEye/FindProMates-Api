package utils

import (
	"slices"
	"strconv"
	"testing"
)

func TestMapTo(t *testing.T) {
	t.Run("TestMapToString", func(t *testing.T) {
		arr := []int{1, 2, 3}
		fn := func(i int) string { return strconv.Itoa(i) }
		want := []string{"1", "2", "3"}
		if got := MapTo(arr, fn); !slices.Equal(got, want) {
			t.Errorf("MapTo() = %v, want %v", got, want)
		}
	})
	t.Run("TestMapToInt", func(t *testing.T) {
		arr := []string{"1", "2", "3"}
		fn := func(s string) int { i, _ := strconv.Atoi(s); return i }
		want := []int{1, 2, 3}
		if got := MapTo(arr, fn); !slices.Equal(got, want) {
			t.Errorf("MapTo() = %v, want %v", got, want)
		}
	})
	t.Run("TestMapToStructToInt", func(t *testing.T) {
		type MyStruct struct {
			ID   int
			Name string
		}
		arr := []MyStruct{
			{ID: 1, Name: "John"},
			{ID: 2, Name: "Jane"},
			{ID: 3, Name: "Alice"},
		}
		fn := func(s MyStruct) int { return s.ID }
		want := []int{1, 2, 3}
		if got := MapTo(arr, fn); !slices.Equal(got, want) {
			t.Errorf("MapTo() = %v, want %v", got, want)
		}
	})

}

func TestAll(t *testing.T) {
	t.Run("TestAllTrue", func(t *testing.T) {
		arr := []int{1, 2, 3}
		fn := func(i int) bool { return i > 0 }
		if got := All(arr, fn); !got {
			t.Errorf("All() = %v, want %v", got, true)
		}
	})
	t.Run("TestAllFalse", func(t *testing.T) {
		arr := []int{1, 2, 3}
		fn := func(i int) bool { return i > 1 }
		if got := All(arr, fn); got {
			t.Errorf("All() = %v, want %v", got, false)
		}
	})
}
func TestAny(t *testing.T) {
	t.Run("TestAnyTrue", func(t *testing.T) {
		arr := []int{1, 2, 3}
		fn := func(i int) bool { return i > 2 }
		if got := Any(arr, fn); !got {
			t.Errorf("Any() = %v, want %v", got, true)
		}
	})
	t.Run("TestAnyFalse", func(t *testing.T) {
		arr := []int{1, 2, 3}
		fn := func(i int) bool { return i > 3 }
		if got := Any(arr, fn); got {
			t.Errorf("Any() = %v, want %v", got, false)
		}
	})
}

func TestTernary(t *testing.T) {
	t.Run("TestTernaryTrue", func(t *testing.T) {
		if got := Ternary(true, 1, 2); got != 1 {
			t.Errorf("Ternary() = %v, want %v", got, 1)
		}
	})
	t.Run("TestTernaryFalse", func(t *testing.T) {
		if got := Ternary(false, 1, 2); got != 2 {
			t.Errorf("Ternary() = %v, want %v", got, 2)
		}
	})
}

func TestElivis(t *testing.T) {
	t.Run("TestElivisTrue", func(t *testing.T) {
		a := "a"
		b := "b"
		if got := Elivis(&a, b); got != a {
			t.Errorf("Elivis() = %v, want %v", got, a)
		}
	})
	t.Run("TestElivisFalse", func(t *testing.T) {
		var a *string
		b := "b"
		if got := Elivis(a, b); got != b {
			t.Errorf("Elivis() = %v, want %v", got, b)
		}
	})
}

func TestKeys(t *testing.T) {
	t.Run("TestKeys", func(t *testing.T) {
		m := map[string]int{"a": 1, "b": 2, "c": 3}
		want := []string{"a", "b", "c"}
		got := Keys(m)
		for _, v := range got {
			if !slices.Contains(want, v) {
				t.Errorf("Value %v not found in %v", v, want)
			}
		}
	})
}

func TestValues(t *testing.T) {
	t.Run("TestValues", func(t *testing.T) {
		m := map[string]int{"a": 1, "b": 2, "c": 3}
		want := []int{1, 2, 3}
		got := Values(m)
		for _, v := range got {
			if !slices.Contains(want, v) {
				t.Errorf("Value %v not found in %v", v, want)
			}
		}
	})
}

func TestSet(t *testing.T) {
	t.Run("TestToSet", func(t *testing.T) {
		a := []int{1, 2, 3, 5, 4, 2, 1, 1}
		want := []int{1, 2, 3, 5, 4}
		mapper := func(i int) (string, error) { return strconv.Itoa(i), nil }
		got, err := ToSet(a, mapper)
		if err != nil {
			t.Errorf("ToSet() = %v, want %v", err, nil)
		}
		if len(got) != len(want) {
			t.Errorf("Set() = %v, want %v", got, want)
		}
		for _, v := range want {
			mapped, _ := mapper(v)
			if !got.Contains(mapped) {
				t.Errorf("Value %v not found in %v", v, want)
			}
		}
	})
	t.Run("TestSetToSlice", func(t *testing.T) {
		a := []int{1, 2, 3, 5, 4, 2, 1, 1}
		want := []int{1, 2, 3, 5, 4}
		mapper := func(i int) (string, error) { return strconv.Itoa(i), nil }
		s, err := ToSet(a, mapper)
		if err != nil {
			t.Errorf("ToSet() = %v, want %v", err, nil)
		}
		got, err := ToSlice(s, func(s string) (int, error) { return strconv.Atoi(s) }, "1", "5", "2")
		if err != nil {
			t.Errorf("ToSlice() = %v, want %v", err, nil)
		}
		if len(got) != len(want) {
			t.Errorf("Set() = %v, want %v", got, want)
		}
		if got[0] != 1 {
			t.Error("Value 1 should be first")
		}

		if got[1] != 5 {
			t.Error("Value 5 should be second")
		}
		if got[2] != 2 {
			t.Error("Value 2 should be third")
		}
		for _, v := range want {
			if !slices.Contains(got, v) {
				t.Errorf("Value %v not found in %v", v, want)
			}
		}
	})
	t.Run("TestSetUnion", func(t *testing.T) {
		a := []int{1, 2, 3, 5, 4, 2, 1, 1}
		b := []int{1, 1, 4, 10, 12}
		want := []int{1, 2, 3, 5, 4, 10, 12}
		mapper := func(i int) (string, error) { return strconv.Itoa(i), nil }
		as, _ := ToSet(a, mapper)
		bs, _ := ToSet(b, mapper)
		got := as.Union(bs)
		if len(got) != len(want) {
			t.Errorf("Set() = %v, want %v", got, want)
		}
		for _, v := range want {
			val, _ := mapper(v)
			if !got.Contains(val) {
				t.Errorf("Value %v not found in %v", v, want)
			}
		}
	})
	t.Run("TestSetIntersection", func(t *testing.T) {
		a := []int{1, 5, 2, 5, 1, 10, 2}
		b := []int{5, 2, 5, 2, 3, 4, 7, 9, 10}
		want := []int{5, 2, 10}
		mapper := func(i int) (string, error) { return strconv.Itoa(i), nil }
		as, _ := ToSet(a, mapper)
		bs, _ := ToSet(b, mapper)
		got := as.Intersection(bs)
		if len(got) != len(want) {
			t.Errorf("Set() = %v, want %v", got, want)
		}
		for _, v := range want {
			val, _ := mapper(v)
			if !got.Contains(val) {
				t.Errorf("Value %v not found in %v", v, want)
			}
		}
	})

}
