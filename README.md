# golang-gin-todolist
gin todolist

migration 參考資源:

https://zhuanlan.zhihu.com/p/69472163
https://zhuanlan.zhihu.com/p/69472163

install migrate: brew install golang-migrate
create:migrate create -ext sql -dir ./db/migrations/ -seq create_tags_table 
migrate: migrate -path db/migrations -database "mysql://root:@/gin_todo" -verbose up
rollback:  migrate -path db/migrations -database "mysql://root:@/gin_todo" -verbose down

JWT 參考資源:
1. github.com/dgrijalva/jwt-go

Log 參考資源:
1. https://www.cnblogs.com/xinliangcoder/p/11212573.html
2. https://mojotv.cn/2018/12/27/golang-logrus-tutorial
