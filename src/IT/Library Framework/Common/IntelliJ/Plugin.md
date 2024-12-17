# Plugin

<aside>
💡

IntelliJ에서 사용할 수 있는 IDE의 플러그인.

</aside>

**plugin.xml**

- META-INF에 정의하는 플러그인 구성 파일.
- 플러그인 내의 액션이나 의존성을 구성.

**Action**

- 사용자가 플러그인의 기능을 호출하는 가장 일반적인 방법.
- AnAction, actionPerformed()로 작성.

**Extension**

- 플러그인이 메뉴나 도구 모음에 동작을 추가하는 것만큼 간단하지 않은 IntelliJ 플랫폼 기능을 확장하는 방법.