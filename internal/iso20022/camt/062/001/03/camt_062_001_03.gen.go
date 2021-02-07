// Code generated by main. DO NOT EDIT.

package camt_062_001_03

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

// May be one of ADDR, PBOX, HOME, BIZZ, MLTO, DLVY
type AddressType2Code string

type AgreedRate2 struct {
	XchgRate float64      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.062.001.03 XchgRate"`
	UnitCcy  CurrencyCode `xml:"urn:iso:std:iso:20022:tech:xsd:camt.062.001.03 UnitCcy"`
	QtdCcy   CurrencyCode `xml:"urn:iso:std:iso:20022:tech:xsd:camt.062.001.03 QtdCcy"`
}

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type AnyBICIdentifier string

type BalanceStatus2 struct {
	Bal ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:camt.062.001.03 Bal"`
}

type ClearingSystemIdentification2Choice struct {
	Cd    ExternalClearingSystemIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.062.001.03 Cd"`
	Prtry Max35Text                                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.062.001.03 Prtry"`
}

// Must match the pattern [A-Z]{2,2}
type CountryCode string

// Must match the pattern [A-Z]{3,3}
type CurrencyCode string

type CurrencyFactors1 struct {
	Ccy         CurrencyCode `xml:"urn:iso:std:iso:20022:tech:xsd:camt.062.001.03 Ccy"`
	ShrtPosLmt  float64      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.062.001.03 ShrtPosLmt"`
	MinPayInAmt float64      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.062.001.03 MinPayInAmt"`
	VoltlyMrgn  float64      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.062.001.03 VoltlyMrgn"`
	Rate        AgreedRate2  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.062.001.03 Rate,omitempty"`
}

type Document struct {
	PayInSchdl PayInScheduleV03 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.062.001.03 PayInSchdl"`
}

// May be one of TRIA, OFFI, REQU
type Entry2Code string

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

// Must be at least 1 items long
type ExternalClearingSystemIdentification1Code string

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

// Must be at least 1 items long
type Max105Text string

// Must be at least 1 items long
type Max16Text string

// Must be at least 1 items long
type Max34Text string

// Must be at least 1 items long
type Max350Text string

// Must be at least 1 items long
type Max35Text string

// Must be at least 1 items long
type Max70Text string

type NameAndAddress8 struct {
	Nm         Max350Text     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.062.001.03 Nm"`
	Adr        PostalAddress1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.062.001.03 Adr,omitempty"`
	AltrntvIdr []Max35Text    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.062.001.03 AltrntvIdr,omitempty"`
}

type PartyIdentification44 struct {
	AnyBIC     AnyBICIdentifier `xml:"urn:iso:std:iso:20022:tech:xsd:camt.062.001.03 AnyBIC"`
	AltrntvIdr []Max35Text      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.062.001.03 AltrntvIdr,omitempty"`
}

type PartyIdentification59 struct {
	PtyNm      Max34Text                           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.062.001.03 PtyNm,omitempty"`
	AnyBIC     PartyIdentification44               `xml:"urn:iso:std:iso:20022:tech:xsd:camt.062.001.03 AnyBIC,omitempty"`
	AcctNb     Max34Text                           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.062.001.03 AcctNb,omitempty"`
	Adr        Max105Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.062.001.03 Adr,omitempty"`
	ClrSysId   ClearingSystemIdentification2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.062.001.03 ClrSysId,omitempty"`
	LglNttyIdr LEIIdentifier                       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.062.001.03 LglNttyIdr,omitempty"`
}

type PartyIdentification73Choice struct {
	NmAndAdr NameAndAddress8       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.062.001.03 NmAndAdr"`
	AnyBIC   PartyIdentification44 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.062.001.03 AnyBIC"`
	PtyId    PartyIdentification59 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.062.001.03 PtyId"`
}

type PayInFactors1 struct {
	AggtShrtPosLmt ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:camt.062.001.03 AggtShrtPosLmt"`
	CcyFctrs       []CurrencyFactors1      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.062.001.03 CcyFctrs"`
}

type PayInScheduleItems1 struct {
	Amt  ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:camt.062.001.03 Amt"`
	Ddln ISODateTime             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.062.001.03 Ddln"`
}

type PayInScheduleV03 struct {
	PtyId            PartyIdentification73Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.062.001.03 PtyId"`
	RptData          ReportData4                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.062.001.03 RptData"`
	PayInSchdlLngBal []BalanceStatus2            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.062.001.03 PayInSchdlLngBal,omitempty"`
	PayInSchdlItm    []PayInScheduleItems1       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.062.001.03 PayInSchdlItm,omitempty"`
	PayInFctrs       PayInFactors1               `xml:"urn:iso:std:iso:20022:tech:xsd:camt.062.001.03 PayInFctrs,omitempty"`
	SplmtryData      []SupplementaryData1        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.062.001.03 SplmtryData,omitempty"`
}

type PostalAddress1 struct {
	AdrTp       AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.062.001.03 AdrTp,omitempty"`
	AdrLine     []Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.062.001.03 AdrLine,omitempty"`
	StrtNm      Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.062.001.03 StrtNm,omitempty"`
	BldgNb      Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.062.001.03 BldgNb,omitempty"`
	PstCd       Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.062.001.03 PstCd,omitempty"`
	TwnNm       Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.062.001.03 TwnNm,omitempty"`
	CtrySubDvsn Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.062.001.03 CtrySubDvsn,omitempty"`
	Ctry        CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.062.001.03 Ctry"`
}

type ReportData4 struct {
	MsgId       Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.062.001.03 MsgId"`
	ValDt       ISODate                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.062.001.03 ValDt"`
	DtAndTmStmp ISODateTime            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.062.001.03 DtAndTmStmp"`
	Tp          Entry2Code             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.062.001.03 Tp"`
	SchdlTp     Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:camt.062.001.03 SchdlTp"`
	SttlmSsnIdr Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:camt.062.001.03 SttlmSsnIdr,omitempty"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.062.001.03 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.062.001.03 Envlp"`
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
