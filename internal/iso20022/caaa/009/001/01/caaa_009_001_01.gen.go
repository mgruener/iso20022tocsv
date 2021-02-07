// Code generated by main. DO NOT EDIT.

package caaa_009_001_01

import (
	"bytes"
	"encoding/base64"
	"encoding/xml"
	"time"
)

type AcceptorReconciliationRequest1 struct {
	Envt CardPaymentEnvironment7    `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.009.001.01 Envt"`
	Tx   TransactionReconciliation1 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.009.001.01 Tx"`
}

type AcceptorReconciliationRequestV01 struct {
	Hdr        Header1                        `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.009.001.01 Hdr"`
	RcncltnReq AcceptorReconciliationRequest1 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.009.001.01 RcncltnReq"`
	SctyTrlr   ContentInformationType3        `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.009.001.01 SctyTrlr"`
}

type Acquirer1 struct {
	Id         GenericIdentification32 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.009.001.01 Id,omitempty"`
	ParamsVrsn ISODateTime             `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.009.001.01 ParamsVrsn"`
}

// May be one of MACC, MCCS, UKPT, DKPT, E3DC, HS25, ERS2, ERSA
type Algorithm1Code string

type AlgorithmIdentification1 struct {
	Algo  Algorithm1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.009.001.01 Algo"`
	Param Parameter1     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.009.001.01 Param,omitempty"`
}

// May be one of CNAT, LATT, OATT, OUAT, CATT
type AttributeType1Code string

type AuthenticatedData1 struct {
	Vrsn        float64                  `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.009.001.01 Vrsn,omitempty"`
	Rcpt        []Recipient1Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.009.001.01 Rcpt"`
	MACAlgo     AlgorithmIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.009.001.01 MACAlgo"`
	NcpsltdCntt EncapsulatedContent1     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.009.001.01 NcpsltdCntt"`
	MAC         Max35Binary              `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.009.001.01 MAC"`
}

type CardPaymentEnvironment7 struct {
	Acqrr    Acquirer1               `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.009.001.01 Acqrr"`
	MrchntId GenericIdentification32 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.009.001.01 MrchntId,omitempty"`
	POIId    GenericIdentification32 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.009.001.01 POIId,omitempty"`
}

type CertificateIdentifier1 struct {
	IssrAndSrlNb IssuerAndSerialNumber1 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.009.001.01 IssrAndSrlNb"`
}

type CertificateIssuer1 struct {
	RltvDstngshdNm []RelativeDistinguishedName1 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.009.001.01 RltvDstngshdNm"`
}

type ContentInformationType3 struct {
	CnttTp       ContentType1Code     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.009.001.01 CnttTp"`
	AuthntcdData []AuthenticatedData1 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.009.001.01 AuthntcdData,omitempty"`
}

// May be one of DATA, SIGN, EVLP, DGST, ECRP, AUTH
type ContentType1Code string

// Must match the pattern [A-Z]{3,3}
type CurrencyCode string

type Document struct {
	AccptrRcncltnReq AcceptorReconciliationRequestV01 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.009.001.01 AccptrRcncltnReq"`
}

type EncapsulatedContent1 struct {
	CnttTp ContentType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.009.001.01 CnttTp"`
	Cntt   Max10000Binary   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.009.001.01 Cntt,omitempty"`
}

// May be no more than 10 items long
type Exact10Text string

// Must match the pattern [0-9]{4}
type Exact4NumericText string

type GenericIdentification31 struct {
	Id     Max35Text      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.009.001.01 Id"`
	Tp     PartyType3Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.009.001.01 Tp"`
	Issr   PartyType4Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.009.001.01 Issr,omitempty"`
	ShrtNm Max35Text      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.009.001.01 ShrtNm,omitempty"`
}

type GenericIdentification32 struct {
	Id     Max35Text      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.009.001.01 Id"`
	Tp     PartyType3Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.009.001.01 Tp,omitempty"`
	Issr   PartyType4Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.009.001.01 Issr,omitempty"`
	ShrtNm Max35Text      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.009.001.01 ShrtNm,omitempty"`
}

type Header1 struct {
	MsgFctn    MessageFunction1Code    `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.009.001.01 MsgFctn"`
	PrtcolVrsn Max6Text                `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.009.001.01 PrtcolVrsn"`
	XchgId     Max3NumericText         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.009.001.01 XchgId"`
	CreDtTm    ISODateTime             `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.009.001.01 CreDtTm"`
	InitgPty   GenericIdentification32 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.009.001.01 InitgPty"`
	RcptPty    GenericIdentification32 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.009.001.01 RcptPty,omitempty"`
	Tracblt    []Traceability1         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.009.001.01 Tracblt,omitempty"`
}

type ISODateTime time.Time

func (t *ISODateTime) UnmarshalText(text []byte) error {
	return (*xsdDateTime)(t).UnmarshalText(text)
}
func (t ISODateTime) MarshalText() ([]byte, error) {
	return xsdDateTime(t).MarshalText()
}

