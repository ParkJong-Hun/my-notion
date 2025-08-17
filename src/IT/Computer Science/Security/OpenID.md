# OpenID

<aside>
💡

분산 인증 시스템으로, 사용자가 하나의 디지털 신원으로 여러 웹 사이트로 로그인 할 수 있게 해주는 개방향 표준.

기본 개념은 SSO, 분산 인증, 사용자 중심.

</aside>

### OpenID 1.0/2.0 (레거시)

- 초기 OpenID는 URL 기반 식별자 사용
- 동작 방식
    - 사용자가 웹사이트에 OpenID URL 입력
    - 웹사이트가 OpenID 제공자와 통신해 인증 요청
    - 사용자가 OpenID 제공자에게 인증
    - 인증 결과가 원래 웹사이트로 전달

### OpenID Connect (현재 표준)

- OAuth 2.0 기반 현대적인 인증 계층
- 동작 방식
    - 사용자가 클라이언트 앱에 로그인 요청
    - 클라이언트가 OpenID Provider에 인증 요청 (인가 코드)
    - OpenID Provider가 사용자에게 로그인 페이지 제공
    - 사용자가 OpenID Provider에게 인증 정보 입력
    - OpenID Provider가 클라이언트에게 인가 코드 전달
    - 클라이언트가 OpenID Provider의 인가 코드를 Access Token과 ID Token으로 교환
    - 클라이언트가 ID Token으로 사용자 정보 추출s