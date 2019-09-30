package tsRand

import (
	"math/rand"
	"time"
)

type RandMaker struct {
}

func Seed(value int64) {
	rand.Seed(value)
}

/*
 * @brief 获得随机数，在设定范围
 *
 * @param min 最小
 * @param max 最大
 * @return 随机数
 */
func RandInt(min int, max int) int {
	return min + rand.Intn(max-min+1)
}

//随机数获取
func RandNum(intn int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(intn)
}

// 生成随机数 - 必须要动态变化种子
func GenerateRangeNum(min int, max int) int {
	if min == max {
		return min
	}
	rand.Seed(time.Now().Unix())
	randNum := rand.Intn(max-min) + min
	return randNum
}

// 传入权重列表返回随机到的下标和是否为唯一值
func RandomByWeigh(weigh []int) (bool, int) {
	len := len(weigh)
	if len == 0 {
		return false, -1
	}
	if len == 1 {
		return true, 0
	}
	sum := 0
	for _, v := range weigh {
		sum += v
	}
	r := rand.Intn(sum-0) + 0
	t := 0
	for i, v := range weigh {
		t += v
		if t >= r {
			return false, i
		}
	}
	return false, -1
}
