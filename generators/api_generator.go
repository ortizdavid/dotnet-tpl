package generators

import (
	apiTemplates "github.com/ortizdavid/dotnet-tpl/templates/api"
	colTemplates "github.com/ortizdavid/dotnet-tpl/templates/api/api_collections"
	dbTemplates "github.com/ortizdavid/dotnet-tpl/templates/api/database"
	fileTemplates "github.com/ortizdavid/dotnet-tpl/templates/api/files"
	extTemplates "github.com/ortizdavid/dotnet-tpl/templates/api/extensions"
	helperTemplates "github.com/ortizdavid/dotnet-tpl/templates/api/helpers"
	modelTemplates "github.com/ortizdavid/dotnet-tpl/templates/api/models"
	repoTemplates "github.com/ortizdavid/dotnet-tpl/templates/api/repositories"
	contTemplates "github.com/ortizdavid/dotnet-tpl/templates/api/controllers"
	"github.com/ortizdavid/go-nopain/filemanager"
)

type ApiGenerator struct {
}

func (ApiGenerator) Generate(templateType string) error {
	var fileManager filemanager.FileManager
	// templates Structs
	var apiTpl apiTemplates.ApiTemplate
	var colTpl colTemplates.ApiCollectionsTemplate
	var dbTpl dbTemplates.ApiDatabaseTemplate
	var fileTpl fileTemplates.ApiFilesTemplate
	var extTpl extTemplates.ApiExtensionsTemplate
	var helTpl helperTemplates.ApiHelpersTemplate
	var modTpl modelTemplates.ApiModelsTemplate
	var repTpl repoTemplates.ApiRepositoriesTemplate
	var conTpl contTemplates.ApiControllersTemplate

	// Folders---------------------------------------------------------
	apiCollectionsFolder := "_Api_Collections"
	uploadsFolder := "_Uploads"
	databaseFolder := "_Database"
	filesFolder := "_Files"
	extensionsFolder := "Extensions"
	helpersFolder := "Helpers"
	modelsFolder := "Models"
	repositoriesFolder := "Repositories"
	controllersFolder := "Controllers"

	// Files outter folder ----------------------------------------------
	dockerFile := "Dockerfile"
	dockerignoreFile := ".dockerignore"
	gitignoreFile := ".gitignore"
	readmeMdFile := "README.md"
	appsettingsFile := "appsettings.json"
	packagesFile := "packages.sh"
	programCsFile := "Program.cs"

	// Files inner folder ------------------------------
	// files from /_Api_Collections
	postmanCollectionFile := "postman_collection.json"
	// files from /_Database
	databaseFile := "db_dotnet_product.sql"
	// files from /_Files
	categoriesCsvFile := "categories-for-import.csv"
	productsCsvFile := "products-for-import.csv"
	suppliersCsvFile := "suppliers-for-import.csv"
	// files from /Controllers
	authControllerFile := "AuthController.cs"
	categoriesControllerFile := "CategoriesController.cs"
	productsControllerFile := "ProductsControllers.cs"
	productsReportControllerFile := "ProductsReportController.cs"
	suppliersControllerFile := "SuppliersController.cs"
	usersControllerFile := "UsersController.cs"
	// files from /Extensions
	authExtensionsFile := "AuthExtensions.cs"
	repositoryExtensionsFile := "RepositoryExtensions.cs"
	// files from /Helpers
	capacityUnitFile := "CapacityUnit.cs"
	encryptionFile := "Encryption.cs"
	fileExtensionsFile := "FileExtensions.cs"
	fileUploaderFile := "FileUploader.cs"
	passwordHelperFile := "PasswordHelper.cs"
	// files from /Models
	appDbContextFile := "AppDbContext.cs"
	categoryFile := "Category.cs"
	imageFile := "Image.cs"
	loginRequestFile := "LoginRequest.cs"
	productFile := "Product.cs"
	productDataFile := "ProductData.cs"
	productReportFile := "ProductReport.cs"
	supplierFile := "Supplier.cs"
	userFile := "User.cs"
	// files from /Repositories
	iRepositoryFile := "IRepository.cs"
	categoryRepositoryFile := "CategoryRepository.cs"
	imageRepositoryFile := "ImageRepository.cs"
	productRepositoryFile := "ProductRepository.cs"
	productReportRepositoryFile := "ProductReportRepository.cs"
	supplierRepositoryFile := "SupplierRepository.cs"
	userRepositoryFile := "UserRepository.cs"

	// Create Files--------------------------------------------------
	// Create Files outter folder
	fileManager.CreateSingleFile(".", dockerFile)
	fileManager.CreateSingleFile(".", dockerignoreFile)
	fileManager.CreateSingleFile(".", gitignoreFile)
	fileManager.CreateSingleFile(".", readmeMdFile)
	fileManager.CreateSingleFile(".", appsettingsFile)
	fileManager.CreateSingleFile(".", packagesFile)
	fileManager.CreateSingleFile(".", programCsFile)

	// Create folders
	fileManager.CreateSingleFolder(apiCollectionsFolder)
	fileManager.CreateSingleFolder(uploadsFolder)
	fileManager.CreateSingleFolder(databaseFolder)
	fileManager.CreateSingleFolder(filesFolder)
	fileManager.CreateSingleFolder(extensionsFolder)
	fileManager.CreateSingleFolder(helpersFolder)
	fileManager.CreateSingleFolder(modelsFolder)
	fileManager.CreateSingleFolder(repositoriesFolder)
	fileManager.CreateSingleFolder(controllersFolder)

	// Create Files in folders
	// files from /_Api_Collections
	fileManager.CreateSingleFile(apiCollectionsFolder, postmanCollectionFile)
	// files from /_Database
	fileManager.CreateSingleFile(databaseFolder, databaseFile)
	// files from /_Files
	fileManager.CreateSingleFile(filesFolder, categoriesCsvFile)
	fileManager.CreateSingleFile(filesFolder, productsCsvFile)
	fileManager.CreateSingleFile(filesFolder, suppliersCsvFile)
	// files from /Controllers
	fileManager.CreateSingleFile(controllersFolder, authControllerFile)
	fileManager.CreateSingleFile(controllersFolder, categoriesControllerFile)
	fileManager.CreateSingleFile(controllersFolder, productsControllerFile)
	fileManager.CreateSingleFile(controllersFolder, productsReportControllerFile)
	fileManager.CreateSingleFile(controllersFolder, suppliersControllerFile)
	fileManager.CreateSingleFile(controllersFolder, usersControllerFile)
	// files from /Extensions
	fileManager.CreateSingleFile(extensionsFolder, authExtensionsFile)
	fileManager.CreateSingleFile(extensionsFolder, repositoryExtensionsFile)
	// files from /Helpers
	fileManager.CreateSingleFile(helpersFolder, capacityUnitFile)
	fileManager.CreateSingleFile(helpersFolder, encryptionFile)
	fileManager.CreateSingleFile(helpersFolder, fileExtensionsFile)
	fileManager.CreateSingleFile(helpersFolder, fileUploaderFile)
	fileManager.CreateSingleFile(helpersFolder, passwordHelperFile)
	// files from /Models
	fileManager.CreateSingleFile(modelsFolder, appDbContextFile)
	fileManager.CreateSingleFile(modelsFolder, categoryFile)
	fileManager.CreateSingleFile(modelsFolder, imageFile)
	fileManager.CreateSingleFile(modelsFolder, loginRequestFile)
	fileManager.CreateSingleFile(modelsFolder, productFile)
	fileManager.CreateSingleFile(modelsFolder, productDataFile)
	fileManager.CreateSingleFile(modelsFolder, productReportFile)
	fileManager.CreateSingleFile(modelsFolder, supplierFile)
	fileManager.CreateSingleFile(modelsFolder, userFile)
	// files from /Repositories
	fileManager.CreateSingleFile(repositoriesFolder, iRepositoryFile)
	fileManager.CreateSingleFile(repositoriesFolder, categoryRepositoryFile)
	fileManager.CreateSingleFile(repositoriesFolder, imageRepositoryFile)
	fileManager.CreateSingleFile(repositoriesFolder, productRepositoryFile)
	fileManager.CreateSingleFile(repositoriesFolder, productReportRepositoryFile)
	fileManager.CreateSingleFile(repositoriesFolder, supplierRepositoryFile)
	fileManager.CreateSingleFile(repositoriesFolder, userRepositoryFile)

	// Write to files --------------------------------------------------------------
	// Write to files from outer folder
	fileManager.WriteFile(".", dockerFile, apiTpl.Dockerfile())
	fileManager.WriteFile(".", dockerignoreFile, apiTpl.Dockerignore())
	fileManager.WriteFile(".", gitignoreFile, apiTpl.Gitignore())
	fileManager.WriteFile(".", readmeMdFile, apiTpl.ReadmeMd())
	fileManager.WriteFile(".", appsettingsFile, apiTpl.Appsettings())
	fileManager.WriteFile(".", packagesFile, apiTpl.Packages())
	fileManager.WriteFile(".", programCsFile, apiTpl.ProgramCs())
	// Write to files from /_Api_Collections
	fileManager.WriteFile(apiCollectionsFolder, postmanCollectionFile, colTpl.PostmanCollectionsJson())
	// Write to files from /_Database
	fileManager.WriteFile(databaseFolder, databaseFile, dbTpl.DbDotnetProductSQL())

	// Write to files from /_Files
	fileManager.WriteFile(filesFolder, categoriesCsvFile, fileTpl.CategoriesForImportCsv())
	fileManager.WriteFile(filesFolder, productsCsvFile, fileTpl.ProductsForImportCsv())
	fileManager.WriteFile(filesFolder, suppliersCsvFile, fileTpl.SuppliersForImportCsv())

	// Write to files from /Controllers
	fileManager.WriteFile(controllersFolder, authControllerFile, conTpl.AuthControllerCs())
	fileManager.WriteFile(controllersFolder, categoriesControllerFile, conTpl.CategoriesControllerCs())
	fileManager.WriteFile(controllersFolder, productsControllerFile, conTpl.ProductsControllerCs())
	fileManager.WriteFile(controllersFolder, productsReportControllerFile, conTpl.ProductsReportControllerCs())
	fileManager.WriteFile(controllersFolder, suppliersControllerFile, conTpl.SuppliersControllerCs())
	fileManager.WriteFile(controllersFolder, usersControllerFile, conTpl.UsersControllerCs())

	// Write to files from /Extensions
	fileManager.WriteFile(extensionsFolder, authExtensionsFile, extTpl.AuthExtensionsCs())
	fileManager.WriteFile(extensionsFolder, repositoryExtensionsFile, extTpl.RepositoryExtensionsCs())

	// Write to files from /Helpers
	fileManager.WriteFile(helpersFolder, capacityUnitFile, helTpl.CapacityUnitCs())
	fileManager.WriteFile(helpersFolder, encryptionFile, helTpl.EncryptionCs())
	fileManager.WriteFile(helpersFolder, fileExtensionsFile, helTpl.FileExtensionsCs())
	fileManager.WriteFile(helpersFolder, fileUploaderFile, helTpl.FileUploaderCs())
	fileManager.WriteFile(helpersFolder, passwordHelperFile, helTpl.PasswordHelperCs())

	// Write to files from /Models
	fileManager.WriteFile(modelsFolder, appDbContextFile, modTpl.AppDbContextCs())
	fileManager.WriteFile(modelsFolder, categoryFile, modTpl.CategoryCs())
	fileManager.WriteFile(modelsFolder, imageFile, modTpl.ImageCs())
	fileManager.WriteFile(modelsFolder, loginRequestFile, modTpl.LoginRequestCs())
	fileManager.WriteFile(modelsFolder, productFile, modTpl.ProductCs())
	fileManager.WriteFile(modelsFolder, productDataFile, modTpl.ProductDataCs())
	fileManager.WriteFile(modelsFolder, productReportFile, modTpl.ProductReportCs())
	fileManager.WriteFile(modelsFolder, supplierFile, modTpl.SupplierCs())
	fileManager.WriteFile(modelsFolder, userFile, modTpl.UserCs())

	// Write to files from /Repositories
	fileManager.WriteFile(repositoriesFolder, iRepositoryFile, repTpl.IRepositoryCs())
	fileManager.WriteFile(repositoriesFolder, categoryRepositoryFile, repTpl.CategoryRepositoryCs())
	fileManager.WriteFile(repositoriesFolder, imageRepositoryFile, repTpl.ImageRepositoryCs())
	fileManager.WriteFile(repositoriesFolder, productRepositoryFile, repTpl.ProductRepositoryCs())
	fileManager.WriteFile(repositoriesFolder, productReportRepositoryFile, repTpl.ProductReportRepositoryCs())
	fileManager.WriteFile(repositoriesFolder, supplierRepositoryFile, repTpl.SupplierRepositoryCs())
	fileManager.WriteFile(repositoriesFolder, userRepositoryFile, repTpl.UserRepositoryCs())

	return nil
}
