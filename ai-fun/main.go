package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()

	// Perform addition, squaring, concatenation, and reversing
	adderResult := add(ctx, 38, 4)
	squaredResult := square(ctx, adderResult)
	concatResult := concatenate(ctx, "Hello", "World")
	reversedResult := reverse(ctx, concatResult)

	// Print final results
	fmt.Printf("Squared result: %d\n", squaredResult)
	fmt.Printf("Concatenated and reversed result: \"%s\"\n", reversedResult)
}
