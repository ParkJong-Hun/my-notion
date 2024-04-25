# Shere sheet

<aside>
💡 공유, 퍼가기.
Intent를 생성해 action 속성을 `Intent.ACTION_SEND`로 설정하고, `putExtra(Intent.EXTRA_보낼유형, 내용)`으로 보낼 데이터를 Intent에 추가.
이후 type 속성에서 보낼 내용의 파일 확장자(MIME)를 지정.
 `Intent.createChooser`를 사용하지 않아도 심플하게 사용 가능하지만, 사용하면 제목을 지정한 Intent로 반환 가능.

</aside>

### 기본 형태

```kotlin
val sendIntent = Intent().apply {
  action = Intent.ACTION_SEND
  putExtra(Intent.EXTRA_TEXT, "text")
  type = "text/plain"
}
val shareIntent = Intent.createChooser(sendIntent, "텍스트 공유")
startActivity(shareIntent)
```

### MIME 종류

- text/plain
- text/rtf
- text/html
- text/json
- image/jpg
- image/png
- image/gif
- video/mp4
- video/3gp
- application/pdf

### 바이너리 콘텐츠 공유

```kotlin
val sendIntent = Intent().apply {
  action = Intent.ACTION_SEND
  putExtra(Intent.EXTRA_STREAM, stream)
  type = "image/jpeg"
}
val shareIntent = Intent.createChooser(sendIntent, "이미지 공유")
startActivity(shareIntent)
```

### 여러 콘텐츠 공유

```kotlin
val sendIntent = Intent().apply {
  action = Intent.ACTION_SEND
  putParcelableArrayListExtra(Intent.EXTRA_STREAM, arrayListOf(uri1, uri2))
  type = "image/jpeg"
}
val shareIntent = Intent.createChooser(sendIntent, "여러 이미지 공유")
startActivity(shareIntent)
```

### 미리보기 추가

```kotlin
val sendIntent = Intent().apply {
  action = Intent.ACTION_SEND
  putExtra(Intent.EXTRA_TEXT, "text")
  putExtra(Intent.EXTRA_TITLE, "preview title")
  data = contentUri
  flags = Intent.FLAG_GRANT_READ_URI_PERMISSION
}
val shareIntent = Intent.createChooser(sendIntent, "미리보기 제목")
startActivity(shareIntent)
```

### 맞춤 타겟 추가

```kotlin
val shareIntent = Intent.createChooser(sendIntent, null).apply {
  putExtra(Intent.EXTRA_CHOOSER_TARGETS, myChooserTargetArray)
  putExtra(Intent.EXTRA_INITIAL_INTENTS, myInitialIntentArray)
}
startActivity(shareIntent)
```

### 구성요소별 타겟 제거

```kotlin
val shareIntent = Intent.createChooser(sendIntent, null).apply {
  putExtra(Intent.EXTRA_EXCLUDE_COMPONENTS, myComponentArray)
}
startActivity(shareIntent)
```

### 공유 정보 가져오기

```kotlin
val pendingIntent = PendingIntent.getBroadcast(myContext, requestCode, Intent(myContext, MyBroadcastReceiver.class),
Intent.FLAG_UPDATE_CURRENT)
Intent.createChooser(Intent(Intent.ACTION_SEND), null, pendingIntent.intentSender)
```

### 콜백 수신

```kotlin
override fun onReceive(context: Context, intent: Intent) {
    val clickedComponent : ComponentName = intent.getParcelableExtra(EXTRA_CHOSEN_COMPONENT)
}
```