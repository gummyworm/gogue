package game

type Attacker interface {
	Attack(target interface{})
}

type Quaffer interface {
	Quaff(target interface{})
}

type Caster interface {
	Cast(target interface{})
}

type Seer interface {
	See() [80][25]byte
}
