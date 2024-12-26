package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/kaelCoding/toyBE/internal/models"
    "github.com/kaelCoding/toyBE/internal/utils"
    "github.com/dgrijalva/jwt-go"
    "gorm.io/gorm"
    "os"
    "time"
    "encoding/json"
    "errors"
    "log"
)

func RegisterUser(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        var newUser models.User
        if err := json.NewDecoder(c.Request.Body).Decode(&newUser); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
            return
        }

        // Hash the password
        hash, err := utils.GenerateFromPassword(newUser.Password, &utils.HashParams{ // Pass newUser.Password and HashParams directly
            Memory:      64 * 1024,
            Iterations:  3,
            Parallelism: 2,
            SaltLength:  16,
            KeyLength:   32,
        })
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        newUser.Password = hash

        // Create the user in the database
        result := db.Create(&newUser)
        if result.Error != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
            return
        }

        c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
    }
}

func LoginUser(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        var loginData models.Login
        if err := c.ShouldBindJSON(&loginData); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
            return
        }

        user := &models.User{}
        result := db.Where(&models.User{Username: loginData.Username}).First(user)
        if result.Error != nil {
            if errors.Is(result.Error, gorm.ErrRecordNotFound) {
                c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
            } else {
                c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"}) // Lỗi cơ sở dữ liệu chung
                // Có thể log lỗi chi tiết hơn ở đây để debug
                log.Println("Database error:", result.Error)
            }
            return
        }

        // Sử dụng hàm VerifyPassword từ utils
        match, err := loginData.VerifyPassword(user.Password)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Error verifying password"})
            return
        }

        if !match {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
            return
        }

        // Nếu xác thực thành công, tạo token và trả về
        claims := &models.CustomJWTClaims{
            Username: user.Username,
            Email: user.Email,
            ID:   user.ID,
            StandardClaims: jwt.StandardClaims{
                ExpiresAt: time.Now().Add(time.Minute * 30).Unix(),
            },
        }

        token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

        jwtSecret := os.Getenv("JWT_SECRET")
        tokenString, err := token.SignedString([]byte(jwtSecret))
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating token"})
            return
        }

        c.JSON(http.StatusOK, gin.H{"token": tokenString})
    }
}