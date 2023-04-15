package redis

import "github.com/gomodule/redigo/redis"

type stringRds struct {
}

// 设置值
func (s *stringRds) Set(key string, value interface{}, expire ...int64) *Reply {
	conn := pool.Get()
	defer conn.Close()
	if len(expire) == 0 {
		return getReply(conn.Do("set", key, value))
	}
	return getReply(conn.Do("set", key, value, "ex", expire[0]))
}

// 获取值
func (s *stringRds) Get(key string) *Reply {
	conn := pool.Get()
	defer conn.Close()
	return getReply(conn.Do("get", key))
}

// key不存在是在设置值
func (s *stringRds) SetNX(key string, value interface{}) *Reply {
	conn := pool.Get()
	defer conn.Close()
	return getReply(conn.Do("setnx", key, value))
}

// 	设置并返回旧值
func (s *stringRds) GetSet(key string, value interface{}) *Reply {
	conn := pool.Get()
	defer conn.Close()
	return getReply(conn.Do("getset", key, value))
}

// 	设置key并指定生存时间
func (s *stringRds) SetEX(key string, value interface{}, seconds int64) *Reply {
	conn := pool.Get()
	defer conn.Close()
	return getReply(conn.Do("setex", key, seconds, value))
}

// 	设置key值并指定生存时间(毫秒)
func (s *stringRds) PSetEX(key string, value interface{}, milliseconds int64) *Reply {
	conn := pool.Get()
	defer conn.Close()
	return getReply(conn.Do("psetex", key, milliseconds, value))
}

// 设置子字符串
func (s *stringRds) SetRange(key string, value interface{}, offset int64) *Reply {
	conn := pool.Get()
	defer conn.Close()
	return getReply(conn.Do("setrange", key, offset, value))
}

// 	获取子字符串
func (s *stringRds) GetRange(key string, start, end int64) *Reply {
	conn := pool.Get()
	defer conn.Close()
	return getReply(conn.Do("getrange", key, start, end))
}

// 设置多个值
func (s *stringRds) MSet(kv map[string]interface{}) *Reply {
	conn := pool.Get()
	defer conn.Close()
	return getReply(conn.Do("mset", redis.Args{}.AddFlat(kv)))
}

// key不存在时设置多个值
func (s *stringRds) MSetNx(kv map[string]interface{}) *Reply {
	conn := pool.Get()
	defer conn.Close()
	return getReply(conn.Do("msetnx", redis.Args{}.AddFlat(kv)))
}

// 返回多个key的值
func (s *stringRds) MGet(keys []string) *Reply {
	conn := pool.Get()
	defer conn.Close()
	return getReply(conn.Do("mget", redis.Args{}.AddFlat(keys)...))
}

// 自增
func (s *stringRds) Incr(key string) *Reply {
	conn := pool.Get()
	defer conn.Close()
	return getReply(conn.Do("incr", key))
}

// 增加指定值
func (s *stringRds) IncrBy(key string, increment int64) *Reply {
	conn := pool.Get()
	defer conn.Close()
	return getReply(conn.Do("incrby", key, increment))
}

// 增加一个浮点值
func (s *stringRds) IncrByFloat(key string, increment float64) *Reply {
	conn := pool.Get()
	defer conn.Close()
	return getReply(conn.Do("incrbyfloat", key, increment))
}

// 自减
func (s *stringRds) Decr(key string) *Reply {
	conn := pool.Get()
	defer conn.Close()
	return getReply(conn.Do("decr", key))
}

// 自减指定值
func (s *stringRds) DecrBy(key string, increment int64) *Reply {
	conn := pool.Get()
	defer conn.Close()
	return getReply(conn.Do("decrby", key, increment))
}
