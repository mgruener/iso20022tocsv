// Code generated by main. DO NOT EDIT.

package caam_007_001_01

import (
	"bytes"
	"encoding/base64"
	"encoding/xml"
	"time"
)

// ATMCommandIdentification1
//
// Identification of the entity issuing the command.
type ATMCommandIdentification1 struct {
	Orgn Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:caam.007.001.01 Orgn,omitempty"`
	Ref  Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:caam.007.001.01 Ref,omitempty"`
	Prcr Max140Text `xml:"urn:iso:std:iso:20022:tech:xsd:caam.007.001.01 Prcr,omitempty"`
}

// ATMEnvironment9
//
// Environment of the ATM.
type ATMEnvironment9 struct {
	Acqrr    Acquirer7               `xml:"urn:iso:std:iso:20022:tech:xsd:caam.007.001.01 Acqrr,omitempty"`
	ATMMgrId Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:caam.007.001.01 ATMMgrId,omitempty"`
	ATM      AutomatedTellerMachine7 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.007.001.01 ATM"`
}

// ATMMessageFunction1
//
// Identifies the type of process related to an ATM message.
type ATMMessageFunction1 struct {
	Fctn     MessageFunction7Code `xml:"urn:iso:std:iso:20022:tech:xsd:caam.007.001.01 Fctn"`
	ATMSvcCd Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:caam.007.001.01 ATMSvcCd,omitempty"`
	HstSvcCd Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:caam.007.001.01 HstSvcCd,omitempty"`
}

// Acquirer7
//
// Acquirer of the withdrawal transaction, in charge of the funds settlement with the issuer.
type Acquirer7 struct {
	AcqrgInstn Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:caam.007.001.01 AcqrgInstn,omitempty"`
	Brnch      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:caam.007.001.01 Brnch,omitempty"`
}

// May be one of HS25, HS38, HS51, HS01
type Algorithm11Code string

// May be one of MACC, MCCS, CMA1, MCC1, CMA9, CMA5
type Algorithm12Code string

// May be one of EA2C, E3DC, DKP9, UKPT, UKA1, EA9C, EA5C
type Algorithm13Code string

// May be one of EA2C, E3DC, EA9C, EA5C
type Algorithm15Code string

// May be one of ERSA, RSAO
type Algorithm7Code string

// May be one of MGF1
type Algorithm8Code string

// AlgorithmIdentification11
//
// Cryptographic algorithms and parameters for the protection of transported keys by an asymmetric key.
type AlgorithmIdentification11 struct {
	Algo  Algorithm7Code `xml:"urn:iso:std:iso:20022:tech:xsd:caam.007.001.01 Algo"`
	Param Parameter4     `xml:"urn:iso:std:iso:20022:tech:xsd:caam.007.001.01 Param,omitempty"`
}

// AlgorithmIdentification12
//
// Mask generator function cryptographic algorithm and parameters.
type AlgorithmIdentification12 struct {
	Algo  Algorithm8Code `xml:"urn:iso:std:iso:20022:tech:xsd:caam.007.001.01 Algo"`
	Param Parameter5     `xml:"urn:iso:std:iso:20022:tech:xsd:caam.007.001.01 Param,omitempty"`
}

// AlgorithmIdentification13
//
// Cryptographic algorithm and parameters for the protection of the transported key.
type AlgorithmIdentification13 struct {
	Algo  Algorithm13Code `xml:"urn:iso:std:iso:20022:tech:xsd:caam.007.001.01 Algo"`
	Param Parameter6      `xml:"urn:iso:std:iso:20022:tech:xsd:caam.007.001.01 Param,omitempty"`
}

// AlgorithmIdentification14
//
// Cryptographic algorithm and parameters for encryptions with a symmetric cryptographic key.
type AlgorithmIdentification14 struct {
	Algo  Algorithm15Code `xml:"urn:iso:std:iso:20022:tech:xsd:caam.007.001.01 Algo"`
	Param Parameter6      `xml:"urn:iso:std:iso:20022:tech:xsd:caam.007.001.01 Param,omitempty"`
}

