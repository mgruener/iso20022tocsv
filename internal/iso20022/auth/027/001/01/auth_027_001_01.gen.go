// Code generated by main. DO NOT EDIT.

package auth_027_001_01

import (
	"bytes"
	"encoding/xml"
	"time"
)

// May be one of ADDR, PBOX, HOME, BIZZ, MLTO, DLVY
type AddressType2Code string

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type AnyBICIdentifier string

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type BICFIIdentifier string

type BranchAndFinancialInstitutionIdentification5 struct {
	FinInstnId FinancialInstitutionIdentification8 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 FinInstnId"`
	BrnchId    BranchData2                         `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 BrnchId,omitempty"`
}

type BranchData2 struct {
	Id      Max35Text      `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 Id,omitempty"`
	Nm      Max140Text     `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 Nm,omitempty"`
	PstlAdr PostalAddress6 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 PstlAdr,omitempty"`
}

type ClearingSystemIdentification2Choice struct {
	Cd    ExternalClearingSystemIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 Cd"`
	Prtry Max35Text                                 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 Prtry"`
}

type ClearingSystemMemberIdentification2 struct {
	ClrSysId ClearingSystemIdentification2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 ClrSysId,omitempty"`
	MmbId    Max35Text                           `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 MmbId"`
}

type ContactDetails2 struct {
	NmPrfx   NamePrefix1Code `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 NmPrfx,omitempty"`
	Nm       Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 Nm,omitempty"`
	PhneNb   PhoneNumber     `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 PhneNb,omitempty"`
	MobNb    PhoneNumber     `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 MobNb,omitempty"`
	FaxNb    PhoneNumber     `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 FaxNb,omitempty"`
	EmailAdr Max2048Text     `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 EmailAdr,omitempty"`
	Othr     Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 Othr,omitempty"`
}

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type CurrencyControlGroupStatus1 struct {
	OrgnlRefs OriginalMessage3                             `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 OrgnlRefs"`
	RptgPty   TradeParty2                                  `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 RptgPty"`
	RegnAgt   BranchAndFinancialInstitutionIdentification5 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 RegnAgt"`
	RptgPrd   Period4Choice                                `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 RptgPrd,omitempty"`
	Sts       StatisticalReportingStatus1Code              `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 Sts,omitempty"`
	StsRsn    []ValidationStatusReason1                    `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 StsRsn,omitempty"`
	StsDtTm   ISODateTime                                  `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 StsDtTm,omitempty"`
}

type CurrencyControlHeader2 struct {
	MsgId    Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 MsgId"`
	CreDtTm  ISODateTime                                  `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 CreDtTm"`
	NbOfItms Max15NumericText                             `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 NbOfItms"`
	RcvgPty  PartyIdentification77                        `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 RcvgPty"`
	RegnAgt  BranchAndFinancialInstitutionIdentification5 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 RegnAgt"`
}

type CurrencyControlPackageStatus1 struct {
	PackgId Max35Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 PackgId"`
	Sts     StatisticalReportingStatus1Code `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 Sts"`
	StsRsn  []ValidationStatusReason1       `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 StsRsn,omitempty"`
	StsDtTm ISODateTime                     `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 StsDtTm,omitempty"`
	RcrdSts []CurrencyControlRecordStatus1  `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 RcrdSts,omitempty"`
}

type CurrencyControlRecordStatus1 struct {
	RcrdId  Max35Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 RcrdId"`
	Sts     StatisticalReportingStatus1Code `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 Sts"`
	StsRsn  []ValidationStatusReason1       `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 StsRsn,omitempty"`
	StsDtTm ISODateTime                     `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 StsDtTm,omitempty"`
	DocId   DocumentIdentification28        `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 DocId,omitempty"`
}

type CurrencyControlStatusAdviceV01 struct {
	GrpHdr      CurrencyControlHeader2          `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 GrpHdr"`
	GrpSts      []CurrencyControlGroupStatus1   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 GrpSts"`
	PackgSts    []CurrencyControlPackageStatus1 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 PackgSts,omitempty"`
	SplmtryData []SupplementaryData1            `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 SplmtryData,omitempty"`
}

