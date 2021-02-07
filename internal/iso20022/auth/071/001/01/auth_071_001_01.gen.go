// Code generated by main. DO NOT EDIT.

package auth_071_001_01

import (
	"bytes"
	"encoding/xml"
	"time"
)

type ActiveOrHistoricCurrencyAndAmount struct {
	Value float64                      `xml:",chardata"`
	Ccy   ActiveOrHistoricCurrencyCode `xml:"Ccy,attr"`
}

// Must match the pattern [A-Z]{3,3}
type ActiveOrHistoricCurrencyCode string

// Must match the pattern [A-Z0-9]{4,4}[A-Z]{2,2}[A-Z0-9]{2,2}([A-Z0-9]{3,3}){0,1}
type AnyBICDec2014Identifier string

type CashReuseData1 struct {
	RinvstdCsh      []ReinvestedCashTypeAndAmount1 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.071.001.01 RinvstdCsh"`
	CshRinvstmtRate float64                        `xml:"urn:iso:std:iso:20022:tech:xsd:auth.071.001.01 CshRinvstmtRate"`
}

type CollateralType12 struct {
	Scty []SecurityReuseData1 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.071.001.01 Scty,omitempty"`
	Csh  []CashReuseData1     `xml:"urn:iso:std:iso:20022:tech:xsd:auth.071.001.01 Csh,omitempty"`
}

type CounterpartyData46 struct {
	RptSubmitgNtty    OrganisationIdentification9Choice `xml:"urn:iso:std:iso:20022:tech:xsd:auth.071.001.01 RptSubmitgNtty"`
	RptgCtrPty        OrganisationIdentification9Choice `xml:"urn:iso:std:iso:20022:tech:xsd:auth.071.001.01 RptgCtrPty"`
	NttyRspnsblForRpt OrganisationIdentification9Choice `xml:"urn:iso:std:iso:20022:tech:xsd:auth.071.001.01 NttyRspnsblForRpt,omitempty"`
}

type Document struct {
	SctiesFincgRptgTxReusdCollDataRpt SecuritiesFinancingReportingTransactionReusedCollateralDataReportV01 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.071.001.01 SctiesFincgRptgTxReusdCollDataRpt"`
}

type FundingSource1 struct {
	Tp     FundingSourceType1Code            `xml:"urn:iso:std:iso:20022:tech:xsd:auth.071.001.01 Tp"`
	MktVal ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:auth.071.001.01 MktVal"`
}

// May be one of SECL, FREE, OTHR, BSHS, CSHS, REPO, UBOR
type FundingSourceType1Code string

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

// Must match the pattern [A-Z0-9]{18,18}[0-9]{2,2}
type LEIIdentifier string

// Must be at least 1 items long
type Max140Text string

// Must be at least 1 items long
type Max350Text string

// Must be at least 1 items long
type Max50Text string

type OrganisationIdentification9Choice struct {
	LEI    LEIIdentifier           `xml:"urn:iso:std:iso:20022:tech:xsd:auth.071.001.01 LEI"`
	ClntId Max50Text               `xml:"urn:iso:std:iso:20022:tech:xsd:auth.071.001.01 ClntId"`
	AnyBIC AnyBICDec2014Identifier `xml:"urn:iso:std:iso:20022:tech:xsd:auth.071.001.01 AnyBIC"`
}

type ReinvestedCashTypeAndAmount1 struct {
	Tp            ReinvestmentType1Code             `xml:"urn:iso:std:iso:20022:tech:xsd:auth.071.001.01 Tp"`
	RinvstdCshAmt ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:auth.071.001.01 RinvstdCshAmt"`
}

// May be one of OTHR, OCMP, MMFT, REPM, SDPU
type ReinvestmentType1Code string

// May be one of NOTX
type ReportPeriodActivity1Code string

