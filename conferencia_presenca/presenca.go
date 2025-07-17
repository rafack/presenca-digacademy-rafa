package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

type Commit struct {
	Commit struct {
		Message string `json:"message"`
		Date    string `json:"committer"`
	} `json:"commit"`
}

func main() {
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		panic("GITHUB_TOKEN not set")
	}
	aula := os.Getenv("AULA")
	if token == "" {
		panic("AULA not set")
	}

	file, err := os.Open("repos.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	fmt.Println("Alunas presentes: ")
	for scanner.Scan() {
		line := scanner.Text()
		name, repoURL, _ := strings.Cut(line, " | ")
		parts := strings.Split(repoURL, "github.com/")
		if len(parts) < 2 {
			fmt.Println(name, " - Erro ao dividir url do repositório em partes")
			continue
		}
		repoPath := parts[1]
		apiURL := fmt.Sprintf("https://api.github.com/repos/%s/commits", repoPath)

		req, _ := http.NewRequest("GET", apiURL, nil)
		req.Header.Set("Authorization", "Bearer "+token)
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println(name, " - Erro na requisição")
			continue
		}
		defer resp.Body.Close()

		var commits []Commit
		json.NewDecoder(resp.Body).Decode(&commits)

		found := false
		for _, c := range commits {
			if strings.Contains(c.Commit.Message, fmt.Sprint("aula ", aula, " -")) {
				// fmt.Printf("[✅] %s commitou: %s\n", name, c.Commit.Message)
				fmt.Printf("%s\t\t%d\n", name, 1)
				found = true
				break
			}
		}
		if !found {
			fmt.Printf("%s\t\t%d\n", name, 0)
		}
	}
}
