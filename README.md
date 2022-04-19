# Collatz conjecture
## Assessment test

This assessment test consists in:

1. create a go program that given a number n calculates the number of steps for n to reach 1 based on the Collatz conjecture<https://en.wikipedia.org/wiki/Collatz_conjecture>; 
2. set up a kubernetes cluster (e.g. using minikube) and expose the code as a service;
3. add a readme describing how to set it up and a script you can use to query the service for all values from 1...K for a given number K.

## Executing the service

Unfortunately, I couldn't finish the Kubernetes Cluster and Service Expose topic. I wrote a script to build an image and run a application container using Docker (I'll keep trying Kubernetes anyway, thanks for this challenge).

To run the applicantion, first you must clone this repository: `git clone git@github.com:lgmontenegro/collatz.git .`

After clone the repository, enter the directory you cloned it (it is collatz by default) and run `chmod +x runme.sh && chmod +x collatz.sh`

Now run `./runme.sh` to build and run the application and after that run `./collatz.sh N` where the N is the desirable number.

You can also open you browser and try `http://localhost:8080/collatz/N` where N is the number you want to check the Collatz conjecture.

**PS**: you **MUST** have Git and Docker installed.