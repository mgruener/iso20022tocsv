// Code generated by main. DO NOT EDIT.

package acmt_015_001_01

import (
	"bytes"
	"encoding/xml"
	"time"
)

type AccountContract2 struct {
	TrgtGoLiveDt ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 TrgtGoLiveDt,omitempty"`
	TrgtClsgDt   ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 TrgtClsgDt,omitempty"`
	UrgcyFlg     bool    `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 UrgcyFlg,omitempty"`
}

type AccountExcludedMandateMaintenanceRequestV01 struct {
	Refs             References4                                  `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 Refs"`
	CtrctDts         AccountContract2                             `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 CtrctDts,omitempty"`
	UndrlygMstrAgrmt ContractDocument1                            `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 UndrlygMstrAgrmt,omitempty"`
	Acct             CustomerAccount1                             `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 Acct"`
	AcctSvcrId       BranchAndFinancialInstitutionIdentification4 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 AcctSvcrId"`
	Org              []Organisation6                              `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 Org"`
	AddtlMsgInf      AdditionalInformation5                       `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 AddtlMsgInf,omitempty"`
	DgtlSgntr        []PartyAndSignature1                         `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 DgtlSgntr,omitempty"`
}

type AccountIdentification4Choice struct {
	IBAN IBAN2007Identifier            `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 IBAN"`
	Othr GenericAccountIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 Othr"`
}

type AccountSchemeName1Choice struct {
	Cd    ExternalAccountIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 Cd"`
	Prtry Max35Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 Prtry"`
}

// May be one of ENAB, DISA, DELE, FORM
type AccountStatus3Code string

// Must match the pattern [A-Z]{3,3}
type ActiveCurrencyCode string

type AdditionalInformation5 struct {
	Inf []Max256Text `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 Inf"`
}

// May be one of ADDR, PBOX, HOME, BIZZ, MLTO, DLVY
type AddressType2Code string

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type AnyBICIdentifier string

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type BICIdentifier string

type BranchAndFinancialInstitutionIdentification4 struct {
	FinInstnId FinancialInstitutionIdentification7 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 FinInstnId"`
	BrnchId    BranchData2                         `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 BrnchId,omitempty"`
}

type BranchData2 struct {
	Id      Max35Text      `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 Id,omitempty"`
	Nm      Max140Text     `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 Nm,omitempty"`
	PstlAdr PostalAddress6 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 PstlAdr,omitempty"`
}

type CashAccountType2 struct {
	Cd    CashAccountType4Code `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 Cd"`
	Prtry Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 Prtry"`
}

// May be one of CASH, CHAR, COMM, TAXE, CISH, TRAS, SACC, CACC, SVGS, ONDP, MGLD, NREX, MOMA, LOAN, SLRY, ODFT
type CashAccountType4Code string

type ClearingSystemIdentification2Choice struct {
	Cd    ExternalClearingSystemIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 Cd"`
	Prtry Max35Text                                 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 Prtry"`
}

type ClearingSystemMemberIdentification2 struct {
	ClrSysId ClearingSystemIdentification2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 ClrSysId,omitempty"`
	MmbId    Max35Text                           `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 MmbId"`
}

type CodeOrProprietary1Choice struct {
	Cd    Max4Text                `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 Cd"`
	Prtry GenericIdentification13 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 Prtry"`
}

type ContactDetails2 struct {
	NmPrfx   NamePrefix1Code `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 NmPrfx,omitempty"`
	Nm       Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 Nm,omitempty"`
	PhneNb   PhoneNumber     `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 PhneNb,omitempty"`
	MobNb    PhoneNumber     `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 MobNb,omitempty"`
	FaxNb    PhoneNumber     `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 FaxNb,omitempty"`
	EmailAdr Max2048Text     `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 EmailAdr,omitempty"`
	Othr     Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 Othr,omitempty"`
}

