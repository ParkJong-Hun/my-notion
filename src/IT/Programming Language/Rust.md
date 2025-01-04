# Rust

<aside>
💡

모질라 재단에서 개발한 저수준 시스템 프로그래밍 언어.
메모리 관리 안전성 강조.

</aside>

[Cargo%20125f37315c44800086a9c4fc48d7b453](Cargo%20125f37315c44800086a9c4fc48d7b453)

### Type

**u**

- unsigned integer.

**i**

- signed integer.
- 타입을 정의하지 않고 정수를 작성하면 기본적으로 i32.

**f**

- floating point.
- 웬만하면 계산 안정성을 위해 64가 좋음.

**&str**

- string slice(string literal).
- 문자열을 담는 불변형 타입.

**Vec<>**

- vector.
- 사이즈 변경 가능한 배열 타입.

**tuple**

- 여러 값을 하나의 단위로 묶어 표현할 수  있는 타입.

<aside>
⚠️

char는 무조건 4bytes.
u8으로 캐스팅 가능.

숫자 뒤에 u8 등 타입을 기입하면 해당 타입으로 인식됨.

</aside>

### Keyword

**let**

- 런타임 변수 선언.
- 기본적으로 immutable.

**const**

- 컴파일 타임 상수 선언.

**static**

- 전역 변수 선언.

**mut**

- mutable variable.

**&mut**

- mutable reference.
- 변경 가능하도록 참조를 전달.

**{}**

- display print.
- 문자열 안에 다른 타입의 값을 넣을 때.

**fn name(paramters) → returnType { code }**

- function.
- 내부 코드는 return이 없으면 마지막 작성한 내용이 return.

**()**

- empty tuple.
- unit type(void).

**{:?}**

- debug print.
- 문자열 안에 다른 타입의 type 이름이나 프로퍼티.

**unsafe**

- 플랫폼 의존적 코드나 메모리에 직접 접근 하는 코드의 wrapper.
- 생 포인터에 역참조.
- unsafe 함수, 메소드 호출.
- mutable static에 접근, 수정.

**trait**

- 인터페이스.
- impl로 다른 타입으로서 사용할 수 있도록 구현 가능.

**match**

- 스위치문.

**struct**

- 구조체.
- **unit struct**
    - 메모리 0인 아무것도 가지지 않는 struct

**enum**

- 열거형.

**loop**

- 반복문.

**impl**

- 구조체나 열거형, Trait 생성, 메서드 작성을 담당하는 것.

**Self**

- 자기 자신.

**use**

- Enum을 사용할 때, 타입 이름 생략 가능하도록 선언.

**where**

- 복잡한 generics 선언.