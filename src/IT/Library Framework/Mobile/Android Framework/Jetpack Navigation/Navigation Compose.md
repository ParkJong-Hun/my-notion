# Navigation Compose

### **NavOptions**

**popUpTo**

- 지정한 루트와 일치하는 화면을 백스택에서 찾아서 백스택에서 그 화면으로부터 현재 화면까지의 스택을 없애고, 화면을 이동함.
- 만약 백스택에 루트가 없으면 액션은 무시됨.
- **inclusive**
    - true로 설정하면 스택에서 지정한 루트도 포함해서 삭제.
- **saveState**
    - 백스택에서 제거되는 화면들의 상태를 저장.
        - remember는 저장되지 않고, rememberSaveable만 저장.
            - remember도 저장하고싶다면, 서드파티 라이브러리인 Circuit이나 Rin의 rememberRetained를 사용.

**restoreState**

- 화면 이동시, popUpTo로 해당 화면으로 이전에 saveState로 저장한 백스택에서 삭제한 화면들의 상태를 복원해서 현재 백스택의 상단에 추가.
    - popUpTo를 했을 때 inclusive 설정에 따라 사용 가능한 화면이 제한됨.
    - 1회 밖에 할 수 없음.