# Pass

### 핵심 개념

**Pass**

- 유저에게 발급되어 Google Wallet에 저장되는 Passes Object 인스턴스.
- 탑승권, 이벤트 티켓, 신분증 등 여러 유형 지원.
- **Generic Pass**
    - 특별히 지원되지 않는 패스를 만드는 데 사용.

**Pass Issuer**

- Pass를 생성해 유저에게 발급해 Google Wallet에 저장하는 엔티티.
- Pass 소유, 생성, 발급, 업데이트 가능.
- 개발자, 회사, 조직일 수 있음.

**Passes Class**

- Pass를 생성하는 공유 템플릿.
- 스타일, 모양, Smart Tap, 등록, 로그인 같은 기능 지원.

**Passes Object**

- 유저에게 발급되어 Google Wallet에 저장되는 개별 Pass 정의.
- 종종 사용자별 정보 포함.
    - ex/ 잔액, 만료일

**Private Passes**

- 민감한 유저 데이터를 안전하게 유지하기 위해 추가 보호가 필요한 Pass.
- Generic Private Pass를 사용.

**Smart Tap**

- 모바일 기기와 NFC 간 데이터를 전달하기 위한 Google 독점 근거리 통신 프로토콜.