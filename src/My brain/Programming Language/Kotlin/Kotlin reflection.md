# Kotlin reflection

<aside>
💡 코틀린의 런타임에 프로그램의 클래스, 프로퍼티, 메소드에 접근하기 위해 사용되는 방법.
동적으로 클래스를 사용할 때, 작성 시점에 어떤 클래스를 사용해야할지 모르지만 런타임 시점에서는 실행해야하는 경우 필요.
IntelliJ에서는 기본적으로 의존성을 가짐.
런타임 과정에서만 알 수 있는 것들(어노테이션, Json 직렬화 등)

</aside>

**KClass**

- {클래스 명}::class로 획득 가능.
- 클래스 내의 정보를 담고 있음.

**KCallable**

- KFunction과 KProperty를 아울러 말하는 인터페이스.

**KFunction**

- ::{메소드 명}으로 획득 가능.
- 메소드 내의 정보를 담고 있음.

**KProperty**

- {클래스 명}::{프로퍼티 명}으로 획득 가능.
- 프로퍼티 내의 정보를 담고 있음.