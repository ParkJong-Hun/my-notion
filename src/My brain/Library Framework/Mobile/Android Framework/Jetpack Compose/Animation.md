# Animation

### **High Level Animation API**

**AnimatedVisibility**

- 표시하거나 숨길 때 나타남, 사라짐 애니메이션

**AnimatedContent**

- 다른 컴포저블 간에 애니메이션을 적용

**animateContentSize()**

- 자동 크기 변경 애니메이션

### Low Level Animation API

**animateHogeAsState**

- 단일 값을 애니메이션 처리하는 가장 간단한 방법

**Transition**

- 한 번에 여러 값에 애니메이션 적용

**InfiniteTransition**

- 지속적으로 애니메이션 적용

**Animatable**

- 다른 애니메이션 API으로는 실현불가능한 제어가 필요한 경우에 구현하기 위해 사용

**AnimatedImageVector**

- 드로어블 파일을 로드하는 등으로 만드는 움직이는 이미지 벡터

**AnimationSpec**

- 애니메이션 사양(물리학 설정 등)

**MotionCompose**

- MotionLayout의 Compose 버전

**MotionScene**

- MotionLayout 애니메이션에 대한 다양한 제약 조건 세트, 전환, 키 프레임 정의

**ConstraintSet**

- MotionLayout 초기, 중간, 최종 레이아웃 상태 정의 제약 조건 집합

**Shared Element Transition**

- 같은 컴포넌트를 두 화면에서 자연스럽게 전환해 사용하도록 하는 방법

**Navigation Animation**

- EnterTransition, ExitTransition, PopEnterTransition, PopExitTransition으로 정의.
- animationSpec으로 별도 관리 가능.