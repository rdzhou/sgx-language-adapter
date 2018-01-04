---------------------
Purpose of Sgx_Go
---------------------
The project demonstrates several fundamental example APIs created with Go.
These example APIs enables Go to perform enclave operations in Intel(R) Software Guard Extensions (SGX):
    - Initialize and destroy an enclave.
    - Generate a random byte array using SGX trusted library function (sgx_read_rand).

------------------------------------
How to Build/Execute the Sample Code
------------------------------------
- Install Intel(R) SGX SDK for Linux* OS
- export LD_LIBRARY_PATH=$LD_LIBRARY_PATH:`pwd`
- Build the project with the prepared Makefile:
    a. Hardware Mode, Debug build:
        $ make
    b. Hardware Mode, Pre-release build:
        $ make SGX_PRERELEASE=1 SGX_DEBUG=0
    c. Hardware Mode, Release build:
        $ make SGX_DEBUG=0
    d. Simulation Mode, Debug build:
        $ make SGX_MODE=SIM
    e. Simulation Mode, Pre-release build:
        $ make SGX_MODE=SIM SGX_PRERELEASE=1 SGX_DEBUG=0
    f. Simulation Mode, Release build:
        $ make SGX_MODE=SIM SGX_DEBUG=0
- Execute example Go app:
    $ go run GoApp.go [optional: size of byte array]
- Remember to "make clean" before switching build mode

------------------------------------------
Brief Introduction of Each Folder and File
------------------------------------------
- Enclave: defines the functions running in the enclave
- LibSgxC: defines ecall/ocall functions that communicates to the enclave
- LibSgxGo: Go package that wraps the shared library
- GoApp.go: the example Go application
