# Raylib-Collision

## What is it?

Raylib-Collision is a simple 2D collision detection and resolution library.
Collision is a pretty straightforward task and easy to implement. However, there is no state-of-the-art library that detects the collision and, if you want to, handles it for you as well.

Raylib-Collision is not something like Box2D or other physics framework where you just need to outline the, e.g., tilemap. No, you can easily add polygons or rectangles for your tilemap, and this library does the rest for you.

## How to install?

`go get github.com/florianwoelki/raylib-collision`

## How to use it?

Using it is simple. You can look into the [examples section](https://github.com/FlorianWoelki/raylib-collision/tree/main/examples) for more information.

There are some functions you can use to implement your custom collision handling. For that, you can use, e.g., the `CollideWithRect` function:
```go
func CollideWithRect(rl.Rectangle, rl.Rectangle) bool
```

Otherwise, if you have the most basic collision handling, you can use the predefined `Move` function:
```go
func MoveRect(*rl.Rectangle, float32, float32, []rl.Rectangle) collisionDirection
```

Where the first parameter checks the collision of the passed array. The return value is a `collisionDirection`, which returns the direction where the collision happened.
```go
type collisionDirection struct {
	Top    bool
	Bottom bool
	Right  bool
	Left   bool
}
```
