// Code generated by main. DO NOT EDIT.

package seev_038_001_02

import (
	"bytes"
	"encoding/xml"
	"time"
)

type AccountIdentification10 struct {
	IdCd SafekeepingAccountIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.038.001.02 IdCd"`
}

type AccountIdentification14Choice struct {
	ForAllAccts         AccountIdentification10   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.038.001.02 ForAllAccts"`
	AcctsListAndBalDtls []AccountIdentification18 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.038.001.02 AcctsListAndBalDtls"`
}

type AccountIdentification18 struct {
	SfkpgAcct Max35Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.038.001.02 SfkpgAcct"`
	AcctOwnr  PartyIdentification36Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.038.001.02 AcctOwnr,omitempty"`
	SfkpgPlc  SafekeepingPlaceFormat2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.038.001.02 SfkpgPlc,omitempty"`
	ConfdBal  BalanceFormat1Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.038.001.02 ConfdBal"`
}

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type AnyBICIdentifier string

type BalanceFormat1Choice struct {
	Bal         SignedQuantityFormat1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.038.001.02 Bal"`
	ElgblBal    SignedQuantityFormat2 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.038.001.02 ElgblBal"`
	NotElgblBal SignedQuantityFormat2 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.038.001.02 NotElgblBal"`
}

type CorporateActionGeneralInformation10 struct {
	CorpActnEvtId      Max35Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.038.001.02 CorpActnEvtId"`
	OffclCorpActnEvtId Max35Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.038.001.02 OffclCorpActnEvtId,omitempty"`
	NrrtvTp            CorporateActionNarrative1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.038.001.02 NrrtvTp,omitempty"`
}

type CorporateActionNarrative1Choice struct {
	Cd    CorporateActionNarrative1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.038.001.02 Cd"`
	Prtry GenericIdentification20       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.038.001.02 Prtry"`
}

// May be one of TAXE, REGI, WTRC, RFMC, PAUT, CTIN
type CorporateActionNarrative1Code string

type CorporateActionNarrativeV02 struct {
	AcctDtls       AccountIdentification14Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.038.001.02 AcctDtls,omitempty"`
	UndrlygScty    UnderlyingSecurity3                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.038.001.02 UndrlygScty,omitempty"`
	CorpActnGnlInf CorporateActionGeneralInformation10 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.038.001.02 CorpActnGnlInf"`
	AddtlInf       UpdatedAdditionalInformation2       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.038.001.02 AddtlInf"`
	SplmtryData    []SupplementaryData1                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.038.001.02 SplmtryData,omitempty"`
}

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type Document struct {
	CorpActnNrrtv CorporateActionNarrativeV02 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.038.001.02 CorpActnNrrtv"`
}

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

// Must be at least 1 items long
type ExternalFinancialInstrumentIdentificationType1Code string

type FinancialInstrumentQuantity1Choice struct {
	Unit     float64 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.038.001.02 Unit"`
	FaceAmt  float64 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.038.001.02 FaceAmt"`
	AmtsdVal float64 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.038.001.02 AmtsdVal"`
}

type GenericIdentification19 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.038.001.02 Id"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.038.001.02 Issr"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.038.001.02 SchmeNm,omitempty"`
}

type GenericIdentification20 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:seev.038.001.02 Id"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.038.001.02 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.038.001.02 SchmeNm,omitempty"`
}

type GenericIdentification21 struct {
	Tp GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.038.001.02 Tp"`
	Id Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.038.001.02 Id,omitempty"`
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

