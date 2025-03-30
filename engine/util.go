package engine

type PropMap map[EntityProp]interface{}

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

	// switch i.(type) {
	// case int:
	// 	return float32(i.(int))
	// case float32:
	// 	return i.(float32)
	// case float64:
	// 	return float32(i.(float64))
	// default:
	// 	return 0
	// }
}
