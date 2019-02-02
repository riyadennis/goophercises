package lib

import (
	"testing"
)

type scenario struct {
	name        string
	fileName    string
	expectedErr string
}

func TestRead(t *testing.T) {
	scenarios := []scenario{
		{
			name:        "invalidFile",
			fileName:    "invalid.csv",
			expectedErr: "open invalid.csv: no such file or directory",
		},
		{
			name:        "validFile",
			fileName:    "problems.csv",
			expectedErr: "",
		},
	}

	f := func(t *testing.T, sc scenario) {
		cnt, err := Read(sc.fileName)
		if err == nil && cnt == "" {
			t.Error("empty file")
		}
		if err != nil && err.Error() != sc.expectedErr {
			t.Errorf("error got %v, want %v", err.Error(), sc.expectedErr)
		}
	}

	for _, sc := range scenarios {
		t.Run(sc.name, func(t *testing.T) {
			f(t, sc)
		})
	}
}
