# GetX

<aside>
💡 간단하고 경량화된 방법으로 상태를 관리하는 Flutter 패키지.
DI 형식을 기반으로 간단하게 의존성 주입, 상태 관리.
간단하며 퍼포먼스가 좋고 적은 보일러플레이트 코드가 장점.
스트롱 타이핑을 제공하지 않으며, 테스트나 확장이 없다는 것이 단점.

</aside>

**.obs**

Stream으로 변환해 값을 관찰가능하게(observable) 만듦.

**Obx**

StreamBuilder를 대신해 Observer에 전달할 것.

**ever**

Rx 변수가 새 값을 내보낼 때 호출.

onEach 같은 것.

**everAll**

Rx 변수가 새 값을 내보낼 때마다 Rx 값 리스트를 사용해 호출.

**once**

Rx 변수가 새 값을 내보낼 때 딱 한 번 첫번째 변경에서만 호출.

**debounce**

사용자가 타이핑을 마칠 때만 호출.

**interval**

정한 타이머 시간이 지난 후 마지막 변경만 내보냄.

**Get**

경로 관리, 의존성 주입에 사용.

**to**

새 화면으로 이동.

**toNamed**

지정한 명칭으로 새 화면에 이동.

**back**

스낵바, 다이얼로그 등을 닫음.

**off**

이전 화면으로 돌아가기.

**offAll**

다음 화면으로 이동하고 이전 화면들을 모두 닫음.

**put**

Controller 주입.

**find**

주입한 Controller 반환.

**tr**

현재 지역값을 이용해 국제화된 문자열 반환.

**GetMaterialApp**

**changeTheme**

테마 변경.

**GetConnect**

통신에 사용.