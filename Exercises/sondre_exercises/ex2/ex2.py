# Python 3.3.3
# python3 helloworld_python.py

from threading import Thread, Lock

# Potentially useful thing:
#   In Python you "import" a global variable, instead of "export"ing it when you declare it
#   (This is probably an effort to make you feel bad about typing the word "global")

num = 0

def threadIncNum(lock):
    global num
    for i in range(1000000):
        lock.acquire(True)
        num += 1
        lock.release()

def threadDecNum(lock):
    global num
    for i in range(1000001):
        lock.acquire(True)
        num -= 1
        lock.release()

def main():
    lock = Lock()

    thread1 = Thread(target = threadIncNum, args = (lock,))
    thread1.start()

    thread2 = Thread(target = threadDecNum, args = (lock,))
    thread2.start()

    thread1.join()
    thread2.join()

    print("Num: " + str(num))

main()