// Code generated by main. DO NOT EDIT.

package auth_025_001_02

import (
	"bytes"
	"encoding/base64"
	"encoding/xml"
	"time"
)

type ActiveCurrencyAndAmount struct {
	Value float64            `xml:",chardata"`
	Ccy   ActiveCurrencyCode `xml:"Ccy,attr"`
}

// Must match the pattern [A-Z]{3,3}
type ActiveCurrencyCode string

// May be one of ADDR, PBOX, HOME, BIZZ, MLTO, DLVY
type AddressType2Code string

type AddressType3Choice struct {
	Cd    AddressType2Code        `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 Prtry"`
}

// Must match the pattern [A-Z0-9]{4,4}[A-Z]{2,2}[A-Z0-9]{2,2}([A-Z0-9]{3,3}){0,1}
type AnyBICDec2014Identifier string

// Must match the pattern [A-Z0-9]{4,4}[A-Z]{2,2}[A-Z0-9]{2,2}([A-Z0-9]{3,3}){0,1}
type BICFIDec2014Identifier string

type BinaryFile1 struct {
	MIMETp         Max35Text     `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 MIMETp,omitempty"`
	NcodgTp        Max35Text     `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 NcodgTp,omitempty"`
	CharSet        Max35Text     `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 CharSet,omitempty"`
	InclBinryObjct Max100KBinary `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 InclBinryObjct,omitempty"`
}

type BranchAndFinancialInstitutionIdentification6 struct {
	FinInstnId FinancialInstitutionIdentification18 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 FinInstnId"`
	BrnchId    BranchData3                          `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 BrnchId,omitempty"`
}

type BranchData3 struct {
	Id      Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 Id,omitempty"`
	LEI     LEIIdentifier   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 LEI,omitempty"`
	Nm      Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 Nm,omitempty"`
	PstlAdr PostalAddress24 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 PstlAdr,omitempty"`
}

type ClearingSystemIdentification2Choice struct {
	Cd    ExternalClearingSystemIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 Cd"`
	Prtry Max35Text                                 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 Prtry"`
}

type ClearingSystemMemberIdentification2 struct {
	ClrSysId ClearingSystemIdentification2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 ClrSysId,omitempty"`
	MmbId    Max35Text                           `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 MmbId"`
}

type Contact4 struct {
	NmPrfx    NamePrefix2Code             `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 NmPrfx,omitempty"`
	Nm        Max140Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 Nm,omitempty"`
	PhneNb    PhoneNumber                 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 PhneNb,omitempty"`
	MobNb     PhoneNumber                 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 MobNb,omitempty"`
	FaxNb     PhoneNumber                 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 FaxNb,omitempty"`
	EmailAdr  Max2048Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 EmailAdr,omitempty"`
	EmailPurp Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 EmailPurp,omitempty"`
	JobTitl   Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 JobTitl,omitempty"`
	Rspnsblty Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 Rspnsblty,omitempty"`
	Dept      Max70Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 Dept,omitempty"`
	Othr      []OtherContact1             `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 Othr,omitempty"`
	PrefrdMtd PreferredContactMethod1Code `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 PrefrdMtd,omitempty"`
}

type ContractRegistrationReference1Choice struct {
	RegdCtrctId Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 RegdCtrctId"`
	Ctrct       DocumentIdentification28 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 Ctrct"`
}

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type CurrencyControlHeader5 struct {
	MsgId    Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 MsgId"`
	CreDtTm  ISODateTime                                  `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 CreDtTm"`
	NbOfItms Max15NumericText                             `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 NbOfItms"`
	InitgPty Party40Choice                                `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 InitgPty"`
	FwdgAgt  BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 FwdgAgt,omitempty"`
}

type CurrencyControlSupportingDocumentDeliveryV02 struct {
	GrpHdr      CurrencyControlHeader5 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 GrpHdr"`
	SpprtgDoc   []SupportingDocument2  `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 SpprtgDoc"`
	SplmtryData []SupplementaryData1   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 SplmtryData,omitempty"`
}

type DateAndPlaceOfBirth1 struct {
	BirthDt     ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 BirthDt"`
	PrvcOfBirth Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 PrvcOfBirth,omitempty"`
	CityOfBirth Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 CityOfBirth"`
	CtryOfBirth CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 CtryOfBirth"`
}

type Document struct {
	CcyCtrlSpprtgDocDlvry CurrencyControlSupportingDocumentDeliveryV02 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 CcyCtrlSpprtgDocDlvry"`
}

