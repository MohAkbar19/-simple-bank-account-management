package bank

import (
    "testing"
    . "github.com/onsi/ginkgo/v2"
    . "github.com/onsi/gomega"
)

func TestAccount(t *testing.T) {
    RegisterFailHandler(Fail)
    RunSpecs(t, "Account Suite")
}

var _ = Describe("Account", func() {
    var account Account

    BeforeEach(func() {
        account = Account{}
    })

    Describe("Deposit", func() {
        It("harus menambah saldo ketika jumlah positif didepositkan", func() {
            err := account.Deposit(100)
            Expect(err).To(BeNil())
            Expect(account.GetBalance()).To(Equal(100.0))
        })

        It("harus mengembalikan error untuk jumlah deposit nol atau negatif", func() {
            err := account.Deposit(-10)
            Expect(err).To(MatchError("jumlah deposit harus positif"))
            Expect(account.GetBalance()).To(Equal(0.0))
        })
    })

    Describe("Withdraw", func() {
        BeforeEach(func() {
            account.Deposit(100)
        })

        It("harus mengurangi saldo ketika jumlah yang valid ditarik", func() {
            err := account.Withdraw(50)
            Expect(err).To(BeNil())
            Expect(account.GetBalance()).To(Equal(50.0))
        })

        It("harus mengembalikan error untuk jumlah penarikan nol atau negatif", func() {
            err := account.Withdraw(-10)
            Expect(err).To(MatchError("jumlah penarikan harus positif"))
            Expect(account.GetBalance()).To(Equal(100.0))
        })

        It("harus mengembalikan error saat menarik lebih dari saldo", func() {
            err := account.Withdraw(200)
            Expect(err).To(MatchError("saldo tidak mencukupi"))
            Expect(account.GetBalance()).To(Equal(100.0))
        })
    })

    Describe("GetBalance", func() {
        It("harus mengembalikan saldo yang benar", func() {
            account.Deposit(100)
            Expect(account.GetBalance()).To(Equal(100.0))
        })
    })
})
