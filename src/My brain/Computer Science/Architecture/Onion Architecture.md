# Onion Architecture

<aside>
💡 제어의 역전 원칙을 기반으로 도메인, 서비스 계층을 애플리케이션의 중심에 배치하고, 인프라스럭처를 외부에 배치하는 아키텍처.

</aside>

<aside>
💡 기존의 3계층 아키텍처는 모든 계층이 Data Access 레이어 위에 존재해 해당 레이어에 변경사항이 발생하면 모든 레이어에 변경이 발생하는 단점이 있음.
Onion Architecture에서는 데이터베이스 유형에 따라 달리지지 않는 저수준의 객체 모델만이 존재해, 데이터베이스의 실제 유형과 데이터 저장 방법은 Infrastructure 계층에서 결정

</aside>

**Domain Model**

가장 중심에 있는 레이어이며, 다른 레이어에 의존할 수 없음.

엔티티가 위치.

**Domain Services**

모델 레이어에 정의된 규약의 구현이 위치함.

**Application Services**

Infrastructure 레이어와 Domain Service 레이어의 브리지 역할을 함.

비즈니스 기능을 완성시키는 레이어.

Domain Service 레이어에 의존해 정의된 규약을 지켜야 함.

**Infrastructure/API** 

데이터베이스, 메시징, 알림, UI 등 도메인과 관련 없는 레이어.