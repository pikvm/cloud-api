all:
	protoc --go_out=. --go-grpc_out=. \
		--go_opt=paths=source_relative --go-grpc_opt=paths=source_relative \
		hive_for_agent/*.proto \
		proxy_for_agent/*.proto