type DateAndPlaceOfBirth struct {
	BirthDt     ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 BirthDt"`
	PrvcOfBirth Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 PrvcOfBirth,omitempty"`
	CityOfBirth Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 CityOfBirth"`
	CtryOfBirth CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 CtryOfBirth"`
}

type Document struct {
	CcyCtrlStsAdvc CurrencyControlStatusAdviceV01 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 CcyCtrlStsAdvc"`
}

type DocumentIdentification28 struct {
	Id       Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 Id,omitempty"`
	DtOfIsse ISODate   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 DtOfIsse"`
}

// Must be at least 1 items long
type ExternalClearingSystemIdentification1Code string

// Must be at least 1 items long
type ExternalFinancialInstitutionIdentification1Code string

// Must be at least 1 items long
type ExternalOrganisationIdentification1Code string

// Must be at least 1 items long
type ExternalPersonIdentification1Code string

// Must be at least 1 items long
type ExternalStatusReason1Code string

// Must be at least 1 items long
type ExternalValidationRuleIdentification1Code string

type FinancialIdentificationSchemeName1Choice struct {
	Cd    ExternalFinancialInstitutionIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 Cd"`
	Prtry Max35Text                                       `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 Prtry"`
}

type FinancialInstitutionIdentification8 struct {
	BICFI       BICFIIdentifier                     `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 BICFI,omitempty"`
	ClrSysMmbId ClearingSystemMemberIdentification2 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 ClrSysMmbId,omitempty"`
	Nm          Max140Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 Nm,omitempty"`
	PstlAdr     PostalAddress6                      `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 PstlAdr,omitempty"`
	Othr        GenericFinancialIdentification1     `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 Othr,omitempty"`
}

type GenericFinancialIdentification1 struct {
	Id      Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 Id"`
	SchmeNm FinancialIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 SchmeNm,omitempty"`
	Issr    Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 Issr,omitempty"`
}

type GenericOrganisationIdentification1 struct {
	Id      Max35Text                                   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 Id"`
	SchmeNm OrganisationIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 SchmeNm,omitempty"`
	Issr    Max35Text                                   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 Issr,omitempty"`
}

type GenericPersonIdentification1 struct {
	Id      Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 Id"`
	SchmeNm PersonIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 SchmeNm,omitempty"`
	Issr    Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 Issr,omitempty"`
}

type GenericValidationRuleIdentification1 struct {
	Id      Max35Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 Id"`
	Desc    Max350Text                      `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 Desc,omitempty"`
	SchmeNm ValidationRuleSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 SchmeNm,omitempty"`
	Issr    Max35Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 Issr,omitempty"`
}

type ISODate time.Time

func (t *ISODate) UnmarshalText(text []byte) error {
	return (*xsdDate)(t).UnmarshalText(text)
}
func (t ISODate) MarshalText() ([]byte, error) {
	return xsdDate(t).MarshalText()
}

type ISODateTime time.Time

func (t *ISODateTime) UnmarshalText(text []byte) error {
	return (*xsdDateTime)(t).UnmarshalText(text)
}
func (t ISODateTime) MarshalText() ([]byte, error) {
	return xsdDateTime(t).MarshalText()
}

type LegalOrganisation2 struct {
	Id           Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 Id,omitempty"`
	Nm           Max140Text `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 Nm,omitempty"`
	EstblishmtDt ISODate    `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 EstblishmtDt,omitempty"`
	RegnDt       ISODate    `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 RegnDt,omitempty"`
}

// Must be at least 1 items long
type Max105Text string

// Must be at least 1 items long
type Max140Text string

// Must match the pattern [0-9]{1,15}
type Max15NumericText string

// Must be at least 1 items long
type Max16Text string

// Must be at least 1 items long
type Max2048Text string

