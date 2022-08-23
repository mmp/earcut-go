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

// Polygon represents a general multipolygon; given by rings of vertices
// (of any winding, courtesy of mapbox)
type Polygon struct {
	Rings [][]Vertex
}

// The result of the triangulation is returned as a slice of Triangles.
type Triangle struct {
	Vertices [3]Vertex
}

// Triangulate takes a Polygon and returns a slice of Triangles, representing
// the polygon's triangulation.
func Triangulate(p Polygon) []Triangle {
	k := 0
	vs := []Vertex{}
	ss := []int{}
	es := []int{}
	poly := C.NewPolygon()

	for _, ring := range p.Rings {
		s := k
		e := k + len(ring)
		k = e
		vs = append(vs, ring...)
		ss = append(ss, s)
		es = append(es, e)

		// Vertex above is layout-compatible with the Vertex type in earcut.h,
		// so we can pass the pointer directly to C.Triangulate...
		uvs := (*C.struct_Vertex)(unsafe.Pointer(&vs[0]))
		C.AddRing(poly, uvs, C.int(s), C.int(e))
	}

	uvs := (*C.struct_Vertex)(unsafe.Pointer(&vs[0]))
	ctris := C.Triangulate(poly, uvs)

	// And since Triangle is layout-compatible as well, we can #yolo our
	// way stright to a slice.
	ts := unsafe.Slice((*Triangle)(unsafe.Pointer(ctris.tri)), ctris.n)
	tris := make([]Triangle, len(ts))
	copy(tris, ts)

	// Free the storage allocated in the C++ code.
	C.free(unsafe.Pointer(ctris.tri))
	C.DeletePolygon(poly)

	return tris
}
