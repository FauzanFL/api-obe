package main

import (
	"api-obe/controller"
	"api-obe/db"
	"api-obe/middleware"
	repo "api-obe/repository"
	"fmt"
	"log"
	"os"

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
	dbUser := db.GetDBConnection(os.Getenv("DB_USER_IDENTIFICATION_NAME"))
	dbKurikulum := db.GetDBConnection(os.Getenv("DB_KURIKULUM_IDENTIFICATION_NAME"))
	dbPenilaian := db.GetDBConnection(os.Getenv("DB_PENILAIAN_IDENTIFICATION_NAME"))
	perancanganObeRepo := repo.NewPerancanganObeRepo(dbKurikulum)
	ploRepo := repo.NewPloRepo(dbKurikulum)
	cloRepo := repo.NewCloRepo(dbKurikulum)
	ploCloRepo := repo.NewPloCloRepo(dbKurikulum)
	kurikulumRepo := repo.NewKurikulumRepo(dbKurikulum)
	mataKuliahRepo := repo.NewMataKuliahRepo(dbKurikulum)
	plottingDosenRepo := repo.NewPlottingDosenMkRepo(dbKurikulum)
	lembarAssessmentRepo := repo.NewLembarAssessmentRepo(dbKurikulum)
	jenisAssessmentRepo := repo.NewJenisAssessmentRepo(dbKurikulum)
	userRepo := repo.NewUserRepo(dbUser)
	dosenRepo := repo.NewDosenRepo(dbUser)
	mkMahasiswaRepo := repo.NewMkMahasiswaRepo(dbPenilaian)
	penilaianRepo := repo.NewPenilaianRepo(dbPenilaian)
	tahunAjaranRepo := repo.NewTahunAjaranRepo(dbPenilaian)
	beritaAcaraRepo := repo.NewBeritaAcaraRepo(dbPenilaian)
	kelasRepo := repo.NewKelasRepo(dbPenilaian)
	mahasiswaRepo := repo.NewMahasiswaRepo(dbPenilaian)

	authMiddleware := middleware.NewAuthMiddleware(userRepo)

	userController := controller.NewUserController(userRepo, dosenRepo)
	perancanganObeController := controller.NewPerancanganObeController(perancanganObeRepo)
	ploController := controller.NewPloController(ploRepo)
	cloController := controller.NewCloController(cloRepo)
	ploCloController := controller.NewPloCloController(ploCloRepo)
	kurikulumController := controller.NewKurikulumController(kurikulumRepo)
	mataKuliahController := controller.NewMataKuliahController(mataKuliahRepo, plottingDosenRepo, dosenRepo)
	plottingDosenController := controller.NewPlottingDosenMkController(plottingDosenRepo)
	lembarAssessmentController := controller.NewLembarAssessmentController(lembarAssessmentRepo)
	jenisAssessmentController := controller.NewJenisAssessmentController(jenisAssessmentRepo)
	dosenController := controller.NewDosenController(dosenRepo, mataKuliahRepo)
	mkMahasiswaController := controller.NewMkMahasiswaController(mkMahasiswaRepo)
	penilaianController := controller.NewPenilaianController(penilaianRepo)
	tahunAjaranController := controller.NewTahunAjaranController(tahunAjaranRepo)
	beritaAcaraController := controller.NewBeritaAcaraController(beritaAcaraRepo)
	kelasController := controller.NewKelasController(kelasRepo)
	mahasiswaController := controller.NewMahasiswaController(mahasiswaRepo)

	userRouter := r.Group("/users")
	{
		userRouter.GET("/", authMiddleware.RequireAuth, userController.GetUser)
		userRouter.POST("/", authMiddleware.RequireAuth, userController.AddUser)
		userRouter.DELETE("/delete/:id", authMiddleware.RequireAuth, userController.DeleteUser)
		userRouter.PUT("/update/:id", authMiddleware.RequireAuth, userController.UpdateUser)
		userRouter.POST("/login", userController.Login)
		userRouter.GET("/logout", authMiddleware.RequireAuth, userController.Logout)
	}

	dosenRouter := r.Group("/dosen")
	{
		dosenRouter.GET("/", authMiddleware.RequireAuth, dosenController.GetDosen)
		dosenRouter.GET("/:id/obe/:obeId", authMiddleware.RequireAuth, dosenController.GetMataKuliahObe)
	}

	perancanganObeRouter := r.Group("/perancangan_obe")
	{
		perancanganObeRouter.GET("/", authMiddleware.RequireAuth, perancanganObeController.GetPerancanganObe)
		perancanganObeRouter.GET("/:id", authMiddleware.RequireAuth, perancanganObeController.GetPerancanganObeById)
		perancanganObeRouter.POST("/", authMiddleware.RequireAuth, perancanganObeController.CreatePerancanganObe)
		perancanganObeRouter.DELETE("/delete/:id", authMiddleware.RequireAuth, perancanganObeController.DeletePerancanganObe)
		perancanganObeRouter.PUT("/update/:id", authMiddleware.RequireAuth, perancanganObeController.UpdatePerancanganObe)
	}

	ploRouter := r.Group("/plo")
	{
		ploRouter.GET("/", authMiddleware.RequireAuth, ploController.GetPlo)
		ploRouter.GET("/:id", authMiddleware.RequireAuth, ploController.GetPloById)
		ploRouter.POST("/", authMiddleware.RequireAuth, ploController.CreatePlo)
		ploRouter.DELETE("/delete/:id", authMiddleware.RequireAuth, ploController.DeletePlo)
		ploRouter.PUT("/update/:id", authMiddleware.RequireAuth, ploController.UpdatePlo)
	}

	cloRouter := r.Group("/clo")
	{
		cloRouter.GET("/", authMiddleware.RequireAuth, cloController.GetClo)
		cloRouter.GET("/:id", authMiddleware.RequireAuth, cloController.GetCloById)
		cloRouter.POST("/", authMiddleware.RequireAuth, cloController.CreateClo)
		cloRouter.DELETE("/delete/:id", authMiddleware.RequireAuth, cloController.DeleteClo)
		cloRouter.PUT("/update/:id", authMiddleware.RequireAuth, cloController.UpdateClo)
	}

	ploCloRouter := r.Group("/plo_clo")
	{
		ploCloRouter.GET("/", authMiddleware.RequireAuth, ploCloController.GetPloClo)
		ploCloRouter.POST("/", authMiddleware.RequireAuth, ploCloController.CreatePloClo)
		ploCloRouter.DELETE("/delete/:id", authMiddleware.RequireAuth, ploCloController.DeletePloClo)
	}

	kurikulumRouter := r.Group("/kurikulum")
	{
		kurikulumRouter.GET("/", authMiddleware.RequireAuth, kurikulumController.GetKurikulum)
	}

	mataKuliahRouter := r.Group("/mata_kuliah")
	{
		mataKuliahRouter.GET("/", authMiddleware.RequireAuth, mataKuliahController.GetMataKuliah)
		mataKuliahRouter.GET("/:id", authMiddleware.RequireAuth, mataKuliahController.GetMataKuliahById)
		mataKuliahRouter.GET("/obe/:obeId", authMiddleware.RequireAuth, mataKuliahController.GetMataKuliahByObeId)
		mataKuliahRouter.POST("/", authMiddleware.RequireAuth, mataKuliahController.CreateMataKuliah)
		mataKuliahRouter.DELETE("/delete/:id", authMiddleware.RequireAuth, mataKuliahController.DeleteMataKuliah)
		mataKuliahRouter.PUT("/update/:id", authMiddleware.RequireAuth, mataKuliahController.UpdateMataKuliah)
		mataKuliahRouter.POST("/rps/:id", authMiddleware.RequireAuth, mataKuliahController.GetRPS)
	}

	plottingDosenMkRouter := r.Group("/plotting_dosen_mk")
	{
		plottingDosenMkRouter.GET("/", authMiddleware.RequireAuth, plottingDosenController.GetPlottingDosenMk)
		plottingDosenMkRouter.POST("/", authMiddleware.RequireAuth, plottingDosenController.CreatePlottingDosenMk)
		plottingDosenMkRouter.DELETE("/delete/:id", authMiddleware.RequireAuth, plottingDosenController.DeletePlottingDosenMk)
	}

	lembarAssessmentRouter := r.Group("/lembar_assessment")
	{
		lembarAssessmentRouter.GET("/", authMiddleware.RequireAuth, lembarAssessmentController.GetLembarAssessment)
		lembarAssessmentRouter.GET("/:id", authMiddleware.RequireAuth, lembarAssessmentController.GetLembarAssessmentById)
		lembarAssessmentRouter.POST("/", authMiddleware.RequireAuth, lembarAssessmentController.CreateLembarAssessment)
		lembarAssessmentRouter.DELETE("/delete/:id", authMiddleware.RequireAuth, lembarAssessmentController.DeleteLembarAssessment)
		lembarAssessmentRouter.PUT("/update/:id", authMiddleware.RequireAuth, lembarAssessmentController.UpdateLembarAssessment)
	}

	jenisAssessmentRouter := r.Group("/jenis_assessment")
	{
		jenisAssessmentRouter.GET("/", authMiddleware.RequireAuth, jenisAssessmentController.GetJenisAssessment)
		jenisAssessmentRouter.POST("/", authMiddleware.RequireAuth, jenisAssessmentController.CreateJenisAssessment)
		jenisAssessmentRouter.DELETE("/delete/:id", authMiddleware.RequireAuth, jenisAssessmentController.DeleteJenisAssessment)
	}

	mkMahasiswaRouter := r.Group("/mk_mahasiswa")
	{
		mkMahasiswaRouter.GET("/", authMiddleware.RequireAuth, mkMahasiswaController.GetMkMahasiswa)
	}

	penilaianRouter := r.Group("/penilaian")
	{
		penilaianRouter.GET("/", authMiddleware.RequireAuth, penilaianController.GetPenilaian)
		penilaianRouter.GET("/:id", authMiddleware.RequireAuth, penilaianController.GetPenilaianById)
		penilaianRouter.GET("/kelas/:kelasid", authMiddleware.RequireAuth, penilaianController.GetPenilaianByKelas)
		penilaianRouter.POST("/", authMiddleware.RequireAuth, penilaianController.CreatePenilaian)
		penilaianRouter.DELETE("/delete/:id", authMiddleware.RequireAuth, penilaianController.DeletePenilaian)
		penilaianRouter.PUT("/update/:id", authMiddleware.RequireAuth, penilaianController.UpdatePenilaian)
		penilaianRouter.POST("/upload", authMiddleware.RequireAuth, penilaianController.UploadEvidence)
	}

	tahunAjaranRouter := r.Group("/tahun_ajaran")
	{
		tahunAjaranRouter.GET("/", authMiddleware.RequireAuth, tahunAjaranController.GetTahunAjaran)
	}

	beritaAcaraRouter := r.Group("/berita_acara")
	{
		beritaAcaraRouter.GET("/", authMiddleware.RequireAuth, beritaAcaraController.GetBeritaAcara)
		beritaAcaraRouter.POST("/", authMiddleware.RequireAuth, beritaAcaraController.CreateBeritaAcara)
		beritaAcaraRouter.DELETE("/delete/:id", authMiddleware.RequireAuth, beritaAcaraController.DeleteBeritaAcara)
	}

	kelasRouter := r.Group("/kelas")
	{
		kelasRouter.GET("/", authMiddleware.RequireAuth, kelasController.GetKelas)
	}

	mahasiswaRouter := r.Group("/mahasiswa")
	{
		mahasiswaRouter.GET("/", authMiddleware.RequireAuth, mahasiswaController.GetMahasiswa)
		mahasiswaRouter.GET("/mata_kuliah/:mkId", authMiddleware.RequireAuth, mahasiswaController.GetMahasiswaByMataKuliah)
		mahasiswaRouter.POST("/mata_kuliah/:mkId/kelas/:kelasId", authMiddleware.RequireAuth, mahasiswaController.GetMahasiswaByKelasMataKuliah)
	}

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, "Hello World")
	})
	r.Run()
}
