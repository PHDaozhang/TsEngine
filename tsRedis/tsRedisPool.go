package tsRedis

import (
	"errors"
	"github.com/astaxie/beego/logs"
	"sync"
	"time"

	"github.com/go-redis/redis"
)

var gTsRedisPool *TsRedisPool

type TsRedisPool struct {
	objs  []*redis.Client
	busy  []bool
	addr  string
	db    int
	pwd   string
	count int

	lock sync.Mutex
}

// Deprecated: use tsRedis.InitPool instead.
func InitTsRedisPool(addr string, db int, pwd string, count int) error {
	return InitPool(addr, db, pwd, count)
}

func InitPool(addr string, db int, pwd string, count int) error {
	ret := &TsRedisPool{
		objs:  make([]*redis.Client, count),
		busy:  make([]bool, count),
		addr:  addr,
		db:    db,
		pwd:   pwd,
		count: count,
	}

	for i := 0; i < count; i++ {
		obj := redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: pwd,
			DB:       db,
		})

		_, err := obj.Ping().Result()

		if err != nil {
			return err
		}

		ret.objs[i] = obj
		ret.busy[i] = false

		obj.Get("test...test...")
	}

	gTsRedisPool = ret

	logs.Info("redis pool size:", count, " addr:", addr, " index:", db)
	return nil
}

func getRedis() *redis.Client {
	gTsRedisPool.lock.Lock()
	defer gTsRedisPool.lock.Unlock()
	for i := 0; i < len(gTsRedisPool.busy); i++ {
		if !gTsRedisPool.busy[i] {
			gTsRedisPool.busy[i] = true
			return gTsRedisPool.objs[i]
		}
	}
	return nil
}

func giveBackRedis(rb *redis.Client) {
	gTsRedisPool.lock.Lock()
	defer gTsRedisPool.lock.Unlock()
	for i := 0; i < len(gTsRedisPool.busy); i++ {
		if gTsRedisPool.objs[i] == rb {
			gTsRedisPool.busy[i] = false
			return
		}
	}
}

func Set(key string, value interface{}, second int64) error {
	cb := getRedis()
	if cb == nil {
		return errors.New("no free db")
	}
	defer giveBackRedis(cb)

	t := time.Duration(second) * time.Second

	err := cb.Set(key, value, t).Err()
	return err
}

func SetNX(key string, value interface{}, second int64) (err error) {
	cb := getRedis()
	if cb == nil {
		return errors.New("no free db")
	}
	defer giveBackRedis(cb)

	t := time.Duration(second) * time.Second

	err = cb.SetNX(key, value, t).Err()

	return
}

/**
- 不影响现在模式下，拷贝出来的
- 这个函数的精髓就是要返回是否设置成功
- 重复setnx则不允许，说明当前key还在使用中
*/
func SetNX2(key string, value interface{}, second int64) (isOk bool, err error) {
	cb := getRedis()
	if cb == nil {
		return false, errors.New("no free db")
	}
	defer giveBackRedis(cb)

	t := time.Duration(second) * time.Second

	isOk, err = cb.SetNX(key, value, t).Result()

	return
}

func HSet(key string, SubKey string, value interface{}, exp ...int) error {
	cb := getRedis()
	if cb == nil {
		return errors.New("no free db")
	}
	defer giveBackRedis(cb)

	err := cb.HSet(key, SubKey, value).Err()
	return err
}

func HGet(key string, SubKey string) (string, error) {

	cb := getRedis()

	defer giveBackRedis(cb)

	cmd := cb.HGet(key, SubKey)
	value := cmd.Val()
	err := cmd.Err()

	return value, err
}

func HExists(key string, SubKey string) (bool, error) {

	cb := getRedis()

	defer giveBackRedis(cb)

	return cb.HExists(key, SubKey).Result()
}

func HDel(key string, SubKey string) (delCount int64, err error) {

	cb := getRedis()

	defer giveBackRedis(cb)

	return cb.HDel(key, SubKey).Result()
}

func HMSet(key string, mapValue map[string]interface{}, second int64) error {
	cb := getRedis()
	if cb == nil {
		return errors.New("no free db")
	}
	defer giveBackRedis(cb)

	t := time.Duration(second) * time.Second

	err := cb.HMSet(key, mapValue).Err()
	if err != nil {
		return err
	}

	return cb.Expire(key, t).Err()
}

func HGetAll(key string) (map[string]string, error) {
	cb := getRedis()
	if cb == nil {
		return map[string]string{}, errors.New("no free db")
	}
	defer giveBackRedis(cb)

	return cb.HGetAll(key).Result()
}

func HIncrBy(key string, filed string, num int64) error {
	cb := getRedis()
	if cb == nil {
		return errors.New("no free db")
	}
	defer giveBackRedis(cb)

	return cb.HIncrBy(key, filed, num).Err()
}

func Get(key string) (string, error) {

	cb := getRedis()

	defer giveBackRedis(cb)

	val, err := cb.Get(key).Result()

	return val, err
}

func MGet(keys ...string) []interface{} {

	cb := getRedis()

	defer giveBackRedis(cb)

	vals := cb.MGet(keys...).Val()

	return vals
}

func Del(key string) error {
	cb := getRedis()
	if cb == nil {
		return errors.New("no free db")
	}
	defer giveBackRedis(cb)

	err := cb.Del(key).Err()
	return err
}

func Exists(keys ...string) (int64, error) {
	cb := getRedis()
	if cb == nil {
		return 0, errors.New("no free db")
	}
	defer giveBackRedis(cb)

	cmd := cb.Exists(keys...)
	ex := cmd.Val()
	err := cmd.Err()
	return ex, err
}

func Expire(key string, second int64) error {
	cb := getRedis()
	if cb == nil {
		return errors.New("no free db")
	}
	defer giveBackRedis(cb)
	cmd := cb.Expire(key, time.Duration(second)*time.Second)
	err := cmd.Err()
	return err
}

func Keys(key string) ([]string, error) {
	cb := getRedis()
	if cb == nil {
		return nil, errors.New("no free db")
	}
	defer giveBackRedis(cb)
	cmd := cb.Keys(key)
	values := cmd.Val()
	err := cmd.Err()
	return values, err
}

func RPush(key string, value string) error {
	cb := getRedis()
	if cb == nil {
		return errors.New("no free db")
	}
	defer giveBackRedis(cb)
	err := cb.RPush(key, value).Err()
	return err
}

func LRange(key string, start int64, end int64) ([]string, error) {
	cb := getRedis()
	defer giveBackRedis(cb)
	result, err := cb.LRange(key, start, end).Result()

	return result, err
}
