
-- Adicionando restrição para não realizar empréstimos indevidos

CREATE UNIQUE INDEX book_loans_single_loan 
ON book_loans (book_id, from_user_id, to_user_id, (returned_at IS NULL))
WHERE returned_at IS NULL;