package gengine

import "math"

// for pointers in in place modifications the funtion names would be
// add()
// and for the ones that modify and return a new its - added()

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
// returns a float64
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

// Add adds another vector to this one (v = v + other)
// "v" is the receiver. "other" is the argument.
func (v *Vector3D) Add(other Vector3D) {
	v.X += other.X
	v.Y += other.Y
	v.Z += other.Z
}

// add another vector original vector ir modifying it
func (v *Vector3D) Substract(other Vector3D) {
	v.X -= other.X
	v.Y -= other.Y
	v.Z -= other.Z
}

// returns new vector
// this retunrs a new vector
func (v Vector3D) Added(other Vector3D) Vector3D {

	return Vector3D{
		X: v.X + other.X,
		Y: v.Y + other.Y,
		Z: v.Z + other.Z,
	}
}

// retunrs new vector
// this returns a new vector
func (v *Vector3D) Substracted(other Vector3D) Vector3D {
	return Vector3D{
		X: v.X - other.X,
		Y: v.Y - other.Y,
		Z: v.Z - other.Z,
	}
}

// component product
// does the component product or cross product and returns a new vector
func (v *Vector3D) ComponentProductUpdate(other Vector3D) Vector3D {
	return Vector3D{
		X: v.X * other.X,
		Y: v.Y * other.Y,
		Z: v.Z * other.Z,
	}
}

func (v *Vector3D) ComponentProduct(other Vector3D) {
	v.X *= other.X
	v.Y *= other.Y
	v.Z *= other.Z
}

/// scalar product
