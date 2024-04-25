# remember

<aside>
💡 상태 관리를 위해 사용되는 기능.
컴포저블이 리컴포지션되어도 이전 State를 유지.
컴포저블의 생명주기를 이해.

</aside>

**rememberSavable**

- 앱의 구성 변경(화면 회전 등)에도 이전 State를 유지.

**rememberRetained**

- 보통의 remember에서는 백스택에 보관되지 않았던 State를 Navigation 변경 후 되돌아와도 이전 State를 유지.