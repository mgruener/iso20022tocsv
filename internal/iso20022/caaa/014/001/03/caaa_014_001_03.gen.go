// Code generated by main. DO NOT EDIT.

package caaa_014_001_03

import (
	"bytes"
	"encoding/base64"
	"encoding/xml"
	"time"
)

type AcceptorDiagnosticResponse3 struct {
	Envt     CardPaymentEnvironment31 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.014.001.03 Envt"`
	TMSTrggr TMSTrigger1              `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.014.001.03 TMSTrggr,omitempty"`
}

type AcceptorDiagnosticResponseV03 struct {
	Hdr        Header7                     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.014.001.03 Hdr"`
	DgnstcRspn AcceptorDiagnosticResponse3 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.014.001.03 DgnstcRspn"`
	SctyTrlr   ContentInformationType8     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.014.001.03 SctyTrlr"`
}

type Acquirer2 struct {
	Id         GenericIdentification32 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.014.001.03 Id,omitempty"`
	ParamsVrsn Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.014.001.03 ParamsVrsn"`
}

// May be one of MACC, MCCS, CMA1
type Algorithm10Code string

// May be one of HS25, HS38, HS51
type Algorithm5Code string

// May be one of ERSA, RSAO
type Algorithm7Code string

// May be one of MGF1
type Algorithm8Code string

// May be one of EA2C, E3DC, DKP9, UKPT, UKA1
type Algorithm9Code string

type AlgorithmIdentification10 struct {
	Algo  Algorithm10Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.014.001.03 Algo"`
	Param Parameter1      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.014.001.03 Param,omitempty"`
}

type AlgorithmIdentification7 struct {
	Algo  Algorithm7Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.014.001.03 Algo"`
	Param Parameter2     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.014.001.03 Param,omitempty"`
}

type AlgorithmIdentification8 struct {
	Algo  Algorithm8Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.014.001.03 Algo"`
	Param Parameter3     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.014.001.03 Param,omitempty"`
}

type AlgorithmIdentification9 struct {
	Algo  Algorithm9Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.014.001.03 Algo"`
	Param Parameter1     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.014.001.03 Param,omitempty"`
}

// May be one of CNAT, LATT, OATT, OUAT, CATT
type AttributeType1Code string

type AuthenticatedData3 struct {
	Vrsn        float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.014.001.03 Vrsn,omitempty"`
	Rcpt        []Recipient3Choice        `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.014.001.03 Rcpt"`
	MACAlgo     AlgorithmIdentification10 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.014.001.03 MACAlgo"`
	NcpsltdCntt EncapsulatedContent2      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.014.001.03 NcpsltdCntt"`
	MAC         Max35Binary               `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.014.001.03 MAC"`
}

type CardPaymentEnvironment31 struct {
	Acqrr      Acquirer2               `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.014.001.03 Acqrr"`
	AcqrrAvlbl bool                    `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.014.001.03 AcqrrAvlbl,omitempty"`
	MrchntId   GenericIdentification32 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.014.001.03 MrchntId,omitempty"`
	POIId      GenericIdentification32 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.014.001.03 POIId,omitempty"`
}

type CertificateIdentifier1 struct {
	IssrAndSrlNb IssuerAndSerialNumber1 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.014.001.03 IssrAndSrlNb"`
}

type CertificateIssuer1 struct {
	RltvDstngshdNm []RelativeDistinguishedName1 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.014.001.03 RltvDstngshdNm"`
}

type ContentInformationType8 struct {
	CnttTp       ContentType1Code     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.014.001.03 CnttTp"`
	AuthntcdData []AuthenticatedData3 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.014.001.03 AuthntcdData,omitempty"`
}

// May be one of DATA, SIGN, EVLP, DGST, ECRP, AUTH
type ContentType1Code string

type Document struct {
	AccptrDgnstcRspn AcceptorDiagnosticResponseV03 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.014.001.03 AccptrDgnstcRspn"`
}

type EncapsulatedContent2 struct {
	CnttTp ContentType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.014.001.03 CnttTp"`
	Cntt   Max100KBinary    `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.014.001.03 Cntt,omitempty"`
}

// May be no more than 10 items long
type Exact10Text string

type GenericIdentification31 struct {
	Id     Max35Text      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.014.001.03 Id"`
	Tp     PartyType3Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.014.001.03 Tp"`
	Issr   PartyType4Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.014.001.03 Issr,omitempty"`
	ShrtNm Max35Text      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.014.001.03 ShrtNm,omitempty"`
}

type GenericIdentification32 struct {
	Id     Max35Text      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.014.001.03 Id"`
	Tp     PartyType3Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.014.001.03 Tp,omitempty"`
	Issr   PartyType4Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.014.001.03 Issr,omitempty"`
	ShrtNm Max35Text      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.014.001.03 ShrtNm,omitempty"`
}

