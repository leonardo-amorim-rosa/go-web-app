package models

import "github.com/loja/db"

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscarTodosOsProdutos() []Produto {
	db := db.ConectaComBanco()
	resultado, err := db.Query("select * from produtos order by id asc")

	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}
	for resultado.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err := resultado.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade
		produtos = append(produtos, p)
	}

	defer db.Close() // fecha a conexão com o banco de dados
	return produtos
}

func CriarNovoProduto(nome, descricao string, preco float64, quantidade int) {
	db := db.ConectaComBanco()

	insereDadosNoBanco, err := db.Prepare("Insert into produtos (nome, descricao, preco, quantidade) values ($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	insereDadosNoBanco.Exec(nome, descricao, preco, quantidade)
	defer db.Close()
}

func ExcluirProduto(id string) {
	db := db.ConectaComBanco()

	excluirProduto, err := db.Prepare("Delete from produtos where id=$1")
	if err != nil {
		panic(err.Error())
	}

	excluirProduto.Exec(id)
	defer db.Close()
}

func EditarProduto(id string) Produto {
	db := db.ConectaComBanco()

	produtoParaEditar, err := db.Query("select * from produtos where id=$1", id)
	if err != nil {
		panic(err.Error())
	}

	produtoParaAtualizar := Produto{}

	for produtoParaEditar.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		produtoParaEditar.Scan(&id, &nome, &descricao, &preco, &quantidade)

		produtoParaAtualizar.Id = id
		produtoParaAtualizar.Nome = nome
		produtoParaAtualizar.Descricao = descricao
		produtoParaAtualizar.Preco = preco
		produtoParaAtualizar.Quantidade = quantidade
	}

	defer db.Close()
	return produtoParaAtualizar
}

func AtualizarProduto(id int, nome, descricao string, preco float64, quantidade int) {
	db := db.ConectaComBanco()

	atualizaProdutoNoBanco, err := db.Prepare("update produtos set nome=$1, descricao=$2, preco=$3, quantidade=$4 where id=$5")
	if err != nil {
		panic(err.Error())
	}

	atualizaProdutoNoBanco.Exec(nome, descricao, preco, quantidade, id)
	defer db.Close()
}
