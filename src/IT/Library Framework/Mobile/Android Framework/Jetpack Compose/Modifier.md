# Modifier

<aside>
💡 수정자. Jepack Compose에서 컴포넌트의 크기, 배경, 클릭, 스크롤 등을 지정하는 것.

</aside>

**Modifier(아무런 확장함수도 작성하지 않았을 때)**

기본적으로 해당 컴포넌트가 표시될 정도의 크기만을 제공한다. 즉, 기존의 `wrap_content`와 유사하게 작동한다.

[Modifier%20Order%205296c719cc2b44b089ebe5e75398a27b](Modifier%20Order%205296c719cc2b44b089ebe5e75398a27b)

**weight**

같은 그룹(같은 Row나 Column, Box 레이어 하)에 존재하는 컴포넌트들의 기본 크기를 계산하고 나머지 값을 가중치에 따라 계산하여 나눠 가짐. 

예를 들어 기본 크기를 계산하고 나머지 값을 가중치가 2f인 컴포넌트 하나와, 1f인 컴포넌트 두 개가 있으면,  (해당 컴포넌트의 가중치 / 다른 값을 가진 가중치들을 합친 값)가 해당 컴포넌트의 크기에 추가되므로 (2/3).

가중치를 가진 컴포넌트가 해당 그룹에서 단 하나라면, 해당 그룹의 컴포넌트들의 기본 크기를 계산하고 나머지 값을 위의 계산식에 따라 1/1 즉, 전부 갖게 됨. 

<aside>
⚠️ 같은 그룹 내에 wrapContentSize를 사용한 뷰가 있다면, 그 계산도 끝마친 이후에 weight가 사용할 나머지 값이 반영됨.

</aside>

**pointerInput { detectTapGestures }**

onPress, onLongPress 등 눌렀을 때 어떠한 처리를 하고 싶을 때 사용.

tryAwaitRelease() 이후의 처리는 탭을 한 후 종료되기전에 실행.

**layout**

컴포넌트가 측정되고 배치되는 방식을 수정.

- **mesaureable**
    - measure를 통해 측정할 수 있는 컴포넌트.
- **constraints**
    - 컴포저블에 수신된 제약 조건.
    - measure에 사용됨.
- placeable
    - 화면에 래핑된 컴포넌트를 place로 배치 가능.