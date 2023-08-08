package postgres

const getCustomerAccountsFundByID = `
	SELECT
		c.id,
		c.customer_account_no,
		c.fund_id,
		c.total_amount
	FROM customer_accounts_funds c
	WHERE c.id = $1;
`

const getTotalAmount = `
	SELECT c.total_amount
	FROM customer_accounts_funds c
	WHERE c.id = $1;
`
