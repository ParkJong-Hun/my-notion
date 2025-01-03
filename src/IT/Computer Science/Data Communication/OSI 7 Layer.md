# OSI 7 Layer

<aside>
💡 데이터 통신의 규격과 프로토콜을 통일하려고 했던 ISO가 규격 통일을 표준화 하여 선언한 데이터 통신 단게 구성도.
실제로 ISO 표준화는 실패했지만, 데이터 통신을 설명하는 데 유용.
데이터 통신을 7단계로 나눔.
네트워크에 의한 데이터 통신은 단계마다 복수의 프로토콜로 실현.
각 계층이 독립되어 있어서 어떤 계층의 프로토콜 변경은 다른 계층에 영향을 끼치지 않음.
하위 계층은 상위 계층을 위해서 일하고 상위 계층은 하위 계층에 관여하지 않음.

</aside>

### Application (Layer 7)

- 유저가 OSI 환경에 접근할 수 있도록 서비스 제공.
- 프로토콜을 이용한 서비스 사용을 의미.
- HTTP, FTP, DNS, Telnet, DHCP.

### Presentation (Layer 6)

- 알아볼 수 있게 표현.
- 응용 계층에서 받은 데이터를 세션 계층에 맞게 변환하거나, 그 반대를 행함.
- 코드 변화, 데이터 암호화, 데이터 압축, 구문 검색, 정보 형식 변환.
- ASCII, MPEG, JPEG, MIDI.

### Session (Layer 5)

- 송수신측 간 관련성 유지, 대화 제어 담당.
- 동기 제어, 데이터 교환 관리.
- SSH, TLS, NetBIOS.

### Trasport (Layer 4)

- 전송할 수 있게 연결을 설정, 해제.
- 데이터가 실제로 전송됨.
- 게이트웨이.
- TCP, UDP, ARP.

### Network (Layer 3)

- 네트워크 연결 관리. (경로 설정, 연결, 해제, 패킷 전송)
- 데이터 교환 중계.
- 라우터.
- IP, IPX.

### Data Link (Layer 2)

- 두 인접한 개방 시스템들 끼리 신뢰성 있고 효율적인 데이터 전송을 가능하게 함.
- 흐름 제어, 프레임 동기화, 오류 제어, 순서 제어.
- 링크 확립, 유지, 단절 제어.
- 랜카드, 브릿지, 스위치.
- MAC, Ethernet, FDDI.

### Physical (Layer 1)

- 전송에 필요한 두 장치간 실제 접속과 절단, 기계적, 전기적, 기능적, 절차적 특성 정의.
- 전기, 제어, 클럭 신호 전송.
- 리피터, 허브.
- Ethernet.RS-232C