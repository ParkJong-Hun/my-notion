# Project

<aside>
💡

Gradle 빌드의 기본 단위.
build.gradle 파일을 이용해 한 프로젝트의 구성과 동작 정의.

</aside>

- **extensions**
    - Project 안에서 등록된 Extension 모임.
    - **extension**
        - 플러그인이나 스크립트의 추가 설정.
- **configurations**
    - Project 안에서 등록된 의존성 모임.
    - **configuration**
        - 의존성 단위.
- **pluginManager**
    - PluginAware를 관리하는 객체.