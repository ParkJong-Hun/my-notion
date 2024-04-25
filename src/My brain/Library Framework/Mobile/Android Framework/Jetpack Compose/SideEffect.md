# SideEffect

- **derivedStateOf**
    - State를 가공해 이를 관찰할 수 있도록 하는 함수.
    - 주어진 코드 블록 내에서 사용된 State에 대한 종속성을 자동으로 감지해 해당 State가 변경될 때 코드 블록을 다시 실행.
    
- **snapshotFlow**
    - Flow value를 캡처해 이를 Composable 함수 내에서 사용할 수 있도록 하는 함수.

- **produceState**
    - State를 생성하고 관리하기 위한 함수.