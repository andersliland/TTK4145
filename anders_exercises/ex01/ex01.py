# Python 3.3.3 and 2.7.6
# python helloworld_python.py

from threading import Thread

Num = 0

def incNumFunc():
    global Num
    for i in range(1000000):
    	Num = Num + 1

def decNumFunc():
	global Num
	for i in range(1000000):
		Num = Num -1

def main():
    incNumThread = Thread(target = incNumFunc, args = (),)
    incNumThread.start()

    decNumThread = Thread(target = decNumFunc, args = (),)
    decNumThread.start()
    
    incNumThread.join()
    decNumThread.join()

    print("Hello from main!")
    print("Num:" + str(Num))


main()