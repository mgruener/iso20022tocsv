// Code generated by main. DO NOT EDIT.

package auth_009_001_02

import (
	"bytes"
	"encoding/xml"
	"time"
)

type ActiveCurrencyAndAmount struct {
	Value float64            `xml:",chardata"`
	Ccy   ActiveCurrencyCode `xml:"Ccy,attr"`
}

// Must match the pattern [A-Z]{3,3}
type ActiveCurrencyCode string

type ActiveOrHistoricCurrencyAnd13DecimalAmount struct {
	Value float64                      `xml:",chardata"`
	Ccy   ActiveOrHistoricCurrencyCode `xml:"Ccy,attr"`
}

// Must match the pattern [A-Z]{3,3}
type ActiveOrHistoricCurrencyCode string

// May be one of ADDR, PBOX, HOME, BIZZ, MLTO, DLVY
type AddressType2Code string

type AlternateSecurityIdentification1 struct {
	Id         Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.009.001.02 Id"`
	DmstIdSrc  CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:auth.009.001.02 DmstIdSrc"`
	PrtryIdSrc Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.009.001.02 PrtryIdSrc"`
}

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type AnyBICIdentifier string

// Must match the pattern [A-Z]{1,6}
type CFIIdentifier string

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type CurrencyAndAmount struct {
	Value float64      `xml:",chardata"`
	Ccy   CurrencyCode `xml:"Ccy,attr"`
}

// Must match the pattern [A-Z]{3,3}
type CurrencyCode string

type Document struct {
	RgltryTxRptCxlReq RegulatoryTransactionReportCancellationRequestV02 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.009.001.02 RgltryTxRptCxlReq"`
}

type DocumentIdentification8 struct {
	Id      Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.009.001.02 Id"`
	CreDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:auth.009.001.02 CreDtTm,omitempty"`
}

type DocumentIdentification9 struct {
	Id Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:auth.009.001.02 Id"`
}

type FinancialInstrument15 struct {
	Id              SecurityIdentification6Choice  `xml:"urn:iso:std:iso:20022:tech:xsd:auth.009.001.02 Id"`
	InstrmDesc      SecurityInstrumentDescription2 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.009.001.02 InstrmDesc,omitempty"`
	UndrlygInstrmId SecurityIdentification6Choice  `xml:"urn:iso:std:iso:20022:tech:xsd:auth.009.001.02 UndrlygInstrmId,omitempty"`
}

type GenericIdentification1 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:auth.009.001.02 Id"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:auth.009.001.02 SchmeNm,omitempty"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:auth.009.001.02 Issr,omitempty"`
}

