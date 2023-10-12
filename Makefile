help:
	@echo "Available targets:"
	@echo "  build            Compile the Go code in the current directory."
	@echo "  buildall         Compile all Go code in all subdirectories."
	@echo "  generate         Run the Go generate tool on all packages."
	@echo "  lint             Run various linting and formatting tools."
	@echo "  test             Execute tests using gotestsum."
	@echo "  diagram          Generate a database schema diagram."
	@echo "  languagetool     Check or install LanguageTool and process spelling."
	@echo "  godoclint        Check Go source files for specific comment patterns."

check-go:
	@which go > /dev/null || (echo "Go is not installed. Please install Go." && exit 1)

install-lint-deps:
	@which golangci-lint > /dev/null || (echo "Required golangci-lint not found. Installing it..." && GO111MODULE=on go get github.com/golangci/golangci-lint/cmd/golangci-lint@latest)
	@which staticcheck > /dev/null || (echo "Required staticcheck not found. Installing it..." && GO111MODULE=on go get honnef.co/go/tools/cmd/staticcheck)

install-test-deps:
	@which gotestsum > /dev/null || (echo "Installing gotestsum..." && GO111MODULE=on go get gotest.tools/gotestsum@latest)

build: check-go
	go build -o singularity .

buildall: check-go
	go build ./...

generate: check-go
	go generate ./...

lint: check-go install-lint-deps
	gofmt -s -w .
	golangci-lint run --no-config --fix --disable-all -E tagalign
	golangci-lint run --fix
	staticcheck ./...

test: check-go install-test-deps
	go run gotest.tools/gotestsum@latest --format testname ./...

diagram: build
	./singularity admin init
	schemacrawler.sh --server=sqlite --database=./singularity.db --command=schema --output-format=svg --output-file=docs/database-diagram.svg --info-level=maximum

DIR_NAME := $(shell ls -d LanguageTool-* 2>/dev/null | head -n 1)

languagetool:
	if [ -z "$(DIR_NAME)" ]; then \
	    echo "LanguageTool is not installed. Installing..." ; \
		curl -L https://raw.githubusercontent.com/languagetool-org/languagetool/master/install.sh | bash ; \
	else \
		echo "LanguageTool seems to be already installed in $$DIR_NAME. Skipping installation." ; \
	fi
	-cp -f ./$(DIR_NAME)/org/languagetool/resource/en/hunspell/spelling.txt.bak ./$(DIR_NAME)/org/languagetool/resource/en/hunspell/spelling.txt
	cp -f ./$(DIR_NAME)/org/languagetool/resource/en/hunspell/spelling.txt ./$(DIR_NAME)/org/languagetool/resource/en/hunspell/spelling.txt.bak
	echo >>./$(DIR_NAME)/org/languagetool/resource/en/hunspell/spelling.txt
	cat ./spelling.txt >> ./$(DIR_NAME)/org/languagetool/resource/en/hunspell/spelling.txt
	java -jar ./$(DIR_NAME)/languagetool-commandline.jar -l en-US docs/en/README.md

godoclint:
	find . -path ./client -prune -o -name '*.go' ! -name '*_test.go' -exec grep -EHn '^// -|^// 1.|^//\s+[a-z]+:' {} \;
