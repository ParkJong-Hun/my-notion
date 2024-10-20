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

<aside>
⚠️

char는 무조건 4bytes.
u8으로 캐스팅 가능.

숫자 뒤에 u8 등 타입을 기입하면 해당 타입으로 인식됨.

</aside>

### Keyword

**mut**

- mutable variable.

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