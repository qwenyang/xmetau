import random
def orderList(num):
    orderNums = []
    for i in range(num):
        orderNums.append(i+1)
    
    for i in range(num):
        ri = random.randint(0, num - 1)
        orderNums[i], orderNums[ri] = orderNums[ri], orderNums[i]
    return orderNums

