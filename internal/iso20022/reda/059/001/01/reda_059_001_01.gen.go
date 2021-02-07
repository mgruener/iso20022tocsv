// Code generated by main. DO NOT EDIT.

package reda_059_001_01

import (
	"bytes"
	"encoding/xml"
	"time"
)

type AccountIdentification26 struct {
	Prtry SimpleIdentificationInformation4 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.059.001.01 Prtry"`
}

// Must match the pattern [A-Z]{3,3}
type ActiveCurrencyCode string

// May be one of ADDR, PBOX, HOME, BIZZ, MLTO, DLVY
type AddressType2Code string

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type AnyBICIdentifier string

// Must match the pattern [A-Z]{6,6}
type CFIOct2015Identifier string

type ClassificationType1Choice struct {
	ClssfctnFinInstrm CFIOct2015Identifier   `xml:"urn:iso:std:iso:20022:tech:xsd:reda.059.001.01 ClssfctnFinInstrm"`
	AltrnClssfctn     GenericIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.059.001.01 AltrnClssfctn"`
}

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type Document struct {
	StgSttlmInstrCxl StandingSettlementInstructionCancellationV01 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.059.001.01 StgSttlmInstrCxl"`
}

type EffectiveDate1 struct {
	FctvDt      ISODate                             `xml:"urn:iso:std:iso:20022:tech:xsd:reda.059.001.01 FctvDt"`
	FctvDtParam ExternalEffectiveDateParameter1Code `xml:"urn:iso:std:iso:20022:tech:xsd:reda.059.001.01 FctvDtParam,omitempty"`
}

// Must be at least 1 items long
type ExternalEffectiveDateParameter1Code string

// Must be at least 1 items long
type ExternalMarketArea1Code string

// Must be at least 1 items long
type ExternalSecuritiesPurpose1Code string

type GenericIdentification1 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:reda.059.001.01 Id"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:reda.059.001.01 SchmeNm,omitempty"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:reda.059.001.01 Issr,omitempty"`
}

type ISODate time.Time

func (t *ISODate) UnmarshalText(text []byte) error {
	return (*xsdDate)(t).UnmarshalText(text)
}
func (t ISODate) MarshalText() ([]byte, error) {
	return xsdDate(t).MarshalText()
}

type MarketIdentification87 struct {
	Ctry       CountryCode               `xml:"urn:iso:std:iso:20022:tech:xsd:reda.059.001.01 Ctry"`
	ClssfctnTp ClassificationType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:reda.059.001.01 ClssfctnTp"`
	SttlmPurp  Purpose3Choice            `xml:"urn:iso:std:iso:20022:tech:xsd:reda.059.001.01 SttlmPurp,omitempty"`
}

type MarketIdentificationOrCashPurpose1Choice struct {
	SttlmInstrMktId MarketIdentification87    `xml:"urn:iso:std:iso:20022:tech:xsd:reda.059.001.01 SttlmInstrMktId"`
	CshSSIPurp      []ExternalMarketArea1Code `xml:"urn:iso:std:iso:20022:tech:xsd:reda.059.001.01 CshSSIPurp"`
}

// Must be at least 1 items long
type Max16Text string

// Must be at least 1 items long
type Max350Text string

// Must be at least 1 items long
type Max35Text string

// Must be at least 1 items long
type Max70Text string

type NameAndAddress5 struct {
	Nm  Max350Text     `xml:"urn:iso:std:iso:20022:tech:xsd:reda.059.001.01 Nm"`
	Adr PostalAddress1 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.059.001.01 Adr,omitempty"`
}

type PartyIdentification63 struct {
	PtyId  PartyIdentification75Choice `xml:"urn:iso:std:iso:20022:tech:xsd:reda.059.001.01 PtyId"`
	PrcgId Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:reda.059.001.01 PrcgId,omitempty"`
}

type PartyIdentification75Choice struct {
	AnyBIC   AnyBICIdentifier `xml:"urn:iso:std:iso:20022:tech:xsd:reda.059.001.01 AnyBIC"`
	NmAndAdr NameAndAddress5  `xml:"urn:iso:std:iso:20022:tech:xsd:reda.059.001.01 NmAndAdr"`
	Ctry     CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:reda.059.001.01 Ctry"`
}

type PartyOrCurrency1Choice struct {
	Dpstry   PartyIdentification63 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.059.001.01 Dpstry"`
	SttlmCcy ActiveCurrencyCode    `xml:"urn:iso:std:iso:20022:tech:xsd:reda.059.001.01 SttlmCcy"`
}

type PostalAddress1 struct {
	AdrTp       AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:reda.059.001.01 AdrTp,omitempty"`
	AdrLine     []Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:reda.059.001.01 AdrLine,omitempty"`
	StrtNm      Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:reda.059.001.01 StrtNm,omitempty"`
	BldgNb      Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:reda.059.001.01 BldgNb,omitempty"`
	PstCd       Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:reda.059.001.01 PstCd,omitempty"`
	TwnNm       Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:reda.059.001.01 TwnNm,omitempty"`
	CtrySubDvsn Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:reda.059.001.01 CtrySubDvsn,omitempty"`
	Ctry        CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:reda.059.001.01 Ctry"`
}

type Purpose3Choice struct {
	SctiesPurpCd ExternalSecuritiesPurpose1Code `xml:"urn:iso:std:iso:20022:tech:xsd:reda.059.001.01 SctiesPurpCd"`
	Prtry        GenericIdentification1         `xml:"urn:iso:std:iso:20022:tech:xsd:reda.059.001.01 Prtry"`
}

type SimpleIdentificationInformation4 struct {
	Id Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:reda.059.001.01 Id"`
}

type StandingSettlementInstructionCancellationV01 struct {
	MsgRefId    Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:reda.059.001.01 MsgRefId"`
	FctvDtDtls  EffectiveDate1                           `xml:"urn:iso:std:iso:20022:tech:xsd:reda.059.001.01 FctvDtDtls,omitempty"`
	AcctId      []AccountIdentification26                `xml:"urn:iso:std:iso:20022:tech:xsd:reda.059.001.01 AcctId"`
	MktId       MarketIdentificationOrCashPurpose1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:reda.059.001.01 MktId"`
	SttlmDtls   PartyOrCurrency1Choice                   `xml:"urn:iso:std:iso:20022:tech:xsd:reda.059.001.01 SttlmDtls"`
	PrvsMsgRef  Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:reda.059.001.01 PrvsMsgRef"`
	SplmtryData []SupplementaryData1                     `xml:"urn:iso:std:iso:20022:tech:xsd:reda.059.001.01 SplmtryData,omitempty"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.059.001.01 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.059.001.01 Envlp"`
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