package redis

import (
	"fmt"

	"github.com/go-redis/redis"
)

func main() {
	// 连接到 redis
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
}

// ============================ string 类型 ============================

// 设置key value expire
func set(client redis.Client) {
	err := client.Set("key", "value", 0).Err()
	if err != nil {
		return
	}
}

// 根据 key 获取 value
func get(client redis.Client) {
	value, err := client.Get("key").Result()
	if err == redis.Nil {
		fmt.Println("redis key is not existed")
	} else if err != nil {
		return
	} else {
		fmt.Println(value)
	}
}

// 根据 key 删除 value
func del(client redis.Client) {
	n, err := client.Del("key1", "key2", "key3").Result()
	if err != nil {
		return
	}
	fmt.Printf("成功删除了 %v 个key", n)
}

// 判断是否存在某 key
func exists(client redis.Client) {
	n, err := client.Exists("key1", "key2").Result()
	if err != nil {
		return
	}
	if n > 0 {
		fmt.Println("存在")
	} else {
		fmt.Println("不存在")
	}
}

// 设置key value，如果key已存在，返回false，如果key不存在，设置成功
func setnx(client redis.Client) {
	res, err := client.SetNX("key", "value", 0).Result()
	if err != nil {
		return
	}
	if res {
		fmt.Println("设置成功")
	} else {
		fmt.Println("设置失败，无事发生")
	}
}

// 根据key获取value，value+1
func incr(client redis.Client) {
	value, err := client.Incr("key").Result()
	if err != nil {
		return
	}
	fmt.Println(value)
}

// 根据key获取value，value+n，n可以自定义
func incrBy(client redis.Client) {
	value, err := client.IncrBy("key", 10).Result()
	if err != nil {
		return
	}
	fmt.Println(value)
}

// 根据key获取value，value-1
func decr(client redis.Client) {
	value, err := client.Decr("key").Result()
	if err != nil {
		return
	}
	fmt.Println(value)
}

// 根据key获取value，value-n，n可以自定义
func decrBy(client redis.Client) {
	value, err := client.DecrBy("key", 10).Result()
	if err != nil {
		return
	}
	fmt.Println(value)
}

// ============================ list 类型 ============================

// 从左侧将数据压入列表
func lpush(client redis.Client) {
	n, err := client.LPush("list", 1, 1, 2).Result()
	if err != nil {
		return
	}
	fmt.Println(n)
}

// 在指定值的前后插入数据
func lInsert(client redis.Client) {
	err := client.LInsert("list", "before", "100", 123).Err()
	if err != nil {
		return
	}
}

// 设置某下标的值为新的数据
func lSet(client redis.Client) {
	err := client.LSet("list", 1, 100).Err()
	if err != nil {
		return
	}
}

// 返回列表长度
func lLen(client redis.Client) {
	len, err := client.LLen("list").Result()
	if err != nil {
		return
	}
	fmt.Println(len)
}

// 根据下标返回值
func lIndex(client redis.Client) {
	err := client.LIndex("list", 0).Err()
	if err != nil {
		return
	}
}

// 根据下标范围返回值
func lRange(client redis.Client) {
	res, err := client.LRange("list", 0, 2).Result()
	if err != nil {
		return
	}
	fmt.Println(res)
}

// 从左侧弹出数据
func lPop(client redis.Client) {
	value, err := client.LPop("list").Result()
	if err != nil {
		return
	}
	fmt.Println(value)
}

// 删除n个值为xxx的数据
func lRem(client redis.Client) {
	n, err := client.LRem("list", 2, 100).Result()
	if err != nil {
		return
	}
	fmt.Println(n)
}

// ============================ set 类型 ============================

// 集合添加数据
func sAdd(client redis.Client) {
	client.SAdd("set", "value1", "value2")
}

// 删除集合中的数据
func sRem(client redis.Client) {
	n, err := client.SRem("set", "value1").Result()
	if err != nil {
		return
	}
	fmt.Println(n)
}

// 获取集合中所有的成员
func sMembers(client redis.Client) {
	value, err := client.SMembers("set").Result()
	if err != nil {
		return
	}
	fmt.Println(value)
}

// 判断成员是否在集合中
func sIsMember(client redis.Client) {
	exist, err := client.SIsMember("set", "value1").Result()
	if err != nil {
		return
	}
	if exist {
		fmt.Println("存在集合")
	} else {
		fmt.Println("不存在集合")
	}
}

// 返回集合有多少个成员
func sCard(client redis.Client) {
	count, err := client.SCard("set").Result()
	if err != nil {
		return
	}
	fmt.Println(count)
}

// ============================ zset 类型 ============================

// 集合添加数据
func zAdd(client redis.Client) {
	client.ZAdd("zset", redis.Z{
		Score:  0,
		Member: "pengjj",
	})
	client.ZAdd("zset", redis.Z{
		Score:  2,
		Member: "jjpeng",
	})
}

// 集合修改数据的score，+100
func zIncrBy(client redis.Client) {
	client.ZIncrBy("zset", 100, "pengjj")
}

// 返回集合中得分从小到大的前三个数据
func zRange(client redis.Client) {
	res, err := client.ZRange("zset", 0, 2).Result()
	if err != nil {
		return
	}
	fmt.Println(res)
}

// 返回集合中得分从大到小的前三个数据
func zRevRange(client redis.Client) {
	res, err := client.ZRevRange("zset", 0, 2).Result()
	if err != nil {
		return
	}
	fmt.Println(res)
}

// 返回集合中得分范围的数据
func zRangeByScore(client redis.Client) {
	res, err := client.ZRangeByScore("zset", redis.ZRangeBy{
		Min: "20",
		Max: "56",
	}).Result()
	if err != nil {
		return
	}
	fmt.Println(res)
}

// 返回集合的长度
func zCard(client redis.Client) {
	n, err := client.ZCard("zset").Result()
	if err != nil {
		return
	}
	fmt.Println(n)
}

// 删除集合的数据
func zRem(client redis.Client) {
	client.ZRem("zset", "2")
}

// ============================ hash 类型（存储对象） ============================

// 设置key+field的值为value
func hSet(client redis.Client) {
	client.HSet("user", "name", "pengjj")
}

// 根据key+field获取value
func hGet(client redis.Client) {
	name, err := client.HGet("user", "name").Result()
	if err != nil {
		return
	}
	fmt.Println(name)
}

// 根据key获取整个对象，返回是以map[string]interface{}的形式返回
func hGetAll(client redis.Client) {
	user, err := client.HGetAll("user").Result()
	if err != nil {
		return
	}
	fmt.Println(user)
}

// 根据key删除对象
func hDel(client redis.Client) {
	client.HDel("user")
}

// 判断是否存在key+field
func hExists(client redis.Client) {
	res, _ := client.HExists("user", "name").Result()
	if res {
		fmt.Println("存在")
	} else {
		fmt.Println("不存在")
	}
}
