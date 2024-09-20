generate:
	buf generate --path ./protobuf

publish:
	cd ./protobuf/gen/es/protobuf; npm publish