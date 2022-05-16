package helpers

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"time"

	"gopkg.in/yaml.v3"
)

type Post struct {
	Title       string    `yaml:"title"`
	Description string    `yaml:"description"`
	Date        time.Time `yaml:"date"`
}

/*
Will have an array of post. Needed for not having to go through
all the file within the post directory to fetch de post name, description
and date
*/
type Config struct {
	BlogPath string `yaml:"blogPath"`
	BlogName string `yaml:"blogName"`
	Posts    []Post `yaml:"posts"`
}

type AppConfig struct {
	BlogPath string `yaml:"blogPath"`
}

/*
@brief Get the blogPath from the app configuration file.
*/
func GetBlogPath() string {

	var blogConfig map[string]string

	configFilePath, err := filepath.Abs("./config.yaml")

	if err != nil {
		log.Fatal("Couldn't create the file path to the app config file" + err.Error())
	}

	config, err2 := ioutil.ReadFile(configFilePath)

	if err2 != nil {
		fmt.Println("Couldn't not read the app config file" + err2.Error())
	}

	err3 := yaml.Unmarshal(config, &blogConfig)

	if err3 != nil {
		log.Fatal("Coudln't unmarshal the app config file " + err3.Error())
	}

	return blogConfig["blogPath"]
}

/*
@brief Initialization of the configuration file
Create a directory at the destination provided by the user.
Create a yaml file that will contain the configuration of the blog.
Add the blogname to the app config file.
*/
func InitConfig(blogpath string) error {
	var err error
	configFilePath := blogpath + "config.yaml"
	postsDirectoryPath := blogpath + "posts"

	// blog folder creation
	err = os.Mkdir(blogpath, 0755)

	if err != nil {
		log.Fatal("Error creating the config directory " + err.Error())
	}

	// posts folder creation
	err = os.Mkdir(postsDirectoryPath, 0755)

	if err != nil {
		log.Fatal("Error creating the posts directory with the blog directory" + err.Error())
	}

	// Create the configuration file within the blog directory
	_, err = os.Create(configFilePath)

	if err != nil {
		log.Fatal("Error creating the config file " + err.Error())
	}

	config := Config{}

	var initalPosts []Post

	config.BlogPath = blogpath

	config.Posts = initalPosts

	// Update the blog config with the new blog path
	UpdateBlogConfig(config)

	appConfig := map[string]string{"blogPath": blogpath}

	ymlAppConfig, err3 := yaml.Marshal(appConfig)

	if err3 != nil {
		return err3
	}

	err = ioutil.WriteFile("config.yaml", ymlAppConfig, 0)

	if err != nil {
		fmt.Printf("Error writting to the app config file %s", err.Error())
		return err
	}

	return nil
}

/*
@brief Creation of a blog post. In order to create a blog post, we need to create
a markdown file inside the post folder and add the post title to the posts array
inside the blog configuration file.
*/
func CreatePost(postTitle string, postDescription string) error {
	var newPost Post
	var tempConfig Config

	postDate := time.Now()

	newPost.Date = postDate
	newPost.Title = postTitle
	newPost.Description = postDescription

	blogPath := GetBlogPath()

	blogConfigPath := path.Join(blogPath, "config.yaml")

	blogConfig, err := ioutil.ReadFile(blogConfigPath)

	if err != nil {
		log.Fatal("Couldn't read blog config : " + err.Error())
	}

	err2 := yaml.Unmarshal(blogConfig, &tempConfig)

	if err2 != nil {
		log.Fatal("Coudln't unmarshal the app config file " + err2.Error())
	}

	tempConfig.Posts = append(tempConfig.Posts, newPost)

	UpdateBlogConfig(tempConfig)

	return nil
}

/*
@brief Update the blog configuration by overwritting the
existing blog configuration
*/
func UpdateBlogConfig(blogConfig Config) {

	blogConfigPath := path.Join(blogConfig.BlogPath, "config.yaml")

	fmt.Println(blogConfigPath)

	ymlUserConfig, err := yaml.Marshal(blogConfig)

	if err != nil {
		log.Fatal("Error while marshaling the blog config into yaml : " + err.Error())
	}

	err2 := ioutil.WriteFile(blogConfigPath, ymlUserConfig, 0)

	if err2 != nil {
		log.Fatal("Error while writting to the blog's config : " + err2.Error())
	}
}
