# Cache Implementation in Go

This README provides a step-by-step explanation of the cache system implemented in Go. It includes an overview of the logic and detailed descriptions of each function.

---

## Overview of Logic

The cache system is based on a Least Recently Used (LRU) strategy. The primary components are:

- **Cache**: Maintains the queue and hash map.
- **Queue**: A doubly linked list for tracking the order of elements.
- **Hash**: A map for quick access to elements by their value.

### Key Features

- Efficient O(1) operations for adding, removing, and accessing elements.
- Dynamic cache size management.
- Functions for cache inspection and manipulation.

---

## Function Documentation

## Cache Functions

### `NewCache`

Creates a new cache with a specified maximum size.

#### **Parameters**:

- `size int`: The maximum number of elements the cache can hold.

#### **Returns**:

- `Cache`: An initialized cache instance.

#### **Usage**:

```go
c := NewCache(5) // Creates a cache with a maximum size of 5
```

---

### `Check`

Adds a value to the cache or refreshes its position if it already exists.

#### **Parameters**:

- `str string`: The value to be added or refreshed.

#### **Logic**:

1. If the value exists in the hash, remove it to refresh its position.
2. Otherwise, create a new node for the value.
3. Add the node to the front of the queue.
4. Update the hash map.

#### **Usage**:

```go
c.Check("abc")
```

---

### `Add`

Adds a new node to the front of the queue.

#### **Parameters**:

- `n *Node`: The node to be added.

#### **Logic**:

- Adjust the links in the queue to insert the node at the front.
- Increment the queue length.
- If the queue exceeds the maximum size, remove the least recently used node.

#### **Usage**:

```go
n := &Node{Val: "abc"}
c.Add(node)
```

---

### `Remove`

Removes a node from the cache.

#### **Parameters**:

- `n *Node`: The node to be removed.

#### **Logic**:

- Update the links of the node’s neighbors to exclude it.
- Decrement the queue length.
- Remove the value from the hash map.

#### **Returns**:

- `*Node`: The removed node.

#### **Usage**:

```go
r := c.Remove(node)
```

---

### `Display`

Prints the current state of the cache.

#### **Logic**:

- Traverse the queue from the head to the tail.
- Print each value in order.

#### **Usage**:

```go
c.Display()
```

---

### `Clear`

Removes all elements from the cache.

#### **Logic**:

- Reinitialize the queue.
- Clear the hash map.

#### **Usage**:

```go
c.Clear()
```

---

### `Exists`

Checks if a value exists in the cache.

#### **Parameters**:

- `str string`: The value to check.

#### **Returns**:

- `bool`: `true` if the value exists, `false` otherwise.

#### **Usage**:

```go
exists := c.Exists("abc")
```

---

### `Get`

Retrieves a value from the cache without removing it.

#### **Parameters**:

- `str string`: The value to retrieve.

#### **Returns**:

- `string`: The value if it exists.
- `bool`: `true` if the value exists, `false` otherwise.

#### **Usage**:

```go
val, found := c.Get("abc")
```

---

### `Length`

Returns the current number of elements in the cache.

#### **Returns**:

- `int`: The length of the cache.

#### **Usage**:

```go
length := c.Length()
```

---

### `Peek`

Returns the most recently used value without removing it.

#### **Returns**:

- `string`: The value of the most recently used item.
- `bool`: `true` if the cache is not empty, `false` otherwise.

#### **Usage**:

```go
val, found := c.Peek()
```

---

### `Update`

Updates an existing value in the cache.

#### **Parameters**:

- `old string`: The value to be updated.
- `new string`: The new value.

#### **Returns**:

- `bool`: `true` if the update was successful, `false` otherwise.

#### **Usage**:

```go
updated := c.Update("old", "new")
```

---

### `Keys`

Retrieves all keys currently in the cache.

#### **Returns**:

- `[]string`: A slice of all keys.

#### **Usage**:

```go
keys := c.Keys()
```

---

### `Values`

Retrieves all values currently in the cache.

#### **Returns**:

