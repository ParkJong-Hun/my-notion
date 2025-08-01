# Actor Model

<aside>
💡

Erlang에서 유래된 Actor 들이 서로에게 메시지를 비동기적으로 주고받으며 상태를 관리하는 구조.

</aside>

### 구성 요소

- Actor
    - 독립적인 상태를 가진 실행 단위
- Message
    - Actor에게 전달되는 명령 또는 요청
- Context
    - Actor의 실행 환경
- Handler
    - 메세지를 받아서 처리하는 로직