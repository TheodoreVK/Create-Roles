init:
	go mod init ksuser
	
gen:
	protoc proto/roles/*.proto --go_out=pb --go_opt=paths=source_relative --go-grpc_out=require_unimplemented_servers=false:pb --go-grpc_opt=paths=source_relative -I=proto --experimental_allow_proto3_optional
.PHONY: init gen