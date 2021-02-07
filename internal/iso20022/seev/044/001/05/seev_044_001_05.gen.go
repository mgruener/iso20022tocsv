// Code generated by main. DO NOT EDIT.

package seev_044_001_05

import (
	"bytes"
	"encoding/xml"
	"time"
)

type AccountIdentification10 struct {
	IdCd SafekeepingAccountIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.05 IdCd"`
}

type AccountIdentification13Choice struct {
	ForAllAccts AccountIdentification10   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.05 ForAllAccts"`
	AcctsList   []AccountIdentification15 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.05 AcctsList"`
}

type AccountIdentification15 struct {
	SfkpgAcct Max35Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.05 SfkpgAcct"`
	AcctOwnr  PartyIdentification36Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.05 AcctOwnr,omitempty"`
	SfkpgPlc  SafekeepingPlaceFormat2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.05 SfkpgPlc,omitempty"`
}

// May be one of ADDR, PBOX, HOME, BIZZ, MLTO, DLVY
type AddressType2Code string

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type AnyBICIdentifier string

type CorporateAction13 struct {
	DtDtls  CorporateActionDate30                  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.05 DtDtls,omitempty"`
	EvtStag CorporateActionEventStageFormat6Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.05 EvtStag,omitempty"`
	LtryTp  LotteryTypeFormat1Choice               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.05 LtryTp,omitempty"`
}

type CorporateActionDate30 struct {
	RcrdDt   DateFormat19Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.05 RcrdDt,omitempty"`
	ExDvddDt DateFormat19Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.05 ExDvddDt,omitempty"`
}

// May be one of FULL, PART, RESC
type CorporateActionEventStage4Code string

type CorporateActionEventStageFormat6Choice struct {
	Cd    CorporateActionEventStage4Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.05 Cd"`
	Prtry GenericIdentification20        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.05 Prtry"`
}

// May be one of ACCU, ACTV, ATTI, BRUP, DFLT, BONU, EXRI, CAPD, CAPG, CAPI, DRCA, DVCA, CHAN, CLSA, COOP, CONS, CONV, CREV, DECR, DETI, DSCL, DVOP, DRIP, DRAW, DTCH, EXOF, REDM, MCAL, INCR, PPMT, INTR, PRII, RHDI, LIQU, EXTM, MRGR, NOOF, CERT, ODLT, OTHR, PARI, PCAL, PRED, PINK, PLAC, PDEF, PRIO, BPUT, REDO, REMK, BIDS, SPLR, RHTS, DVSC, SHPR, SMAL, SOFF, DVSE, SPLF, TREC, TEND, DLST, SUSP, EXWA, WTRC, WRTH
type CorporateActionEventType13Code string

type CorporateActionEventType14Choice struct {
	Cd    CorporateActionEventType13Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.05 Cd"`
	Prtry GenericIdentification20        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.05 Prtry"`
}

type CorporateActionGeneralInformation53 struct {
	CorpActnEvtId      Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.05 CorpActnEvtId"`
	OffclCorpActnEvtId Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.05 OffclCorpActnEvtId,omitempty"`
	ClssActnNb         Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.05 ClssActnNb,omitempty"`
	EvtTp              CorporateActionEventType14Choice         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.05 EvtTp"`
	MndtryVlntryEvtTp  CorporateActionMandatoryVoluntary1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.05 MndtryVlntryEvtTp"`
	FinInstrmId        SecurityIdentification14                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.05 FinInstrmId"`
}

type CorporateActionMandatoryVoluntary1Choice struct {
	Cd    CorporateActionMandatoryVoluntary1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.05 Cd"`
	Prtry GenericIdentification20                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.05 Prtry"`
}

// May be one of MAND, CHOS, VOLU
type CorporateActionMandatoryVoluntary1Code string

type CorporateActionMovementPreliminaryAdviceCancellationAdviceV05 struct {
	MvmntPrlimryAdvcId DocumentIdentification15            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.05 MvmntPrlimryAdvcId"`
	CorpActnGnlInf     CorporateActionGeneralInformation53 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.05 CorpActnGnlInf"`
	AcctDtls           AccountIdentification13Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.05 AcctDtls"`
	CorpActnDtls       CorporateAction13                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.05 CorpActnDtls,omitempty"`
	IssrAgt            []PartyIdentification46Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.05 IssrAgt,omitempty"`
	PngAgt             []PartyIdentification46Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.05 PngAgt,omitempty"`
	SubPngAgt          []PartyIdentification46Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.05 SubPngAgt,omitempty"`
	Regar              PartyIdentification46Choice         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.05 Regar,omitempty"`
	RsellngAgt         []PartyIdentification46Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.05 RsellngAgt,omitempty"`
	PhysSctiesAgt      PartyIdentification46Choice         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.05 PhysSctiesAgt,omitempty"`
	DrpAgt             PartyIdentification46Choice         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.05 DrpAgt,omitempty"`
	SlctnAgt           []PartyIdentification46Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.05 SlctnAgt,omitempty"`
	InfAgt             PartyIdentification46Choice         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.05 InfAgt,omitempty"`
	SplmtryData        []SupplementaryData1                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.05 SplmtryData,omitempty"`
}

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type DateAndDateTimeChoice struct {
	Dt   ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.05 Dt"`
	DtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.05 DtTm"`
}