type GenericIdentification3 struct {
	Id   Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:auth.009.001.02 Id"`
	Issr Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:auth.009.001.02 Issr,omitempty"`
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

// Must match the pattern [A-Z0-9]{4,4}
type MICIdentifier string

// Must be at least 1 items long
type Max16Text string

// Must be at least 1 items long
type Max350Text string

// Must be at least 1 items long
type Max35Text string

// Must be at least 1 items long
type Max70Text string

type NameAndAddress5 struct {
	Nm  Max350Text     `xml:"urn:iso:std:iso:20022:tech:xsd:auth.009.001.02 Nm"`
	Adr PostalAddress1 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.009.001.02 Adr,omitempty"`
}

type OffMarket1Choice struct {
	OffMktInd    OffMarket1Code   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.009.001.02 OffMktInd"`
	SystmtcIntlr AnyBICIdentifier `xml:"urn:iso:std:iso:20022:tech:xsd:auth.009.001.02 SystmtcIntlr"`
}

// May be one of XOFF, XXXX
type OffMarket1Code string

// May be one of CALL, PUTO
type OptionTypeCode string

// May be one of BUYI, SELL
type OrderDriverCode string

type PartyIdentification11Choice struct {
	BICOrBEI    AnyBICIdentifier       `xml:"urn:iso:std:iso:20022:tech:xsd:auth.009.001.02 BICOrBEI"`
	CntrlCtrPty MICIdentifier          `xml:"urn:iso:std:iso:20022:tech:xsd:auth.009.001.02 CntrlCtrPty"`
	PrtryId     GenericIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.009.001.02 PrtryId"`
}

type PartyIdentification23 struct {
	BICOrBEI AnyBICIdentifier       `xml:"urn:iso:std:iso:20022:tech:xsd:auth.009.001.02 BICOrBEI"`
	PrtryId  GenericIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.009.001.02 PrtryId"`
	NmAndAdr NameAndAddress5        `xml:"urn:iso:std:iso:20022:tech:xsd:auth.009.001.02 NmAndAdr,omitempty"`
}

type PartyIdentification23Choice struct {
	BICOrBEI AnyBICIdentifier       `xml:"urn:iso:std:iso:20022:tech:xsd:auth.009.001.02 BICOrBEI"`
	PrtryId  GenericIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.009.001.02 PrtryId"`
}

type PartyIdentification24Choice struct {
	BICOrBEI AnyBICIdentifier `xml:"urn:iso:std:iso:20022:tech:xsd:auth.009.001.02 BICOrBEI"`
	MIC      MICIdentifier    `xml:"urn:iso:std:iso:20022:tech:xsd:auth.009.001.02 MIC"`
}

type PartyIdentification2Choice struct {
	BICOrBEI AnyBICIdentifier       `xml:"urn:iso:std:iso:20022:tech:xsd:auth.009.001.02 BICOrBEI"`
	PrtryId  GenericIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.009.001.02 PrtryId"`
	NmAndAdr NameAndAddress5        `xml:"urn:iso:std:iso:20022:tech:xsd:auth.009.001.02 NmAndAdr"`
}

type PlaceOfTradeIdentification2Choice struct {
	MktId  MICIdentifier    `xml:"urn:iso:std:iso:20022:tech:xsd:auth.009.001.02 MktId"`
	OffMkt OffMarket1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:auth.009.001.02 OffMkt"`
}

type PostalAddress1 struct {
	AdrTp       AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:auth.009.001.02 AdrTp,omitempty"`
	AdrLine     []Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:auth.009.001.02 AdrLine,omitempty"`
	StrtNm      Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:auth.009.001.02 StrtNm,omitempty"`
	BldgNb      Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:auth.009.001.02 BldgNb,omitempty"`
	PstCd       Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:auth.009.001.02 PstCd,omitempty"`
	TwnNm       Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:auth.009.001.02 TwnNm,omitempty"`
	CtrySubDvsn Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:auth.009.001.02 CtrySubDvsn,omitempty"`
	Ctry        CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:auth.009.001.02 Ctry"`
}

type PriceRateOrAmountChoice struct {
	Rate float64                                    `xml:"urn:iso:std:iso:20022:tech:xsd:auth.009.001.02 Rate"`
	Amt  ActiveOrHistoricCurrencyAnd13DecimalAmount `xml:"urn:iso:std:iso:20022:tech:xsd:auth.009.001.02 Amt"`
}

type RegulatoryTransactionReportCancellationRequestV02 struct {
	Id           DocumentIdentification8     `xml:"urn:iso:std:iso:20022:tech:xsd:auth.009.001.02 Id"`
	RptgInstn    PartyIdentification23Choice `xml:"urn:iso:std:iso:20022:tech:xsd:auth.009.001.02 RptgInstn"`
	RptgAgt      PartyIdentification24Choice `xml:"urn:iso:std:iso:20022:tech:xsd:auth.009.001.02 RptgAgt,omitempty"`
	CxlByTxDtls  []TransactionDetails3       `xml:"urn:iso:std:iso:20022:tech:xsd:auth.009.001.02 CxlByTxDtls"`
	PrvsRef      DocumentIdentification9     `xml:"urn:iso:std:iso:20022:tech:xsd:auth.009.001.02 PrvsRef"`
	CxlByTradRef []TransactionDetails2       `xml:"urn:iso:std:iso:20022:tech:xsd:auth.009.001.02 CxlByTradRef"`
}

type SecurityClassificationType1Choice struct {
	CFI           CFIIdentifier          `xml:"urn:iso:std:iso:20022:tech:xsd:auth.009.001.02 CFI"`
	AltrnClssfctn GenericIdentification3 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.009.001.02 AltrnClssfctn"`
}

type SecurityIdentification6Choice struct {
	ISIN       ISINIdentifier                   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.009.001.02 ISIN"`
	OthrId     AlternateSecurityIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.009.001.02 OthrId"`
	InstrmDesc SecurityInstrumentDescription2   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.009.001.02 InstrmDesc"`
}

type SecurityInstrumentDescription2 struct {
	Desc       Max350Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:auth.009.001.02 Desc,omitempty"`
	ClssfctnTp SecurityClassificationType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:auth.009.001.02 ClssfctnTp,omitempty"`
	PlcOfListg MICIdentifier                     `xml:"urn:iso:std:iso:20022:tech:xsd:auth.009.001.02 PlcOfListg,omitempty"`
	ExrcDt     ISODate                           `xml:"urn:iso:std:iso:20022:tech:xsd:auth.009.001.02 ExrcDt,omitempty"`
	MtrtyDt    ISODate                           `xml:"urn:iso:std:iso:20022:tech:xsd:auth.009.001.02 MtrtyDt,omitempty"`
	OptnTp     OptionTypeCode                    `xml:"urn:iso:std:iso:20022:tech:xsd:auth.009.001.02 OptnTp,omitempty"`
	StrkPric   PriceRateOrAmountChoice           `xml:"urn:iso:std:iso:20022:tech:xsd:auth.009.001.02 StrkPric,omitempty"`
	Mltplr     float64                           `xml:"urn:iso:std:iso:20022:tech:xsd:auth.009.001.02 Mltplr,omitempty"`
}

// May be one of PRIN, CPRN, RISP, PROP, AGEN, CAGN, OAGN, PRAG
type TradingCapacity3Code string

type TransactionDetails2 struct {
	TradRef Max70Text `xml:"urn:iso:std:iso:20022:tech:xsd:auth.009.001.02 TradRef"`
}

type TransactionDetails3 struct {
	TradRef        Max70Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:auth.009.001.02 TradRef"`
	AssoctdTradRef []Max70Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:auth.009.001.02 AssoctdTradRef,omitempty"`
	PlcOfTrad      PlaceOfTradeIdentification2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:auth.009.001.02 PlcOfTrad"`
	TradDtTm       ISODateTime                       `xml:"urn:iso:std:iso:20022:tech:xsd:auth.009.001.02 TradDtTm"`
	FinInstrmDtls  FinancialInstrument15             `xml:"urn:iso:std:iso:20022:tech:xsd:auth.009.001.02 FinInstrmDtls"`
	Sd             OrderDriverCode                   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.009.001.02 Sd"`
	TxRptMrkr      []PartyIdentification24Choice     `xml:"urn:iso:std:iso:20022:tech:xsd:auth.009.001.02 TxRptMrkr,omitempty"`
	CtrPty         PartyIdentification11Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:auth.009.001.02 CtrPty"`
	Clnt           PartyIdentification23             `xml:"urn:iso:std:iso:20022:tech:xsd:auth.009.001.02 Clnt,omitempty"`
	Cpcty          TradingCapacity3Code              `xml:"urn:iso:std:iso:20022:tech:xsd:auth.009.001.02 Cpcty"`
	ExctdTradPric  PriceRateOrAmountChoice           `xml:"urn:iso:std:iso:20022:tech:xsd:auth.009.001.02 ExctdTradPric"`
	ExctdTradQty   UnitOrFaceAmountChoice            `xml:"urn:iso:std:iso:20022:tech:xsd:auth.009.001.02 ExctdTradQty"`
	SttlmAmt       ActiveCurrencyAndAmount           `xml:"urn:iso:std:iso:20022:tech:xsd:auth.009.001.02 SttlmAmt,omitempty"`
	SttlmDt        ISODateTime                       `xml:"urn:iso:std:iso:20022:tech:xsd:auth.009.001.02 SttlmDt,omitempty"`
	PrxyHldr       PartyIdentification2Choice        `xml:"urn:iso:std:iso:20022:tech:xsd:auth.009.001.02 PrxyHldr,omitempty"`
	AddtlInf       Max350Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:auth.009.001.02 AddtlInf,omitempty"`
}

type UnitOrFaceAmountChoice struct {
	Unit    float64           `xml:"urn:iso:std:iso:20022:tech:xsd:auth.009.001.02 Unit"`
	FaceAmt CurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:auth.009.001.02 FaceAmt"`
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
