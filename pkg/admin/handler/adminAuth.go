package handler

import (
    "context"
    "fmt"
    "log"
    "net/http"

    "github.com/Shakezidin/middleware"
    dto "github.com/Shakezidin/pkg/DTO"
    pb "github.com/Shakezidin/pkg/admin/pb"
    "github.com/gin-gonic/gin"
    "github.com/go-playground/validator/v10"
)

func AdminLoginHandler(ctx *gin.Context, client pb.AdminClient, role string) {
    var login dto.Login

    if err := ctx.BindJSON(&login); err != nil {
        errMsg := "error binding JSON"
        log.Println(errMsg, err)
        ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
            "status": http.StatusBadRequest,
            "error":  errMsg,
        })
        return
    }

    validate := validator.New()

    if err := validate.Struct(login); err != nil {
        errMsg := "validation error"
        log.Println(errMsg, err)
        ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
            "status": http.StatusBadRequest,
            "error":  errMsg,
        })
        return
    }

    ctxt := context.Background()
    response, err := client.AdminLoginRequest(ctxt, &pb.AdminLogin{
        Email:    login.Email,
        Password: login.Password,
        Role:     role,
    })

    if err != nil {
        errMsg := "error logging in user"
        log.Printf("%s %s: %v", errMsg, login.Email, err)
        ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
            "status": http.StatusInternalServerError,
            "error":  errMsg,
        })
        return
    }

    ctx.JSON(http.StatusOK, gin.H{
        "status":  http.StatusOK,
        "message": fmt.Sprintf("%s logged in successfully", login.Email),
        "data":    response,
    })
}

func ViewDashboard(ctx *gin.Context, client pb.AdminClient) {
    email, _, err := middleware.ValidateToken(ctx, "admin")
    if err != nil {
        errMsg := "error validating token"
        log.Println(errMsg, err)
        ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
            "status": http.StatusUnauthorized,
            "error":  errMsg,
        })
        return
    }

    response, err := client.AdminViewDashBord(ctx, &pb.AdminView{Status: email})
    if err != nil {
        errMsg := "error fetching dashboard"
        log.Printf("%s: %v", errMsg, err)
        ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
            "status": http.StatusInternalServerError,
            "error":  errMsg,
        })
        return
    }

    ctx.JSON(http.StatusOK, gin.H{
        "status": http.StatusOK,
        "data":   response,
    })
}
