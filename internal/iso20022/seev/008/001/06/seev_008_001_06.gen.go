// Code generated by main. DO NOT EDIT.

package seev_008_001_06

import (
	"bytes"
	"encoding/xml"
	"time"
)

// May be one of ADDR, PBOX, HOME, BIZZ, MLTO, DLVY
type AddressType2Code string

// Must match the pattern [A-Z0-9]{4,4}[A-Z]{2,2}[A-Z0-9]{2,2}([A-Z0-9]{3,3}){0,1}
type AnyBICDec2014Identifier string

type CommunicationAddress11 struct {
	EmailAdr Max256Text  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 EmailAdr,omitempty"`
	URLAdr   Max2048Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 URLAdr,omitempty"`
}

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type Document struct {
	MtgRsltDssmntn MeetingResultDisseminationV06 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 MtgRsltDssmntn"`
}

type EligiblePosition7 struct {
	AcctId    Max35Text                      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 AcctId"`
	AcctOwnr  PartyIdentification228Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 AcctOwnr,omitempty"`
	HldgBal   []HoldingBalance9              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 HldgBal,omitempty"`
	RghtsHldr []PartyIdentification227Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 RghtsHldr,omitempty"`
}

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

// Must be at least 1 items long
type ExternalFinancialInstrumentIdentificationType1Code string

type FinancialInstrumentQuantity18Choice struct {
	Unit    float64 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 Unit"`
	FaceAmt float64 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 FaceAmt"`
}

type GenericIdentification13 struct {
	Id      Max4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 Id"`
	SchmeNm Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 SchmeNm,omitempty"`
	Issr    Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 Issr"`
}

type GenericIdentification30 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 Id"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 SchmeNm,omitempty"`
}

type GenericIdentification36 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 Id"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 Issr"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 SchmeNm,omitempty"`
}

type GenericIdentification78 struct {
	Tp GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 Tp"`
	Id Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 Id,omitempty"`
}

type HoldingBalance9 struct {
	Bal      FinancialInstrumentQuantity18Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 Bal"`
	BalTp    SecuritiesEntryType2Code            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 BalTp"`
	SfkpgPlc SafekeepingPlaceFormat28Choice      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 SfkpgPlc,omitempty"`
	Dt       ISODate                             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 Dt,omitempty"`
}

// Must match the pattern [A-Z]{2,2}[A-Z0-9]{9,9}[0-9]{1,1}
type ISINOct2015Identifier string

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

type IdentificationSource3Choice struct {
	Cd    ExternalFinancialInstrumentIdentificationType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 Cd"`
	Prtry Max35Text                                          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 Prtry"`
}

type IdentificationType45Choice struct {
	Cd    TypeOfIdentification4Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 Cd"`
	Prtry GenericIdentification30   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 Prtry"`
}

// Must match the pattern [A-Z0-9]{18,18}[0-9]{2,2}
type LEIIdentifier string

// Must be at least 1 items long
type Max140Text string

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

// Must match the pattern [a-zA-Z0-9]{1,4}
type Max4AlphaNumericText string

// Must be at least 1 items long
type Max50Text string

// Must be at least 1 items long
type Max70Text string

type MeetingReference8 struct {
	MtgId      Max35Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 MtgId,omitempty"`
	IssrMtgId  Max35Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 IssrMtgId,omitempty"`
	MtgDtAndTm ISODateTime                      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 MtgDtAndTm"`
	Tp         MeetingType4Code                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 Tp"`
	Clssfctn   MeetingTypeClassification2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 Clssfctn,omitempty"`
	Lctn       []PostalAddress1                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 Lctn,omitempty"`
	Issr       PartyIdentification129Choice     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 Issr,omitempty"`
}

type MeetingResultDisseminationV06 struct {
	MtgRsltsDssmntnTp     NotificationType2Code  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 MtgRsltsDssmntnTp"`
	PrvsMtgRsltsDssmntnId Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 PrvsMtgRsltsDssmntnId,omitempty"`
	MtgRef                MeetingReference8      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 MtgRef"`
	Scty                  []SecurityPosition10   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 Scty"`
	VoteRslt              []Vote12               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 VoteRslt"`
	Prtcptn               Participation5         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 Prtcptn,omitempty"`
	AddtlInf              CommunicationAddress11 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 AddtlInf,omitempty"`
	SplmtryData           []SupplementaryData1   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 SplmtryData,omitempty"`
}

