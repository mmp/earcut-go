// Earcut is a simple wrapper around the high-quality earcut polygon
// triangulation library from mapbox: https://github.com/mapbox/earcut.hpp.
package earcut

// #cgo CXXFLAGS: -std=c++14
// #include "earcut.h"
// #include <stdlib.h>
import "C"

import (
	"unsafe"
)

// Each vertex in both the original polygon and the returned triangles
// is represented using a Vertex.
type Vertex struct {
	P [2]float64
}

// Polygon represents a general polygon; the given vertices specify its
// extent. Currently, only solid polygons without holes can be specified,
// though earcut.hpp is able to handle holes; that functionality could be
// easily added to this wrapper if needed.
type Polygon struct {
	Vertices []Vertex
	// TODO: support holes, since earcut.hpp does...
}

// The result of the triangulation is returned as a slice of Triangles.
type Triangle struct {
	Vertices [3]Vertex
}

// Triangulate takes a Polygon and returns a slice of Triangles, representing
// the polygon's triangulation.
func Triangulate(p Polygon) []Triangle {
	// Vertex above is layout-compatible with the Vertex type in earcut.h,
	// so we can pass the pointer directly to C.Triangulate...
	v := (*C.struct_Vertex)(unsafe.Pointer(&p.Vertices[0]))
	ctris := C.Triangulate(v, C.int(len(p.Vertices)))

	// And since Triangle is layout-compatible as well, we can #yolo our
	// way stright to a slice.
	ts := unsafe.Slice((*Triangle)(unsafe.Pointer(ctris.tri)), ctris.n)
	tris := make([]Triangle, len(ts))
	copy(tris, ts)

	// Free the storage allocated in the C++ code.
	C.free(unsafe.Pointer(ctris.tri))

	return tris
}
