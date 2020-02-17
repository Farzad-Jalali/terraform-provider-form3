package form3

import (
	"fmt"
	"os"
	"testing"

	form3 "github.com/form3tech-oss/terraform-provider-form3/api"
	"github.com/form3tech-oss/terraform-provider-form3/client/associations"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	uuid "github.com/satori/go.uuid"
)

func TestAccSepaLiquidityAssociation_basic(t *testing.T) {
	parentOrganisationID := os.Getenv("FORM3_ORGANISATION_ID")
	sponsorOrganisationID := uuid.NewV4().String()
	sponsoredOrganisationID := uuid.NewV4().String()
	sponsorAssociationID := uuid.NewV4().String()
	sponsoredAssociationID := uuid.NewV4().String()

	sponsor_assoc_path := "form3_sepaliquidity_association.sponsor_association"
	sponsored_assoc_path := "form3_sepaliquidity_association.sponsored_association"

	config := fmt.Sprintf(
		testForm3SepaLiquidityAssociationConfigA,
		parentOrganisationID,
		sponsorOrganisationID,
		sponsoredOrganisationID,
		sponsorAssociationID,
		sponsoredAssociationID,
	)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSepaLiquidityAssociationDestroy,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSepaLiquidityAssociationExists(sponsor_assoc_path),
					resource.TestCheckResourceAttr(sponsor_assoc_path, "association_id", sponsorAssociationID),
					resource.TestCheckResourceAttr(sponsor_assoc_path, "organisation_id", sponsorOrganisationID),
					resource.TestCheckResourceAttr(sponsor_assoc_path, "name", "Sponsor company"),
					resource.TestCheckResourceAttr(sponsor_assoc_path, "technical_bic", "TESTBIC1"),
					resource.TestCheckResourceAttr(sponsor_assoc_path, "settlement_bic", "BICBIC99"),
					resource.TestCheckResourceAttr(sponsor_assoc_path, "settlement_iban", "GB22ABCD19283700000001"),
					resource.TestCheckResourceAttr(sponsor_assoc_path, "address_street", "Harp Ln"),
					resource.TestCheckResourceAttr(sponsor_assoc_path, "address_building_number", "7"),
					resource.TestCheckResourceAttr(sponsor_assoc_path, "address_city", "London"),
					resource.TestCheckResourceAttr(sponsor_assoc_path, "address_country", "United Kingdom"),
					resource.TestCheckResourceAttr(sponsor_assoc_path, "sponsor_id", ""),

					testAccCheckSepaLiquidityAssociationExists(sponsored_assoc_path),
					resource.TestCheckResourceAttr(sponsored_assoc_path, "association_id", sponsoredAssociationID),
					resource.TestCheckResourceAttr(sponsored_assoc_path, "organisation_id", sponsoredOrganisationID),
					resource.TestCheckResourceAttr(sponsored_assoc_path, "name", "Sponsored company"),
					resource.TestCheckResourceAttr(sponsored_assoc_path, "sponsored_bics.0", "BICBIC50XXX"),
					resource.TestCheckResourceAttr(sponsored_assoc_path, "sponsored_bics.1", "BICBIC51XXX"),
					resource.TestCheckResourceAttr(sponsored_assoc_path, "settlement_iban", "GB22ABCD19283700000002"),
					resource.TestCheckResourceAttr(sponsored_assoc_path, "address_street", "Harp Ln"),
					resource.TestCheckResourceAttr(sponsored_assoc_path, "address_building_number", "7"),
					resource.TestCheckResourceAttr(sponsored_assoc_path, "address_city", "London"),
					resource.TestCheckResourceAttr(sponsored_assoc_path, "address_country", "United Kingdom"),
					resource.TestCheckResourceAttr(sponsored_assoc_path, "sponsor_id", sponsorAssociationID),
				),
			},
		},
	})
}

func testAccCheckSepaLiquidityAssociationDestroy(state *terraform.State) error {
	client := testAccProvider.Meta().(*form3.AuthenticatedClient)

	for _, rs := range state.RootModule().Resources {
		if rs.Type != "form3_sepaliquidity_association" {
			continue
		}

		response, err := client.AssociationClient.Associations.GetSepaLiquidityID(associations.NewGetSepaLiquidityIDParams().
			WithID(strfmt.UUID(rs.Primary.ID)))

		if err == nil {
			return fmt.Errorf("sepa liquidity record %s still exists, %+v", rs.Primary.ID, response)
		}
	}

	return nil
}

func testAccCheckSepaLiquidityAssociationExists(resourceKey string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceKey]

		if !ok {
			return fmt.Errorf("not found: %s", resourceKey)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no sepa liquidity Record ID is set")
		}

		client := testAccProvider.Meta().(*form3.AuthenticatedClient)

		foundRecord, err := client.AssociationClient.Associations.GetSepaLiquidityID(associations.NewGetSepaLiquidityIDParams().
			WithID(strfmt.UUID(rs.Primary.ID)))

		if err != nil {
			return err
		}

		if foundRecord.Payload.Data.ID.String() != rs.Primary.ID {
			return fmt.Errorf("sepa liquidity record not found expected %s found %s", rs.Primary.ID, foundRecord.Payload.Data.ID.String())
		}

		return nil
	}
}

const testForm3SepaLiquidityAssociationConfigA = `
locals {
	parent_organisation_id    = "%s"

	sponsor_organisation_id   = "%s"
	sponsored_organisation_id = "%s"

	sponsor_association_id    = "%s"
	sponsored_association_id  = "%s"
}

resource "form3_organisation" "sponsor" {
	organisation_id        = "${local.sponsor_organisation_id}"
	parent_organisation_id = "${local.parent_organisation_id}"
	name 		               = "terraform-sponsor-organisation"
}

resource "form3_organisation" "sponsored" {
	organisation_id        = "${local.sponsored_organisation_id}"
	parent_organisation_id = "${local.parent_organisation_id}"
	name 		               = "terraform-sponsored-organisation"
}

resource "form3_sepaliquidity_association" "sponsor_association" {
	organisation_id         = "${form3_organisation.sponsor.organisation_id}"
	association_id          = "${local.sponsor_association_id}"
	name                    = "Sponsor company"
	technical_bic           = "TESTBIC1"
	settlement_bic          = "BICBIC99"
	settlement_iban         = "GB22ABCD19283700000001"
	address_street          = "Harp Ln"
	address_building_number = "7"
	address_city            = "London"
	address_country         = "United Kingdom"
}

resource "form3_sepaliquidity_association" "sponsored_association" {
	organisation_id         = "${form3_organisation.sponsored.organisation_id}"
	association_id          = "${local.sponsored_association_id}"
	name                    = "Sponsored company"
	sponsored_bics          = ["BICBIC50XXX", "BICBIC51XXX"]
	settlement_iban         = "GB22ABCD19283700000002"
	address_street          = "Harp Ln"
	address_building_number = "7"
	address_city            = "London"
	address_country         = "United Kingdom"
	sponsor_id              = "${form3_sepaliquidity_association.sponsor_association.association_id}"
}
`
