package map3D

import "github.com/MJKWoolnough/equaler"

// Region allows for filling based on custom rules
type Region interface {
	Extend(bool)
	Get(*Map3D, int32, uint8) equaler.Equaler
	Set(*Map3D, int32, equaler.Equaler)
	NewRegion(uint8, int32) Region
	GetNext() Region
	SetNext(Region)
	Distance() uint32
}

type region struct {
	distance uint32
	next     Region
	prev     uint8
}

func (r *region) Distance() uint32 {
	return r.distance
}

func (r *region) GetNext() Region {
	return r.next
}

func (r *region) SetNext(reg Region) {
	if r.next != nil {
		reg.SetNext(r.next)
	}
	r.next = reg
}
