package main

import (
	"fmt"
	"log"
	"os"
	"Codivy/internal/github/requests"
	"Codivy/internal/modules"
	"bufio"
	"strings"
	"encoding/json"
)

func main() {
	var err error

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("What will be the name of your repository?: ")
	repoName, _ := reader.ReadString('\n')
	repoName = strings.TrimSpace(repoName)

	fmt.Print("Do you want the repository to be private? (y/n): ")
	priv, _ := reader.ReadString('\n')
	priv = strings.ToLower(strings.TrimSpace(priv))

	isPriv := false
	if priv == "y" || priv == "yes" {
		isPriv = true
	}

	fmt.Print("What will be the description of the repository?: ")
	repoDesc, _ := reader.ReadString('\n')
	repoDesc = strings.TrimSpace(repoDesc)

	fmt.Print("What type of project license? ")
	repoLicense, _ := reader.ReadString('\n')
	repoLicense = strings.TrimSpace(repoLicense)

	var repoToken string
	for {
    	fmt.Print("Enter the token to create the repository (required): ")
    	repoToken, _ = reader.ReadString('\n')
    	repoToken = strings.TrimSpace(repoToken)
    	if repoToken != "" {
        	break
    	}
    	fmt.Println("Token cannot be empty. Please enter a valid GitHub token.")
	}

	fmt.Print("Choose a language for the template (python/go/node): ")
	lang, _ := reader.ReadString('\n')
	lang = strings.TrimSpace(strings.ToLower(lang))

	os.Mkdir(repoName, 0755)
	templatePath := fmt.Sprintf("template/%s", lang)
	err = modules.CopyDir(templatePath, repoName)

	if err != nil {
		log.Fatalf("Failed to copy template: %v", err)
	}

	fmt.Println("Template files copied to local folder!")

	client := requests.NewClient(repoToken)

	payload := map[string]any{
		"name":             repoName,
		"private":          isPriv,
		"description":      repoDesc,
		"license_template": repoLicense,
	}

	resp, err := client.Post("/user/repos", payload)
	if err != nil {
		log.Fatalf("Error creating repository: %v", err)
	}

	type RepoResponse struct {
		HTMLURL string `json:"html_url"`
		CloneURL string `json:"clone_url"`
	}

	var repoData RepoResponse
	if err := json.NewDecoder(resp.Body).Decode(&repoData); err == nil {
		fmt.Printf("\nRepository created successfully!\nURL: %s\n", repoData.HTMLURL)
	} else {
		fmt.Printf("[-] Created, but couldn't parse the response.\n")
	}

	err = modules.SetupRepo(repoData.CloneURL, repoName, lang)
	if err != nil {
		log.Fatalf("[-] Error when configuring the repository: %v", err)
	}
}
