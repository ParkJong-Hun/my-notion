# RxCocoa

[RxSwift%208cbccbb2c31c4577bd9a8c4571c31e93](RxSwift%208cbccbb2c31c4577bd9a8c4571c31e93)

<aside>
💡 ReactiveX 프레임워크를 Swift 언어로 구현한 것.
RxSwift를 기반으로 iOS, macOS 앱 개발을 위해 확장한 라이브러리.
UIKit, Cocoa 프레임워크의 클래스, 프로퍼티를 RxSwift와 통합해 iOS와 macOS 개발에서 리액티브 프로그래밍을 더 쉽게 함.

</aside>

**Driver**

- UI 컴포넌트와 안전하게 데이터 바인딩.
- 메인 스레드에서만 실행됨.
- 에러를 방출하지 않음.

**Signal**

- 단방향 일회성 이벤트.

**Relay**

- Observable를 대체하는 간단한 버전의 Subject
- Observable를 생성하고 데이터를 방출.
- 중첩된 데이터는 방출되지 않고 onError를 방출하지 않음.
- accept를 사용.

**BehaviorRelay**

- 현재 값을 유지하면서 새 구독자에게 현재 값과 그 이후 값들을 전달.
- 초기값을 가짐.
- 종료되지 않고 항상 최신 값을 유지.

**PublishRelay**

- 현재 값을 유지하지 않고, 새 값만 전달.

**accept**

- Relay에서 onNext 대신 사용.

**bind**

- subscribe와 거의 동일.
- 디버그 모드에서 에러가 발생할 경우 fatalError가 발생하고, 릴리즈 모드에서 에러가 발생할 경우 일반 error 기록.