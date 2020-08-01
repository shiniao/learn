import os

from work import work_on, Worker, Helper, work_on_env
from unittest import TestCase, mock


class TestWorkMockingModule(TestCase):

    @mock.patch('work.os')
    def test_work_on(self, work_os):
        work_on()
        work_os.getcwd.assert_called_once()

    def test_class_worker(self):
        # 加入autospec选项代表纠正调用错误
        with mock.patch('work.Helper', autospec=True) as mock_helper:

           # mocking the helper return value 'testing'
            mock_helper.return_value.get_path.return_value = 'testing'
            worker = Worker()

           # 测试 Helper 是否被调用了一次
            mock_helper.assert_called_once_with('db')

            # 测试 Worker 的返回值
            self.assertEqual(worker.work(), 'testing')

    def test_partial_patching(self):
        # patch 也可以只在某些方法上，而不是这个类
        with mock.patch.object(Helper, 'get_path', return_value='testing'):
            worker = Worker()
            self.assertEqual(worker.helper.path, 'db')
            self.assertEqual(worker.work(), 'testing')

    @mock.patch('os.getcwd', return_value='/home/')
    @mock.patch('work.print')
    @mock.patch.dict('os.environ', {'MY_VAR':'testing'})
    def test_work_on_env(self, mock_print, mock_getcwd):
        self.assertEqual(work_on_env(), '/home/testing')
        mock_print.assert_called_once_with('Working on /home/testing')
