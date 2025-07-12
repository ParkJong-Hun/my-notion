# Consumer

하나의 입력 값을 받아들이고, 해당 값을 처리하는 작업을 수행.

주로 Observable이 발행한 데이터를 구독하여 처리할 때 사용.

Observable.subscribe(Consumer())의 형태로 구독한다.

**accept**

받아들인 값을 처리하는 작업 수행.

accept의 오버라이드로 해당 함수를 사용할 때 LiveData.observe처럼 처리할 내용을 실행시킨다.