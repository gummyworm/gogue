package game

import ()

const (
	FIRE = iota
	ICE
	LIGHT
	DARK
)

// Spell is a struct containing information about a magickal spell or ability.
type Spell struct {
	cost        int //the spell's MP cost.
	channelTime int //the number of action points required to channel.
	level       int //the level of the spell.
}

//Fireball.
type Fireball struct {
	Spell
}

func (sp *Fireball) Cast(target interface{}) {
	unit, ok := target.(Unit)
	if ok {
		unit.TakeDamage(sp.level, FIRE)
	}
}

//Light.
type Light struct {
	Spell
}

func (sp *Light) Cast(target interface{}) {
	unit, ok := target.(Unit)
	if ok {
		unit.TakeDamage(sp.level, LIGHT)
	}
}
