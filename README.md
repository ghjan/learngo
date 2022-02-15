# learngo

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

## worker可以多个
### 第二个爬取服务
cd /e/go-workspace/imooc_go/src/learngo
go run crawler_distributed/worker/crawlerworkerserver/crawlerworkerserver.go -port 1236

go run crawler_distributed/crawler_distributed.go -workerHosts=localhost:1234,localhost:1236 -itemSaverHost=localhost:1235

##开启前端
cd /e/go-workspace/imooc_go/src/learngo
go run crawler/frontend/starter.go -port=8888
http://localhost:8888