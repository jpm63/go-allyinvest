package allyinvest

import (
	"testing"
)

func TestValidate(t *testing.T) {
	type x struct {
		IgnoreNotPointer   string  `json:"ignoreNotPointer"`
		IgnorePointer      *string `json:"ignorePointer"`
		NotPointer         string  `validate:"required"`
		Pointer            *int    `validate:"required"`
		OptionalNotPointer int     `validate:"optional"`
		OptionalPointer    *string `validate:"optional"`
	}

	s := "string"
	i := 1
	tests := []struct {
		name string
		in   x
		err  bool
	}{
		{
			name: "Fail",
			err:  true,
			in:   x{},
		},
		{
			name: "Pass",
			in: x{
				NotPointer: s,
				Pointer:    &i,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := validate(tt.in)
			if tt.err && got == nil {
				t.Errorf("got %v, want err", got)
			} else if !tt.err && got != nil {
				t.Errorf("got %v, want nil", got)
			}
		})
	}
}
