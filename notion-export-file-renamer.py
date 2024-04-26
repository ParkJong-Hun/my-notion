import os
import re

file_regexp = r'( \w{32,32}\b)'  # 파일 이름에서 사용한 정규식 패턴입니다.

def main():
    for root, dirs, files in os.walk(os.getcwd(), topdown=False):
        for f in files:
            f_path = os.path.join(root, f)
            if f.endswith('.md'):  # .md 파일인지 확인합니다.
                with open(f_path, 'r', encoding='utf-8') as file:
                    content = file.read()
                substituted_content = re.sub(r'\[.*?\]\((.*?)\)', lambda x: '[' + os.path.splitext(os.path.basename(x.group(1)))[0] + '](' + os.path.splitext(os.path.basename(x.group(1)))[0] + ')', content)
                # 위의 코드는 []() 패턴을 찾아서 파일명과 동일한 조건으로 변경합니다.
                with open(f_path, 'w', encoding='utf-8') as file:
                    file.write(substituted_content)

                # 파일 내의 링크 URL에 적용한 후에 파일명 변경을 수행합니다.
                subtitute_f_path = re.sub(file_regexp, '', os.path.abspath(f_path))
                os.renames(f_path, subtitute_f_path)

if __name__ == "__main__":
    main()
