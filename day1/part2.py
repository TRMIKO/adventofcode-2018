


# z=tc+r
# r=z-tc

# variables

cont = []
c = 0
val = 0
t = 0


# Funciones 

def valor(t,c,r):
    return t*c+r

def busca(t,c,z):
    return z-t*c

def pares(t,c1,r1):
    res=valor(t,c1,r1)
    print("buscare",res)
    for i in range(0,c1+1):
        for j in cont:
            #print(valor(t,i,j))
            if c1==i and r1==j:
                return False
            if res==valor(t,i,j):
                return True

def run(c,cont,t):
    while True:
        for r in cont:
            if pares(t,c,r):
                print(valor(t,c,r))
                return True
        c=c+1
        print("columna :",c)



fp = open("input.txt","r")

while True:
    b = fp.readline()
    if b == '':
        break
    t=t+int(b)
    cont.append(t)

fp.close()

# arr = [3,3,4,-2,-4]

# for i in arr:
#     t=t+i
#     cont.append(t)

run(c,cont,t)


