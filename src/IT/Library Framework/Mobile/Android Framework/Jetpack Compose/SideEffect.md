# SideEffect

- **LaunchedEffect**
    - Composition이 시작하면 실행.
    - Composable scope에서 suspend fun 실행.
    - 키의 내용이 변경되면 기존 Coroutine을 취소하고 새 Coroutine에서 다시 실행.

- **SideEffect**
    - recomposition이 성공하면 실행.
    - Compose에서 관리하지 않는 객체에 Compose 상태를 이용할 때 사용.

- **DisabledEffect**
    - Composition이 종료된 후 실행.
    - 키의 내용이 변경되면 기존 Effect를 취소하고 새 Effect를 다시 실행.

- **EffectName**
    - Effect에서 실행 중인 Effect를 취소하는 데 인수로 사용되는 키.

- **produceState**
    - State를 생성하고 관리하기 위한 함수.

- **derivedStateOf**
    - State를 가공해 이를 관찰할 수 있도록 하는 함수.
    - 주어진 코드 블록 내에서 사용된 State에 대한 종속성을 자동으로 감지해 해당 State가 변경될 때 코드 블록을 다시 실행.

- **rememberCoroutineScope**
    - CompositionScope를 이용해 Composable 외부에서 Coroutine을 실행.

- **snapshotFlow**
    - Flow value를 캡처해 이를 Composable 함수 내에서 사용할 수 있도록 하는 함수.