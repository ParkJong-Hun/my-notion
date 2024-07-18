# Compose Structure

Sources → Compiler —(Code Transformation)→ Runtime ←(In-memory Representation)→ UI (Rendering Layout trees)

**Compose Compiler**

- Jetpack Compose 내의 중추적인 역할을 하는 핵심 구성 요소.
- 순수한 Kotlin으로 작성된 Kotlin 컴파일러 플러그인.
- KAPT나 KSP 같은 annotation processor와 달리 FIR을 통해 개발자가 작성한 코드를 정적 분석, 변형.
- 2024년 4월부터 JetBrains에서 관리.
- Kotlin 2.0부터는 Kotlin과 버전 통합.
- Composable 함수에 사용된 매개변수에 stable 혹은 unstable 타입 부여.

**Compose Runtime**

- Compose 모델, 상태 관리의 초석 역할.
- Gap Buffer라는 데이터 구조에서 파생된 Slot Table을 사용해 상태 저장 등 메모리 운용.
- Composable 함수에 대한 실질적 생명주기 관리, UI에 대한 정보를 담고 있는 메모리 표현.

**Compose UI**

- 개발자가 레이아웃을 생성할 수 있도록 여러 UI 라이브러리를 통해 컴포넌트 제공.
- 각 컴포넌트는 LayoutNode를 생성하고, 결과적으로 Compose Layout Tree 구성을 용이하게 함.
- Compose Runtime에 의해 소비되는 것이 기본 원칙.