type ContractDocument1 struct {
	Ref      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 Ref"`
	SgnOffDt ISODate   `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 SgnOffDt,omitempty"`
	Vrsn     Max6Text  `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 Vrsn,omitempty"`
}

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type CustomerAccount1 struct {
	Id            AccountIdentification4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 Id"`
	Nm            Max70Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 Nm,omitempty"`
	Sts           AccountStatus3Code           `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 Sts,omitempty"`
	Tp            CashAccountType2             `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 Tp,omitempty"`
	Ccy           ActiveCurrencyCode           `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 Ccy"`
	MnthlyPmtVal  float64                      `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 MnthlyPmtVal,omitempty"`
	MnthlyRcvdVal float64                      `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 MnthlyRcvdVal,omitempty"`
	MnthlyTxNb    Max5NumericText              `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 MnthlyTxNb,omitempty"`
	AvrgBal       float64                      `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 AvrgBal,omitempty"`
	AcctPurp      Max140Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 AcctPurp,omitempty"`
	FlrNtfctnAmt  float64                      `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 FlrNtfctnAmt,omitempty"`
	ClngNtfctnAmt float64                      `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 ClngNtfctnAmt,omitempty"`
	StmtCycl      Frequency3Code               `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 StmtCycl,omitempty"`
	ClsgDt        ISODate                      `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 ClsgDt,omitempty"`
	Rstrctn       []Restriction1               `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 Rstrctn,omitempty"`
}

type DateAndPlaceOfBirth struct {
	BirthDt     ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 BirthDt"`
	PrvcOfBirth Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 PrvcOfBirth,omitempty"`
	CityOfBirth Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 CityOfBirth"`
	CtryOfBirth CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 CtryOfBirth"`
}

type Document struct {
	AcctExcldMndtMntncReq AccountExcludedMandateMaintenanceRequestV01 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 AcctExcldMndtMntncReq"`
}

// Must be at least 1 items long
type ExternalAccountIdentification1Code string

// Must be at least 1 items long
type ExternalClearingSystemIdentification1Code string

// Must be at least 1 items long
type ExternalFinancialInstitutionIdentification1Code string

// Must be at least 1 items long
type ExternalOrganisationIdentification1Code string

// Must be at least 1 items long
type ExternalPersonIdentification1Code string

type FinancialIdentificationSchemeName1Choice struct {
	Cd    ExternalFinancialInstitutionIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 Cd"`
	Prtry Max35Text                                       `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 Prtry"`
}

type FinancialInstitutionIdentification7 struct {
	BIC         BICIdentifier                       `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 BIC,omitempty"`
	ClrSysMmbId ClearingSystemMemberIdentification2 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 ClrSysMmbId,omitempty"`
	Nm          Max140Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 Nm,omitempty"`
	PstlAdr     PostalAddress6                      `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 PstlAdr,omitempty"`
	Othr        GenericFinancialIdentification1     `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 Othr,omitempty"`
}

// May be one of YEAR, MNTH, QURT, MIAN, WEEK, DAIL, TEND
type Frequency3Code string

type GenericAccountIdentification1 struct {
	Id      Max34Text                `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 Id"`
	SchmeNm AccountSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 SchmeNm,omitempty"`
	Issr    Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 Issr,omitempty"`
}

type GenericFinancialIdentification1 struct {
	Id      Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 Id"`
	SchmeNm FinancialIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 SchmeNm,omitempty"`
	Issr    Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 Issr,omitempty"`
}

type GenericIdentification13 struct {
	Id      Max4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 Id"`
	SchmeNm Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 SchmeNm,omitempty"`
	Issr    Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 Issr"`
}

type GenericOrganisationIdentification1 struct {
	Id      Max35Text                                   `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 Id"`
	SchmeNm OrganisationIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 SchmeNm,omitempty"`
	Issr    Max35Text                                   `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 Issr,omitempty"`
}

type GenericPersonIdentification1 struct {
	Id      Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 Id"`
	SchmeNm PersonIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 SchmeNm,omitempty"`
	Issr    Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 Issr,omitempty"`
}

// Must match the pattern [A-Z]{2,2}[0-9]{2,2}[a-zA-Z0-9]{1,30}
type IBAN2007Identifier string

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

// Must be at least 1 items long
type Max140Text string

// Must be at least 1 items long
type Max16Text string

// Must be at least 1 items long
type Max2048Text string

// Must be at least 1 items long
type Max256Text string

// Must be at least 1 items long
type Max34Text string

// Must be at least 1 items long
type Max350Text string

// Must be at least 1 items long
type Max35Text string

// Must match the pattern [a-zA-Z0-9]{1,4}
type Max4AlphaNumericText string

// Must be at least 1 items long
type Max4Text string

// Must match the pattern [0-9]{1,5}
type Max5NumericText string

// Must be at least 1 items long
type Max6Text string

// Must be at least 1 items long
type Max70Text string

