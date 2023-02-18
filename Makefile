.POSIX:

AWK = awk
GO = go
GOFMT = gofmt

all:

build:
	@echo "Building vt"
	@$(GO) build -v ./...

check:
	@echo "Testing vt"
	@$(GO) test -v ./...

check-fmt:
	@echo "Checking formatting"
	@$(GOFMT) -l . | $(AWK) '{ print } END { exit(NR > 0) }'

depends:
	@echo "Get dependencies"
	@$(GO) get -v -t -d ./...

fmt:
	@echo "Formatting Go files"
	@$(GO) fmt ./...
