package map3D

import "github.com/MJKWoolnough/equaler"

type regionYZ struct {
	*region
	x int32
	y int32
	z int32
}

func (r *regionYZ) Extend(forward bool) {
	r.distance++
	if !forward {
		r.y--
	}
}

func (r *regionYZ) Get(theMap *Map3D, add int32, dir uint8) equaler.Equaler {
	if dir == 0 {
		return theMap.Get(r.x, r.y+add, r.z)
	} else if dir == 1 && r.prev != 2 {
		return theMap.Get(r.x, r.y+add, r.z+1)
	} else if dir == 2 && r.prev != 1 {
		return theMap.Get(r.x, r.y+add, r.z-1)
	}
	return &equaler.EFalse
}

func (r *regionYZ) Set(theMap *Map3D, add int32, toSet equaler.Equaler) {
	theMap.Set(r.x, r.y+add, r.z, toSet)
}

func (r *regionYZ) NewRegion(dir uint8, start int32) Region {
	x, y, z := r.x, r.y+start, r.z
	if dir == 1 {
		z++
	} else if dir == 2 {
		z--
	}
	next := &regionYZ{new(region), x, y, z}
	next.prev = dir
	r.SetNext(next)
	return next
}
