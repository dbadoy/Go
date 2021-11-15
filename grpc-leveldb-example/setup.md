# install protoc & craete pb
go get -u google.golang.org/protobuf/cmd/protoc-gen-go  
go install google.golang.org/protobuf/cmd/protoc-gen-go  

go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc  
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc  

protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative KVstore.proto
