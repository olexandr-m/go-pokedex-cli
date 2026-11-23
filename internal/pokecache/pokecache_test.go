package pokecache

import (
	"slices"
	"testing"
)

func TestAddGet(t *testing.T) {
	type testCase struct {
		key      string
		val      []byte
		expected []byte
	}

	cases := []testCase{
		{
			key:      "key1",
			val:      []byte("testdata"),
			expected: []byte("testdata"),
		},
		{
			key:      "key2",
			val:      []byte("moretestdata"),
			expected: []byte("moretestdata"),
		},
	}

	cache := NewCache()

	for _, c := range cases {
		cache.Add(c.key, c.val)
		actual, ok := cache.Get(c.key)
		if !ok {
			t.Errorf("Get(%q) ok = false; want true", c.key)
			continue
		}
		if !slices.Equal(actual, c.expected) {
			t.Errorf("Get(%q) = %q; want %q", c.key, actual, c.expected)
		}
	}

	for _, c := range cases {
		actual, ok := cache.Get(c.key)
		if !ok {
			t.Errorf("Get(%q) ok = false; want true", c.key)
			continue
		}
		if !slices.Equal(actual, c.expected) {
			t.Errorf("Get(%q) = %q; want %q", c.key, actual, c.expected)
		}
	}
}

func TestGetMissing(t *testing.T) {
	type testCase struct {
		key      string
		expected []byte
		found    bool
	}

	cases := []testCase{
		{
			key:      "https://example.com",
			expected: nil,
			found:    false,
		},
		{
			key:      "",
			expected: nil,
			found:    false,
		},
	}

	cache := NewCache()

	for _, c := range cases {
		actual, ok := cache.Get(c.key)
		if ok != c.found {
			t.Errorf("Get(%q) ok = %v; want %v", c.key, ok, c.found)
		}
		if !slices.Equal(actual, c.expected) {
			t.Errorf("Get(%q) = %v; want %v", c.key, actual, c.expected)
		}
	}
}

func TestReplaceValue(t *testing.T) {
	type testCase struct {
		key        string
		initialVal []byte
		updatedVal []byte
		expected   []byte
	}

	cases := []testCase{
		{
			key:        "key1",
			initialVal: []byte("initial"),
			updatedVal: []byte("updated"),
			expected:   []byte("updated"),
		},
		{
			key:        "key2",
			initialVal: []byte("first"),
			updatedVal: []byte("second"),
			expected:   []byte("second"),
		},
	}

	cache := NewCache()

	for _, c := range cases {
		cache.Add(c.key, c.initialVal)
		cache.Add(c.key, c.updatedVal)

		actual, ok := cache.Get(c.key)
		if !ok {
			t.Errorf("Get(%q) ok = false; want true", c.key)
			continue
		}
		if !slices.Equal(actual, c.expected) {
			t.Errorf("Get(%q) = %q; want %q", c.key, actual, c.expected)
		}
	}
}
