package main

import (
	"reflect"
	"testing"
)

func Test_searchAnagram(t *testing.T) {
	tests := []struct {
		name string
		arr  []string
		want map[string][]string
	}{
		{
			name: "Test 1",
			arr: []string{"слиток",
				"автобус",
				"пятка",
				"Столик",
				"Столик",
				"тяпка"},
			want: map[string][]string{
				"пятка":  {"пятка", "тяпка"},
				"слиток": {"слиток", "столик"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchAnagram(tt.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("searchAnagram() = %v, want %v", got, tt.want)
			}
		})
	}
}
