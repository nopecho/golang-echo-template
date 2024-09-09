# Go Server Template

![Go](https://golang.org/doc/gopher/frontpage.png)

## Introduction
Go 서버 애플리케이션 개발 시 사용할 수 있는 [echo](https://echo.labstack.com/) 프레임워크 기반 템플릿 프로젝트 입니다.

다음을 포함합니다.
* [버전 관리](#version-management)
* [패키지 구조](#package-layout)
* [각종 스크립트](#getting-started)
* 샘플 코드

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
├── cmd
│   ├── web
│   ├── worker
│   └── ...
├── internal
│   ├── app
│   └── pkg
├── build
├── pkg
├── examples
└── local
```
패키지 구조는 [Standard Go Project Layout](https://github.com/golang-standards/project-layout)을 따릅니다.

* `cmd`: 실행 파일 패키지입니다.
* `internal`: 애플리케이션에서 사용되는 내부 패키지입니다.
  * `app`: 내부 애플리케이션 코드가 존재하는 패키지입니다. 
  * `pkg`: 내부 라이브러리 패키지입니다.
* `build`: 빌드 파일이 존재하는 패키지입니다.
* `pkg`: 외부 라이브러리 패키지입니다.
* `examples`: 예시 코드가 존재하는 패키지입니다.
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
