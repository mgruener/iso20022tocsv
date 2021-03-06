// Code generated by main. DO NOT EDIT.

package tsmt_011_001_04

import (
	"bytes"
	"encoding/xml"
	"time"
)

// May be one of SBTW, RSTW, RSBS, ARDM, ARCS, ARES, WAIT, UPDT, SBDS, ARBA, ARRO, CINR
type Action2Code string

type BICIdentification1 struct {
	BIC BICIdentifier `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.011.001.04 BIC"`
}

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type BICIdentifier string

type BaselineReportV04 struct {
	RptId             MessageIdentification1          `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.011.001.04 RptId"`
	RltdMsgRef        MessageIdentification1          `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.011.001.04 RltdMsgRef,omitempty"`
	RptTp             ReportType2                     `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.011.001.04 RptTp"`
	TxId              SimpleIdentificationInformation `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.011.001.04 TxId"`
	EstblishdBaselnId DocumentIdentification6         `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.011.001.04 EstblishdBaselnId"`
	TxSts             TransactionStatus4              `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.011.001.04 TxSts"`
	UsrTxRef          []DocumentIdentification5       `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.011.001.04 UsrTxRef,omitempty"`
	Buyr              PartyIdentification26           `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.011.001.04 Buyr"`
	Sellr             PartyIdentification26           `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.011.001.04 Sellr"`
	BuyrBk            BICIdentification1              `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.011.001.04 BuyrBk"`
	SellrBk           BICIdentification1              `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.011.001.04 SellrBk"`
	RptdLineItm       LineItem14                      `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.011.001.04 RptdLineItm"`
	ReqForActn        PendingActivity2                `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.011.001.04 ReqForActn,omitempty"`
}

// May be one of PROP, CLSD, PMTC, ESTD, ACTV, COMP, AMRQ, RARQ, CLRQ, SCRQ, SERQ, DARQ
type BaselineStatus3Code string

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type CurrencyAndAmount struct {
	Value float64      `xml:",chardata"`
	Ccy   CurrencyCode `xml:"Ccy,attr"`
}

// Must match the pattern [A-Z]{3,3}
type CurrencyCode string

type Document struct {
	BaselnRpt BaselineReportV04 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.011.001.04 BaselnRpt"`
}

type DocumentIdentification5 struct {
	Id     Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.011.001.04 Id"`
	IdIssr BICIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.011.001.04 IdIssr"`
}

type DocumentIdentification6 struct {
	Id          Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.011.001.04 Id"`
	Vrsn        float64         `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.011.001.04 Vrsn"`
	AmdmntSeqNb Max3NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.011.001.04 AmdmntSeqNb,omitempty"`
}

type GenericIdentification4 struct {
	Id   Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.011.001.04 Id"`
	IdTp Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.011.001.04 IdTp"`
}

type ISODateTime time.Time

func (t *ISODateTime) UnmarshalText(text []byte) error {
	return (*xsdDateTime)(t).UnmarshalText(text)
}
func (t ISODateTime) MarshalText() ([]byte, error) {
	return xsdDateTime(t).MarshalText()
}

type LineItem14 struct {
	LineItmDtls           []LineItemDetails12 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.011.001.04 LineItmDtls"`
	OrdrdLineItmsTtlAmt   CurrencyAndAmount   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.011.001.04 OrdrdLineItmsTtlAmt"`
	AccptdLineItmsTtlAmt  CurrencyAndAmount   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.011.001.04 AccptdLineItmsTtlAmt"`
	OutsdngLineItmsTtlAmt CurrencyAndAmount   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.011.001.04 OutsdngLineItmsTtlAmt"`
	PdgLineItmsTtlAmt     CurrencyAndAmount   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.011.001.04 PdgLineItmsTtlAmt"`
	OrdrdTtlNetAmt        CurrencyAndAmount   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.011.001.04 OrdrdTtlNetAmt"`
	AccptdTtlNetAmt       CurrencyAndAmount   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.011.001.04 AccptdTtlNetAmt"`
	OutsdngTtlNetAmt      CurrencyAndAmount   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.011.001.04 OutsdngTtlNetAmt"`
	PdgTtlNetAmt          CurrencyAndAmount   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.011.001.04 PdgTtlNetAmt"`
}

type LineItemDetails12 struct {
	LineItmId  Max70Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.011.001.04 LineItmId"`
	PdctNm     Max70Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.011.001.04 PdctNm,omitempty"`
	PdctIdr    []ProductIdentifier2Choice      `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.011.001.04 PdctIdr,omitempty"`
	PdctChrtcs []ProductCharacteristics1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.011.001.04 PdctChrtcs,omitempty"`
	PdctCtgy   []ProductCategory1Choice        `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.011.001.04 PdctCtgy,omitempty"`
	OrdrdQty   Quantity9                       `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.011.001.04 OrdrdQty"`
	AccptdQty  Quantity9                       `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.011.001.04 AccptdQty"`
	OutsdngQty Quantity9                       `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.011.001.04 OutsdngQty"`
	PdgQty     Quantity9                       `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.011.001.04 PdgQty"`
	QtyTlrnce  PercentageTolerance1            `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.011.001.04 QtyTlrnce,omitempty"`
	OrdrdAmt   CurrencyAndAmount               `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.011.001.04 OrdrdAmt"`
	AccptdAmt  CurrencyAndAmount               `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.011.001.04 AccptdAmt"`
	OutsdngAmt CurrencyAndAmount               `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.011.001.04 OutsdngAmt"`
	PdgAmt     CurrencyAndAmount               `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.011.001.04 PdgAmt"`
	PricTlrnce PercentageTolerance1            `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.011.001.04 PricTlrnce,omitempty"`
}

