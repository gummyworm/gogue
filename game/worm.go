package game

type Worm struct {
	Unit
}

func (worm *Worm) Update() {
}

func (worm *Worm) See() (ret [][]byte) {
	render := Render(worm.X, worm.Y, 80, 25, 0)

	ret = make([][]byte, len(render))
	for i, col := range render {
		copy(ret[i], col)
	}
	return ret
}
