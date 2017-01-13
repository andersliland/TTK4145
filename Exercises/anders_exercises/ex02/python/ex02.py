# Python 3.3.3 and 2.7.6
# python helloworld_python.py

from threading import Thread, Lock


Num = 0

def incNumFunc(lock):
    global Num
    for i in range(10):
        lock.aquire(True)
        Num += 1
        print("incNum:" + str(Num))
        lock.release() #release lock, no matter what

def decNumFunc(lock):
    global Num
    for i in range(11):
        lock.aquire(True)
        Num -= 1
        print("denNum:" + str(Num))
        lock.release()

def main():
    lock = Lock()

    incNumThread = Thread(target = incNumFunc, args = (lock),)
    incNumThread.start()

    decNumThread = Thread(target = decNumFunc, args = (lock),)
    decNumThread.start()
    
    incNumThread.join()
    decNumThread.join()

    print("Hello from main!")
    print("Num:" + str(Num))


main()