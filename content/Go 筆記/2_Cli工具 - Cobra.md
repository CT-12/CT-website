<!-- 
date: 2024-01-29
update: 2024-01-29
Tag: 
    - Go
    - Cli
    - Package
-->

# Cobra

## Install

可以先用 `go install` 安裝 `cobra-cli` 的命令行工具

```shell
go install github.com/spf13/cobra-cli@latest
```

## Introduction

Cobra 有三個部分

- Commands : 用來表示要執行的動作
- Args : 執行動作所需的參數
- Flags : 這些動作額外的行為

```shell
# 例如

git clone URL -bare

# git : 根命令
# clone : 子命令，表示要執行的動作（Command）
# URL : 執行該動作所需要的參數 (Arg)
# -bare : 執行 clone 這個動作額外的行為 (Flag)
```

## Start 

假設我要做一個 Todo list 的命令行工具

```shell
mkdir Todo
cd Todo
go mod init Todo
cobra-cli init 
```

使用 cobra-cli 初始化專案後會在專案下看到以下檔案

```
Todo/
|-cmd/
|  |- root.go
|- main.go
```

在 main.go 就只是很單純的呼叫 root.go 的 Execute 函式。而在 root.go 會有根命令的一些內容。

```go
// In root.go

/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

const (
	VERSION string = "0.1.0"
)

var (
	show_version bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "Todo",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) { 
		if show_version {
			cmd.Println("Version: ", VERSION)
		} else {
			cmd.Help()
		}
	},
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.To-do-list.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolVarP(&show_version, "version", "v", false, "Show version")
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
```

在 init 函式中可以定義 Flag。而 Run 則是使用根命令時會執行的函式，其中在函式簽名裡有個 `args`，傳給命令的參數會在裡面。

### 新增子命令

如果想要有子命令可以用下面的指令創建：

```shell
cobra-cli add <New Command>
```

這個指令會在 `cmd/` 目錄下創建一個與子命令同名的新的檔案，同時在該檔案會看到： 

```go
// 假設子命令名稱叫 show 
// 在 cmd/show.go 會看到在 init 函式中子命令被加到根命令

func init() {
	rootCmd.AddCommand(showCmd)
}
```

接下來就可以在子命令的 Run 定義要執行的動作了

# References

1. [go命令行库-cobra](https://blog.csdn.net/zy010101/article/details/127397143)
2. [万字长文——Go 语言现代命令行框架 Cobra 详解](https://xie.infoq.cn/article/915006cf3760c99ad0028d895)
3. [在 Golang 中使用 Cobra 创建 CLI 应用](https://www.qikqiak.com/post/create-cli-app-with-cobra/)