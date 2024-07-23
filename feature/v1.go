package feature

import (
	"net/http"

	_ "github.com/cesc1802/onboarding-and-volunteer-service/docs"
	authStorage "github.com/cesc1802/onboarding-and-volunteer-service/feature/authentication/storage"
	authTransport "github.com/cesc1802/onboarding-and-volunteer-service/feature/authentication/transport"
	authUsecase "github.com/cesc1802/onboarding-and-volunteer-service/feature/authentication/usecase"
	countryStorage "github.com/cesc1802/onboarding-and-volunteer-service/feature/country/storage"
	countryTransport "github.com/cesc1802/onboarding-and-volunteer-service/feature/country/transport"
	countryUsecase "github.com/cesc1802/onboarding-and-volunteer-service/feature/country/usecase"
	deptStorage "github.com/cesc1802/onboarding-and-volunteer-service/feature/department/storage"
	deptTransport "github.com/cesc1802/onboarding-and-volunteer-service/feature/department/transport"
	deptUsecase "github.com/cesc1802/onboarding-and-volunteer-service/feature/department/usecase"
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/middleware"
	roleStorage "github.com/cesc1802/onboarding-and-volunteer-service/feature/role/storage"
	userStorage "github.com/cesc1802/onboarding-and-volunteer-service/feature/user/storage"
	userTransport "github.com/cesc1802/onboarding-and-volunteer-service/feature/user/transport"
	userUsecase "github.com/cesc1802/onboarding-and-volunteer-service/feature/user/usecase"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	appliIdentityStorage "github.com/cesc1802/onboarding-and-volunteer-service/feature/user_identity/storage"
	appliIdentityTransport "github.com/cesc1802/onboarding-and-volunteer-service/feature/user_identity/transport"
	appliIdentityUsecase "github.com/cesc1802/onboarding-and-volunteer-service/feature/user_identity/usecase"

	roleTransport "github.com/cesc1802/onboarding-and-volunteer-service/feature/role/transport"
	roleUsecase "github.com/cesc1802/onboarding-and-volunteer-service/feature/role/usecase"
	volunteerStorage "github.com/cesc1802/onboarding-and-volunteer-service/feature/volunteer/storage"
	volunteerTransport "github.com/cesc1802/onboarding-and-volunteer-service/feature/volunteer/transport"
	volunteerUsecase "github.com/cesc1802/onboarding-and-volunteer-service/feature/volunteer/usecase"

	"github.com/cesc1802/share-module/system"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// @host localhost:8080
// @BasePath /api/v1
func RegisterHandlerV1(mono system.Service) {
	router := mono.Router()
	secretKey := authStorage.GetSecretKey()
	router.Use(cors.Default())
	// add swagger
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]any{
			"data": "success",
		})
	})
	v1 := router.Group("/api/v1")
	// Initialize repository
	authRepo := authStorage.NewAuthenticationRepository(mono.DB())
	userRepo := userStorage.NewAdminRepository(mono.DB())
	applicantRepo := userStorage.NewApplicantRepository(mono.DB())
	applicantRequestRepo := userStorage.NewApplicantRequestRepository(mono.DB())
	applicantIdentityRepo := appliIdentityStorage.NewUserIdentityRepository(mono.DB())
	volunteerRepo := volunteerStorage.NewVolunteerRepository(mono.DB())
	volunteerRequestRepo := userStorage.NewVolunteerRequestRepository(mono.DB())
	roleRepo := roleStorage.NewRoleRepository(mono.DB())
	deptRepo := deptStorage.NewDepartmentRepository(mono.DB())
	countryRepo := countryStorage.NewCountryRepository(mono.DB())
	// Initialize usecase
	authUseCase := authUsecase.NewUserUsecase(authRepo, secretKey)
	userUseCase := userUsecase.NewAdminUsecase(userRepo)
	applicantUseCase := userUsecase.NewApplicantUsecase(applicantRepo)
	applicantRequestUseCase := userUsecase.NewApplicantRequestUsecase(applicantRequestRepo)
	applicantIdenityUseCase := appliIdentityUsecase.NewUserIdentityUsecase(applicantIdentityRepo)
	volunteerUseCase := volunteerUsecase.NewVolunteerUsecase(volunteerRepo)
	volunteerRequestUseCase := userUsecase.NewVolunteerRequestUsecase(volunteerRequestRepo)
	roleUseCase := roleUsecase.NewRoleUsecase(roleRepo)
	deptUseCase := deptUsecase.NewDepartmentUsecase(deptRepo)
	countryUseCase := countryUsecase.NewCountryUsecase(countryRepo)
	// Initialize handler
	authHandler := authTransport.NewAuthenticationHandler(authUseCase)
	userHandler := userTransport.NewAuthenticationHandler(userUseCase)
	applicantHandler := userTransport.NewApplicantHandler(applicantUseCase)
	applicantRequestHandler := userTransport.NewApplicantRequestHandler(applicantRequestUseCase)
	applicantIdentityHandler := appliIdentityTransport.NewUserIdentityHandler(applicantIdenityUseCase)
	volunteerHandler := volunteerTransport.NewVolunteerHandler(volunteerUseCase)
	volunteerRequestHandler := userTransport.NewVolunteerRequestHandler(volunteerRequestUseCase)
	roleHandler := roleTransport.NewRoleHandler(roleUseCase)
	deptHandler := deptTransport.NewDepartmentHandler(deptUseCase)
	countryHandler := countryTransport.NewCountryHandler(countryUseCase)
	auth := v1.Group("/auth")
	{
		auth.POST("/login", authHandler.Login)

		auth.POST("/register", authHandler.Register)
	}

	admin := v1.Group("/admin")
	admin.Use(middleware.AuthMiddleware(secretKey))
	{
		admin.GET("/list-request", userHandler.GetListRequest)
		admin.GET("/request/:id", userHandler.GetRequestById)
		admin.GET("/list-pending-request", userHandler.GetListPendingRequest)
		admin.GET("/pending-request/:id", userHandler.GetPendingRequestById)
		admin.POST("/approve-request/:id", userHandler.ApproveRequest)
		admin.POST("/reject-request/:id", userHandler.RejectRequest)
		admin.POST("/add-reject-notes/:id", userHandler.AddRejectNotes)
		admin.DELETE("/delete-request/:id", userHandler.DeleteRequest)
	}

	applicant := v1.Group("/applicant")
	{
		applicant.POST("/", applicantHandler.CreateApplicant)
		applicant.PUT("/:id", applicantHandler.UpdateApplicant)
		applicant.DELETE("/:id", applicantHandler.DeleteApplicant)
		applicant.GET("/:id", applicantHandler.FindApplicantByID)
	}

	appliRequest := v1.Group("/applicant-request")
	{
		appliRequest.POST("/", applicantRequestHandler.CreateApplicantRequest)
	}

	appliIdentity := v1.Group("applicant-identity")
	{
		appliIdentity.POST("/", applicantIdentityHandler.CreateUserIdentity)
		appliIdentity.GET("/:id", applicantIdentityHandler.FindUserIdentity)
		appliIdentity.PUT("/:id", applicantIdentityHandler.UpdateUserIdentity)
	}

	volunteer := v1.Group("/volunteer")
	{
		volunteer.POST("/", volunteerHandler.CreateVolunteer)
		volunteer.PUT("/:id", volunteerHandler.UpdateVolunteer)
		volunteer.DELETE("/:id", volunteerHandler.DeleteVolunteer)
		volunteer.GET("/:id", volunteerHandler.FindVolunteerByID)
		volunteer.GET("/", volunteerHandler.GetAllVolunteers)
	}

	volRequest := v1.Group("/volunteer-request")
	{
		volRequest.POST("/", volunteerRequestHandler.CreateVolunteerRequest)
	}

	role := v1.Group("/role")
	{
		role.POST("/", roleHandler.CreateRole)
		role.PUT("/:id", roleHandler.UpdateRole)
		role.DELETE("/:id", roleHandler.DeleteRole)
		role.GET("/:id", roleHandler.GetRoleByID)
		role.GET("/", roleHandler.GetAllRoles)
	}

	dept := v1.Group("/departments")
	{
		dept.POST("/", deptHandler.CreateDepartment)
		dept.PUT("/:id", deptHandler.UpdateDepartment)
		dept.DELETE("/:id", deptHandler.DeleteDepartment)
		dept.GET("/:id", deptHandler.GetDepartmentByID)
		dept.GET("/", deptHandler.GetAllDepartments)
	}

	country := v1.Group("/countries")
	{
		country.POST("/", countryHandler.CreateCountry)
		country.PUT("/:id", countryHandler.UpdateCountry)
		country.DELETE("/:id", countryHandler.DeleteCountry)
		country.GET("/:id", countryHandler.GetCountryByID)
		country.GET("/", countryHandler.GetAllCountries)
	}
}
