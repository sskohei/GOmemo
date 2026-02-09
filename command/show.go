package command

import (
	"fmt"
	"sample_go/memo"
)

func Show(id int) {
	memos, _ := memo.LoadMemos()
	for _, m := range memos {
		if m.ID == id {
			fmt.Printf("%d: %s\n", m.ID, m.Text)
			return
		}
	}
	fmt.Println("指定したIDのメモは見つかりません")
}
