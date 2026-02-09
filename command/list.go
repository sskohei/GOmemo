package command

import (
	"fmt"
	"sample_go/memo"
)

func List() {
	memos, _ := memo.LoadMemos()
	if len(memos) == 0 {
		fmt.Println("メモはありません")
		return
	}

	for _, m := range memos {
		fmt.Printf("%d: %s\n", m.ID, m.Text)
	}
}
