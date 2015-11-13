package game

type Caster interface {
	Cast(interface{})
}

type Attacker interface {
	Attack(interface{})
}

type Quaffer interface {
	Quaff(interface{})
}
