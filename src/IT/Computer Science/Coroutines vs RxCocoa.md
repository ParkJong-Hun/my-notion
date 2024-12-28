# Coroutines vs RxCocoa

| Coroutines | RxCocoa | 비고 |
| --- | --- | --- |
| Flow |  | 콜드 스트림. |
| StateFlow | BehaviorRelay |  |
| SharedFlow | PublishRelay |  |
|  | Signal |  |
|  | Driver |  |
| emit | accept |  |
| collect | bind |  |
| flowOn | observeOn / subscribeOn |  |
| collectLatest | flatMapLatest |  |
| retry / retryWhen | retry / retryWhen |  |
| catch | catchError |  |
| combine | combineLatest |  |
| zip | zip |  |
| debounce | debounce |  |
| distinctUntilChanged | distinctUntilChanged |  |
| mapLatest | switchMap |  |