# Swift-Dependencies

<aside>
💡

Point-Free에서 개발한 Swift 앱에서 의존성 주입을 안전하고 효율적으로 관리하기 위해 설계된 패키지.

</aside>

`@Dependency` 

- 의존성을 선언적으로 주입
- `testValue`
    - 테스트 환경에서 별도의 설정 없이 목 또는 스텁 구현 자동 주입

`DependencyKey` 

- 의존성을 정의할 때 사용해야하는 키 프로토콜
- `liveValue`
    - 실제 앱에서 사용된 운영 환경 구현체
- `testValue`
    - 테스트 환경에서 사용될 테스트용 구현체
- `previewValue`
    - Swift 프리뷰를 위한 구현체

`@DependencyClient` 

- 의존성 인터페이스를 정의하는 데 필요한 상용구 코드를 자동으로 생성