package sets

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestNewSetInt(t *testing.T) {
	tests := []struct {
		name string
		want *SetInt
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, NewSetInt())
			if got := NewSetInt(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSetInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetInt_Add(t *testing.T) {
	type fields struct {
		Elements map[int]struct{}
	}
	type args struct {
		element int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{"positive", {{1: struct{}, 2: struct{}}}, {3}
		}
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			set := SetInt{
				Elements: tt.fields.Elements,
			}
			set.Add(tt.args.element)
		})
	}
}

func TestSetInt_Delete(t *testing.T) {
	type fields struct {
		Elements map[int]struct{}
	}
	type args struct {
		element int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			set := SetInt{
				Elements: tt.fields.Elements,
			}
			set.Delete(tt.args.element)
		})
	}
}

func TestSetInt_Difference(t *testing.T) {
	type fields struct {
		Elements map[int]struct{}
	}
	type args struct {
		otherSet SetInt
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *SetInt
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			set := SetInt{
				Elements: tt.fields.Elements,
			}
			if got := set.Difference(tt.args.otherSet); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Difference() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetInt_Intersection(t *testing.T) {
	type fields struct {
		Elements map[int]struct{}
	}
	type args struct {
		otherSet SetInt
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *SetInt
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			set := SetInt{
				Elements: tt.fields.Elements,
			}
			if got := set.Intersection(tt.args.otherSet); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Intersection() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetInt_PrintElements(t *testing.T) {
	type fields struct {
		Elements map[int]struct{}
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			set := SetInt{
				Elements: tt.fields.Elements,
			}
			set.PrintElements()
		})
	}
}

func TestSetInt_Union(t *testing.T) {
	type fields struct {
		Elements map[int]struct{}
	}
	type args struct {
		otherSet SetInt
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *SetInt
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			set := SetInt{
				Elements: tt.fields.Elements,
			}
			if got := set.Union(tt.args.otherSet); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Union() = %v, want %v", got, tt.want)
			}
		})
	}
}
