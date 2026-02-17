package main

type ForceGenerator interface {
	UpdateForce(p *Particle, duration float64)
}

// ForceRegistration holds the link between ONE particle and ONE force.
type ForceRegistration struct {
	Particle  *Particle
	Generator ForceGenerator
}

// ForceRegistry holds all the active connections.
// * Holds the list of registrations.
type ForceRegistry struct {
	registrations []ForceRegistration
}

// addply the force to the particle
func (r *ForceRegistry) Add(p *Particle, fg ForceGenerator) {
	r.registrations = append(r.registrations, ForceRegistration{
		Particle:  p,
		Generator: fg,
	})
}

// UpdateForces loops through the list and applies ALL forces.
func (r *ForceRegistry) UpdateForces(duration float64) {
	for _, reg := range r.registrations {
		reg.Generator.UpdateForce(reg.Particle, duration)
	}
}

func (r *ForceRegistry) Clear() {
	r.registrations = []ForceRegistration{}
}

type GravityGenerator struct {
	Gravity Vector3D
}

// DragGenerator simulates air resistance.
type DragGenerator struct {
	K1 float64 // Velocity drag coefficient
	K2 float64 // Velocity squared drag coefficient
}

func (g *GravityGenerator) UpdateForce(p *Particle, duration float64) {
	// If object has infinite mass (InverseMass = 0), gravity doesn't move it.
	if p.InverseMass == 0 {
		return
	}

	// F = m * g
	// Since we store 1/m, we calculate m = 1 / InverseMass
	mass := 1.0 / p.InverseMass

	// Apply the force: Force = Gravity * Mass
	// We use Scaled() to get a new vector, then add it to the particle.
	force := g.Gravity.Scaled(mass)
	p.AddForce(force)
}

// UpdateForce applies drag force
func (d *DragGenerator) UpdateForce(p *Particle, duration float64) {
	force := p.Velocity

	// Calculate total drag coefficient
	dragCoeff := force.Magnitude()
	dragCoeff = d.K1*dragCoeff + d.K2*dragCoeff*dragCoeff

	// Calculate final force direction and magnitude
	force.Normalize()
	force.ScalarMultiply(-dragCoeff)

	p.AddForce(force)
}
