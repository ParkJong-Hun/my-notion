# settings.gradle

<aside>
💡

Gradle 멀티 프로젝트 빌드에서 프로젝트 구성을 정의.

계층 구조를 정의하기 위해 사용.

</aside>

- **include**
    - 프로젝트 구성을 지정해 특정 하위 프로젝트를 포함.
- **repositories**
    - 의존성을 해결하고 필요한 외부 라이브러리, 플러그인을 다운로드하는 방법 지정.
- **includeBuild**
    - 다른 Gradle 프로젝트를 포함하고 그 프로젝트의 설정, 빌드를 재사용.
    - Gradle 7.0부터 도입.
- **repositoriesMode**
    - 의존성 해결 시 어떤 저장소를 검색할지 제어하는데 사용.
    - Gradle 7.0부터 도입.
- **pluginManagement**
    - 플러그인 버전을 중앙화 관리.
    - Gradle 7.0부터 도입.
- **dependencyResolutionManagement**
    - 프로젝트의 의존성 해결 규칙 관리.
    - resolutionStrategy
        - 의존성 해결 규칙 설정.
    - Gradle 7.0부터 도입.
- **versionCatalogs**
    - 의존성 버전 관리를 위한 설정 옵션.
    - Gradle 7.2부터 도입.