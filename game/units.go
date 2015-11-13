package game

import (
	//"github.com/beefsack/go-astar"
	"math"
)

// Stats is a struct containing attributes pertaining to a unit's combative abilities.
type Stats struct {
	Ap    int //the number of action points the unit has.
	Hp    int //the number of hit points the unit has.
	Speed int //the number of tiles the unit moves per action point.
}

// Unit is a structure containing statistics and methods common to all units.
type Unit struct {
	Vec3            //the position of this unit.
	Stats           //the current stats of this unit.
	baseStats Stats //the base (unaltered) stats of this unit.

	Name string //the unit's name
	Desc string //a description of the unit.

	dest     Vec3    //the location this unit is headed to.
	destDist float64 //the distance to this unit's destination.
	destDir  Vec3    //the direction the unit is moving (either -1 or 1)

	weak   int32 //bitmask of weaknesses.
	resist int32 //bitmask of resistances.

	spells []*Spell //the unit's spells or abiliites.
}

//MoveTo moves u to pos.
func (u *Unit) MoveTo(target Vec3) {
	u.dest = target
	distX := float64(target.X - u.X)
	distY := float64(target.Y - u.Y)
	distZ := float64(target.Z - u.Z)
	u.destDist = math.Sqrt(distX*distX + distY*distY + distZ*distZ)
}

//TakeDamage causes u to receive amount points of damage (pre-armor-mitigation).
func (u *Unit) TakeDamage(amount int, element int32) {
	u.Hp -= amount
}

//Update is called per-frame to update the unit.
func (u *Unit) Update() {
	//XXX use bresenham's
	dx := 1
	dy := 1
	if u.X < u.dest.X {
		dx = -1
	}
	if u.Y < u.dest.Y {
		dy = -1
	}
	for ap := u.Ap; ap > 0; ap-- {
		if u.X != u.dest.X {
			u.X += dx
		} else if u.Y != u.dest.Y {
			u.Y += dy
		} else {
			break
		}
	}
}

// Worm
func NewWorm(level int) *Unit {
	return &Unit{
		Name:      "Worm",
		Desc:      "A slimy creature",
		weak:      FIRE | ICE,
		resist:    DARK,
		baseStats: Stats{Hp: 10 + level, Ap: 3}}
}
