# Flow

<aside>
💡 오직 한 값만 반환하는 정지 함수와 달리 여러 값을 순차적으로 내보낼 수 있는 타입.
데이터를 한 번에 하나씩 순차적으로 방출하며, 다수의 구독자에게 동시에 방출.
기본적으로 상태를 유지하지 않고, 새로운 구독자는 방출된 데이터의 현재 상태를 받지 않음.
단방향 데이터 흐름에 사용되며, 데이터를 처리하는 중간 단계에서 추가 작업을 요함.

</aside>

코루틴을 기반으로 빌드되며 여러 값을 제공할 수 있음.

비동기식으로 계산할 수 있는 데이터 스트림 개념.

내보낸 값은 동일한 유형.

## **Cold? Hot?**

**Cold Flow**

- 구독자가 구독하면 그 때부터 처음 데이터를 발행시키는 것. (구독자가 이전 데이터를 모르고 구독하면 그 때부터 자신에 의해 발행되는 것 처럼 인식한다는 의미)

**Hot Flow**

- 구독자를 신경쓰지 않고 데이터를 계속 발행시키는 것.

<aside>
⚠️ Flow는 자신을 호출한 코루틴의 context를 그대로 가지게 된다. (**Context preservation**)
호출한 곳과 다른 context를 사용할 때 `withContext`를 사용하면 `Flow invariant is violated`라는 에러가 발생하므로, `flowOn`을 사용해야 한다.

</aside>

## 파생 Flow

[StateFlow%20d586a489087b4534b323b21209e05ad2](StateFlow%20d586a489087b4534b323b21209e05ad2)

[SharedFlow%20d781747ff0c348cbbd2ab136b800507b](SharedFlow%20d781747ff0c348cbbd2ab136b800507b)

## 역할

**Producer**

스트림에 추가되는 데이터를 생산.

코루틴 덕분에 Flow에서 비동기적으로 데이터가 생산될 수 있음.

(선택사항) **Intermediary(중개자)**

스트림에서 내보내는 각각의 값이나 스트림 자체를 수정할 수 있음.

**Consumer**

스트림의 값을 사용.

## **함수**

**flow**

Flow를 생성하는 함수.

**emit**

flow 함수 내에서 값을 방출하는 함수.

**flowOf**

주어진 값들로 Flow를 생성.

**asFlow**

다른 컬렉션이나 시퀀스로부터 Flow를 생성.

**collect**

Flow에서 값을 소비해 데이터를 처리.

방출된 데이터를 모두 순차적으로 처리하고싶을 경우 이 것을 사용.

**launchIn**

Flow에서 값을 비동기적으로 소비해, 이를 지정한 CoroutineScope에서 처리.

**collectLatest**

Flow의 이전 값을 취소하거나 무시하고 가장 최근 값을 소비해 데이터를 처리.

최신 값만 관심이 있다면 이 것을 사용.

**map**

Flow의 각 값을 변환해 새로운 Flow를 생성.

**filter**

조건을 만족하는 값만으로 새로운 Flow를 생성.

**transform**

Flow의 값들을 변환하거나 추가적인 작업 수행.

**zip**

두 개 이상의 Flow를 결합해 하나의 Flow로 만듦.

각 Flow에서 방출되는 데이터가 한 Flow에서 한 짝으로서 방출됨.

쌍으로 묶일 데이터들이 전부 방출될 때까지 기다렸다가 방출.

**merge**

두 개 이상의 Flow를 결합해 하나의 Flow로 만듦.

각 Flow에서 방출되는 데이터가 한 Flow로서 방출됨.

**combine**

두 개 이상의 Flow를 결합해 하나의 Flow로 만듦.

각 Flow에서 방출되는 데이터가 한 Flow에서 한 짝으로서 방출됨.

각 Flow의 데이터가 방출될 때마다 기존 데이터에 새 데이터를 이용해 한 짝으로 만들어 방출.

**flatMapConcat**

각 값들을 Flow로 변환해 이를 순차적으로 연결해 새로운 Flow 생성.

**flatMapMerge**

각 값들을 Flow로 변환해 이를 병렬적으로 실행해 새로운 Flow 생성.

**flatMapLatest**

최신 Flow만 유지해, 새로운 Flow를 생성.

**distinctUntilChanged**

연속적으로 중복되는 값 제거.

**catch**

Flow에서 발생하는 예외를 처리.

**onCompletion**

Flow의 완료 시점에 작업 수행.

**onEach**

Flow의 각 요소에 대해 추가적인 작업 수행.

**buffer**

Flow의 값을 버퍼링해 처리 속도 조절.

**debounce**

일정 시간동안 값의 변화가 없을 때에만 값을 전달.

**flowOn**

Flow의 실행을 특정 스레드 혹은 스레드 풀로 이동.