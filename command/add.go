package command

import (
	"fmt"
	"sample_go/memo"
)

func Add(text string) {
	memos, _ := memo.LoadMemos()

	id := 1
	if len(memos) > 0 {
		id = memos[len(memos)-1].ID + 1
	}

	memos = append(memos, memo.Memo{
		ID:   id,
		Text: text,
	})
	memo.SaveMemos(memos)
	fmt.Println("メモを追加しました 👍")
}
