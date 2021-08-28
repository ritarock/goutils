PWD=$(shell pwd)
test:
	cd $(PWD)/lib/array;go test
	cd $(PWD)/lib/date;go test
	cd $(PWD)/lib/file;go test
