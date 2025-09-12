# C#

### 언어적 특징

- Properties
    - Java의 getter/setter를 언어 차원에서 지원
    - 
    
    ```csharp
    public class Person 
    {
        public string Name { get; set; }  // 자동 프로퍼티
        
        private string _email;
        public string Email 
        {
            get => _email;
            set => _email = value?.ToLower();  // 커스텀 로직
        }
    }
    ```
    
- LINQ (Language Integrated Query)
    - 컬렉션 조작을 SQL과 유사한 구문으로 가능
    - Kotlin보다 더 강력
    - 
    
    [Expression%20Tree%2026cf37315c4480ce86adf32024165075](Expression%20Tree%2026cf37315c4480ce86adf32024165075)
    
    ```csharp
    var adults = people
        .Where(p => p.Age >= 18)
        .Select(p => new { p.Name, p.Email })
        .OrderBy(p => p.Name);
    ```
    
- Null 관련 기능
    - Kotlin의 Null safety와 비슷하지만 Opt-in
    - 
    
    ```csharp
    string? nullableString = null;  // nullable reference type
    string nonNullString = "";      // non-nullable
    
    // null-coalescing operators
    string result = nullableString ?? "default";
    nullableString ??= "assigned if null";
    ```
    
- async/await
    - JS와 유사하지만 더 체계적
    - 
    
    [ConfigureAwait%EC%99%80%20SynchronizationContext%2026cf37315c4480ddb830ef3927c67f41](ConfigureAwait%EC%99%80%20SynchronizationContext%2026cf37315c4480ddb830ef3927c67f41)
    
    [IAsyncEnumerable%2026cf37315c448064a78cd3d157e18b42](IAsyncEnumerable%2026cf37315c448064a78cd3d157e18b42)
    
    ```csharp
    public async Task<string> FetchDataAsync()
    {
        using var client = new HttpClient();
        return await client.GetStringAsync("https://api.example.com");
    }
    ```
    
- Pattern Matching
    - Rust의 match와 비슷하지만 더 유연
    - 
    
    ```csharp
    var result = value switch
    {
        0 => "Zero",
        > 0 and < 10 => "Single digit",
        string s when s.Length > 5 => "Long string",
        _ => "Other"
    };
    ```
    
- Expression-bodied members
    - 
    
    ```csharp
    public string FullName => $"{FirstName} {LastName}";
    public void Log(string msg) => Console.WriteLine($"[LOG] {msg}");
    ```
    

### 메모리 관리

- Garbage Collection
    - Java와 유사하지만, `struct`를 통해 스택 할당 가능
    - 
    
    ```csharp
    struct Point  // 값 타입, 스택에 할당
    {
        public int X, Y;
    }
    
    class Rectangle  // 참조 타입, 힙에 할당
    {
        public Point TopLeft;
    }
    ```
    
- using과 IDisposable
    - Go의 defer나 Swift의 defer와 비슷
    - 
    
    ```csharp
    using var file = File.OpenRead("data.txt");
    // 스코프를 벗어날 때 자동으로 file.Dispose() 호출
    ```
    

### 고성능 기능

- Span<T>와 Memory<T>
    - Rust의 slice와 비슷하게 zero-copy 조작 가능
    - 
    
    [Span%20T%20%EC%99%80%20Memory%20T%20%EC%98%88%2026cf37315c4480a398a7fcc07c40a31b](Span%20T%20%EC%99%80%20Memory%20T%20%EC%98%88%2026cf37315c4480a398a7fcc07c40a31b)
    
    ```csharp
    ReadOnlySpan<char> span = "Hello World".AsSpan(6, 5);  // "World"
    ```
    
- unsafe
    - 필요시 포인터 직접 다룰 수 있음
    - 
    
    ```csharp
    unsafe
    {
        fixed (byte* ptr = buffer)
        {
            // 포인터 직접 조작
        }
    }
    ```
    

### 함수형 프로그래밍

- Delegates와 Events
    - 
    
    [Event%EC%9D%98%20Thread%20Safety%2026cf37315c4480cbb6abfa82ecc11420](Event%EC%9D%98%20Thread%20Safety%2026cf37315c4480cbb6abfa82ecc11420)
    
    [Multicast%20Delegate%2026cf37315c4480cdbaa0ffede1976907](Multicast%20Delegate%2026cf37315c4480cdbaa0ffede1976907)
    
    ```csharp
    Action<string> logger = msg => Console.WriteLine(msg);
    Func<int, int, int> add = (x, y) => x + y;
    
    // 이벤트 시스템
    public event Action<string> OnMessageReceived;
    ```
    
- Records
    - 불변 데이터 클래스를 쉽게 작성 가능
    - 
    
    ```csharp
    public record Person(string Name, int Age);
    
    var person1 = new Person("John", 30);
    var person2 = person1 with { Age = 31 };  // 새 인스턴스 생성
    ```
    

### 그외

- Source Generators
    - (Annotation으로 Compile time에 코드 생성하는 느낌)
    
    ```csharp
    [Generator]
    public class AutoNotifyGenerator : ISourceGenerator
    {
        public void Initialize(GeneratorInitializationContext context)
        {
            context.RegisterForSyntaxNotifications(() => new SyntaxReceiver());
        }
        
        public void Execute(GeneratorExecutionContext context)
        {
            if (!(context.SyntaxContextReceiver is SyntaxReceiver receiver))
                return;
            
            foreach (var classDeclaration in receiver.CandidateClasses)
            {
                var model = context.Compilation.GetSemanticModel(classDeclaration.SyntaxTree);
                var classSymbol = model.GetDeclaredSymbol(classDeclaration);
                
                var source = GeneratePropertyNotification(classSymbol);
                context.AddSource($"{classSymbol.Name}_AutoNotify.cs", source);
            }
        }
        
        private string GeneratePropertyNotification(INamedTypeSymbol classSymbol)
        {
            var namespaceName = classSymbol.ContainingNamespace.ToDisplayString();
            var className = classSymbol.Name;
            
            return $@"
    using System.ComponentModel;
    
    namespace {namespaceName}
    {{
        public partial class {className} : INotifyPropertyChanged
        {{
            public event PropertyChangedEventHandler PropertyChanged;
            
            protected virtual void OnPropertyChanged(string propertyName)
            {{
                PropertyChanged?.Invoke(this, new PropertyChangedEventArgs(propertyName));
            }}
        }}
    }}";
        }
    }
    
    // 사용
    [AutoNotify]
    public partial class Person
    {
        private string _name;
        public string Name
        {
            get => _name;
            set
            {
                _name = value;
                OnPropertyChanged(); // Source Generator가 생성한 메서드
            }
        }
    }
    ```