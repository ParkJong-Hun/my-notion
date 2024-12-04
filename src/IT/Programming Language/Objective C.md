# Objective C

<aside>
💡 주로 Apple의 앱을 만드는데 사용되는 객체 지향형 언어.
LLVM 컴파일러를 사용해 LLVM 바이트 코드로 변환된다.

</aside>

[h%20101f37315c44808da519c6953a1b9b76](h%20101f37315c44808da519c6953a1b9b76)

[m%20101f37315c4480b58377d88d80096e60](m%20101f37315c4480b58377d88d80096e60)

### 기본 문법

**@autoreleasepool**

메모리 할당, 해제 자동 처리 영역.

**@interface Vehicle : NSObject {**

**// member variable**

**int wheels;**

**}**

**-(void)setWheels:(int)w;**

**@end**

정의

**@implementation Vehicle**

**-(void)setWheels:(int)w {**

**wheels = w;**

**}**

**@end**

구현

**@property**

setter, getter를 자동으로 만드는 변수.

**@synthesize**

프로퍼티에 대해 자동으로 getter, setter를 구현시켜줌.(최근에는 작성할 필요 없음)