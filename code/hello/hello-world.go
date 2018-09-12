package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	people := "People!"
	fmt.Println("Hello,", people)
	i18n := "هذا المؤتمر那次会议အဲဒီညီလာခံ"
	fmt.Println("Does", i18n, "work?") // in Arabic, Chinese, and Burmese?
}
