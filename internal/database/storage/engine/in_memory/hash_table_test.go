package in_memory

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSet(t *testing.T) {
	table := &HashTable{
		data: map[string]string{
			"key_1": "value_1",
			"key_2": "value_2",
		},
	}

	testCases := []struct {
		name  string
		key   string
		value string
	}{
		{name: "test set not existing key", key: "key_3", value: "new_value"},
		{name: "test set existing key", key: "key_2", value: "new_value"},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			table.Set(tc.key, tc.value)

			value, found := table.data[tc.key]
			require.Equal(t, tc.value, value)
			require.True(t, found)
		})
	}
}

func TestGet(t *testing.T) {
	table := &HashTable{
		data: map[string]string{
			"key_1": "value_1",
			"key_2": "value_2",
		},
	}

	testCases := []struct {
		name          string
		key           string
		expectedFound bool
		expectedValue string
	}{
		{name: "test get not existing key", key: "key_3", expectedFound: false, expectedValue: ""},
		{name: "test get existing key", key: "key_1", expectedFound: true, expectedValue: "value_1"},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			value, found := table.Get(tc.key)

			require.Equal(t, tc.expectedValue, value)
			require.Equal(t, tc.expectedFound, found)
		})
	}
}

func TestDel(t *testing.T) {
	table := &HashTable{
		data: map[string]string{
			"key_1": "value_1",
			"key_2": "value_2",
		},
	}

	testCases := []struct {
		name string
		key  string
	}{
		{name: "test del not existing key", key: "key_3"},
		{name: "test del existing key", key: "key_1"},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			table.Del(tc.key)

			_, found := table.data[tc.key]
			require.False(t, found)
			_, found = table.data["key_2"]
			require.True(t, found)
			if tc.key == "key_3" {
				_, found = table.data["key_1"]
				require.True(t, found)
			}
		})
	}
}
