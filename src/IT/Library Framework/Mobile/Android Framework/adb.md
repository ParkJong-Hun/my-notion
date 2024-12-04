# adb

**adb shell**

Android Shell 진입

**adb shell activity activities**

OS에서 현재 실행 중인 모든 액티비티 목록 표시

**adb shell dumpsys -l**

열람 가능한 시스템 서비스 목록 표시

**adb shell screenrecord**

디바이스 화면 녹화

**adb install**

PC의 앱을 디바이스에 설치

**adb shell screencap**

디바이스 화면 캡처

**adb pull**

캡처한 것을 PC로 보내기

**adb shell pm grant**

앱에 권한 부여

**adb shell pm revoke**

앱의 권한 취소

### Android Shell

**dumpsys activity | grep -i run**

OS에서 실행중인 Activity 필터링해 표시

**dumpsys activity**

OS에서 실행중인 시스템 상태, Activity, Task 표시