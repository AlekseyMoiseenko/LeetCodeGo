test:
	go test -v -cover ./...

test_hash:
	go test -v -cover ./leetcode/hash_table/...

test_sorting:
	go test -v -cover ./leetcode/sorting/...

.PHONY: test test_hash test_sorting