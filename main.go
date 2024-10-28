package main

import (
	"fmt"
	"strconv"

	"github.com/charmbracelet/lipgloss"
	myform "github.com/concerned-doggo/quizTUI/MyForm"
)


func main() {
    result := myform.MakeForm()
   

    var titleStyle = lipgloss.NewStyle().
        Bold(true).
        Foreground(lipgloss.Color("#7D56F4")).
        PaddingRight(1)
    var resultStyle = lipgloss.NewStyle().
        Bold(true).
        Foreground(lipgloss.Color("#FAFAFA"))

    fmt.Print(titleStyle.Render("Your Score is"))
    fmt.Println(resultStyle.Render(strconv.Itoa(result) + " / 5" ))
}

