# Python 3.3.3
# python3 helloworld_python.py

from threading import Thread

# Potentially useful thing:
#   In Python you "import" a global variable, instead of "export"ing it when you declare it
#   (This is probably an effort to make you feel bad about typing the word "global")

#global num
num = 0

def threadIncNum():
    global num
    for i in range(1000000):
        num += 1

def threadDecNum():
    global num
    for i in range(1000000):
        num -= 1


def main():
    thread1 = Thread(target = threadIncNum, args = (),)
    thread1.start()

    thread2 = Thread(target = threadDecNum, args = (),)
    thread2.start()

    thread1.join()
    thread2.join()

    print("Num: " + str(num))

main()