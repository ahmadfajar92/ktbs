.PHONY : all clean format test cover

ALL_PACKAGES=$(shell go list ./...|grep -v -e mocks -e vendor -e schemas -e jsonschema -e external -e protogo -e grpc -e docs -e mocks_shared|tr '\n' ' ')

cover:
	@echo "mode: cover" > coverage-all.txt
	$(foreach pkg, $(ALL_PACKAGES) notifications/src/sms/vendors, \
	go test -coverprofile=coverage.out -covermode=set $(pkg);\
	tail -n +2 coverage.out >> coverage-all.txt;)
	go tool cover -func=coverage-all.txt | grep total:
	rm coverage.out

test:
	$(foreach pkg, $(ALL_PACKAGES) notifications/src/sms/vendors,\
	go test -race -short $(pkg);)

docker:
	docker-compose up

run:
	go run main.go
