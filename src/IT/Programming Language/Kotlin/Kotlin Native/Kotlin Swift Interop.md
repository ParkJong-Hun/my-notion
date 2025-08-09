# Kotlin/Swift Interop

# Class and Interfaces

|  | Kotlin | Swift | Good? | 비고 |
| --- | --- | --- | --- | --- |
| Abstract class | `abstract class` | `class Hoge : AbstractClass`  | ⛔️ | 상속할 수 있지만, `abstract fun`을 오버라이드 해야한다고 말하지 않음.
이 기능을 사용하려고 하면 `NSGenericException`이 발생해 앱이 크래쉬. |
| Annotation class | `annotation class`  | - | ⛔️ | 사용 불가능. |
| Companion object | `Class.companion object.value`  | `Class.` `value` | ✅ |  |
| Data class | `data class` 
`copy`
`equals` 
`toString`  | `class` 
`doCopy` 
`isEquals`
`description`  | ✅ | destructing은 지원 안함. |
| Enum class | `enum class`  | `class`  | ⚠️ | Switch에서 편리하게 사용은 불가능. |
| Fun interface | `fun interface`  | - | ⛔️ | 사용 불가능.
다만, 비슷한 방법으로 사용할 수 있지만 상당한 오버헤드 발생. |
| Inline class | `inline class` |  | ⛔️ | 사용 불가능. |
| Inner class | `AClass(a).BClass()`  | `AClass.BClass(a)` | ✅ | 부모가 아닌, 자식에게 구문 전달. |
| Interface | `interface`  | `protocol`  | ✅ |  |
| Object | `object A`  | `A.shared`  | ✅ |  |
| Open class | `open class`  | `class`  | ⚠️ | `protected`, `open fun` 재정의 가능.
final method를 재정의하면 `NSException`으로 앱이 중단됨 . |
| Sealed class | `sealed class`  | `class` | ⚠️ | 자식 클래스가 일반 클래스로 변환되어 switch에서 default 필요. |
| Sealed interface | `sealed interface`  | `protocol` | ⚠️ | 별도의 관계없는 프로토콜로 분리됨. |

# Coroutines

|  | Kotlin | Swift | Good? | 비고 |
| --- | --- | --- | --- | --- |
| Flow | `collect<Int>` | `collect<Any?>`  | ⚠️ | Flow: Flow는 제네릭타입을 잃고Any?로서 수집해야 함. 
Task.cancel: 작업이 취소되어도 흐름은 계속 요소 생성. |
| Suspend function | `suspend fun a`  | `a(completionHandler: { ReturnType, ErrorType in … }`  or (Swift 5.5 ~)
`async await` | ⚠️ | Task.cancel: 작업이 취소되어도 suspend 함수는 평소처럼 반환. |

# Extensions

|  | Kotlin | Swift | Good? | 비고 |
| --- | --- | --- | --- | --- |
| Extension function over platform class | `fun String.hoge()`  | `FileKt.hoge(a: String)`  | ⚠️ | 기본 유형이나 특정 플랫폼 특화 클래스의 확장 함수는 파일명Kt.함수로서 호출 가능. |
| Extension function over usual class | `fun A.hoge()`  | `A().hoge()` | ✅ |  |
| Extension properties for companion object of platform class | `String.Companion.CONST_VAL`  | - | ⛔️ | 기본 유형이나 특정 플랫폼 특화 클래스의 Companion object는Objective-C 헤더 파일에는 존재하지만 접근 불가능. |
| Extension properties for companion object of usual class | `A.Companion.CONST_VAL` | `A.companion.CONST_VAL` | ✅ |  |
| Extension properties over platform class | `val String.a` | `FileKt.a(s: String)` | ⚠️ | 기본 유형이나 특정 플랫폼 특화 클래스의 확장 프로퍼티는 파일명Kt.함수로서 호출 가능. |
| Extension properties over usual class | `val A.a`  | `A().a`  | ✅ |  |

# Function and Properties

