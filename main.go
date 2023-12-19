package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"regexp"
	"strings"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Printf("Failed to get cwd: %s\n", err)
		os.Exit(1)
	}

	matchfile := path.Join(cwd, ".plexmatch")
	content, err := os.ReadFile(matchfile)

	if err == nil {
		fmt.Printf("The .plexmatch file already exists:\n%s", string(content))
		os.Exit(1)
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Please input IMDb id (e.g. tt0903747): ")
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)
	matched, _ := regexp.MatchString(`tt\d+`, line)
	if !matched {
		fmt.Printf("Invalid IMDb id: %s\n", line)
		os.Exit(1)
	}
	imdbid := line

	fmt.Printf("Please input season number (e.g. 02, defaults to 01): ")
	line, _ = reader.ReadString('\n')
	line = strings.TrimSpace(line)
	if line == "" {
		line = "01"
	}
	matched, _ = regexp.MatchString(`\d+`, line)
	if !matched {
		fmt.Printf("Invalid season number: %s\n", line)
		os.Exit(1)
	}
	season := line

	file, err := os.Create(matchfile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	data := "imdbid: " + imdbid + "\n" + "season: " + season + "\n"
	_, err = file.WriteString(data)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Successfully wrote .plexmatch with the following content:\n%s", data)
}
