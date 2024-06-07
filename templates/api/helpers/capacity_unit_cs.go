package templates

import "github.com/ortizdavid/dotnet-tpl/helpers"

func (ApiHelpersTemplate) CapacityUnitCs() string  {
return `namespace `+helpers.GetCurrentFolder()+`.Helpers
{
    public class CapacityUnit
    {
        public const long KILO_BYTE = 1024;
        public const long MEGA_BYTE = KILO_BYTE * 1024;
        public const long GIGA_BYTE = MEGA_BYTE * 1024;
    }
}
`	
}