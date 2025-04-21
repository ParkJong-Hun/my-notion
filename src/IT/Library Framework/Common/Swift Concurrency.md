# Swift Concurrency

<aside>
💡

Swift 5.5에서 도입된 언어 수준의 비동기 프로그래밍 기능.
콜백 지옥을 피하며 코드 가독성을 높이고, 컴파일 타임의 많은 동시성 문제를 감지.

</aside>

### 주요 기능

- async/await
    - 비동기 코드를 동기 코드처럼 작성하는 문법
- Task
    - 비동기 작업의 생성과 관리를 위한 단위
- Actor
    - 공유 상태에 대한 안전한 접근을 보장하는 동시성 단위
- Structured Concurrency
    - 비동기 작업 간 관계를 명확하게 정의