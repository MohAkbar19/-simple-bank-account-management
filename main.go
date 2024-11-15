package main

import (
    "fmt"
    "my_project/bank"
)

func main() {
    account := bank.Account{}
    
    // Menambahkan deposit
    if err := account.Deposit(100.50); err != nil {
        fmt.Println("Error:", err)
    }
    
    // Mencoba melakukan penarikan
    if err := account.Withdraw(50); err != nil {
        fmt.Println("Error:", err)
    }
    
    // Mencoba menarik lebih dari saldo yang tersedia
    if err := account.Withdraw(100); err != nil {
        fmt.Println("Error:", err)
    }

    fmt.Println("Saldo Akhir:", account.GetBalance())
}
