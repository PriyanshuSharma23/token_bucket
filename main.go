package main

import (
	"fmt"
	"time"

	"github.com/PriyanshuSharma23/token_bucket/bucket"
)

func main() {
	b := bucket.NewBucket(2, 10)

	fmt.Println(b.Check(4))
	fmt.Println(b.Size())

	time.Sleep(time.Second)

	fmt.Println(b.Check(8))
	fmt.Println(b.Size())

	fmt.Println(b.Check(8))
	fmt.Println(b.Size())
}
