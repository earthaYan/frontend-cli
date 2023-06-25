package cmd

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
	flag "github.com/spf13/pflag"
)

var forceVar bool
var frameVar string
var questions = []*survey.Question{
	{
		Name: "overWrite",
		Prompt: &survey.Select{
			Message: "the project dir existed,would you like to over write it?",
			Options: []string{"yes", "no"},
			Default: "yes",
		},
	},
}

func init() {
	// flag 的定义需要在 Cobra 配置前完成,否则 flag 就无法被 Cobra 命令识别
	flag.BoolVar(&forceVar, "force", false, "overwrite dir directly")
	flag.StringVar(&frameVar, "frame", "vue", "which framework u wanna use,react or vue")
	flag.Parse()
}
func createProject(cmd *cobra.Command, args []string) {
	createProjectWithCertainName(args[0])
}
func create(projectName string, targetFilePath string) {
	repoList, repoNameList := getTemplatesDataList(frameVar)
		// 获取仓库信息-模板
	currentRepoItem:=getRepoInfo(repoList, repoNameList)
	// 获取标签信息-版本信息
	currentRepoTag,err:=getTagInfo(currentRepoItem)
	// 下载模板
	if err != nil {
		log.Fatal("get info failed")
	} else {
		err := downloadTemplate(currentRepoItem, currentRepoTag, targetFilePath)
		// 展示模板使用提示
		if err == nil {
			fmt.Printf("Successfully created project %s\n", projectName)
			fmt.Printf("cd %s\r\n", projectName)
			fmt.Println("npm install")
			fmt.Println("npm run serve")
		}
	}

}
func createProjectWithCertainName(projectName string) {
	currentDir, _ := os.Getwd()
	targetFilePath := path.Join(currentDir, projectName)
	answers := struct {
		OverWrite string
		Template  string
	}{}
	// 如果文件名不存在,os.Stat()会返回错误
	if _, error := os.Stat(targetFilePath); error == nil {
		// 文件已存在
		// 如果有force
		if forceVar {
			// 强制执行
			// 删除同名目录
			if err := os.RemoveAll(targetFilePath); err != nil {
				log.Fatal("delete files failed....")
			}
		} else {
			if err := survey.Ask(questions, &answers); err != nil {
				fmt.Println(err.Error())
				return
			}
			fmt.Println(answers)
			//交互式
			if answers.OverWrite == "no" {
				// 不重写项目目录
				fmt.Println("cancel ")
				return
			} else {
				// 重写目录
				if err := os.RemoveAll(targetFilePath); err != nil {
					log.Fatal("overwrite failed....")
				}
			}
		}
	}
	// 创建项目
	create(projectName, targetFilePath)
}
