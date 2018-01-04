package LibSgxGo
/*
#cgo LDFLAGS: -ldl
#include <dlfcn.h>
#include <stdio.h>
#include <stdlib.h>

int (*fp_ctypes_initialize_enclave)(void);
void (*fp_ctypes_sgx_destroy_enclave)(void);
unsigned char* (*fp_ctypes_ecall_sgx_read_rand)(int len, unsigned char *byteArrayPtr);

int ctypes_initialize_enclave() {
    return fp_ctypes_initialize_enclave();
}

void ctypes_sgx_destroy_enclave(void) {
    fp_ctypes_sgx_destroy_enclave();
}

unsigned char* ctypes_ecall_sgx_read_rand(int len, void*byteArrayPtr) {
    return fp_ctypes_ecall_sgx_read_rand(len, byteArrayPtr);
}

void *handle;

void loadLib() {
    char *error;

    handle = dlopen ("./libSgx.so", RTLD_LAZY);
    if (!handle) {
        fputs (dlerror(), stderr);
        exit(1);
    }

    fp_ctypes_initialize_enclave = dlsym(handle, "ctypes_initialize_enclave");
    if ((error = dlerror()) != NULL)  {
        fputs(error, stderr);
        exit(1);
    }

    fp_ctypes_sgx_destroy_enclave = dlsym(handle, "ctypes_sgx_destroy_enclave");
    if ((error = dlerror()) != NULL)  {
        fputs(error, stderr);
        exit(1);
    }

    fp_ctypes_ecall_sgx_read_rand = dlsym(handle, "ctypes_ecall_sgx_read_rand");
    if ((error = dlerror()) != NULL)  {
        fputs(error, stderr);
        exit(1);
    }
}

void closeLib() {
    if(NULL != handle) {
        dlclose(handle);
        handle = NULL;
    }
}
*/
import "C"

import (
    "fmt"
    "unsafe"
    "os"
)

func InitializeLibSgx() {
    fmt.Printf("loading DLL...\n")
    C.loadLib()
}


// Initialize the enclave
func Sgxfunction_initialize_enclave() {
    if (C.ctypes_initialize_enclave() < 0){
        fmt.Println("Failed to initiate enclave! Exiting...")
        os.Exit(-1)
    }
    fmt.Println("Successfully initialized enclave!")
    return
}   

// Ecall random number generator in enclave
func Sgxfunction_ecall_sgx_read_rand(bytes_number int) []byte{
    array := make([]byte, bytes_number)
    buf := C.CBytes(array)
    length := C.int(bytes_number)
    C.ctypes_ecall_sgx_read_rand(length, buf)
    ret := C.GoBytes(buf, length)
    C.free(unsafe.Pointer(buf))
    return ret
}

// Destroy the enclave
func Sgxfunction_sgx_destroy_enlave() {
    C.ctypes_sgx_destroy_enclave()
    return
}
