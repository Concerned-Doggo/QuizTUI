package main

import (
	"fmt"

	"github.com/charmbracelet/huh"
	apicall "github.com/concerned-doggo/quizTUI/APICall"
	myform "github.com/concerned-doggo/quizTUI/MyForm"
)



func main() {
    fmt.Println("ready to go")

    var category string
    form := huh.NewForm(
        huh.NewGroup(
            // Ask the user for a base burger and toppings.
            huh.NewSelect[string]().
                Title("Choose your Category").
                Options(
                huh.NewOption("Linux", "linux"),
                huh.NewOption("Code", "code"),
                huh.NewOption("Bash", "bash"),
                huh.NewOption("MySql", "sql"),
                ).
                Value(&category),
            ),
        )
    form.Run()

    data := apicall.Question(category)

    myform.MakeForm(data)

}
