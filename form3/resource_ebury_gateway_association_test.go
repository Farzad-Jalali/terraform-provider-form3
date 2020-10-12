package form3

import (
	"fmt"
	"os"
	"testing"

	form3 "github.com/form3tech-oss/terraform-provider-form3/api"
	"github.com/form3tech-oss/terraform-provider-form3/client/associations"
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccEburyAssociation_basic(t *testing.T) {
	parentOrganisationId := os.Getenv("FORM3_ORGANISATION_ID")
	organisationId := uuid.New().String()
	associationId := uuid.New().String()
	organisationLocation := "GB"
	fundingCurrency := "GBP"
	eburyContactId := "CON100000"
	eburyClientId := "CLI100000"
	partyPaymentFee := "2.00"
	organisationPaymentFee := "3.00"
	organisationKYCModel := "Reliance"
	partyName := "Test party"
	partyAddress := []string{"Test", "party", "address"}
	partyCity := "Test city"
	partyPostCode := "TE1 2ST"

	locationUpdate := "FR"
	currencyUpdate := "EUR"
	partyFeeUpdate := "3.50"
	orgFeeUpdate := "4.50"
	nameUpdate := "nom"
	addressUpdate := []string{"1 Le Rue", "ou", "Paris"}
	cityUpdate := "Paris"
	postcodeUpdate := "L42956"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckEburyAssociationDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testForm3EburyAssociationConfig, organisationId, parentOrganisationId, organisationLocation, associationId,
					fundingCurrency, eburyContactId, eburyClientId, partyPaymentFee, organisationPaymentFee, organisationKYCModel,
					partyName, partyAddress[0], partyAddress[1], partyAddress[2], partyCity, partyPostCode),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckEburyAssociationExists("form3_ebury_association.association"),
					resource.TestCheckResourceAttr("form3_ebury_association.association", "association_id", associationId),
					resource.TestCheckResourceAttr("form3_ebury_association.association", "organisation_id", organisationId),
					resource.TestCheckResourceAttr("form3_ebury_association.association", "organisation_location", organisationLocation),
					resource.TestCheckResourceAttr("form3_ebury_association.association", "ebury_contact_id", eburyContactId),
					resource.TestCheckResourceAttr("form3_ebury_association.association", "ebury_client_id", eburyClientId),
					resource.TestCheckResourceAttr("form3_ebury_association.association", "party_payment_fee", partyPaymentFee),
					resource.TestCheckResourceAttr("form3_ebury_association.association", "organisation_payment_fee", organisationPaymentFee),
					resource.TestCheckResourceAttr("form3_ebury_association.association", "organisation_kyc_model", organisationKYCModel),
					resource.TestCheckResourceAttr("form3_ebury_association.association", "party_name", partyName),
					resource.TestCheckResourceAttr("form3_ebury_association.association", "party_address.0", partyAddress[0]),
					resource.TestCheckResourceAttr("form3_ebury_association.association", "party_address.1", partyAddress[1]),
					resource.TestCheckResourceAttr("form3_ebury_association.association", "party_address.2", partyAddress[2]),
					resource.TestCheckResourceAttr("form3_ebury_association.association", "party_city", partyCity),
					resource.TestCheckResourceAttr("form3_ebury_association.association", "party_post_code", partyPostCode),
				),
			},
			{
				Config: fmt.Sprintf(testForm3EburyAssociationConfig, organisationId, parentOrganisationId, locationUpdate, associationId,
					currencyUpdate, eburyContactId, eburyClientId, partyFeeUpdate, orgFeeUpdate, organisationKYCModel,
					nameUpdate, addressUpdate[0], addressUpdate[1], addressUpdate[2], cityUpdate, postcodeUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckEburyAssociationExists("form3_ebury_association.association"),
					resource.TestCheckResourceAttr("form3_ebury_association.association", "association_id", associationId),
					resource.TestCheckResourceAttr("form3_ebury_association.association", "organisation_id", organisationId),
					resource.TestCheckResourceAttr("form3_ebury_association.association", "organisation_location", locationUpdate),
					resource.TestCheckResourceAttr("form3_ebury_association.association", "ebury_contact_id", eburyContactId),
					resource.TestCheckResourceAttr("form3_ebury_association.association", "ebury_client_id", eburyClientId),
					resource.TestCheckResourceAttr("form3_ebury_association.association", "party_payment_fee", partyFeeUpdate),
					resource.TestCheckResourceAttr("form3_ebury_association.association", "organisation_payment_fee", orgFeeUpdate),
					resource.TestCheckResourceAttr("form3_ebury_association.association", "organisation_kyc_model", organisationKYCModel),
					resource.TestCheckResourceAttr("form3_ebury_association.association", "party_name", nameUpdate),
					resource.TestCheckResourceAttr("form3_ebury_association.association", "party_address.0", addressUpdate[0]),
					resource.TestCheckResourceAttr("form3_ebury_association.association", "party_address.1", addressUpdate[1]),
					resource.TestCheckResourceAttr("form3_ebury_association.association", "party_address.2", addressUpdate[2]),
					resource.TestCheckResourceAttr("form3_ebury_association.association", "party_city", cityUpdate),
					resource.TestCheckResourceAttr("form3_ebury_association.association", "party_post_code", postcodeUpdate),
				),
			},
		},
	})
}

func testAccCheckEburyAssociationDestroy(state *terraform.State) error {
	client := testAccProvider.Meta().(*form3.AuthenticatedClient)

	for _, rs := range state.RootModule().Resources {
		if rs.Type != "form3_ebury_association" {
			continue
		}

		response, err := client.AssociationClient.Associations.GetEburyID(associations.NewGetEburyIDParams().
			WithID(strfmt.UUID(rs.Primary.ID)))

		if err == nil {
			return fmt.Errorf("ebury record %s still exists, %+v", rs.Primary.ID, response)
		}
	}

	return nil
}

func testAccCheckEburyAssociationExists(resourceKey string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceKey]

		if !ok {
			return fmt.Errorf("not found: %s", resourceKey)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no ebury Record ID is set")
		}

		client := testAccProvider.Meta().(*form3.AuthenticatedClient)

		foundRecord, err := client.AssociationClient.Associations.GetEburyID(associations.NewGetEburyIDParams().
			WithID(strfmt.UUID(rs.Primary.ID)))

		if err != nil {
			return err
		}

		if foundRecord.Payload.Data.ID.String() != rs.Primary.ID {
			return fmt.Errorf("ebury record not found, expected %s but found %s", rs.Primary.ID, foundRecord.Payload.Data.ID.String())
		}

		return nil
	}
}

const testForm3EburyAssociationConfig = `
resource "form3_organisation" "organisation" {
	organisation_id        = "%s"
	parent_organisation_id = "%s"
	name 		           = "terraform-organisation"
}

resource "form3_ebury_association" "association" {
	organisation_id  		 = "${form3_organisation.organisation.organisation_id}"
	organisation_location 	 = "%s"
	association_id           = "%s"
	funding_currency 		 = "%s"
	ebury_contact_id 		 = "%s"
	ebury_client_id  		 = "%s"
	party_payment_fee		 = "%s"
	organisation_payment_fee = "%s"
	organisation_kyc_model   = "%s"
	party_name 				 = "%s"
	party_address  			 = ["%s", "%s", "%s"]
	party_city 				 = "%s"
	party_post_code 		 = "%s"
}`
