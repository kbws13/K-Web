package command

import (
	"KWeb/framework/cobra"
	"KWeb/framework/util"
	"bytes"
	"context"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/google/go-github/v39/github"
	"github.com/spf13/cast"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

// new相关的名称
func initNewCommand() *cobra.Command {
	return newCommand
}

// 创建一个新的应用
var newCommand = &cobra.Command{
	Use:     "new",
	Aliases: []string{"create", "init"},
	Short:   "创建一个新的应用",
	RunE: func(c *cobra.Command, args []string) error {
		currentPath := util.GetExecDirectory()

		var name string
		var folder string
		var mod string
		var version string
		var release *github.RepositoryRelease
		{
			prompt := &survey.Input{
				Message: "请输入目录名称: ",
			}
			err := survey.AskOne(prompt, &name)
			if err != nil {
				return nil
			}

			folder = filepath.Join(currentPath, name)
			if util.Exists(folder) {
				isForce := false
				prompt2 := &survey.Confirm{
					Message: "目录" + folder + "已经存在，是否删除重新创建？(确认后立即执行删除操作！)",
					Default: false,
				}
				err := survey.AskOne(prompt2, &isForce)
				if err != nil {
					return nil
				}

				if isForce {
					if err := os.RemoveAll(folder); err != nil {
						return err
					}
				} else {
					fmt.Println("目录已存在，创建应用失败")
					return nil
				}
			}
		}
		{
			prompt := &survey.Input{
				Message: "请输入模块名称(go.mod中的module, 默认为文件夹名称)：",
			}
			err := survey.AskOne(prompt, &mod)
			if err != nil {
				return err
			}
			if mod == "" {
				mod = name
			}
		}
		{
			// 获取KWen的版本
			client := github.NewClient(nil)
			prompt := &survey.Input{
				Message: "请输入版本名称(参考 https://github.com/kbws13/K-Web/releases，默认为最新版本)：",
			}
			err := survey.AskOne(prompt, &version)
			if err != nil {
				return err
			}
			if version != "" {
				// 确认版本是否正确
				release, _, err = client.Repositories.GetReleaseByTag(context.Background(), "kbws13", "K-Web", version)
				if err != nil || release == nil {
					fmt.Println("版本不存在，创建应用失败，请参考 https://github.com/kbws13/K-Web/releases")
					return nil
				}
			}
			if version == "" {
				release, _, err = client.Repositories.GetLatestRelease(context.Background(), "kbws13", "K-Web")
				version = release.GetTagName()
			}
		}
		fmt.Println("====================================================")
		fmt.Println("开始进行创建应用操作")
		fmt.Println("创建目录：", folder)
		fmt.Println("应用名称：", mod)
		fmt.Println("K-Web框架版本：", release.GetTagName())

		templateFolder := filepath.Join(currentPath, "template-kWeb-"+version+"-"+cast.ToString(time.Now().Unix()))
		os.Mkdir(templateFolder, os.ModePerm)
		fmt.Println("创建临时目录", templateFolder)

		// 拷贝template项目
		url := release.GetZipballURL()
		err := util.DownloadFile(filepath.Join(templateFolder, "template.zip"), url)
		if err != nil {
			return err
		}
		fmt.Println("下载zip包到template.zip")

		_, err = util.Unzip(filepath.Join(templateFolder, "template.zip"), templateFolder)
		if err != nil {
			return err
		}

		// 获取folder下的相关解压目录
		fInfos, err := os.ReadDir(templateFolder)
		if err != nil {
			return err
		}
		for _, fInfo := range fInfos {
			// 找到解压后的文件夹
			if fInfo.IsDir() && strings.Contains(fInfo.Name(), "goKWeb-KWeb-") {
				if err := os.Rename(filepath.Join(templateFolder, fInfo.Name()), folder); err != nil {
					return err
				}
			}
		}
		fmt.Println("解压zip包")

		if err := os.RemoveAll(templateFolder); err != nil {
			return err
		}
		fmt.Println("删除临时文件夹", templateFolder)

		os.RemoveAll(path.Join(folder, ".git"))
		fmt.Println("删除.git目录")

		// 删除framework 目录
		os.RemoveAll(path.Join(folder, "framework"))
		fmt.Println("删除framework目录")

		filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
				return nil
			}

			c, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			if path == filepath.Join(folder, "go.mod") {
				fmt.Println("更新文件:" + path)
				c = bytes.ReplaceAll(c, []byte("module KWeb"), []byte("module "+mod))
				c = bytes.ReplaceAll(c, []byte("require ("), []byte("require (\n\tKWeb "+version))
				err = os.WriteFile(path, c, 0644)
				if err != nil {
					return err
				}
				return nil
			}

			isContain := bytes.Contains(c, []byte("github.com/kbws13/KWeb/app"))
			if isContain {
				fmt.Println("更新文件:" + path)
				c = bytes.ReplaceAll(c, []byte("github.com/kbws13/KWeb/app"), []byte(mod+"/app"))
				err = os.WriteFile(path, c, 0644)
				if err != nil {
					return err
				}
			}

			return nil
		})
		fmt.Println("创建应用结束")
		fmt.Println("目录：", folder)
		fmt.Println("====================================================")
		return nil
	},
}
