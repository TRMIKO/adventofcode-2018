file = open("input.txt",'r')

r=0

while True:
    b = file.readline()
    if b == '':
        break
    r=r+int(b)

print(r)