// AlgorithmIdentification15
//
// Identification of a cryptographic algorithm and parameters for the MAC computation.
type AlgorithmIdentification15 struct {
	Algo  Algorithm12Code `xml:"urn:iso:std:iso:20022:tech:xsd:caam.007.001.01 Algo"`
	Param Parameter7      `xml:"urn:iso:std:iso:20022:tech:xsd:caam.007.001.01 Param,omitempty"`
}

// May be one of CNAT, LATT, OATT, OUAT, CATT
type AttributeType1Code string

// AuthenticatedData4
//
// Message authentication code (MAC), computed on the data to protect with an encryption key.
type AuthenticatedData4 struct {
	Vrsn        float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:caam.007.001.01 Vrsn,omitempty"`
	Rcpt        []Recipient4Choice        `xml:"urn:iso:std:iso:20022:tech:xsd:caam.007.001.01 Rcpt"`
	MACAlgo     AlgorithmIdentification15 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.007.001.01 MACAlgo"`
	NcpsltdCntt EncapsulatedContent3      `xml:"urn:iso:std:iso:20022:tech:xsd:caam.007.001.01 NcpsltdCntt"`
	MAC         Max140Binary              `xml:"urn:iso:std:iso:20022:tech:xsd:caam.007.001.01 MAC"`
}

// AutomatedTellerMachine7
//
// ATM information.
type AutomatedTellerMachine7 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:caam.007.001.01 Id"`
	AddtlId Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:caam.007.001.01 AddtlId,omitempty"`
	SeqNb   Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:caam.007.001.01 SeqNb,omitempty"`
}

// May be one of LNGT, NUL8, NULG, NULL, RAND
type BytePadding1Code string

// CertificateIssuer1
//
// Certificate issuer name (see X.509).
type CertificateIssuer1 struct {
	RltvDstngshdNm []RelativeDistinguishedName1 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.007.001.01 RltvDstngshdNm"`
}

// ContentInformationType10
//
// General cryptographic message syntax (CMS) containing encrypted data.
type ContentInformationType10 struct {
	CnttTp     ContentType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:caam.007.001.01 CnttTp"`
	EnvlpdData EnvelopedData4   `xml:"urn:iso:std:iso:20022:tech:xsd:caam.007.001.01 EnvlpdData"`
}

// ContentInformationType15
//
// General cryptographic message syntax (CMS) containing authenticated data.
type ContentInformationType15 struct {
	CnttTp       ContentType2Code   `xml:"urn:iso:std:iso:20022:tech:xsd:caam.007.001.01 CnttTp"`
	AuthntcdData AuthenticatedData4 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.007.001.01 AuthntcdData"`
}

// May be one of DATA, SIGN, EVLP, DGST, AUTH
type ContentType2Code string

type Document struct {
	HstToATMReq HostToATMRequestV01 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.007.001.01 HstToATMReq"`
}

// EncapsulatedContent3
//
// Data to authenticate.
type EncapsulatedContent3 struct {
	CnttTp ContentType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:caam.007.001.01 CnttTp"`
	Cntt   Max100KBinary    `xml:"urn:iso:std:iso:20022:tech:xsd:caam.007.001.01 Cntt,omitempty"`
}

// EncryptedContent3
//
// Encrypted data with an encryption key.
type EncryptedContent3 struct {
	CnttTp         ContentType2Code          `xml:"urn:iso:std:iso:20022:tech:xsd:caam.007.001.01 CnttTp"`
	CnttNcrptnAlgo AlgorithmIdentification14 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.007.001.01 CnttNcrptnAlgo"`
	NcrptdData     Max100KBinary             `xml:"urn:iso:std:iso:20022:tech:xsd:caam.007.001.01 NcrptdData"`
}

// May be one of TR31, TR34
type EncryptionFormat1Code string

// EnvelopedData4
//
// Encrypted data with encryption key.
type EnvelopedData4 struct {
	Vrsn       float64            `xml:"urn:iso:std:iso:20022:tech:xsd:caam.007.001.01 Vrsn,omitempty"`
	Rcpt       []Recipient4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:caam.007.001.01 Rcpt"`
	NcrptdCntt EncryptedContent3  `xml:"urn:iso:std:iso:20022:tech:xsd:caam.007.001.01 NcrptdCntt,omitempty"`
}

