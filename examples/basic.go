package main

import (
	rlC "github.com/florianwoelki/raylib-collision"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	screenWidth := int32(900)
	screenHeight := int32(600)

	rl.InitWindow(screenWidth, screenHeight, "basic window")
	rl.SetTargetFPS(60)

	player := rl.NewRectangle(5, 5, 32, 64)
	box := rl.NewRectangle(float32(screenWidth/2), float32(screenHeight/2), 64, 64)

	for !rl.WindowShouldClose() {
		var dx float32 = 0
		var dy float32 = 0
		var speed float32 = 5
		if rl.IsKeyDown(rl.KeyA) {
			dx -= speed
		} else if rl.IsKeyDown(rl.KeyD) {
			dx += speed
		}
		if rl.IsKeyDown(rl.KeyW) {
			dy -= speed
		} else if rl.IsKeyDown(rl.KeyS) {
			dy += speed
		}

		rlC.MoveRect(&player, dx, dy, []rl.Rectangle{box})

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		// draw rectangles
		rl.DrawRectangle(int32(player.X), int32(player.Y), int32(player.Width), int32(player.Height), rl.Black)
		rl.DrawRectangle(int32(box.X), int32(box.Y), int32(box.Width), int32(box.Height), rl.Red)

		rl.DrawFPS(screenWidth-90, screenHeight-30)
		rl.EndDrawing()
	}

	rl.CloseWindow()
}
