# Multicast Delegate

```csharp
Action action1 = () => Console.WriteLine("First");
Action action2 = () => Console.WriteLine("Second");
Action action3 = () => Console.WriteLine("Third");

// 내부적으로 linked list 구조
Action combined = action1 + action2 + action3;

// 실행 순서 보장
combined(); // First, Second, Third 순서로 실행

// 예외 처리 주의점
Action riskyAction = () => throw new Exception("Error");
Action safeAction = () => Console.WriteLine("Safe");

Action combinedRisky = riskyAction + safeAction;
// combinedRisky(); // 첫 번째에서 예외 발생 시 두 번째 실행 안됨
```

수동으로 안전한 실행

```csharp
public static void SafeInvoke(this MulticastDelegate multicast, params object[] args)
{
    if (multicast == null) return;
    
    foreach (Delegate del in multicast.GetInvocationList())
    {
        try
        {
            del.Method.Invoke(del.Target, args);
        }
        catch (Exception ex)
        {
            Console.WriteLine($"Error in {del.Method.Name}: {ex.Message}");
            // 다른 delegate들은 계속 실행
        }
    }
}
```