import os
import sys


def list2slice(s: str):
    res = "[]int" + s.replace('[', '{').replace(']', '}') # TODO: support more types
    return "[]" + res if s.startswith("[[") else res

if __name__ == '__main__':
    inputs = " ".join(sys.argv[1:]).strip()
    print(list2slice(inputs))
