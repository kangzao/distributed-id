protoc --proto_path=. --micro_out=. --go_out=. greeter.proto

--proto_path是greeter.proto文件所在的路径
--micro_out是生成.go文件的所在目录，该文件被用于创建服务
--go_out是生成.go文件的所在目录，该文件被用于做数据序列化传输