|  | Kotlin | Swift | Good? | 비고 |
| --- | --- | --- | --- | --- |
| Constructor | `Class(a = a)`  | `Class(a: a)`  | ✅ |  |
| Function expecting lambda argument | `a(x = { ... })`  | `FileKt.a(x: { ... })`  | ✅ |  |
| Function returning function type | `returnLambda()()`  | `FileKt.returnLambda()()` | ✅ |  |
| Member function | `class.a()`  | `class.a()`  | ✅ |  |
| Mutable member properties | `class.a = “”`  | `class.a = “”`  | ✅ |  |
| Read-only member properties | `class.a`  | `class.a`  | ✅ |  |
| Top-level mutable var properties | `a = 1`  | `FileKt.a = 1`  | ⚠️ | 탑 레벨은 Swift에서 import로 표현할 수 없으므로 파일명을 지정해야함. |
| Top-level val properties | `a`  | `FileKt.a` | ⚠️ | 탑 레벨은 Swift에서 import로 표현할 수 없으므로 파일명을 지정해야함. |

# Generics

|  | Kotlin | Swift | Good? | 비고 |
| --- | --- | --- | --- | --- |
| Bounded generic | `class A<T: B>()` | -  | ⚠️ | 제네릭 정보가 전달되지 않으므로 직접 스위프트 내에서 클래스 상속 관계를 정의해야 함. |
| Contravariant generic |  |  |  |  |
| Covariant generic |  |  |  |  |
| Generic class |  |  |  |  |
| Generic function |  |  |  |  |
| Generic interface |  |  |  |  |
| Reified function |  |  |  |  |
| Star projection |  |  |  |  |

# More about Functions

|  | Kotlin | Swift | Good? | 비고 |
| --- | --- | --- | --- | --- |
| Constructor with default argument | `class A(p : Boolean = true)`  | `class A(p: Boolean)`  | ⛔️ | 디폴트 값이 사라짐. |
| Function expecting lambda with receiver | `a {
  …
}` | `a(x: { hoge in
 …
}`
   | ⚠️ | 확장 함수가 람다로 변경되고, 첫번째 파라미터가 리시버. |
| Function with default argument | `fun a(p : Boolean = true)`  | `func a(p : Boolean)` | ⛔️ | 디폴트 값이 사라짐. |
| Function with overload | `fun a(param: Long)
fun a(param: Float)` | `func a(param: Long)
fun a(param_: Float)` | ⚠️ | 스위프트에서는 같은 파라미터가 같은 이름으로 오버로딩이 불가능해서, 이름에 _이 추가됨. |
| Function with receiver |  |  |  |  |
| Function with value class parameter | `value class A(val x: Int)` | `x: Int32`  | ⚠️ | 내부의 타입으로 변형. |
| Function with vararg parameter | `fun a(vararg  item: String)` | `func a(item: KotlinArray<NSString>)`  | ⚠️ | KotlinArray로 변형됨 |
| Inline function | `a()`  | `a()` | ✅ |  |

# Types

|  | Kotlin | Swift | Good? | 비고 |
| --- | --- | --- | --- | --- |
| Basic type | `String` `Boolean` 
`Byte` 
`UByte`
`Short`
`UShort`
`Int`
`UInt`
`Long`
`ULong`
`Float`
`Double` 
`Char` | `String` 
`Boolean` 
`Int8`
`UInt8`
`Int16`
`UInt16`
`Int32`
`UInt32`
`Int64`
`UInt64` 
`Float`
`Double` 
`unichar` | ✅ |  |
| Collection with basic type | `List<Byte>`
`List<UByte>`
`List<Short>`
`List<UShort>`
`List<Int>`
`List<UInt>`
`List<Long>`
`List<ULong>`
`List<Float>`
`List<Double>`
`List<Boolean>`
`List<String>`
`List<Char>` | `[KotlinByte]`
`[KotlinUByte]`
`[KotlinShort]`
`[KotlinUShort]`
`[KotlinInt]`
`[KotlinUInt]`
`[KotlinLong]`
`[KotlinULong]`
`[KotlinFloat]`
`[KotlinDouble]`
`[KotlinBoolean]`
`[String]`
`[Any]` | ⚠️ | String과 Char는 원시형 타입과 비슷 |
| Collection with custom type | `List<A>`
`Set<A>`
`Map<String, A>` | `[A]` 
`Set<A>`
`[String: A]` | ✅ |  |
| Mutable, immutable collection | `MutableList<Int>` | `var a: [KotlinInt]` | ✅ |  |
| Optional basic type | `?` | `?` | ✅ |  |
| Unit and Nothing | `Unit` 
`Nothing`  | `KotlinUnit` 
`KotlinNothing`  | ⚠️ | KotlinUnit은 생성 가능하지만, KotlinNothing은 앱을 크래쉬 시킴. |