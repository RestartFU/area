package area

import "github.com/go-gl/mathgl/mgl64"

// Vec2 represents a 2D area.
type Vec2 struct {
	// minX is the minimum X value.`
	minX,
	// maxX is the maximum X value.
	maxX,
	// minY is the minimum Y value.
	minY,
	// maxY is the maximum Y value.
	maxY float64
}

// Max returns a mgl64.Vec2 with the maximum X and Y values.
func (v Vec2) Max() mgl64.Vec2 { return mgl64.Vec2{v.maxX, v.maxY} }

// Min returns a mgl64.Vec2 with the minimum X and Y values.
func (v Vec2) Min() mgl64.Vec2 { return mgl64.Vec2{v.minX, v.minY} }

// NewVec2 returns a new Vec2 area with the minimum and maximum X and Y values.
func NewVec2(b1, b2 mgl64.Vec2) Vec2 {
	return Vec2{
		minX: minBound(b1.X(), b2.X()),
		maxX: maxBound(b1.X(), b2.X()),

		minY: minBound(b1.Y(), b2.Y()),
		maxY: maxBound(b1.Y(), b2.Y()),
	}
}

// Vec2Within returns true if the given mgl64.Vec2 is within the area.
func (v Vec2) Vec2Within(vec mgl64.Vec2) bool {
	return vec.X() > v.minX && vec.X() < v.maxX && vec.Y() > v.minY && vec.Y() < v.maxY
}

// Vec3WithinXZ returns true if the given mgl64.Vec3 is within the area.
func (v Vec2) Vec3WithinXZ(vec mgl64.Vec3) bool {
	return vec.X() > v.minX && vec.X() < v.maxX && vec.Z() > v.minY && vec.Z() < v.maxY
}

// Vec2WithinOrEqual returns true if the given mgl64.Vec2 is within or equal to the bounds of the area.
func (v Vec2) Vec2WithinOrEqual(vec mgl64.Vec2) bool {
	return vec.X() >= v.minX && vec.X() <= v.maxX && vec.Y() >= v.minY && vec.Y() <= v.maxY
}

// Vec3WithinOrEqualXZ returns true if the given mgl64.Vec3 is within or equal to the bounds of the area.
func (v Vec2) Vec3WithinOrEqualXZ(vec mgl64.Vec3) bool {
	return vec.X() >= v.minX && vec.X() <= v.maxX && vec.Z() >= v.minY && vec.Z() <= v.maxY
}
