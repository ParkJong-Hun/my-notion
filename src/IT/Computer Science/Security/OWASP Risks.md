# OWASP Risks

<aside>
💡

Open Web Application Security Project(OWASP)에서 주관해 발표하는 웹 앱, 소프트웨어 보안 관련 주요 위협 요소.

</aside>

### 2021년 기준

- **Broken Access Control**
    - 적절한 권한 검증이 이루어지지 않아 비인가 사용자가 민감한 데이터에 접근하거나 기능을 사용할 수 있는 취약점.
- **Cryptographic Failures**
    - 암호화 관련 문제로 인해 데이터가 노출되거나 변조될 위험이 있는 경우.
- **Injection**
    - SQL, NoSQL, OS 명령어, LDAP 등을 포함한 코드에 비정상적인 데이터를 주입하여 악성 코드를 실행하는 취약점.
- **Insecure Design**
    - 보안 요구사항을 충분히 고려하지 않고 설계된 시스템으로 인해 발생하는 문제.
- **Security Misconfiguration**
    - 설정 오류나 잘못된 디폴트 설정으로 인해 시스템이 공격에 노출되는 경우.
- **Vulnerable and Outdated Components**
    - 알려진 취약점을 가진 소프트웨어 또는 오래된 라이브러리를 사용하는 경우.
- **Identification and Authentication Failures**
    - 인증 및 세션 관리의 부적절한 구현으로 인해 사용자 계정이 탈취되는 문제.
- **Software and Data Integrity Failures**
    - 무결성 검증 부족으로 인해 악의적으로 수정된 소프트웨어 또는 데이터가 사용되는 경우.
- **Security Logging and Monitoring Failures**
    - 적절한 로깅 및 모니터링 부재로 인해 침입 탐지나 사고 대응이 어려운 경우.
- **Server-Side Request Forgery (SSRF)**
    - 서버가 신뢰하지 않는 요청을 처리하여 악의적인 명령을 실행하는 문제.