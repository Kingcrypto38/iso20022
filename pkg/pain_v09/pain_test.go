// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package pain_v09

import (
	"encoding/json"
	"encoding/xml"
	"testing"
	"time"

	"github.com/moov-io/iso20022/pkg/common"
	"github.com/moov-io/iso20022/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestDocumentPain00800109(t *testing.T) {
	sample := DocumentPain00800109{}
	err := sample.Validate()
	assert.NotNil(t, err)

	testTime, _ := time.Parse(time.RFC3339, utils.TestTimeString)
	sample = DocumentPain00800109{
		Xmlns: sample.NameSpace(),
		CstmrDrctDbtInitn: CustomerDirectDebitInitiationV09{
			GrpHdr: GroupHeader83{
				MsgId:   "MsgId",
				CreDtTm: common.ISODateTime(testTime),
				NbOfTxs: "001",
			},
		},
	}
	err = sample.Validate()
	assert.Nil(t, err)

	buf, err := json.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `{"Xmlns":"urn:iso:std:iso:20022:tech:xsd:pain.008.001.09","CstmrDrctDbtInitn":{"GrpHdr":{"MsgId":"MsgId","CreDtTm":"2014-11-12T11:45:26.371","NbOfTxs":"001","InitgPty":{}}}}`)

	buf, err = xml.Marshal(&sample)
	assert.Nil(t, err)
	assert.Equal(t, string(buf), `<Document xmlns="urn:iso:std:iso:20022:tech:xsd:pain.008.001.09" xmlns:xs="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><CstmrDrctDbtInitn><GrpHdr><MsgId>MsgId</MsgId><CreDtTm>2014-11-12T11:45:26.371</CreDtTm><NbOfTxs>001</NbOfTxs><InitgPty></InitgPty></GrpHdr></CstmrDrctDbtInitn></Document>`)
}

