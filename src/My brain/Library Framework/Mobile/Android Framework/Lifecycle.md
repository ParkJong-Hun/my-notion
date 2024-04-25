# Lifecycle

<aside>
💡 액티비티나 프래그먼트와 같은 구성요소와 수명주기 상태관련 정보를 포함해 다른 객체가 이 상태를 관찰할 수 있게 하는 클래스

</aside>

[수명 주기 인식 구성요소로 수명 주기 처리  |  Android 개발자  |  Android Developers](https://developer.android.com/topic/libraries/architecture/lifecycle?hl=ko#lco)

**이벤트**

프레임워크, `Lifecycle` 클래스에서 전달되는 수명 주기 이벤트.

액티비티나 프래그먼트의 콜백 이벤트에 매핑됨.

ex) onCreate, onStart, onResume, onPause, onStop, onDestroy

**상태**

`Lifecycle` 객체가 추적한 구성요소의 현재 상태

ex) INITIALIZED(initial state), CREATED, STARTED, RESUMED, DESTROYED(dead state)

<aside>
💡 클래스는 DefaultLifecycleObserver를 구현해 onCreate, onStart 등에 상응하는 메소드를 재정의해 컴포넨트의 수명 주기 상태를 모니터링할 수 있다.
LifecycleOwner는 이 것을 옵저버로 추가하는 메소드를 갖고 있다.

</aside>

```
class MyObserver : DefaultLifecycleObserver {
    override fun onResume(owner: LifecycleOwner) {
        connect()
    }

    override fun onPause(owner: LifecycleOwner) {
        disconnect()
    }
}

myLifecycleOwner.getLifecycle().addObserver(MyObserver())
```