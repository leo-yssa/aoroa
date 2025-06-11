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
go run main.go
```

서버는 기본적으로 8080 포트에서 실행됩니다.

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

### 2. 이슈 목록 조회 [GET] /issues
```bash
# 전체 이슈 조회
curl http://localhost:8080/issues

# 상태별 필터링
curl http://localhost:8080/issues?status=PENDING
```

### 3. 이슈 상세 조회 [GET] /issue/:id
```bash
curl http://localhost:8080/issue/1
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