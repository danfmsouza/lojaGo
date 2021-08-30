package models

import "github.com/danfmsouza/lojaGo/db"

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func FindAllProdutos() []Produto {
	db := db.ConectaComDB()

	selectProdutos, err := db.Query("select * from produtos order by id asc")
	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}
	for selectProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)

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
	defer db.Close()
	return produtos
}

func GravarProduto(nome, descricao string, preco float64, quantidade int) {
	db := db.ConectaComDB()

	insereProduto, err := db.Prepare("insert into produtos(nome, descricao, preco, quantidade) values($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}
	insereProduto.Exec(nome, descricao, preco, quantidade)
	defer db.Close()
}
func RemoveProduto(id string) {
	db := db.ConectaComDB()

	deletarProduto, err := db.Prepare("delete from produtos where id=$1")
	if err != nil {
		panic(err.Error())
	}
	deletarProduto.Exec(id)
}

func AlteraProduto(id string) Produto {
	db := db.ConectaComDB()

	selectProduto, err := db.Query("select * from produtos where id=$1", id)
	if err != nil {
		panic(err.Error())
	}

	atualizaProduto := Produto{}

	for selectProduto.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectProduto.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}
		atualizaProduto.Id = id
		atualizaProduto.Nome = nome
		atualizaProduto.Descricao = descricao
		atualizaProduto.Preco = preco
		atualizaProduto.Quantidade = quantidade
	}
	defer db.Close()
	return atualizaProduto
}

func AtualizarProduto(id int, nome, descricao string, preco float64, quantidade int) {
	db := db.ConectaComDB()

	AtualizaProduto, err := db.Prepare("update produtos set nome=$1, descricao=$2, preco=$3, quantidade=$4 where id=$5")
	if err != nil {
		panic(err.Error())
	}
	AtualizaProduto.Exec(nome, descricao, preco, quantidade, id)
	defer db.Close()
}
