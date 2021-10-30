package raylibcollision

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func CollideWithRect(r1 rl.Rectangle, r2 rl.Rectangle) bool {
	if r1.X+r1.Width > r2.X && r1.X < r2.X+r2.Width && r1.Y+r1.Height > r2.Y && r1.Y < r2.Y+r2.Height {
		return true
	}
	return false
}
