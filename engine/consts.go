package engine

func getDefaultGameConfig() map[ConfigureType]ConfigureVal {
	return map[ConfigureType]ConfigureVal{
		InputMode: FOLLOWMOUSE,
	}
}

type ConfigureType int

const (
	_ ConfigureType = iota
	InputMode
)

type ConfigureVal int

const (
	_ ConfigureVal = iota
	WASD
	FOLLOWMOUSE
)
