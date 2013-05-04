# map3D
--
    import "github.com/MJKWoolnough/drawing-map3D"


## Usage

#### type Map

```go
type Map interface {
	Get(int32, int32, int32) equaler.Equaler
	Set(int32, int32, int32, equaler.Equaler) error
}
```

Map allows for a custom data storage medium.

#### type Map3D

```go
type Map3D struct {
}
```


#### func  NewCustomMap

```go
func NewCustomMap(minX, minY, minZ, maxX, maxY, maxZ int32, data Map) *Map3D
```
NewCustomMap creates a new map with a custom storage medium.

#### func  NewMap3D

```go
func NewMap3D(minX, minY, minZ, maxX, maxY, maxZ int32) *Map3D
```
NewMap3D creates a new map with the default storage medium.

#### func (*Map3D) Copy

```go
func (m *Map3D) Copy(r *Map3D) error
```
Copy will copy data from one map to the other, even if it has the same
underlying data structure.

#### func (*Map3D) DrawLine

```go
func (m *Map3D) DrawLine(x1, y1, z1, x2, y2, z2 int32, block equaler.Equaler) error
```
DrawLine will draw a 3D line, based on Bresenham's line algorithm.

#### func (*Map3D) DrawLineFunc

```go
func (m *Map3D) DrawLineFunc(x1, y1, z1, x2, y2, z2 int32, posFunc func(int32, int32, int32, int32) equaler.Equaler) error
```
DrawLineFunc will draw a 3D line but will use the given function to determine
what is put at each coordinate.

#### func (*Map3D) DrawPlane

```go
func (m *Map3D) DrawPlane(x1, y1, z1, x2, y2, z2, x3, y3, z3, x4, y4, z4 int32, block equaler.Equaler) error
```
DrawPlane will draw a plane with the given bounds.

#### func (*Map3D) DrawPlaneFunc

```go
func (m *Map3D) DrawPlaneFunc(x1, y1, z1, x2, y2, z2, x3, y3, z3, x4, y4, z4 int32, posFunc func(int32, int32, int32, int32) equaler.Equaler) error
```
DrawPlaneFunc will draw a plane and use the given function to determine what
will be put in each coordinate.

#### func (*Map3D) Fill

```go
func (m *Map3D) Fill(toSet equaler.Equaler)
```
Fill will simply set each coordinate within the bounds to the specified data

#### func (*Map3D) FillArea

```go
func (m *Map3D) FillArea(a *Map3D)
```
FillArea will copy (repeatedly if necessary) the given data to fill the map.

#### func (*Map3D) FillFunc

```go
func (m *Map3D) FillFunc(eFunc func(int32, int32, int32) equaler.Equaler)
```
FillFun will use the given function to set each coordinate within the bounds.

#### func (*Map3D) FloodReplace3D

```go
func (m *Map3D) FloodReplace3D(x, y, z int32, toSet equaler.Equaler)
```
FloodReplace3D will do a flood replace in all dimensions.

#### func (*Map3D) FloodReplaceXY

```go
func (m *Map3D) FloodReplaceXY(x, y, z int32, toSet equaler.Equaler)
```
FloodReplaceXY will do a flood replace using the XY plane.

#### func (*Map3D) FloodReplaceYZ

```go
func (m *Map3D) FloodReplaceYZ(x, y, z int32, toSet equaler.Equaler)
```
FloodReplaceYZ will do a flood replace using the YZ plane.

#### func (*Map3D) FloodReplaceZX

```go
func (m *Map3D) FloodReplaceZX(x, y, z int32, toSet equaler.Equaler)
```
FloodReplaceZX will do a flood replace using the ZX plane.

#### func (*Map3D) Get

```go
func (m *Map3D) Get(x, y, z int32) equaler.Equaler
```
Get returns the data in the given coordinates

#### func (*Map3D) GetBounds

```go
func (m *Map3D) GetBounds() (int32, int32, int32, int32, int32, int32)
```
GetBounds returns the minimum and maximum X, Y & Z values allowed.

#### func (*Map3D) PaintReplace

```go
func (m *Map3D) PaintReplace(toSet equaler.Equaler, region Region)
```
PaintReplace will do a flood replace using the rules defined in the region.

#### func (*Map3D) Set

```go
func (m *Map3D) Set(x, y, z int32, b equaler.Equaler) error
```
Set puts the data in the given coordinates

#### func (*Map3D) SubMap

```go
func (m *Map3D) SubMap(minX, minY, minZ, maxX, maxY, maxZ int32) *Map3D
```
SubMap will produce a new map with the same underlying data structure but with
new bounds.

#### type Region

```go
type Region interface {
	Extend(bool)
	Get(*Map3D, int32, uint8) equaler.Equaler
	Set(*Map3D, int32, equaler.Equaler)
	NewRegion(uint8, int32) Region
	GetNext() Region
	SetNext(Region)
	Distance() uint32
}
```

Region allows for filling based on custom rules
