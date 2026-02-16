package main

import (
	"math/rand"
)

type Firework struct {
	Particle

	Type int

	Age float64
}

type Payload struct {
	Type  int
	Count int
	Color float64
}

type FireworkRule struct {
	Type int // The ID this rule manages

	MinAge float64
	MaxAge float64

	MinVelocity Vector3D
	MaxVelocity Vector3D

	Damping float64

	Payloads []Payload
}

func (f *Firework) Update(duration float64) bool {
	f.Integrate(duration) // sets position and velocity

	f.Age -= 1

	return f.Age < 0 || f.Position.Y < 0
}

func (rule *FireworkRule) Create(f *Firework, parent *Firework) {
	f.Type = rule.Type

	f.Age = randomRange(rule.MinAge, rule.MaxAge)

	// Setup Physics
	if parent != nil {
		f.Position = parent.Position
		f.Velocity = parent.Velocity
	} else {
		f.Position = Vector3D{X: 0, Y: 0, Z: 0}
		f.Velocity = Vector3D{X: 0, Y: 0, Z: 0}
	}

	f.Velocity.Add(randomVector(rule.MinVelocity, rule.MaxVelocity))

	f.InverseMass = 1.0

	f.Damping = rule.Damping
	f.Accelaration = Vector3D{X: 0, Y: -9.8, Z: 0} // Gravity

	f.ClearAccumulator()
}

func randomRange(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

// Returns a random Vector between min and max
func randomVector(min, max Vector3D) Vector3D {
	return Vector3D{
		X: randomRange(min.X, max.X),
		Y: randomRange(min.Y, max.Y),
		Z: randomRange(min.Z, max.Z),
	}
}
