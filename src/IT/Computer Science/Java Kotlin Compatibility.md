# Java / Kotlin Compatibility

`@JvmOverloads`

- 코틀린 함수의 기본값 매개변수를 자바에 쉽게 호출
- 컴파일러가 매개변수의 조합에 따라 오버로드된 메서드를 자동 생성

`@JvmStatic`

- 코틀린의 컴패니언 객체나 객체 선언 내의 함수를 자바의 정적(static) 메서드로 노출

`@JvmField`

- `lateinit`이나 `const`가 아닌 일반 프로퍼티를 자동으로 `private` 필드와 `public getter`/`setter` 메서드로 변환하지 않고 필드로 노출

`@JvmName` 

- 코틀린에서 함수나 프로퍼티의 컴파일된 이름을 변경 가능

`@JvmDefaultWithoutCompatibility`

- 자바 8 이전 버전과의 호환성을 고려하지 않고 순수한 자바 8 스타일의 디폴트 메서드로 컴파일

`@JvmRecord`

- 코틀린의 `data class`를 자바에 최근에 도입된 `Record` 로 변환

`@JvmInline`

- JVM 바이트코드 수준에서 `value class`의 최적화

`@JvmMultifileClass`

- 여러 파일에 분산된 최상위 함수들을 하나의 클래스로 묶음

`@Throws`

- 코틀린의 메서드가 던질 수 있는 체크드 예외를 명시적으로 선언 (컴파일에서 예외 처리하도록 경고)

`@Volatile`

- 코틀린 프로퍼티를 `volatile` 필드로 컴파일 (멀티 스레드 환경 보장)

`@Strictfp`

- 부동 소수점 연산이 모든 플랫폼에서 동일한 결과를 내도록 IEEE 754 표준을 준수