package fvm_test

import (
	"encoding/hex"
	"errors"
	"fmt"
	"testing"

	"github.com/onflow/cadence/runtime"
	"github.com/onflow/cadence/runtime/sema"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"

	"github.com/onflow/flow-go/engine/execution/testutil"
	"github.com/onflow/flow-go/fvm"
	"github.com/onflow/flow-go/fvm/state"
	"github.com/onflow/flow-go/model/flow"
	"github.com/onflow/flow-go/utils/unittest"
)

func makeTwoAccounts(t *testing.T, aPubKeys []flow.AccountPublicKey, bPubKeys []flow.AccountPublicKey) (flow.Address, flow.Address, *state.State) {

	ledger := state.NewMapLedger()
	st := state.NewState(ledger)

	a := flow.HexToAddress("1234")
	b := flow.HexToAddress("5678")

	//create accounts
	accounts := state.NewAccounts(st)
	err := accounts.Create(aPubKeys, a)
	require.NoError(t, err)
	err = accounts.Create(bPubKeys, b)
	require.NoError(t, err)

	err = st.Commit()
	require.NoError(t, err)

	return a, b, st
}

func TestAccountFreezing(t *testing.T) {

	t.Run("setFrozenAccount can be enabled", func(t *testing.T) {

		address, _, st := makeTwoAccounts(t, nil, nil)
		accounts := state.NewAccounts(st)

		// account should no be frozen
		frozen, err := accounts.GetAccountFrozen(address)
		require.NoError(t, err)
		require.False(t, frozen)

		rt := runtime.NewInterpreterRuntime()
		log := zerolog.Nop()
		vm := fvm.New(rt)
		txInvocator := fvm.NewTransactionInvocator(log)

		code := fmt.Sprintf(`
			transaction {
				execute {
					setAccountFrozen(0x%s, true)
				}
			}
		`, address.String())

		proc := fvm.Transaction(&flow.TransactionBody{Script: []byte(code)}, 0)

		context := fvm.NewContext(log, fvm.WithAccountFreezeAvailable(false))

		err = txInvocator.Process(vm, &context, proc, st)
		require.Error(t, err)
		require.Contains(t, err.Error(), "cannot find")
		require.Contains(t, err.Error(), "setAccountFrozen")

		context = fvm.NewContext(log, fvm.WithAccountFreezeAvailable(true))

		err = txInvocator.Process(vm, &context, proc, st)
		require.NoError(t, err)

		// account should be frozen now
		frozen, err = accounts.GetAccountFrozen(address)
		require.NoError(t, err)
		require.True(t, frozen)
	})

	t.Run("frozen account is rejected", func(t *testing.T) {

		txChecker := fvm.NewTransactionAccountFrozenChecker()

		frozenAddress, notFrozenAddress, st := makeTwoAccounts(t, nil, nil)
		accounts := state.NewAccounts(st)

		// freeze account
		err := accounts.SetAccountFrozen(frozenAddress, true)
		require.NoError(t, err)

		// make sure freeze status is correct
		frozen, err := accounts.GetAccountFrozen(frozenAddress)
		require.NoError(t, err)
		require.True(t, frozen)

		frozen, err = accounts.GetAccountFrozen(notFrozenAddress)
		require.NoError(t, err)
		require.False(t, frozen)

		// Authorizers

		// no account associated with tx so it should work
		tx := fvm.Transaction(&flow.TransactionBody{}, 0)
		err = txChecker.Process(nil, &fvm.Context{}, tx, st)
		require.NoError(t, err)

		tx = fvm.Transaction(&flow.TransactionBody{Authorizers: []flow.Address{notFrozenAddress}}, 0)
		err = txChecker.Process(nil, &fvm.Context{}, tx, st)
		require.NoError(t, err)

		tx = fvm.Transaction(&flow.TransactionBody{Authorizers: []flow.Address{frozenAddress}}, 0)
		err = txChecker.Process(nil, &fvm.Context{}, tx, st)
		require.Error(t, err)

		// all addresses must not be frozen
		tx = fvm.Transaction(&flow.TransactionBody{Authorizers: []flow.Address{frozenAddress, notFrozenAddress}}, 0)
		err = txChecker.Process(nil, &fvm.Context{}, tx, st)
		require.Error(t, err)

		// Payer should be part of authorizers account, but lets check it separately for completeness

		tx = fvm.Transaction(&flow.TransactionBody{Payer: notFrozenAddress}, 0)
		err = txChecker.Process(nil, &fvm.Context{}, tx, st)
		require.NoError(t, err)

		tx = fvm.Transaction(&flow.TransactionBody{Payer: frozenAddress}, 0)
		err = txChecker.Process(nil, &fvm.Context{}, tx, st)
		require.Error(t, err)

		// Proposal account

		tx = fvm.Transaction(&flow.TransactionBody{ProposalKey: flow.ProposalKey{Address: frozenAddress}}, 0)
		err = txChecker.Process(nil, &fvm.Context{}, tx, st)
		require.Error(t, err)

		tx = fvm.Transaction(&flow.TransactionBody{ProposalKey: flow.ProposalKey{Address: notFrozenAddress}}, 0)
		err = txChecker.Process(nil, &fvm.Context{}, tx, st)
		require.NoError(t, err)
	})

	t.Run("code from frozen account cannot be loaded", func(t *testing.T) {

		frozenAddress, notFrozenAddress, st := makeTwoAccounts(t, nil, nil)
		accounts := state.NewAccounts(st)

		rt := runtime.NewInterpreterRuntime()

		log := zerolog.Nop()
		txInvocator := fvm.NewTransactionInvocator(log)
		vm := fvm.New(rt)

		// deploy code to accounts
		whateverContractCode := `
			pub contract Whatever {
				pub fun say() {
					log("Düsseldorf")
				}
			}
		`

		deployContract := []byte(fmt.Sprintf(
			`
			 transaction {
			   prepare(signer: AuthAccount) {
				   signer.contracts.add(name: "Whatever", code: "%s".decodeHex())
			   }
			 }
	   `, hex.EncodeToString([]byte(whateverContractCode)),
		))

		procFrozen := fvm.Transaction(&flow.TransactionBody{Script: deployContract, Authorizers: []flow.Address{frozenAddress}, Payer: frozenAddress}, 0)
		procNotFrozen := fvm.Transaction(&flow.TransactionBody{Script: deployContract, Authorizers: []flow.Address{notFrozenAddress}, Payer: notFrozenAddress}, 0)
		deployContext := fvm.NewContext(zerolog.Nop(), fvm.WithServiceAccount(false), fvm.WithRestrictedDeployment(false), fvm.WithCadenceLogging(false))
		deployTxInvocator := fvm.NewTransactionInvocator(zerolog.Nop())
		deployRt := runtime.NewInterpreterRuntime()

		deployVm := fvm.New(deployRt)

		err := deployTxInvocator.Process(deployVm, &deployContext, procFrozen, st)
		require.NoError(t, err)
		err = deployTxInvocator.Process(deployVm, &deployContext, procNotFrozen, st)
		require.NoError(t, err)

		// both contracts should load now
		context := fvm.NewContext(log, fvm.WithCadenceLogging(true))

		code := func(a flow.Address) []byte {
			return []byte(fmt.Sprintf(`
				import Whatever from 0x%s
	
				transaction {
					execute {
						Whatever.say()
					}
				}
			`, a.String()))
		}

		// code from not frozen loads fine
		proc := fvm.Transaction(&flow.TransactionBody{Script: code(frozenAddress)}, 0)

		err = txInvocator.Process(vm, &context, proc, st)
		require.NoError(t, err)
		require.Len(t, proc.Logs, 1)
		require.Contains(t, proc.Logs[0], "Düsseldorf")

		proc = fvm.Transaction(&flow.TransactionBody{Script: code(notFrozenAddress)}, 0)

		err = txInvocator.Process(vm, &context, proc, st)
		require.NoError(t, err)
		require.Len(t, proc.Logs, 1)
		require.Contains(t, proc.Logs[0], "Düsseldorf")

		// freeze account
		err = accounts.SetAccountFrozen(frozenAddress, true)
		require.NoError(t, err)

		// make sure freeze status is correct
		frozen, err := accounts.GetAccountFrozen(frozenAddress)
		require.NoError(t, err)
		require.True(t, frozen)

		frozen, err = accounts.GetAccountFrozen(notFrozenAddress)
		require.NoError(t, err)
		require.False(t, frozen)

		// loading code from frozen account triggers error
		proc = fvm.Transaction(&flow.TransactionBody{Script: code(frozenAddress)}, 0)

		err = txInvocator.Process(vm, &context, proc, st)
		require.Error(t, err)

		// find frozen account specific error
		require.IsType(t, runtime.Error{}, err)
		err = err.(runtime.Error).Err

		require.IsType(t, &runtime.ParsingCheckingError{}, err)
		err = err.(*runtime.ParsingCheckingError).Err

		require.IsType(t, &sema.CheckerError{}, err)
		checkerErr := err.(*sema.CheckerError)

		checkerErrors := checkerErr.ChildErrors()

		require.Len(t, checkerErrors, 2)
		require.IsType(t, &sema.ImportedProgramError{}, checkerErrors[0])

		importedCheckerError := checkerErrors[0].(*sema.ImportedProgramError).Err

		accountFrozenError := &state.AccountFrozenError{}

		require.True(t, errors.As(importedCheckerError, &accountFrozenError))
		require.Equal(t, frozenAddress, accountFrozenError.Address)
	})

	t.Run("default settings allow only service account to freeze accounts", func(t *testing.T) {

		rt := runtime.NewInterpreterRuntime()
		log := zerolog.Nop()
		vm := fvm.New(rt)
		// create default context
		context := fvm.NewContext(log)

		ledger := testutil.RootBootstrappedLedger(vm, context)

		privateKeys, err := testutil.GenerateAccountPrivateKeys(1)
		require.NoError(t, err)

		// Bootstrap a ledger, creating accounts with the provided private keys and the root account.
		accounts, err := testutil.CreateAccounts(vm, ledger, privateKeys, context.Chain)
		require.NoError(t, err)

		address := accounts[0]

		code := fmt.Sprintf(`
			transaction {
				execute {
					setAccountFrozen(0x%s, true)
				}
			}
		`, address.String())

		txBody := &flow.TransactionBody{Script: []byte(code)}

		err = testutil.SignTransaction(txBody, accounts[0], privateKeys[0], 0)
		require.NoError(t, err)

		tx := fvm.Transaction(txBody, 0)
		err = vm.Run(context, tx, ledger)
		require.NoError(t, err)
		require.Error(t, tx.Err)

		require.Contains(t, tx.Err.Error(), "cannot find")
		require.Contains(t, tx.Err.Error(), "setAccountFrozen")
		require.Equal(t, (&fvm.ExecutionError{}).Code(), tx.Err.Code())

		// sign tx by service account now
		txBody = &flow.TransactionBody{Script: []byte(code)}
		txBody.SetProposalKey(context.Chain.ServiceAddress(), 0, 0)
		txBody.SetPayer(context.Chain.ServiceAddress())

		err = testutil.SignPayload(txBody, accounts[0], privateKeys[0])
		require.NoError(t, err)

		err = testutil.SignEnvelope(txBody, context.Chain.ServiceAddress(), unittest.ServiceAccountPrivateKey)
		require.NoError(t, err)

		tx = fvm.Transaction(txBody, 0)
		err = vm.Run(context, tx, ledger)
		require.NoError(t, err)

		require.NoError(t, tx.Err)

	})

	t.Run("service account cannot freeze itself", func(t *testing.T) {

		rt := runtime.NewInterpreterRuntime()
		log := zerolog.Nop()
		vm := fvm.New(rt)
		// create default context
		context := fvm.NewContext(log)

		ledger := testutil.RootBootstrappedLedger(vm, context)

		privateKeys, err := testutil.GenerateAccountPrivateKeys(1)
		require.NoError(t, err)

		// Bootstrap a ledger, creating accounts with the provided private keys and the root account.
		accounts, err := testutil.CreateAccounts(vm, ledger, privateKeys, context.Chain)
		require.NoError(t, err)

		address := accounts[0]

		codeAccount := fmt.Sprintf(`
			transaction {
				execute {
					setAccountFrozen(0x%s, true)
				}
			}
		`, address.String())

		serviceAddress := context.Chain.ServiceAddress()
		codeService := fmt.Sprintf(`
			transaction {
				execute {
					setAccountFrozen(0x%s, true)
				}
			}
		`, serviceAddress.String())

		// sign tx by service account now
		txBody := &flow.TransactionBody{Script: []byte(codeAccount)}
		txBody.SetProposalKey(serviceAddress, 0, 0)
		txBody.SetPayer(serviceAddress)

		err = testutil.SignPayload(txBody, accounts[0], privateKeys[0])
		require.NoError(t, err)

		err = testutil.SignEnvelope(txBody, serviceAddress, unittest.ServiceAccountPrivateKey)
		require.NoError(t, err)

		tx := fvm.Transaction(txBody, 0)
		err = vm.Run(context, tx, ledger)
		require.NoError(t, err)
		require.NoError(t, tx.Err)

		accountsService := state.NewAccounts(state.NewState(ledger))

		frozen, err := accountsService.GetAccountFrozen(address)
		require.NoError(t, err)
		require.True(t, frozen)

		// make sure service account is not frozen before
		frozen, err = accountsService.GetAccountFrozen(serviceAddress)
		require.NoError(t, err)
		require.False(t, frozen)

		// service account cannot be frozen
		txBody = &flow.TransactionBody{Script: []byte(codeService)}
		txBody.SetProposalKey(serviceAddress, 0, 1)
		txBody.SetPayer(serviceAddress)

		err = testutil.SignPayload(txBody, accounts[0], privateKeys[0])
		require.NoError(t, err)

		err = testutil.SignEnvelope(txBody, serviceAddress, unittest.ServiceAccountPrivateKey)
		require.NoError(t, err)

		tx = fvm.Transaction(txBody, 0)
		err = vm.Run(context, tx, ledger)
		require.NoError(t, err)
		require.Error(t, tx.Err)

		accountsService = state.NewAccounts(state.NewState(ledger))

		frozen, err = accountsService.GetAccountFrozen(serviceAddress)
		require.NoError(t, err)
		require.False(t, frozen)
	})
}
