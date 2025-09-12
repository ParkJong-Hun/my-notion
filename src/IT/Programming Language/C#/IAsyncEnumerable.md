# IAsyncEnumerable

대용량 데이터 스트리밍

```csharp
public class DataProcessor
{
    // 전체를 메모리에 로드하지 않고 스트리밍
    public async IAsyncEnumerable<ProcessedData> ProcessLargeDatasetAsync(
        string connectionString,
        [EnumeratorCancellation] CancellationToken cancellationToken = default)
    {
        using var connection = new SqlConnection(connectionString);
        await connection.OpenAsync(cancellationToken);
        
        using var command = new SqlCommand("SELECT * FROM LargeTable", connection);
        using var reader = await command.ExecuteReaderAsync(cancellationToken);
        
        while (await reader.ReadAsync(cancellationToken))
        {
            cancellationToken.ThrowIfCancellationRequested();
            
            var rawData = new RawData
            {
                Id = reader.GetInt32("Id"),
                Name = reader.GetString("Name"),
                // ... 다른 필드들
            };
            
            var processed = await ProcessSingleItemAsync(rawData, cancellationToken);
            yield return processed;
            
            // 메모리 사용량을 일정하게 유지
        }
    }
    
    private async Task<ProcessedData> ProcessSingleItemAsync(RawData raw, CancellationToken ct)
    {
        // 비동기 처리 로직
        await Task.Delay(10, ct); // 예시
        return new ProcessedData { /* ... */ };
    }
}

// 사용
await foreach (var item in processor.ProcessLargeDatasetAsync(connectionString)
    .WithCancellation(cancellationToken))
{
    Console.WriteLine($"Processed: {item.Id}");
    // 한 번에 하나씩 처리, 메모리 효율적
}
```

백프레셔 처리

```csharp
public async IAsyncEnumerable<T> WithBackpressure<T>(
    this IAsyncEnumerable<T> source,
    int maxConcurrency,
    [EnumeratorCancellation] CancellationToken cancellationToken = default)
{
    using var semaphore = new SemaphoreSlim(maxConcurrency, maxConcurrency);
    
    await foreach (var item in source.WithCancellation(cancellationToken))
    {
        await semaphore.WaitAsync(cancellationToken);
        
        try
        {
            yield return item;
        }
        finally
        {
            semaphore.Release();
        }
    }
}
```