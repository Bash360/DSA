package main

import (
	"fair-money/queue"
	"fmt"
)

func main() {
jobs:=queue.New[string]()
jobs.Enqueue("task start node")
jobs.Enqueue("npm install")
jobs.Enqueue("npm run")
fmt.Println(jobs.Peek())
}
