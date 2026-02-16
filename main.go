package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// CONSTANTS FOR VIEWING
const (
	ScreenWidth  = 800
	ScreenHeight = 600
	Scale        = 10.0 // 1 meter = 10 pixels
)

func main() {

	rl.InitWindow(ScreenWidth, ScreenHeight, "Go Physics Engine - Fireworks")

	rl.SetTargetFPS(60)

	rules := []FireworkRule{
		{
			Type:   1, // THE ROCKET
			MinAge: 0.5, MaxAge: 1.5,
			MinVelocity: Vector3D{X: -5, Y: 25, Z: -5},
			MaxVelocity: Vector3D{X: 5, Y: 35, Z: 5},
			Damping:     0.99,
			Payloads: []Payload{
				{Type: 2, Count: 20},
			},
		},
		{
			Type:   2, // THE EXPLOSION
			MinAge: 0.5, MaxAge: 1.0,
			MinVelocity: Vector3D{X: -20, Y: -20, Z: -20},
			MaxVelocity: Vector3D{X: 20, Y: 20, Z: 20},
			Damping:     0.9,
		},
	}

	fireworks := []*Firework{}
	for !rl.WindowShouldClose() {

		if rl.IsKeyPressed(rl.KeySpace) {
			rocket := &Firework{}
			rules[0].Create(rocket, nil)
			rocket.Position = Vector3D{X: 0, Y: 0, Z: 0}
			fireworks = append(fireworks, rocket)
		}

		dt := 0.016 // Fixed time step (60 FPS)
		nextFrameFireworks := []*Firework{}

		for _, fw := range fireworks {
			isDead := fw.Update(dt)

			if isDead {
				// Explosion Logic
				ruleIndex := fw.Type - 1
				if ruleIndex < len(rules) && ruleIndex >= 0 {
					rule := rules[ruleIndex]
					for _, payload := range rule.Payloads {
						for i := 0; i < payload.Count; i++ {
							childRule := rules[payload.Type-1]
							spark := &Firework{}
							childRule.Create(spark, fw)
							nextFrameFireworks = append(nextFrameFireworks, spark)
						}
					}
				}
			} else {
				nextFrameFireworks = append(nextFrameFireworks, fw)
			}
		}
		fireworks = nextFrameFireworks

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		rl.DrawText("Press SPACE to Launch Firework", 10, 10, 20, rl.White)
		rl.DrawText(fmt.Sprintf("Active Particles: %d", len(fireworks)), 10, 30, 20, rl.Gray)

		for _, fw := range fireworks {
			// Convert Physics (Meters) to Screen (Pixels)
			// X: Center 0 is at ScreenWidth/2
			screenX := int32(ScreenWidth/2 + fw.Position.X*Scale)
			// Y: Physics Up is Screen Down, so we invert it
			screenY := int32(ScreenHeight - (fw.Position.Y * Scale))

			// Draw the particle
			rl.DrawCircle(screenX, screenY, 3.0, fw.Color)
		}

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
