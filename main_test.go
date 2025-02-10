package main

import (
	"fmt"
	"testing"
)

// Testes DN4 - Detecção de Contas Duplicadas e Fraudes

// ValidateCPF verifica se um CPF está na lista de banidos.
// Retorna false se o CPF estiver banido, true caso contrário.
func ValidateCPF(cpf string) bool {
	banned := CheckCPFInBanList(cpf) // Função fictícia que verifica se o CPF está banido.
	return !banned
}

// TestCPFBlock verifica se um CPF banido é corretamente identificado pelo sistema.
func TestCPFBlock(t *testing.T) {
	cpf := "123.456.789-00"
	if ValidateCPF(cpf) {
		t.Errorf("Esperava-se que o CPF %s fosse banido, mas não foi", cpf)
	}
}

// BenchmarkValidationSpeed mede o tempo necessário para validar CPFs.
func BenchmarkValidationSpeed(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ValidateCPF("987.654.321-00")
	}
}

// Testes DN5 - Experiência do Entregador e Feedback Estruturado

// SubmitFeedback simula a submissão de um feedback no sistema.
func SubmitFeedback(userID string, rating int, comment string) error {
	// Simulando submissão de feedback sem erros.
	return nil
}

// TestFeedbackSubmission verifica se a função SubmitFeedback funciona corretamente.
func TestFeedbackSubmission(t *testing.T) {
	err := SubmitFeedback("123", 4, "Ótima experiência!")
	if err != nil {
		t.Errorf("Falha na submissão do feedback: %v", err)
	}
}

// GetResponseRate retorna a taxa atual de resposta dos entregadores.
func GetResponseRate() int {
	return 28 // Simulando uma taxa de resposta baixa.
}

// Alert exibe uma mensagem de alerta quando necessário.
func Alert(message string) {
	fmt.Println("ALERTA:", message)
}

// TestLowResponseRateAlert verifica se um alerta é gerado quando a taxa de resposta está baixa.
func TestLowResponseRateAlert(t *testing.T) {
	if GetResponseRate() < 30 {
		Alert("Baixa taxa de resposta detectada, tome medidas.")
	}
}
