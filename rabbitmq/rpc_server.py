import pika

credentials = pika.PlainCredentials('admin', 'admin')
parameters = pika.ConnectionParameters('localhost', credentials=credentials)
connection = pika.BlockingConnection(parameters)

channel = connection.channel()

channel.queue_declare(queue='rpc')


def fib(n):
    if n == 0:
        return 0

    elif n == 1:
        return 1

    else:
        return fib(n - 1) + fib(n - 2)


def on_request(ch, method, props, body):
    n = int(body)  # body是传的值
    print(f'[. ] fib {n}')
    response = fib(n)  # res 计算fib结果

    # publish给client
    ch.basic_publish(exchange='', routing_key=props.reply_to,
                     properties=pika.BasicProperties(correlation_id=props.correlation_id),
                     body=str(response))

    ch.basic_ack(delivery_tag=method.delivery_tag)


channel.basic_qos(prefetch_count=1)
# consume：调用on_request,返回结果给client
channel.basic_consume(queue='rpc', on_message_callback=on_request)

print('Awaiting RPC requests...')
channel.start_consuming()
