# RxJava

<aside>
💡 ReactiveX 라이브러리의 Java 구현체 중 하나.

</aside>

[Java%208daf8a644670414a8300b057c041cba8](Java%208daf8a644670414a8300b057c041cba8)

### Single

정확히 하나의 데이터만 방출하거나 에러를 발생시킴.

성공 시 onSuccess 호출.

**just**

여러 개의 값을 받아들여 순서대로 발행.

**error**

즉시 에러를 발행.

### Maybe

0개 또는 1개의 데이터를 방출하고 완료되거나, 없거나, 에러를 발생시킴.

### **Completable**

데이터을 방출하지 않음.

단순히 Complete 혹은 Error만 발생.

**fromRunnable**

Java의Runnable을RxJava의Completable로 변환.

동기 처리를 스트림화하는데에 사용.

## 연속적 데이터 스트림

### **Observable**

데이터를 생성하고, 해당 데이터를 구독자들에게 전달하는 역할.

**subscribe**

옵저버를 구독하고, 구독자에게 데이터를 전달해 한 함수를 실행.

**doOnNext**

Observable이 값을 발행할 때마다 누군가 구독자가 있으면 여러 실행문을 실행.

**andThen**

Completable의 완료만을 기다리고, 데이터는 전달하지 않음.

하나의 작업이 완료된 후 다음 작업을 실행할 때 사용.

**onErrorResumeNext**

에러가 발생했을 때 대체 스트림으로 전환. 없으면 원본 스트림 진행.

**map**

데이터를 변환하는 함수를 적용해 새로운 Observable을 생성.

**filter**

주어진 조건에 따라 데이터를 필터링해 새로운 Observable을 생성.

**subscribeOn**

데이터를 생성, 처리할 스레드 또는 스케줄러를 지정.

**observeOn**

데이터를 후처리할 스레드 또는 스케줄러를 지정.

**dispose**

옵저버의 구독을 취소하고, Observable의 동작을 중단.

**merge**

Observable에서 발행된 아이템들을 순서대로 합침.

**zip**

Observable에서 발행된 아이템들을 조합해 새 아이템을 발행.

**create**

새 데이터를 발행.

### **Flowable**

Observable과 유사하지만 backpressure 지원.

생산자가 소비자보다 빠르게 아이템을 방출할 때 발생하는 문제 해결.

대량의 데이터나 파일처리, 네트워크 요청에 적합.