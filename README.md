# Go Server Template

## Introduction
TBD

### Version
- Go `1.23.0`

### Package Overview
```
├── cmd
│   ├── web
│   ├── worker
│   └── ...
├── internal
│   ├── app
│   └── pkg
├── pkg
├── examples
└── local
```
패키지 구조는 [Standard Go Project Layout](https://github.com/golang-standards/project-layout)을 따릅니다.

* `cmd`: 실행 파일 패키지입니다.
  * `web`: 웹 서버 패키지입니다.
  * `worker`: 워커 서버 패키지입니다.
* `internal`: 내부 패키지입니다.
  * `app`: 내부 애플리케이션 코드가 존재하는 패키지입니다. 
  * `pkg`: 내부 라이브러리 패키지입니다.
* `pkg`: 외부 라이브러리 패키지입니다.
* `examples`: 예시 코드가 존재하는 패키지 입니다.
* `local`: 로컬 개발을 위한 스크립트 및 설정 파일이 존재하는 패키지 입니다.

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
