package game

import ()

type Player struct {
	Selection []interface{}
}

func (p *Player) Cmd(r rune, target interface{}) {
	for _, u := range p.Selection {
		switch r {
		case 'a':
			a, ok := u.(Attacker)
			if ok {
				a.Attack(target)
			}
		case 'c':
			c, ok := u.(Caster)
			if ok {
				c.Cast(target)
			}
		case 'q':
			q, ok := u.(Quaffer)
			if ok {
				q.Quaff(target)
			}
		}
	}
}

func (p *Player) See() (ret [80][25]byte) {
	for _, u := range p.Selection {
		l, ok := u.(Seer)
		if ok {
			ret = l.See()
		}
	}
	return ret
}
