package redis

import (
	redis "github.com/go-redis/redis/v8"
	logging "github.com/ipfs/go-log/v2"
	"github.com/switfs/switfs-block/config"
)

var (
	log      = logging.Logger("redis")
	RdClient *redis.Client
)

func New() error {
	// 创建Redis客户端
	RdClient = redis.NewClient(&redis.Options{
		Addr:     config.LotusConfig.Redis.Host,   // Redis服务器地址和端口
		Password: "",                              // Redis密码，如果没有设置密码则为空
		DB:       config.LotusConfig.Redis.DbName, // 选择要使用的数据库，默认为0
	})

	// 使用Ping命令检查与Redis服务器的连接

	return nil
	//
	//// 存储数据
	//err = rdb.Set(context.Background(), "key", "value", 0).Err()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//// 检索数据
	//val, err := rdb.Get(context.Background(), "key").Result()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println("key 的值为:", val)
	//
	//// 设置过期时间
	//err = rdb.Set(context.Background(), "key", "value", 5*time.Second).Err()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//// 检查key是否存在
	//exists, err := rdb.Exists(context.Background(), "key").Result()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println("key 是否存在:", exists)
	//
	//// 删除key
	//err = rdb.Del(context.Background(), "key").Err()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//// 再次检查key是否存在
	//exists, err = rdb.Exists(context.Background(), "key").Result()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println("key 是否存在:", exists)
}
