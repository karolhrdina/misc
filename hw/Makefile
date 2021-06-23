
.PHONY: build all
.PRECIOUS: bin/%

bin/%:
	scripts/build.sh $(@:bin/%=%)

clean/%:
	rm -f bin/$(@:clean/%=%)

vet:
	go vet -source -all ./...

lint:
	golangci-lint run
