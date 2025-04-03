package engine

import "math"

type PropMap map[EntityProp]interface{}
type CBMap map[EntityCB]interface{}

func num(i interface{}) float32 {
	switch t := i.(type) {
	case int:
		return float32(t)
	case float32:
		return t
	case float64:
		return float32(t)
	default:
		return 0
	}
}

func dist(x1, y1, x2, y2 float32) float32 {
	return float32(math.Hypot(float64(x2-x1), float64(y2-y1)))
}
