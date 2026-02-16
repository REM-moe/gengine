package main

import "math"

type Particle struct {
	Position Vector3D

	Velocity Vector3D

	Accelaration Vector3D

	// this is friction or drag that a particle needs
	Damping float64

	InverseMass float64

	ForceAccum Vector3D
}

// we update the position
// we update the velocity
func (p *Particle) Integrate(duration float64) {
	// 1. Sanity check: We can't integrate backward in time
	if duration <= 0 {
		return
	}

	// 2. Update Position
	// position = position + (velocity * duration)
	p.Position.AddScaledVector(p.Velocity, duration)

	// 3. Calculate Acceleration from Forces
	// Start with base acceleration (like gravity) [ 0, -g , 0]
	resultingAcc := p.Accelaration

	// Add the forces we accumulated (a = F/m)
	// resultingAcc += force * inverseMass
	resultingAcc.AddScaledVector(p.ForceAccum, p.InverseMass)

	// 4. Update Velocity
	// velocity = velocity + (acceleration * duration)
	p.Velocity.AddScaledVector(resultingAcc, duration)

	// 5. Apply Damping (Drag)
	// velocity *= damping^duration
	// We use math.Pow to make drag frame-rate independent
	p.Velocity.ScalarMultiply(math.Pow(p.Damping, duration))

}
