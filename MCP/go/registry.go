package main

import (
	"github.com/aws-account/mcp-server/config"
	"github.com/aws-account/mcp-server/models"
	tools_deletealternatecontact "github.com/aws-account/mcp-server/tools/deletealternatecontact"
	tools_disableregion "github.com/aws-account/mcp-server/tools/disableregion"
	tools_enableregion "github.com/aws-account/mcp-server/tools/enableregion"
	tools_getcontactinformation "github.com/aws-account/mcp-server/tools/getcontactinformation"
	tools_listregions "github.com/aws-account/mcp-server/tools/listregions"
	tools_putcontactinformation "github.com/aws-account/mcp-server/tools/putcontactinformation"
	tools_getalternatecontact "github.com/aws-account/mcp-server/tools/getalternatecontact"
	tools_getregionoptstatus "github.com/aws-account/mcp-server/tools/getregionoptstatus"
	tools_putalternatecontact "github.com/aws-account/mcp-server/tools/putalternatecontact"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_deletealternatecontact.CreateDeletealternatecontactTool(cfg),
		tools_disableregion.CreateDisableregionTool(cfg),
		tools_enableregion.CreateEnableregionTool(cfg),
		tools_getcontactinformation.CreateGetcontactinformationTool(cfg),
		tools_listregions.CreateListregionsTool(cfg),
		tools_putcontactinformation.CreatePutcontactinformationTool(cfg),
		tools_getalternatecontact.CreateGetalternatecontactTool(cfg),
		tools_getregionoptstatus.CreateGetregionoptstatusTool(cfg),
		tools_putalternatecontact.CreatePutalternatecontactTool(cfg),
	}
}
