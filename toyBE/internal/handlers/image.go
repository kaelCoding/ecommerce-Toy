package handlers

import (
    "os"
    "net/http"
	"io"
    "path/filepath"
    "errors"
    "fmt"
    "mime/multipart"
    "gorm.io/gorm"


    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
	"github.com/kaelCoding/toyBE/internal/models"
	"github.com/kaelCoding/toyBE/internal/database"
)

const (
    maxUploadSize = 5 * 1024 * 1024 // 5MB (điều chỉnh nếu cần)
    uploadPath    = "uploads"       // Đường dẫn thư mục uploads
)

func UploadImage(c *gin.Context) {
    db := database.GetDB()
    if c.Request.ContentLength > maxUploadSize {
        c.JSON(http.StatusRequestEntityTooLarge, gin.H{"error": "Kích thước tệp vượt quá giới hạn 5MB"})
        return
    }

    form, err := c.MultipartForm()
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Không thể đọc dữ liệu form"})
        return
    }

    files := form.File["image"] 
    var uploadedImages []models.Image
    var uploadErrors []string

    for _, file := range files {
        if err := validateFile(file); err != nil {
            uploadErrors = append(uploadErrors, err.Error())
            continue 
        }

        image, err := uploadAndSaveImage(db, file)
        if err != nil {
            uploadErrors = append(uploadErrors, fmt.Sprintf("Lỗi khi tải lên %s: %v", file.Filename, err))
            continue
        }
        uploadedImages = append(uploadedImages, *image)
    }

    if len(uploadErrors) > 0 {
        c.JSON(http.StatusInternalServerError, gin.H{"errors": uploadErrors})
        return
    }

    c.JSON(http.StatusCreated, uploadedImages)
}

func validateFile(file *multipart.FileHeader) error {
    if file.Size > maxUploadSize {
        return errors.New("kích thước tệp vượt quá giới hạn")
    }

    contentType := file.Header.Get("Content-Type")
    switch contentType {
    case "image/jpeg", "image/png", "image/gif", "image/webp":
    default:
        return errors.New("loại tệp không được hỗ trợ")
    }

    return nil
}

func uploadAndSaveImage(db *gorm.DB, file *multipart.FileHeader) (*models.Image, error) {
    fileName := file.Filename
    ext := filepath.Ext(fileName)
    newFileName := uuid.New().String() + ext
    destination := filepath.Join(uploadPath, newFileName) 

    if err := os.MkdirAll(uploadPath, os.ModePerm); err != nil && !os.IsExist(err) {
        return nil, fmt.Errorf("không thể tạo thư mục uploads: %w", err)
    }

    openedFile, err := file.Open()
    if err != nil {
        return nil, fmt.Errorf("failed to open file: %w", err)
    }
    defer openedFile.Close()

    out, err := os.Create(destination)
    if err != nil {
        return nil, fmt.Errorf("không thể tạo tệp: %w", err)
    }
    defer out.Close()

    _, err = io.Copy(out, openedFile)
    if err != nil {
        return nil, fmt.Errorf("không thể sao chép dữ liệu tệp: %w", err)
    }

    fileInfo, err := os.Stat(destination)
    if err != nil {
        return nil, fmt.Errorf("không thể lấy thông tin tệp: %w", err)
    }

    fileSize := fileInfo.Size()
    contentType := file.Header.Get("Content-Type")

    image := models.Image{
        ID:        uuid.New(),
        Name:      fileName,
        Type:      contentType,
        Size:      fileSize,
        Link:      destination,
    }

    if err := db.Create(&image).Error; err != nil {
        os.Remove(destination)
        return nil, fmt.Errorf("không thể lưu ảnh vào database: %w", err)
    }

    return &image, nil
}