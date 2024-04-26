# Coroutine

<aside>
💡 비동기 프로그래밍을 위한 라이브러리.
스레드보다 효율적으로 동작하며, 코드를 더 읽기 쉽고 유지보수하기 쉽게 만들어줌.

</aside>

**Scope**

- 코루틴이 실행될 수 있는 범위를 나타냄.
- `GlobalScope`는 앱 전체의 생명주기에 걸쳐 사용될 수 있으나 권장되지 않음.
- `CoroutineScope`가 일반적으로 권장됨.

[CoroutineScope%20db6ab3b6de76425581f4c45841409067](CoroutineScope%20db6ab3b6de76425581f4c45841409067)

[CoroutineContext%206067466184cf447d9528591cd92636e7](CoroutineContext%206067466184cf447d9528591cd92636e7)

[Job%2034891f35fd4f4867ac9ddf70799ca098](Job%2034891f35fd4f4867ac9ddf70799ca098)

**launch**

- 가장 간단한 형태의 코루틴 빌더.
- 비동기적으로 코루틴을 실행.
- Job 객체를 반환.

**async**

- 반환 값이 있는 Deferred 객체를 반환.
- 코루틴이 완료될 때 값을 반환하거나 예외를 던짐.
- 병렬로 여러 작업을 수행해 `Defered.await`으로 결과를 받아올 수 있음.

**withContext**

- 코루틴 컨텍스트를 변경하는 데 사용.(주로 스레드를 변경)
- 블록 안에서 실행되는 코드는 지정된 컨텍스트에서 실행됨.

**yield**

코루틴이 실행되는 도중 다른 코루틴에게 양보함을 의미.

다른 코루틴이 실행되고 완료될 때까지 현재 코루틴이 일시적으로 중지되며, 실행이 계속됨.

**delay**

코루틴 내에서 일시적으로 일시 중단 하는 데 사용.

**continuation**

return하지 않고 대신에 지속적으로 이어나갈 새 context.