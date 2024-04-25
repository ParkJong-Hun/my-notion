# Launch Mode

<aside>
💡 AndroidManifest에서 액티비티와 태스크를 관리하는 모드.

</aside>

**Standard**

기본값. 해당 모드 Activity는 한 Task 안에서 여러 번 인스턴스화 될 수 있으며, 여러 다른 Task 안에 속할 수 있음.

**SingleTop**

Task의 Top에서 연달아 존재할 수 없음.

**SingleTask**

Task에서 해당 모드의 Activity가 라우팅되면, 자신보다 이전의 Activity는 전부 Task에서 사라짐.

단, 뒤로가기 버튼을 눌렀을 때 이전 Activity로 돌아감.

**SingleInstance**

SingleInstance로 지정된 Activity가 있는 Task에는 다른 Activity가 추가되지 않음.