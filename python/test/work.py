import os


def work_on():
    path = os.getcwd()
    print(f'Working on {path}')
    return path

def work_on_env():
    path = os.path.join(os.getcwd(), os.environ('MY_VAR'))
    print(f'Working on {path}')
    return path



class Helper:

    def __init__(self, path):
        self.path = path

    def get_path(self):
        base_path = os.getcwd()
        return os.path.join(base_path, self.path)


class Worker:

    def __init__(self):
        self.helper = Helper('db')

    def work(self):
        path = self.helper.get_path()
        print('Working on {path}')
        return path
