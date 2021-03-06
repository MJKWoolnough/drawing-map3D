package map3D

import "github.com/MJKWoolnough/equaler"

type regionXY struct {
	*region
	x int32
	y int32
	z int32
}

func (r *regionXY) Extend(forward bool) {
	r.distance++
	if !forward {
		r.x--
	}
}

func (r *regionXY) Get(theMap *Map3D, add int32, dir uint8) equaler.Equaler {
	if dir == 0 {
		return theMap.Get(r.x+add, r.y, r.z)
	} else if dir == 1 && r.prev != 2 {
		return theMap.Get(r.x+add, r.y+1, r.z)
	} else if dir == 2 && r.prev != 1 {
		return theMap.Get(r.x+add, r.y-1, r.z)
	}
	return &equaler.EFalse
}

func (r *regionXY) Set(theMap *Map3D, add int32, toSet equaler.Equaler) {
	theMap.Set(r.x+add, r.y, r.z, toSet)
}

func (r *regionXY) NewRegion(dir uint8, start int32) Region {
	x, y, z := r.x+start, r.y, r.z
	if dir == 1 {
		y++
	} else if dir == 2 {
		y--
	}
	next := &regionXY{new(region), x, y, z}
	next.prev = dir
	r.SetNext(next)
	return next
}
