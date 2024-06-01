# Annotation

**종류**

**Built in Annotation**

- Kotlin, Android에 내장되어 있는 Annotation.
- Deprecated
    - 특정 함수, 클래스, 필드, 생성자를 더 이상 사용하지 말라고 경고.
- IntRange
    - 특정 함수, 파라미터, 필드 등의 Int 값의 범위를 제한.
- JvmOverloads
    - 함수 또는 생성자 파라미터에 default value가 설정되어 있으면 Compiler가 default value 만큼 오버로딩 함수를 만듦.

**Meta Annotation**

- Annotation에 대한 정보를 나타내기 위한 Annotation.
- Target
    - 어디에 사용 가능한지를 제한.
    - CLASS, FUNCTION, FIELD, TYPE
- Retention
    - Scope를 제한.
    - SOURCE
        - binary에는 포함하지 않음.
    - BINARY
        - reflection에서 접근 금지.
    - RUNTIME
        - 어디에든 접근 가능.
- Repeatable
    - 한 요소에 중복 사용될 수 있는지
- MustBeDocumented
    - Generated Documentation에 해당 Annotation도 포함될 수 있는지

**Custom Annotation**

- 개발자가 직접 만든 Annotation.
- 주로 Reflection 활용, 코드 생성에 사용.