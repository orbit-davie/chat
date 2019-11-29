all: compile

compile:
	protoc --go_out=../protocol ./*.proto
	protoc --go_out=plugins=grpc:../protocol ./*.proto
