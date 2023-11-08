package cmd

type RedisCmd string

const (
	SET           RedisCmd = "set"
	GET           RedisCmd = "get"
	DEL           RedisCmd = "del"
	HSET          RedisCmd = "hset"
	HGET          RedisCmd = "hget"
	HGETALL       RedisCmd = "hgetall"
	HKEYS         RedisCmd = "hkeys"
	HVALS         RedisCmd = "hvals"
	HDEL          RedisCmd = "hdel"
	LPUSH         RedisCmd = "lpush"
	LPOP          RedisCmd = "lpop"
	RPUSH         RedisCmd = "rpush"
	RPOP          RedisCmd = "rpop"
	LRANGE        RedisCmd = "lrange"
	LLEN          RedisCmd = "llen"
	SADD          RedisCmd = "sadd"
	SMEMBERS      RedisCmd = "smembers"
	SPOP          RedisCmd = "spop"
	SREM          RedisCmd = "srem"
	ZADD          RedisCmd = "zadd"
	ZRANGEBYSCORE RedisCmd = "zrangebyscore"
	ZSCORE        RedisCmd = "zscore"
	ZREM          RedisCmd = "zrem"
)

const (
	EXTRA_EXPIRE int = 0
)

var (
	CmdSet       []RedisCmd
	CmdAvailable map[RedisCmd]bool
)

func init() {
	CmdSet = []RedisCmd{
		SET, GET, DEL, // string
		HSET, HGET, HGETALL, HKEYS, HVALS, HDEL, // hash
		LPUSH, LPOP, RPUSH, RPOP, LRANGE, LLEN, // list
		SADD, SMEMBERS, SPOP, SREM, // set
		ZADD, ZRANGEBYSCORE, ZSCORE, ZREM, // zset
	}
	// 所有可用命令
	CmdAvailable = make(map[RedisCmd]bool)
	for _, v := range CmdSet {
		CmdAvailable[v] = true
	}
}
