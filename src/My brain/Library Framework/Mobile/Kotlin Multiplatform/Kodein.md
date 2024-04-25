# Kodein

**DI**

DI Container 블록.

**DIAware**

많은 기능을 잠금 해제한 DI.

**DI.module**

DI Container에 가져올 수 있는 모듈.

**Direct**

기본적으로 DI는 lazy하지만, 즉각적으로 원하는 경우에 Direct를 사용하면 됨.

**bindProvider**

항상 새 구현 인스턴스 생성해서 반환.

**bindSingleton**

구현 인스턴스를 하나만 생성하고 항상 이 인스턴스를 반환.

**bindEagerSingleton**

제공된 함수로 인스턴스 생성이 바로 실행된다는 점 외에는 일반적인 싱글톤과 같음.

**bindMultiton**

어떤 클래스의 인스턴스가 여러 개 필요한 경우 사용. Singleton과 달리 Multiton은 이름 또는 키를 사용해 각각의 인스턴스를 식별.

ex / `bindMultiton<식별자, 인터페이스> = 구현체(식별자, …)`로 작성

**bindIsntance**

이미 존재하는 인스턴스를 타입을 통해 바인드.

**bindConstant**

특정 태그를 이용해 상수값(원시 타입)을 반환.

**bindFactory**

정의된 타입의 인수를 가져와서 바인딩된 유형의 객체를 반환.

**instance**

bind한 것에 대한 인스턴스를 가져올 때 사용. multiton일 경우 키를 전달해야 함.

**injection**

종속성이 주입되면, 생성 시 클래스에 종속성이 제공됨.

**retrieved**

종속성이 검색되면, 클래스는 자체 종속성을 가져오는 책임을 가짐.