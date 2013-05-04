package map3D

import "github.com/MJKWoolnough/equaler"

type regionZX struct {
	*region
	x int32
	y int32
	z int32
}

func (r *regionZX) Extend(forward bool) {
	r.distance++
	if !forward {
		r.z--
	}
}

func (r *regionZX) Get(theMap *Map3D, add int32, dir uint8) equaler.Equaler {
	if dir == 0 {
		return theMap.Get(r.x, r.y, r.z+add)
	} else if dir == 1 && r.prev != 2 {
		return theMap.Get(r.x+1, r.y, r.z+add)
	} else if dir == 2 && r.prev != 1 {
		return theMap.Get(r.x-1, r.y, r.z+add)
	}
	return &equaler.EFalse
}

func (r *regionZX) Set(theMap *Map3D, add int32, toSet equaler.Equaler) {
	theMap.Set(r.x, r.y, r.z+add, toSet)
}

func (r *regionZX) NewRegion(dir uint8, start int32) Region {
	x, y, z := r.x, r.y, r.z+start
	if dir == 1 {
		x++
	} else if dir == 2 {
		x--
	}
	next := &regionZX{new(region), x, y, z}
	next.prev = dir
	r.SetNext(next)
	return next
}
