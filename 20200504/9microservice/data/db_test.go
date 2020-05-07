package data

import (
	"testing"

	_ "github.com/lib/pq"
)

func Test_open(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "check connection",
	},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Open()
			if err != nil{
				t.Error("failed to open db connection: ", err)
			}
		})
	}
}
