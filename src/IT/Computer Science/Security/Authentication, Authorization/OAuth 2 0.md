# OAuth 2.0

<aside>
💡 OAuth 프로토콜 버전 중 하나로, 이전 버전인 1.0의 몇 가지 단점을 보완해 간단하고 유연한 인증 프로세스 제공.
주로 웹, 모바일 애플리케이션에 API의 인증, 권한을 부여하기 위해 사용

</aside>

**순서**

1. 서비스 접근 및 이용 시도(사용자 → 서비스)
2. Client ID, Redirect URI(서비스 → 사용자)
3. 로그인 페이지 요청 Client ID, Redirect URI(사용자 → 인증 서버)
4. 로그인 페이지 제공(인증 서버 → 사용자)
5. ID / PW 입력(사용자 → 인증 서버)
6. Authorization code 발급(인증 서버 → 사용자)
7. Redirect URI로 Authorization code 전달(사용자 → 서비스)
8. Authorization code로 Access Token 요청(서비스 → 인증 서버)
9. Access Token 발급(인증 서버 → 서비스)
10. 인증 완료 및 로그인 성공(서비스 → 사용자)
11. 서비스 요청(사용자 → 서비스)
12. Access Token으로 API 호출(서비스 → 리소스 서버)
13. Authorization code 검증 및 서비스 제공(리소스 서버 → 서비스)
    1. scope를 활용해 필요한 권한만 제공
14. 서비스 제공(서비스 → 사용자)