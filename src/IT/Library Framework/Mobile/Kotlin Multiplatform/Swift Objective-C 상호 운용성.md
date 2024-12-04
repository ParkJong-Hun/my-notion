# Swift/Objective-C 상호 운용성

| Kotlin | Swift | Objective-C |
| --- | --- | --- |
| class | class | @interface |
| interface | protocol; | @protocol |
| constuctor/create | Initializer | Initializer |
| Property | Property | Property |
| Method | Method | Method |
| enum class | class | @interface |
| suspend → | completionHandler: / async | completionHandler: |
| @Throws fun | throws | error: (NSError**)error |
| Extension | Extension | Category member |
| companion member ← | Class method or property | Class method or property |
| null | nil | nil |
| Singleton | shared or companion property | shared or companion property |
| Primitive type | Primitive type / NSNumber |  |
| Unit return type | Void | void |
| String | String | NSString |
| String | NSMutableString | NSMutableString |
| List | Array | NSArray |
| MutableList | NSMutableArray | NSMutableArray |
| Set | Set | NSSet |
| MutableSet | NSMutableSet | NSMutaebleSet |
| Map | Dictionary | NSDictionary |
| MutableMap | NSMutableDictionary | NSMutableDictionary |
| Function type | Function type | Block pointer type |
| Inline classes | Unsupported | Unsupported |