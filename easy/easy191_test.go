package easy

import (
	"fmt"
	"testing"
)

func TestEasy191(t *testing.T) {
	res := hammingWeight(00000000000000000000000010000000)
	fmt.Printf("res: %v", res)
}
