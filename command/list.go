package command

import (
	"fmt"
	"sample_go/memo"
)

func List(num int) {
	memos, _ := memo.LoadMemos()
	if num > len(memos) {
		num = len(memos)
	}
	if len(memos) == 0 {
		fmt.Println("メモはありません")
		return
	}

	for i := len(memos) - 1; i >= len(memos)-num; i-- {
		m := memos[i]
		fmt.Printf("%d: %s\n", m.ID, m.Text)
	}
}
