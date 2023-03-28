package util

import (
	"go-fp-crowdfundchat/model"
	"strconv"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

func GetTransactionURL(transaction *model.Transaction, user *model.User) (string, error) {
	midtrans.ServerKey = "SB-Mid-server-TJcg2sGMzLetZbvhcHpUzFjf"
	midtrans.Environment = midtrans.Sandbox
	var s snap.Client

	s.New(midtrans.ServerKey, midtrans.Sandbox)

	req := & snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
		  OrderID:  strconv.Itoa(transaction.ID),
		  GrossAmt: int64(transaction.Amount),
		},
		CustomerDetail: &midtrans.CustomerDetails{
		  FName: user.Name,
		  Phone: user.PhoneNo,
		},
	  }
	
	snapResp, err := s.CreateTransaction(req)
	if err != nil {
		return "", err
	}

	return snapResp.RedirectURL, nil
}