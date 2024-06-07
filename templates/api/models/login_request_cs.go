package templates

func (ApiModelsTemplate) LoginRequestCs() string {
return `namespace AspNetCoreRestApi.Models
{
    public class LoginRequest
    {
        public string? Username { get; set; }
        public string? Password { get; set; }
    }
}
`
}