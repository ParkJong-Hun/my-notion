# Hilt

<aside>
💡 안드로이드 애플리케이션에서 DI를 구현하기 위한 라이브러리.
Hilt는 Dagger를 기반으로 작동하며, Dagger Component의 단순화된 버전을 제공하여 의존성 주입을 쉽게 설정 가능.

</aside>

## Annotation

**HiltAndroidApp**

Application 클래스에 적용되는 어노테이션. Hilt의 기능을 활성화할 수 있음. 앱이 실행될 때 Hilt 컴포넌트와 의존성 주입을 구성하기 위해 필요한 코드를 생성.

**HiltViewModel**

Android Jetpack의 ViewModel과 함께 사용되는 어노테이션. 해당 ViewModel 인스턴스에 의존성 주입을 적용할 수 있음.

**AndroidEntryPoint**

Android Framework 주요 컴포넌트(Activity, Fragment, Service 등)에 적용할 수 있는 어노테이션. AndroidEntryPoint 어노테이션을 사용하면 해당 컴포넌트에 대해 Hilt 컴포넌트와 의존성 주입을 적용할 수 있음.

## Component

**ApplicationComponent**

애플리케이션의 전역 범위로 구성되는 컴포넌트. 주로 애플리케이션의 수명 주기 동안 유지되는 객체들을 관리.

Application Module과 @AndroidEntryPoint이 지정된 클래스를 사용해 구성됨.

**ActivityComponent**

액티비티의 수명 주기와 관련된 컴포넌트. 각 액티비티에 대해 별도의 컴포넌트 인스턴스가 생성되며, 액티비티 내에서 사용되는 의존성을 관리함. Activity Module과 @AndroidEntryPoint이 지정된 클래스를 사용해 구성됨.

**ActivityRetainedComponent**

액티비티의 생명주기와 관련이 없는 컴포넌트. 각 액티비티 인스턴스와 관계없이 한 번 생성되고, 그 이후로 유지. ActivityRetainedComponent는 주로 액티비티 간 공유되는 의존성을 관리하는 데 사용.

**FragmentComponent**

프래그먼트의 수명 주기와 관련된 컴포넌트. 각 프래그먼트에 대해 별도의 컴포넌트 인스턴스가 생성되며, 프래그먼트 내에서 사용되는 의존성을 관리. Fragment Module과 @AndroidEntryPoint이 지정된 클래스를 사용해 구성됨.

**ViewComponent**

사용자 정의 뷰의 수명 주기와 관련된 컴포넌트. 각 사용자 정의 뷰에 대해 별도의 컴포넌트 인스턴스가 생성되며, 해당 뷰 내에서 사용되는 의존성을 관리. View Module과 @AndroidEntryPoint이 지정된 클래스를 사용해 구성됨.

**ServiceComponent**

서비스의 수명 주기와 관련된 컴포넌트. 각 서비스에 대해 별도의 컴포넌트 인스터스가 생성되며, 서비스 내에서 사용되는 의존성을 관리. Service Module과 @AndroidEntryPoint이 지정된 클래스를 사용해 구성됨.

**SingleComponent**

애플리케이션 전역 범위의 의존성을 관리하기 위해 사용. 애플리케이션 전체에서 사용되는 의존성 객체를 주입할 수 있음.