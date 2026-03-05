package policy

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Parent func for invdividual stating functions
func updateState(d *schema.ResourceData, resp *jamfpro.ResourcePolicy) diag.Diagnostics {
	var diags diag.Diagnostics

	// Log decoded API response when TF_LOG=DEBUG (helps debug API omissions e.g. self_service.notification_type)
	if respJSON, err := json.MarshalIndent(resp, "", "  "); err == nil {
		log.Printf("[DEBUG] Policy API response (decoded by SDK) for id=%d:\n%s", resp.General.ID, string(respJSON))
	}

	if err := d.Set("id", strconv.Itoa(resp.General.ID)); err != nil {
		diags = append(diags, diag.FromErr(err)...)
	}

	// General/Root level
	stateGeneral(d, resp, &diags)

	// Scope
	stateScope(d, resp, &diags)

	// Self Service
	stateSelfService(d, resp, &diags)

	// Payloads
	statePayloads(d, resp, &diags)

	return diags
}
