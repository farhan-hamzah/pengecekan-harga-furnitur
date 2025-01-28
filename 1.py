N = float(input())

if N > 1000:
    N = N*0.8
elif N >= 500 and N <= 1000:
    N = N*0.85
print(N)