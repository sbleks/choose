package choose

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var DefaultPrompt = `#? `

func readline() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

/*
type Chooser[T any] interface {
	Choose() (T, error)
}
*/

type Choices[T any] []T

//From prompts terminal user to choose from one of the choices.
func (c Choices[T]) Choose() (T, error) {
	for i, v := range c {
		fmt.Printf("%v. %v\n", i+1, v)
	}

	for {
		fmt.Print(DefaultPrompt)
		resp := readline()
		n, _ := strconv.Atoi(resp)
		if 0 < n && n < len(c)+1 {
			return c[n-1], nil
		}
	}

}

func From[T any](choices []T) (T, error) {
	return Choices[T](choices).Choose()
}