// GenericIdentification77
//
// Identification of an entity.
type GenericIdentification77 struct {
	Id     Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:caam.007.001.01 Id"`
	Tp     PartyType12Code   `xml:"urn:iso:std:iso:20022:tech:xsd:caam.007.001.01 Tp"`
	Issr   PartyType12Code   `xml:"urn:iso:std:iso:20022:tech:xsd:caam.007.001.01 Issr,omitempty"`
	Ctry   Min2Max3AlphaText `xml:"urn:iso:std:iso:20022:tech:xsd:caam.007.001.01 Ctry,omitempty"`
	ShrtNm Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:caam.007.001.01 ShrtNm,omitempty"`
}

// Header20
//
// Information related to the protocol management on a segment of the path from the ATM to the acquirer.
type Header20 struct {
	MsgFctn    ATMMessageFunction1 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.007.001.01 MsgFctn"`
	PrtcolVrsn Max6Text            `xml:"urn:iso:std:iso:20022:tech:xsd:caam.007.001.01 PrtcolVrsn"`
	XchgId     Max3NumericText     `xml:"urn:iso:std:iso:20022:tech:xsd:caam.007.001.01 XchgId"`
	CreDtTm    ISODateTime         `xml:"urn:iso:std:iso:20022:tech:xsd:caam.007.001.01 CreDtTm"`
	InitgPty   Max35Text           `xml:"urn:iso:std:iso:20022:tech:xsd:caam.007.001.01 InitgPty"`
	RcptPty    Max35Text           `xml:"urn:iso:std:iso:20022:tech:xsd:caam.007.001.01 RcptPty,omitempty"`
	PrcStat    Max35Text           `xml:"urn:iso:std:iso:20022:tech:xsd:caam.007.001.01 PrcStat,omitempty"`
	Tracblt    []Traceability4     `xml:"urn:iso:std:iso:20022:tech:xsd:caam.007.001.01 Tracblt,omitempty"`
}

// HostToATMRequest1
//
// Information related to the request to an ATM to contact the ATM manager.
type HostToATMRequest1 struct {
	Envt         ATMEnvironment9           `xml:"urn:iso:std:iso:20022:tech:xsd:caam.007.001.01 Envt"`
	CmdId        ATMCommandIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.007.001.01 CmdId,omitempty"`
	XpctdMsgFctn MessageFunction8Code      `xml:"urn:iso:std:iso:20022:tech:xsd:caam.007.001.01 XpctdMsgFctn"`
}

// HostToATMRequestV01
//
// The HostToATMRequest message is sent by a host to an ATM to request the ATM to contact a host by sending of a maintenance messages.
type HostToATMRequestV01 struct {
	Hdr               Header20                 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.007.001.01 Hdr"`
	PrtctdHstToATMReq ContentInformationType10 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.007.001.01 PrtctdHstToATMReq,omitempty"`
	HstToATMReq       HostToATMRequest1        `xml:"urn:iso:std:iso:20022:tech:xsd:caam.007.001.01 HstToATMReq,omitempty"`
	SctyTrlr          ContentInformationType15 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.007.001.01 SctyTrlr,omitempty"`
}

// ISODateTime
//
// A particular point in the progression of time defined by a mandatory date and a mandatory time component, expressed in either UTC time format (YYYY-MM-DDThh:mm:ss.sssZ), local time with UTC offset format (YYYY-MM-DDThh:mm:ss.sss+/-hh:mm), or local time format (YYYY-MM-DDThh:mm:ss.sss). These representations are defined in "XML Schema Part 2: Datatypes Second Edition - W3C Recommendation 28 October 2004" which is aligned with ISO 8601.
// Note on the time format:
// 1) beginning / end of calendar day
// 00:00:00 = the beginning of a calendar day
// 24:00:00 = the end of a calendar day
// 2) fractions of second in time format
// Decimal fractions of seconds may be included. In this case, the involved parties shall agree on the maximum number of digits that are allowed.
type ISODateTime time.Time

