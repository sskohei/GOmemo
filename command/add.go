package command

import (
	"flag"
	"fmt"
	"sample_go/memo"
)

func Add(text []string) {
	memos, _ := memo.LoadMemos()
	id := 1
	if len(memos) > 0 {
		id = memos[len(memos)-1].ID + 1
	}

	fs := flag.NewFlagSet("test", flag.ExitOnError)
	me := fs.String("m", "", "メモの内容")
	tag := fs.String("t", "", "タグの内容")
	fs.Parse(text)
	if *me == "" {
		fmt.Println("メモ内容を-mで指定してください")
		return
	}

	memos = append(memos, memo.Memo{
		ID:   id,
		Tag:  *tag,
		Text: *me,
	})

	memo.SaveMemos(memos)

	fmt.Println("メモ内容:", *me)
	fmt.Println("タグ内容:", *tag)
}
