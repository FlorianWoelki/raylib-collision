package raylibcollision

import rl "github.com/gen2brain/raylib-go/raylib"

func CollideWithRects(rect rl.Rectangle, rects []rl.Rectangle) []rl.Rectangle {
	var hitList []rl.Rectangle
	for _, rec := range rects {
		if CollideWithRect(rect, rec) {
			hitList = append(hitList, rec)
		}
	}
	return hitList
}

type collisionDirection struct {
	Top    bool
	Bottom bool
	Right  bool
	Left   bool
}

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
