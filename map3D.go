// Package map3D provides numerous three dimensional drawing functions.
package map3D

import (
	"fmt"
	"github.com/MJKWoolnough/equaler"
)

// Map allows for a custom data storage medium.
type Map interface {
	Get(int32, int32, int32) equaler.Equaler
	Set(int32, int32, int32, equaler.Equaler) error
}

type mapData map[int32]map[int32]map[int32]equaler.Equaler

func (mD *mapData) Get(x, y, z int32) equaler.Equaler {
	if (*mD)[y] == nil || (*mD)[y][z] == nil || (*mD)[y][z][x] == nil {
		return &equaler.EThis
	}
	return (*mD)[y][z][x]
}

func (mD *mapData) Set(x, y, z int32, b equaler.Equaler) error {
	if (*mD)[y] == nil {
		(*mD)[y] = make(map[int32]map[int32]equaler.Equaler)
	}
	if (*mD)[y][z] == nil {
		(*mD)[y][z] = make(map[int32]equaler.Equaler)
	}
	(*mD)[y][z][x] = b
	return nil
}

func newMapData() *mapData {
	a := mapData(make(map[int32]map[int32]map[int32]equaler.Equaler, 0))
	return &a
}

type Map3D struct {
	minX int32
	minY int32
	minZ int32
	maxX int32
	maxY int32
	maxZ int32
	data Map
}

// Get returns the data in the given coordinates
func (m *Map3D) Get(x, y, z int32) equaler.Equaler {
	if !m.inBounds(x, y, z) {
		return &equaler.EFalse
	}
	return m.data.Get(x, y, z)
}

// Set puts the data in the given coordinates
func (m *Map3D) Set(x, y, z int32, b equaler.Equaler) error {
	if !m.inBounds(x, y, z) {
		return fmt.Errorf("Coordinates not within map bounds")
	}
	return m.data.Set(x, y, z, b)
}

func (m *Map3D) inBounds(x, y, z int32) bool {
	if x >= m.minX && x < m.maxX && y >= m.minY && y < m.maxY && z >= m.minZ && z < m.maxZ {
		return true
	}
	return false
}

// GetBounds returns the minimum and maximum X, Y & Z values allowed.
func (m *Map3D) GetBounds() (int32, int32, int32, int32, int32, int32) {
	return m.minX, m.minY, m.minZ, m.maxX, m.maxY, m.maxZ
}

// Copy will copy data from one map to the other, even if it has the same underlying data structure.
func (m *Map3D) Copy(r *Map3D) error {
	if m == nil || r == nil {
		return fmt.Errorf("Nil map received")
	}
	minX, minY, minZ, maxX, maxY, maxZ := r.GetBounds()
	if m.maxX-m.minX != maxX-minX || m.maxY-m.minY != maxY-minY || m.maxZ-m.minZ != maxZ-minZ {
		return fmt.Errorf("Regions must be of same dimensions")
	}
	var err error
	for x := int32(0); x < m.maxX-m.minX; x++ {
		for y := int32(0); y < m.maxY-m.minY; y++ {
			for z := int32(0); z < m.maxZ-m.minZ; z++ {
				if e := m.Set(m.minX+x, m.minY+y, m.minZ+z, r.Get(minX+x, minY+y, minZ+z)); err != nil {
					err = e
				}
			}
		}
	}
	return err
}

// DrawLine will draw a 3D line, based on Bresenham's line algorithm.
func (m *Map3D) DrawLine(x1, y1, z1, x2, y2, z2 int32, block equaler.Equaler) error {
	return m.DrawLineFunc(x1, y1, z1, x2, y2, z2,
		func(i, j, k, l int32) equaler.Equaler {
			return block

		})
}