// Must be at least 1 items long
type Max140Text string

// Must match the pattern [0-9]{1,15}
type Max15NumericText string

// Must be at least 1 items long
type Max16Text string

// Must be at least 1 items long
type Max35Text string

// Must match the pattern [0-9]{1,3}
type Max3NumericText string

// Must be at least 1 items long
type Max70Text string

type MessageIdentification1 struct {
	Id      Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.011.001.04 Id"`
	CreDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.011.001.04 CreDtTm"`
}

type PartyIdentification26 struct {
	Nm      Max70Text              `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.011.001.04 Nm"`
	PrtryId GenericIdentification4 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.011.001.04 PrtryId,omitempty"`
	PstlAdr PostalAddress5         `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.011.001.04 PstlAdr"`
}

type PendingActivity2 struct {
	Tp   Action2Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.011.001.04 Tp"`
	Desc Max140Text  `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.011.001.04 Desc,omitempty"`
}

type PercentageTolerance1 struct {
	PlusPct float64 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.011.001.04 PlusPct"`
	MnsPct  float64 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.011.001.04 MnsPct"`
}

type PostalAddress5 struct {
	StrtNm      Max70Text   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.011.001.04 StrtNm,omitempty"`
	PstCdId     Max16Text   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.011.001.04 PstCdId,omitempty"`
	TwnNm       Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.011.001.04 TwnNm,omitempty"`
	CtrySubDvsn Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.011.001.04 CtrySubDvsn,omitempty"`
	Ctry        CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.011.001.04 Ctry"`
}

type ProductCategory1 struct {
	Tp   ProductCategory1Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.011.001.04 Tp"`
	Ctgy Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.011.001.04 Ctgy"`
}

type ProductCategory1Choice struct {
	StrdPdctCtgy ProductCategory1       `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.011.001.04 StrdPdctCtgy"`
	OthrPdctCtgy GenericIdentification4 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.011.001.04 OthrPdctCtgy"`
}

// May be one of HRTR, QOTA, PRGP, LOBU, GNDR
type ProductCategory1Code string

type ProductCharacteristics1 struct {
	Tp     ProductCharacteristics1Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.011.001.04 Tp"`
	Chrtcs Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.011.001.04 Chrtcs"`
}

type ProductCharacteristics1Choice struct {
	StrdPdctChrtcs ProductCharacteristics1 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.011.001.04 StrdPdctChrtcs"`
	OthrPdctChrtcs GenericIdentification4  `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.011.001.04 OthrPdctChrtcs"`
}

// May be one of BISP, CHNR, CLOR, EDSP, ENNR, OPTN, ORCR, PCTV, SISP, SIZE, SZRG, SPRM, STOR, VINR
type ProductCharacteristics1Code string

type ProductIdentifier2 struct {
	Tp  ProductIdentifier2Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.011.001.04 Tp"`
	Idr Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.011.001.04 Idr"`
}

type ProductIdentifier2Choice struct {
	StrdPdctIdr ProductIdentifier2     `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.011.001.04 StrdPdctIdr"`
	OthrPdctIdr GenericIdentification4 `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.011.001.04 OthrPdctIdr"`
}

// May be one of BINR, COMD, EANC, HRTR, MANI, MODL, PART, QOTA, STYL, SUPI, UPCC
type ProductIdentifier2Code string

type Quantity9 struct {
	UnitOfMeasr UnitOfMeasure3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.011.001.04 UnitOfMeasr"`
	Val         float64              `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.011.001.04 Val"`
	Fctr        Max15NumericText     `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.011.001.04 Fctr,omitempty"`
}

type ReportType2 struct {
	Tp ReportType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.011.001.04 Tp"`
}

// May be one of PREC, CURR
type ReportType2Code string

type SimpleIdentificationInformation struct {
	Id Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.011.001.04 Id"`
}

type TransactionStatus4 struct {
	Sts BaselineStatus3Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.011.001.04 Sts"`
}

type UnitOfMeasure3Choice struct {
	UnitOfMeasrCd   UnitOfMeasure4Code `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.011.001.04 UnitOfMeasrCd"`
	OthrUnitOfMeasr Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.011.001.04 OthrUnitOfMeasr"`
}

// May be one of KGM, EA, LTN, MTR, INH, LY, GLI, GRM, CMT, MTK, FOT, 1A, INK, FTK, MIK, ONZ, PTI, PT, QTI, QT, GLL, MMT, KTM, YDK, MMK, CMK, KMK, MMQ, CLT, LTR, LBR, STN, BLL, BX, BO, CT, CH, CR, INQ, MTQ, OZI, OZA, BG, BL, TNE
type UnitOfMeasure4Code string

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
