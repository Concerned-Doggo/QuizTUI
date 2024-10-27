package apicall

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)


func Question(category string)([]map[string]interface{}){
    godotenv.Load()
    API_KEY := os.Getenv("API_KEY")
    apiUrl := "https://quizapi.io/api/v1/questions?apiKey=" + API_KEY + "&limit=5&category=" + category

    res, err := http.Get(apiUrl)
    if err != nil {
        panic(err)
    }
    defer res.Body.Close()

    var data []map[string]interface{}
    json.NewDecoder(res.Body).Decode(&data)

    // for pretty printing json data in the terminal
    // jsonBytes, _ := json.MarshalIndent(data, "", "   ")
    // fmt.Println(string(jsonBytes))
    fmt.Println()
    return data
}
