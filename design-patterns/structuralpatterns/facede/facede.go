package facede

import (
	"fmt"
	"log"
	"structuralpatterns/facede/models"
)

func Facede() {

	walletFacade := models.NewWalletFacade("abc", 1234)
	fmt.Println()
	err := walletFacade.AddMoneyToWallet("abc", 1234, 10)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}
	fmt.Println()
	err = walletFacade.DeductMoneyFromWallet("ab", 1234, 5)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}
}
