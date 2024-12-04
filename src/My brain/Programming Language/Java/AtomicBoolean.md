# AtomicBoolean

스레드 간의 안전성을 제공하기 위한 동기화된 원자적 연산을 지원하는 클래스.

여러 스레드에서 동시에 접근하더라도 안전하게 값을 읽고 쓸 수 있음.

안드로이드의 ViewModel이나 다른 구성 요소에서 스레드 안전성이 필요한 경우 유용함.

여러 스레드에서 동시에 ViewModel의 플래그를 업데이트하고 이를 관찰할 필요가 있다면 AtomicBoolean을 사용 할 수 있음.