// Code generated by main. DO NOT EDIT.

package sese_018_001_05

import (
	"bytes"
	"encoding/xml"
	"time"
)

type Account19 struct {
	Id    Max35Text                   `xml:"urn:swift:xsd:sese.018.001.05 Id,omitempty"`
	Dsgnt Max35Text                   `xml:"urn:swift:xsd:sese.018.001.05 Dsgnt,omitempty"`
	Svcr  PartyIdentification70Choice `xml:"urn:swift:xsd:sese.018.001.05 Svcr,omitempty"`
}

type AccountHoldingInformationV05 struct {
	MsgRef           MessageIdentification1         `xml:"urn:swift:xsd:sese.018.001.05 MsgRef"`
	PoolRef          AdditionalReference6           `xml:"urn:swift:xsd:sese.018.001.05 PoolRef,omitempty"`
	PrvsRef          AdditionalReference6           `xml:"urn:swift:xsd:sese.018.001.05 PrvsRef,omitempty"`
	RltdRef          AdditionalReference6           `xml:"urn:swift:xsd:sese.018.001.05 RltdRef,omitempty"`
	BizFlowDrctnTp   BusinessFlowDirectionType1Code `xml:"urn:swift:xsd:sese.018.001.05 BizFlowDrctnTp,omitempty"`
	PmryIndvInvstr   IndividualPerson8              `xml:"urn:swift:xsd:sese.018.001.05 PmryIndvInvstr,omitempty"`
	ScndryIndvInvstr IndividualPerson8              `xml:"urn:swift:xsd:sese.018.001.05 ScndryIndvInvstr,omitempty"`
	OthrIndvInvstr   []IndividualPerson8            `xml:"urn:swift:xsd:sese.018.001.05 OthrIndvInvstr,omitempty"`
	PmryCorpInvstr   Organisation21                 `xml:"urn:swift:xsd:sese.018.001.05 PmryCorpInvstr,omitempty"`
	ScndryCorpInvstr Organisation21                 `xml:"urn:swift:xsd:sese.018.001.05 ScndryCorpInvstr,omitempty"`
	OthrCorpInvstr   []Organisation21               `xml:"urn:swift:xsd:sese.018.001.05 OthrCorpInvstr,omitempty"`
	TrfrAcct         Account19                      `xml:"urn:swift:xsd:sese.018.001.05 TrfrAcct"`
	NmneeAcct        Account19                      `xml:"urn:swift:xsd:sese.018.001.05 NmneeAcct,omitempty"`
	Trfee            PartyIdentification70Choice    `xml:"urn:swift:xsd:sese.018.001.05 Trfee"`
	PdctTrf          []ISATransfer23                `xml:"urn:swift:xsd:sese.018.001.05 PdctTrf"`
	MktPrctcVrsn     MarketPracticeVersion1         `xml:"urn:swift:xsd:sese.018.001.05 MktPrctcVrsn,omitempty"`
	Xtnsn            []Extension1                   `xml:"urn:swift:xsd:sese.018.001.05 Xtnsn,omitempty"`
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

type AdditionalReference6 struct {
	Ref     Max35Text                   `xml:"urn:swift:xsd:sese.018.001.05 Ref"`
	RefIssr PartyIdentification90Choice `xml:"urn:swift:xsd:sese.018.001.05 RefIssr,omitempty"`
	MsgNm   Max35Text                   `xml:"urn:swift:xsd:sese.018.001.05 MsgNm,omitempty"`
}

// May be one of ADDR, PBOX, HOME, BIZZ, MLTO, DLVY
type AddressType2Code string

// May be one of LIQU, NLIQ
type AllOtherCash1Code string

type AlternateSecurityIdentification7 struct {
	Id    Max35Text                   `xml:"urn:swift:xsd:sese.018.001.05 Id"`
	IdSrc IdentificationSource1Choice `xml:"urn:swift:xsd:sese.018.001.05 IdSrc"`
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
	CurYrTp       ISAType1Code    `xml:"urn:swift:xsd:sese.018.001.05 CurYrTp"`
	XtndedCurYrTp Extended350Code `xml:"urn:swift:xsd:sese.018.001.05 XtndedCurYrTp"`
}

type Document struct {
	AcctHldgInf AccountHoldingInformationV05 `xml:"urn:swift:xsd:sese.018.001.05 AcctHldgInf"`
}

// Must be at least 1 items long
type EuroclearClearstreamIdentifier string

// Must be at least 1 items long
type Extended350Code string

type Extension1 struct {
	PlcAndNm Max350Text `xml:"urn:swift:xsd:sese.018.001.05 PlcAndNm"`
	Txt      Max350Text `xml:"urn:swift:xsd:sese.018.001.05 Txt"`
}

type FinancialInstrument47 struct {
	Id             SecurityIdentification23Choice    `xml:"urn:swift:xsd:sese.018.001.05 Id"`
	Nm             Max350Text                        `xml:"urn:swift:xsd:sese.018.001.05 Nm,omitempty"`
	ShrtNm         Max35Text                         `xml:"urn:swift:xsd:sese.018.001.05 ShrtNm,omitempty"`
	Qty            Quantity12Choice                  `xml:"urn:swift:xsd:sese.018.001.05 Qty,omitempty"`
	AvrgAcqstnPric ActiveOrHistoricCurrencyAndAmount `xml:"urn:swift:xsd:sese.018.001.05 AvrgAcqstnPric,omitempty"`
	TtlBookVal     ActiveOrHistoricCurrencyAndAmount `xml:"urn:swift:xsd:sese.018.001.05 TtlBookVal,omitempty"`
	TrfeeAcct      Account19                         `xml:"urn:swift:xsd:sese.018.001.05 TrfeeAcct,omitempty"`
	SubAcctDtls    SubAccount5                       `xml:"urn:swift:xsd:sese.018.001.05 SubAcctDtls,omitempty"`
	DlvrgAgtDtls   PartyIdentificationAndAccount125  `xml:"urn:swift:xsd:sese.018.001.05 DlvrgAgtDtls,omitempty"`
}

// May be one of MALE, FEMA
type GenderCode string

type GenericIdentification1 struct {
	Id      Max35Text `xml:"urn:swift:xsd:sese.018.001.05 Id"`
	SchmeNm Max35Text `xml:"urn:swift:xsd:sese.018.001.05 SchmeNm,omitempty"`
	Issr    Max35Text `xml:"urn:swift:xsd:sese.018.001.05 Issr,omitempty"`
}

type ISAPortfolio1Choice struct {
	ISA   ISAYearsOfIssue4 `xml:"urn:swift:xsd:sese.018.001.05 ISA"`
	Prtfl Portfolio1       `xml:"urn:swift:xsd:sese.018.001.05 Prtfl"`
}

type ISATransfer23 struct {
	MstrRef             Max35Text               `xml:"urn:swift:xsd:sese.018.001.05 MstrRef,omitempty"`
	TrfId               Max35Text               `xml:"urn:swift:xsd:sese.018.001.05 TrfId"`
	TrfConfId           Max35Text               `xml:"urn:swift:xsd:sese.018.001.05 TrfConfId,omitempty"`
	RsdlCsh             ResidualCash1Code       `xml:"urn:swift:xsd:sese.018.001.05 RsdlCsh,omitempty"`
	Prtfl               ISAPortfolio1Choice     `xml:"urn:swift:xsd:sese.018.001.05 Prtfl,omitempty"`
	AllOthrCsh          AllOtherCash1Code       `xml:"urn:swift:xsd:sese.018.001.05 AllOthrCsh,omitempty"`
	FinInstrmAsstForTrf []FinancialInstrument47 `xml:"urn:swift:xsd:sese.018.001.05 FinInstrmAsstForTrf,omitempty"`
}

// May be one of MINE, MAXI, MINC
type ISAType1Code string

type ISAYearsOfIssue4 struct {
	CurYr          CurrentYearType1Choice   `xml:"urn:swift:xsd:sese.018.001.05 CurYr,omitempty"`
	CshCmpntInd    bool                     `xml:"urn:swift:xsd:sese.018.001.05 CshCmpntInd"`
	PrvsYrs        PreviousYear2            `xml:"urn:swift:xsd:sese.018.001.05 PrvsYrs,omitempty"`
	CurYrSbcptDtls SubscriptionInformation1 `xml:"urn:swift:xsd:sese.018.001.05 CurYrSbcptDtls,omitempty"`
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

type ISOYearMonth time.Time

func (t *ISOYearMonth) UnmarshalText(text []byte) error {
	return (*xsdGYearMonth)(t).UnmarshalText(text)
}
func (t ISOYearMonth) MarshalText() ([]byte, error) {
	return xsdGYearMonth(t).MarshalText()
}

type IdentificationSource1Choice struct {
	Dmst  CountryCode `xml:"urn:swift:xsd:sese.018.001.05 Dmst"`
	Prtry Max35Text   `xml:"urn:swift:xsd:sese.018.001.05 Prtry"`
}

type IndividualPerson8 struct {
	Nm            Max35Text       `xml:"urn:swift:xsd:sese.018.001.05 Nm"`
	GvnNm         Max35Text       `xml:"urn:swift:xsd:sese.018.001.05 GvnNm"`
	NmPrfx        NamePrefix1Code `xml:"urn:swift:xsd:sese.018.001.05 NmPrfx,omitempty"`
	NmSfx         Max35Text       `xml:"urn:swift:xsd:sese.018.001.05 NmSfx,omitempty"`
	Gndr          GenderCode      `xml:"urn:swift:xsd:sese.018.001.05 Gndr,omitempty"`
	BirthDt       ISODate         `xml:"urn:swift:xsd:sese.018.001.05 BirthDt,omitempty"`
	SclSctyNb     Max35Text       `xml:"urn:swift:xsd:sese.018.001.05 SclSctyNb,omitempty"`
	IndvInvstrAdr PostalAddress1  `xml:"urn:swift:xsd:sese.018.001.05 IndvInvstrAdr"`
}

type MarketPracticeVersion1 struct {
	Nm Max35Text    `xml:"urn:swift:xsd:sese.018.001.05 Nm"`
	Dt ISOYearMonth `xml:"urn:swift:xsd:sese.018.001.05 Dt,omitempty"`
	Nb Max35Text    `xml:"urn:swift:xsd:sese.018.001.05 Nb,omitempty"`
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
	Id      Max35Text   `xml:"urn:swift:xsd:sese.018.001.05 Id"`
	CreDtTm ISODateTime `xml:"urn:swift:xsd:sese.018.001.05 CreDtTm"`
}

type NameAndAddress5 struct {
	Nm  Max350Text     `xml:"urn:swift:xsd:sese.018.001.05 Nm"`
	Adr PostalAddress1 `xml:"urn:swift:xsd:sese.018.001.05 Adr,omitempty"`
}

// May be one of DOCT, MIST, MISS, MADM
type NamePrefix1Code string

type Organisation21 struct {
	Nm            Max140Text                  `xml:"urn:swift:xsd:sese.018.001.05 Nm"`
	Id            PartyIdentification72Choice `xml:"urn:swift:xsd:sese.018.001.05 Id,omitempty"`
	Purp          Max35Text                   `xml:"urn:swift:xsd:sese.018.001.05 Purp,omitempty"`
	TaxtnCtry     CountryCode                 `xml:"urn:swift:xsd:sese.018.001.05 TaxtnCtry,omitempty"`
	RegnCtry      CountryCode                 `xml:"urn:swift:xsd:sese.018.001.05 RegnCtry,omitempty"`
	RegnDt        ISODate                     `xml:"urn:swift:xsd:sese.018.001.05 RegnDt,omitempty"`
	TaxIdNb       Max35Text                   `xml:"urn:swift:xsd:sese.018.001.05 TaxIdNb,omitempty"`
	NtlRegnNb     Max35Text                   `xml:"urn:swift:xsd:sese.018.001.05 NtlRegnNb,omitempty"`
	CorpInvstrAdr PostalAddress1              `xml:"urn:swift:xsd:sese.018.001.05 CorpInvstrAdr"`
}

type PartyIdentification70Choice struct {
	AnyBIC   AnyBICIdentifier       `xml:"urn:swift:xsd:sese.018.001.05 AnyBIC"`
	PrtryId  GenericIdentification1 `xml:"urn:swift:xsd:sese.018.001.05 PrtryId"`
	NmAndAdr NameAndAddress5        `xml:"urn:swift:xsd:sese.018.001.05 NmAndAdr"`
}

type PartyIdentification72Choice struct {
	AnyBIC  AnyBICIdentifier       `xml:"urn:swift:xsd:sese.018.001.05 AnyBIC"`
	PrtryId GenericIdentification1 `xml:"urn:swift:xsd:sese.018.001.05 PrtryId"`
}

type PartyIdentification90Choice struct {
	AnyBIC   AnyBICIdentifier       `xml:"urn:swift:xsd:sese.018.001.05 AnyBIC"`
	PrtryId  GenericIdentification1 `xml:"urn:swift:xsd:sese.018.001.05 PrtryId"`
	NmAndAdr NameAndAddress5        `xml:"urn:swift:xsd:sese.018.001.05 NmAndAdr"`
}

type PartyIdentificationAndAccount125 struct {
	PtyId      PartyIdentification70Choice `xml:"urn:swift:xsd:sese.018.001.05 PtyId,omitempty"`
	AcctId     Max35Text                   `xml:"urn:swift:xsd:sese.018.001.05 AcctId,omitempty"`
	PlcOfSttlm PartyIdentification70Choice `xml:"urn:swift:xsd:sese.018.001.05 PlcOfSttlm"`
}

type Portfolio1 struct {
	PrtflInf []Max350Text `xml:"urn:swift:xsd:sese.018.001.05 PrtflInf,omitempty"`
}

type PostalAddress1 struct {
	AdrTp       AddressType2Code `xml:"urn:swift:xsd:sese.018.001.05 AdrTp,omitempty"`
	AdrLine     []Max70Text      `xml:"urn:swift:xsd:sese.018.001.05 AdrLine,omitempty"`
	StrtNm      Max70Text        `xml:"urn:swift:xsd:sese.018.001.05 StrtNm,omitempty"`
	BldgNb      Max16Text        `xml:"urn:swift:xsd:sese.018.001.05 BldgNb,omitempty"`
	PstCd       Max16Text        `xml:"urn:swift:xsd:sese.018.001.05 PstCd,omitempty"`
	TwnNm       Max35Text        `xml:"urn:swift:xsd:sese.018.001.05 TwnNm,omitempty"`
	CtrySubDvsn Max35Text        `xml:"urn:swift:xsd:sese.018.001.05 CtrySubDvsn,omitempty"`
	Ctry        CountryCode      `xml:"urn:swift:xsd:sese.018.001.05 Ctry"`
}

// Must match the pattern ALL
type PreviousAll string

type PreviousYear1Choice struct {
	AllPrvsYrs   PreviousAll `xml:"urn:swift:xsd:sese.018.001.05 AllPrvsYrs"`
	SpcfcPrvsYrs []ISOYear   `xml:"urn:swift:xsd:sese.018.001.05 SpcfcPrvsYrs"`
}

type PreviousYear2 struct {
	PrvsYrs     PreviousYear1Choice `xml:"urn:swift:xsd:sese.018.001.05 PrvsYrs"`
	CshCmpntInd bool                `xml:"urn:swift:xsd:sese.018.001.05 CshCmpntInd"`
}

type Quantity12Choice struct {
	Unit     float64 `xml:"urn:swift:xsd:sese.018.001.05 Unit,omitempty"`
	PctgRate float64 `xml:"urn:swift:xsd:sese.018.001.05 PctgRate,omitempty"`
}

// Must be at least 1 items long
type RICIdentifier string

// May be one of NRCT, RCTR
type ResidualCash1Code string

type SecurityIdentification23Choice struct {
	ISIN        ISINIdentifier                        `xml:"urn:swift:xsd:sese.018.001.05 ISIN"`
	SEDOL       string                                `xml:"urn:swift:xsd:sese.018.001.05 SEDOL"`
	CUSIP       string                                `xml:"urn:swift:xsd:sese.018.001.05 CUSIP"`
	RIC         RICIdentifier                         `xml:"urn:swift:xsd:sese.018.001.05 RIC"`
	TckrSymb    TickerIdentifier                      `xml:"urn:swift:xsd:sese.018.001.05 TckrSymb"`
	Blmbrg      BloombergIdentifier                   `xml:"urn:swift:xsd:sese.018.001.05 Blmbrg"`
	CTA         ConsolidatedTapeAssociationIdentifier `xml:"urn:swift:xsd:sese.018.001.05 CTA"`
	QUICK       string                                `xml:"urn:swift:xsd:sese.018.001.05 QUICK"`
	Wrtppr      string                                `xml:"urn:swift:xsd:sese.018.001.05 Wrtppr"`
	Dtch        string                                `xml:"urn:swift:xsd:sese.018.001.05 Dtch"`
	Vlrn        string                                `xml:"urn:swift:xsd:sese.018.001.05 Vlrn"`
	SCVM        string                                `xml:"urn:swift:xsd:sese.018.001.05 SCVM"`
	Belgn       string                                `xml:"urn:swift:xsd:sese.018.001.05 Belgn"`
	Cmon        EuroclearClearstreamIdentifier        `xml:"urn:swift:xsd:sese.018.001.05 Cmon"`
	OthrPrtryId AlternateSecurityIdentification7      `xml:"urn:swift:xsd:sese.018.001.05 OthrPrtryId"`
}

type SubAccount5 struct {
	Id    Max35Text `xml:"urn:swift:xsd:sese.018.001.05 Id"`
	Nm    Max35Text `xml:"urn:swift:xsd:sese.018.001.05 Nm,omitempty"`
	Chrtc Max35Text `xml:"urn:swift:xsd:sese.018.001.05 Chrtc,omitempty"`
}

type SubscriptionInformation1 struct {
	DtOfFrstSbcpt ISODate                 `xml:"urn:swift:xsd:sese.018.001.05 DtOfFrstSbcpt"`
	EqtyCmpnt     ActiveCurrencyAndAmount `xml:"urn:swift:xsd:sese.018.001.05 EqtyCmpnt,omitempty"`
	CshCmpnt      ActiveCurrencyAndAmount `xml:"urn:swift:xsd:sese.018.001.05 CshCmpnt,omitempty"`
	TtlAmtYrToDt  ActiveCurrencyAndAmount `xml:"urn:swift:xsd:sese.018.001.05 TtlAmtYrToDt"`
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

type xsdGYearMonth time.Time

func (t *xsdGYearMonth) UnmarshalText(text []byte) error {
	return _unmarshalTime(text, (*time.Time)(t), "2006-01")
}
func (t xsdGYearMonth) MarshalText() ([]byte, error) {
	return _marshalTime((time.Time)(t), "2006-01")
}
func (t xsdGYearMonth) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if (time.Time)(t).IsZero() {
		return nil
	}
	m, err := t.MarshalText()
	if err != nil {
		return err
	}
	return e.EncodeElement(m, start)
}
func (t xsdGYearMonth) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if (time.Time)(t).IsZero() {
		return xml.Attr{}, nil
	}
	m, err := t.MarshalText()
	return xml.Attr{Name: name, Value: string(m)}, err
}