package game

type Player struct {
	selection []interface{}
}

func (p *Player) Cmd(action rune, target interface{}) {
	for _, u := range p.selection {
		switch action {
		case 'l':
			l, ok := u.(Seer)
			if ok {
				l.See(target)
			}
		case 'a':
			a, ok := u.(Attacker)
			if ok {
				a.Attack(target)
			}
		case 'q':
			q, ok := u.(Quaffer)
			if ok {
				q.Quaff(target)
			}
		default:
		}
	}
}
