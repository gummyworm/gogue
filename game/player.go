package game

type Player struct {
	selection []interface{}
}

func (p *Player) Cmd(action Rune, target interface{}) {
	for _, u := range(player.selection) {
	switch(action) {
		case 'l':
			l, ok := u.(Seer)
			if ok {
				l.See()
			}
		case 'a':
			a, ok := u.(Attacker)
			if ok {
				a.Attack()
			}
		case 'q':
			q, ok := u.(Quaffer)
			if ok {
				q.Quaff()
			}
		default:
	}
}

