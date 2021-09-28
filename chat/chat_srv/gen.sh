protoc -I=./proto --go_out=plugins=grpc:. proto/chat.proto
# go get github.com/favadi/protoc-go-inject-tag
protoc-go-inject-tag -input=proto/gen/chat_pb/chat.pb.go
cp proto/chat.proto ../chat_api/proto/
cp proto/gen/chat_pb/chat.pb.go ../chat_api/proto/gen/go/chat_pb/chat.pb.go