# Channel

<aside>
💡 BlockingQueue와 다르게 suspend하게 데이터를 전송해 두 개의 Coroutine 사이를 연결하는 파이프.

</aside>

**send**

- 송신측에서 데이터 전달.

**receive**

- 수신측에서 데이터 취득.

**Channel Buffer Type**

- Rendezvous(Unbueffered)
    - 기본값
    - 버퍼가 없기 때문에 수신측과 송신측이 모두 가능한 상태로 모일 때까지 suspend.
- Conflate
    - 수신측이 송신측을 따라잡지 못했다면, 송신측이 새로운 값을 버퍼의 마지막에 덮어씌움.
    - 수신측이 다음 값을 받을 차례가 되면, 송신측이 보낸 마지막 값을 받음.
    - 수신측은 채널 버퍼에 값이 올 때까지 suspend.
- Buffered
    - 고정된 크기인 Array 형식의 버퍼 생성.
    - 송신측은 버퍼가 꽉 차면 새로운 값을 보내는 것을 중단.
    - 수신측은 버퍼가 빌 때까지 계속 꺼내서 수행.
- Unlimited(Linked List)
    - 제한 없는 크기인 Linked List 형식의 버퍼를 가짐.
    - OOM을 일으킬 수 있음.
    - 송신측은 suspend하지 않지만, 수신측은 버퍼가 비면 suspend.