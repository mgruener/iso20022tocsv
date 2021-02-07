// Code generated by main. DO NOT EDIT.

package sese_018_001_02

import (
	"bytes"
	"encoding/xml"
	"time"
)

type Account5 struct {
	Id    Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 Id,omitempty"`
	Dsgnt Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 Dsgnt,omitempty"`
	Svcr  PartyIdentification2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 Svcr"`
}

type Account6 struct {
	Id    Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 Id"`
	Dsgnt Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 Dsgnt,omitempty"`
	Svcr  PartyIdentification2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 Svcr,omitempty"`
}

type AccountHoldingInformationV02 struct {
	MsgRef           MessageIdentification1         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 MsgRef"`
	PoolRef          AdditionalReference3           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 PoolRef,omitempty"`
	PrvsRef          AdditionalReference3           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 PrvsRef,omitempty"`
	RltdRef          AdditionalReference3           `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 RltdRef,omitempty"`
	BizFlowDrctnTp   BusinessFlowDirectionType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 BizFlowDrctnTp,omitempty"`
	PmryIndvInvstr   IndividualPerson8              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 PmryIndvInvstr,omitempty"`
	ScndryIndvInvstr IndividualPerson8              `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 ScndryIndvInvstr,omitempty"`
	OthrIndvInvstr   []IndividualPerson8            `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 OthrIndvInvstr,omitempty"`
	PmryCorpInvstr   Organisation4                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 PmryCorpInvstr,omitempty"`
	ScndryCorpInvstr Organisation4                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 ScndryCorpInvstr,omitempty"`
	OthrCorpInvstr   []Organisation4                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 OthrCorpInvstr,omitempty"`
	TrfrAcct         Account5                       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 TrfrAcct"`
	NmneeAcct        Account6                       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 NmneeAcct,omitempty"`
	Trfee            PartyIdentification2Choice     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 Trfee"`
	PdctTrf          []ISATransfer4                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 PdctTrf"`
	Xtnsn            []Extension1                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 Xtnsn,omitempty"`
}

type ActiveCurrencyAndAmount struct {
	Value float64            `xml:",chardata"`
	Ccy   ActiveCurrencyCode `xml:"Ccy,attr"`
}

// Must match the pattern [A-Z]{3,3}
type ActiveCurrencyCode string

type ActiveOrHistoricCurrencyAndAmount struct {
	Value float64                      `xml:",chardata"`
	Ccy   ActiveOrHistoricCurrencyCode `xml:"Ccy,attr"`
}

// Must match the pattern [A-Z]{3,3}
type ActiveOrHistoricCurrencyCode string

type AdditionalReference3 struct {
	Ref     Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 Ref"`
	RefIssr PartyIdentification2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 RefIssr,omitempty"`
	MsgNm   Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 MsgNm,omitempty"`
}

// May be one of ADDR, PBOX, HOME, BIZZ, MLTO, DLVY
type AddressType2Code string

type AlternateSecurityIdentification1 struct {
	Id         Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 Id"`
	DmstIdSrc  CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 DmstIdSrc"`
	PrtryIdSrc Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 PrtryIdSrc"`
}

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type AnyBICIdentifier string

// Must be at least 1 items long
type BloombergIdentifier string

// May be one of ADLV, ARCV
type BusinessFlowDirectionType1Code string

// Must be at least 1 items long
type ConsolidatedTapeAssociationIdentifier string

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type CurrentYearType1Choice struct {
	CurYrTp       ISAType1Code    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 CurYrTp"`
	XtndedCurYrTp Extended350Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 XtndedCurYrTp"`
}

type Document struct {
	AcctHldgInf AccountHoldingInformationV02 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 AcctHldgInf"`
}

// Must be at least 1 items long
type EuroclearClearstreamIdentifier string

// Must be at least 1 items long
type Extended350Code string

type Extension1 struct {
	PlcAndNm Max350Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 PlcAndNm"`
	Txt      Max350Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 Txt"`
}

type FinancialInstrument26 struct {
	Id             SecurityIdentification3Choice     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 Id"`
	Nm             Max350Text                        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 Nm,omitempty"`
	Qty            Quantity12Choice                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 Qty,omitempty"`
	AvrgAcqstnPric ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 AvrgAcqstnPric,omitempty"`
	TtlBookVal     ActiveOrHistoricCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 TtlBookVal,omitempty"`
	TrfeeAcct      Account6                          `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 TrfeeAcct,omitempty"`
}

// May be one of MALE, FEMA
type GenderCode string

type GenericIdentification1 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 Id"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 SchmeNm,omitempty"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 Issr,omitempty"`
}

type ISAPortfolio1Choice struct {
	ISA   ISAYearsOfIssue4 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 ISA"`
	Prtfl Portfolio1       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 Prtfl"`
}

type ISATransfer4 struct {
	MstrRef             Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 MstrRef,omitempty"`
	TrfId               Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 TrfId"`
	RsdlCsh             ResidualCash1Code       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 RsdlCsh,omitempty"`
	Prtfl               ISAPortfolio1Choice     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 Prtfl"`
	AllOthrCsh          bool                    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 AllOthrCsh"`
	FinInstrmAsstForTrf []FinancialInstrument26 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 FinInstrmAsstForTrf,omitempty"`
}

// May be one of MINE, MAXI, MINC
type ISAType1Code string

