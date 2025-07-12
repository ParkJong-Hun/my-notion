# Runnable

<aside>
💡

Java의 핵심 인터페이스 중 하나.
실행 가능한 작업을 정의하는 함수형 인터페이스.
멀티스레딩과 비동기 처리의 기초가 되는 단위.

</aside>

```java
@FunctionalInterface
public interface Runnable {
  public abstract void run();
}
```

<aside>
⚠️

Java 8 이후 람다가 도입되고 잘 사용되지 않음.

</aside>