func TestNestedTypes(t *testing.T) {
	assert.NotNil(t, AccountIdentification4Choice{}.Validate())
	assert.NotNil(t, AccountSchemeName1Choice{}.Validate())
	assert.NotNil(t, ActiveOrHistoricCurrencyAndAmount{}.Validate())
	assert.NotNil(t, AddressType3Choice{}.Validate())
	assert.Nil(t, AdviceType1{}.Validate())
	assert.NotNil(t, AdviceType1Choice{}.Validate())
	assert.Nil(t, AmendmentInformationDetails13{}.Validate())
	assert.NotNil(t, Authorisation1Choice{}.Validate())
	assert.Nil(t, BranchAndFinancialInstitutionIdentification6{}.Validate())
	assert.Nil(t, BranchData3{}.Validate())
	assert.NotNil(t, CashAccount38{}.Validate())
	assert.NotNil(t, CashAccountType2Choice{}.Validate())
	assert.NotNil(t, CategoryPurpose1Choice{}.Validate())
	assert.NotNil(t, ClearingSystemIdentification2Choice{}.Validate())
	assert.NotNil(t, ClearingSystemMemberIdentification2{}.Validate())
	assert.Nil(t, Contact4{}.Validate())
	assert.Nil(t, CreditorReferenceInformation2{}.Validate())
	assert.NotNil(t, CreditorReferenceType1Choice{}.Validate())
	assert.NotNil(t, CreditorReferenceType2{}.Validate())
	assert.NotNil(t, CustomerDirectDebitInitiationV09{}.Validate())
	assert.NotNil(t, DateAndPlaceOfBirth1{}.Validate())
	assert.Nil(t, DatePeriod2{}.Validate())
	assert.Nil(t, DirectDebitTransaction10{}.Validate())
	assert.NotNil(t, DirectDebitTransactionInformation23{}.Validate())
	assert.NotNil(t, DiscountAmountAndType1{}.Validate())
	assert.NotNil(t, DiscountAmountType1Choice{}.Validate())
	assert.NotNil(t, DocumentAdjustment1{}.Validate())
	assert.Nil(t, DocumentLineIdentification1{}.Validate())
	assert.Nil(t, DocumentLineInformation1{}.Validate())
	assert.NotNil(t, DocumentLineType1{}.Validate())
	assert.NotNil(t, DocumentLineType1Choice{}.Validate())
	assert.NotNil(t, FinancialIdentificationSchemeName1Choice{}.Validate())
	assert.Nil(t, FinancialInstitutionIdentification18{}.Validate())
	assert.NotNil(t, Frequency36Choice{}.Validate())
	assert.NotNil(t, FrequencyAndMoment1{}.Validate())
	assert.NotNil(t, FrequencyPeriod1{}.Validate())
	assert.NotNil(t, Garnishment3{}.Validate())
	assert.NotNil(t, GarnishmentType1{}.Validate())
	assert.NotNil(t, GarnishmentType1Choice{}.Validate())
	assert.NotNil(t, GenericAccountIdentification1{}.Validate())
	assert.NotNil(t, GenericFinancialIdentification1{}.Validate())
	assert.NotNil(t, GenericIdentification30{}.Validate())
	assert.NotNil(t, GenericOrganisationIdentification1{}.Validate())
	assert.NotNil(t, GenericPersonIdentification1{}.Validate())
	assert.NotNil(t, GroupHeader83{}.Validate())
	assert.NotNil(t, LocalInstrument2Choice{}.Validate())
	assert.Nil(t, MandateRelatedInformation14{}.Validate())
	assert.NotNil(t, MandateSetupReason1Choice{}.Validate())
	assert.NotNil(t, NameAndAddress16{}.Validate())
	assert.Nil(t, OrganisationIdentification29{}.Validate())
	assert.NotNil(t, OrganisationIdentificationSchemeName1Choice{}.Validate())
	assert.NotNil(t, OtherContact1{}.Validate())
	assert.Nil(t, Party38Choice{}.Validate())
	assert.Nil(t, PartyIdentification135{}.Validate())
	assert.NotNil(t, PaymentIdentification6{}.Validate())
	assert.NotNil(t, PaymentInstruction37{}.Validate())
	assert.Nil(t, PaymentTypeInformation29{}.Validate())
	assert.Nil(t, PersonIdentification13{}.Validate())
	assert.NotNil(t, PersonIdentificationSchemeName1Choice{}.Validate())
	assert.Nil(t, PostalAddress24{}.Validate())
	assert.NotNil(t, ProxyAccountIdentification1{}.Validate())
	assert.NotNil(t, ProxyAccountType1Choice{}.Validate())
	assert.NotNil(t, Purpose2Choice{}.Validate())
	assert.Nil(t, ReferredDocumentInformation7{}.Validate())
	assert.NotNil(t, ReferredDocumentType3Choice{}.Validate())
	assert.NotNil(t, ReferredDocumentType4{}.Validate())
	assert.Nil(t, RegulatoryAuthority2{}.Validate())
	assert.Nil(t, RegulatoryReporting3{}.Validate())
	assert.Nil(t, RemittanceAmount2{}.Validate())
	assert.Nil(t, RemittanceAmount3{}.Validate())
	assert.Nil(t, RemittanceInformation16{}.Validate())
	assert.Nil(t, RemittanceLocation7{}.Validate())
	assert.NotNil(t, RemittanceLocationData1{}.Validate())
	assert.NotNil(t, ServiceLevel8Choice{}.Validate())
	assert.Nil(t, StructuredRegulatoryReporting3{}.Validate())
	assert.Nil(t, StructuredRemittanceInformation16{}.Validate())
	assert.Nil(t, SupplementaryData1{}.Validate())
	assert.Nil(t, SupplementaryDataEnvelope1{}.Validate())
	assert.Nil(t, TaxAmount2{}.Validate())
	assert.NotNil(t, TaxAmountAndType1{}.Validate())
	assert.NotNil(t, TaxAmountType1Choice{}.Validate())
	assert.Nil(t, TaxAuthorisation1{}.Validate())
	assert.Nil(t, TaxInformation7{}.Validate())
	assert.Nil(t, TaxInformation8{}.Validate())
	assert.Nil(t, TaxParty1{}.Validate())
	assert.Nil(t, TaxParty2{}.Validate())
	assert.Nil(t, TaxPeriod2{}.Validate())
	assert.Nil(t, TaxRecord2{}.Validate())
	assert.NotNil(t, TaxRecordDetails2{}.Validate())
}

