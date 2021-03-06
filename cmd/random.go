package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

// randomCmd represents the random command
var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "Get a random dad joke",
	Long:  `Command for getting an random dad joke`,
	Run: func(cmd *cobra.Command, args []string) {
		getRandomJoke()
	},
}

func init() {
	rootCmd.AddCommand(randomCmd)
}

var baseURL string = "https://icanhazdadjoke.com/"

type Joke struct {
	ID     string `json:"id"`
	Joke   string `json:"joke"`
	Status string `json:"status"`
}

func getRandomJoke() {
	jokeData := fetchJokeData()
	joke := Joke{}
	// Unmarshal the response
	json.Unmarshal(jokeData, &joke)

	fmt.Println(string(joke.Joke))
}

func fetchJokeData() []byte {
	request, error := http.NewRequest(http.MethodGet, baseURL, nil)

	if error != nil {
		log.Printf("An error has ocurred %v", error)
	}

	request.Header.Add("Accept", "application/json")

	response, error := http.DefaultClient.Do(request)

	if error != nil {
		log.Printf("An error has ocurred %v", error)
	}

	resBytes, error := ioutil.ReadAll(response.Body)

	if error != nil {
		log.Printf("An error has ocurred %v", error)
	}

	return resBytes
}
