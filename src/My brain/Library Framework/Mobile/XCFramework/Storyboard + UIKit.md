# Storyboard + UIKit

## View Controller

### 라이프 사이클

- **init()**

인스턴스 초기화.

- **loadView()**

인터페이스 초기화, 뷰 계층 구조 설정.

- **viewDidLoad()**

추가적인 초기화, 데이터 로드.

- **viewWillAppear(_:)**

화면이 나타나기 전 필요한 작업 수행.

- **viewDidAppear(_:)**

애니메이션을 시작하거나 외부 데이터를 업데이트하는 등의 작업 수행.

- **viewWillDisappear(_:)**

화면이 사라지기 전 필요한 작업 수행.

- **viewDidDisappear(_:)**

리소스를 해제하거나 다른 화면으로 전환하는 등의 작업 수행.

- **didReceiveMemoryWarning()**

메모리를 해제하고 불필요한 데이터 정리.

- **deinit()**

메모리 해제.