// DrawLineFunc will draw a 3D line but will use the given function to determine what is put at each coordinate.
func (m *Map3D) DrawLineFunc(x1, y1, z1, x2, y2, z2 int32, posFunc func(int32, int32, int32, int32) equaler.Equaler) error {
	if !m.inBounds(x1, y1, z1) || !m.inBounds(x2, y2, z2) {
		return fmt.Errorf("Point out of bounds")
	}
	pos := [6]int32{x1, y1, z1, x2, y2, z2}
	var (
		lengths, steps [3]int32
		maxLength      int32
	)
	for i := 0; i < 3; i++ {
		lengths[i] = pos[i+3] - pos[i]
		if lengths[i] > 0 {
			steps[i] = 1
		} else if lengths[i] < 0 {
			steps[i] = -1
			lengths[i] = -lengths[i]
		}
		if lengths[i] > maxLength {
			maxLength = lengths[i]
		}
	}
	errors := [3]int32{
		maxLength >> 1,
		maxLength >> 1,
		maxLength >> 1,
	}
	for i := int32(0); i <= maxLength; i++ {
		if err := m.Set(pos[0], pos[1], pos[2], posFunc(errors[0], errors[1], errors[2], maxLength)); err != nil {
			return err
		}
		for j := 0; j < 3; j++ {
			errors[j] -= lengths[j]
			if errors[j] < 0 {
				pos[j] += steps[j]
				errors[j] += maxLength
			}
		}
	}
	return nil
}

// DrawPlane will draw a plane with the given bounds.
func (m *Map3D) DrawPlane(x1, y1, z1, x2, y2, z2, x3, y3, z3, x4, y4, z4 int32, block equaler.Equaler) error {
	return m.DrawPlaneFunc(
		x1, y1, z1, x2, y2, z2, x3, y3, z3, x4, y4, z4,
		func(i, j, k, l int32) equaler.Equaler {
			return block

		})
}

