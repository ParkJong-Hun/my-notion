# Factory Pattern

<aside>
💡 객체 생성을 캡슐화하고 객체를 생성하기 위한 인터페이스를 제공하는 패턴.
클라이언트 코드는 구체적인 객체 생성 로직을 알 필요 없이 팩토리 메서드를 호출해 객체를 얻을 수 있음.
객체 생성 과정이 복잡하거나, 생성하는 방식을 변경하려는 경우에 사용.

</aside>

**Factory**

- 어떤 객체를 만들지를 서브클래스에 위임.

**AbstractFactory**

- 생성자는 비교적 간단하지만, 생성 객체에 여러 타입 패턴이 있는 경우 사용.
- 데이터를 넘기면 알아서 적합한 타입을 추론해 객체 생성.