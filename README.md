# Package map3D
    import "/home/michael/Programming/Go/src/github.com/MJKWoolnough/drawing-map3D"


#TYPES

	type Map interface {
	    Get(int32, int32, int32) equaler.Equaler
	    Set(int32, int32, int32, equaler.Equaler) error
	}
Map allows for a custom data storage medium.

	type Map3D struct {
	    // contains filtered or unexported fields
	}

	func NewCustomMap(minX, minY, minZ, maxX, maxY, maxZ int32, data Map) *Map3D
NewCustomMap creates a new map with a custom storage medium.

	func NewMap3D(minX, minY, minZ, maxX, maxY, maxZ int32) *Map3D
NewMap3D creates a new map with the default storage medium.

	func (m *Map3D) Copy(r *Map3D) error
Copy will copy data from one map to the other, even if it has the same underlying data structure.

	func (m *Map3D) DrawLine(x1, y1, z1, x2, y2, z2 int32, block equaler.Equaler) error
DrawLine will draw a 3D line, based on Bresenham's line algorithm.

	func (m *Map3D) DrawLineFunc(x1, y1, z1, x2, y2, z2 int32, posFunc func(int32, int32, int32, int32) equaler.Equaler) error
DrawLineFunc will draw a 3D line but will use the given function to determine what is put at each coordinate.

	func (m *Map3D) DrawPlane(x1, y1, z1, x2, y2, z2, x3, y3, z3, x4, y4, z4 int32, block equaler.Equaler) error
DrawPlane will draw a plane with the given bounds.

	func (m *Map3D) DrawPlaneFunc(x1, y1, z1, x2, y2, z2, x3, y3, z3, x4, y4, z4 int32, posFunc func(int32, int32, int32, int32) equaler.Equaler) error
DrawPlaneFunc will draw a plane and use the given function to determine what will be put in each coordinate.

	func (m *Map3D) Fill(toSet equaler.Equaler)
Fill will simply set each coordinate within the bounds to the specified data

	func (m *Map3D) FillArea(a *Map3D)
FillArea will copy (repeatedly if necessary) the given data to fill the map.

	func (m *Map3D) FillFunc(eFunc func(int32, int32, int32) equaler.Equaler)
FillFun will use the given function to set each coordinate within the bounds.

	func (m *Map3D) FloodReplace3D(x, y, z int32, toSet equaler.Equaler)
FloodReplace3D will do a flood replace in all dimensions.

	func (m *Map3D) FloodReplaceXY(x, y, z int32, toSet equaler.Equaler)
FloodReplaceXY will do a flood replace using the XY plane.

	func (m *Map3D) FloodReplaceYZ(x, y, z int32, toSet equaler.Equaler)
FloodReplaceYZ will do a flood replace using the YZ plane.

	func (m *Map3D) FloodReplaceZX(x, y, z int32, toSet equaler.Equaler)
FloodReplaceZX will do a flood replace using the ZX plane.

	func (m *Map3D) Get(x, y, z int32) equaler.Equaler
Get returns the data in the given coordinates

	func (m *Map3D) GetBounds() (int32, int32, int32, int32, int32, int32)
GetBounds returns the minimum and maximum X, Y & Z values allowed.

	func (m *Map3D) PaintReplace(toSet equaler.Equaler, region Region)
PaintReplace will do a flood replace using the rules defined in the region.

	func (m *Map3D) Set(x, y, z int32, b equaler.Equaler) error
Set puts the data in the given coordinates

	func (m *Map3D) SubMap(minX, minY, minZ, maxX, maxY, maxZ int32) *Map3D
SubMap will produce a new map with the same underlying data structure but with new bounds.

	type Region interface {
	    Extend(bool)
	    Get(*Map3D, int32, uint8) equaler.Equaler
	    Set(*Map3D, int32, equaler.Equaler)
	    NewRegion(uint8, int32) Region
	    GetNext() Region
	    SetNext(Region)
	    Distance() uint32
	}
Region allows for filling based on custom rules


