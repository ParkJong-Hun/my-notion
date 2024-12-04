# LifecycleOwner

<aside>
💡 클래스에 `Lifecycle` 이 있음을 나타내는 단일 메서드 인터페이스.
클래스에서 구현해야 하는 `getLifecycle()` 메서드가 하나 있음.
대신 전체 애플리케이션 프로세스의 수명 주기를 관리하려는 경우 `ProcessLifecycleOwner`를 참고.

</aside>

[Lifecycle%208af5874f03e84f6386306cf064c67b20](Lifecycle%208af5874f03e84f6386306cf064c67b20)

`Fragment`와 `AppCompatActivity`와 같은 개별 클래스에서 `Lifecycle`의 소유권을 추출해, 함께 작동하는 컴포넨트를 작성할 수 있게 함. 모든 맞춤 애플리케이션 클래스는 `LifecycleOwner` 인터페이스를 구현할 수 있음.

관찰자가 관찰을 위해 등록할 수 있는 수명 주기를 소유자가 제공할 수 있으므로, `DefaultLifecycleObserver`를 구현하는 컴포넌트는 `LifecycleOwner`를 구현하는 컴포넌트와 원활하게 작동.

라이브러리에서 Android 수명 주기와 작동하는 데 필요한 클래스를 제공한다면 수명 주기 인식 컴포넌트를 사용하는 것이 좋음. 클라이언트 측에서 수명 주기를 관리하지 않아도 라이브러리 클라이언트는 이러한 컴포넌트를 쉽게 통합할 수 있음.

## **맞춤 LifecycleOwner 구현**

지원 라이브러리 26.1.0 이상의 프래그먼트와 액티비티는 이미 `LifecyclerOwner`가 구현되어 있음.

`LifecycleOwner`를 만들려는 맞춤 클래스가 있다면 `LifecycleRegistry` 클래스를 사용할 수 있지만, 이벤트를 전달해야 함.