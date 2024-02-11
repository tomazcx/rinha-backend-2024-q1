CREATE TABLE IF NOT EXISTS cliente (id INTEGER PRIMARY KEY, limite INTEGER, saldo INTEGER);

INSERT INTO cliente (id, limite, saldo)
VALUES
(1, 1000 * 100, 0),
(2, 800 * 100, 0),
(3, 10000 * 100, 0),
(4, 100000 * 100, 0),
(5, 5000 * 100, 0);

CREATE TABLE IF NOT EXISTS transacao (id SERIAL PRIMARY KEY, valor INTEGER, tipo CHAR, descricao TEXT, realizada_em TIMESTAMP DEFAULT(NOW()), id_cliente INTEGER, CONSTRAINT fk_cliente FOREIGN KEY(id_cliente) REFERENCES cliente(id));
