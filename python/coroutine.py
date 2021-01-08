import asyncio

async def say_what(what):
    return what


async def main():

    task1 = asyncio.create_task(say_what('hello')) 
    task2 = asyncio.create_task(say_what('world'))

    print(await task1)
    print(await task2)

asyncio.run(main())

