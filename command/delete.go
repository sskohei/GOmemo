package command

import (
	"fmt"
	"sample_go/memo"
)

func Delete(id int) {
	memos, _ := memo.LoadMemos()
	newMemos := []memo.Memo{}

	found := false
	for _, m := range memos {
		if m.ID != id {
			newMemos = append(newMemos, m)
		} else {
			found = true
		}
	}

	if !found {
		fmt.Println("指定したIDのメモは見つかりません")
		return
	}

	memo.SaveMemos(newMemos)
	fmt.Println("メモを削除しました 🗑")
}
