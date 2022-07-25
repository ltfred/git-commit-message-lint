package cmd

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
)

var LintCmd = &cobra.Command{
	Use:   "gcml",
	Short: "git commit message lint",
	Run:   lint,
}

type TypeAndDes struct {
	Type  string
	Des   string
	Color func(format string, a ...interface{})
}

var typeAndDes = []TypeAndDes{
	{
		Type:  "feat",
		Des:   "新增/修改功能",
		Color: color.HiGreen,
	},
	{
		Type:  "fix",
		Des:   "修补bug",
		Color: color.HiRed,
	},
	{
		Type:  "docs",
		Des:   "文档（documentation）",
		Color: color.HiWhite,
	},
	{
		Type:  "style",
		Des:   "格式（不影响代码运行的变动）",
		Color: color.Yellow,
	},
	{
		Type:  "refactor",
		Des:   "重构（即不是新增功能，也不是修改bug的代码变动）",
		Color: color.Blue,
	},
	{
		Type:  "perf",
		Des:   "优化相关，比如提升性能、体验",
		Color: color.Magenta,
	},
	{
		Type:  "test",
		Des:   "增加测试",
		Color: color.HiBlue,
	},
	{
		Type:  "chore",
		Des:   "构建过程或辅助工具的变动",
		Color: color.Cyan,
	},
	{
		Type:  "revert",
		Des:   "回滚到上一个版本",
		Color: color.White,
	},
}

func lint(cmd *cobra.Command, args []string) {
	for i := range typeAndDes {
		t := typeAndDes[i]
		t.Color("%-10s %s\n", fmt.Sprintf("%d."+t.Type, i), t.Des)
	}
	var num int
	fmt.Print("请选择一个type: ")
	_, err := fmt.Scanln(&num)
	if err != nil {
		os.Exit(1)
	}
	fmt.Print("请输入message: ")
	reader := bufio.NewReader(os.Stdin)
	msg, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	m := typeAndDes[num].Type + ": " + msg
	command := exec.Command("git", "add", ".")
	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	err = command.Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	command = exec.Command("git", "commit", "-m", m)
	err = command.Run()
	if err != nil {
		fmt.Println(err)
	}
}