// Must be at least 1 items long
type Max350Text string

// Must be at least 1 items long
type Max35Text string

// Must be at least 1 items long
type Max70Text string

// May be one of DOCT, MIST, MISS, MADM
type NamePrefix1Code string

type OrganisationIdentification8 struct {
	AnyBIC AnyBICIdentifier                     `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 AnyBIC,omitempty"`
	Othr   []GenericOrganisationIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 Othr,omitempty"`
}

type OrganisationIdentificationSchemeName1Choice struct {
	Cd    ExternalOrganisationIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 Cd"`
	Prtry Max35Text                               `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 Prtry"`
}

type OriginalMessage3 struct {
	OrgnlSndr    Party28Choice `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 OrgnlSndr,omitempty"`
	OrgnlMsgId   Max35Text     `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 OrgnlMsgId"`
	OrgnlMsgNmId Max35Text     `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 OrgnlMsgNmId"`
	OrgnlCreDtTm ISODateTime   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 OrgnlCreDtTm,omitempty"`
}

type Party11Choice struct {
	OrgId  OrganisationIdentification8 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 OrgId"`
	PrvtId PersonIdentification5       `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 PrvtId"`
}

type Party28Choice struct {
	Pty PartyIdentification77                        `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 Pty"`
	Agt BranchAndFinancialInstitutionIdentification5 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 Agt"`
}

type PartyIdentification77 struct {
	Nm        Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 Nm,omitempty"`
	PstlAdr   PostalAddress19 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 PstlAdr,omitempty"`
	Id        Party11Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 Id,omitempty"`
	CtryOfRes CountryCode     `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 CtryOfRes,omitempty"`
	CtctDtls  ContactDetails2 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 CtctDtls,omitempty"`
}

type Period2 struct {
	FrDt ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 FrDt"`
	ToDt ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 ToDt"`
}

type Period4Choice struct {
	Dt       ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 Dt"`
	FrDt     ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 FrDt"`
	ToDt     ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 ToDt"`
	FrDtToDt Period2 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 FrDtToDt"`
}

type PersonIdentification5 struct {
	DtAndPlcOfBirth DateAndPlaceOfBirth            `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 DtAndPlcOfBirth,omitempty"`
	Othr            []GenericPersonIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 Othr,omitempty"`
}

type PersonIdentificationSchemeName1Choice struct {
	Cd    ExternalPersonIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 Cd"`
	Prtry Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 Prtry"`
}

// Must match the pattern \+[0-9]{1,3}-[0-9()+\-]{1,30}
type PhoneNumber string

type PostalAddress19 struct {
	AdrTp       AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 AdrTp,omitempty"`
	Dept        Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 Dept,omitempty"`
	SubDept     Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 SubDept,omitempty"`
	StrtNm      Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 StrtNm,omitempty"`
	BldgNb      Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 BldgNb,omitempty"`
	PstCd       Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 PstCd,omitempty"`
	TwnNm       Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 TwnNm,omitempty"`
	TwnLctnNm   Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 TwnLctnNm,omitempty"`
	DstrctNm    Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 DstrctNm,omitempty"`
	CtrySubDvsn Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 CtrySubDvsn,omitempty"`
	Ctry        CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 Ctry,omitempty"`
	AdrLine     []Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 AdrLine,omitempty"`
}

type PostalAddress6 struct {
	AdrTp       AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 AdrTp,omitempty"`
	Dept        Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 Dept,omitempty"`
	SubDept     Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 SubDept,omitempty"`
	StrtNm      Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 StrtNm,omitempty"`
	BldgNb      Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 BldgNb,omitempty"`
	PstCd       Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 PstCd,omitempty"`
	TwnNm       Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 TwnNm,omitempty"`
	CtrySubDvsn Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 CtrySubDvsn,omitempty"`
	Ctry        CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 Ctry,omitempty"`
	AdrLine     []Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 AdrLine,omitempty"`
}

