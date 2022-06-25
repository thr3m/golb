package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"path"
)

type AbstractSyntaxTree struct {
	root *AbstractSyntaxNode
}

type AbstractSyntaxNode struct {
	nodeType string
	content  string
	children []*AbstractSyntaxNode
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
			createNodeFromLine(line)
		}
	}
}

func createNodeFromLine(line string) AbstractSyntaxNode {
	var node AbstractSyntaxNode

	nodeType, contentIndex := getNodeType(line)
	lineContent := line[contentIndex:]
	node.nodeType = nodeType

	_ = findChildren(lineContent, nodeType)

	return node
}

// handleModifier return a node and the index of which that modifier ends.
func handleModifier(line string, modifier byte) (string, int) {
	endOfModifierIndex := 0
	var buf []byte
	modifiedText := bytes.NewBuffer(buf)

	// Check if the modifier is the last caracter of the line. Lenght of line will be 2
	// since the end of line is finished by an enter -> ASCII = 32
	if len(line) == 2 && line[1] == 32 {
		return "", 0
	}

	for i := 1; i < len(line); i++ {
		if line[i] == modifier {
			endOfModifierIndex = i
			break
		}

		modifiedText.WriteByte(line[i])
	}

	if endOfModifierIndex == 0 {
		return "", 0
	}

	fmt.Println(modifiedText)

	return modifiedText.String(), endOfModifierIndex
}

func findChildren(line string, nodeType string) []AbstractSyntaxNode {
	var children []AbstractSyntaxNode
	var naturalNodeContent []byte

	for i := 0; i < len(line); i++ {
		switch line[i] {
		case '*':
			modifiedText, lastModifierIndex := handleModifier(line[i:], '*')

			// If lastModifierIndex is 0, it means that the matching modifier is at
			// the end of the string or there's no matching modifier
			if lastModifierIndex != 0 {
				var naturalChild AbstractSyntaxNode
				var child AbstractSyntaxNode

				naturalChild.content = string(naturalNodeContent)
				naturalChild.nodeType = nodeType
				naturalChild.children = nil
				naturalNodeContent = []byte{}

				child.nodeType = ITALIC
				child.content = modifiedText
				child.children = nil

				children = append(children, naturalChild)
				children = append(children, child)
			}
		default:
			fmt.Println("Natual Node Content : " + string(naturalNodeContent))
			naturalNodeContent = append(naturalNodeContent, line[i])
		}
	}

	fmt.Println(children)

	return children
}

func getNodeType(line string) (string, int) {
	var nodeType string
	var indexContent int

	switch line[0] {
	case '#':
		titleType := 1
		for i := 0; i < len(line); i++ {
			if line[i] != ' ' {
				titleType++
			} else {
				break
			}
		}

		indexContent = titleType

		switch titleType {
		case 1:
			nodeType = TITLEONE
		case 2:
			nodeType = TITLETWO
		case 3:
			nodeType = TITLETREE
		case 4:
			nodeType = TITLEFOUR
		case 5:
			nodeType = TITLEFIVE
		}
	default:
		nodeType = TEXT
	}

	return nodeType, indexContent
}
