## Question 1: First Duplicated Number

```go
// Given an array = [2, 5, 1, 2, 3, 5, 1, 2, 4]
// It should return 2
```

```go
// Given an array = [2, 1, 1, 2, 3, 5, 1, 2, 4]
// It should return 1
```

```go
// Given an array = [2, 3, 4, 5]
// It should return nil
```

### My thoughts

- We will loop though the array to find
- For each element of array we should mark if this element is existed already or not
- We will use hash table to instant look up for each element
