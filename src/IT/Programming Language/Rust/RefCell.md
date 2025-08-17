# RefCell

<aside>
💡

Cell보다 더 유연하고 강력한 내부 가변성 타입.

값을 직접 복사(get, set)하지 않고, 런타임에서 borrow를 체크(Ref/RefMut)하는 Cell.

Copy 타입 외의 모든 타입에 대해서도 작동 (참조를 통해 값 변경)

대여 규칙을 런타임에 검사하며 참조(&mut, &를 동적으로 관리)

스레드 안전하지 않음.

</aside>

### 대여 규칙

get할 때 Ref 혹은 RefMut으로만 접근 가능.

set할 때 borrow_mut()을 사용해서 참조 생성 후 값을 변경해야 함.

- borrow (&)은 여러 번 사용 가능.
    - 불변 참조이므로
- borrow_mut(&mut)은 한 번만 가능.
    - 가변 참조이므로
    - 여러번 가변 참조를 생성하면 런타임 에러(패닉)이 발생.
    - borrow가 있는데, borrow_mut을 생성하면 런타임 패닉 발생.

주로 String이나 Vec<T> 같은 복사 불가능한 데이터를 내부적으로 변경할 때 사용.

```rust
struct Counter {
	// 기본 타입이므로 (copy 가능하므로) Cell로 가능.
	count: Cell<i32>
}
```

```rust
struct DynamicList {
	// 기본 타입이 아니므로 Cell로는 불가능.
	count: RefCell<Vec<String>>
}
```