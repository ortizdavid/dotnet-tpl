package templates

import "github.com/ortizdavid/dotnet-tpl/helpers"

func (ApiExtensionsTemplate) RepositoryExtensionsCs() string {
	return `using Microsoft.Extensions.DependencyInjection;
using ` + helpers.GetCurrentFolder() + `.Repositories;

namespace ` + helpers.GetCurrentFolder() + `.Extensions
{
    public static class RepositoryExtensions
    {
        public static void AddRepositories(this IServiceCollection services)
        {
            services.AddScoped<CategoryRepository>();
            services.AddScoped<ProductRepository>();
            services.AddScoped<ImageRepository>();
            services.AddScoped<ProductReportRepository>();
            services.AddScoped<SupplierRepository>();
            services.AddScoped<UserRepository>();
        }
    }
}
`
}