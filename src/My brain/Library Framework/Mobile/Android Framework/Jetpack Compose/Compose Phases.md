# Compose Phases

1. **Composition**

함수를 실행해 UI를 나타내는 트리 구조 출력.

1. **Layout**
    1. Measurement
    2. Subcomposition
    3. Placement

위의 단계에서 만들어진 UI 트리를 사용해 각 레이아웃 노드의 하위 컴포넌트 측정, 크기 측정, 위치 결정, 배치해, 레이아웃 노드에 width, height, x, y를 지정.

1. **Drawing**

UI 트리를 위에서부터 아래로 다시 순회해 실제로 UI를 그림.