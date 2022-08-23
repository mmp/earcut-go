
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

extern void *NewPolygon();
extern void DeletePolygon(void *poly);
extern void AddRing(void *poly, Vertex *v, int s, int e);
extern Triangles Triangulate(void *poly, Vertex *v);

#ifdef __cplusplus
}
#endif // __cplusplus
