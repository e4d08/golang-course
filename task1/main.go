package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"time"
)

const (
	Reset   = "\033[0m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	Gray    = "\033[37m"
	White   = "\033[97m"
)

var re = regexp.MustCompile(`^http(?:s)?:\/\/github\.com\/([^\/]*)\/([^\/]*)`)
var client = &http.Client{
	Timeout: 5 * time.Second,
}

type Repository struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Stars       int       `json:"stargazers_count"`
	Forks       int       `json:"forks_count"`
	CreatedAt   time.Time `json:"created_at"`
	License     License   `json:"license"`
}

type License struct {
	Name string `json:"name"`
}

func colored(s string, color string) string {
	return color + s + Reset
}

func printError(message string) {
	fmt.Println(colored(message, Red))
}

func printUsage() {
	fmt.Println("Usage:", os.Args[0], "<github link>")
}

func crash() {
	os.Exit(1)
}

func parseRepoFromLink(link string) (string, string, error) {
	match := re.FindStringSubmatch(link)

	if len(match) < 3 {
		return "", "", errors.New("failed to find GitHub repo in the link")
	}

	return match[1], match[2], nil
}

func fetchGithub(path string, target any) error {
	link := fmt.Sprintf("https://api.github.com/%s", path)

	res, err := client.Get(link)
	if err != nil {
		return errors.New("failed to fetch resource: " + err.Error())
	}
	defer res.Body.Close()

	switch res.StatusCode {
	case http.StatusOK:
	case http.StatusNotFound:
		return errors.New("repository not found")
	default:
		return errors.New("bad response from server")
	}

	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return err
	}

	return nil
}

func fetchGithubRepo(owner string, name string) (Repository, error) {
	var result Repository

	err := fetchGithub(fmt.Sprintf("repos/%s/%s", owner, name), &result)
	if err != nil {
		return Repository{}, err
	}

	if result.License.Name == "" {
		result.License.Name = "None"
	}

	return result, nil
}

func main() {
	if len(os.Args) < 2 {
		printError("link not specified")
		printUsage()
		crash()
	}

	fmt.Print(colored("Parsing link...\r", Blue))
	var link string = os.Args[1]
	repoOwner, repoName, err := parseRepoFromLink(link)
	if err != nil {
		printError("bad link: " + err.Error())
		crash()
	}
	fmt.Println(colored(fmt.Sprintf("Parsed link! Owner: %s | Name: %s", repoOwner, repoName), Green))

	fmt.Print(colored("Fetching GitHub...\r", Blue))
	repo, err := fetchGithubRepo(repoOwner, repoName)
	if err != nil {
		printError("failed to fetch GitHub repository: " + err.Error())
		crash()
	}
	fmt.Println(colored("Fetched successfully!", Green))

	var layout string
	if time.Now().Year() != repo.CreatedAt.Year() {
		layout = "Jan 2, 2006"
	} else {
		layout = "Jan 2"
	}

	fmt.Println(colored("Name: ", Blue) + repo.Name)
	fmt.Println(colored("Description: ", Blue) + repo.Description)
	fmt.Println(colored("Stars: ", Blue) + strconv.Itoa(repo.Stars))
	fmt.Println(colored("Forks: ", Blue) + strconv.Itoa(repo.Forks))
	fmt.Println(colored("Created at: ", Blue) + repo.CreatedAt.Format(layout))
	fmt.Println(colored("License: ", Blue) + repo.License.Name)
}
