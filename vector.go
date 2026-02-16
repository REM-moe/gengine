package gengine

import "math"

// Vector3D represents a point or displacement in 3D space.
type Vector3D struct {
	X, Y, Z float64
}

// this retunrs a new struct
func NewVector3D(x, y, z float64) Vector3D {
	return Vector3D{X: x, Y: y, Z: z}
}

// Invert flips the direction of the vector in-place.
// We use a pointer receiver (*Vector3D) because we are changing the data.
func (v *Vector3D) Invert() {
	v.X = -v.X
	v.Y = -v.Y
	v.Z = -v.Z
}

// AddScaledVector is a classic Millington function: v += other * scale
func (v *Vector3D) AddScaledVector(other Vector3D, scale float64) {
	v.X += other.X * scale
	v.Y += other.Y * scale
	v.Z += other.Z * scale
}

// in place we scalar multiply
func (v *Vector3D) ScalarMultiply(scale float64) {
	v.X *= scale
	v.Y *= scale
	v.Z *= scale
}

// returns a new vector
func (v Vector3D) Scaled(s float64) Vector3D {
	return Vector3D{
		X: v.X * s,
		Y: v.Y * s,
		Z: v.Z * s,
	}
}

// magnitude of the vector root( x2 + y2 + z2 )
func (v Vector3D) Magnitude() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

// squareMagnitude is the faster version for comparisons
// if we just wanna compare disctance this is enough
func (v Vector3D) SquareMagnitude() float64 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

// normalize changes the vector to have a length of 1
func (v *Vector3D) Normalize() {
	length := v.Magnitude()
	if length > 0 {
		inv := 1.0 / length
		v.X *= inv
		v.Y *= inv
		v.Z *= inv
	}
}
