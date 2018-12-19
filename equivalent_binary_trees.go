package main

import ("golang.org/x/tour/tree"
        "fmt")


// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walker(t *tree.Tree, ch chan int) {
    if t.Left != nil {
        Walker(t.Left, ch)
    }
    ch <- t.Value 
    if t.Right != nil {
        Walker(t.Right,ch)
    }
}

func Walk(t *tree.Tree, ch chan int){
   Walker(t,ch)
   close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
    return true
}

func main() {
    ch := make(chan int)
    go Walk(tree.New(1), ch)
    for i := range ch {
        fmt.Println(i)
    }
}

