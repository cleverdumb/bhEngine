package engine

func getDefaultGameConfig() map[ConfigureType]interface{} {
	return map[ConfigureType]interface{}{
		InputMode: FollowMouse,
		PlayerSpd: float32(7),
	}
}

type ConfigureType int

const (
	_ ConfigureType = iota
	InputMode
	PlayerSpd
)

type ConfigureVal int

const (
	_ ConfigureVal = iota
	WASD
	FollowMouse
	Snap
)
