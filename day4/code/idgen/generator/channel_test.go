package generator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChanGeneratorCommon(t *testing.T) {
	chanGenerator := NewChanGenerator()
	expected := []int{1, 2, 3, 4, 5}
	for _, exp := range expected {
		if act := chanGenerator.Generate(); exp != act {
			t.Errorf("Expected: %d, Actual: %d", exp, act)
		}
	}
}

func TestChanGeneratorTestify(t *testing.T) {
	chanGenerator := NewChanGenerator()
	expected := []int{1, 2, 3, 4, 5}
	for _, exp := range expected {
		assert.Equal(t, exp, chanGenerator.Generate())
	}
}

func TestChanGeneratorImplements(t *testing.T) {
	chanGenerator := NewChanGenerator()
	assert.Implements(t, (*Generator)(nil), chanGenerator)
}
