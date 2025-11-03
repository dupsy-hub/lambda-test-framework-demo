# List only handler packages for tests & coverage
TEST_PKGS := $(shell go list ./internal/handlers/...)
COVERPKG  := $(TEST_PKGS)

COVERFILE := coverage.out
COVERHTML := coverage.html

.PHONY: all tools mocks unit cover bench clean

all: unit

tools:
	@command -v mockgen >/dev/null 2>&1 || go install github.com/golang/mock/mockgen@v1.6.0

mocks: tools
	mockgen -destination=mocks/mock_s3.go -package=mocks \
		unit-test/internal/awsiface S3PutObject

unit:
	# Run tests only for handlers and calculate coverage only on handlers
	go test -race -covermode=atomic -coverpkg=$(COVERPKG) -coverprofile=$(COVERFILE) $(TEST_PKGS)
	./scripts/coverage_check.sh 90 $(COVERFILE)

cover: unit
	go tool cover -func=$(COVERFILE)
	go tool cover -html=$(COVERFILE) -o $(COVERHTML)

bench:
	go test -run=^$ -bench=. -benchmem $(TEST_PKGS)

clean:
	rm -f $(COVERFILE) $(COVERHTML)
