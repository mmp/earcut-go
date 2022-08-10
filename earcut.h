
typedef struct Vertex {
    double p[2];
} Vertex;

typedef struct Triangle {
    Vertex v[3];
} Triangle;

typedef struct Triangles {
    Triangle *tri;
    int n;
} Triangles;

#ifdef __cplusplus
extern "C" {
#endif // __cplusplus

extern Triangles Triangulate(Vertex *v, int nv);

#ifdef __cplusplus
}
#endif // __cplusplus
