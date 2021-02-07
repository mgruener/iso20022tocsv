// Code generated by main. DO NOT EDIT.

package auth_015_001_01

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

type CounterpartyIdentification2Choice struct {
	LEI  LEIIdentifier                `xml:"urn:iso:std:iso:20022:tech:xsd:auth.015.001.01 LEI"`
	Othr ReportedPartyIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.015.001.01 Othr"`
}

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type DateAndDateTimeChoice struct {
	Dt   ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:auth.015.001.01 Dt"`
	DtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:auth.015.001.01 DtTm"`
}

type DateTimePeriod1 struct {
	FrDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:auth.015.001.01 FrDtTm"`
	ToDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:auth.015.001.01 ToDtTm"`
}

type Document struct {
	MnyMktOvrnghtIndxSwpsSttstclRpt MoneyMarketOvernightIndexSwapsStatisticalReportV01 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.015.001.01 MnyMktOvrnghtIndxSwpsSttstclRpt"`
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

// Must be at least 1 items long
type Max105Text string

// Must be at least 1 items long
type Max350Text string

// Must be at least 1 items long
type Max70Text string

type MoneyMarketOvernightIndexSwapsStatisticalReportV01 struct {
	RptHdr             MoneyMarketReportHeader1  `xml:"urn:iso:std:iso:20022:tech:xsd:auth.015.001.01 RptHdr"`
	OvrnghtIndxSwpsRpt OvernightIndexSwap3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:auth.015.001.01 OvrnghtIndxSwpsRpt"`
	SplmtryData        []SupplementaryData1      `xml:"urn:iso:std:iso:20022:tech:xsd:auth.015.001.01 SplmtryData,omitempty"`
}

type MoneyMarketReportHeader1 struct {
	RptgAgt LEIIdentifier   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.015.001.01 RptgAgt"`
	RefPrd  DateTimePeriod1 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.015.001.01 RefPrd"`
}

type NameOrSector1Choice struct {
	Nm   Max70Text `xml:"urn:iso:std:iso:20022:tech:xsd:auth.015.001.01 Nm"`
	Sctr string    `xml:"urn:iso:std:iso:20022:tech:xsd:auth.015.001.01 Sctr"`
}

type OvernightIndexSwap3Choice struct {
	DataSetActn ReportPeriodActivity1Code        `xml:"urn:iso:std:iso:20022:tech:xsd:auth.015.001.01 DataSetActn"`
	Tx          []OvernightIndexSwapTransaction3 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.015.001.01 Tx"`
}

type OvernightIndexSwapTransaction3 struct {
	RptdTxSts       TransactionOperationType1Code     `xml:"urn:iso:std:iso:20022:tech:xsd:auth.015.001.01 RptdTxSts"`
	BrnchId         LEIIdentifier                     `xml:"urn:iso:std:iso:20022:tech:xsd:auth.015.001.01 BrnchId,omitempty"`
	UnqTxIdr        Max105Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:auth.015.001.01 UnqTxIdr,omitempty"`
	PrtryTxId       Max105Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:auth.015.001.01 PrtryTxId"`
	CtrPtyPrtryTxId Max105Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:auth.015.001.01 CtrPtyPrtryTxId,omitempty"`
	CtrPtyId        CounterpartyIdentification2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:auth.015.001.01 CtrPtyId"`
	TradDt          DateAndDateTimeChoice             `xml:"urn:iso:std:iso:20022:tech:xsd:auth.015.001.01 TradDt"`
	StartDt         ISODate                           `xml:"urn:iso:std:iso:20022:tech:xsd:auth.015.001.01 StartDt"`
	MtrtyDt         ISODate                           `xml:"urn:iso:std:iso:20022:tech:xsd:auth.015.001.01 MtrtyDt"`
	FxdIntrstRate   Rate2                             `xml:"urn:iso:std:iso:20022:tech:xsd:auth.015.001.01 FxdIntrstRate"`
	TxTp            OvernightIndexSwapType1Code       `xml:"urn:iso:std:iso:20022:tech:xsd:auth.015.001.01 TxTp"`
	TxNmnlAmt       ActiveCurrencyAndAmount           `xml:"urn:iso:std:iso:20022:tech:xsd:auth.015.001.01 TxNmnlAmt"`
	SplmtryData     []SupplementaryData1              `xml:"urn:iso:std:iso:20022:tech:xsd:auth.015.001.01 SplmtryData,omitempty"`
}

// May be one of PAID, RECE
type OvernightIndexSwapType1Code string

type Rate2 struct {
	Sgn  bool    `xml:"urn:iso:std:iso:20022:tech:xsd:auth.015.001.01 Sgn,omitempty"`
	Rate float64 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.015.001.01 Rate"`
}

// May be one of NOTX
type ReportPeriodActivity1Code string

type ReportedPartyIdentification1 struct {
	NmOrSctr NameOrSector1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:auth.015.001.01 NmOrSctr"`
	Lctn     CountryCode         `xml:"urn:iso:std:iso:20022:tech:xsd:auth.015.001.01 Lctn"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.015.001.01 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.015.001.01 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

// May be one of AMND, CANC, CORR, NEWT
type TransactionOperationType1Code string

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
