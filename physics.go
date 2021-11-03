package raylibcollision

import rl "github.com/gen2brain/raylib-go/raylib"

// CollideWithRects is a utility function and checks whether a rectangle collided with some rectangle in the array or not.
// It returns a hit slice that includes all the collided rectangles.
func CollideWithRects(rect rl.Rectangle, rects []rl.Rectangle) []rl.Rectangle {
	var hitSlice []rl.Rectangle
	for _, rec := range rects {
		if CollideWithRect(rect, rec) {
			hitSlice = append(hitSlice, rec)
		}
	}
	return hitSlice
}

type collisionDirection struct {
	Top    bool
	Bottom bool
	Right  bool
	Left   bool
}

// MoveRect will handle the collision and return the direction struct, which includes boolean values of the directions.
// The function checks the future movement whether it collided with some rectangle in the passed slice.
func MoveRect(rect *rl.Rectangle, dx, dy float32, rects []rl.Rectangle) collisionDirection {
	collision := collisionDirection{Top: false, Bottom: false, Right: false, Left: false}

	rect.X += dx
	hitList := CollideWithRects(*rect, rects)
	for _, tile := range hitList {
		if dx > 0 {
			rect.X = tile.X - rect.Width
			collision.Right = true
		} else if dx < 0 {
			rect.X = tile.X + tile.Width
			collision.Left = true
		}
	}

	rect.Y += dy
	hitList = CollideWithRects(*rect, rects)
	for _, tile := range hitList {
		if dy > 0 {
			rect.Y = tile.Y - rect.Height
			collision.Bottom = true
		} else if dy < 0 {
			rect.Y = tile.Y + tile.Height
			collision.Top = true
		}
	}

	return collision
}
