package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

type repoRequest struct {
	RepoName    string `json:"name"`
	Description string `json:"description"`
}

type repoResponse struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type githubError struct {
	Message string `json:"message"`
}

func main() {
	repo := repoRequest{
		RepoName:        "test102",
		Description: "we are testing",
	}
	err := CreateRepo(repo)
	if err != nil {
		log.Fatal(err)
	}
}

const url = "https://api.github.com/user/repos"

func CreateRepo(repo repoRequest) error {

	headers := http.Header{}
	headers.Set("Authorization", "token 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")

	repoJson, _ := json.Marshal(repo)

	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(repoJson))

	if err != nil {
		return err
	}
	client := http.Client{}
	request.Header = headers
	resp, err := client.Do(request)

	if err != nil {
		return err
	}
	githubResponse, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return err
	}

	if resp.StatusCode > 299 {
		var githubErr githubError
		json.Unmarshal(githubResponse, &githubErr)
		fmt.Println(githubErr)
		return errors.New("failed")
	}

	var result repoResponse

	err = json.Unmarshal(githubResponse, &result)
	if err != nil {
		return err
	}

	fmt.Println(result)
	return nil
}
