package main

import "Sgx_Go/LibSgxGo"
import "fmt"
import "os"
import "strconv"

func printArray(array []byte) {
    fmt.Printf("\nGo byte array array[%x] after SGX running\n", len(array))

    for _, n := range(array) {
        fmt.Printf("%x ", n)
    }
    fmt.Println()
}


func main() {
    length := 8
    if len(os.Args) > 1 {
        res, _ := strconv.ParseInt(os.Args[1], 10, 32)
        length = int(res)
        if(length <= 0) {
            length = 8
        }
    }

    LibSgxGo.InitializeLibSgx()
    LibSgxGo.Sgxfunction_initialize_enclave()
    array := LibSgxGo.Sgxfunction_ecall_sgx_read_rand(length)
    printArray(array)
    LibSgxGo.Sgxfunction_sgx_destroy_enlave()
}
