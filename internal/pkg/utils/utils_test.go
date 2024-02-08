package utils_test

import (
	"example/FindProMates-Api/internal/pkg/utils"
	"testing"
)

func TestMapTo(t *testing.T) {
	type args struct {
		arr []string
		fn  func(string) string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "TestMapTo",
			args: args{
				arr: []string{"a", "b", "c"},
				fn:  func(s string) string { return s + "1" },
			},
			want: []string{"a1", "b1", "c1"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := utils.MapTo(tt.args.arr, tt.args.fn); !utils.Equal(got, tt.want) {
				t.Errorf("MapTo() = %v, want %v", got, tt.want)
			}
		})
	}
}
