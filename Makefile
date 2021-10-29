

all:
	protoc --proto_path=proto --go_out=protocol  proto/*.proto