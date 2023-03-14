package area

import "github.com/go-gl/mathgl/mgl64"

// Vec3 represents a 3D area.
type Vec3 struct {
	// minX is the minimum X value.
	minX,
	// maxX is the maximum X value.
	maxX,
	// minY is the minimum Y value.
	minY,
	// maxY is the maximum Y value.
	maxY,
	// minZ is the minimum Z value.
	minZ,
	// maxZ is the maximum Z value.
	maxZ float64
}

// Max returns a mgl64.Vec3 with the maximum X, Y, and Z values.
func (v Vec3) Max() mgl64.Vec3 { return mgl64.Vec3{v.maxX, v.maxY, v.maxZ} }

// Min returns a mgl64.Vec3 with the minimum X, Y, and Z values.
func (v Vec3) Min() mgl64.Vec3 { return mgl64.Vec3{v.minX, v.minY, v.minZ} }

// NewVec3 returns a new Vec3 area with the minimum and maximum X, Y, and Z values.
func NewVec3(b1, b2 mgl64.Vec3) Vec3 {
	return Vec3{
		minX: minBound(b1.X(), b2.X()),
		maxX: maxBound(b1.X(), b2.X()),

		minY: minBound(b1.Y(), b2.Y()),
		maxY: maxBound(b1.Y(), b2.Y()),

		minZ: minBound(b1.Z(), b2.Z()),
		maxZ: maxBound(b1.Z(), b2.Z()),
	}
}

// Vec3Within returns true if the given mgl64.Vec3 is within the area.
func (v Vec3) Vec3Within(vec mgl64.Vec3) bool {
	return vec.X() > v.minX && vec.X() < v.maxX && vec.Y() > v.minY && vec.Y() < v.maxY && vec.Z() > v.minZ && vec.Z() < v.maxZ
}

// Vec3WithinOrEqual returns true if the given mgl64.Vec3 is within or equal to the bounds of the area.
func (v Vec3) Vec3WithinOrEqual(vec mgl64.Vec3) bool {
	return vec.X() >= v.minX && vec.X() <= v.maxX && vec.Y() >= v.minY && vec.Y() <= v.maxY && vec.Z() >= v.minZ && vec.Z() <= v.maxZ
}
