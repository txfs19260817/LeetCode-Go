import os


def remove_spaces(path):
    if not os.path.exists(path):
        raise FileNotFoundError("No such file or directory: ", path)
    for curDir, dirs, files in os.walk(path):
        for f in files:
            filename = os.path.join(curDir, f)
            os.rename(os.path.join(path, filename), os.path.join(path, filename.replace(' ', '')))


if __name__ == "__main__":
    remove_spaces("./leetcode") # TODO: input path from console
