package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path"
)

type AbstractSyntaxTree struct {
	root *AbstractSyntaxNode
}

type AbstractSyntaxNode struct {
	token  string
	childs []*AbstractSyntaxNode
}

const (
	TEXT      = "text"
	TITLEONE  = "titleOne"
	TITLETWO  = "titleTwo"
	TITLETREE = "titleTree"
	TITLEFOUR = "titleFour"
	TITLEFIVE = "titleFive"
	BOLD      = "bold"
	ITALIC    = "italic"
)

func getPostsPath() []string {
	var posts []string
	blogPath := GetBlogPath()
	postsPath := path.Join(blogPath, "posts")

	items, _ := os.ReadDir(postsPath)

	for _, item := range items {
		posts = append(posts, path.Join(postsPath, item.Name()))
	}

	return posts
}

/*
@brief Read a markdown file line per line and creates an
abstract syntax tree
*/
func ParseMarkdown() {
	posts := getPostsPath()
	for _, postPath := range posts {
		file, err := os.Open(postPath)

		if err != nil {
			log.Fatal("Couldn't read the post located at " + postPath + " err")
		}

		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			line := scanner.Text()
			firstCharacter := line[0]
			switch firstCharacter {
			case '#':
				fmt.Println("Possible title")
			}
		}
	}
}
