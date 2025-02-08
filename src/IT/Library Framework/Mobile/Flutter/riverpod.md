# riverpod

<aside>
💡 리액티브 캐싱, 데이터 바인딩 프레임워크.
Provider를 더 개선한 것.

</aside>

### 기본

**Notifier**

StateNotifier를 개선한 클래스로 riverpod 2.0부터 제공.

build()로 객체 초기화.

제네릭 작성 필요 없음.

어떠한 상태를 보유.

함수를 갖고 있어서 어떠한 상태를 갱신하도록 가능.

**@riverpod**

동시에 여러 위젯에서 액세스될 수 있는 공유 스테이트 클래스를 의미.

클래스의 어노테이션으로 사용해 정의.

**hogeProvider**

@riverpod으로 정의한 클래스를 감시하기 위한 것으로 NotifierProvider.

RiverpodGenerator으로 생성됨.

캐싱, 오류/로딩 처리, 리스너 추가, 데이터가 변경 되면 자동적으로 리빌딩 지원.

**(WidgetRef).watch({provider이름})**

공유 스테이트의 데이터 변화를 감시해서 최근 값을 상수로서 사용.

감시하는 스테이트의 값이 변하면 리빌드 되는 방식.

**ProviderScope**

앱에서 Riverpod이 작동하도록 하는 스코프.

보통 가장 최상단 위젯을 래핑.

**StateNotifier**

ValueNotifier와 비슷하지만, 더 정교하고 강력한 상태 관리 기능 제공.

불변 상태를 권장.

### 부수 작용

**hogeNotifier**

@riverpod과 함께 이름을 작성하여 정의 가능.

provider들의 stateful 위젯.

build() 함수를 오버라이드 해야 함.

**(WidgetRef).read({provider.notifier})**

Notifier를 불러들임.

**(WidgetRef).listen({provider}, (previousValue, newValue) { logic… })**

Provider의 값이 변경되면 logic을 실행.