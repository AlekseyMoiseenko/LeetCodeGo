test:
	go test -v -cover ./...

test_hash:
	go test -v -cover ./leetcode/hash_table/...

.PHONY: test test_hash