Theo như đề bài:

- (1) The left subtree of a node contains only nodes with keys less than the node's key.
- (2) The right subtree of a node contains only nodes with keys greater than the node's key.
- Both the left and right subtrees must also be binary search trees.

Thì ta chỉ cần check 1 node có validate bằng cách kiểm tra điều kiện (1) và (2), 1 tree valid khi tất cả các nodes là valid, invalid khi có 1 node invalid.

- Vì quy luật check tất cả các nodes là giống nhau nên ta sẽ dùng đệ quy, để check đến nodes cuối cùng.
