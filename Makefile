LOCAL_GOARCH := arm64
LOCAL_GOOS := darwin
BUILD_PATH := cmd
LOCAL_INFRA_PATH := $(CURDIR)/local
LOCAL_DOCKER_COMPOSE := $(LOCAL_INFRA_PATH)/docker-compose.yml

asdf:
	@asdf install && \
	while read -r tool version; do \
		asdf local $$tool $$version; \
	done < ./.tool-versions

.PHONY: up
up:
	docker-compose -f $(LOCAL_DOCKER_COMPOSE) up -d

.PHONY: down
down:
	docker-compose -f $(LOCAL_DOCKER_COMPOSE) down

.PHONY: test
test:
	go test -v ./...

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: clean
clean:
	go clean -cache -modcache && go mod download

.PHONY: load
load:
	go mod download

.PHONY: pprof
pprof:
ifdef source
	@pprof -http=localhost:9999 $(source)
else
	@echo "Usage: make pprof source=<source> (e.g. http://localhost:6060/debug/pprof/heap)"
	@exit 1
endif

.PHONY: pprof-cpu
pprof-cpu:
	@pprof -http=localhost:9999 http://localhost:6060/debug/pprof/profile

.PHONY: pprof-heap
pprof-heap:
	@pprof -http=localhost:9999 http://localhost:6060/debug/pprof/heap

.PHONY: pprof-go
pprof-go:
	@pprof -http=localhost:9999 http://localhost:6060/debug/pprof/goroutine

.PHONY: build
build:
ifdef module
	CGO_ENABLED=0 \
	GOOS=$(LOCAL_GOOS) \
	GOARCH=$(LOCAL_GOARCH) \
	go build -o server $(BUILD_PATH)/$(module)/main.go
else
	@echo "Usage: make build module=<module name>"
	@exit 1
endif

.PHONY: docker
docker:
ifdef module
ifdef tag
	docker build \
	-t $(tag) \
	--build-arg BUILD_PATH=$(BUILD_PATH)/$(module)/main.go \
	--build-arg APP_NAME=$(module) \
	.
else
	@echo "Usage: make docker module=<module name> tag=<tag name>"
	@exit 1
endif
else
	@echo "Usage: make docker module=<module name> tag=<tag name>"
	@exit 1
endif

run: up
ifdef module
	go run $(BUILD_PATH)/$(module)/main.go
else
	@echo "Usage: make run module=<module name>! starting default api server"
	go run $(BUILD_PATH)/api/main.go
endif