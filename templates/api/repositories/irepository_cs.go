package templates

import "github.com/ortizdavid/dotnet-tpl/helpers"

func (ApiRepositoriesTemplate) IRepositoryCs() string {
return `using System;
using System.Collections.Generic;
using System.Threading.Tasks;

namespace `+helpers.GetCurrentFolder()+`.Repositories
{
    public interface IRepository<T> where T : class
    {
        Task CreateAsync(T entity);
        Task CreateBatchAsync(List<T> entities);
        Task UpdateAsync(T entity);
        Task DeleteAsync(T entity);
        Task<List<T>> GetAllAsync();
        Task<T?> GetByIdAsync(int id);
        Task<T?> GetByUniqueIdAsync(Guid uniqueId);
        Task<bool> ExistsAsync(string? predicate);
    }
}
`
}