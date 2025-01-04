# Annotation

<aside>
💡

Jetpack의 어노테이션 관련 라이브러리.

</aside>

**@MainThread**

- 메인 스레드에서 실행되어야 함을 의미
- UI 업데이트, 이벤트 처리

**@WorkerThread**

- 백그라운드 스레드에서 실행되어야 함을 의미
- 네트워크, DB 작업

**@UiThread**

- UI 스레드에서 실행되어야 함을 의미
- UI 관련

**@AnyThread**

- 어느 스레드에서든 호출 가능을 의미
- 스레드 안전한 작업

**@BinderThread**

- 바인더 스레드에서 실행되어야 함을 의미
- AIDL, IPC 작업

**@RestrictTo**

- 특정 스코프의 스레드에서만 호출되어야 하며, 이를 명시적으로 제한