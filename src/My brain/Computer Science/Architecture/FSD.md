# FSD

<aside>
💡 Feature-Sliced Design.
기존 모듈식 아키텍처를 대체하는 기능 분할 설계 아키텍처.
모듈식 아키텍처의 암묵적인 모듈 간 연결의 복잡성을 해결하기 위해 나타남.
아키텍처 구성 요소를 쉽게 유지보수가능, 방법론이 비즈니스 지향적이라는 장점.
높은 진입 장벽, 팀 문화, 문제 해결을 즉시 해야한다는 문제점이 있다.

</aside>

### Layer

최상위 디렉토리.

앱 분할의 첫 단계.

레이어 수는 최대 7개로 제한.

app, processes, pages, widgets, features, entitles, shared 등으로 분리하며, app이 가장 상위 레이어.

**App**

애플리케이션의 진입점 역할.

**Processes(deprecated, optional)**

여러 페이지에 걸쳐있는 프로세스를 처리.

**Pages**

전체 페이지를 구성하기 위한 구성.

**Widgets**

의미있는 블록으로 결합.

**Features(optional)**

유저 시나리오와 같은 여러 로직 뭉치를 결합한 기능을 핸들링.

**Entities(optional)**

Business entities.

**Shared**

특정 비즈니스 로직에 종속되지 않은 재사용가능한 컴포넌트와 유틸리티가 포함.

### Slice

앱 분할의 두번째 단계.

레이어의 하위 디렉토리.

특정 비즈니스 엔티티에 관한 것.

### Segment

앱 분할의 세번째 단계.

슬라이스의 하위 디렉토리.

슬라이스 내의 코드를 목적에 맞게 분리하는 것.

ex/ api, ui, model, lib, config, consts