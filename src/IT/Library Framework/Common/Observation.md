# Observation

<aside>
💡

기존 `@StateObject`와 `@ObservableObject`를 대체하는 Swift의 새로운 동적 데이터 감지 프레임워크.
변경 가능한 데이터의 변화를 자동으로 감지해 뷰를 업데이트하는 메커니즘 제공.

</aside>

`@Observable` 

- 컴파일러가 해당 타입의 속성 변경을 추적하는 코드 자동 생성
- 타입의 내부 속성이 변경될 때마다 알림 보낼 준비

`@Binadble`

- `@Observable` 타입의 인스턴스 뷰에 바인딩할 때 사용
- 데이터의 변경 사항을 뷰에 실시간 반영하도록 연결하는 역할

`withObservationTracking`

- 특정 클로저 내에서 어떤 속성들이 접근되었는지 감지하고, 해당 속성에 변경이 발생할 때마다 특정 동작 실행

`@ObservationIgnored` 

- 특정 속성을 관찰 대상에서 제외
- 앱 성능 최적화에 중요한 역할