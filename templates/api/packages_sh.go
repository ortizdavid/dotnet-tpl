package templates

func (ApiTemplate) Packages() string {
return `dotnet add package Dapper
dotnet add package Microsoft.AspNetCore.Authentication.JwtBearer
dotnet add package Microsoft.AspNetCore.Identity.EntityFrameworkCore
dotnet add package Npgsql.EntityFrameworkCore.PostgreSQL
dotnet add package Microsoft.EntityFrameworkCore.SqlServer
dotnet add package Microsoft.VisualStudio.Azure.Containers.Tools.Targets
dotnet add package Swashbuckle.AspNetCore
dotnet add package System.Data.SqlClient
dotnet add package System.IdentityModel.Tokens.Jwt
dotnet add package BCrypt.Net-Next
dotnet add package DinkToPdf
dotnet add package EPPlus
dotnet add package iTextSharp
dotnet add package Swashbuckle.AspNetCore.Annotations
dotnet add package CsvHelper
`
}