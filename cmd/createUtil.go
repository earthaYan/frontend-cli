package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"

	"github.com/AlecAivazis/survey/v2"
)

func getRepoInfo(repoList []templateType, repoNameList []string) (templateType) {
	prompt := &survey.Select{
		Message: "Choose a template:",
		Options: repoNameList,
	}
	var currentRepo string
	survey.AskOne(prompt, &currentRepo)
	var currentRepoItem templateType
	for _, repo := range repoList {
		if repo.name == currentRepo {
			currentRepoItem=repo
			break
		}
	}
	return currentRepoItem
}

func getTagInfo(repoItem templateType) (string, error) {
	var tagNames []string
	var url = "https://api.github.com/repos/" +repoItem.name + "/tags"
	if resp, err := http.Get(url); err != nil {
		fmt.Println(err)
		os.Exit(1)
		return "", err
	} else {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			os.Exit(1)
			return "", err
		}
		var tagList []struct {
			Name string `json:"name"`
		}

		json.Unmarshal(body, &tagList)
		for _, tag := range tagList {
			tagNames = append(tagNames, tag.Name)
		}
		if len(tagNames)<1{
			return "",nil
		}
		prompt := &survey.Select{
			Message: "Choose a version:",
			Options: tagNames,
		}
		var currentTag string
		survey.AskOne(prompt, &currentTag)
		return currentTag, nil
	}
}
func downloadTemplate(repo  templateType, tag string, targetFilePath string) error {
	var templateUrl=repo.downloadUrl
	var cmd *exec.Cmd
	if tag!=""{
		cmd = exec.Command("git", "clone", "-b", tag, templateUrl, targetFilePath)
	}else{
		cmd=exec.Command("git", "clone", templateUrl, targetFilePath)
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func getTemplatesDataList(category string) ([]templateType, []string) {
	var vueTemplates []templateType
	var vueTemplateNames []string
	var reactTemplates []templateType
	var reactTemplateNames []string
	for _, template := range TemplatesSource {
		if template.category == "vue" {
			vueTemplates = append(vueTemplates, template)
			vueTemplateNames = append(vueTemplateNames, template.name)
		} else {
			reactTemplates = append(reactTemplates, template)
			reactTemplateNames = append(reactTemplateNames, template.name)
		}
	}
	if category == "vue" {
		return vueTemplates, vueTemplateNames
	} else {
		return reactTemplates, reactTemplateNames
	}
}
