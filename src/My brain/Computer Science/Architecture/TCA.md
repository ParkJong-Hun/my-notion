# TCA

<aside>
💡

The Composable Architecture.
pointfree에서 만들어낸 아키텍처.
iOS에서 주로 사용.

</aside>

**View(Presenter)**

상태를 표시.

**Action**

Enum으로 정의된 여러 UI 이벤트.

**Reducer**

Action이 오면 다음 상태와 부수 효과를 관할하는 주체.

**Effect**

외부에서 발생하는 부수 효과.

**Environment**

API나 라이브러리 등 의존성.

**Store**

Reducer, State, Environment 등을 작동시키는 단위.