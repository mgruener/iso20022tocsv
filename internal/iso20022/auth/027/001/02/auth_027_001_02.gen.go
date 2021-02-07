// Code generated by main. DO NOT EDIT.

package auth_027_001_02

import (
	"bytes"
	"encoding/xml"
	"time"
)

// May be one of ADDR, PBOX, HOME, BIZZ, MLTO, DLVY
type AddressType2Code string

type AddressType3Choice struct {
	Cd    AddressType2Code        `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 Prtry"`
}

// Must match the pattern [A-Z0-9]{4,4}[A-Z]{2,2}[A-Z0-9]{2,2}([A-Z0-9]{3,3}){0,1}
type AnyBICDec2014Identifier string

// Must match the pattern [A-Z0-9]{4,4}[A-Z]{2,2}[A-Z0-9]{2,2}([A-Z0-9]{3,3}){0,1}
type BICFIDec2014Identifier string

type BranchAndFinancialInstitutionIdentification6 struct {
	FinInstnId FinancialInstitutionIdentification18 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 FinInstnId"`
	BrnchId    BranchData3                          `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 BrnchId,omitempty"`
}

type BranchData3 struct {
	Id      Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 Id,omitempty"`
	LEI     LEIIdentifier   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 LEI,omitempty"`
	Nm      Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 Nm,omitempty"`
	PstlAdr PostalAddress24 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 PstlAdr,omitempty"`
}

type ClearingSystemIdentification2Choice struct {
	Cd    ExternalClearingSystemIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 Cd"`
	Prtry Max35Text                                 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 Prtry"`
}

type ClearingSystemMemberIdentification2 struct {
	ClrSysId ClearingSystemIdentification2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 ClrSysId,omitempty"`
	MmbId    Max35Text                           `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 MmbId"`
}

type Contact4 struct {
	NmPrfx    NamePrefix2Code             `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 NmPrfx,omitempty"`
	Nm        Max140Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 Nm,omitempty"`
	PhneNb    PhoneNumber                 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 PhneNb,omitempty"`
	MobNb     PhoneNumber                 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 MobNb,omitempty"`
	FaxNb     PhoneNumber                 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 FaxNb,omitempty"`
	EmailAdr  Max2048Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 EmailAdr,omitempty"`
	EmailPurp Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 EmailPurp,omitempty"`
	JobTitl   Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 JobTitl,omitempty"`
	Rspnsblty Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 Rspnsblty,omitempty"`
	Dept      Max70Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 Dept,omitempty"`
	Othr      []OtherContact1             `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 Othr,omitempty"`
	PrefrdMtd PreferredContactMethod1Code `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 PrefrdMtd,omitempty"`
}

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type CurrencyControlGroupStatus2 struct {
	OrgnlRefs OriginalMessage5                             `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 OrgnlRefs"`
	RptgPty   TradeParty5                                  `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 RptgPty"`
	RegnAgt   BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 RegnAgt"`
	RptgPrd   Period4Choice                                `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 RptgPrd,omitempty"`
	Sts       StatisticalReportingStatus1Code              `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 Sts,omitempty"`
	StsRsn    []ValidationStatusReason2                    `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 StsRsn,omitempty"`
	StsDtTm   ISODateTime                                  `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 StsDtTm,omitempty"`
}

type CurrencyControlHeader6 struct {
	MsgId    Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 MsgId"`
	CreDtTm  ISODateTime                                  `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 CreDtTm"`
	NbOfItms Max15NumericText                             `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 NbOfItms"`
	RcvgPty  PartyIdentification135                       `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 RcvgPty"`
	RegnAgt  BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 RegnAgt"`
}

type CurrencyControlPackageStatus2 struct {
	PackgId Max35Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 PackgId"`
	Sts     StatisticalReportingStatus1Code `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 Sts"`
	StsRsn  []ValidationStatusReason2       `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 StsRsn,omitempty"`
	StsDtTm ISODateTime                     `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 StsDtTm,omitempty"`
	RcrdSts []CurrencyControlRecordStatus2  `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 RcrdSts,omitempty"`
}

type CurrencyControlRecordStatus2 struct {
	RcrdId  Max35Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 RcrdId"`
	Sts     StatisticalReportingStatus1Code `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 Sts"`
	StsRsn  []ValidationStatusReason2       `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 StsRsn,omitempty"`
	StsDtTm ISODateTime                     `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 StsDtTm,omitempty"`
	DocId   DocumentIdentification28        `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 DocId,omitempty"`
}