type ReuseDataReport3Choice struct {
	New          ReuseDataReportNew3        `xml:"urn:iso:std:iso:20022:tech:xsd:auth.071.001.01 New"`
	Err          ReuseDataReportError3      `xml:"urn:iso:std:iso:20022:tech:xsd:auth.071.001.01 Err"`
	Crrctn       ReuseDataReportCorrection7 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.071.001.01 Crrctn"`
	CollReuseUpd ReuseDataReportCorrection7 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.071.001.01 CollReuseUpd"`
}

type ReuseDataReportCorrection7 struct {
	TechRcrdId  Max140Text           `xml:"urn:iso:std:iso:20022:tech:xsd:auth.071.001.01 TechRcrdId,omitempty"`
	RptgDtTm    ISODateTime          `xml:"urn:iso:std:iso:20022:tech:xsd:auth.071.001.01 RptgDtTm"`
	CtrPtyData  CounterpartyData46   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.071.001.01 CtrPtyData"`
	CollCmpnt   []CollateralType12   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.071.001.01 CollCmpnt"`
	EvtDay      ISODate              `xml:"urn:iso:std:iso:20022:tech:xsd:auth.071.001.01 EvtDay"`
	FndgSrc     []FundingSource1     `xml:"urn:iso:std:iso:20022:tech:xsd:auth.071.001.01 FndgSrc,omitempty"`
	SplmtryData []SupplementaryData1 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.071.001.01 SplmtryData,omitempty"`
}

type ReuseDataReportError3 struct {
	TechRcrdId  Max140Text           `xml:"urn:iso:std:iso:20022:tech:xsd:auth.071.001.01 TechRcrdId,omitempty"`
	RptgDtTm    ISODateTime          `xml:"urn:iso:std:iso:20022:tech:xsd:auth.071.001.01 RptgDtTm"`
	CtrPtyData  CounterpartyData46   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.071.001.01 CtrPtyData"`
	SplmtryData []SupplementaryData1 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.071.001.01 SplmtryData,omitempty"`
}

type ReuseDataReportNew3 struct {
	TechRcrdId  Max140Text           `xml:"urn:iso:std:iso:20022:tech:xsd:auth.071.001.01 TechRcrdId,omitempty"`
	RptgDtTm    ISODateTime          `xml:"urn:iso:std:iso:20022:tech:xsd:auth.071.001.01 RptgDtTm"`
	CtrPtyData  CounterpartyData46   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.071.001.01 CtrPtyData"`
	CollCmpnt   []CollateralType12   `xml:"urn:iso:std:iso:20022:tech:xsd:auth.071.001.01 CollCmpnt"`
	EvtDay      ISODate              `xml:"urn:iso:std:iso:20022:tech:xsd:auth.071.001.01 EvtDay"`
	FndgSrc     []FundingSource1     `xml:"urn:iso:std:iso:20022:tech:xsd:auth.071.001.01 FndgSrc,omitempty"`
	SplmtryData []SupplementaryData1 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.071.001.01 SplmtryData,omitempty"`
}

type ReuseValue1Choice struct {
	Actl   ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:auth.071.001.01 Actl"`
	Estmtd ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:auth.071.001.01 Estmtd"`
}

type SecuritiesFinancingReportingTransactionReusedCollateralDataReportV01 struct {
	TradData    TradeData9Choice     `xml:"urn:iso:std:iso:20022:tech:xsd:auth.071.001.01 TradData"`
	SplmtryData []SupplementaryData1 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.071.001.01 SplmtryData,omitempty"`
}

type SecurityReuseData1 struct {
	ISIN     ISINOct2015Identifier `xml:"urn:iso:std:iso:20022:tech:xsd:auth.071.001.01 ISIN"`
	ReuseVal ReuseValue1Choice     `xml:"urn:iso:std:iso:20022:tech:xsd:auth.071.001.01 ReuseVal"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.071.001.01 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.071.001.01 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

type TradeData9Choice struct {
	DataSetActn ReportPeriodActivity1Code `xml:"urn:iso:std:iso:20022:tech:xsd:auth.071.001.01 DataSetActn"`
	Rpt         []ReuseDataReport3Choice  `xml:"urn:iso:std:iso:20022:tech:xsd:auth.071.001.01 Rpt"`
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