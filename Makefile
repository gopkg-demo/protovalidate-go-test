.PHONY: generate
generate:
	@ rm -rf ./generated/protos
	@ buf generate

.PHONY: deps
deps:
	@ go mod tidy --compat=1.21