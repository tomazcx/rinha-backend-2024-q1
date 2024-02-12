package db

import (
	"context"
	"database/sql"
	"errors"

	"github.com/tomazcx/rinha-backend-2024-q1/config"
)

var ErrInvalidValue = errors.New("Invalid value")

type ClientRepository struct {
	db *sql.DB
}

type Balance struct {
	Total       int    `json:"total"`
	DataExtrato string `json:"data_extrato"`
	Limite      int    `json:"limite"`
}

type Transaction struct {
	Valor       int    `json:"valor"`
	Tipo        string `json:"tipo"`
	Descricao   string `json:"descricao"`
	RealizadaEm string `json:"realizada_em"`
}

type ClientExtract struct {
	Saldo             Balance       `json:"saldo"`
	UltimasTransacoes []Transaction `json:"ultimas_transacoes"`
}

func (r *ClientRepository) GetExtract(id int) (*ClientExtract, error) {
	rows, err := r.db.Query("SELECT c.saldo, NOW() as data_extrato, C.limite, T.valor, T.tipo, T.descricao, T.realizada_em FROM cliente C LEFT JOIN transacao T ON T.id_cliente = C.id WHERE C.id = $1;", id)

	if err != nil {
		return nil, err
	}

	var result ClientExtract
	result.UltimasTransacoes = make([]Transaction, 0)

	for rows.Next() {
		var transaction Transaction
		if err := rows.Scan(&result.Saldo.Total, &result.Saldo.DataExtrato, &result.Saldo.Limite, &transaction.Valor, &transaction.Tipo, &transaction.Descricao, &transaction.RealizadaEm); err != nil {
			break
		}

		result.UltimasTransacoes = append(result.UltimasTransacoes, transaction)
	}

	return &result, nil
}

type CreateTransactionDTO struct {
	Valor     int    `json:"valor"`
	Tipo      string `json:"tipo"`
	Descricao string `json:"descricao"`
}

type ClientData struct {
	Limite int `json:"limite"`
	Saldo  int `json:"saldo"`
}

func (r *ClientRepository) CreateTransaction(clientId int, transaction CreateTransactionDTO) (*ClientData, error) {
	tx, err := r.db.BeginTx(context.Background(), nil)	

	if err != nil {
		return nil, err
	}

	var clientBudget, clientLimit int 
	err = tx.QueryRow("SELECT saldo, limite FROM cliente WHERE id = $1 FOR UPDATE", clientId).Scan(&clientBudget, &clientLimit)

	if err != nil {
		return nil, tx.Rollback()
	}

	if (clientLimit < -(clientBudget - transaction.Valor)) && transaction.Tipo == "d" {
		_ = tx.Rollback()
		return nil, ErrInvalidValue
	}

	_, err = tx.Exec("INSERT INTO transacao (valor, tipo, descricao, id_cliente) VALUES ($1, $2, $3, $4)", transaction.Valor, transaction.Tipo, transaction.Descricao, clientId)

	if err != nil {
		return nil, tx.Rollback()
	}

	var amount int
	if transaction.Tipo == "d" {
		amount = -transaction.Valor
	} else {
		amount = transaction.Valor
	}

	var result ClientData
	row := tx.QueryRow("UPDATE cliente SET saldo = saldo + $1 WHERE id = $2 RETURNING saldo, limite", amount, clientId)
	err = row.Scan(&result.Saldo, &result.Limite)

	if err != nil {
		return nil, tx.Rollback()
	}

	return &result, tx.Commit()
}

func NewClientRepository() *ClientRepository {
	return &ClientRepository{
		db: config.GetDBConn(),
	}
}
