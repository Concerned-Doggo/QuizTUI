package apicall

import (
	"encoding/json"
	// "fmt"
	"net/http"
)

type QuestionList struct {
    ResponseCode int `json:"response_code,omitempty"`
    Results      []Question `json:"results,omitempty"`
}

type Question struct {
    Type             string   `json:"type,omitempty"`
    Difficulty       string   `json:"difficulty,omitempty"`
    Category         string   `json:"category,omitempty"`
    Question         string   `json:"question,omitempty"`
    CorrectAnswer    string   `json:"correct_answer,omitempty"`
    IncorrectAnswers []string `json:"incorrect_answers,omitempty"`

}


func GetData()(QuestionList){
    apiUrl := "https://opentdb.com/api.php?amount=5&category=18&difficulty=easy&type=multiple"

    res, err := http.Get(apiUrl)
    if err != nil {
        panic(err)
    }
    defer res.Body.Close()

    var data QuestionList
    json.NewDecoder(res.Body).Decode(&data)

    // prettyData, _ := json.MarshalIndent(data, "", "   ")
    // fmt.Println(string(prettyData))

    return data
}
