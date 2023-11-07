package zset

import (
	"encoding/json"
	"fmt"

	"github.com/armnerd/zcache/internal/hash"
)

// Zset 有序集合
type Zset struct {
	scoreSet *SkipList
	keyMap   *hash.Map
}

// Iterm 数据单元
type Iterm struct {
	score int
	id    string
}

// Less 比较函数
func (u *Iterm) Less(other interface{}) bool {
	if u.score > other.(*Iterm).score {
		return true
	}
	return false
}

// NewZset instantiates a zset.
func NewZset() *Zset {
	return &Zset{
		scoreSet: NewSkipList(),
		keyMap:   hash.New(),
	}
}

// Put inserts element into the zset.
func (z *Zset) Put(key string, score int) {
	var prevScore, found = z.keyMap.Get(key)
	if found {
		if prevScore != score {
			z.scoreSet.Delete(&Iterm{prevScore.(int), key})
			z.scoreSet.Insert(&Iterm{score, key})
		}
	} else {
		z.scoreSet.Insert(&Iterm{score, key})
	}
	z.keyMap.Put(key, score)
	return
}

// Members get all member from zset.
func (z *Zset) Members(min int, max int) string {
	var mapInstances []map[string]interface{}
	for e := z.scoreSet.Front(); e != nil; e = e.Next() {
		instance := map[string]interface{}{fmt.Sprint(e.Value.(*Iterm).id): e.Value.(*Iterm).score}
		mapInstances = append(mapInstances, instance)
	}
	jsonStr, _ := json.Marshal(mapInstances)
	return string(jsonStr)
}

// GetScore find score by member.
func (z *Zset) GetScore(member string) string {
	var value, err = z.keyMap.Get(member)
	if !err {
		return "not found"
	}
	return fmt.Sprint(value)
}

// RemoveMember remove member.
func (z *Zset) RemoveMember(member string) string {
	var score, err = z.keyMap.Get(member)
	if !err {
		return "not found"
	}
	z.keyMap.Remove(member)
	z.scoreSet.Delete(&Iterm{score.(int), member})
	return "ok"
}
