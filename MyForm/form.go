package myform

import (
	"math/rand"

	"github.com/charmbracelet/huh"
	apicall "github.com/concerned-doggo/quizTUI/APICall"
)


func getOptions(data apicall.QuestionList) ([][]string, [][]bool){
    var options  [][]string
    var correctAns  [][]bool

    for i := 0; i < len(data.Results); i++ {
        answers := append([]string{data.Results[i].CorrectAnswer}, data.Results[i].IncorrectAnswers...)
        x := shuffleOptions(answers)
        options = append(options, x)

        var tempCorrectAns [4]bool
        for j:= 0; j < len(x); j++{
            tempCorrectAns[j] = x[j] == data.Results[i].CorrectAnswer
        }

        correctAns = append(correctAns, tempCorrectAns[:])
    }
    return options, correctAns
}

func shuffleOptions(options []string) ([]string){

    rand.Shuffle(len(options), func(i, j int){
        options[i], options[j] = options[j], options[i]
    })

    return options
}


func getQuestions(data apicall.QuestionList) ([]string){
    var questions []string
    for i := 0; i < len(data.Results); i++ {
        questions = append(questions, data.Results[i].Question)
    }
    return questions
}

func MakeForm() (int){
    data := apicall.GetData()
    options, correctAns := getOptions(data)
    questions := getQuestions(data)
    var userResponse [5]bool

    for i := 0; i < len(questions); i++ {
        huh.NewSelect[bool]().
            Title(questions[i]).
            Options(
            huh.NewOption(options[i][0], correctAns[i][0]),
            huh.NewOption(options[i][1], correctAns[i][1]),
            huh.NewOption(options[i][2], correctAns[i][2]),
            huh.NewOption(options[i][3], correctAns[i][3]),
            ).
            Value(&userResponse[i]).
            Run()
    }


    var totalScore int
    for i := 0; i < len(userResponse); i++ {
        if userResponse[i] {
            totalScore++;
        }
    }
    return totalScore
}
