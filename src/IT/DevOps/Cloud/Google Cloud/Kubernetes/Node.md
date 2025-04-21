# Node

<aside>
💡

Kubernetes의 클러스터 안에서 실제 컨테이너가 실행되는 서버.

</aside>

[Cluster%201dcf37315c4480d8aca0e1d415d22dc2](Cluster%201dcf37315c4480d8aca0e1d415d22dc2)

**Master Node (Control Plane)**

[Control%20Plane%201dcf37315c44805eb625f8ac17b3f1fc](Control%20Plane%201dcf37315c44805eb625f8ac17b3f1fc)

- 클러스터 전체를 관리.

**Worker Node**

- 실제로 앱(Pod)이 실행됨.
- **kubelet**
    - 각 Node에서 Pod을 실행하고 관리하는 에이전트.
    - kube-apiserver와 통신해 상태 보고.
- **kube-proxy**
    - 클러스터 내에서 서비스와 라우팅.
    - Service 리소스를 기반으로 트래픽을 올바른 Pod으로 전달.
- Container Runtime
    - 실제로 컨테이너를 실행시키는 도구.