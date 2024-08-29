package activities

import (
	"context"
	"time"

	"github.com/billowdev/exclusive-go-hexa/internal/adapters/temporal/dto"
)

// func RegisterDataActivity(ctx context.Context, data dto.RegistrationData) (*pagination.Pagination[[]models.SystemField], error) {
// 	// Implement the logic to register data
// 	tx := database.ExtractTx(ctx)
// 	if tx != nil {
// 		// transactorRepo := database.NewTransactorRepo(tx)
// 		sfRepo := repositories.NewSystemFieldRepo(tx)
// 		data, err := sfRepo.GetSystemFields(ctx)
// 		if err != nil {
// 			return nil, err
// 		}
// 		return data, nil
// 	}
// 	// time.Sleep(2 * time.Second)
// 	return nil, nil
// }

func RegisterDataActivity(ctx context.Context, data dto.RegistrationData) error {
	// Implement the logic to register data

	time.Sleep(2 * time.Second)
	return nil
}
func CheckDataActivity(ctx context.Context, data dto.RegistrationData) (string, error) {
	// Implement the logic to check the data status
	time.Sleep(2 * time.Second)

	// return "pending", nil // Example status
	return "pending", nil // Example status
}

func VerifyIdentitiesActivity(ctx context.Context, data dto.RegistrationData) error {
	// Implement the logic to verify identities
	time.Sleep(4 * time.Second)
	return nil
}

func UpdateStatusActivity(ctx context.Context, status string) error {
	// Implement the logic to update the registration status
	time.Sleep(1 * time.Second)
	return nil
}
