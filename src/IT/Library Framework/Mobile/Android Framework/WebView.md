# WebView

**WebView**

- handler
    - obtainMessage
        - 메시지를 생성.
- requestFocusNodeHref
    - 링크에 포커스가 갔을 때 해당 링크의 URL을 받아옴.
- canGoBack
    - 뒤로가기 가능한지 확인.
- goBack
    - 뒤로가기.
- goForward
    - 앞으로 가기.
- canGoForward
    - 앞으로가기 가능한지 확인.
- reload
    - 새로고침.
- scrollTo
    - 스크롤.
- pageUp
    - 페이지 상단 이동.
- pageDown
    - 페이지 하단 이동.

**Handler**

- UI 스레드에서 관리되며, 다른 스레드에서 UI 스레드를 업데이트 할 때 사용.
- 메시지 큐를 갖음.
- sendMessage
    - Message를 Handler에 직접 보냄.

**Message**

- 메시지 큐에서 사용되는 단위로 실제 큐에 전달되는 데이터.
- sendToTarget
    - 이미 생성된 Message를 다른 Handler로 전달.

**WebViewClient**

- 웹 페이지가 로드될 때 발생하는 여러 이벤트를 처리하는 데에 사용.
    - 링크 클릭, 페이지 로딩, 에러 처리 등.
- 웹 페이지의 네비게이션 관리.
- shouldOverrideUrlLoading
    - 웹 페이지 내 링크가 클릭되었을 때
- shouldInterceptRequest
    - 웹 페이지에서 요청이 발생했을 때
- onPageStarted
    - 페이지 로딩이 시작될 때
- onPageFinished
    - 페이지 로딩이 완료될 때
- onReceivedError
    - 페이지 로딩 중 에러 발생 시
- onReceivedSslError
    - 페이지에서 SSL 통신 에러 발생 시
- doUpdateVisitedHistory
    - 페이지에서 히스토리가 갱신되었는지
        - 페이지가 SPA로 구성되어서 shouldOverrideUrlLoading이 호출되지 않을 때 렌더링이 리로드되었는지 확인

WebChromeClient

- 웹 페이지의 UI와 자바스크립트 관련 이벤트를 처리하는 데에 사용.
    - 제목, 자바스크립트 경고, 비디오 처리 등.
- onCreateWindow
    - 윈도우가 생성되었을 때
- onProgressChanged
    - 페이지 로딩 진행 상태
- onReceivedTitle
    - 페이지 제목 처리
- onReceivedIcon
    - 페이지 아이콘 처리
- onConsoleMessage
    - 자바스크립트에서 발생한 콘솔 메시지 처리