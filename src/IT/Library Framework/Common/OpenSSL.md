# OpenSSL

<aside>
💡

데이터 통신을 위한 TLS, SSL 프로토콜을 이용할 수 있는 OSS 라이브러리.

</aside>

**openssl enc -aes-256-cbc -salt -in HOGE -out FUGA**

- 암호화
- -aes-256-cbc
    - cbc 모드로 256 비트 크기의 aes 알고리즘을 사용함을 의미.
- -salt
    - 보다 복잡한 암호화
- -in
    - 인풋
- -out
    - 아웃풋

**openssl enc -d -aes-256-cbc -in FUGA -out HOGE**

- 복호화
- -d
    - 해독을 의미.

**openssl s_client -connect [www.example](www.example)**

- SSL/TLS 연결.
- s_client
    - OpenSSL의 클라이언트 도구 활성화.
- connect
    - 접속할 host 포트 지정.