// May be one of XMET, GMET, MIXD, SPCL, BMET, CMET
type MeetingType4Code string

type MeetingTypeClassification2Choice struct {
	Cd    MeetingTypeClassification2Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 Cd"`
	Prtry GenericIdentification13        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 Prtry"`
}

// May be one of AMET, CLAS, ISSU, OMET, VRHI
type MeetingTypeClassification2Code string

type NameAndAddress5 struct {
	Nm  Max350Text     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 Nm"`
	Adr PostalAddress1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 Adr,omitempty"`
}

type NaturalPersonIdentification1 struct {
	Id   Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 Id"`
	IdTp IdentificationType45Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 IdTp,omitempty"`
}

// May be one of NEWM, REPL
type NotificationType2Code string

type OtherIdentification1 struct {
	Id  Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 Id"`
	Sfx Max16Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 Sfx,omitempty"`
	Tp  IdentificationSource3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 Tp"`
}

type Participation5 struct {
	TtlNbOfVtngRghts     float64                             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 TtlNbOfVtngRghts,omitempty"`
	PctgOfVtngRghts      float64                             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 PctgOfVtngRghts,omitempty"`
	TtlNbOfSctiesOutsdng FinancialInstrumentQuantity18Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 TtlNbOfSctiesOutsdng,omitempty"`
	ClctnDt              ISODate                             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 ClctnDt,omitempty"`
}

type PartyIdentification129Choice struct {
	AnyBIC   AnyBICDec2014Identifier `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 AnyBIC"`
	PrtryId  GenericIdentification36 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 PrtryId"`
	NmAndAdr NameAndAddress5         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 NmAndAdr"`
	LEI      LEIIdentifier           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 LEI"`
}

type PartyIdentification198Choice struct {
	NtlRegnNb Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 NtlRegnNb"`
	LEI       LEIIdentifier           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 LEI"`
	AnyBIC    AnyBICDec2014Identifier `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 AnyBIC"`
	ClntId    Max50Text               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 ClntId"`
	PrtryId   GenericIdentification36 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 PrtryId"`
}

type PartyIdentification221 struct {
	NmAndAdr PersonName2                  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 NmAndAdr"`
	EmailAdr Max256Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 EmailAdr,omitempty"`
	Id       PartyIdentification198Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 Id"`
}

type PartyIdentification222 struct {
	NmAndAdr PersonName1                  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 NmAndAdr"`
	EmailAdr Max256Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 EmailAdr,omitempty"`
	Id       NaturalPersonIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 Id"`
}

type PartyIdentification224 struct {
	NmAndAdr PersonName2                  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 NmAndAdr"`
	EmailAdr Max256Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 EmailAdr,omitempty"`
	Id       PartyIdentification198Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 Id,omitempty"`
}

type PartyIdentification225 struct {
	NmAndAdr PersonName1                  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 NmAndAdr"`
	EmailAdr Max256Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 EmailAdr,omitempty"`
	Id       NaturalPersonIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 Id,omitempty"`
}

type PartyIdentification227Choice struct {
	LglPrsn  PartyIdentification224   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 LglPrsn"`
	NtrlPrsn []PartyIdentification225 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 NtrlPrsn"`
}

type PartyIdentification228Choice struct {
	LglPrsn  PartyIdentification221   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 LglPrsn"`
	NtrlPrsn []PartyIdentification222 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 NtrlPrsn"`
}

type PersonName1 struct {
	FrstNm Max350Text      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 FrstNm"`
	Srnm   Max350Text      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 Srnm"`
	Adr    PostalAddress26 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 Adr,omitempty"`
}

type PersonName2 struct {
	Nm  Max350Text      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 Nm"`
	Adr PostalAddress26 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 Adr,omitempty"`
}