- `[]string`: A slice of all values.

#### **Usage**:

```go
values := c.Values()
```

---

### `IsFull`

Checks if the cache has reached its maximum capacity.

#### **Returns**:

- `bool`: `true` if the cache is full, `false` otherwise.

#### **Usage**:

```go
isFull := c.IsFull()
```

---

### `RemoveLeastRecentlyUsed`

Removes the least recently used item from the cache.

#### **Logic**:

1. If the cache is not empty, remove the last node in the queue.

#### **Usage**:

```go
c.RemoveLeastRecentlyUsed()
```

---

### `Promote`

Moves a specific value to the most recently used position.

#### **Parameters**:

- `val string`: The value to promote.

#### **Returns**:

- `bool`: `true` if the promotion was successful, `false` otherwise.

#### **Usage**:

```go
pmt := c.Promote("abc")
```

---

### `SetMaxSize`

Updates the maximum size of the cache.

#### **Parameters**:

- `size int`: The new maximum size.

#### **Logic**:

1. Update the size attribute.
2. Remove least recently used items until the cache size is within the new limit.

#### **Usage**:

```go
c.SetMaxSize(10)
```

---

## Queue Functions

### `Display`

#### **Parameters:**

- None.

#### **Logic:**

- Iterates through the queue and prints each element's value.

#### **Returns:**

- None.

#### **Usage:**

```go
q.Display()
```

---

### `Enqueue`

#### **Parameters:**

- `val string`: The value to add to the queue.

#### **Logic:**

- Adds a new node to the end of the queue.

#### **Returns:**

- None.

#### **Usage:**

```go
q.Enqueue("example")
```

### `Dequeue`

#### **Parameters:**

- None.

#### **Logic:**

- Removes and returns the value at the front of the queue.

#### **Returns:**

- `(string, bool)`: The value and a boolean indicating if the queue was non-empty.

#### **Usage:**

```go
val, success := q.Dequeue()
```

---

### `IsEmpty`

#### **Parameters:**

- None.

#### **Logic:**

- Checks if the queue is empty.

#### **Returns:**

- `bool`: True if empty, false otherwise.

#### **Usage:**

```go
empty := q.IsEmpty()
```

---

### `Peek`

#### **Parameters:**

- None.

#### **Logic:**

- Returns the value at the front of the queue without removing it.

#### **Returns:**

- `(string, bool)`: The value and a boolean indicating if the queue is non-empty.

#### **Usage:**

```go
val, success := q.Peek()
```

---

### `Clear`

#### **Parameters:**

- None.

#### **Logic:**

- Removes all elements from the queue.

#### **Returns:**

- None.

#### **Usage:**

```go
q.Clear()
```

---

### `Contains`

#### **Parameters:**

- `val string`: The value to check in the queue.

#### **Logic:**

- Checks if the value exists in the queue.

#### **Returns:**

- `bool`: True if the value exists, false otherwise.

#### **Usage:**

```go
exists := q.Contains("example")
```

---

### `Size`

#### **Parameters:**

- None.

#### **Logic:**

- Returns

---

### `Remove`

#### **Parameters:**

- `val string`: The value to check in the queue.

#### **Logic:**

- Removes a specific value from the queue.
- Finds the node containing the value, adjusts its neighbors' links, and decreases the length.

#### **Returns:**

- `bool`: True if the value exists, false otherwise.

#### **Usage:**

```go
rm := q.Remove("B")
```

---

### `Reverse`

#### **Parameters:**

- None.

#### **Logic:**

- Reverses the order of the elements in the queue.
- Swaps the left and right pointers of all nodes and exchanges the head and tail links.

#### **Returns:**

- None.

#### **Usage:**

```go
q.Reverse()
```

---

### `ToSlice`

#### **Parameters:**

- None.

#### **Logic:**

- Converts the queue into a slice of strings.
- Iterates through the queue and appends each value to a slice.

#### **Returns:**

- `[]string`: The slice containing all queue elements.

#### **Usage:**

```go
slice := q.ToSlice()
```
