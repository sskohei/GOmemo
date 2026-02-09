package command

import (
	"fmt"
	"sample_go/memo"
	"strings"
)

func Search(keyword string) {
	memos, _ := memo.LoadMemos()
	for _, m := range memos {
		if strings.Contains(m.Text, keyword) {
			fmt.Printf("%d: %s\n", m.ID, m.Text)
			return
		}
	}
	fmt.Println("その言葉のメモはありません")
}
