# SOLID Rules

**S = Seperation of Responsibility(책임 분담)**

클래스 또는 모듈을 변경하려면 오직 한 가지 이유가 있어야 하거나 혹은 동등하게 한 행위자에게만 책임이 있어야 함.

**O = Open-Closed(개폐)**

작성한 코드는 기존 코드를 수정하는 것이 아닌, 코드를 새롭게 추가해 새로운 기능을 추가할 수 있도록 해야 함.

**I = Interface Segregation(인터페이스 분리)**

클라이언트에게 적합하지 않은 인터페이스를 사용하도록 강요해서는 안됨.

하나의 큰 메서드 인터페이스 대신 하나, 두 개의 작은 메서드 인터페이스를 많이 갖는 것이 나쁠 것은 없음.

**L = Liskov Substitution(리스코프 대체)**

MIT의 컴퓨터 과학자인 Barbara Liskov의 이름을 따서 명명됨. 이 원칙은 기본 클래스 대신 파생 클래스를 사용할 수 있어야 함을 나타냄. 가장 중요한 것은 파생된 기본 클래스의 의미를 변경하려고 시도하면 안됨.

**D = Dependency Inversion(의존성 역전)**

높은 수준의 클래스는 낮은 수준의 클래스에 의존해서는 안됨. 대신 둘 다 추상화에 의존해야 함.