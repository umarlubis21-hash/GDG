package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func isValid(s string) bool {
    countG := strings.Count(s, "G")
    countC := strings.Count(s, "C")
    
    // Syarat 1: Jumlah G harus sama dengan C
    if countG != countC {
        return false
    }
    
    // Syarat 2: Tidak boleh mengandung urutan "DGD"
    if strings.Contains(s, "DGD") {
        return false
    }
    
    return true
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    
    fmt.Println("==============================================")
    fmt.Println("   VALIDATOR KODE RAHASIA GDGoC UNSRI         ")
    fmt.Println("==============================================")
    fmt.Println()
    
    fmt.Print("Masukkan jumlah kode rahasia: ")
    scanner.Scan()
    var a int
    fmt.Sscanf(scanner.Text(), "%d", &a)

    fmt.Println()
    fmt.Println("----------------------------------------------")
    fmt.Println()

    for i := 0; i < a; i++ {
        fmt.Printf("Kode #%d: ", i+1)
        scanner.Scan()
        kode := scanner.Text()

        fmt.Print("  Hasil Validasi: ")
        if isValid(kode) {
            fmt.Println("Valid")
        } else {
            fmt.Println("Tidak Valid")
        }
        fmt.Println()
    }
    
    fmt.Println("==============================================")
    fmt.Println("   TERIMA KASIH TELAH MENGGUNAKAN PROGRAM INI ")
    fmt.Println("==============================================")
}