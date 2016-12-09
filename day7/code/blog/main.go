package main

import b "./blog-service"

func main() {
	service := b.NewService("./theme/index.mustache", "./test.db", ":8080")
	service.Run()
}
