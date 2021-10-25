
create:
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		proto/article.proto


db:
	sudo service postgresql stop
	sudo docker run --name postgres -p 5432:5432 -e POSTGRES_PASSWORD=123456p  --hostname db --network mynet -d postgres

load:
	go run cmd/micro-videogame/main.go