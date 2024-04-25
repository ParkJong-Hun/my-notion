# App Link

<aside>
💡 안드로이드의 딥 링크.

</aside>

[Deep Link](../../../Computer%20Science%20915624a6e5204e1c92de58c7899350dc/Deep%20Link%20fab576d7837c4034a7b0a93c1bebf099.md)

### **추가하기**

1. 앱의 특정 콘텐츠로 연결하는 딥 링크 만들기
    - AndroidManifest.xml의 액티비티 내에 URI 인텐트 필터를 만들고, 인텐트에서 얻을 데이터 구성을 작성해 사용자를 올바른 콘텐츠로 안내.
2. 딥 링크의 인증 추가하기
    - 앱 링크 인증을 요청하도록 함. 이후 Google Search Console로 앱 소유권을 확인하도록 **Digital Asset Link**라는 json 파일을 웹 사이트에 게시.
    

**Digital Asset Link**

앱의 정보를 기입한 json 문서.

해당 파일을 보유함으로서 앱 소유권을 확인하고 해당 앱이 아닌 다른 앱을 필터링할 수 있음.

namespace, 패키지 이름, sha256 서명 핑거 프린트가 필요.