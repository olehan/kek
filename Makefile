export FORCE_COLOR = true

PRE_COMMIT_HOOK_PATH = .git/hooks/pre-commit
PRE_COMMIT_TEMPLATE_HOOK_PATH = hooks/pre-commit.sh

.PHONY: bootstrap bootstrap-hooks test lint check

lint:
	golint -set_exit_status buffer colors config ds formatters levels names pool printer .

test:
	go test -v -race -coverprofile=coverage.txt -covermode=atomic ./...

check:
	make lint
	make test

bootstrap-hooks:
	make check-pre-commit-hook
	cp $(PRE_COMMIT_TEMPLATE_HOOK_PATH) $(PRE_COMMIT_HOOK_PATH)
	chmod +x $(PRE_COMMIT_HOOK_PATH)

bootstrap:
	make check-go-mod
	go mod download
	go mod vendor
	make check
	make bootstrap-hooks

check-go-mod:
ifneq ($(GO111MODULE), on)
	echo "Set GO111MODULE to 'on' first"
	exit 1
else
	echo "GO111MODULE is set"
endif

check-pre-commit-hook:
ifneq ("$(wildcard $(PRE_COMMIT_HOOK_PATH))", "")
	echo "Pre-commit hook already exists, reseting"
	rm $(PRE_COMMIT_HOOK_PATH)
else
	echo "Pre-commit hook is missing"
endif
