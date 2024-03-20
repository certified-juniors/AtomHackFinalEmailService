package handler

import (
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// AcceptDocument принимает новый документ.
// @Summary Принимает новый документ.
// @Description Принимает новый документ с параметрами id, title, owner, createdAt, payload и files.
// @Tags Документы
// @Accept json
// @Produce json
// @Param id formData int true "ID документа"
// @Param title formData string true "Заголовок документа"
// @Param owner formData string true "Владелец документа"
// @Param createdAt formData string true "Дата и время создания документа в формате RFC3339"
// @Param payload formData string true "Payload документа"
// @Param files formData file true "Файлы, прикрепленные к документу"
// @Success 200 {object} model.AcceptDocument "Успешный ответ"
// @Failure 400 {object} model.ErrorResponse "Ошибка в запросе"
// @Failure 500 {object} model.ErrorResponse "Внутренняя ошибка сервера"
// @Router /document/send-to-support [post]
// SendToSupport принимает новый документ и отправляет его на поддержку.
func (h *Handler) SendToSupport(c *gin.Context) {
    // Получение данных о документе из POST-запроса
    id, err := strconv.Atoi(c.PostForm("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get id from request"})
        return
    }
    title := c.PostForm("title")
    owner := c.PostForm("owner")
    createdAtStr := c.PostForm("createdAt")
    payload := c.PostForm("payload")

    // Получение файлов из POST-запроса
    form, err := c.MultipartForm()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse form data"})
        return
    }
    files := form.File["files"]

    // Создание карты вложений для метода SendMail
    attachments := make(map[string][]byte)
    for _, file := range files {
        // Открытие файла
        src, err := file.Open()
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open file: " + file.Filename})
            return
        }
        defer src.Close()

        // Считывание содержимого файла
        content, err := ioutil.ReadAll(src)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read file: " + file.Filename})
            return
        }

        // Добавление содержимого файла в карту вложений
        attachments[file.Filename] = content
    }

    // Отправка сообщения на почту
    err = h.s.SendMailToSupport("New Document Received", "ID: "+strconv.Itoa(id)+"\nTitle: "+title+"\nOwner: "+owner+"\nCreated At: "+createdAtStr+"\nPayload: "+payload, attachments)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send email: " + err.Error()})
        return
    }

    // Отправляем ответ клиенту
    c.JSON(http.StatusOK, gin.H{"message": "Email sent successfully"})
}