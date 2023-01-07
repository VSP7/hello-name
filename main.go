package main

import (
	"bufio"
	"fmt"
	"os"
)


func main() {

	fmt.Print("Введите ваше имя(отчество, фамилию по желанию): ")
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    name := scanner.Text() // запишем вcё введеное с пробелами

	fmt.Print("Hello "+name+"!")

}