//DrawPlaneFunc will draw a plane and use the given function to determine what will be put in each coordinate.
func (m *Map3D) DrawPlaneFunc(x1, y1, z1, x2, y2, z2, x3, y3, z3, x4, y4, z4 int32, posFunc func(int32, int32, int32, int32) equaler.Equaler) error {
	if !m.inBounds(x1, y1, z1) || !m.inBounds(x2, y2, z2) || !m.inBounds(x3, y3, z3) || !m.inBounds(x4, y4, z4) {
		return fmt.Errorf("Point out of bounds")
	}
	pos := [12]int32{x1, y1, z1, x3, y3, z3, x2, y2, z2, x4, y4, z4}
	var (
		lengths, steps [6]int32
		maxLength      int32
	)
	for i := 0; i < 6; i++ {
		lengths[i] = pos[i+6] - pos[i]
		if lengths[i] > 0 {
			steps[i] = 1
		} else if lengths[i] < 0 {
			steps[i] = -1
			lengths[i] = -lengths[i]
		}
		if lengths[i] > maxLength {
			maxLength = lengths[i]
		}
	}
	hml := maxLength >> 1
	errors := [6]int32{
		hml,
		hml,
		hml,
		hml,
		hml,
		hml,
	}
	// 	fmt.Println("Drawing plane\n	Pos:", pos, "\n	Steps:", steps, "\n	MaxLength:", maxLength)
	if err := m.drawPlaneFuncLine([6]int32{pos[0], pos[1], pos[2], pos[3], pos[4], pos[5]}, errors, maxLength, posFunc); err != nil {
		return err
	}
	for a := int32(0); a < maxLength; a++ {
		for b := 0; b < 3; b++ {
			changed := false
			errors[b] -= lengths[b]
			if errors[b] < 0 {
				changed = true
				pos[b] += steps[b]
				errors[b] += maxLength
			}
			c := b + 3
			errors[c] -= lengths[c]
			if errors[c] < 0 {
				changed = true
				pos[c] += steps[c]
				errors[c] += maxLength
			}
			if changed {
				if err := m.drawPlaneFuncLine([6]int32{pos[0], pos[1], pos[2], pos[3], pos[4], pos[5]}, errors, maxLength, posFunc); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (m *Map3D) drawPlaneFuncLine(pos, errors [6]int32, maxLength int32, posFunc func(int32, int32, int32, int32) equaler.Equaler) error {
	var (
		llengths, lsteps [3]int32
		llengthf         [3]int64
		lmaxLength       int32
		order            [3]uint8
	)
	maxNum := 0
	for i := 0; i < 3; i++ {
		llengths[i] = pos[i+3] - pos[i]
		llengthf[i] = int64(maxLength>>1) + int64(errors[i+3]-errors[i])
		if llengths[i] > 0 {
			lsteps[i] = 1
		} else if llengths[i] < 0 {
			lsteps[i] = -1
			llengths[i] = -llengths[i]
		}
		if llengths[i] > lmaxLength {
			lmaxLength = llengths[i]
			maxNum = i
		}
	}
	switch maxNum {
	case 0:
		order[0] = 0
		if llengths[1] > llengths[2] {
			order[1] = 1
			order[2] = 2
		} else {
			order[1] = 2
			order[2] = 1
		}
	case 1:
		order[0] = 1
		if llengths[0] > llengths[2] {
			order[1] = 0
			order[2] = 2
		} else {
			order[1] = 2
			order[2] = 0
		}
	case 2:
		order[0] = 2
		if llengths[0] > llengths[1] {
			order[1] = 0
			order[2] = 1
		} else {
			order[1] = 1
			order[2] = 0
		}
	}
	lerrors := [3]int32{
		lmaxLength >> 1,
		lmaxLength >> 1,
		lmaxLength >> 1,
	}
	lerrorf := [3]int64{
		int64(errors[0]) * int64(lmaxLength),
		int64(errors[1]) * int64(lmaxLength),
		int64(errors[2]) * int64(lmaxLength),
	}
	lml := int64(maxLength) * int64(lmaxLength)
	// 	fmt.Println("	Drawing line\n		Pos:", pos, "\n		Errors:", errors, "\n		MaxLength:", maxLength, "\n		LLengths:", llengths, "\n		LSteps:", lsteps, "\n		lMaxLength:", lmaxLength, "\n		lErrorf:", lerrorf, "\n		llengthf:", llengthf)
	if err := m.Set(pos[0], pos[1], pos[2], posFunc(errors[0], errors[1], errors[2], maxLength)); err != nil {
		return err
	}
	for i := int32(0); i < lmaxLength; i++ {
		for l := 0; l < 3; l++ {
			k := order[l]
			lerrors[k] -= llengths[k]
			lerrorf[k] -= llengthf[k]
			if lerrorf[k] >= lml {
				lerrorf[k] -= lml
				lerrors[k]--
			} else if lerrorf[k] < 0 {
				lerrorf[k] += lml
				lerrors[k]++
			}
			if lerrors[k] < 0 {
				pos[k] += lsteps[k]
				lerrors[k] += lmaxLength
				if err := m.Set(pos[0], pos[1], pos[2], posFunc(lerrors[0], lerrors[1], lerrors[2], lmaxLength)); err != nil {
					return err
				}
			}
		}
	}
	return m.Set(pos[3], pos[4], pos[5], posFunc(errors[3], errors[4], errors[5], maxLength))
}

// PaintReplace will do a flood replace using the rules defined in the region.
func (m *Map3D) PaintReplace(toSet equaler.Equaler, region Region) {
	for toReplace := region.Get(m, 0, 0); region != nil; region = region.GetNext() {
		if !region.Get(m, 0, 0).Equal(toReplace) {
			continue
		}
		regions := make([]Region, 6)
		if region.Get(m, -1, 0).Equal(toReplace) {
			regions[0] = region.NewRegion(0, -1)
			for regions[0].Get(m, -1, 0).Equal(toReplace) {
				regions[0].Extend(false)
			}
		}
		distance := int32(region.Distance())
		distance++
		if region.Get(m, distance, 0).Equal(toReplace) {
			regions[5] = region.NewRegion(0, distance)
			for d := int32(1); regions[5].Get(m, d, 0).Equal(toReplace); d++ {
				regions[5].Extend(true)
			}

		}
		for i := int32(0); i < distance; i++ {
			region.Set(m, i, toSet)
			for j := uint8(1); j < 5; j++ {
				if region.Get(m, i, j).Equal(toReplace) {
					if regions[j] == nil {
						regions[j] = region.NewRegion(j, i)
					} else {
						regions[j].Extend(true)
					}
				} else {
					regions[j] = nil
				}
			}
		}
	}
}

// FloodReplaceXY will do a flood replace using the XY plane.
func (m *Map3D) FloodReplaceXY(x, y, z int32, toSet equaler.Equaler) {
	m.PaintReplace(toSet, &regionXY{new(region), x, y, z})
}

// FloodReplaceYZ will do a flood replace using the YZ plane.
func (m *Map3D) FloodReplaceYZ(x, y, z int32, toSet equaler.Equaler) {
	m.PaintReplace(toSet, &regionYZ{new(region), x, y, z})
}

// FloodReplaceZX will do a flood replace using the ZX plane.
func (m *Map3D) FloodReplaceZX(x, y, z int32, toSet equaler.Equaler) {
	m.PaintReplace(toSet, &regionZX{new(region), x, y, z})
}

// FloodReplace3D will do a flood replace in all dimensions.
func (m *Map3D) FloodReplace3D(x, y, z int32, toSet equaler.Equaler) {
	m.PaintReplace(toSet, &region3D{new(region), x, y, z})
}

// FloodReplace3D will do a flood replace in all dimensions.
func (m *Map3D) Replace(replace, with equaler.Equaler) error {
	for i := m.minX; i < m.maxX; i++ {
		for j := m.minY; j < m.maxY; j++ {
			for k := m.minZ; k < m.maxZ; k++ {
				if b := m.Get(i, j, k); b != nil && b.Equal(replace) {
					if err := m.Set(i, j, k, with); err != nil {
						return err
					}
				}
			}
		}
	}
	return nil
}

// Fill will simply set each coordinate within the bounds to the specified data
func (m *Map3D) Fill(toSet equaler.Equaler) {
	for i := m.minX; i < m.maxX; i++ {
		for j := m.minY; j < m.maxY; j++ {
			for k := m.minZ; k < m.maxZ; k++ {
				m.Set(i, j, k, toSet)
			}
		}
	}
}

// FillArea will copy (repeatedly if necessary) the given data to fill the map.
func (m *Map3D) FillArea(a *Map3D) {
	y := a.minY
	for j := m.minY; j < m.maxY; j++ {
		x := a.minX
		for i := m.minX; i < m.maxX; i++ {
			z := a.minZ
			for k := m.minZ; k < m.maxZ; k++ {
				m.Set(i, j, k, a.Get(x, y, z))
				if z++; z == a.maxZ {
					z = a.minZ
				}
			}
			if x++; x == a.maxX {
				x = a.minX
			}
		}
		if y++; y == a.maxY {
			y = a.minY
		}
	}
}

// FillFun will use the given function to set each coordinate within the bounds.
func (m *Map3D) FillFunc(eFunc func(int32, int32, int32) equaler.Equaler) {
	mdx, mdy, mdz := m.maxX-m.minX, m.maxY-m.minY, m.maxZ-m.minZ
	for i := int32(0); i < mdx; i++ {
		mX := m.minX + i
		for j := int32(0); j < mdy; j++ {
			mY := m.minY + j
			for k := int32(0); k < mdz; k++ {
				mZ := m.minZ + j
				m.Set(mX, mY, mZ, eFunc(i, j, k))
			}
		}
	}
}

// SubMap will produce a new map with the same underlying data structure but with new bounds.
func (m *Map3D) SubMap(minX, minY, minZ, maxX, maxY, maxZ int32) *Map3D {
	if minX >= maxX || minY >= maxY || minZ >= maxZ {
		return nil
	}
	return &Map3D{minX, minY, minZ, maxX, maxY, maxZ, m.data}
}

// NewMap3D creates a new map with the default storage medium.
func NewMap3D(minX, minY, minZ, maxX, maxY, maxZ int32) *Map3D {
	if minX >= maxX || minY >= maxY || minZ >= maxZ {
		return nil
	}
	return &Map3D{minX, minY, minZ, maxX, maxY, maxZ, newMapData()}
}

// NewCustomMap creates a new map with a custom storage medium.
func NewCustomMap(minX, minY, minZ, maxX, maxY, maxZ int32, data Map) *Map3D {
	if minX >= maxX || minY >= maxY || minZ >= maxZ {
		return nil
	}
	return &Map3D{minX, minY, minZ, maxX, maxY, maxZ, data}
}
