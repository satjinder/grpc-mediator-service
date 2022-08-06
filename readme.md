#GRPC Mediator Service
This service is a gRPC gateway that provides gRPC facade on HTTP based APIs. All it needs from every use case to define protobufs and provide a path to the the protos file descriptors.

go to med8r folder and run following command

generate file descriptor set for stats.proto
cd med8r
protoc --include_imports --descriptor_set_out="schemas/register/stats.proto-registry.pb" -Ischemas schemas/statsservice/stats.proto

cd schemas
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative **/*.proto

go to gserver folder

go run .


go to gclient folder

go run .
