# Google Play

# 테스트 종류

### Internal Testing

내부 테스트.

앱을 개발한 팀 내에서 제한된 그룹에게 앱을 배포해 테스트하는 과정.

알파 테스트보다 더 제한된 범위.

### Closed Testing

비공개 테스트(알파 테스트).

앱 출시 전에 더 많은 테스터와 출시 전 버전을 테스트해 더 많은 타겟 피드백을 수집.

공개 릴리스 테스트로 확장할 수 있음.

### Open Testing

공개 테스트(베타 테스트).

대규모 그룹을 대상으로 테스트를 실행해 Google Play 앱의 테스트 버전을 공개 가능.

누구나 테스트 프로그램에 참여해 익명 피드백을 제출 가능.

# 결제

**Google Play Services**를 이용해 앱에서 직접 Google Play Backend를 이용하거나, 앱의 백엔드에서 Google Play Backend나 GCP를 이용.

**Google Play**

- 사용자가 앱, 기타 디지털 제품을 다운로드할 수 있는 온라인 상점.

**Google Play Console**

- Google Play에 앱을 게시할 수 있는 인터페이스를 제공하는 플랫폼.
- Google Play를 통해 판매하는 제품이나 콘텐츠를 비롯해 앱에 관한 세부정보 표시.

**Google Cloud Console**

- Google Play Developer API 같은 백엔드 API를 관리하는 플랫폼.

**Google Play Billing Library**

- Google Play 결제 시스템을 앱에 통합하는 데 사용하는 API.

**Google Play Developer API**

- 게시 앱 관리 작업을 프로그래매틱하게 처리하는 데 사용하는 REST API.

**Realtime Developer Notification**

- Cloud Pub/Sub를 사용해 Google Play에서 관리하는 정기 결제의 상태 변경을 실시간으로 모니터링해주는 메커니즘.

**Security Backend Server**

- Google Play 결제 시스템을 앱에 통합하는 일환으로 보안 백엔드 서버
- 이를 사용해 구매 확인, 정기 결제 관련 기능, 실시간 개발자 알림 처리 같은 결제 관련 작업 구현하는 것이 좋음.

**Google Play Store App**

- Google Play와 관련된 모든 작업을 관리하는 앱.

**License Tester**

- 서명되지 않아 Google Play에 업로드되지 않은 앱의 경우 Google Play 결제 라이브러리가 일반적으로 차단되지만 이 검사를 건너뛰게 할 수 있음.
- 패키지 이름은 Google Play 용으로 구성된 앱 이름과 일치해야 하며 Google 계정은 Google Play Console 계정의 라이선스 테스터여야 함.
- 테스터에 실제 구매 비용을 청구하지 않는 테스트 결제 수단에 액세스 가능.

**Play Billing Lab**

- 개발자가 Google Play 결제 시스템과 통합을 테스트할 수 있는 앱.

# 통합

**Play Integrity API**

- Google Play의 무결성, 서명 서비스를 위한 API.
- 정품 앱 바이너리, Play 설치, Android 기기를 증명.

**Play App Signing**

- Google의 보안 인프라에서 앱 서명 키를 관리하고 보호해 보안을 강화하는 업그레이드 옵션 제공.
- 앱 서명 키를 사용해 AAB에서 최적화된 배포 APK 생성.

**Google Play Instant**

- 앱을 설치 없이 체험적으로 실행 가능하도록 하는 것