type DocumentAmendment1 struct {
	CrrctnId   float64   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 CrrctnId"`
	OrgnlDocId Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 OrgnlDocId,omitempty"`
}

type DocumentGeneralInformation3 struct {
	DocTp           ExternalDocumentType1Code  `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 DocTp"`
	DocNb           Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 DocNb"`
	SndrRcvrSeqId   Max140Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 SndrRcvrSeqId,omitempty"`
	IsseDt          ISODate                    `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 IsseDt,omitempty"`
	URL             Max256Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 URL,omitempty"`
	LkFileHash      SignatureEnvelopeReference `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 LkFileHash,omitempty"`
	AttchdBinryFile BinaryFile1                `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 AttchdBinryFile"`
}

type DocumentIdentification22 struct {
	Id       Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 Id"`
	DtOfIsse ISODate   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 DtOfIsse,omitempty"`
}

type DocumentIdentification28 struct {
	Id       Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 Id,omitempty"`
	DtOfIsse ISODate   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 DtOfIsse"`
}

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

// Must be at least 1 items long
type ExternalClearingSystemIdentification1Code string

// Must be at least 1 items long
type ExternalDocumentType1Code string

// Must be at least 1 items long
type ExternalFinancialInstitutionIdentification1Code string

// Must be at least 1 items long
type ExternalOrganisationIdentification1Code string

// Must be at least 1 items long
type ExternalPersonIdentification1Code string

// Must be at least 1 items long
type ExternalShipmentCondition1Code string

type FinancialIdentificationSchemeName1Choice struct {
	Cd    ExternalFinancialInstitutionIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 Cd"`
	Prtry Max35Text                                       `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 Prtry"`
}

type FinancialInstitutionIdentification18 struct {
	BICFI       BICFIDec2014Identifier              `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 BICFI,omitempty"`
	ClrSysMmbId ClearingSystemMemberIdentification2 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 ClrSysMmbId,omitempty"`
	LEI         LEIIdentifier                       `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 LEI,omitempty"`
	Nm          Max140Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 Nm,omitempty"`
	PstlAdr     PostalAddress24                     `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 PstlAdr,omitempty"`
	Othr        GenericFinancialIdentification1     `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 Othr,omitempty"`
}

type GenericFinancialIdentification1 struct {
	Id      Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 Id"`
	SchmeNm FinancialIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 SchmeNm,omitempty"`
	Issr    Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 Issr,omitempty"`
}

type GenericIdentification30 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 Id"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 SchmeNm,omitempty"`
}

type GenericOrganisationIdentification1 struct {
	Id      Max35Text                                   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 Id"`
	SchmeNm OrganisationIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 SchmeNm,omitempty"`
	Issr    Max35Text                                   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 Issr,omitempty"`
}

