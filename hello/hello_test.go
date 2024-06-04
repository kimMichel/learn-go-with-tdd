package hello

import "testing"

func TestHello(t *testing.T) {
	checkCorrectText := func(t *testing.T, result, expected string) {
		t.Helper()
		if result != expected {
			t.Errorf("result: '%s', expected: '%s'", result, expected)
		}
	}

	t.Run("say hello to people", func(t *testing.T) {
		result := Hello("Michel", "")
		expected := "Hello Michel"

		checkCorrectText(t, result, expected)
	})

	t.Run("say 'Hello World' when the parameter is empty", func (t *testing.T) {
		result := Hello("", "")
		expected := "Hello World"

		checkCorrectText(t, result, expected)
	})

	t.Run("say Hello in portuguese", func (t *testing.T)  {
		result := Hello("Michel", "portuguese")
		expected := "Ola Michel"

		checkCorrectText(t, result, expected)
	})

	t.Run("say Hello in spanish", func (t *testing.T)  {
		result := Hello("Michel", "spanish")
		expected := "Hola Michel"

		checkCorrectText(t, result, expected)
	})

	t.Run("say Hello in french", func (t *testing.T)  {
		result := Hello("Michel", "french")
		expected := "Bonjour Michel"

		checkCorrectText(t, result, expected)
	})
}