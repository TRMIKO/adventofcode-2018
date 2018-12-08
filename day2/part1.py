# part 1

fp = open("input.txt")
par=[0,0]
while True:
    b = fp.readline()
    if b == '':
        break
    #print(b)

    pares=[False,False]

    for i in range(0,len(b)):
        z=1
        for j in range(0,len(b)):
            if i == j:
                continue
            if b[i] == b[j]:
                z=z+1
        if z==2:
            pares[0]=True
        if z==3:
            pares[1]=True
    if pares[0]==True:
        par[0]=par[0]+1
    if pares[1]==True:
        par[1]=par[1]+1

print(par)
print(par[0]*par[1])
    