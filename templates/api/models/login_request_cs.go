package templates

import "github.com/ortizdavid/dotnet-tpl/helpers"

func (ApiModelsTemplate) LoginRequestCs() string {
return `namespace `+helpers.GetCurrentFolder()+`.Models
{
    public class LoginRequest
    {
        public string? Username { get; set; }
        public string? Password { get; set; }
    }
}
`
}