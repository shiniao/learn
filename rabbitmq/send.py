import pika


credentials = pika.PlainCredentials('admin', 'admin')
parameters = pika.ConnectionParameters('localhost', credentials=credentials)
connection = pika.BlockingConnection(parameters)

channel = connection.channel()

channel.queue_declare(queue='hello')

channel.basic_publish(exchange='', routing_key='hello', body='Hello World ggg')
print('[x] Sent Hello World')

connection.close()


