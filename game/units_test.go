package game

import (
	"testing"
)

func TestWorm(t *testing.T) {
	worm := NewWorm(2)
	t.Log("HP is ", worm.Hp)
}
