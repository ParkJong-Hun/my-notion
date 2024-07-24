# BLoC

<aside>
💡 Business Logic Component.
Stream을 통해 상태를 업데이트.
stateful, stateless 개념이 사라지고 setState도 필요 없어짐.
각 UI 객체들은 BLoC 객체를 구독.
BLoC 객체의 상태가 변경되면 상태를 구독중인 UI 객체 들은 그 즉시 해당 상태로 UI를 변경.
BLoC 객체는 UI 객체로부터 이벤트를 전달 받으면, 필요한 Provider나 Repository로부터 데이터를 전달받아서 Business Logic을 처리.
러닝 커브가 높고, 보일러 플레이트 코드가 많다는 것이 단점.

</aside>