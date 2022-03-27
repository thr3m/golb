package helpers

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Post struct {
	Title       string
	Description string
	Date        string
}

/*
Will have an array of post. Needed for not having to go through
all the file within the post directory to fetch de post name, description
and date
*/
type Config struct {
	BlogPath string
	BlogName string
	Posts    []Post
}

/*
@brief Initialization of the configuration file
Will create a directory at the destination provided
by the user.
Will create a yaml file that will contain the configuration
of the blog
*/
func InitConfig(blogpath string) error {

	configFilePath := blogpath + "config.yml"

	err := os.Mkdir(blogpath, 0755)

	if err != nil {
		log.Fatal("Error creating the config directory " + err.Error())
	}

	_, err2 := os.Create(configFilePath)

	if err2 != nil {
		log.Fatal("Error creating the config file " + err.Error())
	}

	config := Config{}

	var initalPosts []Post

	config.BlogPath = blogpath

	config.Posts = initalPosts

	data, err3 := yaml.Marshal(config)

	if err3 != nil {
		return err3
	}

	err4 := ioutil.WriteFile(configFilePath, data, 0)

	if err4 != nil {
		fmt.Print("Error writting to config file %s", err4.Error())
		return err4
	}

	return nil
}
