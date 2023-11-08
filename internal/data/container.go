package data

import (
	"github.com/armnerd/zcache/pkg/hash"
	"github.com/armnerd/zcache/pkg/list"
	"github.com/armnerd/zcache/pkg/set"
	"github.com/armnerd/zcache/pkg/zset"
)

// StringContainer string容器
var StringContainer = make(map[string]string)

// HashContainer hash容器
var HashContainer = make(map[string]*hash.Map)

// ListContainer list容器
var ListContainer = make(map[string]*list.List)

// SetContainer set容器
var SetContainer = make(map[string]*set.Set)

// ZsetContainer zset容器
var ZsetContainer = make(map[string]*zset.Zset)
