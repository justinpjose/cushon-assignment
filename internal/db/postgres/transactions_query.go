package postgres

const createTransaction = `
	INSERT INTO transactions
		(customer_accounts_funds_id, amount)
	VALUES
		($1, $2);
`

const updateTotalAmountInFund = `
	UPDATE customer_accounts_funds
	SET total_amount = $1
	WHERE id = $2;
`
