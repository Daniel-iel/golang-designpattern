package main

import "fmt"

func main() {
	lib.InterateBooks(myBookCallback)

	lib.InterateBooks(func(b Book) error {
		fmt.Println("Book author:", b.Author)
		return nil
	})
}

func myBookCallback(b Book) error {
	fmt.Println("Book title:", b.Name)
	return nil
}
