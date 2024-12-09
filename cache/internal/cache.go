package cache

type (
	Cache struct {
		Queue Queue
		Hash  Hash
		Size  int
	}
)

// Check adds a value to the cache or refreshes its position if it exists
func (c *Cache) Check(str string) {
	n := &Node{}

	if val, ok := c.Hash[str]; ok {
		n = c.Remove(val) // Remove it to refresh its position
	} else {
		n = &Node{Val: str} // Create a new node for the value
	}

	c.Add(n)        // Add the node to the cache
	c.Hash[str] = n // Update the hash map
}

// Add adds a new node to the front of the queue
func (c *Cache) Add(n *Node) {
	tmp := c.Queue.Head.Right // Get the current first node
	c.Queue.Head.Right = n    // Set the new node as the first node
	n.Left = c.Queue.Head     // Link the new node to the head
	n.Right = tmp             // Link the new node to the former first node
	tmp.Left = n              // Link the former first node to the new node

	c.Queue.Length++

	if c.Queue.Length > c.Size {
		c.Remove(c.Queue.Tail.Left)
	}
}

// Remove removes a node from the cache
func (c *Cache) Remove(n *Node) *Node {
	l := n.Left  // Get the left neighbor
	r := n.Right // Get the right neighbor

	r.Left = l  // Update right neighbor's left pointer
	l.Right = r // Update left neighbor's right pointer

	c.Queue.Length--

	delete(c.Hash, n.Val) // Remove the value from the hash map

	return n
}

// Display displays the current state of the cache
func (c *Cache) Display() {
	c.Queue.Display()
}

// Clear removes all elements from the cache
func (c *Cache) Clear() {
	c.Queue = NewQueue()
	c.Hash = Hash{}
}

// Exists checks if a value is in the cache
func (c *Cache) Exists(str string) bool {
	_, found := c.Hash[str]

	return found
}

// Get retrieves a value from the cache without removing it
func (c *Cache) Get(str string) (string, bool) {
	if node, found := c.Hash[str]; found {
		return node.Val, true
	}

	return "", false
}

// Length returns the current number of elements in the cache
func (c *Cache) Length() int {
	return c.Queue.Length
}

// Peek returns the most recently used value without removing it
func (c *Cache) Peek() (string, bool) {
	if c.Queue.Length == 0 {
		return "", false
	}

	return c.Queue.Head.Right.Val, true // Return the value of the first node
}

// Update updates an existing value in the cache
func (c *Cache) Update(old, new string) bool {
	if node, found := c.Hash[old]; found {
		c.Remove(node)       // Remove the old node
		n := &Node{Val: new} // Create a new node with the updated value
		c.Add(n)             // Add the new node to the cache
		c.Hash[new] = n      // Update the hash with the new value
		return true
	}

	return false
}

// Keys retrieves all keys currently in the cache
func (c *Cache) Keys() []string {
	keys := make([]string, 0, c.Length())
	n := c.Queue.Head.Right // Start from the first actual node

	for n != c.Queue.Tail {
		keys = append(keys, n.Val) // Append each value to the slice
		n = n.Right                // Move to the next node
	}

	return keys
}

// Values retrieves all values currently in the cache
func (c *Cache) Values() []string {
	return c.Keys()
}

// IsFull checks if the cache has reached its maximum capacity
func (c *Cache) IsFull() bool {
	return c.Queue.Length >= c.Size
}

// RemoveLeastRecentlyUsed removes the least recently used item from the cache
func (c *Cache) RemoveLeastRecentlyUsed() {
	if c.Queue.Length > 0 {
		c.Remove(c.Queue.Tail.Left) // Remove the last node
	}
}

// Promote moves a specific value to the most recently used position
func (c *Cache) Promote(val string) bool {
	if n, exists := c.Hash[val]; exists {
		c.Remove(n) // Remove it from its current position
		c.Add(n)    // Add it to the front

		return true
	}

	return false
}

// SetMaxSize updates the maximum size of the cache
func (c *Cache) SetMaxSize(size int) {
	c.Size = size

	for c.Queue.Length > c.Size {
		c.RemoveLeastRecentlyUsed() // Remove least recently used items
	}
}

// NewCache creates a new cache with a given maximum size
func NewCache(size int) Cache {
	return Cache{Queue: NewQueue(), Hash: Hash{}, Size: size} // Initialize with empty queue, hash, and size
}
