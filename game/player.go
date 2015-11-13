package game

type Player struct {
	Name      string //the name of the player
	X, Y      int    // the location of the player cam
	W, H      int    // the width and height of the player cam
	selection []interface{}
	Room      *Room //the room this player is in
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

func (p *Player) See() [][]int32 {
	// merge the views of all the selected units
	for _, u := range p.selection {
		l, ok := u.(Seer)
		if ok {
			return l.See(nil)
		}
	}
	view := make([][]int32, p.H)
	// no units selected- render with omniscent camera
	for i := 0; i < p.H; i++ {
		view[i] = make([]int32, p.W)
		for j := 0; j < p.W; j++ {
			view[i][j] = p.Room.GetTile(p.X+j, p.Y+i).Glyph
		}
	}
	return view
}
