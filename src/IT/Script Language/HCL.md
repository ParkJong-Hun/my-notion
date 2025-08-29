# HCL

<aside>
💡

HashiCorp Configuration Language.
하시코프사가 개발한 설정 언어.
인프라의 구조와 설정을 정의하는 데 특화.
대표적인 사례: 테라폼.

</aside>

```hcl
# AWS EC2 인스턴스를 정의하는 블록
resource "aws_instance" "example_server" {
  ami           = "ami-0c55b159cbfafe1f0" # 사용할 서버 이미지 ID
  instance_type = "t2.micro"             # 서버 사양

  tags = {
    Name = "HelloWorld"
  }
}

# 동일한 내용의 JSON 코드
{
  "resource": {
    "aws_instance": {
      "example_server": {
        "ami": "ami-0c55b159cbfafe1f0",
        "instance_type": "t2.micro",
        "tags": {
          "Name": "HelloWorld"
        }
      }
    }
  }
}
```