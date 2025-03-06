# NfcAdapter

<aside>
💡

NFC 기능을 제어하는 클래스.

</aside>

**Foreground Dispatch System**

- 포어그라운드에서 앱이 태그를 읽을 수 있도록 함.
- 태그 감지 우선 순위가 매우 높음.

**NFC Reader Mode**

- 백그라운드에서도 앱이 태그를 읽을 수 있도록 함.
- 설정이 다소 복잡함.
- 태그 감지 우선 순위가 낮음.
- 배터리 사용량이 높음.

**Ndef**

- NFC 태그에 데이터를 읽거나 쓸 때 사용.