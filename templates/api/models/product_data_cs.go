package templates

import "github.com/ortizdavid/dotnet-tpl/helpers"

func (ApiModelsTemplate) ProductDataCs() string {
return `using System.ComponentModel.DataAnnotations.Schema;

namespace `+helpers.GetCurrentFolder()+`.Models
{
    public class ProductData
    {   
        [Column("product_id")]
        public int ProductId { get; set; }

        [Column("unique_id")]
        public Guid UniqueId { get; set; }

        [Column("product_name")]
        public string? ProductName { get; set; }

        [Column("code")]
        public string? Code { get; set; }

        [Column("unit_price")]
        public double UnitPrice { get; set; }

        [Column("description")]
        public string? Description { get; set; }

        [Column("created_at")]
        public DateTime CreatedAt { get; set; }

        [Column("updated_at")]
        public DateTime UpdatedAt { get; set; }

        [Column("category_id")]
        public int CategoryId { get; set; }

        [Column("category_name")]
        public string? CategoryName { get; set; }

        [Column("supplier_id")]
        public int SupplierId { get; set; }

        [Column("supplier_name")]
        public string? SupplierName { get; set; }
    }
}
`
}