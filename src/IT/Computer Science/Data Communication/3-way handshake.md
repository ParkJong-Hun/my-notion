# 3-way handshake

<aside>
💡

TCP/IP에서 각 호스트가 연결될 때까지의 과정.

</aside>

1. Client의 Squence Number가 1000, Server Sequence Number 4000.
2. Server가 Listen.
3. Client가 SYN_SENT
    1. **Client에서 Server로 SYN(번호) 전송.**
4. Server가 SYN_RCVD
    1. **Server에서 Client로 SYN(4000) + ACK(1001) 전송.**
5. Client에서 ESTABLISHED
    1. **Client에서 Server에 ACK(4001) 전송.**
6. Server에서 ESTABLISHED