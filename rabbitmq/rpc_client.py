import pika
import uuid


class FibonacciRpcClient(object):

    def __init__(self):
        self.response = None
        self.corr_id = str(uuid.uuid4())  # 关联id

        credentials = pika.PlainCredentials('admin', 'admin')  # 认证
        parameters = pika.ConnectionParameters('localhost', credentials=credentials)  # 连接的变量

        self.connection = pika.BlockingConnection(parameters)

        self.channel = self.connection.channel()

        result = self.channel.queue_declare(queue='', exclusive=True)

        self.callback_queue = result.method.queue

        self.channel.basic_consume(
            queue=self.callback_queue,
            on_message_callback=self.on_response,
            auto_ack=True
        )

    def on_response(self, ch, method, props, body):
        # 检测correlation_id 是否一致
        # 如果一致保存到响应
        if self.corr_id == props.correlation_id:
            self.response = body

    # RPC request
    def call(self, n):
        # publish
        self.channel.basic_publish(
            exchange='',
            routing_key='rpc',
            # 属性信息
            # replay_to 回调队列
            # correlation_id 可以绑定request和response
            # 保证res属于req
            properties=pika.BasicProperties(
                reply_to=self.callback_queue,
                correlation_id=self.corr_id
            ),
            body=str(n)
        )

        while self.response is None:
            self.connection.process_data_events()
        return int(self.response)


fibonacci_rpc = FibonacciRpcClient()

print(" [x] Requesting fib(10)")
response = fibonacci_rpc.call(10)
print(" [.] Got %r" % response)