type CurrencyControlStatusAdviceV02 struct {
	GrpHdr      CurrencyControlHeader6          `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 GrpHdr"`
	GrpSts      []CurrencyControlGroupStatus2   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 GrpSts"`
	PackgSts    []CurrencyControlPackageStatus2 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 PackgSts,omitempty"`
	SplmtryData []SupplementaryData1            `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 SplmtryData,omitempty"`
}

type DateAndPlaceOfBirth1 struct {
	BirthDt     ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 BirthDt"`
	PrvcOfBirth Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 PrvcOfBirth,omitempty"`
	CityOfBirth Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 CityOfBirth"`
	CtryOfBirth CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 CtryOfBirth"`
}

type Document struct {
	CcyCtrlStsAdvc CurrencyControlStatusAdviceV02 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 CcyCtrlStsAdvc"`
}

type DocumentIdentification28 struct {
	Id       Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 Id,omitempty"`
	DtOfIsse ISODate   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 DtOfIsse"`
}

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

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
	Cd    ExternalFinancialInstitutionIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 Cd"`
	Prtry Max35Text                                       `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 Prtry"`
}

type FinancialInstitutionIdentification18 struct {
	BICFI       BICFIDec2014Identifier              `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 BICFI,omitempty"`
	ClrSysMmbId ClearingSystemMemberIdentification2 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 ClrSysMmbId,omitempty"`
	LEI         LEIIdentifier                       `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 LEI,omitempty"`
	Nm          Max140Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 Nm,omitempty"`
	PstlAdr     PostalAddress24                     `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 PstlAdr,omitempty"`
	Othr        GenericFinancialIdentification1     `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 Othr,omitempty"`
}

type GenericFinancialIdentification1 struct {
	Id      Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 Id"`
	SchmeNm FinancialIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 SchmeNm,omitempty"`
	Issr    Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 Issr,omitempty"`
}

type GenericIdentification30 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 Id"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 SchmeNm,omitempty"`
}

type GenericOrganisationIdentification1 struct {
	Id      Max35Text                                   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 Id"`
	SchmeNm OrganisationIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 SchmeNm,omitempty"`
	Issr    Max35Text                                   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 Issr,omitempty"`
}

type GenericPersonIdentification1 struct {
	Id      Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 Id"`
	SchmeNm PersonIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 SchmeNm,omitempty"`
	Issr    Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 Issr,omitempty"`
}

