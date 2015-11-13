package game

import (
	"testing"
)

var u Unit

func TestFireball(t *testing.T) {
	var sp Fireball

	t.Log("testing fireball")
	sp.Cast(u)
}

func TestLight(t *testing.T) {
	var sp Light

	t.Log("testing light")
	sp.Cast(u)
}
