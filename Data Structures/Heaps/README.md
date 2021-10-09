## Heaps
- Good for implementing priority queues.
- Normal queues: first in first out
- Priority queues: take out the highest priority first
- Can express as a complete tree (every level of tree are full - expect lowest level, but all nodes to the left)
- Largest key --> root, parents have higher key than children
- BTS: heap stored as an array
- Parent: (i-1)/2
- Left child: 2i + 1
- Right child: 2i + 2

## Insert and Heapify
### Insert
- add at bottom of the right of the tree
- after insert --> heapify bottom up to rearrange to keep heap property(parents have higher key than children)

### Extract max
- get the root out
- put last on root
- heapify top down

## Time complexity
- Take out first index, add at the end are **Not Dominant**
- Heap insert & extract: O(log n)