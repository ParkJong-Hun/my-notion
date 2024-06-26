# Dagger

<aside>
💡 DI를 위한 자바, 안드로이드 프레임워크.
Google에서 개발한 DI 프레임워크로, 컴파일 타임에 DI 하는 코드를 자동으로 생성.
런타임에 동적으로 의존성을 주입하는 일반적인 방법과는 다름.

</aside>

## Module

의존성을 제공하는 방법을 정의하는 클래스. 각 모듈은 @Module로 표시되며, 모듈 내에서 @Provides 어노테이션을 사용해 의존성을 생성, 제공.

## Component

의존성을 요청하는 곳과 의존성을 제공하는 모듈을 연결하는 인터페이스. 각 컴포넌트는 @Component로 표시되며, @Component 어노테이션에 modules 매개변수를 통해 사용할 모듈을 지정. 컴포넌트는 의존성을 주입받을 수 있는 메서드를 가지고 있음.

## Inject

의존성 주입을 받을 필드, 생성자 또는 메서드에 @Inject 어노테이션을 사용.

[Hilt%20726c4a96eb9342fd99b3d10dd8d08b78](Hilt%20726c4a96eb9342fd99b3d10dd8d08b78)