# Build Lifecycle

<aside>
💡

의존성 기반 프로그래밍을 위한 Gradle이 작업이 종속성 순서대로 실행되고 각 작업이 한 번만 실행되도록 보장하기 위해 종속성을 위한 방향성 비순환 그래프를 형성하고 빌드하는 과정.

</aside>

### Build phases

1. Initialization
    1. 빌드에 참여할 프로젝트를 결정하고 각 프로젝트에 대한 Project 인스턴스 생성.
2. Configuration
    1. 프로젝트 개체가 구성. 빌드에 포함된 모든 프로젝트의 빌드 스크립트 실행.
    2. 생성, gradle 명령과 현재 디렉토리에 전달된 작업 이름 인수에 의해 구성된 작업의 하위 집합을 결정.
3. Execution
    1. 선택한 각 작업을 실행.

### Initialization

다중 프로젝트의 경우, settings.gradle에서 다중 프로젝트 빌드를 트리거해 Gradle이 빌드 구성.

setting.gradle이 없는 프로젝트 내에서 Gradle을 실행하면 상위 디렉토리에서 settings.gradle을 찾음.

찾을 수 없다면 단일 프로젝트 빌드로 실행.