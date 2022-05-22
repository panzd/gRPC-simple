gen:
	protoc --go_out=.\pb --go-grpc_out=.  .\proto\*.proto

clean:
	rm pb\*.go

run:
	go run main.go