PWD=$(shell pwd)
test:
	cd $(PWD)/internal/array;go test
	cd $(PWD)/internal/date;go test
	cd $(PWD)/internal/file;go test