func TestTypes(t *testing.T) {
	var type1 ExternalAccountIdentification1Code
	assert.NotNil(t, type1.Validate())
	type1 = "test"
	assert.Nil(t, type1.Validate())

	var type2 ExternalCashAccountType1Code
	assert.NotNil(t, type2.Validate())
	type2 = "test"
	assert.Nil(t, type2.Validate())

	var type3 ExternalCategoryPurpose1Code
	assert.NotNil(t, type3.Validate())
	type3 = "test"
	assert.Nil(t, type3.Validate())

	var type4 ExternalClearingSystemIdentification1Code
	assert.NotNil(t, type4.Validate())
	type4 = "test"
	assert.Nil(t, type4.Validate())

	var type6 ExternalDiscountAmountType1Code
	assert.NotNil(t, type6.Validate())
	type6 = "test"
	assert.Nil(t, type6.Validate())

	var type8 ExternalDocumentLineType1Code
	assert.NotNil(t, type8.Validate())
	type8 = "test"
	assert.Nil(t, type8.Validate())

	var type10 ExternalFinancialInstitutionIdentification1Code
	assert.NotNil(t, type10.Validate())
	type10 = "test"
	assert.Nil(t, type10.Validate())

	var type11 ExternalGarnishmentType1Code
	assert.NotNil(t, type11.Validate())
	type11 = "test"
	assert.Nil(t, type11.Validate())

	var type12 ExternalLocalInstrument1Code
	assert.NotNil(t, type12.Validate())
	type12 = "test"
	assert.Nil(t, type12.Validate())

	var type13 ExternalMandateSetupReason1Code
	assert.NotNil(t, type13.Validate())
	type13 = "test"
	assert.Nil(t, type13.Validate())

	var type14 ExternalOrganisationIdentification1Code
	assert.NotNil(t, type14.Validate())
	type14 = "test"
	assert.Nil(t, type14.Validate())

	var type15 ExternalPersonIdentification1Code
	assert.NotNil(t, type15.Validate())
	type15 = "test"
	assert.Nil(t, type15.Validate())

	var type16 ExternalProxyAccountType1Code
	assert.NotNil(t, type16.Validate())
	type16 = "test"
	assert.Nil(t, type16.Validate())

	var type17 ExternalPurpose1Code
	assert.NotNil(t, type17.Validate())
	type17 = "test"
	assert.Nil(t, type17.Validate())

	var type18 ExternalServiceLevel1Code
	assert.NotNil(t, type18.Validate())
	type18 = "test"
	assert.Nil(t, type18.Validate())

	var type19 ExternalTaxAmountType1Code
	assert.NotNil(t, type19.Validate())
	type19 = "test"
	assert.Nil(t, type19.Validate())

	var type24 AdviceType1Code
	assert.NotNil(t, type24.Validate())
	type24 = "test"
	assert.NotNil(t, type24.Validate())
	type24 = "ADWD"
	assert.Nil(t, type24.Validate())

	var type25 ChargeBearerType1Code
	assert.NotNil(t, type25.Validate())
	type25 = "test"
	assert.NotNil(t, type25.Validate())
	type25 = "DEBT"
	assert.Nil(t, type25.Validate())

	var type28 DocumentType3Code
	assert.NotNil(t, type28.Validate())
	type28 = "test"
	assert.NotNil(t, type28.Validate())
	type28 = "RADM"
	assert.Nil(t, type28.Validate())

	var type29 DocumentType6Code
	assert.NotNil(t, type29.Validate())
	type29 = "test"
	assert.NotNil(t, type29.Validate())
	type29 = "MSIN"
	assert.Nil(t, type29.Validate())

	var type30 Frequency6Code
	assert.NotNil(t, type30.Validate())
	type30 = "test"
	assert.NotNil(t, type30.Validate())
	type30 = "YEAR"
	assert.Nil(t, type30.Validate())

	var type31 PaymentMethod2Code
	assert.NotNil(t, type31.Validate())
	type31 = "test"
	assert.NotNil(t, type31.Validate())
	type31 = "DD"
	assert.Nil(t, type31.Validate())

	var type32 PreferredContactMethod1Code
	assert.NotNil(t, type32.Validate())
	type32 = "test"
	assert.NotNil(t, type32.Validate())
	type32 = "LETT"
	assert.Nil(t, type32.Validate())

	var type33 Priority2Code
	assert.NotNil(t, type33.Validate())
	type33 = "test"
	assert.NotNil(t, type33.Validate())
	type33 = "HIGH"
	assert.Nil(t, type33.Validate())

	var type34 RegulatoryReportingType1Code
	assert.NotNil(t, type34.Validate())
	type34 = "test"
	assert.NotNil(t, type34.Validate())
	type34 = "CRED"
	assert.Nil(t, type34.Validate())

	var type35 RemittanceLocationMethod2Code
	assert.NotNil(t, type35.Validate())
	type35 = "test"
	assert.NotNil(t, type35.Validate())
	type35 = "FAXI"
	assert.Nil(t, type35.Validate())

	var type36 TaxRecordPeriod1Code
	assert.NotNil(t, type36.Validate())
	type36 = "test"
	assert.NotNil(t, type36.Validate())
	type36 = "MM01"
	assert.Nil(t, type36.Validate())

	var type37 SequenceType3Code
	assert.NotNil(t, type37.Validate())
	type37 = "test"
	assert.NotNil(t, type37.Validate())
	type37 = "FRST"
	assert.Nil(t, type37.Validate())
}
