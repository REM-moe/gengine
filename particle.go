package main

type Particle struct {
	Position Vector3D

	Velocity Vector3D

	Accelaration Vector3D

	// this is friction or drag that a particle needs
	Damping float64

	InverseMass float64
}