func (t *ISODateTime) UnmarshalText(text []byte) error {
	return (*xsdDateTime)(t).UnmarshalText(text)
}
func (t ISODateTime) MarshalText() ([]byte, error) {
	return xsdDateTime(t).MarshalText()
}

// IssuerAndSerialNumber1
//
// Certificate issuer name and serial number  (see X.509).
type IssuerAndSerialNumber1 struct {
	Issr  CertificateIssuer1 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.007.001.01 Issr"`
	SrlNb Max35Binary        `xml:"urn:iso:std:iso:20022:tech:xsd:caam.007.001.01 SrlNb"`
}

// KEK4
//
// Key encryption key (KEK), using previously distributed symmetric key.
type KEK4 struct {
	Vrsn          float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:caam.007.001.01 Vrsn,omitempty"`
	KEKId         KEKIdentifier2            `xml:"urn:iso:std:iso:20022:tech:xsd:caam.007.001.01 KEKId"`
	KeyNcrptnAlgo AlgorithmIdentification13 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.007.001.01 KeyNcrptnAlgo"`
	NcrptdKey     Max500Binary              `xml:"urn:iso:std:iso:20022:tech:xsd:caam.007.001.01 NcrptdKey"`
}

// KEKIdentifier2
//
// Identification of a key encryption key (KEK), using previously distributed symmetric key.
type KEKIdentifier2 struct {
	KeyId     Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:caam.007.001.01 KeyId"`
	KeyVrsn   Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:caam.007.001.01 KeyVrsn"`
	SeqNb     float64         `xml:"urn:iso:std:iso:20022:tech:xsd:caam.007.001.01 SeqNb,omitempty"`
	DerivtnId Min5Max16Binary `xml:"urn:iso:std:iso:20022:tech:xsd:caam.007.001.01 DerivtnId,omitempty"`
}

// KeyTransport4
//
// Key encryption key (KEK), encrypted with a previously distributed asymmetric public key.
type KeyTransport4 struct {
	Vrsn          float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:caam.007.001.01 Vrsn,omitempty"`
	RcptId        Recipient5Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:caam.007.001.01 RcptId"`
	KeyNcrptnAlgo AlgorithmIdentification11 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.007.001.01 KeyNcrptnAlgo"`
	NcrptdKey     Max5000Binary             `xml:"urn:iso:std:iso:20022:tech:xsd:caam.007.001.01 NcrptdKey"`
}

// Max100KBinary
//
// Binary data of 100K maximum.
type Max100KBinary []byte

