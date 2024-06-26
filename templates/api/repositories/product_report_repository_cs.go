package templates

import "github.com/ortizdavid/dotnet-tpl/helpers"

func (ApiRepositoriesTemplate) ProductReportRepositoryCs() string {
return `using `+helpers.GetCurrentFolder()+`.Models;

namespace `+helpers.GetCurrentFolder()+`.Repositories
{
    public class ProductReportRepository 
    {
        public IEnumerable<ProductReport> GetFieldsForReport(IEnumerable<ProductData> products)
        {
            return products.Select(p => new ProductReport
            {
                ProductId = p.ProductId,
                ProductName = p.ProductName,
                Code = p.Code,
                UnitPrice = p.UnitPrice,
                CategoryName = p.CategoryName,
                Description = p.Description,
                CreatedAt = p.CreatedAt,
                UpdatedAt = p.UpdatedAt
            });
        }
    }
}
`
}