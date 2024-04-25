# CompositionLocal

<aside>
💡 Composable 트리 내에서 데이터를 공유하고 전파하는 데 사용.
주로 테마, 로케일, 특정 데이터 등을 요소 사이에서 전달하고 공유하는 데 유용.

</aside>

**CompositionLocalProvider**

- 데이터를 제공하는 곳. 상위 Composable에서 데이터를 설정하고 하위 Composable에서 사용 가능하게 만듦.

**CompositionLocal**

- 데이터를 사용하려는 하위 Composable에서 CompositionLocal을 통해 데이터를 읽음.

**CompositionLocalAmbient**

- 데이터를 가져올 때 사용되는 Composable에서 CompositionLocalAmbient를 사용해 데이터를 읽음.

**staticCompositionLocalOf()**

- 자주 변경되지 않는 상태를 저장할 때 이용.
- 상태가 변경되면 할당된 변경에 영향이 있는 모든 컴포저블을 리컴포지션.

**compositonLocalOf()**

- 자주 변경되는 상태를 저장할 때 이용.
- 현재 상태에 접근하는 컴포저블만 리컴포지션.