// May be one of ACPT, ACTC, PART, PDNG, RCVD, RJCT, RMDR, INCF, CRPT
type StatisticalReportingStatus1Code string

type StatusReason6Choice struct {
	Cd    ExternalStatusReason1Code `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 Cd"`
	Prtry Max35Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 Prtry"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

// May be one of NONE, MASA, MISA, SISA, IISA, CUYP, PRYP, ASTR, EMPY, EMCY, EPRY, ECYE, NFPI, NFQP, DECP, IRAC, IRAR, KEOG, PFSP, 401K, SIRA, 403B, 457X, RIRA, RIAN, RCRF, RCIP, EIFP, EIOP
type TaxExemptReason1Code string

type TaxExemptionReasonFormatChoice struct {
	Ustrd Max140Text           `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 Ustrd"`
	Strd  TaxExemptReason1Code `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 Strd"`
}

type TaxParty3 struct {
	TaxId       Max35Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 TaxId,omitempty"`
	TaxTp       Max35Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 TaxTp,omitempty"`
	RegnId      Max35Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 RegnId,omitempty"`
	TaxXmptnRsn []TaxExemptionReasonFormatChoice `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 TaxXmptnRsn,omitempty"`
}

type TradeParty2 struct {
	PtyId  PartyIdentification77 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 PtyId"`
	LglOrg LegalOrganisation2    `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 LglOrg,omitempty"`
	TaxPty []TaxParty3           `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 TaxPty,omitempty"`
}

type ValidationRuleSchemeName1Choice struct {
	Cd    ExternalValidationRuleIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 Cd"`
	Prtry Max35Text                                 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 Prtry"`
}

type ValidationStatusReason1 struct {
	Orgtr     PartyIdentification77                  `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 Orgtr,omitempty"`
	Rsn       StatusReason6Choice                    `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 Rsn,omitempty"`
	VldtnRule []GenericValidationRuleIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 VldtnRule,omitempty"`
	AddtlInf  []Max105Text                           `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.01 AddtlInf,omitempty"`
}

type xsdDate time.Time

func (t *xsdDate) UnmarshalText(text []byte) error {
	return _unmarshalTime(text, (*time.Time)(t), "2006-01-02")
}
func (t xsdDate) MarshalText() ([]byte, error) {
	return _marshalTime((time.Time)(t), "2006-01-02")
}
func (t xsdDate) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if (time.Time)(t).IsZero() {
		return nil
	}
	m, err := t.MarshalText()
	if err != nil {
		return err
	}
	return e.EncodeElement(m, start)
}
func (t xsdDate) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if (time.Time)(t).IsZero() {
		return xml.Attr{}, nil
	}
	m, err := t.MarshalText()
	return xml.Attr{Name: name, Value: string(m)}, err
}
func _unmarshalTime(text []byte, t *time.Time, format string) (err error) {
	s := string(bytes.TrimSpace(text))
	*t, err = time.Parse(format, s)
	if _, ok := err.(*time.ParseError); ok {
		*t, err = time.Parse(format+"Z07:00", s)
	}
	return err
}
func _marshalTime(t time.Time, format string) ([]byte, error) {
	return []byte(t.Format(format + "Z07:00")), nil
}

type xsdDateTime time.Time

func (t *xsdDateTime) UnmarshalText(text []byte) error {
	return _unmarshalTime(text, (*time.Time)(t), "2006-01-02T15:04:05.999999999")
}
func (t xsdDateTime) MarshalText() ([]byte, error) {
	return _marshalTime((time.Time)(t), "2006-01-02T15:04:05.999999999")
}
func (t xsdDateTime) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if (time.Time)(t).IsZero() {
		return nil
	}
	m, err := t.MarshalText()
	if err != nil {
		return err
	}
	return e.EncodeElement(m, start)
}
func (t xsdDateTime) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if (time.Time)(t).IsZero() {
		return xml.Attr{}, nil
	}
	m, err := t.MarshalText()
	return xml.Attr{Name: name, Value: string(m)}, err
}
