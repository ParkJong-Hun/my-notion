# @Bean

<aside>
💡

Spring 컨테이너가 관리하는 자바 객체.
Spring 앱에서 필요한 객체를 생성하고, 관리하는 단위.

기본적으로 자동적으로 생성되며, DI로 여러 곳에 사용 가능한 싱글톤.

</aside>

### @Scope

**singleton**

- 기본 값.

**prototype**

- 요청할 때마다 새 객체 생성.

**request**

- HTTP 요청마다 새 객체 생성.

**session**

- 세션마다 새 객체 생성.