import os
import sys
import platform


def title2path(title: str):
    path = title.replace(' ', '', 1)
    path = path.replace(' ', '-')
    path = '0' * (4 - len(path.split('.')[0])) + path
    filename = title.replace(' ', '') + ".go"
    return path + '/' + filename


def copy2clipboard(s: str):
    cmd = ' | pbcopy'
    if "win" in platform.system().lower():
        cmd = ' | clip'
    cmd = 'echo ' + s + cmd
    os.system(cmd)


def create_go_file(relative_path: str, base_dir: str = "./leetcode"):
    """
    Create a .go file with 'package leetcode' content.
    
    Args:
        relative_path: The relative path including filename (e.g., "0001.Two-Sum/TwoSum.go")
        base_dir: The base directory where the file should be created (default: "./leetcode")
    """
    full_path = os.path.join(base_dir, relative_path)
    directory = os.path.dirname(full_path)
    
    # Create directory if it doesn't exist
    if not os.path.exists(directory):
        os.makedirs(directory)
        print(f"Created directory: {directory}")
    
    # Create the .go file with package leetcode content
    if not os.path.exists(full_path):
        with open(full_path, 'w') as f:
            f.write("package leetcode\n")
        print(f"Created file: {full_path}")


if __name__ == '__main__':
    res = title2path(" ".join(sys.argv[1:]))
    print(res)
    copy2clipboard(res)
    print("Copied to your clipboard!")
    
    # Create the .go file with package leetcode content
    create_go_file(res)
