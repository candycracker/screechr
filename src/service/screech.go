package service

import (
	"errors"
	"time"
)

/*
● Id – positive integer, upper limit >1 quintillion, uniquely identifies the screech in the
system, cannot be null.
● Content – The content of the screech. 1024 characters, may be all whitespace but
cannot be null.
● Creator Id – The user id who created the screech. Not nullable.
● Date created – The date & time the user was added to the system. Cannot be null.
● Date modified – The most recent date & time the user was added to the system.
*/
type ScreechType struct {
	ID           int64
	Content      string
	CreatorID    int64
	CreatedTime  string //time.RFC3339
	ModifiedTime string //time.RFC3339
}

type screech struct {
	db          map[int64]*ScreechType
	screechList []*ScreechType
	userList    map[int64][]*ScreechType
	total       int64
}

type ScreechService interface {
	GetScreech(id int64) (ScreechType, error)
	GetScreeches(descending bool, uid int64) []ScreechType
	CreateScreech(uid int64, content string) error
	UpdateContent(id int64, content string) error
}

var screechDb = map[int64]*ScreechType{
	1: {
		ID:           1,
		Content:      "aaaaaaaaaa",
		CreatorID:    9987654321,
		CreatedTime:  "2022-08-11T17:41:59-07:00",
		ModifiedTime: "2022-08-11T17:41:59-07:00",
	},
	2: {
		ID:           2,
		Content:      "bbbbbbbbbbb",
		CreatorID:    9987654321,
		CreatedTime:  "2022-08-11T17:41:59-07:00",
		ModifiedTime: "2022-08-11T17:41:59-07:00",
	},
	3: {
		ID:           3,
		Content:      "cccccccccccc",
		CreatorID:    1234567890,
		CreatedTime:  "2022-08-11T17:41:59-07:00",
		ModifiedTime: "2022-08-11T17:41:59-07:00",
	},
}

var sList = []*ScreechType{{
	ID:           1,
	Content:      "aaaaaaaaaa",
	CreatorID:    9987654321,
	CreatedTime:  "2022-08-11T17:41:59-07:00",
	ModifiedTime: "2022-08-11T17:41:59-07:00",
}, {
	ID:           2,
	Content:      "bbbbbbbbbbb",
	CreatorID:    9987654321,
	CreatedTime:  "2022-08-11T17:41:59-07:00",
	ModifiedTime: "2022-08-11T17:41:59-07:00",
}, {
	ID:           3,
	Content:      "cccccccccccc",
	CreatorID:    1234567890,
	CreatedTime:  "2022-08-11T17:41:59-07:00",
	ModifiedTime: "2022-08-11T17:41:59-07:00",
}}
var uList = map[int64][]*ScreechType{
	1234567890: {{
		ID:           3,
		Content:      "cccccccccccc",
		CreatorID:    1234567890,
		CreatedTime:  "2022-08-11T17:41:59-07:00",
		ModifiedTime: "2022-08-11T17:41:59-07:00",
	}},
	9987654321: {
		{
			ID:           1,
			Content:      "aaaaaaaaaa",
			CreatorID:    9987654321,
			CreatedTime:  "2022-08-11T17:41:59-07:00",
			ModifiedTime: "2022-08-11T17:41:59-07:00",
		}, {
			ID:           2,
			Content:      "bbbbbbbbbbb",
			CreatorID:    9987654321,
			CreatedTime:  "2022-08-11T17:41:59-07:00",
			ModifiedTime: "2022-08-11T17:41:59-07:00",
		},
	},
}

func Screech() ScreechService {
	return &screech{
		db:          screechDb,
		screechList: sList,
		userList:    uList,
		total:       3,
	}
}

func (s *screech) GetScreech(id int64) (ScreechType, error) {
	screech, ok := s.db[id]
	if !ok {
		return ScreechType{}, errors.New("unable to find screech")
	}
	return *screech, nil
}

func (s *screech) GetScreeches(descending bool, uid int64) []ScreechType {
	var screeches []*ScreechType
	if uid > 0 {
		screeches = s.userList[uid]
	} else {
		screeches = s.screechList
	}
	var list []ScreechType
	count := 0
	if !descending {
		for _, s := range screeches {
			if count > 500 {
				return list
			}
			list = append(list, *s)
		}

		return list
	}
	for i := len(screeches) - 1; i >= 0; i-- {
		if count > 500 {
			return list
		}
		list = append(list, *screeches[i])
	}

	return list
}

func (s *screech) CreateScreech(uid int64, content string) error {
	var id int64 = s.total + 1
	newScreech := &ScreechType{
		ID:           id,
		Content:      content,
		CreatorID:    uid,
		CreatedTime:  time.Now().Format(time.RFC3339),
		ModifiedTime: time.Now().Format(time.RFC3339),
	}
	s.db[id] = newScreech
	s.screechList = append(s.screechList, newScreech)
	s.total++

	return nil
}
func (s *screech) UpdateContent(id int64, content string) error {
	screech, ok := s.db[id]
	if !ok {
		return errors.New("unable to find screech")
	}
	screech.Content = content
	screech.ModifiedTime = time.Now().Format(time.RFC3339)
	return nil
}