type ISAYearsOfIssue4 struct {
	CurYr          CurrentYearType1Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 CurYr,omitempty"`
	CshCmpntInd    bool                     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 CshCmpntInd"`
	PrvsYrs        PreviousYear2            `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 PrvsYrs,omitempty"`
	CurYrSbcptDtls SubscriptionInformation1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 CurYrSbcptDtls,omitempty"`
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

type ISOYear time.Time

func (t *ISOYear) UnmarshalText(text []byte) error {
	return (*xsdGYear)(t).UnmarshalText(text)
}
func (t ISOYear) MarshalText() ([]byte, error) {
	return xsdGYear(t).MarshalText()
}

type IndividualPerson8 struct {
	Nm            Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 Nm"`
	GvnNm         Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 GvnNm"`
	NmPrfx        NamePrefix1Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 NmPrfx,omitempty"`
	NmSfx         Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 NmSfx,omitempty"`
	Gndr          GenderCode      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 Gndr,omitempty"`
	BirthDt       ISODate         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 BirthDt,omitempty"`
	SclSctyNb     Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 SclSctyNb,omitempty"`
	IndvInvstrAdr PostalAddress1  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 IndvInvstrAdr"`
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

type MessageIdentification1 struct {
	Id      Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 Id"`
	CreDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 CreDtTm"`
}

type NameAndAddress5 struct {
	Nm  Max350Text     `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 Nm"`
	Adr PostalAddress1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 Adr,omitempty"`
}

// May be one of DOCT, MIST, MISS, MADM
type NamePrefix1Code string

type Organisation4 struct {
	Nm            Max140Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 Nm"`
	Id            PartyIdentification4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 Id,omitempty"`
	Purp          Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 Purp,omitempty"`
	TaxtnCtry     CountryCode                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 TaxtnCtry,omitempty"`
	RegnCtry      CountryCode                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 RegnCtry,omitempty"`
	RegnDt        ISODate                    `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 RegnDt,omitempty"`
	TaxIdNb       Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 TaxIdNb,omitempty"`
	NtlRegnNb     Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 NtlRegnNb,omitempty"`
	CorpInvstrAdr PostalAddress1             `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 CorpInvstrAdr"`
}

type PartyIdentification2Choice struct {
	BICOrBEI AnyBICIdentifier       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 BICOrBEI"`
	PrtryId  GenericIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 PrtryId"`
	NmAndAdr NameAndAddress5        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 NmAndAdr"`
}

type PartyIdentification4Choice struct {
	BICOrBEI AnyBICIdentifier       `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 BICOrBEI"`
	PrtryId  GenericIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 PrtryId"`
}

type Portfolio1 struct {
	PrtflInf []Max350Text `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 PrtflInf,omitempty"`
}

type PostalAddress1 struct {
	AdrTp       AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 AdrTp,omitempty"`
	AdrLine     []Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 AdrLine,omitempty"`
	StrtNm      Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 StrtNm,omitempty"`
	BldgNb      Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 BldgNb,omitempty"`
	PstCd       Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 PstCd,omitempty"`
	TwnNm       Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 TwnNm,omitempty"`
	CtrySubDvsn Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 CtrySubDvsn,omitempty"`
	Ctry        CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 Ctry"`
}

// Must match the pattern ALL
type PreviousAll string

type PreviousYear1Choice struct {
	AllPrvsYrs   PreviousAll `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 AllPrvsYrs"`
	SpcfcPrvsYrs []ISOYear   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 SpcfcPrvsYrs"`
}

type PreviousYear2 struct {
	PrvsYrs     PreviousYear1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 PrvsYrs"`
	CshCmpntInd bool                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 CshCmpntInd"`
}

type Quantity12Choice struct {
	Unit     float64 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 Unit,omitempty"`
	PctgRate float64 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 PctgRate,omitempty"`
}

// Must be at least 1 items long
type RICIdentifier string

// May be one of NRCT, RCTR
type ResidualCash1Code string

type SecurityIdentification3Choice struct {
	ISIN        ISINIdentifier                        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 ISIN"`
	SEDOL       string                                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 SEDOL"`
	CUSIP       string                                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 CUSIP"`
	RIC         RICIdentifier                         `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 RIC"`
	TckrSymb    TickerIdentifier                      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 TckrSymb"`
	Blmbrg      BloombergIdentifier                   `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 Blmbrg"`
	CTA         ConsolidatedTapeAssociationIdentifier `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 CTA"`
	QUICK       string                                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 QUICK"`
	Wrtppr      string                                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 Wrtppr"`
	Dtch        string                                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 Dtch"`
	Vlrn        string                                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 Vlrn"`
	SCVM        string                                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 SCVM"`
	Belgn       string                                `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 Belgn"`
	Cmon        EuroclearClearstreamIdentifier        `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 Cmon"`
	OthrPrtryId AlternateSecurityIdentification1      `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 OthrPrtryId"`
}

type SubscriptionInformation1 struct {
	DtOfFrstSbcpt ISODate                 `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 DtOfFrstSbcpt"`
	EqtyCmpnt     ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 EqtyCmpnt,omitempty"`
	CshCmpnt      ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 CshCmpnt,omitempty"`
	TtlAmtYrToDt  ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:sese.018.001.02 TtlAmtYrToDt"`
}

// Must be at least 1 items long
type TickerIdentifier string

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

type xsdGYear time.Time

func (t *xsdGYear) UnmarshalText(text []byte) error {
	return _unmarshalTime(text, (*time.Time)(t), "2006")
}
func (t xsdGYear) MarshalText() ([]byte, error) {
	return _marshalTime((time.Time)(t), "2006")
}
func (t xsdGYear) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if (time.Time)(t).IsZero() {
		return nil
	}
	m, err := t.MarshalText()
	if err != nil {
		return err
	}
	return e.EncodeElement(m, start)
}
func (t xsdGYear) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if (time.Time)(t).IsZero() {
		return xml.Attr{}, nil
	}
	m, err := t.MarshalText()
	return xml.Attr{Name: name, Value: string(m)}, err
}
