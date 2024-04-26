# Content Provider

<aside>
💡 앱이 액세스할 수 있는 다른 모든 영구 저장 위치에 저장 가능한 앱 데이터의 공유형 집합을 관리.
음악이나 사진, 파일 같은 대용량 데이터를 공유할 때 사용.

</aside>

<aside>
⚠️ 작은 데이터들은 인텐트를 통해 공유할 수 있음.

</aside>

**Content Resolver**

각 앱의 ContentProvider를 중계.

Content Resolver가 URI를 요청하면 그에 맞게 Content Provider는 DB에서 데이터를 찾아 Content Resolver를 통해 다른 앱에 전달.

- query
    - uri: content://scheme 방식의 원하는 데이터를 가져오기 위한 정해진 주소.
    - projection: 가져올 컬럼 이름 목록, null이면 모든 컬럼.
    - selection: where 절에 해당하는 내용.
    - selectionArgs: selection에서 ?로 표시한 곳에 들어갈 데이터.
    - sortOrder: 정렬할 위한 order by 구문.

[Intent%2057034274b5f44a93bc6c95321b029778](Intent%2057034274b5f44a93bc6c95321b029778)