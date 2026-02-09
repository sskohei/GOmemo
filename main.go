package main

import (
	"fmt"
	"os"
	"sample_go/command"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("使い方:")
		fmt.Println("  memo add \"内容\"")
		fmt.Println("  memo list")
		fmt.Println("  memo show ID")
		fmt.Println("  memo search keyword")
		fmt.Println("  memo delete ID")
		return
	}

	switch os.Args[1] {

	case "add":
		if len(os.Args) < 3 {
			fmt.Println("メモ内容を指定してください")
			return
		}
		command.Add(os.Args[2])

	case "list":
		id := 10
		if len(os.Args) > 2 {
			id, _ = strconv.Atoi(os.Args[2])
		}
		command.List(id)

	case "show":
		id, _ := strconv.Atoi(os.Args[2])
		command.Show(id)

	case "search":
		if len(os.Args) < 3 {
			fmt.Println("検索する言葉を指定してください")
			return
		}
		command.Search(os.Args[2])

	case "delete":
		id, _ := strconv.Atoi(os.Args[2])
		command.Delete(id)

	default:
		fmt.Println("不明なコマンドです")
	}
}
