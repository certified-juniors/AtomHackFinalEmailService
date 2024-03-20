package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// SendToSupport отпрвялет на почту поддержки.
// @Summary Принимает сообщение (обращение).
// @Description Принимает новое обращение с параметрами mail, title, message, createdAt и files.
// @Tags EmailService
// @Accept json
// @Produce json
// @Param mail formData string true "Email пользователя"
// @Param title formData string true "Заголовок обращения"
// @Param message formData string true "Сообщение"
// @Param createdAt formData string true "Дата и время создания обращения в формате RFC3339"
// @Param files formData file true "Файлы, прикрепленные к обращению"
// @Success 200 {object} model.MessageResponse "Успешный ответ"
// @Failure 400 {object} model.ErrorResponse "Ошибка в запросе"
// @Failure 500 {object} model.ErrorResponse "Внутренняя ошибка сервера"
// @Router /send-to-support [post]
// SendToSupport принимает новый документ и отправляет его на поддержку.
func (h *Handler) SendToSupport(c *gin.Context) {
	email := c.PostForm("mail")
	title := c.PostForm("title")
	message := c.PostForm("message")
	createdAt := c.PostForm("timestamp")

	createdTime, err := time.Parse(time.RFC3339, createdAt)
	if err != nil {
		logrus.Errorf("Failed to parse timestamp: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse timestamp"})
		return
	}

	createdAtFormatted := createdTime.Format("02.01.2006 15:04:05")

	form, err := c.MultipartForm()
	if err != nil {
		logrus.Errorf("Failed to parse form data: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse form data"})
		return
	}
	files := form.File["files"]

	attachments := make(map[string][]byte)
	for _, file := range files {
		src, err := file.Open()
		if err != nil {
			logrus.Errorf("Failed to open file %s: %v", file.Filename, err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to open file: %s", file.Filename)})
			return
		}
		defer src.Close()

		content, err := ioutil.ReadAll(src)
		if err != nil {
			logrus.Errorf("Failed to read file %s: %v", file.Filename, err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to read file: %s", file.Filename)})
			return
		}

		attachments[file.Filename] = content
	}

	body := fmt.Sprintf("%s%s\n%s\n\n%s%s", "Email отправителя: ", email, message, "------------", "Дата обращения: ", createdAtFormatted)
	err = h.s.SendMailToSupport(title, body, attachments)
	if err != nil {
		logrus.Errorf("Failed to send SupportEmail: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to send SupportEmail: %v", err)})
		return
	}
	body = fmt.Sprintf("%s\n\n%s\n\n%s%s\n\nЭто сообщение было успешно отправлено в службу поддержки.", message, "------------", "Дата обращения: ", createdAtFormatted)
	err = h.s.SendMailToClient(title, body, attachments, email)
	if err != nil {
		logrus.Errorf("Failed to send ClientEmail: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to send ClientEmail: %v", err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Email sent successfully"})
}
