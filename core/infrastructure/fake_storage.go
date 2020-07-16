package infrastructure

import (
	"encoding/json"
	"sync"

	"github.com/dc0d/bank_account_kata_go/core"
	"github.com/dc0d/bank_account_kata_go/core/model/account"
)

type FakeStorage struct {
	account.AccountRepo

	storage map[string]map[string][]byte // map bucket name -> id -> value
	mx      sync.RWMutex
}

func NewFakeStorage() *FakeStorage {
	res := &FakeStorage{
		storage: make(map[string]map[string][]byte),
	}

	res.storage[accountsBucketName] = make(map[string][]byte)

	return res
}

func (storage *FakeStorage) FindAccount(id account.AccountID) (*account.Account, error) {
	var (
		result account.Account
		err    error
	)

	storage.mx.RLock()
	defer storage.mx.RUnlock()

	accountsBucket := storage.storage[accountsBucketName]

	data, ok := accountsBucket[string(id)]
	if !ok {
		return nil, core.ErrNotFound
	}

	if err = json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (storage *FakeStorage) SaveAccount(clientAccount *account.Account) error {
	storage.mx.Lock()
	defer storage.mx.Unlock()

	accountsBucket := storage.storage[accountsBucketName]

	if oldData, ok := accountsBucket[string(clientAccount.AccountID)]; ok {
		var oldAccount account.Account
		if err := json.Unmarshal(oldData, &oldAccount); err != nil {
			return err
		}

		var newTransactions account.Transactions
		for i, tx := range clientAccount.Transactions {
			if i >= len(oldAccount.Transactions) {
				newTransactions = append(newTransactions, tx)
				continue
			}

			if oldAccount.Transactions[i] == tx {
				continue
			}

			newTransactions = append(newTransactions, tx)
		}

		clientAccount.Transactions = append(oldAccount.Transactions, newTransactions...)
	}

	data, err := json.Marshal(clientAccount)
	if err != nil {
		return err
	}

	accountsBucket[string(clientAccount.AccountID)] = data

	storage.storage[accountsBucketName] = accountsBucket

	return nil
}

const (
	accountsBucketName = "accounts"
)
