package myform

import (
    "fmt"

    "github.com/charmbracelet/huh"
)

func getOptions(data []map[string]interface{}) ([]map[string]interface{}, []map[string]interface{}) {

    var answerList []map[string]interface{}
    var correctAnswerList []map[string]interface{}

    for i := 0; i < len(data); i++ {
        answerList = append(answerList, data[i]["answers"].(map[string]interface{}))
        correctAnswerList = append(correctAnswerList, data[i]["correct_answers"].(map[string]interface{}))
    }

    return answerList, correctAnswerList
}


func MakeForm(data []map[string]interface{}) (){
    fmt.Println("this is fine")

    // answers := data[0]["answers"].(map[string]interface{})
    answerList, correctAnswerList := getOptions(data)
    

    var userResponse [5]string
    form := huh.NewForm(
        huh.NewGroup(
            huh.NewSelect[string]().
                Title(data[0]["question"].(string)).
                Options(
                huh.NewOption(answerList[0]["answer_a"].(string), correctAnswerList[0]["answer_a_correct"].(string)),
                huh.NewOption(answerList[0]["answer_b"].(string), correctAnswerList[0]["answer_b_correct"].(string)),
                huh.NewOption(answerList[0]["answer_c"].(string), correctAnswerList[0]["answer_c_correct"].(string)),
                huh.NewOption(answerList[0]["answer_d"].(string), correctAnswerList[0]["answer_d_correct"].(string)),
                ).
                Value(&userResponse[0]),

            huh.NewSelect[string]().
                Title(data[1]["question"].(string)).
                Options(

                huh.NewOption(answerList[1]["answer_a"].(string), correctAnswerList[1]["answer_a_correct"].(string)),
                huh.NewOption(answerList[1]["answer_b"].(string), correctAnswerList[1]["answer_b_correct"].(string)),
                huh.NewOption(answerList[1]["answer_c"].(string), correctAnswerList[1]["answer_c_correct"].(string)),
                huh.NewOption(answerList[1]["answer_d"].(string), correctAnswerList[1]["answer_d_correct"].(string)),
                ).
                Value(&userResponse[1]),

            huh.NewSelect[string]().
                Title(data[2]["question"].(string)).
                Options(


                huh.NewOption(answerList[2]["answer_a"].(string), correctAnswerList[2]["answer_a_correct"].(string)),
                huh.NewOption(answerList[2]["answer_b"].(string), correctAnswerList[2]["answer_b_correct"].(string)),
                huh.NewOption(answerList[2]["answer_c"].(string), correctAnswerList[2]["answer_c_correct"].(string)),
                huh.NewOption(answerList[2]["answer_d"].(string), correctAnswerList[2]["answer_d_correct"].(string)),
                ).
                Value(&userResponse[2]),

            huh.NewSelect[string]().
                Title(data[3]["question"].(string)).
                Options(
                huh.NewOption(answerList[3]["answer_a"].(string), correctAnswerList[3]["answer_a_correct"].(string)),
                huh.NewOption(answerList[3]["answer_b"].(string), correctAnswerList[3]["answer_b_correct"].(string)),
                huh.NewOption(answerList[3]["answer_c"].(string), correctAnswerList[3]["answer_c_correct"].(string)),
                huh.NewOption(answerList[3]["answer_d"].(string), correctAnswerList[3]["answer_d_correct"].(string)),
                ).
                Value(&userResponse[3]),

            huh.NewSelect[string]().
                Title(data[4]["question"].(string)).
                Options(
                huh.NewOption(answerList[4]["answer_a"].(string), correctAnswerList[4]["answer_a_correct"].(string)),
                huh.NewOption(answerList[4]["answer_b"].(string), correctAnswerList[4]["answer_b_correct"].(string)),
                huh.NewOption(answerList[4]["answer_c"].(string), correctAnswerList[4]["answer_c_correct"].(string)),
                huh.NewOption(answerList[4]["answer_d"].(string), correctAnswerList[4]["answer_d_correct"].(string)),
                ).
                Value(&userResponse[4]),

            ),
        )
    form.Run()

    var totalScore int
    for i := 0; i < len(userResponse); i++ {
        if userResponse[i] == "true" {
            totalScore++;
        }
    }
    fmt.Println("Your total score is: ", totalScore)
}
