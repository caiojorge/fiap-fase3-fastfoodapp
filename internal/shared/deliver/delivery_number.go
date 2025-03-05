package deliver

import (
	"fmt"
	"math/rand"
	"time"
)

// Função para gerar o número de entrega
func GenerateDeliveryNumber() string {
	// Cria uma nova fonte de aleatoriedade
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	// Gera duas letras aleatórias
	letters := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	randomLetters := string([]rune{
		letters[random.Intn(len(letters))],
		letters[random.Intn(len(letters))],
	})

	// Gera três números aleatórios
	randomNumbers := fmt.Sprintf("%03d", random.Intn(1000))

	// Combina as letras e números
	return randomLetters + randomNumbers
}
