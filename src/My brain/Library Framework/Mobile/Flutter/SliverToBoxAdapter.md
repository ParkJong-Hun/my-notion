# SliverToBoxAdapter

**Sliver**

한 박스 위젯을 갖음.

CustomScrollView를 사용해 결합해 커스텀 스크롤 효과를 만들 수 있는 특수 목적 위젯.

**SilverToBoxAdapter**

일반적인 박스 기반 위젯 중 하나로 다시 연결되는 브리지를 만드는 기본 Sliver.

여러 박스 위젯을 CustomScrollView에서 표시하고 싶을 때 SliverToBoxAdapter를 사용하는 것보다 SliverList, SliverFixedExtentList, SliverPrototypeExtentList, SliverGrid를 사용하는 것이 스크률 뷰의 뷰 포트를 통해 실제로 표시되는 하위 항목만 인스턴스화 하기 때문에 더 효율적.