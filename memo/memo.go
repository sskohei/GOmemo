package memo

import (
	"encoding/json"
	"os"
)

type Memo struct {
	ID   int
	Tag  string
	Text string
}

const dataFile = "memo.json"

func LoadMemos() ([]Memo, error) {
	data, err := os.ReadFile(dataFile)
	if err != nil {
		if os.IsNotExist(err) {
			return []Memo{}, nil
		}
		return nil, err
	}

	var memos []Memo
	err = json.Unmarshal(data, &memos)
	return memos, err
}

func SaveMemos(memos []Memo) error {
	data, err := json.MarshalIndent(memos, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(dataFile, data, 0644)
}
