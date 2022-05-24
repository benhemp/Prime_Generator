# Prime_Generator
Go code to generate candidates for primes based on an idea I had. This POC could then be extended to be used with a sieve program to reduce the amount of time the sieve takes to run. 
General idea: in 6*n-1 and 6*n+1, every prime number greater than 2 or 3 exists. 
I had the idea that (5*n)+(n-1) = 6n-1 and (5*n)+(n+1) = 6n+1 and wondered what would happen if we substituted other prime numbers in for 5 in this form and sure enough, they also generate primes at a greater than circumstance rate. 

First POC generates the first 100 of what I would deem "prime candidates" for the first 10 prime numbers. 

This should help me find any other interesting patterns. 
