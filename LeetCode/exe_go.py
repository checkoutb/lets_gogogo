import os
import sys
import time


def exe_go():
    go = None
    with open("last_go_path.test", "r", encoding='utf-8') as f:
        go = f.readline()
    
    if (go is None):
        return
    
    print(os.system("go run " + go))

if __name__ == "__main__":
    exe_go()