type Header7 struct {
	MsgFctn    MessageFunction4Code    `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.014.001.03 MsgFctn"`
	PrtcolVrsn Max6Text                `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.014.001.03 PrtcolVrsn"`
	XchgId     Max3NumericText         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.014.001.03 XchgId"`
	CreDtTm    ISODateTime             `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.014.001.03 CreDtTm"`
	InitgPty   GenericIdentification32 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.014.001.03 InitgPty"`
	RcptPty    GenericIdentification32 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.014.001.03 RcptPty,omitempty"`
	Tracblt    []Traceability1         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.014.001.03 Tracblt,omitempty"`
}

type ISODateTime time.Time

func (t *ISODateTime) UnmarshalText(text []byte) error {
	return (*xsdDateTime)(t).UnmarshalText(text)
}
func (t ISODateTime) MarshalText() ([]byte, error) {
	return xsdDateTime(t).MarshalText()
}

type IssuerAndSerialNumber1 struct {
	Issr  CertificateIssuer1 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.014.001.03 Issr"`
	SrlNb Max35Binary        `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.014.001.03 SrlNb"`
}

type KEK3 struct {
	Vrsn          float64                  `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.014.001.03 Vrsn,omitempty"`
	KEKId         KEKIdentifier1           `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.014.001.03 KEKId"`
	KeyNcrptnAlgo AlgorithmIdentification9 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.014.001.03 KeyNcrptnAlgo"`
	NcrptdKey     Max140Binary             `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.014.001.03 NcrptdKey"`
}

type KEKIdentifier1 struct {
	KeyId     Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.014.001.03 KeyId"`
	KeyVrsn   Exact10Text     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.014.001.03 KeyVrsn"`
	DerivtnId Min5Max16Binary `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.014.001.03 DerivtnId,omitempty"`
}

type KeyTransport3 struct {
	Vrsn          float64                  `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.014.001.03 Vrsn,omitempty"`
	RcptId        CertificateIdentifier1   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.014.001.03 RcptId"`
	KeyNcrptnAlgo AlgorithmIdentification7 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.014.001.03 KeyNcrptnAlgo"`
	NcrptdKey     Max3000Binary            `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.014.001.03 NcrptdKey"`
}

type Max100KBinary []byte

func (t *Max100KBinary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Max100KBinary) MarshalText() ([]byte, error) {
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

type Max3000Binary []byte

func (t *Max3000Binary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Max3000Binary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

type Max35Binary []byte

func (t *Max35Binary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Max35Binary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

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

// May be one of AUTQ, AUTP, FAUQ, FAUP, CMPV, CMPK, FCMV, FCMK, RVRA, RVRR, FRVA, FRVR, CCAQ, CCAP, CCAV, CCAK, DGNP, DGNQ, RCLQ, RCLP, RJCT, DCCQ, DCCP
type MessageFunction4Code string

type Min5Max16Binary []byte

func (t *Min5Max16Binary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Min5Max16Binary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

type Parameter1 struct {
	InitlstnVctr Max500Binary `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.014.001.03 InitlstnVctr,omitempty"`
}

type Parameter2 struct {
	DgstAlgo     Algorithm5Code           `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.014.001.03 DgstAlgo,omitempty"`
	MskGnrtrAlgo AlgorithmIdentification8 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.014.001.03 MskGnrtrAlgo,omitempty"`
}

type Parameter3 struct {
	DgstAlgo Algorithm5Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.014.001.03 DgstAlgo,omitempty"`
}

// May be one of OPOI, MERC, ACCP, ITAG, ACQR, CISS, DLIS
type PartyType3Code string

// May be one of MERC, ACCP, ITAG, ACQR, CISS, TAXH
type PartyType4Code string

type Recipient3Choice struct {
	KeyTrnsprt KeyTransport3  `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.014.001.03 KeyTrnsprt"`
	KEK        KEK3           `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.014.001.03 KEK"`
	KeyIdr     KEKIdentifier1 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.014.001.03 KeyIdr"`
}

type RelativeDistinguishedName1 struct {
	AttrTp  AttributeType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.014.001.03 AttrTp"`
	AttrVal Max140Text         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.014.001.03 AttrVal"`
}

// May be one of CRIT, ASAP, DTIM
type TMSContactLevel1Code string

type TMSTrigger1 struct {
	TMSCtctLvl  TMSContactLevel1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.014.001.03 TMSCtctLvl"`
	TMSId       Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.014.001.03 TMSId,omitempty"`
	TMSCtctDtTm ISODateTime          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.014.001.03 TMSCtctDtTm,omitempty"`
}

type Traceability1 struct {
	RlayId      GenericIdentification31 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.014.001.03 RlayId"`
	TracDtTmIn  ISODateTime             `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.014.001.03 TracDtTmIn"`
	TracDtTmOut ISODateTime             `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.014.001.03 TracDtTmOut"`
}

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