type PostalAddress1 struct {
	AdrTp       AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 AdrTp,omitempty"`
	AdrLine     []Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 AdrLine,omitempty"`
	StrtNm      Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 StrtNm,omitempty"`
	BldgNb      Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 BldgNb,omitempty"`
	PstCd       Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 PstCd,omitempty"`
	TwnNm       Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 TwnNm,omitempty"`
	CtrySubDvsn Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 CtrySubDvsn,omitempty"`
	Ctry        CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 Ctry"`
}

type PostalAddress26 struct {
	AdrTp       AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 AdrTp,omitempty"`
	AdrLine     []Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 AdrLine,omitempty"`
	StrtNm      Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 StrtNm,omitempty"`
	BldgNb      Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 BldgNb,omitempty"`
	PstBx       Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 PstBx,omitempty"`
	PstCd       Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 PstCd,omitempty"`
	TwnNm       Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 TwnNm,omitempty"`
	CtrySubDvsn Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 CtrySubDvsn,omitempty"`
	Ctry        CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 Ctry"`
}

type ProprietaryVote2 struct {
	Cd  GenericIdentification30             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 Cd"`
	Qty FinancialInstrumentQuantity18Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 Qty"`
}

// May be one of ACPT, REJT, WDRA
type ResolutionStatus2Code string

// May be one of CUST, ICSD, NCSD, SHHE
type SafekeepingPlace1Code string

// May be one of SHHE, ALLP
type SafekeepingPlace2Code string

type SafekeepingPlaceFormat28Choice struct {
	Id      SafekeepingPlaceTypeAndText6           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 Id"`
	Ctry    CountryCode                            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 Ctry"`
	TpAndId SafekeepingPlaceTypeAndIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 TpAndId"`
	Prtry   GenericIdentification78                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 Prtry"`
}

type SafekeepingPlaceTypeAndIdentification1 struct {
	SfkpgPlcTp SafekeepingPlace1Code   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 SfkpgPlcTp"`
	Id         AnyBICDec2014Identifier `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 Id"`
}

type SafekeepingPlaceTypeAndText6 struct {
	SfkpgPlcTp SafekeepingPlace2Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 SfkpgPlcTp"`
	Id         Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 Id,omitempty"`
}

// May be one of BLOK, ELIG, PEND, PENR, NOMI, SETD, BORR, LOAN, SPOS, TRAD, COLI, COLO, UNBA, INBA, REGO
type SecuritiesEntryType2Code string

type SecurityIdentification19 struct {
	ISIN   ISINOct2015Identifier  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 ISIN,omitempty"`
	OthrId []OtherIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 OthrId,omitempty"`
	Desc   Max140Text             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 Desc,omitempty"`
}

type SecurityPosition10 struct {
	FinInstrmId SecurityIdentification19 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 FinInstrmId"`
	Pos         []EligiblePosition7      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 Pos,omitempty"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

// May be one of ARNU, CUST, CORP, DRLC, IDCD, NRIN, CCPT, SOCS, TXID
type TypeOfIdentification4Code string

type Vote12 struct {
	IssrLabl  Max35Text                           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 IssrLabl"`
	RsltnSts  ResolutionStatus2Code               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 RsltnSts"`
	For       FinancialInstrumentQuantity18Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 For,omitempty"`
	Agnst     FinancialInstrumentQuantity18Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 Agnst,omitempty"`
	Abstn     FinancialInstrumentQuantity18Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 Abstn,omitempty"`
	Wthhld    FinancialInstrumentQuantity18Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 Wthhld,omitempty"`
	WthMgmt   FinancialInstrumentQuantity18Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 WthMgmt,omitempty"`
	AgnstMgmt FinancialInstrumentQuantity18Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 AgnstMgmt,omitempty"`
	Dscrtnry  FinancialInstrumentQuantity18Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 Dscrtnry,omitempty"`
	OneYr     FinancialInstrumentQuantity18Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 OneYr,omitempty"`
	TwoYrs    FinancialInstrumentQuantity18Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 TwoYrs,omitempty"`
	ThreeYrs  FinancialInstrumentQuantity18Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 ThreeYrs,omitempty"`
	NoActn    FinancialInstrumentQuantity18Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 NoActn,omitempty"`
	Blnk      FinancialInstrumentQuantity18Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 Blnk,omitempty"`
	Prtry     ProprietaryVote2                    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.008.001.06 Prtry,omitempty"`
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