type GenericValidationRuleIdentification1 struct {
	Id      Max35Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 Id"`
	Desc    Max350Text                      `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 Desc,omitempty"`
	SchmeNm ValidationRuleSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 SchmeNm,omitempty"`
	Issr    Max35Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 Issr,omitempty"`
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

// Must match the pattern [A-Z0-9]{18,18}[0-9]{2,2}
type LEIIdentifier string

type LegalOrganisation2 struct {
	Id           Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 Id,omitempty"`
	Nm           Max140Text `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 Nm,omitempty"`
	EstblishmtDt ISODate    `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 EstblishmtDt,omitempty"`
	RegnDt       ISODate    `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 RegnDt,omitempty"`
}

// Must be at least 1 items long
type Max105Text string

// Must be at least 1 items long
type Max128Text string

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
type Max4Text string

// Must be at least 1 items long
type Max70Text string

// May be one of DOCT, MADM, MISS, MIST, MIKS
type NamePrefix2Code string

type OrganisationIdentification29 struct {
	AnyBIC AnyBICDec2014Identifier              `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 AnyBIC,omitempty"`
	LEI    LEIIdentifier                        `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 LEI,omitempty"`
	Othr   []GenericOrganisationIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 Othr,omitempty"`
}

type OrganisationIdentificationSchemeName1Choice struct {
	Cd    ExternalOrganisationIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 Cd"`
	Prtry Max35Text                               `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 Prtry"`
}

type OriginalMessage5 struct {
	OrgnlSndr    Party40Choice `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 OrgnlSndr,omitempty"`
	OrgnlMsgId   Max35Text     `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 OrgnlMsgId"`
	OrgnlMsgNmId Max35Text     `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 OrgnlMsgNmId"`
	OrgnlCreDtTm ISODateTime   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 OrgnlCreDtTm,omitempty"`
}

type OtherContact1 struct {
	ChanlTp Max4Text   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 ChanlTp"`
	Id      Max128Text `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 Id,omitempty"`
}

type Party38Choice struct {
	OrgId  OrganisationIdentification29 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 OrgId"`
	PrvtId PersonIdentification13       `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 PrvtId"`
}

type Party40Choice struct {
	Pty PartyIdentification135                       `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 Pty"`
	Agt BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 Agt"`
}

type PartyIdentification135 struct {
	Nm        Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 Nm,omitempty"`
	PstlAdr   PostalAddress24 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 PstlAdr,omitempty"`
	Id        Party38Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 Id,omitempty"`
	CtryOfRes CountryCode     `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 CtryOfRes,omitempty"`
	CtctDtls  Contact4        `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 CtctDtls,omitempty"`
}

type Period2 struct {
	FrDt ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 FrDt"`
	ToDt ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 ToDt"`
}

type Period4Choice struct {
	Dt       ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 Dt"`
	FrDt     ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 FrDt"`
	ToDt     ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 ToDt"`
	FrDtToDt Period2 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 FrDtToDt"`
}

type PersonIdentification13 struct {
	DtAndPlcOfBirth DateAndPlaceOfBirth1           `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 DtAndPlcOfBirth,omitempty"`
	Othr            []GenericPersonIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 Othr,omitempty"`
}

type PersonIdentificationSchemeName1Choice struct {
	Cd    ExternalPersonIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 Cd"`
	Prtry Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 Prtry"`
}

// Must match the pattern \+[0-9]{1,3}-[0-9()+\-]{1,30}
type PhoneNumber string

type PostalAddress24 struct {
	AdrTp       AddressType3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 AdrTp,omitempty"`
	Dept        Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 Dept,omitempty"`
	SubDept     Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 SubDept,omitempty"`
	StrtNm      Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 StrtNm,omitempty"`
	BldgNb      Max16Text          `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 BldgNb,omitempty"`
	BldgNm      Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 BldgNm,omitempty"`
	Flr         Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 Flr,omitempty"`
	PstBx       Max16Text          `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 PstBx,omitempty"`
	Room        Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 Room,omitempty"`
	PstCd       Max16Text          `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 PstCd,omitempty"`
	TwnNm       Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 TwnNm,omitempty"`
	TwnLctnNm   Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 TwnLctnNm,omitempty"`
	DstrctNm    Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 DstrctNm,omitempty"`
	CtrySubDvsn Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 CtrySubDvsn,omitempty"`
	Ctry        CountryCode        `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 Ctry,omitempty"`
	AdrLine     []Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 AdrLine,omitempty"`
}

// May be one of LETT, MAIL, PHON, FAXX, CELL
type PreferredContactMethod1Code string

// May be one of ACPT, ACTC, PART, PDNG, RCVD, RJCT, RMDR, INCF, CRPT
type StatisticalReportingStatus1Code string

type StatusReason6Choice struct {
	Cd    ExternalStatusReason1Code `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 Cd"`
	Prtry Max35Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 Prtry"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

// May be one of NONE, MASA, MISA, SISA, IISA, CUYP, PRYP, ASTR, EMPY, EMCY, EPRY, ECYE, NFPI, NFQP, DECP, IRAC, IRAR, KEOG, PFSP, 401K, SIRA, 403B, 457X, RIRA, RIAN, RCRF, RCIP, EIFP, EIOP
type TaxExemptReason1Code string

type TaxExemptionReasonFormat1Choice struct {
	Ustrd Max140Text           `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 Ustrd"`
	Strd  TaxExemptReason1Code `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 Strd"`
}

type TaxParty4 struct {
	TaxId       Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 TaxId,omitempty"`
	TaxTp       Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 TaxTp,omitempty"`
	RegnId      Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 RegnId,omitempty"`
	TaxXmptnRsn []TaxExemptionReasonFormat1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 TaxXmptnRsn,omitempty"`
}

type TradeParty5 struct {
	PtyId  PartyIdentification135 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 PtyId"`
	LglOrg LegalOrganisation2     `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 LglOrg,omitempty"`
	TaxPty []TaxParty4            `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 TaxPty,omitempty"`
}

type ValidationRuleSchemeName1Choice struct {
	Cd    ExternalValidationRuleIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 Cd"`
	Prtry Max35Text                                 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 Prtry"`
}

type ValidationStatusReason2 struct {
	Orgtr     PartyIdentification135                 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 Orgtr,omitempty"`
	Rsn       StatusReason6Choice                    `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 Rsn,omitempty"`
	VldtnRule []GenericValidationRuleIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 VldtnRule,omitempty"`
	AddtlInf  []Max105Text                           `xml:"urn:iso:std:iso:20022:tech:xsd:auth.027.001.02 AddtlInf,omitempty"`
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
