# Macro

<aside>
💡

Rust에서 코드를 입력받아 다른 코드를 생성하는 메타프로그래밍 도구.

런타임에 실행되는 함수와 다르게, 컴파일 타임에 실행되어 코드로 전환되어서 런타임 오버헤드가 없으며, 코드 중복을 제거하고 타입 안정적.

</aside>

### 종류

- Declarative Macros
    - `macro_rules!`로 정의하는 패턴 기반 매크로.
- Procedural Macros
    - Functional-like Macro
        - `my_macro!()`
        - `vec!`
        - `format!`
        - `panic!`
        - Cargo.toml에서 proc-macro를 true로 설정해야 커스텀 Derive 매크로를 만들 수 있음
    - Derive Macro
        - `#[derive(Trait)]`
    - Attribute Macro
        - `#[my_attribute]`

### 사용 사례

- 새로운 구문 생성
- 컴파일 타임 코드 생성
    - 함수 자동 생성
- 가변 개수 인자 처리
- 타입 정보에 기반한 코드 생성
- 조건부 컴파일과 결합

### 사용 불가능한 사례

- 런타임 값에 따른 동적 처리
    - if 문 처리 등
- 값 기반 재귀 호출
- 트레이트와 제네릭을 통한 추상화
- 클로저와 고차 함수