generate:
	buf format -w 
	buf lint
	buf generate --path ./protobuf

publish:
	cd ./protobuf/gen/es/protobuf; npm publish