
#include "earcut.h"
#include "mapbox/earcut.hpp"

#include <stdlib.h>
#include <vector>
#include <array>

Triangles Triangulate(Vertex *v, int nv) {
    using EarPoint = std::array<double, 2>;
    std::vector<EarPoint> points(nv);
    for (int i = 0; i < nv; ++i)
        points[i] = EarPoint{v[i].p[0], v[i].p[1]};

    std::vector<std::vector<EarPoint>> polygon;
    polygon.push_back(points);

    std::vector<int> indices = mapbox::earcut<int>(polygon);

    Triangles t;
    t.n = indices.size() / 3;
    t.tri = (Triangle *)malloc(t.n * sizeof(Triangle));
    for (int i = 0; i < t.n; ++i) {
        for (int j = 0; j < 3; ++j) {
            int idx = indices[3*i + j];
            t.tri[i].v[j].p[0] = points[idx][0];
            t.tri[i].v[j].p[1] = points[idx][1];
        }
    }

    return t;
}
