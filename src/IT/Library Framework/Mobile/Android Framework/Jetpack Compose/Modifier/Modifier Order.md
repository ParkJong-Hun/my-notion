# Modifier Order

기본적으로 다음과 같은 순서로 적용.

1. 기본 Modifier
    - 기본 레이아웃과 스타일을 설정.
2. 레이아웃 Modifier
    - padding, offset, align 같은 것으로 위치와 크기 조정.
3. 그림자, 클립 Modifier
    - 그림자나 클립 관련 영역 정의.
    

ex)

패딩을 무시하고 모든 Modifier 부분 클릭 가능.

`Modifier`

`.clickable()`

`.padding()`

패딩을 적용한 후의 Modifier 부분 클릭 가능.

`Modifier`

`.padding()`

`.clickable()`