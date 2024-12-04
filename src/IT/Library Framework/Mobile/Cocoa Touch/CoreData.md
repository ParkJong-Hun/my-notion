# CoreData

<aside>
💡

Apple의 객체 그래프 관리, 영속성 프레임워크.

</aside>

**NSManagedObjectContext**

핵심 객체.

데이터를 메모리에서 다루는 환경 제공.

CRUD 작업을 행함.

**NSManagedObject**

데이터 모델 객체.

엔티티와 매핑하는 클래스로 데이터 객체를 저장하고 관리하는데 사용.

**NsPersistentStoreCoordinator**

저장소 관리 담당 객체.

데이터를 실제로 저장할 위치 지정.

**NSPersistentContainer**

iOS 10 이상에서 사용되는 객체로 모든 Core Data의 스택을 간편하게 설정하고 관리.

**NSFetchRequest**

데이터를 검색할 때 사용하는 객체.