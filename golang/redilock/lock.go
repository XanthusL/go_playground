// redigo/redis doesn't support redis cluster yet, go-redis/redis does
package redilock

import (
	"time"

	"github.com/gomodule/redigo/redis"
)

type RedisLock struct {
	Expire time.Duration
	Conn   redis.Conn
	Key    string
	Sleep  time.Duration
}

func GetLock(key string, conn redis.Conn, expire time.Duration) *RedisLock {
	return &RedisLock{
		Expire: expire,
		Conn:   conn,
		Key:    key,
		Sleep:  time.Second,
	}

}

func (l *RedisLock) Lock(token string) (bool, error) {
	for {
		ok, err := redis.String(l.Conn.Do("SET",
			l.Key, token, "NX", "PX", int64(l.Expire/time.Millisecond)))
		if ok == "OK" {
			return true, nil
		}
		if err != nil && err != redis.ErrNil {
			return false, err
		}
		time.Sleep(l.Sleep)
	}

}

const luaScript = `if redis.call("get",KEYS[1]) == ARGV[1]
then
    return redis.call("del",KEYS[1])
else
    return 0
end`

func (l *RedisLock) UnlockIfLocked(token string) (int, error) {
	script := redis.NewScript(1, luaScript)
	return redis.Int(script.Do(l.Conn, l.Key, token))
}
