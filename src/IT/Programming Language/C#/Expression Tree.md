# Expression Tree

<aside>
💡

코드를 데이터 구조로 표현하는 방식.

Kotlin의 리플렉션이나 Go의 리플렉션 패키지와 비슷하지만 더 강력.

</aside>

```csharp
// 일반 람다 (컴파일 시 IL 코드로 변환)
Func<int, bool> funcDelegate = x => x > 5;

// Expression Tree (데이터 구조로 표현)
Expression<Func<int, bool>> expression = x => x > 5;
```

### 사용 사례

- 내부 구조 탐색
    - 
    
    ```csharp
    Expression<Func<Person, bool>> expr = p => p.Age > 18 && p.Name.StartsWith("A");
    
    // Expression Tree 분석
    var binaryExpr = (BinaryExpression)expr.Body;
    Console.WriteLine(binaryExpr.NodeType); // AndAlso
    
    var leftSide = (BinaryExpression)binaryExpr.Left;
    Console.WriteLine(leftSide.NodeType); // GreaterThan
    
    var memberAccess = (MemberExpression)leftSide.Left;
    Console.WriteLine(memberAccess.Member.Name); // Age
    
    var constant = (ConstantExpression)leftSide.Right;
    ```
    
- 동적 쿼리 생성
    - 
    
    ```csharp
    public static Expression<Func<T, bool>> CreateFilter<T>(string propertyName, object value)
    {
        var parameter = Expression.Parameter(typeof(T), "x");
        var property = Expression.Property(parameter, propertyName);
        var constant = Expression.Constant(value);
        var equal = Expression.Equal(property, constant);
        
        return Expression.Lambda<Func<T, bool>>(equal, parameter);
    }
    
    // 사용
    var filter = CreateFilter<Person>("Age", 25);
    var people = dbContext.People.Where(filter); // SELECT * FROM People WHERE Age = 25
    ```
    
- Entity Framework에서 활용
    - 
    
    ```csharp
    // 이것이 가능한 이유는 Expression Tree 때문
    var adults = dbContext.People
        .Where(p => p.Age > 18)  // Expression<Func<Person, bool>>
        .Select(p => new { p.Name, p.Email })  // Expression<Func<Person, object>>
        .ToList();
    
    // EF Core가 내부적으로 Expression을 SQL로 변환
    // WHERE Age > 18
    ```