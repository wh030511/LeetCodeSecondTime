package medium

import (
	"fmt"
	"testing"
)

func TestConvert(t *testing.T) {
	res := convert("PAYPALISHIRING", 3)
	fmt.Printf("res: %v", res)
}