type DateCode11Choice struct {
	Cd    DateType8Code           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.05 Cd"`
	Prtry GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.05 Prtry"`
}

type DateFormat19Choice struct {
	Dt   DateAndDateTimeChoice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.05 Dt"`
	DtCd DateCode11Choice      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.05 DtCd"`
}

// May be one of UKWN, ONGO
type DateType8Code string

type Document struct {
	CorpActnMvmntPrlimryAdvcCxlAdvc CorporateActionMovementPreliminaryAdviceCancellationAdviceV05 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.05 CorpActnMvmntPrlimryAdvcCxlAdvc"`
}

type DocumentIdentification15 struct {
	Id    Max35Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.05 Id"`
	LkgTp ProcessingPosition1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.05 LkgTp,omitempty"`
}

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

// Must be at least 1 items long
type ExternalFinancialInstrumentIdentificationType1Code string

type GenericIdentification19 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.05 Id"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.05 Issr"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.05 SchmeNm,omitempty"`
}

type GenericIdentification20 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.05 Id"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.05 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.05 SchmeNm,omitempty"`
}

type GenericIdentification21 struct {
	Tp GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.05 Tp"`
	Id Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.05 Id,omitempty"`
}

// Must match the pattern [A-Z0-9]{12,12}
type ISINIdentifier string

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
	Cd    ExternalFinancialInstrumentIdentificationType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.05 Cd"`
	Prtry Max35Text                                          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.05 Prtry"`
}

// May be one of ORIG, SUPP
type LotteryType1Code string

type LotteryTypeFormat1Choice struct {
	Cd    LotteryType1Code        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.05 Cd"`
	Prtry GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.05 Prtry"`
}

// Must be at least 1 items long
type Max140Text string

// Must be at least 1 items long
type Max16Text string

// Must be at least 1 items long
type Max350Text string

// Must be at least 1 items long
type Max35Text string

// Must be at least 1 items long
type Max70Text string

type NameAndAddress5 struct {
	Nm  Max350Text     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.05 Nm"`
	Adr PostalAddress1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.05 Adr,omitempty"`
}

type OtherIdentification1 struct {
	Id  Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.05 Id"`
	Sfx Max16Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.05 Sfx,omitempty"`
	Tp  IdentificationSource3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.05 Tp"`
}

type PartyIdentification36Choice struct {
	AnyBIC  AnyBICIdentifier        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.05 AnyBIC"`
	PrtryId GenericIdentification19 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.05 PrtryId"`
}

type PartyIdentification46Choice struct {
	AnyBIC   AnyBICIdentifier        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.05 AnyBIC"`
	PrtryId  GenericIdentification19 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.05 PrtryId"`
	NmAndAdr NameAndAddress5         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.05 NmAndAdr"`
}

type PostalAddress1 struct {
	AdrTp       AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.05 AdrTp,omitempty"`
	AdrLine     []Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.05 AdrLine,omitempty"`
	StrtNm      Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.05 StrtNm,omitempty"`
	BldgNb      Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.05 BldgNb,omitempty"`
	PstCd       Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.05 PstCd,omitempty"`
	TwnNm       Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.05 TwnNm,omitempty"`
	CtrySubDvsn Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.05 CtrySubDvsn,omitempty"`
	Ctry        CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.05 Ctry"`
}

type ProcessingPosition1Choice struct {
	Cd    ProcessingPosition3Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.05 Cd"`
	Prtry GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.05 Prtry"`
}

// May be one of AFTE, WITH, BEFO, INFO
type ProcessingPosition3Code string

// May be one of GENR
type SafekeepingAccountIdentification1Code string

// May be one of CUST, ICSD, NCSD, SHHE
type SafekeepingPlace1Code string

// May be one of SHHE, ALLP
type SafekeepingPlace2Code string

type SafekeepingPlaceFormat2Choice struct {
	Id      SafekeepingPlaceTypeAndText2             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.05 Id"`
	Ctry    CountryCode                              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.05 Ctry"`
	TpAndId SafekeepingPlaceTypeAndAnyBICIdentifier1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.05 TpAndId"`
	Prtry   GenericIdentification21                  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.05 Prtry"`
}

type SafekeepingPlaceTypeAndAnyBICIdentifier1 struct {
	SfkpgPlcTp SafekeepingPlace1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.05 SfkpgPlcTp"`
	Id         AnyBICIdentifier      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.05 Id"`
}

type SafekeepingPlaceTypeAndText2 struct {
	SfkpgPlcTp SafekeepingPlace2Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.05 SfkpgPlcTp"`
	Id         Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.05 Id,omitempty"`
}

type SecurityIdentification14 struct {
	ISIN   ISINIdentifier         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.05 ISIN,omitempty"`
	OthrId []OtherIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.05 OthrId,omitempty"`
	Desc   Max140Text             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.05 Desc,omitempty"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.05 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.044.001.05 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
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
