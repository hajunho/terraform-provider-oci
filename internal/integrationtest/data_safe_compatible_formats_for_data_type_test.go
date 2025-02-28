// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
)

var (
	compatibleFormatsForDataTypeSingularDataSourceRepresentation = map[string]interface{}{}

	CompatibleFormatsForDataTypeResourceConfig = ""
)

// issue-routing-tag: data_safe/default
func TestDataSafeCompatibleFormatsForDataTypeResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeCompatibleFormatsForDataTypeResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	singularDatasourceName := "data.oci_data_safe_compatible_formats_for_data_type.test_compatible_formats_for_data_type"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_compatible_formats_for_data_type", "test_compatible_formats_for_data_type", acctest.Required, acctest.Create, compatibleFormatsForDataTypeSingularDataSourceRepresentation) +
				compartmentIdVariableStr + CompatibleFormatsForDataTypeResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(

				resource.TestCheckResourceAttr(singularDatasourceName, "formats_for_data_type.#", "5"),
			),
		},
	})
}
