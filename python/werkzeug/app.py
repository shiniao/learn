from werkzeug import run_simple
from werkzeug.wrappers import Response
from python.werkzeug.shortly import Shortly

def create_app(redis_host='localhost', redis_port=6379):
    app = Shortly({
        'redis_host': redis_host,
        'redis_port': redis_port
    })

    return app


if __name__ == '__main__':
    app = create_app()
    run_simple('127.0.0.1', 5000, app, use_debugger=True, use_reloader=True)