func (t *Max100KBinary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Max100KBinary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

// Max140Binary
//
// Specifies a binary string with a maximum length of 140 binary bytes.
type Max140Binary []byte

func (t *Max140Binary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Max140Binary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

// Must be at least 1 items long
type Max140Text string

// Max35Binary
//
// Specifies a binary string with a maximum length of 35 binary bytes.
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

// Max5000Binary
//
// Specifies a binary string with a maximum length of 5000 binary bytes.
type Max5000Binary []byte

func (t *Max5000Binary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Max5000Binary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

// Max500Binary
//
// Specifies a binary string with a maximum length of 500 binary bytes.
type Max500Binary []byte

func (t *Max500Binary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Max500Binary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

// Must be at least 1 items long
type Max6Text string

// May be one of BALN, CMPA, CMPD, ACMD, DVCC, DIAQ, DIAP, GSTS, INQQ, INQP, KYAQ, KYAP, PINQ, PINP, RJAQ, RJAP, WITV, WITK, WITQ, WITP, INQC, H2AP, H2AQ, TMOP, CSEC, DSEC, SKSC, SSTS
type MessageFunction7Code string

// May be one of BALN, GSTS, DSEC, INQC, KEYQ, SSTS
type MessageFunction8Code string

// Must match the pattern [a-zA-Z]{2,3}
type Min2Max3AlphaText string

// Min5Max16Binary
//
// Specifies a binary string with a minimum length of 5 bytes, and a maximum length of 16 bytes.
type Min5Max16Binary []byte

func (t *Min5Max16Binary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Min5Max16Binary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

// Parameter4
//
// Parameters of the asymmetric encryption algorithm.
type Parameter4 struct {
	NcrptnFrmt   EncryptionFormat1Code     `xml:"urn:iso:std:iso:20022:tech:xsd:caam.007.001.01 NcrptnFrmt,omitempty"`
	DgstAlgo     Algorithm11Code           `xml:"urn:iso:std:iso:20022:tech:xsd:caam.007.001.01 DgstAlgo,omitempty"`
	MskGnrtrAlgo AlgorithmIdentification12 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.007.001.01 MskGnrtrAlgo,omitempty"`
}

// Parameter5
//
// Parameters associated to a mask generator cryptographic function.
type Parameter5 struct {
	DgstAlgo Algorithm11Code `xml:"urn:iso:std:iso:20022:tech:xsd:caam.007.001.01 DgstAlgo,omitempty"`
}

// Parameter6
//
// Parameters associated to a cryptographic encryption algorithm.
type Parameter6 struct {
	NcrptnFrmt   EncryptionFormat1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caam.007.001.01 NcrptnFrmt,omitempty"`
	InitlstnVctr Max500Binary          `xml:"urn:iso:std:iso:20022:tech:xsd:caam.007.001.01 InitlstnVctr,omitempty"`
	BPddg        BytePadding1Code      `xml:"urn:iso:std:iso:20022:tech:xsd:caam.007.001.01 BPddg,omitempty"`
}

// Parameter7
//
// Parameters associated to the MAC algorithm.
type Parameter7 struct {
	InitlstnVctr Max500Binary     `xml:"urn:iso:std:iso:20022:tech:xsd:caam.007.001.01 InitlstnVctr,omitempty"`
	BPddg        BytePadding1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caam.007.001.01 BPddg,omitempty"`
}

// May be one of ACQR, ATMG, CISP, DLIS, HSTG, ITAG, OATM
type PartyType12Code string

// Recipient4Choice
//
// Transport key or key encryption key (KEK) for the recipient.
type Recipient4Choice struct {
	KeyTrnsprt KeyTransport4  `xml:"urn:iso:std:iso:20022:tech:xsd:caam.007.001.01 KeyTrnsprt"`
	KEK        KEK4           `xml:"urn:iso:std:iso:20022:tech:xsd:caam.007.001.01 KEK"`
	KeyIdr     KEKIdentifier2 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.007.001.01 KeyIdr"`
}

// Recipient5Choice
//
// Identification of a cryptographic asymmetric key.
type Recipient5Choice struct {
	IssrAndSrlNb IssuerAndSerialNumber1 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.007.001.01 IssrAndSrlNb"`
	KeyIdr       KEKIdentifier2         `xml:"urn:iso:std:iso:20022:tech:xsd:caam.007.001.01 KeyIdr"`
}

// RelativeDistinguishedName1
//
// Relative distinguished name defined by X.500 and X.509.
type RelativeDistinguishedName1 struct {
	AttrTp  AttributeType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caam.007.001.01 AttrTp"`
	AttrVal Max140Text         `xml:"urn:iso:std:iso:20022:tech:xsd:caam.007.001.01 AttrVal"`
}

// Traceability4
//
// Identification of partners involved in exchange from the ATM to the issuer, with the relative timestamp of their exchanges.
type Traceability4 struct {
	RlayId      GenericIdentification77 `xml:"urn:iso:std:iso:20022:tech:xsd:caam.007.001.01 RlayId"`
	SeqNb       Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:caam.007.001.01 SeqNb,omitempty"`
	TracDtTmIn  ISODateTime             `xml:"urn:iso:std:iso:20022:tech:xsd:caam.007.001.01 TracDtTmIn"`
	TracDtTmOut ISODateTime             `xml:"urn:iso:std:iso:20022:tech:xsd:caam.007.001.01 TracDtTmOut"`
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
