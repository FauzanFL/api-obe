package main

import (
	"api-obe/controller"
	"api-obe/db"
	"api-obe/middleware"
	repo "api-obe/repository"
	"fmt"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dbConfigs := []db.DBConfig{
		{
			IdentificationName: os.Getenv("DB_KURIKULUM_IDENTIFICATION_NAME"),
			DB:                 os.Getenv("DB_KURIKULUM_NAME"),
			User:               os.Getenv("DB_KURIKULUM_USER"),
			Password:           os.Getenv("DB_KURIKULUM_PASSWORD"),
			Host:               os.Getenv("DB_KURIKULUM_HOST"),
			Port:               os.Getenv("DB_KURIKULUM_PORT"),
			TimeZone:           os.Getenv("DB_KURIKULUM_TIMEZONE"),
			Type:               os.Getenv("DB_KURIKULUM_TYPE"),
		},
		{
			IdentificationName: os.Getenv("DB_USER_IDENTIFICATION_NAME"),
			DB:                 os.Getenv("DB_USER_NAME"),
			User:               os.Getenv("DB_USER_USER"),
			Password:           os.Getenv("DB_USER_PASSWORD"),
			Host:               os.Getenv("DB_USER_HOST"),
			Port:               os.Getenv("DB_USER_PORT"),
			TimeZone:           os.Getenv("DB_USER_TIMEZONE"),
			Type:               os.Getenv("DB_USER_TYPE"),
		},
		{
			IdentificationName: os.Getenv("DB_PENILAIAN_IDENTIFICATION_NAME"),
			DB:                 os.Getenv("DB_PENILAIAN_NAME"),
			User:               os.Getenv("DB_PENILAIAN_USER"),
			Password:           os.Getenv("DB_PENILAIAN_PASSWORD"),
			Host:               os.Getenv("DB_PENILAIAN_HOST"),
			Port:               os.Getenv("DB_PENILAIAN_PORT"),
			TimeZone:           os.Getenv("DB_PENILAIAN_TIMEZONE"),
			Type:               os.Getenv("DB_PENILAIAN_TYPE"),
		},
	}
	fmt.Println("Start connection")
	db.InitDBConnections(dbConfigs)
}

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "User-Agent", "Referrer", "Host", "Token", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowMethods:     []string{"PUT", "POST", "GET", "DELETE", "OPTIONS"},
		AllowCredentials: true,
		AllowOrigins:     []string{"http://localhost:5173", "http://localhost:5173/"},
	}))

	dbUser := db.GetDBConnection(os.Getenv("DB_USER_IDENTIFICATION_NAME"))
	dbKurikulum := db.GetDBConnection(os.Getenv("DB_KURIKULUM_IDENTIFICATION_NAME"))
	dbPenilaian := db.GetDBConnection(os.Getenv("DB_PENILAIAN_IDENTIFICATION_NAME"))

	perancanganObeRepo := repo.NewPerancanganObeRepo(dbKurikulum)
	ploRepo := repo.NewPloRepo(dbKurikulum)
	cloRepo := repo.NewCloRepo(dbKurikulum)
	kurikulumRepo := repo.NewKurikulumRepo(dbKurikulum)
	mataKuliahRepo := repo.NewMataKuliahRepo(dbKurikulum)
	plottingDosenRepo := repo.NewPlottingDosenMkRepo(dbKurikulum)
	lembarAssessmentRepo := repo.NewLembarAssessmentRepo(dbKurikulum)
	jenisAssessmentRepo := repo.NewJenisAssessmentRepo(dbKurikulum)
	tahunAjaranRepo := repo.NewTahunAjaranRepo(dbKurikulum)
	userRepo := repo.NewUserRepo(dbUser)
	dosenRepo := repo.NewDosenRepo(dbUser)
	penilaianRepo := repo.NewPenilaianRepo(dbPenilaian)
	beritaAcaraRepo := repo.NewBeritaAcaraRepo(dbPenilaian)
	kelasRepo := repo.NewKelasRepo(dbPenilaian)
	indexPenilaianRepo := repo.NewIndexPenilaianRepo(dbPenilaian)

	authMiddleware := middleware.NewAuthMiddleware(userRepo)

	userController := controller.NewUserController(userRepo, dosenRepo)
	perancanganObeController := controller.NewPerancanganObeController(perancanganObeRepo)
	ploController := controller.NewPloController(ploRepo)
	cloController := controller.NewCloController(cloRepo)
	kurikulumController := controller.NewKurikulumController(kurikulumRepo)
	mataKuliahController := controller.NewMataKuliahController(mataKuliahRepo, plottingDosenRepo, dosenRepo, lembarAssessmentRepo, perancanganObeRepo)
	plottingDosenController := controller.NewPlottingDosenMkController(plottingDosenRepo, perancanganObeRepo, dosenRepo, kelasRepo, mataKuliahRepo)
	lembarAssessmentController := controller.NewLembarAssessmentController(lembarAssessmentRepo, cloRepo)
	jenisAssessmentController := controller.NewJenisAssessmentController(jenisAssessmentRepo)
	dosenController := controller.NewDosenController(dosenRepo, mataKuliahRepo, perancanganObeRepo)
	penilaianController := controller.NewPenilaianController(penilaianRepo, cloRepo, lembarAssessmentRepo, dosenRepo, ploRepo, perancanganObeRepo)
	tahunAjaranController := controller.NewTahunAjaranController(tahunAjaranRepo)
	beritaAcaraController := controller.NewBeritaAcaraController(beritaAcaraRepo, dosenRepo, tahunAjaranRepo, lembarAssessmentRepo, penilaianRepo)
	kelasController := controller.NewKelasController(kelasRepo)
	indexPenilaianController := controller.NewIndexPenilaianController(indexPenilaianRepo)

	apiRouter := r.Group("/api")
	{
		userRouter := apiRouter.Group("/users")
		{
			userRouter.GET("/", authMiddleware.RequireAdminAuth, userController.GetUser)
			userRouter.GET("/role", authMiddleware.RequireAuth, userController.GetUserRole)
			userRouter.GET("/dosen", authMiddleware.RequireAdminAuth, userController.GetUserDosen)
			userRouter.GET("/dosen/search", authMiddleware.RequireAdminAuth, userController.SearchUserDosen)
			userRouter.POST("/", authMiddleware.RequireAdminAuth, userController.AddUser)
			userRouter.DELETE("/delete/:id", authMiddleware.RequireAdminAuth, userController.DeleteUser)
			userRouter.PUT("/update/:id", authMiddleware.RequireAdminAuth, userController.UpdateUser)
			userRouter.POST("/login", userController.Login)
			userRouter.GET("/logout", authMiddleware.RequireAuth, userController.Logout)
		}

		dosenRouter := apiRouter.Group("/dosen")
		{
			dosenRouter.GET("/", authMiddleware.RequireAdminAuth, dosenController.GetDosen)
			dosenRouter.GET("/mata_kuliah", authMiddleware.RequireDosenAuth, dosenController.GetMataKuliah)
			dosenRouter.GET("/mata_kuliah/search", authMiddleware.RequireDosenAuth, dosenController.SearchMataKuliah)
		}

		perancanganObeRouter := apiRouter.Group("/perancangan_obe")
		{
			perancanganObeRouter.GET("/", authMiddleware.RequireAdminAuth, perancanganObeController.GetPerancanganObe)
			perancanganObeRouter.GET("/search", authMiddleware.RequireAdminAuth, perancanganObeController.SearchPerancanganObe)
			perancanganObeRouter.GET("/active", authMiddleware.RequireAdminAuth, perancanganObeController.GetActivePerancanganObe)
			perancanganObeRouter.PUT("/activate/:id", authMiddleware.RequireAdminAuth, perancanganObeController.ActivatePerancangan)
			perancanganObeRouter.GET("/:id", authMiddleware.RequireAdminAuth, perancanganObeController.GetPerancanganObeById)
			perancanganObeRouter.POST("/", authMiddleware.RequireAdminAuth, perancanganObeController.CreatePerancanganObe)
			perancanganObeRouter.DELETE("/delete/:id", authMiddleware.RequireAdminAuth, perancanganObeController.DeletePerancanganObe)
			perancanganObeRouter.PUT("/update/:id", authMiddleware.RequireAdminAuth, perancanganObeController.UpdatePerancanganObe)
		}

		ploRouter := apiRouter.Group("/plo")
		{
			ploRouter.GET("/", authMiddleware.RequireAuth, ploController.GetPlo)
			ploRouter.GET("/:id", authMiddleware.RequireAuth, ploController.GetPloById)
			ploRouter.GET("/obe/:obeId", authMiddleware.RequireAuth, ploController.GetPloByObeId)
			ploRouter.GET("/obe/:obeId/search", authMiddleware.RequireAuth, ploController.SearchPloByObeId)
			ploRouter.POST("/", authMiddleware.RequireAdminAuth, ploController.CreatePlo)
			ploRouter.DELETE("/delete/:id", authMiddleware.RequireAdminAuth, ploController.DeletePlo)
			ploRouter.PUT("/update/:id", authMiddleware.RequireAdminAuth, ploController.UpdatePlo)
		}

		cloRouter := apiRouter.Group("/clo")
		{
			cloRouter.GET("/", authMiddleware.RequireAuth, cloController.GetClo)
			cloRouter.GET("/:id", authMiddleware.RequireAuth, cloController.GetCloById)
			cloRouter.GET("/mk/:mkId", authMiddleware.RequireAuth, cloController.GetCloByMkId)
			cloRouter.POST("/", authMiddleware.RequireDosenAuth, cloController.CreateClo)
			cloRouter.DELETE("/delete/:id", authMiddleware.RequireDosenAuth, cloController.DeleteClo)
			cloRouter.PUT("/update/:id", authMiddleware.RequireDosenAuth, cloController.UpdateClo)
		}

		kurikulumRouter := apiRouter.Group("/kurikulum")
		{
			kurikulumRouter.GET("/", authMiddleware.RequireAdminAuth, kurikulumController.GetKurikulum)
		}

		mataKuliahRouter := apiRouter.Group("/mata_kuliah")
		{
			mataKuliahRouter.GET("/", authMiddleware.RequireAuth, mataKuliahController.GetMataKuliah)
			mataKuliahRouter.GET("/:id", authMiddleware.RequireAuth, mataKuliahController.GetMataKuliahById)
			mataKuliahRouter.GET("/tahun/:tahunId", authMiddleware.RequireAuth, mataKuliahController.GetMataKuliahActiveByTahunId)
			mataKuliahRouter.GET("/tahun/:tahunId/search", authMiddleware.RequireAuth, mataKuliahController.SearchMataKuliahActiveByTahunId)
			mataKuliahRouter.POST("/", authMiddleware.RequireAdminAuth, mataKuliahController.CreateMataKuliah)
			mataKuliahRouter.DELETE("/delete/:id", authMiddleware.RequireAdminAuth, mataKuliahController.DeleteMataKuliah)
			mataKuliahRouter.PUT("/update/:id", authMiddleware.RequireAdminAuth, mataKuliahController.UpdateMataKuliah)
			mataKuliahRouter.GET("/rps/:id", authMiddleware.RequireAuth, mataKuliahController.GetRPS)
		}

		plottingDosenMkRouter := apiRouter.Group("/plotting_dosen_mk")
		{
			plottingDosenMkRouter.GET("/tahun/:tahunId", authMiddleware.RequireAdminAuth, plottingDosenController.GetPlottingDosenMkByTahun)
			plottingDosenMkRouter.GET("/matakuliah/:mkId/kelas", authMiddleware.RequireAdminAuth, plottingDosenController.GetKelasMkByMk)
			plottingDosenMkRouter.GET("/matakuliah/:mkId/kelas/dosen", authMiddleware.RequireDosenAuth, plottingDosenController.GetKelasDosenByMk)
			plottingDosenMkRouter.GET("/tahun/:tahunId/search", authMiddleware.RequireAdminAuth, plottingDosenController.SearchPlottingDosenMkByTahun)
			plottingDosenMkRouter.POST("/", authMiddleware.RequireAdminAuth, plottingDosenController.CreatePlottingDosenMk)
			plottingDosenMkRouter.DELETE("/delete/:id", authMiddleware.RequireAdminAuth, plottingDosenController.DeletePlottingDosenMk)
		}

		lembarAssessmentRouter := apiRouter.Group("/lembar_assessment")
		{
			lembarAssessmentRouter.GET("/", authMiddleware.RequireAuth, lembarAssessmentController.GetLembarAssessment)
			lembarAssessmentRouter.GET("/:id", authMiddleware.RequireAuth, lembarAssessmentController.GetLembarAssessmentById)
			lembarAssessmentRouter.GET("/matakuliah/:mkId", authMiddleware.RequireAuth, lembarAssessmentController.GetLembarAssessmentByMkId)
			lembarAssessmentRouter.GET("/clo/:cloId", authMiddleware.RequireAuth, lembarAssessmentController.GetLembarAssessmentByCloId)
			lembarAssessmentRouter.GET("/clo/:cloId/search", authMiddleware.RequireAuth, lembarAssessmentController.SearchLembarAssessment)
			lembarAssessmentRouter.POST("/", authMiddleware.RequireAuth, lembarAssessmentController.CreateLembarAssessment)
			lembarAssessmentRouter.DELETE("/delete/:id", authMiddleware.RequireAuth, lembarAssessmentController.DeleteLembarAssessment)
			lembarAssessmentRouter.PUT("/update/:id", authMiddleware.RequireAuth, lembarAssessmentController.UpdateLembarAssessment)
		}

		jenisAssessmentRouter := apiRouter.Group("/jenis_assessment")
		{
			jenisAssessmentRouter.GET("/", authMiddleware.RequireAuth, jenisAssessmentController.GetJenisAssessment)
			jenisAssessmentRouter.POST("/", authMiddleware.RequireAuth, jenisAssessmentController.CreateJenisAssessment)
			jenisAssessmentRouter.PUT("/update/:id", authMiddleware.RequireAuth, jenisAssessmentController.UpdateJenisAssessment)
			jenisAssessmentRouter.DELETE("/delete/:id", authMiddleware.RequireAuth, jenisAssessmentController.DeleteJenisAssessment)
		}

		penilaianRouter := apiRouter.Group("/penilaian")
		{
			penilaianRouter.GET("/", authMiddleware.RequireAuth, penilaianController.GetPenilaian)
			penilaianRouter.GET("/:id", authMiddleware.RequireAuth, penilaianController.GetPenilaianById)
			penilaianRouter.GET("/data/matakuliah/:mkId/kelas/:kelasId", authMiddleware.RequireAuth, penilaianController.GetDataPenilaian)
			penilaianRouter.GET("/data/clo_plo/matakuliah/:mkId/tahun/:tahunId", authMiddleware.RequireDosenAuth, penilaianController.GetDataPenilaianCLOPLOByMk)
			penilaianRouter.GET("/data/plo/tahun/:tahunId", authMiddleware.RequireAuth, penilaianController.GetDataPenilaianPLO)
			penilaianRouter.POST("/", authMiddleware.RequireDosenAuth, penilaianController.CreatePenilaian)
			penilaianRouter.DELETE("/delete/:id", authMiddleware.RequireDosenAuth, penilaianController.DeletePenilaian)
			penilaianRouter.PUT("/update/:id", authMiddleware.RequireDosenAuth, penilaianController.UpdatePenilaian)
			penilaianRouter.POST("/upload", authMiddleware.RequireDosenAuth, penilaianController.UploadEvidence)
		}

		tahunAjaranRouter := apiRouter.Group("/tahun_ajaran")
		{
			tahunAjaranRouter.GET("/", authMiddleware.RequireAuth, tahunAjaranController.GetTahunAjaran)
			tahunAjaranRouter.GET("/:id", authMiddleware.RequireAuth, tahunAjaranController.GetTahunAjaranById)
			tahunAjaranRouter.GET("/now", authMiddleware.RequireAuth, tahunAjaranController.GetTahunAjaranNow)
			tahunAjaranRouter.GET("/search", authMiddleware.RequireAuth, tahunAjaranController.SearchTahunAjaran)
			tahunAjaranRouter.POST("/", authMiddleware.RequireAdminAuth, tahunAjaranController.CreateTahunAjaran)
			tahunAjaranRouter.PUT("/update/:id", authMiddleware.RequireAdminAuth, tahunAjaranController.UpdateTahunAjaran)
			tahunAjaranRouter.DELETE("/delete/:id", authMiddleware.RequireAdminAuth, tahunAjaranController.DeleteTahunAjaran)
		}

		beritaAcaraRouter := apiRouter.Group("/berita_acara")
		{
			beritaAcaraRouter.GET("/", authMiddleware.RequireAuth, beritaAcaraController.GetBeritaAcara)
			beritaAcaraRouter.GET("/penilaian/:penilaianId", authMiddleware.RequireAuth, beritaAcaraController.GetBeritaAcaraByPenilaian)
			beritaAcaraRouter.POST("/", authMiddleware.RequireDosenAuth, beritaAcaraController.CreateBeritaAcara)
			beritaAcaraRouter.DELETE("/delete/:id", authMiddleware.RequireAuth, beritaAcaraController.DeleteBeritaAcara)
		}

		kelasRouter := apiRouter.Group("/kelas")
		{
			kelasRouter.GET("/", authMiddleware.RequireAuth, kelasController.GetKelas)
			kelasRouter.GET("/:kelasId", authMiddleware.RequireAuth, kelasController.GetKelasById)
			kelasRouter.GET("/search", authMiddleware.RequireAuth, kelasController.SearchKelas)
			kelasRouter.POST("/", authMiddleware.RequireAdminAuth, kelasController.CreateKelas)
			kelasRouter.PUT("/update/:id", authMiddleware.RequireAdminAuth, kelasController.UpdateKelas)
			kelasRouter.DELETE("/delete/:id", authMiddleware.RequireAdminAuth, kelasController.DeleteKelas)
		}

		indexPenilaianRouter := apiRouter.Group("/index_penilaian")
		{
			indexPenilaianRouter.GET("/", authMiddleware.RequireAuth, indexPenilaianController.GetIndexPenilaian)
			indexPenilaianRouter.GET("/:id", authMiddleware.RequireAuth, indexPenilaianController.GetIndexPenilaianById)
			indexPenilaianRouter.GET("/grade", authMiddleware.RequireAuth, indexPenilaianController.GetIndexPenilaianByNilai)
			indexPenilaianRouter.POST("/", authMiddleware.RequireAdminAuth, indexPenilaianController.CreateIndexPenilaian)
			indexPenilaianRouter.PUT("/update/:id", authMiddleware.RequireAdminAuth, indexPenilaianController.UpdateIndexPenilaian)
			indexPenilaianRouter.DELETE("/delete/:id", authMiddleware.RequireAdminAuth, indexPenilaianController.DeleteIndexPenilaian)
		}
	}

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, "Hello World")
	})
	r.Run()
}
