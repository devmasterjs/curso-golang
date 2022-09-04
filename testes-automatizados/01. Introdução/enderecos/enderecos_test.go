package enderecos_test

// go test ./... --coverprofile cobertura.txt
// go tool cover --func=cobertura.txt
// go tool cover --html=cobertura.txt

import (
	. "introducao-testes/enderecos"
	"testing"
)

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoDeEndereco(t *testing.T) {
	t.Parallel()

	// given
	cenariosDeTeste := []cenarioDeTeste{
		{"Rua ABC", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"Estrada das Missões", "Estrada"},
		{"Rodovia dos Imigrantes", "Rodovia"},
		{"Praça da Luz", "Tipo inválido"},
		{"", "Tipo inválido"},
	}

	// when
	for _, cenario := range cenariosDeTeste {
		tipoDeEnderecoRecebido := TipoDeEndereco(cenario.enderecoInserido)

		// then
		if tipoDeEnderecoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido [%s] é diferente do esperado [%s]", tipoDeEnderecoRecebido, cenario.retornoEsperado)
		}
	}

}

func TipoDeEndereco(s string) {
	panic("unimplemented")
}
