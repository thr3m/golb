package cli

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/thr3m/nojs/helpers"
)

/*
@brief Handle the user input.
Offers 4 options: init, post, delete, deploy
init : Initization of the blog directory
post : Start the creation process of a blog article
delete : Start the deletion process of a blog article
deploy: TBD
*/
func HandleUserInput(argList []string) {
	command := argList[0]

	switch command {
	case "init":
		initBlog()
	case "create":
		createPost()
	case "delete":
		fmt.Println("Deleting a post")
	case "deploy":
		fmt.Println("Deploying a blog")
	}
}

/*
@brief Initialization of the blog directory.
Ask for desired blog path. If none provided, create a directery under the user's
home directory named my_blog
*/
func initBlog() string {
	user_os := runtime.GOOS
	postDirectoryPath, err6 := os.UserHomeDir()
	scanner := bufio.NewScanner(os.Stdin)
	blogName := "my_blog"

	if err6 != nil {
		fmt.Print(err6.Error())
	}

	fmt.Println("Let's init your new blog !")

	fmt.Print("How would you want to name your blog ? (default : " + blogName + " ) : ")
	if scanner.Scan() {
		if scanner.Text() != "" {
			blogName = scanner.Text()
		}
	} else {
		log.Fatal("Please provide a post name")
	}

	fmt.Print("Please provide a path where golb will look for your blog (default : " + postDirectoryPath + " ) : ")
	if scanner.Scan() {
		if scanner.Text() != "" {
			postDirectoryPath = scanner.Text()
		}
	}

	blogPath := filepath.Join(postDirectoryPath, blogName)

	// Adding a backslash at the end of the blog path
	if user_os != "windows" {
		if blogPath[len(blogPath)-1] != '/' {
			blogPath = blogPath + "/"
		}
	} else {
		if blogPath[len(blogPath)-1] != '\\' {
			blogPath = blogPath + "\\"
		}
	}

	// Initialize the blog's config
	err := helpers.InitBlogConfig(blogPath)

	if err != nil {
		log.Fatal(err)
	}

	// Initialize the blog's config
	helpers.InitAppConfig(blogPath)

	return postDirectoryPath
}

/*
@brief Create a blog post.
Ask user for a post name and create a file under /posts directory.
If no post title is provided, use a timestamp as a temporary title.
*/
func createPost() {
	var postTitle string
	var postDescription string

	fmt.Println("Create a post !")

	fmt.Print("Enter a post title : ")
	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		if scanner.Text() != "" {
			postTitle = scanner.Text()
		}
	} else {
		log.Fatal("Please provide a post name")
	}

	fmt.Print("Enter a short description of the blog post (default blank) : ")
	if scanner.Scan() {
		postDescription = scanner.Text()
	}

	err := helpers.CreatePost(postTitle, postDescription)

	if err != nil {
		fmt.Print("Couldn't create the post" + err.Error())
	}
}
