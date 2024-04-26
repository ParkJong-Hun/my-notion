# CoroutineScope

[Dispatcher%201388bc111f414893ac97f99db7b2a34f](Dispatcher%201388bc111f414893ac97f99db7b2a34f)

- 코루틴이 기본적으로 실행되는 스레드나 스레드 풀을 결정.

[Job%2034891f35fd4f4867ac9ddf70799ca098](Job%2034891f35fd4f4867ac9ddf70799ca098)

- 코루틴이 실행되는 동안의 상태를 추적하고, 코루틴을 취소하거나 완료될 때까지 대기하는 데 사용.

**(선택사항) ExceptionHandler**

- 코루틴 내에서 발생하는 예외를 처리하는 ExceptionHandler를 지정.

**(선택사항)** 

[CoroutineContext%206067466184cf447d9528591cd92636e7](CoroutineContext%206067466184cf447d9528591cd92636e7)

- Dispatcher, Job, ExceptionHandler 등을 포함.