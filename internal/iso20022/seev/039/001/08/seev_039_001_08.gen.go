// Code generated by main. DO NOT EDIT.

package seev_039_001_08

import (
	"bytes"
	"encoding/xml"
	"time"
)

type AccountIdentification10 struct {
	IdCd SafekeepingAccountIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.08 IdCd"`
}

type AccountIdentification29Choice struct {
	ForAllAccts AccountIdentification10   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.08 ForAllAccts"`
	AcctsList   []AccountIdentification31 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.08 AcctsList"`
}

type AccountIdentification31 struct {
	SfkpgAcct Max35Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.08 SfkpgAcct"`
	AcctOwnr  PartyIdentification92Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.08 AcctOwnr,omitempty"`
	SfkpgPlc  SafekeepingPlaceFormat8Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.08 SfkpgPlc,omitempty"`
}

// May be one of CONS, FPRE, PPUT, PPRE
type AdditionalBusinessProcess8Code string

type AdditionalBusinessProcessFormat15Choice struct {
	Cd    AdditionalBusinessProcess8Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.08 Cd"`
	Prtry GenericIdentification30        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.08 Prtry"`
}

// May be one of ADDR, PBOX, HOME, BIZZ, MLTO, DLVY
type AddressType2Code string

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type AnyBICIdentifier string

type CorporateAction43 struct {
	DtDtls         CorporateActionDate59                     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.08 DtDtls,omitempty"`
	EvtStag        CorporateActionEventStageFormat14Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.08 EvtStag,omitempty"`
	AddtlBizPrcInd []AdditionalBusinessProcessFormat15Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.08 AddtlBizPrcInd,omitempty"`
	LtryTp         LotteryTypeFormat4Choice                  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.08 LtryTp,omitempty"`
}

type CorporateActionCancellation3 struct {
	CxlRsnCd CorporateActionCancellationReason1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.08 CxlRsnCd"`
	CxlRsn   Max140Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.08 CxlRsn,omitempty"`
	PrcgSts  CorporateActionEventStatus1            `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.08 PrcgSts"`
}

type CorporateActionCancellationAdviceV08 struct {
	CxlAdvcGnlInf  CorporateActionCancellation3         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.08 CxlAdvcGnlInf"`
	CorpActnGnlInf CorporateActionGeneralInformation124 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.08 CorpActnGnlInf"`
	AcctsDtls      AccountIdentification29Choice        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.08 AcctsDtls"`
	CorpActnDtls   CorporateAction43                    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.08 CorpActnDtls,omitempty"`
	AddtlTxt       []Max8000Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.08 AddtlTxt,omitempty"`
	IssrAgt        []PartyIdentification71Choice        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.08 IssrAgt,omitempty"`
	PngAgt         []PartyIdentification71Choice        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.08 PngAgt,omitempty"`
	SubPngAgt      []PartyIdentification71Choice        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.08 SubPngAgt,omitempty"`
	Regar          PartyIdentification71Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.08 Regar,omitempty"`
	RsellngAgt     []PartyIdentification71Choice        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.08 RsellngAgt,omitempty"`
	PhysSctiesAgt  PartyIdentification71Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.08 PhysSctiesAgt,omitempty"`
	DrpAgt         PartyIdentification71Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.08 DrpAgt,omitempty"`
	SlctnAgt       []PartyIdentification71Choice        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.08 SlctnAgt,omitempty"`
	InfAgt         PartyIdentification71Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.08 InfAgt,omitempty"`
	SplmtryData    []SupplementaryData1                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.08 SplmtryData,omitempty"`
}

// May be one of WITH, PROC
type CorporateActionCancellationReason1Code string

type CorporateActionDate59 struct {
	RcrdDt   DateFormat43Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.08 RcrdDt,omitempty"`
	ExDvddDt DateFormat43Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.08 ExDvddDt,omitempty"`
}

// May be one of FULL, PART, RESC
type CorporateActionEventStage4Code string

type CorporateActionEventStageFormat14Choice struct {
	Cd    CorporateActionEventStage4Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.08 Cd"`
	Prtry GenericIdentification30        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.08 Prtry"`
}

type CorporateActionEventStatus1 struct {
	EvtCmpltnsSts EventCompletenessStatus1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.08 EvtCmpltnsSts"`
	EvtConfSts    EventConfirmationStatus1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.08 EvtConfSts"`
}

// May be one of ACTV, ATTI, BRUP, DFLT, BONU, EXRI, CAPD, CAPG, CAPI, DRCA, DVCA, CHAN, COOP, CLSA, CONS, CONV, CREV, DECR, DETI, DSCL, DVOP, DRIP, DRAW, DTCH, EXOF, REDM, MCAL, INCR, PPMT, INTR, PRII, RHDI, LIQU, EXTM, MRGR, NOOF, CERT, ODLT, OTHR, PARI, PCAL, PRED, PINK, PLAC, PDEF, PRIO, BPUT, REDO, REMK, BIDS, SPLR, RHTS, DVSC, SHPR, SMAL, SOFF, DVSE, SPLF, TREC, TEND, DLST, SUSP, EXWA, WTRC, WRTH, ACCU, INFO, TNDP
type CorporateActionEventType26Code string

type CorporateActionEventType76Choice struct {
	Cd    CorporateActionEventType26Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.08 Cd"`
	Prtry GenericIdentification30        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.08 Prtry"`
}

