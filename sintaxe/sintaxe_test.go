package sintaxe

import "testing"

func TestVerificar(t *testing.T) {

	assert_eq := func(got, want bool) {
		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	}

	t.Run("teste 1", func(t *testing.T) {
		got := Verificar("AvB")
		want := true
		assert_eq(got, want)
	})

	t.Run("teste 2", func(t *testing.T) {
		got := Verificar("~AvB")
		want := true
		assert_eq(got, want)
	})

	t.Run("teste 3", func(t *testing.T) {
		got := Verificar("~Av~B")
		want := true
		assert_eq(got, want)
	})

	t.Run("teste 4", func(t *testing.T) {
		got := Verificar("-BvA")
		want := false
		assert_eq(got, want)
	})

	t.Run("teste 5", func(t *testing.T) {
		got := Verificar("~~~~A")
		want := true
		assert_eq(got, want)
	})

	t.Run("teste 6", func(t *testing.T) {
		got := Verificar("~~~~A~~~~~")
		want := false
		assert_eq(got, want)
	})

	t.Run("teste 7", func(t *testing.T) {
		got := Verificar("A")
		want := true
		assert_eq(got, want)
	})

	t.Run("teste 8", func(t *testing.T) {
		got := Verificar("Av~A")
		want := true
		assert_eq(got, want)
	})

	t.Run("teste 9", func(t *testing.T) {
		got := Verificar("~B")
		want := true
		assert_eq(got, want)
	})

	t.Run("teste 10", func(t *testing.T) {
		got := Verificar("~Bv")
		want := false
		assert_eq(got, want)
	})

	t.Run("teste 11", func(t *testing.T) {
		got := Verificar("~~~~~~~~~~")
		want := false
		assert_eq(got, want)
	})

	t.Run("teste 12", func(t *testing.T) {
		got := Verificar("~~~~~~~~~~S=SvT^P-K")
		want := true
		assert_eq(got, want)
	})

	t.Run("teste 13", func(t *testing.T) {
		got := Verificar("~vA")
		want := false
		assert_eq(got, want)
	})

	t.Run("teste 14", func(t *testing.T) {
		got := Verificar("A~A")
		want := false
		assert_eq(got, want)
	})

	t.Run("teste 15", func(t *testing.T) {
		got := Verificar("~BA")
		want := false
		assert_eq(got, want)
	})

	t.Run("teste 16", func(t *testing.T) {
		got := Verificar("~BvBv~~~~~~~A")
		want := true
		assert_eq(got, want)
	})

	t.Run("teste 17", func(t *testing.T) {
		got := Verificar("~B^^A")
		want := false
		assert_eq(got, want)
	})

	t.Run("teste 18", func(t *testing.T) {
		got := Verificar("B~^A")
		want := false
		assert_eq(got, want)
	})

	t.Run("teste 19", func(t *testing.T) {
		got := Verificar("B")
		want := true
		assert_eq(got, want)
	})

	t.Run("teste 20", func(t *testing.T) {
		got := Verificar("A~~~~~~~~A")
		want := false
		assert_eq(got, want)
	})

	t.Run("teste 21", func(t *testing.T) {
		got := Verificar("")
		want := false
		assert_eq(got, want)
	})
}