type IdentificationSource3Choice struct {
	Cd    ExternalFinancialInstrumentIdentificationType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.038.001.02 Cd"`
	Prtry Max35Text                                          `xml:"urn:iso:std:iso:20022:tech:xsd:seev.038.001.02 Prtry"`
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
type Max8000Text string

type OtherIdentification1 struct {
	Id  Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.038.001.02 Id"`
	Sfx Max16Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:seev.038.001.02 Sfx,omitempty"`
	Tp  IdentificationSource3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.038.001.02 Tp"`
}

type PartyIdentification36Choice struct {
	AnyBIC  AnyBICIdentifier        `xml:"urn:iso:std:iso:20022:tech:xsd:seev.038.001.02 AnyBIC"`
	PrtryId GenericIdentification19 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.038.001.02 PrtryId"`
}

type ProprietaryQuantity2 struct {
	Qty     float64                `xml:"urn:iso:std:iso:20022:tech:xsd:seev.038.001.02 Qty"`
	QtyTp   Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:seev.038.001.02 QtyTp"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.038.001.02 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.038.001.02 SchmeNm,omitempty"`
}

type Quantity2Choice struct {
	Qty      FinancialInstrumentQuantity1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.038.001.02 Qty"`
	PrtryQty ProprietaryQuantity2               `xml:"urn:iso:std:iso:20022:tech:xsd:seev.038.001.02 PrtryQty"`
}

// May be one of GENR
type SafekeepingAccountIdentification1Code string

// May be one of CUST, ICSD, NCSD, SHHE
type SafekeepingPlace1Code string

// May be one of SHHE, ALLP
type SafekeepingPlace2Code string

type SafekeepingPlaceFormat2Choice struct {
	Id      SafekeepingPlaceTypeAndText2             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.038.001.02 Id"`
	Ctry    CountryCode                              `xml:"urn:iso:std:iso:20022:tech:xsd:seev.038.001.02 Ctry"`
	TpAndId SafekeepingPlaceTypeAndAnyBICIdentifier1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.038.001.02 TpAndId"`
	Prtry   GenericIdentification21                  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.038.001.02 Prtry"`
}

type SafekeepingPlaceTypeAndAnyBICIdentifier1 struct {
	SfkpgPlcTp SafekeepingPlace1Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.038.001.02 SfkpgPlcTp"`
	Id         AnyBICIdentifier      `xml:"urn:iso:std:iso:20022:tech:xsd:seev.038.001.02 Id"`
}

type SafekeepingPlaceTypeAndText2 struct {
	SfkpgPlcTp SafekeepingPlace2Code `xml:"urn:iso:std:iso:20022:tech:xsd:seev.038.001.02 SfkpgPlcTp"`
	Id         Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.038.001.02 Id,omitempty"`
}

type SecurityIdentification14 struct {
	ISIN   ISINIdentifier         `xml:"urn:iso:std:iso:20022:tech:xsd:seev.038.001.02 ISIN,omitempty"`
	OthrId []OtherIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.038.001.02 OthrId,omitempty"`
	Desc   Max140Text             `xml:"urn:iso:std:iso:20022:tech:xsd:seev.038.001.02 Desc,omitempty"`
}

// May be one of SHOR, LONG
type ShortLong1Code string

type SignedQuantityFormat1 struct {
	ShrtLngPos ShortLong1Code  `xml:"urn:iso:std:iso:20022:tech:xsd:seev.038.001.02 ShrtLngPos"`
	QtyChc     Quantity2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.038.001.02 QtyChc"`
}

type SignedQuantityFormat2 struct {
	ShrtLngPos ShortLong1Code                     `xml:"urn:iso:std:iso:20022:tech:xsd:seev.038.001.02 ShrtLngPos"`
	Qty        FinancialInstrumentQuantity1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:seev.038.001.02 Qty"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.038.001.02 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.038.001.02 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

type UnderlyingSecurity3 struct {
	SctyId SecurityIdentification14 `xml:"urn:iso:std:iso:20022:tech:xsd:seev.038.001.02 SctyId"`
}

type UpdatedAdditionalInformation2 struct {
	UpdDesc  Max140Text    `xml:"urn:iso:std:iso:20022:tech:xsd:seev.038.001.02 UpdDesc,omitempty"`
	UpdDt    ISODate       `xml:"urn:iso:std:iso:20022:tech:xsd:seev.038.001.02 UpdDt,omitempty"`
	AddtlInf []Max8000Text `xml:"urn:iso:std:iso:20022:tech:xsd:seev.038.001.02 AddtlInf"`
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