num = list(map(int, input().split()))
product = num[0] * num[1]
print(num)
if product % 2 == 1:
    print("Odd")
else:
    print("Even")
    