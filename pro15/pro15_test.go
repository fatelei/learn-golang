package pro15

import "testing"
import "fmt"

func TestStack(t *testing.T) {
	stack := new(Stack)

	stack.push(1)

	if stack.ary[0] != 1 {
		t.Error("push failed")
	}

	cur := stack.pop()
	fmt.Println(cur)
	if cur != 1 {
		t.Error("pop failed")
	}
}
