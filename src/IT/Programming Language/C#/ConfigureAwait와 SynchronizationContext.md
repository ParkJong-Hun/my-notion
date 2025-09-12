# ConfigureAwait와 SynchronizationContext

SynchronizationContext의 작동 원리

```csharp
// UI 애플리케이션에서의 문제
public async void Button_Click(object sender, EventArgs e)
{
    // UI 스레드에서 시작
    label.Text = "Loading...";
    
    // 백그라운드 스레드에서 실행
    var data = await FetchDataAsync();
    
    // ConfigureAwait(false) 없으면 UI 스레드로 돌아옴
    label.Text = data; // UI 업데이트는 UI 스레드에서만 가능
}

// 라이브러리 코드 - 성능 최적화
public async Task<string> FetchDataAsync()
{
    // ConfigureAwait(false) 사용 - UI 스레드로 돌아가지 않음
    var response = await httpClient.GetAsync(url).ConfigureAwait(false);
    var content = await response.Content.ReadAsStringAsync().ConfigureAwait(false);
    
    // 현재 스레드: 백그라운드 스레드
    return ProcessData(content); // CPU 집약적 작업은 백그라운드에서
}
```

Custom SynchronizationContext

```csharp
public class CustomSynchronizationContext : SynchronizationContext
{
    private readonly BlockingCollection<(SendOrPostCallback callback, object state)> _queue
        = new BlockingCollection<(SendOrPostCallback, object)>();
    
    private readonly Thread _thread;
    
    public CustomSynchronizationContext()
    {
        _thread = new Thread(ProcessQueue) { IsBackground = false };
        _thread.Start();
    }
    
    public override void Post(SendOrPostCallback d, object state)
    {
        _queue.Add((d, state));
    }
    
    public override void Send(SendOrPostCallback d, object state)
    {
        if (Thread.CurrentThread == _thread)
        {
            d(state);
        }
        else
        {
            var mre = new ManualResetEventSlim();
            Exception exception = null;
            
            _queue.Add(((s) =>
            {
                try { d(s); }
                catch (Exception ex) { exception = ex; }
                finally { mre.Set(); }
            }, state));
            
            mre.Wait();
            if (exception != null) throw exception;
        }
    }
    
    private void ProcessQueue()
    {
        SetSynchronizationContext(this);
        
        foreach (var (callback, state) in _queue.GetConsumingEnumerable())
        {
            callback(state);
        }
    }
}
```

데드락 방지

```csharp
public class DeadlockExample
{
    // 데드락 발생 코드
    public string BadMethod()
    {
        // UI 스레드에서 호출되면 데드락!
        return FetchDataAsync().Result;
    }
    
    // 올바른 방법들
    public async Task<string> GoodMethodAsync()
    {
        return await FetchDataAsync().ConfigureAwait(false);
    }
    
    public string GoodMethodSync()
    {
        // 동기 컨텍스트 제거
        return Task.Run(async () => await FetchDataAsync()).Result;
    }
    
    private async Task<string> FetchDataAsync()
    {
        await Task.Delay(1000).ConfigureAwait(false);
        return "Data";
    }
}
```