package templates

func (ApiTemplate) Appsettings() string {
return `
{
  "Logging": {
    "LogLevel": {
      "Default": "Information",
      "Microsoft.AspNetCore": "Warning",
      "Microsoft.EntityFrameworkCore.Database.Command": "None"
    }
  },

  "AllowedHosts": "*",

  "ConnectionStrings": {
    "DefaultConnection": "Host=localhost;Port=5432;Database=db_dotnet_product;Username=postgres;Password=003334743LA032"
  },

  "JwtSettings": {
    "Issuer": "testIssuer",
    "Audience": "testAudience",
    "SecretKey": "ee05133a-138e-4dfc-827e-9290271c29df",
    "ExpiryMinutes": 30
  },

  "UploadsDirectory": "_Uploads"
}
`	
}