# Go Template

<img alt="Go" height="" src="https://go.dev/doc/gopher/gopherbw.png" width="300"/>

## Introduction
Go 서버 애플리케이션 개발 시 빠르게 사용할 수 있는 [echo](https://echo.labstack.com/) 프레임워크 기반 템플릿 프로젝트 입니다.

다음을 내용을 포함합니다.
* [버전 관리](#version-management)
* [패키지 구조](#package-layout)
* [각종 스크립트](#getting-started)
* 샘플 코드

## Dependency
이 템플릿에서 사용되는 라이브러리는 다음과 같습니다.

> golang 커뮤니티에선 되도록이면 표준 라이브러리리 사용을 권장합니다.
> 
> 하지만 최소한의 개발 편의를 위해 다음 라이브러리를 사용합니다.
> 
> _~~바퀴를 다시 발명하지 않는 것이 좋습니다.~~_

* Logging
  * [zerolog](https://github.com/rs/zerolog)
* Testing
  * [testify](https://github.com/stretchr/testify)
  * [testcontainers](https://golang.testcontainers.org/)
* HTTP Server
  * [echo](https://echo.labstack.com/)
* Database
  * [gorm](https://gorm.io/)

## Version Management

* go version: `1.23.0`

버전 관리 툴로 [asdf](https://asdf-vm.com/guide/introduction.html)를 사용합니다.

asdf가 설치되어 있지 않은 경우 [다음](#asdf)을 참고하여 설치합니다.

### version setting
```shell
make asdf
```

### asdf
>[asdf 설치](https://asdf-vm.com/guide/getting-started.html)
> 
>[golang plugin 설치](https://github.com/asdf-community/asdf-golang)

1. install asdf
```shell
brew install asdf
```
2. install plugin
```shell
asdf plugin add golang https://github.com/asdf-community/asdf-golang.git
```
3. setting `GOROOT`
```shell
. ~/.asdf/plugins/golang/set-env.zsh
```

## Package Layout
```
├─ cmd
│  ├── web
│  ├── worker
│  └── etc..
├─ internal
│  ├── app
│  │   ├── api
│  │   ├── domain
│  │   ├── infra
│  │   └── svc
│  └── pkg
├─ examples
└─ local
```
패키지 구조는 [Standard Go Project Layout](https://github.com/golang-standards/project-layout)을 따릅니다.

* `cmd`: 주요 애플리케이션 진입점. 이 계층은 서버를 시작하고 애플리케이션의 모든 계층을 연결하는 역할을 합니다.

* `internal`: 애플리케이션의 내부 패키지입니다.
  * `app`: 애플리케이션의 핵심 로직이 포함된 패키지입니다.
  * `pkg`: 애플리케이션의 여러 부분에서 사용할 수 있는 패키지입니다.

* `local`: 로컬 개발을 위한 스크립트 및 설정 파일이 존재하는 패키지입니다.

## Getting Started

### Run

```shell
make run module=<module name>
```

### Build

```shell
 make build module=<module name>
```

### Docker Build

```shell
 make docker module=<module name> tag=<tag name>
```

### Test

```shell
 make test
```

### Stress Test (with k6)

1. 스크립트 초기화
    ```shell
    make k6-init name=<script-name>
    ```

2. 스크립트 작성 📍`local/k6/script/<script-name>.js`

3. 스크립트 실행
    ```shell
    make k6-run name=<script-name>
    ```
