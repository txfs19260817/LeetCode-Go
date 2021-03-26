import sys


def title2path(title: str):
    path = title.replace(' ', '', 1)
    path = path.replace(' ', '-')
    path = '0' * (4 - len(path.split('.')[0])) + path
    filename = title.replace(' ', '') + ".go"
    return path + '/' + filename


if __name__ == '__main__':
    print(title2path(" ".join(sys.argv[1:])))
