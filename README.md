# golang-gin-todolist
gin todolist

migration 參考資源:

https://zhuanlan.zhihu.com/p/69472163
https://zhuanlan.zhihu.com/p/69472163

migrate: migrate -path db/migrations -database "mysql://root:@/gin_todo" -verbose up
rollback:  migrate -path db/migrations -database "mysql://root:@/gin_todo" -verbose down