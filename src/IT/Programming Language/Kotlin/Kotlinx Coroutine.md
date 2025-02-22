# Kotlinx Coroutine

<aside>
💡 비동기 프로그래밍을 위한 라이브러리.
스레드보다 효율적으로 동작하며, 코드를 더 읽기 쉽고 유지보수하기 쉽게 만들어줌.

</aside>

**Scope**

- 코루틴이 실행될 수 있는 범위를 나타냄.
- **GlobalScope**
    - Application이 시작하고 종료될 때까지 유지되는 싱글톤 CorotuineScope.
    - 권장되지 않음.

[CoroutineScope%20db6ab3b6de76425581f4c45841409067](CoroutineScope%20db6ab3b6de76425581f4c45841409067)

[CoroutineContext%206067466184cf447d9528591cd92636e7](CoroutineContext%206067466184cf447d9528591cd92636e7)

[Job%2034891f35fd4f4867ac9ddf70799ca098](Job%2034891f35fd4f4867ac9ddf70799ca098)

[Channel%20d1eb6b6f6c074664ac53c7cc18c8cb63](Channel%20d1eb6b6f6c074664ac53c7cc18c8cb63)

[BlockingQueue%20411c50c8e34f4553804eae7dcd49877c](BlockingQueue%20411c50c8e34f4553804eae7dcd49877c)

**runBlocking**

- 처리를 모두 수행할 때까지 메인 스레드를 멈춤.
- 코루틴 스코프 없이도 정지 함수 실행 가능.
- 권장되지 않음.

**launch**

- 가장 간단한 형태의 코루틴 빌더.
- 비동기적으로 코루틴을 실행.
- Job 객체를 반환.

**async**

- 반환 값이 있는 Deferred 객체를 반환.
- 코루틴이 완료될 때 값을 반환하거나 예외를 던짐.
- 병렬로 여러 작업을 수행해 `Defered.await`으로 결과를 받아올 수 있음.

**withContext**

- 코루틴 내에서 코루틴 컨텍스트를 일시적으로 변경하는 데 사용.(주로 스레드를 변경)
- 블록 안에서 실행되는 코드는 지정된 컨텍스트에서 실행됨.
- 순차적으로 작업을 수행.

**yield**

코루틴이 실행되는 도중 다른 코루틴에게 양보함을 의미.

다른 코루틴이 실행되고 완료될 때까지 현재 코루틴이 일시적으로 중지되며, 실행이 계속됨.

**delay**

코루틴 내에서 일시적으로 일시 중단 하는 데 사용.

**continuation**

코루틴의 핵심.

현재 코루틴의 실행 상태를 나타내며, 중단 이후 실행을 재개하기 위해 사용됨.

return하지 않고 대신에 지속적으로 이어나갈 새 context.

[Flow%207ff3904dbbff41fa849c387ffbed116b](Flow%207ff3904dbbff41fa849c387ffbed116b)