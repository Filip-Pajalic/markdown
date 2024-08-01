package markdown_test

import (
	"fmt"
	"testing"

	"github.com/Filip-Pajalic/markdown"
)

type Person struct {
	Name    string `markdown:"header"`
	Age     int    `markdown:"item,Age"`
	Email   string `markdown:"item,Email"`
	Address string `markdown:"item,Address"`
}

const (
	markdownText = `## John Doe
	- **Age**: 30
	- **Email**: john.doe@example.com
	- **Address**: 123 Main St, Anytown, USA
	`
)

func TestEncode(t *testing.T) {
	p := Person{
		Name:    "John Doe",
		Age:     30,
		Email:   "john.doe@example.com",
		Address: "123 Main St, Anytown, USA",
	}

	markdown, err := markdown.Encode(p)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(markdown)
}

func TestDecode(t *testing.T) {
	var person Person
	err := markdown.Decode(markdownText, &person)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Unmarshaled struct: %+v\n", person)

}
