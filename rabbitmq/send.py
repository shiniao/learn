import pika

# 创建连接
credentials = pika.PlainCredentials('admin', 'admin')
parameters = pika.ConnectionParameters('localhost', credentials=credentials)
connection = pika.BlockingConnection(parameters)

channel = connection.channel()


# 创建一个队列：hello
channel.queue_declare(queue='hello')

# publish message
channel.basic_publish(exchange='', routing_key='hello', body='Hello World ggg')
print('[x] Sent Hello World')

# close
connection.close()


