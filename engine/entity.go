package engine

type Entity struct {
	X     float32
	Y     float32
	Shape Shape
	D     PropMap
	CB    CBMap
}

func (E *Entity) CheckCB(g *Game) {
	if f, ok := E.CB[CollidePL]; ok {
		if dist(g.Pl.X, g.Pl.Y, E.X, E.Y) <= num(E.D[R])+g.Pl.R {
			(f.(func(*Player)))(&g.Pl)
		}
	}
}
