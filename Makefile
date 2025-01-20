all:
	# protoc --go_out=. --go-grpc_out=. \
	# 	--go_opt=paths=source_relative --go-grpc_opt=paths=source_relative \
	# 	hive_for_agent/*.proto \
	# 	proxy_for_agent/*.proto

	protoc \
		--go_out=$(CURDIR)/proto --go_opt=paths=source_relative \
		--go-xrpc_out=$(CURDIR)/proto --go-xrpc_opt=paths=source_relative \
		--plugin=../../../my/go-xrpc/go-xrpc/cmd/protoc-gen-go-xrpc/protoc-gen-go-xrpc \
		-I $(CURDIR)/proto \
		common/common.proto \
		hiveagent/hive.proto hiveagent/agent.proto \
		proxyagent/proxy.proto proxyagent/agent.proto
