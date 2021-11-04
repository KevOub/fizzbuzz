# Python program to print Fizz Buzz
 
# Loop for 100 times i.e. range
for fizzbuzz in range(1,10000):
 
    # Number divisible by 15,(divisible
    # by both 3 & 5), print 'FizzBuzz'
    # in place of the number
    if fizzbuzz % 15 == 0:
        print("FizzBuzz")                                        
        continue
 
    # Number divisible by 3, print 'Fizz'
    # in place of the number
    elif fizzbuzz % 3 == 0:    
        print("Fizz")                                        
        continue
 
    # Number divisible by 5,
    # print 'Buzz' in
    # place of the number
    elif fizzbuzz % 5 == 0:        
        print("Buzz")                                    
        continue
 
    # Print numbers
    print(fizzbuzz)
