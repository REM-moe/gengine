package gengine

import "math"

// Vector3D represents a point or displacement in 3D space.
type Vector3D struct {
	X, Y, Z float64
}

// -----------------------------------------------------------------------------
// BASIC MATH (IN-PLACE MODIFIERS)
// These methods modify the existing vector. Use these for physics updates.
// -----------------------------------------------------------------------------

// Invert flips the direction of the vector.
func (v *Vector3D) Invert() {
	v.X = -v.X
	v.Y = -v.Y
	v.Z = -v.Z
}

// Add adds another vector to this one (v += other).
func (v *Vector3D) Add(other Vector3D) {
	v.X += other.X
	v.Y += other.Y
	v.Z += other.Z
}

// Subtract removes another vector from this one (v -= other).
func (v *Vector3D) Subtract(other Vector3D) {
	v.X -= other.X
	v.Y -= other.Y
	v.Z -= other.Z
}

// ScalarMultiply scales the vector by a number (v *= scale).
func (v *Vector3D) ScalarMultiply(scale float64) {
	v.X *= scale
	v.Y *= scale
	v.Z *= scale
}

// AddScaledVector adds a vector scaled by a value (v += other * scale).
// This is critical for the physics integration step.
// if the values were velocity and time then
// it would be position + displacement ( displacement = velocity * time)
func (v *Vector3D) AddScaledVector(other Vector3D, scale float64) {
	v.X += other.X * scale
	v.Y += other.Y * scale
	v.Z += other.Z * scale
}

// ComponentProductUpdate multiplies this vector by another vector component-wise.
// (v.x *= other.x, etc.)
func (v *Vector3D) ComponentProductUpdate(other Vector3D) {
	v.X *= other.X
	v.Y *= other.Y
	v.Z *= other.Z
}

// CrossProduct calculates the cross product and updates this vector.
func (v *Vector3D) CrossProduct(other Vector3D) {
	// We use temp variables to avoid overwriting X before we need it for Z.
	newX := (v.Y * other.Z) - (v.Z * other.Y)
	newY := (v.Z * other.X) - (v.X * other.Z)
	newZ := (v.X * other.Y) - (v.Y * other.X)

	v.X = newX
	v.Y = newY
	v.Z = newZ
}

// Normalize changes the vector to have a length of 1.
func (v *Vector3D) Normalize() {
	length := v.Magnitude()
	if length > 0 {
		inv := 1.0 / length
		v.X *= inv
		v.Y *= inv
		v.Z *= inv
	}
}

// -----------------------------------------------------------------------------
// PURE OPERATIONS (RETURN NEW VECTOR)
// These methods return a NEW vector and do not change the original.
// -----------------------------------------------------------------------------

// Added returns a new vector that is the sum (v + other).
func (v Vector3D) Added(other Vector3D) Vector3D {
	return Vector3D{
		X: v.X + other.X,
		Y: v.Y + other.Y,
		Z: v.Z + other.Z,
	}
}

// Subtracted returns a new vector that is the difference (v - other).
func (v Vector3D) Subtracted(other Vector3D) Vector3D {
	return Vector3D{
		X: v.X - other.X,
		Y: v.Y - other.Y,
		Z: v.Z - other.Z,
	}
}

// Scaled returns a new vector scaled by a number (v * s).
func (v Vector3D) Scaled(s float64) Vector3D {
	return Vector3D{
		X: v.X * s,
		Y: v.Y * s,
		Z: v.Z * s,
	}
}

// ComponentProduct returns a new vector that is the component-wise product.
func (v Vector3D) ComponentProduct(other Vector3D) Vector3D {
	return Vector3D{
		X: v.X * other.X,
		Y: v.Y * other.Y,
		Z: v.Z * other.Z,
	}
}

// Cross returns a new vector that is the cross product (v x other).
func (v Vector3D) Cross(other Vector3D) Vector3D {
	return Vector3D{
		X: (v.Y * other.Z) - (v.Z * other.Y),
		Y: (v.Z * other.X) - (v.X * other.Z),
		Z: (v.X * other.Y) - (v.Y * other.X),
	}
}

// -----------------------------------------------------------------------------
// QUERIES
// -----------------------------------------------------------------------------

// Magnitude returns the length of the vector.
func (v Vector3D) Magnitude() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

// SquareMagnitude returns the squared length (faster for comparisons).
func (v Vector3D) SquareMagnitude() float64 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

// DotProduct (Scalar Product) returns the dot product of two vectors.
// v . other = (x*x + y*y + z*z)
func (v Vector3D) DotProduct(other Vector3D) float64 {
	return v.X*other.X + v.Y*other.Y + v.Z*other.Z
}
