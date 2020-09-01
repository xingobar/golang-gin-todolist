# golang-gin-todolist
gin todolist

migration 參考資源:

https://zhuanlan.zhihu.com/p/69472163
https://zhuanlan.zhihu.com/p/69472163

create:migrate create -ext sql -dir ./db/migrations/ -seq create_tags_table 
migrate: migrate -path db/migrations -database "mysql://root:@/gin_todo" -verbose up
rollback:  migrate -path db/migrations -database "mysql://root:@/gin_todo" -verbose down