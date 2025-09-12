# HTTP

<aside>
💡 Hyper Text Transfer Protocol.
애플리케이션 레이어 프로토콜.
웹 브라우저와 웹 서버간의 통신을 위해 설계되었지만 다른 목적으로도 사용할 수 있음.
기본 포트는 `80`

</aside>

**Header**

- 클라이언트와 서버가 요청 또는 응답으로 부가적인 정보를 전송할 수 있도록 해줌.
- 대소문자를 구분하지 않는 이름과 `:` 다음에 오는 값(줄 바꿈 없이)으로 이루어져있음.
- 값 앞에 붙은 빈 문자열은 무시.
- 커스텀 등록 헤더는 `X-`를 앞에 붙여 추가될 수 있지만, 이 관례는 [rfc6648](rfc6648)에서 비표준 필드가 표준이 되었을때 불편함을 유발하는 이유로 2012년 6월에 폐기.
- 반드시 메시지의 최종 수신자에게 전송되어야 함. 요청에 대해서는 서버, 응답에 대해서는 클라이언트. 중간 프록시는 반드시 종단 간 헤더를 수정되지 않은 상태로 재전송해야하며 캐시는 이를 반드시 저장해야 함.
- [General_header](General_header)
    - HTTP 헤더를 말하는 구식 용어.
    - 요청과 응답 모두에 적용되지만 바디에서 최종적으로 전송되는 데이터와는 관련이 없는 헤더.
    - **Date**
        - 메시지가 생성된 날짜와 시간
    - **Connection**
        - 네트워크 연결 관리
            - keep-alive, close
    - **Transfer-Encoding**
        - 메시지 본문의 인코딩 방식
            - chunked 등
- [Request_header](Request_header)
    - fetch될 리소스나 클라이언트 자체에 대한 자세한 정보를 포함하는 헤더.
    - **Content-Type**
        - 요청 본문의 미디어 타입 지정
            - application/json, text/html, mutlipart/form-data 등이 있음.
    - **Accept**
        - 클라이언트가 받아들일 수 있는 미디어 타입을 서버에 알림
            - application/json, text/html, application/xhtml+xml 등이 있음.
    - **Authorization**
        - 인증 정보 전달 (토큰, 기본 인증 등)
            - Bearer Hoge, Basic Hoge 등
    - **User-Agent**
        - 클라이언트 앱 정보 (브라우저, 운영체제 등)
    - **Host**
        - 요청하는 서버의 호스트명과 포트
    - **Cookie**
        - 클라이언트에 저장된 쿠키를 서버에 전송
    - **Referer**
        - 현재 요청을 보낸 페이지의 URL
- [Response_header](Response_header)
    - 위치 또는 서버 자체에 대한 정보(이름, 버전 등)와 같이 응답에 대한 부가적인 정보를 갖는 헤더.
    - **Content-Type**
    - **Content-Length**
        - 응답 본문의 바이트 크기
    - **Set-Cookie**
        - 클라이언트에 쿠키 설정 지시
    - **Cache-Control**
        - 캐싱 정책 지정
    - **Location**
        - 리다이렉션 시 새로운 URL 지정
    - **Server**
        - 서버 소프트웨어 정보
- [Entity_header](Entity_header)
    - 컨텐츠 길이나 MIME 타입과 같이 엔티티 바디에 대한 자세한 정보를 포함하는 헤더.

**HTTP/0.9 (One-Line Protocol)**

- 요청이 단일 라인. 리소스에 대한 메소드는 GET만 존재.
- 응답이 단순.
- 헤더가 없으며, HTML 파일만 전송 가능.
- 상태 혹은 에러 코드가 없음.

**HTTP/1.0**

- 헤더가 있음.
- 요청, 응답 모두 Meta data 전송을 허용.
- 유연하고 확장 가능하도록 함.
- 버전 정보와 요청 메소드가 함께 전송됨.
- 응답 시작 부분에 상태 코드가 추가되어, 요청 성공/실패 파악 가능.
- Content-Type 도입으로 HTML 이외의 문서 전송 가능.
- 커넥션 하나당 요청 하나와 응답 하나만 처리 가능해, 매우 비효율적이고 서버 부하가 문제.

**HTTP/1.1 (Standard Protocol)**

- 모호함을 명확하게 하고 많은 개선 사항을 도입.
- 영구적인 커넥션.
    - 지정한 타임아웃 동안 커넥션 재사용이 가능하며, 기존 연결에 대한 Handshaking을 생략.
- 파이프라이닝 추가.
    - 앞 요청의 응답을 기다리지 않고 순차적으로 여러 요청을 연속적으로 보내고 그 순서에 맞게 응답.
- HOL로 뒤 요청이 블로킹됨, 하나의 커넥션이 하나의 요청을 처리해, latency가 증가, 헤더에 많은 Meta data가 있는 등 헤더가 거대해짐, 우선순위 없음이 문제.

**HTTP/2**

- HTTP 1.X 버전의 성능 향상에 초점을 맞춘 확장.
- **Google SPDY (Proprietary Solution)**
    - 웹 컨텐츠 전송을 위해 구글이 개발한 비표준 네트워크 프로토콜.
    - 전송 지연을 줄이기 위해 압축, 다중화, 우선순위 설정 등을 지원.
    - SSL/TLS는 암호화되지 않은 연결을 지원하지 않음.
- Binary Framing 계층을 추가해 보내는 메시지를 Frame이라는 단위로 분할해 추가적으로 바이너리로 인코딩.
- 멀티플렉싱이 가능.
    - 요청이 여러개 일 때, 첫 요청이 끝나기 전에 더 빨리 응답이 가능한 요청에 응답하는 등 비동기적으로 응답이 가능.
- 각 요청을 Stream으로 구분해 병렬적으로 처리가 가능하지만, TCL 고유의 HOL Blocking이 아직 존재. 하나의 Stream에서 유실이 발생하거나 문제가 생기면 다른 Stream에도 문제가 해결될 때까지 지연되는 현상이 발생하는 문제가 있음.

**HTTP/3 (Quick UDP Internet Connections)**

- QUIC 기반으로 나온 새로운 HTTP 메이저 버전.
- TCP가 아닌 UDP를 이용해 HOL Blocking을 해결한 네트워크 통신 프로토콜.

[HTTP%20Status%20167f37315c4480c2aca7d48cfa1b8572](HTTP%20Status%20167f37315c4480c2aca7d48cfa1b8572)

[HTTPS%200af386f99ca241e9bde0e0545a127e56](HTTPS%200af386f99ca241e9bde0e0545a127e56)