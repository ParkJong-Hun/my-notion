# PendingIntent

<aside>
💡 안드로이드에서 특정 작업을 나중에 다른 앱이나 컴포넌트가 수행하도록 요청하는 기능.
실행될 Intent와 함께 전달됨.

</aside>

### 플래그

**FLAG_UPDATE_CURRENT**

- 기존 PendingIntent와 동일한 경우 새 Intent로 업데이트.

**FLAG_CANCEL_CURRENT**

- 기존 PendingIntent와 동일한 경우 취소하고 새로 생성.

**FLAG_NO_CREATE**

- PendingIntent가 존재하지 않을 경우, 새로 생성하지 않고 null 반환.

**FLAG_ONE_SHOT**

- 한 번만 사용할 수 있는 PendingIntent 생성.

### 특성

- 나중에 실행
- 권한 전달
- 여러 사용 사례