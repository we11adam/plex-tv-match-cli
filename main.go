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
		fmt.Printf(".plexmatch file already exists:\n%s ", string(content))
		os.Exit(1)
	}

	fmt.Printf("Please input IMDb id (e.g. tt0903747): ")
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	matched, err := regexp.MatchString(`tt\d+`, line)
	if !matched {
		fmt.Printf("Invalid IMDb id: %s\n", line)
		os.Exit(1)
	}

	imdbid := strings.TrimSpace(line)

	fmt.Printf("please input season number (e.g. 02): ")
	line, err = reader.ReadString('\n')

	matched, err = regexp.MatchString(`\d+`, line)

	if !matched {
		fmt.Printf("Invalid season number: %s\n", line)
		os.Exit(1)
	}

	season := strings.TrimSpace(line)

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
	} else {
		fmt.Printf("Successfully wrote to file the following content to %s:\n%s", matchfile, data)
	}
}
