# learngo

# 第一种配置 分布式爬虫配置1
## 开启两个服务：
### 第一个爬取服务
cd /e/go-workspace/imooc_go/src/learngo
go run crawler_distributed/worker/crawlerworkerserver/crawlerworkerserver.go -port 1234
### 第二个持久化服务
cd /e/go-workspace/imooc_go/src/learngo
go run crawler_distributed/persist/itemsaverserver/itemsaverserver.go -port 1235

## 运行分布式爬虫
cd /e/go-workspace/imooc_go/src/learngo
go run crawler_distributed/crawler_distributed.go -workerHosts=localhost:1234 -itemSaverHost=localhost:1235

# 第二种配置 分布式爬虫配置2
## worker可以多个

除了第一种配置里面的两个服务，还加上一个爬取服务
分布式爬虫的运行参数也要做相应调整

### 第二个爬取服务
cd /e/go-workspace/imooc_go/src/learngo
go run crawler_distributed/worker/crawlerworkerserver/crawlerworkerserver.go -port 1236

分布式爬虫的运行参数也要做相应调整:workerHosts里面添加新加的爬取服务
go run crawler_distributed/crawler_distributed.go -workerHosts=localhost:1234,localhost:1236 -itemSaverHost=localhost:1235

#开启前端 
可以搜索框里面输入查询条件查询内容

cd /e/go-workspace/imooc_go/src/learngo
go run crawler/frontend/starter.go -port=8888
http://localhost:8888


###举例而言：几个查询条件
Gender:(男) AND Age:(<=40) AND Age:(>18) AND Height:(>180) AND 20000
Gender:(女) AND Age:(<=30) AND Age:(>=16) AND Height:[150 TO 165] AND Weight:[40 TO 55]

Gender:(男) AND 有房 AND 有车
Gender:(男) AND House:(有房) AND Car:(有车)

###es中文分词插件的安装
/f/tmp/es/分词/中文分词.txt

对比后我们的结论是：ik_smart既能满足英文的要求，又更智能更轻量，占用存储最小，所以首推ik_smart；
standard对英语支持是最好的，
但是对中文是简单暴力每个字建一个反向索引，浪费存储空间而且效果很差；
ik_max_word比ik_smart对中文的支持更全面，但是存储上的开销实在太大，不建议使用。

链接：https://www.jianshu.com/p/bb89ad7a7f7d
