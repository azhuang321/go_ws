protoc -I=./proto --go_out=plugins=grpc:. proto/user.proto
cp proto/user.proto ../user_api/proto/
cp proto/gen/userpb/user.pb.go ../user_api/proto/gen/go/userpb/user.pb.go

cp proto/user.proto ../../chat/chat_api/proto/
cp proto/gen/userpb/user.pb.go ../../chat/chat_api/proto/gen/go/userpb/user.pb.go