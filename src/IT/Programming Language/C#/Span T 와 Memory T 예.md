# Span<T>와 Memory<T> 예

메모리 할당 없는 문자열 처리

```csharp
public static class StringExtensions
{
    // 전통적인 방식 - 많은 할당 발생
    public static string[] SplitOld(this string input, char separator)
    {
        return input.Split(separator); // 새 배열과 문자열들 생성
    }
    
    // Span을 사용한 방식 - 할당 최소화
    public static void SplitSpan(this ReadOnlySpan<char> input, char separator, 
        Span<Range> ranges, out int count)
    {
        count = 0;
        int start = 0;
        
        for (int i = 0; i < input.Length; i++)
        {
            if (input[i] == separator)
            {
                ranges[count++] = new Range(start, i);
                start = i + 1;
            }
        }
        ranges[count++] = new Range(start, input.Length);
    }
}

// 사용
string csv = "name,age,city,country";
Span<Range> ranges = stackalloc Range[10];

csv.AsSpan().SplitSpan(',', ranges, out int count);
for (int i = 0; i < count; i++)
{
    ReadOnlySpan<char> part = csv.AsSpan(ranges[i]);
    // 새 문자열 생성 없이 원본의 일부분 참조
}
```

Binary 데이터 처리

```csharp
public static class BinaryProcessor
{
    public static int ReadInt32(ReadOnlySpan<byte> data, int offset)
    {
        // 안전한 범위 검사
        if (offset + sizeof(int) > data.Length)
            throw new ArgumentOutOfRangeException();
        
        // 직접 메모리 접근
        return BitConverter.ToInt32(data.Slice(offset, sizeof(int)));
    }
    
    public static void WriteInt32(Span<byte> data, int offset, int value)
    {
        if (offset + sizeof(int) > data.Length)
            throw new ArgumentOutOfRangeException();
        
        BitConverter.GetBytes(value).CopyTo(data.Slice(offset, sizeof(int)));
    }
}
```

stackalloc과의 조합

```csharp
public static string ProcessLargeString(string input)
{
    // 작은 버퍼는 스택에 할당
    Span<char> buffer = input.Length < 1024 
        ? stackalloc char[input.Length]
        : new char[input.Length];
    
    // 처리 로직
    for (int i = 0; i < input.Length; i++)
    {
        buffer[i] = char.ToUpper(input[i]);
    }
    
    return new string(buffer);
}
```