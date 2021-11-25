package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "Adrian",
		Price: 1.00,
		SKU:   "avs-vvs-def",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
