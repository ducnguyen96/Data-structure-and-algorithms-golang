## Traersal Quiz

1. If you know a solution is not far from the root of the tree.

- **Answer**: BFS
- **Explain**: Because it starts searching the closest nodes to the parents first

2. If the tree is very deep and solutions are rare.

- **Answer**: BFS
- **Explain**: DFS will take very long, because solutions are rare so it will go over and over.

3. If the tree is very wide

- **Answer**: DFS
- **Explain**: BFS will need to much memory

4. If solutions are frequent but located deep in the tree

- **Answer**: DFS
- **Explain**: Because solutions are frequent so DFS will faster than BFS

5. Determining whether a path exists between two nodes

- **Answer**: DFS
- **Explain**:

6. Finding the shortest path

- **Answer**: BFS
- **Explain**:

## BFS

## DFS

```go
//      9
//   4     20
// 1   6 15   170

// InOrder - [1, 4, 6, 9, 15, 20, 170]
// PreOrder - [9, 4, 1, 6, 20, 15, 170]
// PostOrder - [1, 6, 4, 15, 170, 20, 9]
```
