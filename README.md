# go-cache

Golang 简易 Redis实现 ᕕ( ᐛ )ᕗ

### String

> SDS 一个带长度信息的字节数组

|   命令   |   描述   |
| ---- | ---- |
|   SET key value   |   设置指定 key 的值      |
|   GET key   |   获取指定 key 的值      |
|   DEL key   |   该命令用于在 key 存在时删除 key      |

###  List

* quicklist: 一个大的 ziplist, 切成段, 段与段之间双向链表连接
* ziplist: 元素挨元素, 并记录元素长度, 节省空间

|   命令   |   描述   |
| ---- | ---- |
|   LPUSH key value   |   将一个值插入到列表头部      |
|   LPOP key   |   移出并获取列表的第一个元素      |
|   RPUSH key value   |   将一个值插入到列表尾部      |
|   RPOP key   |   移出并获取列表的最后一个元素      |
|   LRANGE key start stop   |   获取列表指定范围内的元素      |
|   LLEN key   |   获取列表长度      |

### Hash[dict]

* 两个 hashtable 轮替 [扩容缩容时 rehash]
* hashtable: 一维数组加二维链表
* 链表里储存 key 以及 value

|   命令   |   描述   |
| ---- | ---- |
|   HSET key field value   |   将哈希表 key 中的字段 field 的值设为 value      |
|   HGET key field   |   获取存储在哈希表中指定字段的值     |
|   HGETALL key   |   获取在哈希表中指定 key 的所有字段和值     |
|   HKEYS key   |  获取所有哈希表中的字段      |
|   HVALS key   |  获取哈希表中所有值      |
|   HDEL key field   |  删除一个哈希表字段      |

### Set

> 使用 dict 储存 key, value 为 NULL

|   命令   |   描述   |
| ---- | ---- |
|   SADD key member   |   向集合添加一个成员      |
|   SMEMBERS key   |   返回集合中的所有成员      |
|   SPOP key   |   移除并返回集合中的一个随机元素      |
|   SREM key member   |   移除集合中一个成员      |

### Zset

* 使用 dict 储存 value 和 score 值的映射关系
* skiplist: 底层链表方便区间查询, 相当于反方向二叉树, 方便插入及查询

|   命令   |   描述   |
| ---- | ---- |
|   ZADD key score member   |   向有序集合添加一个成员，或者更新已存在成员的分数      |
|   ZRANGEBYSCORE key min max   |   通过分数返回有序集合指定区间内的成员      |
|   ZSCORE key member   |   返回有序集中，成员的分数值      |
|   ZREM key member   |   移除有序集合中的一个成员      |

## Thanks

|   Role   |   Package   |   Link   |
| ---- | ---- | ---- |
|   tcp   |   zinx       |   https://gitee.com/Aceld/zinx     |
|   hash/list/set   |   gods       |   https://github.com/emirpasic/gods     |
|   zset   |   skiplist       |   https://github.com/gansidui/skiplist     |
