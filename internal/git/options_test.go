package git_test

import (
	"reflect"
	"testing"

	"github.com/shipengqi/commitizen/internal/git"
)

func TestOptionsCombineSignOff(t *testing.T) {
	tests := []struct {
		name     string
		signOff  bool
		expected []string
	}{
		{
			name:     "signoff disabled",
			signOff:  false,
			expected: []string{"commit", "-F", "/tmp/msg"},
		},
		{
			name:     "signoff enabled",
			signOff:  true,
			expected: []string{"commit", "-F", "/tmp/msg", "--signoff"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := git.NewOptions()
			o.SignOff = tt.signOff
			got := o.Combine("/tmp/msg")
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("Combine() = %v, want %v", got, tt.expected)
			}
		})
	}
}
