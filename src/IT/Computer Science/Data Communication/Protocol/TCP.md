# TCP

<aside>
💡

Transmission Control Protocol.

전송 제어 프로토콜.

두 기기 간 네트워크를 통한 신뢰성 있는 데이터 전송을 보장하는 통신 프로토콜.
데이터 분할, 전송 중 손실 복구, 순서 보장, 흐름 제어 및 혼잡 제어를 제공.
인터넷의 기초인 TCP/IP 스위트의 일부.

</aside>

### TCP Header

- 20 바이트.
- 16 비트 소스 포트 + 16 비트 목적지 포트.
- 32비트 시퀀스 넘버.
- 32비트 Acknowledgment 넘버.
- Offset + Reversed + 16비트 TCP Flags.