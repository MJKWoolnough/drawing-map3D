package map3D

import "github.com/MJKWoolnough/equaler"

type region3D struct {
	*region
	x int32
	y int32
	z int32
}

func (r *region3D) Extend(forward bool) {
	r.distance++
	if !forward {
		r.x--
	}
}

func (r *region3D) Get(theMap *Map3D, add int32, dir uint8) equaler.Equaler {
	if dir == 0 {
		return theMap.Get(r.x+add, r.y, r.z)
	} else if dir == 1 && r.prev != 2 {
		return theMap.Get(r.x+add, r.y+1, r.z)
	} else if dir == 2 && r.prev != 1 {
		return theMap.Get(r.x+add, r.y-1, r.z)
	} else if dir == 3 && r.prev != 4 {
		return theMap.Get(r.x+add, r.y, r.z+1)
	} else if dir == 4 && r.prev != 3 {
		return theMap.Get(r.x+add, r.y, r.z-1)
	}
	return &equaler.EFalse
}

func (r *region3D) Set(theMap *Map3D, add int32, toSet equaler.Equaler) {
	theMap.Set(r.x+add, r.y, r.z, toSet)
}

func (r *region3D) NewRegion(dir uint8, start int32) Region {
	x, y, z := r.x+start, r.y, r.z
	if dir == 1 {
		y++
	} else if dir == 2 {
		y--
	} else if dir == 3 {
		z++
	} else if dir == 4 {
		z--
	}
	next := &region3D{new(region), x, y, z}
	next.prev = dir
	r.SetNext(next)
	return next
}
