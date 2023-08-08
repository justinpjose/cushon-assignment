package postgres

const getAvailableFunds = `
	SELECT f.id, f."name" 
	FROM funds f 
	
	INNER JOIN companies_funds cf2
	on f.id = cf2.fund_id
	
	INNER JOIN customers c 
	on c.company_id  = cf2.company_id
	
	INNER JOIN customer_accounts ca 
	on ca.customer_id = c.id
	
	WHERE ca.customer_account_no = $1
	AND cf2.active = TRUE
	
	AND f.id NOT IN (
		SELECT fund_id
		FROM customer_accounts_funds caf
		WHERE caf.customer_account_no = $1
	);
`
