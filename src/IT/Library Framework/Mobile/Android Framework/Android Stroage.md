# Android Stroage

### **Internal System and App Data**

**/data/**

내부 데이터를 저장하는 디렉토리.

시스템 파일, 앱 데이터, OS 설정과 상태 정보 포함.

루트 권한을 가진 사용자만 접근 가능. 

**/data/data/<패키지 이름>/**

앱의 내부 데이터를 저장하는 디렉토리.

다른 앱에서는 접근할 수 없음.

SharedPreferenes, Databases, Files, Cache 등이 포함.

보통 앱이 삭제될 때 함께 삭제됨(단, 백업 가능).

**/data/system/**

시스템 관련 설정 파일을 저장하는 디렉토리.

시스템 데이터, 로그인 상태, 시계, 타이머 등이 포함.

**/data/user/<사용자 번호>**

멀티 사용자 환경을 지원하면서 도입된 디렉토리.

각 사용자마다 독립된 데이터 저장 공간을 제공.

숫자로 몇번째 사용자인지 나타냄.

### **External Storage and Media File**

**/storage/**

외부 저장소 디렉토리.

내장된 외부 저장소나 SD 카드 같은 저장 장치의 파일 시스템을 나타냄.

**/storage/emulated/<사용자 번호>**

내장 저장소.

사용자가 직접 파일에 접근할 수 있는 파일, 사진, 비디오, 오디오 등을 포함.

**/stroage/sdcard/**

SD 카드나 외장 메모리.

역할은 위와 같음.

**/storage/emulated/<사용자 번호>/Android/data/<패키지 이름>/**

앱 전용 외부 저장소라고도 불림.

외부 저장소에 있지만, 앱 전용으로만 접근 가능한 영역.

다른 앱에서 접근할 수 없으며, 앱이 삭제될 때 함께 삭제됨.