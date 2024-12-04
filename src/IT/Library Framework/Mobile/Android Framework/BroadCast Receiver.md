# BroadCast Receiver

<aside>
💡 안드로이드 OS로부터 발생하는 각종 이벤트와 정보를 받아와 핸들링함.

</aside>

### 정적 리시버

AndroidManifest에 <receiver>를 작성해 <intent-filter>로 원하는 action을 관리할 수 있음.

### 동적 리시버

`Context.registerReceiver()`로 코드 내에서 작성해 관리 가능. 단, `Context.unRegisterReceiver()`로 리시버를 사용하지 않을 때는 등록을 해제해줘야함.

### 액션 종류

- ACTION_BOOT_COMPLETED: 부팅이 끝났을 때
- ACTION_CAMERA_BUTTON: 카메라 버튼이 눌렸을 때
- ACTION_DATE_CHANGED
- ACTION_TIME_CHANGED: 날짜, 시간이 수동으로 변했을 때
- ACTION_SCREEN_OFF
- ACTION_SCREEN_ON: 화면이 켜질 때, 꺼질 때
- ACTION_AIRPLANE_MODE_CHANGED: 비행기 모드가 바꼈을 때
- ACTION_PACKAGE_ADDED
- ACTION_PACKAGE_CHANGED
- ACTION_PACKAGE_DATA_CLEARED
- ACTION_PACKAGE_INSTALL
- ACTION_PACKAGE_REMOVED
- ACTION_PACKAGE_REPLACED
- ACTION_PACKAGE_RESTARTED: 앱 설치/제거
- ACTION_POWER_CONNECTED
- ACTION_POWER_DISCONNECTED: 충전
- ACTION_REBOOT
- ACTION_SHUTDOWN: 부팅, 종료
- ACTION_TIME_TICK: 매 android.provider.Telephony.SMS_RECEIVED: SMS 수신 (RECEIVE_SMS 권한 필요)