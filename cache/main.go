package main

import (
	"fmt"

	cache "cache/internal"
)

func main() {
	fmt.Println("Cache Operations Start")

	// Initialize the cache with a maximum size of 3
	cache := cache.NewCache(3)

	// Add elements to the cache
	cache.Check("apple")
	cache.Check("banana")
	cache.Check("cherry")
	cache.Display()

	// Access an existing item to move it to the front
	fmt.Println("Promote 'banana':", cache.Promote("banana"))
	cache.Display()

	// Add another element to force eviction
	cache.Check("date")
	cache.Display()

	// Check if an element exists in the cache
	exists := cache.Exists("banana")
	fmt.Println("Does 'banana' exist in the cache?:", exists)

	// Peek at the most recently used element
	if val, found := cache.Peek(); found {
		fmt.Println("Most recently used item:", val)
	} else {
		fmt.Println("Cache is empty.")
	}

	// Get a specific value without removing it
	if val, found := cache.Get("cherry"); found {
		fmt.Println("Value retrieved:", val)
	} else {
		fmt.Println("Value 'cherry' not found in cache.")
	}

	// Remove the least recently used item
	cache.RemoveLeastRecentlyUsed()
	cache.Display()

	// Clear the cache
	cache.Clear()
	cache.Display()

	// Add multiple items and test size adjustments
	cache.Check("kiwi")
	cache.Check("grape")
	cache.Check("lemon")
	cache.Display()

	// Increase cache size
	cache.SetMaxSize(5)
	cache.Check("mango")
	cache.Check("nectarine")
	cache.Display()

	// Decrease cache size, causing eviction
	cache.SetMaxSize(2)
	cache.Display()

	fmt.Println("Cache Operations End")
}