type CorporateActionGeneralInformation124 struct {
	CorpActnEvtId      Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.08 CorpActnEvtId"`
	OffclCorpActnEvtId Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.08 OffclCorpActnEvtId,omitempty"`
	ClssActnNb         Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.08 ClssActnNb,omitempty"`
	EvtTp              CorporateActionEventType76Choice         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.08 EvtTp"`
	MndtryVlntryEvtTp  CorporateActionMandatoryVoluntary3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.08 MndtryVlntryEvtTp"`
	FinInstrmId        SecurityIdentification19                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.08 FinInstrmId"`
}

// May be one of MAND, CHOS, VOLU
type CorporateActionMandatoryVoluntary1Code string

type CorporateActionMandatoryVoluntary3Choice struct {
	Cd    CorporateActionMandatoryVoluntary1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.08 Cd"`
	Prtry GenericIdentification30                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.08 Prtry"`
}

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type DateAndDateTime2Choice struct {
	Dt   ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.08 Dt"`
	DtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.08 DtTm"`
}

type DateCode19Choice struct {
	Cd    DateType8Code           `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.08 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.08 Prtry"`
}

type DateFormat43Choice struct {
	Dt   DateAndDateTime2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.08 Dt"`
	DtCd DateCode19Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.08 DtCd"`
}

// May be one of UKWN, ONGO
type DateType8Code string

type Document struct {
	CorpActnCxlAdvc CorporateActionCancellationAdviceV08 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.08 CorpActnCxlAdvc"`
}

// May be one of COMP, INCO
type EventCompletenessStatus1Code string

// May be one of CONF, UCON
type EventConfirmationStatus1Code string

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

// Must be at least 1 items long
type ExternalFinancialInstrumentIdentificationType1Code string

type GenericIdentification30 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.08 Id"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.08 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.08 SchmeNm,omitempty"`
}

type GenericIdentification36 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.08 Id"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.08 Issr"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.08 SchmeNm,omitempty"`
}

type GenericIdentification78 struct {
	Tp GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.08 Tp"`
	Id Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.08 Id,omitempty"`
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
	Cd    ExternalFinancialInstrumentIdentificationType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.08 Cd"`
	Prtry Max35Text                                          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.08 Prtry"`
}

// May be one of ORIG, SUPP
type LotteryType1Code string

type LotteryTypeFormat4Choice struct {
	Cd    LotteryType1Code        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.08 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.08 Prtry"`
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

// Must be at least 1 items long
type Max8000Text string

type NameAndAddress5 struct {
	Nm  Max350Text     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.08 Nm"`
	Adr PostalAddress1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.08 Adr,omitempty"`
}

type OtherIdentification1 struct {
	Id  Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.08 Id"`
	Sfx Max16Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.08 Sfx,omitempty"`
	Tp  IdentificationSource3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.08 Tp"`
}

type PartyIdentification71Choice struct {
	AnyBIC   AnyBICIdentifier        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.08 AnyBIC"`
	PrtryId  GenericIdentification36 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.08 PrtryId"`
	NmAndAdr NameAndAddress5         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.08 NmAndAdr"`
}

type PartyIdentification92Choice struct {
	AnyBIC  AnyBICIdentifier        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.08 AnyBIC"`
	PrtryId GenericIdentification36 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.08 PrtryId"`
}

type PostalAddress1 struct {
	AdrTp       AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.08 AdrTp,omitempty"`
	AdrLine     []Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.08 AdrLine,omitempty"`
	StrtNm      Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.08 StrtNm,omitempty"`
	BldgNb      Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.08 BldgNb,omitempty"`
	PstCd       Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.08 PstCd,omitempty"`
	TwnNm       Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.08 TwnNm,omitempty"`
	CtrySubDvsn Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.08 CtrySubDvsn,omitempty"`
	Ctry        CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.08 Ctry"`
}

// May be one of GENR
type SafekeepingAccountIdentification1Code string

// May be one of CUST, ICSD, NCSD, SHHE
type SafekeepingPlace1Code string

// May be one of SHHE, ALLP
type SafekeepingPlace2Code string

type SafekeepingPlaceFormat8Choice struct {
	Id      SafekeepingPlaceTypeAndText6             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.08 Id"`
	Ctry    CountryCode                              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.08 Ctry"`
	TpAndId SafekeepingPlaceTypeAndAnyBICIdentifier1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.08 TpAndId"`
	Prtry   GenericIdentification78                  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.08 Prtry"`
}

type SafekeepingPlaceTypeAndAnyBICIdentifier1 struct {
	SfkpgPlcTp SafekeepingPlace1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.08 SfkpgPlcTp"`
	Id         AnyBICIdentifier      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.08 Id"`
}

type SafekeepingPlaceTypeAndText6 struct {
	SfkpgPlcTp SafekeepingPlace2Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.08 SfkpgPlcTp"`
	Id         Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.08 Id,omitempty"`
}

type SecurityIdentification19 struct {
	ISIN   ISINOct2015Identifier  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.08 ISIN,omitempty"`
	OthrId []OtherIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.08 OthrId,omitempty"`
	Desc   Max140Text             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.08 Desc,omitempty"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.08 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.039.001.08 Envlp"`
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