type MessageIdentification1 struct {
	Id      Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 Id"`
	CreDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 CreDtTm"`
}

// May be one of DOCT, MIST, MISS, MADM
type NamePrefix1Code string

type Organisation6 struct {
	FullLglNm    Max350Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 FullLglNm"`
	TradgNm      Max350Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 TradgNm,omitempty"`
	CtryOfOpr    CountryCode                 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 CtryOfOpr"`
	RegnDt       ISODate                     `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 RegnDt,omitempty"`
	OprlAdr      PostalAddress6              `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 OprlAdr,omitempty"`
	BizAdr       PostalAddress6              `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 BizAdr,omitempty"`
	LglAdr       PostalAddress6              `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 LglAdr"`
	OrgId        OrganisationIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 OrgId"`
	RprtvOffcr   []PartyIdentification40     `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 RprtvOffcr,omitempty"`
	TrsrMgr      PartyIdentification40       `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 TrsrMgr,omitempty"`
	MainMndtHldr []PartyIdentification40     `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 MainMndtHldr,omitempty"`
	Sndr         []PartyIdentification40     `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 Sndr,omitempty"`
}

type OrganisationIdentification6 struct {
	BIC  AnyBICIdentifier                     `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 BIC,omitempty"`
	Othr []GenericOrganisationIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 Othr,omitempty"`
}

type OrganisationIdentificationSchemeName1Choice struct {
	Cd    ExternalOrganisationIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 Cd"`
	Prtry Max35Text                               `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 Prtry"`
}

type Party8Choice struct {
	OrgId  OrganisationIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 OrgId"`
	PrvtId PersonIdentification5       `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 PrvtId"`
}

type PartyAndSignature1 struct {
	Pty   PartyIdentification41 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 Pty"`
	Sgntr ProprietaryData3      `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 Sgntr"`
}

type PartyIdentification40 struct {
	Nm        Max140Text            `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 Nm,omitempty"`
	PstlAdr   PostalAddress6        `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 PstlAdr,omitempty"`
	Id        PersonIdentification5 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 Id,omitempty"`
	CtryOfRes CountryCode           `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 CtryOfRes,omitempty"`
	CtctDtls  ContactDetails2       `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 CtctDtls,omitempty"`
}

type PartyIdentification41 struct {
	Nm        Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 Nm,omitempty"`
	PstlAdr   PostalAddress6  `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 PstlAdr,omitempty"`
	Id        Party8Choice    `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 Id,omitempty"`
	CtryOfRes CountryCode     `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 CtryOfRes,omitempty"`
	CtctDtls  ContactDetails2 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 CtctDtls,omitempty"`
}

type PersonIdentification5 struct {
	DtAndPlcOfBirth DateAndPlaceOfBirth            `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 DtAndPlcOfBirth,omitempty"`
	Othr            []GenericPersonIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 Othr,omitempty"`
}

type PersonIdentificationSchemeName1Choice struct {
	Cd    ExternalPersonIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 Cd"`
	Prtry Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 Prtry"`
}

// Must match the pattern \+[0-9]{1,3}-[0-9()+\-]{1,30}
type PhoneNumber string

type PostalAddress6 struct {
	AdrTp       AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 AdrTp,omitempty"`
	Dept        Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 Dept,omitempty"`
	SubDept     Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 SubDept,omitempty"`
	StrtNm      Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 StrtNm,omitempty"`
	BldgNb      Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 BldgNb,omitempty"`
	PstCd       Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 PstCd,omitempty"`
	TwnNm       Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 TwnNm,omitempty"`
	CtrySubDvsn Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 CtrySubDvsn,omitempty"`
	Ctry        CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 Ctry,omitempty"`
	AdrLine     []Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 AdrLine,omitempty"`
}

type ProprietaryData3 struct {
	Item string `xml:",any"`
}

type References4 struct {
	MsgId       MessageIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 MsgId"`
	PrcId       MessageIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 PrcId"`
	AttchdDocNm []Max70Text            `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 AttchdDocNm,omitempty"`
}

type Restriction1 struct {
	RstrctnTp CodeOrProprietary1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 RstrctnTp"`
	VldFr     ISODateTime              `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 VldFr"`
	VldUntil  ISODateTime              `xml:"urn:iso:std:iso:20022:tech:xsd:acmt.015.001.01 VldUntil,omitempty"`
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
