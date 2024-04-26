# RxJava

<aside>
💡 ReactiveX 라이브러리의 Java 구현체 중 하나.

</aside>

[Java%208daf8a644670414a8300b057c041cba8](Java%208daf8a644670414a8300b057c041cba8)

## **Observable**

데이터를 생성하고, 해당 데이터를 구독자들에게 전달하는 역할.

**subscribe**

옵저버를 구독하고, 구독자에게 데이터를 전달해 한 함수를 실행.

**doOnNext**

Observable이 값을 발행할 때마다 누군가 구독자가 있으면 여러 실행문을 실행.

**map**

데이터를 변환하는 함수를 적용해 새로운 Observable을 생성.

**filter**

주어진 조건에 따라 데이터를 필터링해 새로운 Observable을 생성.

**observeOn**

데이터를 처리할 스레드 또는 스케줄러를 지정.

**dispose**

옵저버의 구독을 취소하고, Observable의 동작을 중단.

**merge**

Observable에서 발행된 아이템들을 순서대로 합침.

**zip**

Observable에서 발행된 아이템들을 조합해 새 아이템을 발행.

**just**

여러 개의 값을 받아들여 순서대로 발행.

**create**

새 데이터를 발행.

## Consumer

하나의 입력 값을 받아들이고, 해당 값을 처리하는 작업을 수행.

주로 Observable이 발행한 데이터를 구독하여 처리할 때 사용.

Observable.subscribe(Consumer())의 형태로 구독한다.

**accept**

받아들인 값을 처리하는 작업 수행.

accept의 오버라이드로 해당 함수를 사용할 때 LiveData.observe처럼 처리할 내용을 실행시킨다.