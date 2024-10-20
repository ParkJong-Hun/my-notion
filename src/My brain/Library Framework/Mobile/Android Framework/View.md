# View

<aside>
💡

Android UI의 기본 구성 요소.
자신의 UI를 화면에 렌더링할 책임을 갖고 있으며, 내부적으로 Canvas를 사용.

</aside>

### 렌더링 과정

1. Constructor
    1. 뷰 초기 설정 세팅.
    2. AttributeSet으로 지원.
2. onAttatchedToWindow
    1. 부모 뷰가 addView(childView)를 호출한 뒤 자식 뷰가 윈도우에 Attach.
    2. 뷰 ID로 뷰에 액세스 가능.
3. onMeasure
    1. 뷰의 크기 측정.
4. onLayout
    1. 뷰의 위치와 onMeasure로 결정한 크기를 토대로 최종 크기 결정.
5. onDraw
    1. 뷰를 실제로 렌더링.
    2. Canvas, Paint를 사용.

**invalidate** 

- onDraw부터 다시 그림.

**requestLayout**

- onMeasure부터 다시 그림.