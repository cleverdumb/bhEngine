module example.com/test

go 1.23.3

replace example.com/engine => ../engine

require example.com/engine v0.0.0-00010101000000-000000000000

require (
	github.com/go-gl/gl v0.0.0-20231021071112-07e5d0ea2e71 // indirect
	github.com/go-gl/glfw/v3.3/glfw v0.0.0-20250301202403-da16c1255728 // indirect
)
