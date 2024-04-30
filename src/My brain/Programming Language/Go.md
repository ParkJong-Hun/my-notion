# Go

<aside>
💡 2007년 구글이 개발한 오픈 소스 프로그래밍 언어.

</aside>

## 예약어

### Variable

- `var`
- `const`
- `type`
    - struct를 사용하지 않으면 그냥 typealias.
- `map`
- `struct`
- `range`
    - forEachIndexed.

### Package

- `package`
- `import`

### Function

- `func`
- `interface`
- `defer`
- `return`

<aside>
💡 Go 언어에서는 public 함수를 PascalCase, private 함수를 camelCase로 작성.

</aside>

### Condition

- `if`
- `else`
- `switch`
- `case`
- `default`
- `fallthrough`
    - 다음 case를 체크

### Route

- `goto`
- `for`
- `continue`
- `break`

### Concurrency

- `go`
- `chan`
- `select`
    - 채널에서 수신받았을 때의 처리 switch

## 문법

- `:=`
    - 선언 + 생성
- `*`
    - 포인터 선언, 값 접근
- `&`
    - 주소 접근
- `[{index}]Type`
    - 배열 선언
- `make`
    - 배열 동적 생성
- `<-`
    - 채널에 값 송신
- `close`
    - 채널 닫기
- `nil`
    - Null 값