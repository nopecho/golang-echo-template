LOCAL_GOARCH := arm64
LOCAL_GOOS := darwin
BUILD_PATH := cmd
LOCAL_INFRA_PATH := $(CURDIR)/local
LOCAL_DOCKER_COMPOSE := $(LOCAL_INFRA_PATH)/docker-compose.yml
K6_SCRIPT_PATH := $(LOCAL_INFRA_PATH)/k6/script
K6_REPORT_PATH := $(LOCAL_INFRA_PATH)/k6/report

asdf:
	@asdf install && \
	while read -r tool version; do \
		asdf local $$tool $$version; \
	done < ./.tool-versions

.PHONY: k6-init
k6-init:
ifdef name
	@docker run \
	-v $(K6_SCRIPT_PATH):/script \
	-w /script \
	--rm -i grafana/k6 new $(name).js
else
	@echo "Usage: make k6-init name=<script-name>"
endif

.PHONY: k6-run
k6-run:
ifdef name
	@docker run \
	-v $(K6_REPORT_PATH):/report \
	-e K6_WEB_DASHBOARD=true \
	-e K6_WEB_DASHBOARD_EXPORT=/report/$(NOW)-$(name).html \
	--rm -i grafana/k6 run - <$(K6_SCRIPT_PATH)/$(name).js \
	--summary-trend-stats="p(99),p(95),avg,min,max" && \
	open $(K6_REPORT_PATH)/$(NOW)-$(name).html
else
	@echo "Usage: make k6-run name=<script-name>"
endif

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

.PHONY: build
build:
ifdef module
	CGO_ENABLED=0 \
	GOOS=$(LOCAL_GOOS) \
	GOARCH=$(LOCAL_GOARCH) \
	go build -o server $(BUILD_PATH)/$(module)/server.go
else
	@echo "Usage: make build module=<module name>"
	@exit 1
endif

.PHONY: docker
docker:
ifdef module
ifdef tag
	docker build \
	-f build/Dockerfile \
	-t $(tag) \
	--build-arg BUILD_PATH=$(BUILD_PATH)/$(module)/server.go \
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
	go run $(BUILD_PATH)/$(module)/server.go
else
	@echo "Usage: make run module=<module name>! starting default web server"
	go run $(BUILD_PATH)/web/server.go
endif