package messageservice

import "gorm.io/gorm"

type MessageRepository interface {
	CreateMessage(message Message) (Message, error)
	GetAllMessages() ([]Message, error)
	UpdateMessageByID(id uint, message Message) (Message, error)
	DeleteMessageByID(id uint) error
}

type messageRepository struct {
	db *gorm.DB
}

func NewMessageRepository(db *gorm.DB) *messageRepository {
	return &messageRepository{db: db}
}

func (r *messageRepository) CreateMessage(message Message) (Message, error) {
	result := r.db.Create(&message)
	if result.Error != nil {
		return Message{}, result.Error
	}
	return message, nil
}

func (r *messageRepository) GetAllMessages() ([]Message, error) {
	var Messages []Message
	err := r.db.Find(&Messages).Error
	return Messages, err
}

func (r *messageRepository) UpdateMessageByID(id uint, message Message) (Message, error) {
	err := r.db.Model(&Message{}).Where("id = ?", id).Update("is_done", message.IsDone).Error
	if err != nil {
		return Message{}, err
	}

	var result Message
	err = r.db.First(&result, id).Error
	return result, err
}

func (r *messageRepository) DeleteMessageByID(id uint) error {
	err := r.db.Unscoped().Delete(&Message{}, id).Error
	return err
}
