package gateway

import (
	"context"

	"github.com/eduardotecnologo/ChatGpt/internal/domain/entity"
)

// Contratos (Ex: interfaces)
type ChatGateway interface {
	CreateChat(ctx context.Context, chat *entity.Chat) error
	FindChatByID(ctx context.Context, chatID string) (*entity.Chat, error)
	SaveChat(ctx context.Context, chat *entity.Chat) error
}
