export FORCE_COLOR = true

PRE_COMMIT_HOOK_PATH = .git/hooks/pre-commit
PRE_COMMIT_TEMPLATE_HOOK_PATH = hooks/pre-commit.sh

.PHONY: bootstrap benchmarks bootstrap-hooks test lint check

lint:
	golint -set_exit_status `go list ./... | grep -v formatters`

test:
	go test -v -race -coverprofile=coverage.txt -covermode=atomic `go list ./... | grep -v benchmarks`

check:
	make lint
	make test

bootstrap-hooks:
	make check-pre-commit-hook
	cp $(PRE_COMMIT_TEMPLATE_HOOK_PATH) $(PRE_COMMIT_HOOK_PATH)
	chmod +x $(PRE_COMMIT_HOOK_PATH)

bootstrap:
	make check
	make bootstrap-hooks

check-pre-commit-hook:
ifneq ("$(wildcard $(PRE_COMMIT_HOOK_PATH))", "")
	echo "Pre-commit hook already exists, reseting"
	rm $(PRE_COMMIT_HOOK_PATH)
else
	echo "Pre-commit hook is missing"
endif

benchmarks:
	go test -bench=. ./benchmarks/ > benchmarks.txt
