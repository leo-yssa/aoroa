# 이슈 관리 API

이 프로젝트는 이슈 관리 시스템의 REST API를 구현한 것입니다. Gin Framework를 사용하여 구현되었습니다.

## 요구사항

- Go 1.16 이상
- Gin Framework

## 설치 및 실행

1. 프로젝트 클론
```bash
git clone https://github.com/leo-yssa/aoroa.git
cd aoroa
```

2. 의존성 설치
```bash
go mod download
```

3. 서버 실행
```bash
go run cmd/api/main.go
```

서버는 기본적으로 8080 포트에서 실행됩니다.

## 프로젝트 구조

```
aoroa/
├── cmd/
│   └── api/
│       └── main.go
├── internal/
│   ├── domain/
│   │   ├── issue.go
│   │   └── user.go
│   ├── application/
│   │   └── issue_service.go
│   ├── infrastructure/
│   │   └── repository/
│   │       └── memory_repository.go
│   └── interfaces/
│       ├── http/
│       │   └── handlers/
│       │       └── issue_handler.go
│       └── repository/
│           └── issue_repository.go
├── pkg/
│   └── common/
│       └── errors/
│           └── errors.go
├── go.mod
├── go.sum
└── README.md
```

## 테스트 실행

단위 테스트를 실행하려면 다음 명령어를 사용하세요:
```bash
go test ./...
```

## 초기 사용자

시스템에는 기본적으로 아래 3명의 사용자가 등록되어 있습니다:
```json
[
    { "id": 1, "name": "김개발" },
    { "id": 2, "name": "이디자인" },
    { "id": 3, "name": "박기획" }
]
```

## API 엔드포인트

### 1. 이슈 생성 [POST] /issue
```bash
curl -X POST http://localhost:8080/issue \
  -H "Content-Type: application/json" \
  -d '{
    "title": "버그 수정 필요",
    "description": "로그인 페이지에서 오류 발생",
    "userId": 1
  }'
```

응답 예시 (201 Created):
```json
{
    "id": 1,
    "title": "버그 수정 필요",
    "description": "로그인 페이지에서 오류 발생",
    "status": "IN_PROGRESS",
    "user": { "id": 1, "name": "김개발" },
    "createdAt": "2025-06-02T10:00:00Z",
    "updatedAt": "2025-06-02T10:00:00Z"
}
```

### 2. 이슈 목록 조회 [GET] /issues
```bash
# 전체 이슈 조회
curl http://localhost:8080/issues

# 상태별 필터링
curl http://localhost:8080/issues?status=PENDING
```

응답 예시 (200 OK):
```json
{
    "issues": [
        {
            "id": 1,
            "title": "버그 수정 필요",
            "description": "로그인 페이지에서 오류 발생",
            "status": "PENDING",
            "user": { "id": 1, "name": "김개발" },
            "createdAt": "2025-06-02T10:00:00Z",
            "updatedAt": "2025-06-02T10:05:00Z"
        }
    ]
}
```

### 3. 이슈 상세 조회 [GET] /issue/:id
```bash
curl http://localhost:8080/issue/1
```

응답 예시 (200 OK):
```json
{
    "id": 1,
    "title": "버그 수정 필요",
    "description": "로그인 페이지에서 오류 발생",
    "status": "PENDING",
    "user": { "id": 1, "name": "김개발" },
    "createdAt": "2025-06-02T10:00:00Z",
    "updatedAt": "2025-06-02T10:05:00Z"
}
```

### 4. 이슈 수정 [PATCH] /issue/:id
```bash
curl -X PATCH http://localhost:8080/issue/1 \
  -H "Content-Type: application/json" \
  -d '{
    "title": "로그인 버그 수정",
    "status": "IN_PROGRESS",
    "userId": 2
  }'
```

응답 예시 (200 OK):
```json
{
    "id": 1,
    "title": "로그인 버그 수정",
    "description": "로그인 페이지에서 오류 발생",
    "status": "IN_PROGRESS",
    "user": { "id": 2, "name": "이디자인" },
    "createdAt": "2025-06-02T10:00:00Z",
    "updatedAt": "2025-06-02T10:10:00Z"
}
```

## 이슈 상태

- PENDING: 대기 중
- IN_PROGRESS: 진행 중
- COMPLETED: 완료
- CANCELLED: 취소

## 에러 응답

에러가 발생한 경우 다음과 같은 형식으로 응답됩니다:
```json
{
    "error": "에러 메시지",
    "code": 400
}
```