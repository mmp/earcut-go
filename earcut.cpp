
#include "earcut.h"
#include "mapbox/earcut.hpp"

#include <stdlib.h>
#include <vector>
#include <array>

using EarPoint = std::array<double, 2>;
using EarPoly = std::vector<std::vector<EarPoint>>;

void *NewPolygon() {
    EarPoly* poly = new EarPoly();
    return static_cast<void*>(poly);
}

void DeletePolygon(void *_poly) {
    EarPoly* poly = static_cast<EarPoly*>(_poly);
    delete poly;
}

void AddRing(void *_poly, Vertex *v, int s, int e) {
    EarPoly* poly = static_cast<EarPoly*>(_poly);
    std::vector<EarPoint> points(e-s);
    for (int i = s; i < e; ++i)
        points[i-s] = EarPoint{v[i].p[0], v[i].p[1]};

    poly->push_back(points);
}

Triangles Triangulate(void *_poly, Vertex *v) {
    EarPoly* poly = static_cast<EarPoly*>(_poly);
    std::vector<int> indices = mapbox::earcut<int>(*poly);

    Triangles t;
    t.n = indices.size() / 3;
    t.tri = (Triangle *)malloc(t.n * sizeof(Triangle));
    for (int i = 0; i < t.n; ++i) {
        for (int j = 0; j < 3; ++j) {
            int idx = indices[3*i + j];
            t.tri[i].v[j].p[0] = v[idx].p[0];
            t.tri[i].v[j].p[1] = v[idx].p[1];
        }
    }

    return t;
}
