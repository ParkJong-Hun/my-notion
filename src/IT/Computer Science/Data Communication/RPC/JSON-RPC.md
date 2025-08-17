# JSON-RPC

<aside>
💡

Remote Procedure Call을 JSON 형식으로 실현하는 프로토콜.
클라이언트가 서버측에서 정의된 함수나 메소드를 호출해 그 결과를 받는 구조.

클라이언트가 보내는 요청과 서버가 반환하는 응답은 ID로 연결.

MCP 클라이언트와 MCP 서버 간 통신을 할 때 주로 채용.

</aside>

### Request

```json
{
	"jsonrpc" : "2.0",
	"id": 1,
	"method": "list_resources",
	"params": []
}
```

### Response

```json
{
	"jsonrpc" : "2.0",
	"id": 1,
	"result": [
		{"uri": "example//resource", "name": "Example Resource"}
	]
}
```

### Error

```json
{
  "jsonrpc": "2.0",
  "error": {
    "code": -32601,
    "message": "Method not found"
  },
  "id": 1
}
```

### Notification

클라이언트 혹은 서버가 상대방에게 정보를 송신하는데, 응답을 기대하지 않는 단방향 리퀘스트로, ID를 포함하지 않음.

```json
{
  "jsonrpc": "2.0",
  "method": "someMethod",
  "params": {...}
}
```