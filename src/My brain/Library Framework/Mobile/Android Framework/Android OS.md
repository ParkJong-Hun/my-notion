# Android OS

### 구조

**System Application**

일반 사용자들이 접하는 앱.

기본 제공 앱과 구글 플레이 스토어에서 다운로드 받은 앱들.

- 홈, 연락처, 전화, 브라우저, 카메라 등

**Java API Framework(Application Framework)**

자바로 작성된 앱 관련 프레임워크.

개발자가 코드를 작성하는 앱 계층과 더 아래 수준의 계층이 소통할 수 있도록 하는 다리 역할.

액티비티 생명주기, 언어 설정 등 프레임워크 수준 기능 담당.

- Content Providers: 애플리케이션 간 데이터 공유 관리.
- View System: UI에서 쓰이는 안드로이드 뷰들 관리(AAC, Compose)
- Managers
    - Activity: 앱 안 액티비티들을 관리.
    - Location: GPS, 기지국 신호로 위치 정보 관리.
    - Package: 앱 설치, 업데이트, 삭제 관리.
    - Notification: 알림 관리.
    - Resource: 앱에서 사용하는 리소스 관리.(이미지, 문자열 등)
    - Telephony: 음성 통화 관리.
    - Window: 디스플레이 상의 창을 배치하고 관리.

**Native C/C++ Libraries**

일반적으로 사용하는데 조금 더 저수준에서 동작하는 라이브러리.

- Webkit: 웹 콘텐츠 렌더링 담당.
- OpenMAX AL: 멀티미디어 데이터 처리 담당.
- Libc: C 프로그래밍 표준 라이브러리. 메모리, 파일, 문자열 조작 담당.
- OpenGL ES: 2D/3D 그래픽 담당.
- Surface Manager: 화면 버퍼를 관리하고 화면을 그리는 담당.
- SQLite: 경량화 로컬 DB
- FreeType: 폰트 렌더링
- SGL: 2D 그래픽
- SSL: 인터넷에서 데이터 통신 보안에 사용되는 프로토콜.

**Android Runtime**

실행 중인 앱 코드를 OS가 이해하도록 컴파일.

자바의 저작권, 성능 등의 이유로 JVM을 사용하지 않고 Dalvik Virtual Machine을 사용하다가 롤리팝부 Android Runtime을 사용. 그러다가 누가부터는 다시 달빅의 JIT와 ART를 병용. 

- Android Runtime(ART)
    - 자바 코드를 OS가 이해하는 바이너리 코드 컴파일.

**Hardware Abstraction Layer(HAL)**

하드웨어 공급업체에서 구현해야 하는 표준 인터페이스를 정의.

안드로이드에서 하위 드라이버 구현을 고려하지 않아도 되게 해주는 시스템.

- Audio IF
- Bluetooth IF
- Camera IF
- Sensors IF

**Linux Kernel**

OS가 리눅스 기반으로 설계됐기 때문에 커널도 리눅스로 만들어짐.

하드웨어와 OS를 연결하는 다리 역할을 함.

시스템의 핵심적인 부분을 관리.

하드웨어 추상화, 메모리 관리, 보안 설정, 전원 관리, 다른 하드웨어 장치의 드라이버 관리,

네트워크 서비스 관리

- Drivers
    - Audio
    - Binder(IPC)
    - Display
    - Keypad
    - Bluetooth
    - Camera
    - Shared Memory
    - USB
    - WiFi
- Power Management