type GenericPersonIdentification1 struct {
	Id      Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 Id"`
	SchmeNm PersonIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 SchmeNm,omitempty"`
	Issr    Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 Issr,omitempty"`
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

type Max100KBinary []byte

func (t *Max100KBinary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Max100KBinary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

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
type Max256Text string

// Must be at least 1 items long
type Max350Text string

// Must be at least 1 items long
type Max35Text string

// Must be at least 1 items long
type Max4Text string

// Must be at least 1 items long
type Max500Text string

// Must be at least 1 items long
type Max70Text string

// May be one of DOCT, MADM, MISS, MIST, MIKS
type NamePrefix2Code string

type OrganisationIdentification29 struct {
	AnyBIC AnyBICDec2014Identifier              `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 AnyBIC,omitempty"`
	LEI    LEIIdentifier                        `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 LEI,omitempty"`
	Othr   []GenericOrganisationIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 Othr,omitempty"`
}

type OrganisationIdentificationSchemeName1Choice struct {
	Cd    ExternalOrganisationIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 Cd"`
	Prtry Max35Text                               `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 Prtry"`
}

type OtherContact1 struct {
	ChanlTp Max4Text   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 ChanlTp"`
	Id      Max128Text `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 Id,omitempty"`
}

type Party38Choice struct {
	OrgId  OrganisationIdentification29 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 OrgId"`
	PrvtId PersonIdentification13       `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 PrvtId"`
}

type Party40Choice struct {
	Pty PartyIdentification135                       `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 Pty"`
	Agt BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 Agt"`
}

type PartyIdentification135 struct {
	Nm        Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 Nm,omitempty"`
	PstlAdr   PostalAddress24 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 PstlAdr,omitempty"`
	Id        Party38Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 Id,omitempty"`
	CtryOfRes CountryCode     `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 CtryOfRes,omitempty"`
	CtctDtls  Contact4        `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 CtctDtls,omitempty"`
}

type PersonIdentification13 struct {
	DtAndPlcOfBirth DateAndPlaceOfBirth1           `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 DtAndPlcOfBirth,omitempty"`
	Othr            []GenericPersonIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 Othr,omitempty"`
}

type PersonIdentificationSchemeName1Choice struct {
	Cd    ExternalPersonIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 Cd"`
	Prtry Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 Prtry"`
}

// Must match the pattern \+[0-9]{1,3}-[0-9()+\-]{1,30}
type PhoneNumber string

type PostalAddress24 struct {
	AdrTp       AddressType3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 AdrTp,omitempty"`
	Dept        Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 Dept,omitempty"`
	SubDept     Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 SubDept,omitempty"`
	StrtNm      Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 StrtNm,omitempty"`
	BldgNb      Max16Text          `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 BldgNb,omitempty"`
	BldgNm      Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 BldgNm,omitempty"`
	Flr         Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 Flr,omitempty"`
	PstBx       Max16Text          `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 PstBx,omitempty"`
	Room        Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 Room,omitempty"`
	PstCd       Max16Text          `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 PstCd,omitempty"`
	TwnNm       Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 TwnNm,omitempty"`
	TwnLctnNm   Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 TwnLctnNm,omitempty"`
	DstrctNm    Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 DstrctNm,omitempty"`
	CtrySubDvsn Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 CtrySubDvsn,omitempty"`
	Ctry        CountryCode        `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 Ctry,omitempty"`
	AdrLine     []Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 AdrLine,omitempty"`
}

// May be one of LETT, MAIL, PHON, FAXX, CELL
type PreferredContactMethod1Code string

type ShipmentAttribute1 struct {
	Conds         ExternalShipmentCondition1Code `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 Conds,omitempty"`
	XpctdDt       ISODate                        `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 XpctdDt,omitempty"`
	CtryOfCntrPty CountryCode                    `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 CtryOfCntrPty"`
}

type SignatureEnvelopeReference struct {
	Item string `xml:",any"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

type SupportingDocument2 struct {
	SpprtgDocId Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 SpprtgDocId"`
	OrgnlReqId  Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 OrgnlReqId,omitempty"`
	Cert        DocumentIdentification28                     `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 Cert"`
	AcctOwnr    PartyIdentification135                       `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 AcctOwnr"`
	AcctSvcr    BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 AcctSvcr"`
	Amdmnt      DocumentAmendment1                           `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 Amdmnt,omitempty"`
	CtrctRef    ContractRegistrationReference1Choice         `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 CtrctRef"`
	Ntry        []SupportingDocumentEntry1                   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 Ntry"`
	SplmtryData []SupplementaryData1                         `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 SplmtryData,omitempty"`
}

type SupportingDocumentEntry1 struct {
	NtryId                      Max35Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 NtryId"`
	OrgnlDoc                    DocumentIdentification22      `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 OrgnlDoc"`
	DocTp                       Exact4AlphaNumericText        `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 DocTp"`
	TtlAmt                      ActiveCurrencyAndAmount       `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 TtlAmt,omitempty"`
	TtlAmtAftrShipmnt           ActiveCurrencyAndAmount       `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 TtlAmtAftrShipmnt,omitempty"`
	TtlAmtInCtrctCcy            ActiveCurrencyAndAmount       `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 TtlAmtInCtrctCcy,omitempty"`
	TtlAmtAftrShipmntInCtrctCcy ActiveCurrencyAndAmount       `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 TtlAmtAftrShipmntInCtrctCcy,omitempty"`
	ShipmntAttrbts              ShipmentAttribute1            `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 ShipmntAttrbts"`
	AddtlInf                    Max500Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 AddtlInf,omitempty"`
	Attchmnt                    []DocumentGeneralInformation3 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.02 Attchmnt,omitempty"`
}

type xsdBase64Binary []byte

func (b *xsdBase64Binary) UnmarshalText(text []byte) (err error) {
	*b, err = base64.StdEncoding.DecodeString(string(text))
	return
}
func (b xsdBase64Binary) MarshalText() ([]byte, error) {
	var buf bytes.Buffer
	enc := base64.NewEncoder(base64.StdEncoding, &buf)
	enc.Write([]byte(b))
	enc.Close()
	return buf.Bytes(), nil
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
