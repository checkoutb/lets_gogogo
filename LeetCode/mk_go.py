import os
import sys
import time
import platform

# mk xxxx.go

def mk_go():
    '''
    copy mk_cpp, group by per thousand
    <2000, group by thousand
    >=2000, group by 200
    '''

    lt_name = ""

    if (len(lt_name) < 5):
        lt_name = input("plz input name:")

    # print("-------------")

    lt_name = lt_name.replace("'", "")
    t2 = lt_name[0:lt_name.find('.')]
    mainName2 = "LT" + t2.zfill(4) + "_" + time.strftime("%Y%m%d", time.localtime())
    f_name = "/" + mainName2 + "_" + lt_name[lt_name.find('.') + 2 :].replace(" ", "_").replace("'", "") + ".go"
    dir_name = ""
    if (int(t2) < 2000):
        dir_name = "ge" + str(int(int(int(t2)/1000)*1000))
        if int(t2) < 1000:
            dir_name = "gt0000"
    else:
        dir_name = "ge" + str(int(int(int(t2)/200)*200))

    content = """// package sdq
package main

import (
    "fmt"
)




func main_""" + mainName2 + """() {
// func main() {

    fmt.Println("ans:")


}
"""

    # print(f_name)
    name = dir_name + f_name
    pwd = os.getcwd() + "/" + name
    print(pwd)

    # if os.path.isfile(pwd):
    if os.path.exists(name):
        print("already exists, so exit...")
        sys.exit()

    # 如果是/开头，那么是 根目录文件下。
    if not os.path.isdir(dir_name):
        os.makedirs(dir_name)

    # not a file, 已经能确保不会删除其他文件了。
    if not os.path.isfile(name):
        fd = open(name, mode="a+", encoding='utf-8')
        fd.write(content)
        fd.close()

    # .test is from .gitignore
    with open("last_go_path.test", 'w', encoding='utf-8') as f:
        f.write(name)

    # linux, use visual studio code open file
    if 'linux' in platform.system().lower():
        print(os.system("code " + name))


if __name__ == "__main__":
    mk_go()
    # mk_cpp()
    # print(mk_cpp.__doc__)
