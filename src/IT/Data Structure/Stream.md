# Stream

<aside>
💡 스트림.

</aside>

**핫 스트림**

데이터가 구독자가 없어도 끊임없이 흐르며 실시간으로 처리됨.

**콜드 스트림**

데이터가 구독자가 구독했을 때부터 방출됨.

**Backpressure management**

데이터 스트림에서 생산자와 소비자 간 처리 속도 차이를 조절하는 메커니즘.

- Buffering
- Dropping
- Throttling
- Request-based