# CoroutineContext

<aside>
💡 코루틴의 실행되는 환경 정의.

</aside>

**Key**

- CoroutineContext Element의 키.

**Element**

- CorotuineContext의 싱글톤 요소.
- 고유의 키를 가져서 이 것이 여러 코루틴을 구별 가능하도록 함.

**plus**

- 좌측 CoroutineContext의 Element와 우측 CoroutineContext의 Element를 포함하는 CoroutineContext를 반환.
- 좌측 CoroutineContext와 동일한 Key를 가진 우측 CoroutineContext의 Element는 삭제됨.