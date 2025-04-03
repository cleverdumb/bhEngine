package engine

func getDefaultGameConfig() map[ConfigureType]interface{} {
	return map[ConfigureType]interface{}{
		InputMode:   FollowMouse,
		PlayerSpd:   7,
		ScreenBound: BoundX | BoundY,
	}
}

type ConfigureType int

const (
	_           ConfigureType = iota
	InputMode                 // FollowMouse
	PlayerSpd                 // 7
	ScreenBound               // BoundX | BoundY
)

type ConfigureVal int

const (
	_      ConfigureVal = iota
	BoundX              // fixed
	BoundY              // fixed
	WASD
	FollowMouse
	Snap
)

type Shape int

const (
	_ Shape = iota
	Circle
)

type EntityProp int

const (
	_ EntityProp = iota
	R
)

type EntityCB int

const (
	_ EntityCB = iota
	CollidePL
)
