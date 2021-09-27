protoc -I=proto/ --go_out=. proto/message.proto
protoc -I=proto/ --js_out=import_style=commonjs,binary:proto/gen/js proto/message.proto
cd proto/gen/js/
browserify exports.js -o message_pb_web.js
mv message_pb_web.js ../../../../static/im/other/message_pb_web.js