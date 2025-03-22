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
- JvmStatic
    - companion object에서 만든 메소드는 실행시에 인스턴스화되므로 Java의 static 메소드는 아닌데, 이 것으로 Java의 Static처럼 사용 가능.
- JvmField
    - 속성에 대한 getter, setter를 생성하지 않고 Java처럼 직접 액세스.
- PublishedApi
    - 일반적으로는 internal이지만, inline 함수 내에서 사용되기 위해서 설정.

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
- DslMarker
    - 암묵적 외부 수신자 참조를 방지하여 안전한 DSL을 만들기 위해 사용

**Custom Annotation**

- 개발자가 직접 만든 Annotation.
- 주로 Reflection 활용, 코드 생성에 사용.