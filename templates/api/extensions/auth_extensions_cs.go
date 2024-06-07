package templates

import "github.com/ortizdavid/dotnet-tpl/helpers"

func (ApiExtensionsTemplate) AuthExtensionsCs() string {
return `using System.Text;
using Microsoft.AspNetCore.Authentication.JwtBearer;
using Microsoft.IdentityModel.Tokens;
using Npgsql.Replication;

namespace `+helpers.GetCurrentFolder()+`.Extensions
{
    public static class AuthExtensions
    {
        public static void AddJwtAuthentication(this IServiceCollection services, IConfiguration configuration)  
        {
            var jwtSettings = configuration.GetSection("JwtSettings");

            services.AddAuthentication(options =>
            {
                options.DefaultAuthenticateScheme = JwtBearerDefaults.AuthenticationScheme;
                options.DefaultChallengeScheme = JwtBearerDefaults.AuthenticationScheme;
            })
            .AddJwtBearer(options =>
            {
                options.TokenValidationParameters = new TokenValidationParameters
                {
                    ValidateIssuer = true,
                    ValidateAudience = true,
                    ValidateLifetime = true,
                    ValidateIssuerSigningKey = true,
                    ValidIssuer = jwtSettings["Issuer"],
                    ValidAudience = jwtSettings["Audience"],
                    IssuerSigningKey = new SymmetricSecurityKey(Encoding.UTF8.GetBytes(jwtSettings["SecretKey"]))
                };
            });
        }
    }
}
`	
}