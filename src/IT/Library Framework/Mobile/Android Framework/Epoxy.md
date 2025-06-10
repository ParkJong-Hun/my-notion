# Epoxy

<aside>
💡

Airbnb의 Android RecyclerView를 쉽게 만들기 위한 라이브러리.

선언적 UI, 자동 diffing, 타입 안전성, EpoxyModel 기반, 데이터 바인딩 지원이 특징.

</aside>

### vs Groupie

- `EpoxyModelWithHolder` 구현과 `EpoxyModelClass`라는 Annotation class로 사용할 뷰를 직접 대입.
- `EpoxyAttribute`로 모델 상태를 사용.
- diffing 최적화 더 발달.
- 러닝 커브가 높음.
- `EpoxyController`는 Adapter 역할.

### migrate to RecyclerView.Adapter

1. EpoxyModelClass를 대신해 데이터 클래스 정의.
2. EpoxyHolder를 대신해 ViewHolder 생성.
3. EpoxyController를 대신해 Adapter 구현.
4. 수동으로 DiffUtil 구현.