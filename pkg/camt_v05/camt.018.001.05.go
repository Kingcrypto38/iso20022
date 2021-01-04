// Code generated by xsdgen. DO NOT EDIT.

package camt_v05

type BusinessDayCriteria2 struct {
	NewQryNm Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.018.001.05 NewQryNm,omitempty"`
	SchCrit  []BusinessDaySearchCriteria2 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.018.001.05 SchCrit,omitempty"`
	RtrCrit  BusinessDayReturnCriteria2   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.018.001.05 RtrCrit,omitempty"`
}

type BusinessDayCriteria3Choice struct {
	QryNm   Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.018.001.05 QryNm"`
	NewCrit BusinessDayCriteria2 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.018.001.05 NewCrit"`
}

type BusinessDayQuery2 struct {
	QryTp QueryType2Code             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.018.001.05 QryTp,omitempty"`
	Crit  BusinessDayCriteria3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.018.001.05 Crit,omitempty"`
}

type BusinessDayReturnCriteria2 struct {
	SysDtInd   bool `xml:"urn:iso:std:iso:20022:tech:xsd:camt.018.001.05 SysDtInd,omitempty"`
	SysStsInd  bool `xml:"urn:iso:std:iso:20022:tech:xsd:camt.018.001.05 SysStsInd,omitempty"`
	SysCcyInd  bool `xml:"urn:iso:std:iso:20022:tech:xsd:camt.018.001.05 SysCcyInd,omitempty"`
	ClsrPrdInd bool `xml:"urn:iso:std:iso:20022:tech:xsd:camt.018.001.05 ClsrPrdInd,omitempty"`
	EvtInd     bool `xml:"urn:iso:std:iso:20022:tech:xsd:camt.018.001.05 EvtInd,omitempty"`
	SsnPrdInd  bool `xml:"urn:iso:std:iso:20022:tech:xsd:camt.018.001.05 SsnPrdInd,omitempty"`
	EvtTpInd   bool `xml:"urn:iso:std:iso:20022:tech:xsd:camt.018.001.05 EvtTpInd,omitempty"`
}

type BusinessDaySearchCriteria2 struct {
	SysDt   ISODate                       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.018.001.05 SysDt,omitempty"`
	SysId   []SystemIdentification2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.018.001.05 SysId,omitempty"`
	SysCcy  []ActiveCurrencyCode          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.018.001.05 SysCcy,omitempty"`
	EvtTp   SystemEventType2Choice        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.018.001.05 EvtTp,omitempty"`
	ClsrPrd DateTimePeriod1Choice         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.018.001.05 ClsrPrd,omitempty"`
}

type DateTimePeriod1 struct {
	FrDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:camt.018.001.05 FrDtTm"`
	ToDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:camt.018.001.05 ToDtTm"`
}

type DateTimePeriod1Choice struct {
	FrDtTm ISODateTime     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.018.001.05 FrDtTm"`
	ToDtTm ISODateTime     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.018.001.05 ToDtTm"`
	DtTmRg DateTimePeriod1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.018.001.05 DtTmRg"`
}

type Document struct {
	GetBizDayInf GetBusinessDayInformationV05 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.018.001.05 GetBizDayInf"`
}

// Must be at least 1 items long
type ExternalEnquiryRequestType1Code string

// Must be at least 1 items long
type ExternalMarketInfrastructure1Code string

// Must be at least 1 items long
type ExternalPaymentControlRequestType1Code string

type GenericIdentification1 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.018.001.05 Id"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.018.001.05 SchmeNm,omitempty"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:camt.018.001.05 Issr,omitempty"`
}

type GetBusinessDayInformationV05 struct {
	MsgHdr          MessageHeader9       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.018.001.05 MsgHdr"`
	BizDayInfQryDef BusinessDayQuery2    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.018.001.05 BizDayInfQryDef,omitempty"`
	SplmtryData     []SupplementaryData1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.018.001.05 SplmtryData,omitempty"`
}

type MarketInfrastructureIdentification1Choice struct {
	Cd    ExternalMarketInfrastructure1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.018.001.05 Cd"`
	Prtry Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:camt.018.001.05 Prtry"`
}

type MessageHeader9 struct {
	MsgId   Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.018.001.05 MsgId"`
	CreDtTm ISODateTime        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.018.001.05 CreDtTm,omitempty"`
	ReqTp   RequestType4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.018.001.05 ReqTp,omitempty"`
}

// May be one of ALLL, CHNG, MODF, DELD
type QueryType2Code string

type RequestType4Choice struct {
	PmtCtrl ExternalPaymentControlRequestType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.018.001.05 PmtCtrl"`
	Enqry   ExternalEnquiryRequestType1Code        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.018.001.05 Enqry"`
	Prtry   GenericIdentification1                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.018.001.05 Prtry"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.018.001.05 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.018.001.05 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

type SystemEventType2Choice struct {
	Cd    SystemEventType2Code   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.018.001.05 Cd"`
	Prtry GenericIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.018.001.05 Prtry"`
}

// May be one of LVCO, LVCC, LVRT, EUSU, STSU, LWSU, EUCO, FIRE, STDY, LTNC, CRCO, RECC, LTGC, LTDC, CUSC, IBKC, SYSC, SSSC, REOP, PCOT, NPCT, ESTF
type SystemEventType2Code string

type SystemIdentification2Choice struct {
	MktInfrstrctrId MarketInfrastructureIdentification1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.018.001.05 MktInfrstrctrId"`
	Ctry            CountryCode                               `xml:"urn:iso:std:iso:20022:tech:xsd:camt.018.001.05 Ctry"`
}
