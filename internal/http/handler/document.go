package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/SicParv1sMagna/AtomHackMarsService/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (h *Handler) GetDocuments(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upgrade to websocket connection"})
		return
	}
	defer conn.Close()

	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	statusStr := c.Query("status")

	var status model.Status
	switch statusStr {
	case "DRAFT":
		status = model.StatusDraft
	case "FORMED":
		status = model.StatusFormed
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid status"})
		return
	}

	for {
		documents, err := h.r.GetDocuments(page, pageSize, status)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve documents"})
			return
		}

		if err := conn.WriteJSON(documents); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send documents"})
			return
		}
	}
}

// CreateDocument создает новый документ.
// @Summary Создает новый документ.
// @Description Создает новый документ на основе переданных данных JSON.
// @Tags Документы
// @Accept json
// @Produce json
// @Success 200 {object} model.MessageResponse "Успешный ответ"
// @Failure 400 {object} model.ErrorResponse "Ошибка в запросе"
// @Failure 500 {object} model.ErrorResponse "Внутренняя ошибка сервера"
// @Router /document [post]
func (h *Handler) CreateDocument(c *gin.Context) {
	doc := model.Document{
		Status:    model.StatusDraft,
		CreatedAt: time.Now(),
	}

	if err := h.r.CreateDocument(&doc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Document created successfully",
	})
}

// GetDocumentByID получает документ по его ID.
// @Summary Получает документ по ID.
// @Description Получает документ из репозитория по указанному ID.
// @Tags Документы
// @Accept json
// @Produce json
// @Param docID path int true "ID документа"
// @Success 200 {object} model.Document "Успешный ответ"
// @Failure 400 {object} model.ErrorResponse "Ошибка в запросе"
// @Failure 500 {object} model.ErrorResponse "Внутренняя ошибка сервера"
// @Router /document/{docID} [get]
func (h *Handler) GetDocumentByID(c *gin.Context) {
	docID, err := strconv.Atoi(c.Param("docID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get file from request"})
		return
	}

	doc, err := h.r.GetDocumentByID(uint(docID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"doc": doc})
}

// UpdateDocument обновляет информацию о документе.
// @Summary Обновляет информацию о документе.
// @Description Обновляет информацию о документе на основе переданных данных JSON.
// @Tags Документы
// @Accept json
// @Produce json
// @Param docID path int true "ID документа"
// @Param doc body model.DocumentUpdate true "Пользовательский объект в формате JSON"
// @Success 200 {object} model.MessageResponse "Успешный ответ"
// @Failure 400 {object} model.ErrorResponse "Ошибка в запросе"
// @Failure 500 {object} model.ErrorResponse "Внутренняя ошибка сервера"
// @Router /document/{docID} [put]
func (h *Handler) UpdateDocument(c *gin.Context) {
	docID, err := strconv.Atoi(c.Param("docID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get file from request"})
		return
	}

	var doc model.Document

	if err := c.BindJSON(&doc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to bind JSON"})
		return
	}

	if err := h.r.UpdateDocument(uint(docID), &doc); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Document updated successfully",
		"doc":     doc,
	})
}

// DeleteDocument удаляет документ по его ID.
// @Summary Удаляет документ по ID.
// @Description Удаляет документ из репозитория по указанному ID.
// @Tags Документы
// @Accept json
// @Produce json
// @Param docID path int true "ID документа"
// @Success 200 {object} model.MessageResponse "Успешный ответ"
// @Failure 400 {object} model.ErrorResponse "Ошибка в запросе"
// @Failure 500 {object} model.ErrorResponse "Внутренняя ошибка сервера"
// @Router /document/{docID} [delete]
func (h *Handler) DeleteDocument(c *gin.Context) {
	docID, err := strconv.Atoi(c.Param("docID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get document ID from request"})
		return
	}

	if err := h.r.DeleteDocument(uint(docID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete document: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Document deleted successfully"})
}