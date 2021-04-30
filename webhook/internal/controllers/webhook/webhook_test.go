package webhook

import (
	"errors"
	"testing"

	"github.com/ZupIT/horusec-devkit/pkg/services/database/enums"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/ZupIT/horusec-platform/webhook/internal/entities/webhook"
	enums2 "github.com/ZupIT/horusec-platform/webhook/internal/enums"
	repositoryWebhook "github.com/ZupIT/horusec-platform/webhook/internal/repositories/webhook"
)

func TestController_ListAll(t *testing.T) {
	t.Run("Should return all webhooks without errors", func(t *testing.T) {
		repoMock := &repositoryWebhook.Mock{}
		repoMock.On("ListAll").Return(&[]webhook.Webhook{{}}, nil)
		res, err := NewWebhookController(repoMock).ListAll(uuid.New())
		assert.NoError(t, err)
		assert.NotEmpty(t, res)
	})
	t.Run("Should return error unknown on list all webhooks", func(t *testing.T) {
		repoMock := &repositoryWebhook.Mock{}
		repoMock.On("ListAll").Return(&[]webhook.Webhook{}, errors.New("unexpected error"))
		res, err := NewWebhookController(repoMock).ListAll(uuid.New())
		assert.Error(t, err)
		assert.Empty(t, res)
	})
	t.Run("Should return not error but return empty list if data is nil", func(t *testing.T) {
		repoMock := &repositoryWebhook.Mock{}
		repoMock.On("ListAll").Return(&[]webhook.Webhook{}, nil)
		res, err := NewWebhookController(repoMock).ListAll(uuid.New())
		assert.NoError(t, err)
		assert.Empty(t, res)
	})
}

func TestController_Save(t *testing.T) {
	t.Run("Should save new webhook without error", func(t *testing.T) {
		repoMock := &repositoryWebhook.Mock{}
		repoMock.On("ListOne").Return(&webhook.Webhook{}, nil)
		repoMock.On("Save").Return(nil)
		webhookID, err := NewWebhookController(repoMock).Save(&webhook.Webhook{})
		assert.NoError(t, err)
		assert.NotEqual(t, uuid.Nil, webhookID)
	})
	t.Run("Should save new webhook with error duplicate webhook", func(t *testing.T) {
		repoMock := &repositoryWebhook.Mock{}
		repoMock.On("ListOne").Return(&webhook.Webhook{WebhookID: uuid.New()}, nil)
		repoMock.On("Save").Return(nil)
		webhookID, err := NewWebhookController(repoMock).Save(&webhook.Webhook{})
		assert.Error(t, err)
		assert.Equal(t, enums2.ErrorWebhookDuplicate, err)
		assert.Equal(t, uuid.Nil, webhookID)
	})
	t.Run("Should save new webhook with error unexpected on listone", func(t *testing.T) {
		repoMock := &repositoryWebhook.Mock{}
		repoMock.On("ListOne").Return(&webhook.Webhook{}, errors.New("unexpected error"))
		repoMock.On("Save").Return(nil)
		webhookID, err := NewWebhookController(repoMock).Save(&webhook.Webhook{})
		assert.Error(t, err)
		assert.NotEqual(t, enums2.ErrorWebhookDuplicate, err)
		assert.Equal(t, uuid.Nil, webhookID)
	})
	t.Run("Should save new webhook with error unexpected on save", func(t *testing.T) {
		repoMock := &repositoryWebhook.Mock{}
		repoMock.On("ListOne").Return(&webhook.Webhook{}, nil)
		repoMock.On("Save").Return(errors.New("unexpected error"))
		webhookID, err := NewWebhookController(repoMock).Save(&webhook.Webhook{})
		assert.Error(t, err)
		assert.Equal(t, uuid.Nil, webhookID)
	})
}

func TestController_Update(t *testing.T) {
	t.Run("Should update repository without error", func(t *testing.T) {
		repoMock := &repositoryWebhook.Mock{}
		repoMock.On("Update").Return(nil)
		err := NewWebhookController(repoMock).Update(&webhook.Webhook{}, uuid.New())
		assert.NoError(t, err)
	})
	t.Run("Should update repository with error not found", func(t *testing.T) {
		repoMock := &repositoryWebhook.Mock{}
		repoMock.On("Update").Return(enums.ErrorNotFoundRecords)
		err := NewWebhookController(repoMock).Update(&webhook.Webhook{}, uuid.New())
		assert.Error(t, err)
	})
	t.Run("Should update repository with error unexpected", func(t *testing.T) {
		repoMock := &repositoryWebhook.Mock{}
		repoMock.On("Update").Return(errors.New("unexpected error"))
		err := NewWebhookController(repoMock).Update(&webhook.Webhook{}, uuid.New())
		assert.Error(t, err)
	})
}

func TestController_Remove(t *testing.T) {
	t.Run("Should remove repository without error", func(t *testing.T) {
		repoMock := &repositoryWebhook.Mock{}
		repoMock.On("Remove").Return(nil)
		err := NewWebhookController(repoMock).Remove(uuid.New())
		assert.NoError(t, err)
	})
	t.Run("Should remove repository with error not found", func(t *testing.T) {
		repoMock := &repositoryWebhook.Mock{}
		repoMock.On("Remove").Return(enums.ErrorNotFoundRecords)
		err := NewWebhookController(repoMock).Remove(uuid.New())
		assert.Error(t, err)
	})
	t.Run("Should remove repository with error unexpected", func(t *testing.T) {
		repoMock := &repositoryWebhook.Mock{}
		repoMock.On("Remove").Return(errors.New("unexpected error"))
		err := NewWebhookController(repoMock).Remove(uuid.New())
		assert.Error(t, err)
	})
}
