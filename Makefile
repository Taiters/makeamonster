.PHONY: proto
proto:
	protoc \
		--js_out=import_style=commonjs,binary:client/src \
		--go_out=server \
		--go_opt=module=github.com/Taiters/monstermaker \
		proto/messages.proto
