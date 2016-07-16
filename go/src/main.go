package main

func compare(x interface{}, y interface{}) bool {
	if x.(int) < y.(int) {
		return true
	} else {
		return false
	}
}

func main() {
	tree := New(compare)

	tree.Insert(1)
	tree.Insert(2)
	tree.Insert(3)
}
