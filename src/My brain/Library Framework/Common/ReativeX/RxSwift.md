# RxSwift

[Swift%209eebd2ca15aa466aaa6e9c5f0957d611](Swift%209eebd2ca15aa466aaa6e9c5f0957d611)

<aside>
💡 ReativeX 프레임워크의 반응형 프로그래밍을 위한 스위프트 언어용 라이브러리.
데이터 흐름을 중심으로 애플리케이션을 구축하는 방식으로, 데이터나 이벤트의 변화를 감지하고 이에 반응해 코드를 실행하는 패러다임.

</aside>

### 핵심 개념

- **Observable**

데이터를 발행. 관찰 가능한 객체. 데이터 스트림을 나타냄.

이벤트 스트림이나 비동기 작업의 결과물 등을 나타낼 수 있음.

**disposed()**

DisposeBag의 Observable의 구독을 해제.

Observable을 구독할 때 반환 값인 Disposable이 DisposableBag에 추가되어 수명이 관리됨.

**onNext()**

새로운 데이터를 발행.

**onError()**

에러일 때 실행.

**onCompleted()**

데이터가 성공적으로 발행되었을 때.

- **Observer**

관찰자. Observable를 구독하고, Observable로부터 발생하는 이벤트나 데이터 변경을 처리하는 객체.

- **Operator**

연산자. RxSwift가 제공하는 것으로. Observable을 변환, 조합, 필터링, 시간 지연, 에러 처리 등 다양한 작업을 수행 가능.

- **Subject**

서브젝트. Observable와 Observer의 역할을 모두 수행하는 객체.

데이터를 보내고 받을 수 있음.

- **Scheduler**

스케줄러. RxSwift가 작업을 어떤 스레드나 큐에서 실행할지 제어하기 위한 스케줄러를 제공.

비동기 작업을 관리하고 UI 업데이트 등을 처리할 수 있음.