type IssuerAndSerialNumber1 struct {
	Issr  CertificateIssuer1 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.009.001.01 Issr"`
	SrlNb Max35Binary        `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.009.001.01 SrlNb"`
}

type KEK1 struct {
	Vrsn          float64                  `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.009.001.01 Vrsn,omitempty"`
	KEKId         KEKIdentifier1           `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.009.001.01 KEKId"`
	KeyNcrptnAlgo AlgorithmIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.009.001.01 KeyNcrptnAlgo"`
	NcrptdKey     Max140Binary             `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.009.001.01 NcrptdKey"`
}

type KEKIdentifier1 struct {
	KeyId     Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.009.001.01 KeyId"`
	KeyVrsn   Exact10Text     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.009.001.01 KeyVrsn"`
	DerivtnId Min5Max16Binary `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.009.001.01 DerivtnId,omitempty"`
}

type KeyTransport1 struct {
	Vrsn          float64                  `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.009.001.01 Vrsn"`
	RcptId        CertificateIdentifier1   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.009.001.01 RcptId"`
	KeyNcrptnAlgo AlgorithmIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.009.001.01 KeyNcrptnAlgo"`
	NcrptdKey     Max140Binary             `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.009.001.01 NcrptdKey"`
}

type Max10000Binary []byte

func (t *Max10000Binary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Max10000Binary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

type Max140Binary []byte

func (t *Max140Binary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Max140Binary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

// Must be at least 1 items long
type Max140Text string

type Max35Binary []byte

func (t *Max35Binary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Max35Binary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

// Must match the pattern [0-9]{1,35}
type Max35NumericText string

// Must be at least 1 items long
type Max35Text string

// Must match the pattern [0-9]{1,3}
type Max3NumericText string

type Max500Binary []byte

func (t *Max500Binary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Max500Binary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

// Must be at least 1 items long
type Max6Text string

// Must be at least 1 items long
type Max70Text string

// May be one of AUTQ, AUTP, FAUQ, FAUP, CMPV, CMPK, FCMV, FCMK, RVRA, RVRR, FRVA, FRVR, CCAQ, CCAP, CCAV, CCAK, DGNP, DGNQ, RCLQ, RCLP, RJCT
type MessageFunction1Code string

type Min5Max16Binary []byte

func (t *Min5Max16Binary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Min5Max16Binary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

type Parameter1 struct {
	InitlstnVctr Max500Binary `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.009.001.01 InitlstnVctr,omitempty"`
}

// May be one of OPOI, MERC, ACCP, ITAG, ACQR, CISS, DLIS
type PartyType3Code string

// May be one of MERC, ACCP, ITAG, ACQR, CISS, TAXH
type PartyType4Code string

type Recipient1Choice struct {
	KeyTrnsprt KeyTransport1 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.009.001.01 KeyTrnsprt,omitempty"`
	KEK        KEK1          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.009.001.01 KEK,omitempty"`
}

type RelativeDistinguishedName1 struct {
	AttrTp  AttributeType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.009.001.01 AttrTp"`
	AttrVal Max140Text         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.009.001.01 AttrVal"`
}

type Traceability1 struct {
	RlayId      GenericIdentification31 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.009.001.01 RlayId"`
	TracDtTmIn  ISODateTime             `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.009.001.01 TracDtTmIn"`
	TracDtTmOut ISODateTime             `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.009.001.01 TracDtTmOut"`
}

type TransactionIdentifier1 struct {
	TxDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.009.001.01 TxDtTm"`
	TxRef  Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.009.001.01 TxRef"`
}

type TransactionReconciliation1 struct {
	ClsPrd      bool                   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.009.001.01 ClsPrd,omitempty"`
	RcncltnTxId TransactionIdentifier1 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.009.001.01 RcncltnTxId"`
	RcncltnId   Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.009.001.01 RcncltnId"`
	TxTtls      []TransactionTotals1   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.009.001.01 TxTtls"`
	AddtlTxData Max70Text              `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.009.001.01 AddtlTxData,omitempty"`
}

type TransactionTotals1 struct {
	POIGrpId     Max35Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.009.001.01 POIGrpId,omitempty"`
	CardPdctPrfl Exact4NumericText          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.009.001.01 CardPdctPrfl,omitempty"`
	Ccy          CurrencyCode               `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.009.001.01 Ccy,omitempty"`
	Tp           TypeTransactionTotals1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.009.001.01 Tp"`
	TtlNb        Max35NumericText           `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.009.001.01 TtlNb"`
	CmltvAmt     float64                    `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.009.001.01 CmltvAmt"`
}

// May be one of DEBT, DBTR, CRDT, CRDR
type TypeTransactionTotals1Code string

type xsdBase64Binary []byte

func (b *xsdBase64Binary) UnmarshalText(text []byte) (err error) {
	*b, err = base64.StdEncoding.DecodeString(string(text))
	return
}
func (b xsdBase64Binary) MarshalText() ([]byte, error) {
	var buf bytes.Buffer
	enc := base64.NewEncoder(base64.StdEncoding, &buf)
	enc.Write([]byte(b))
	enc.Close()
	return buf.Bytes(), nil
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