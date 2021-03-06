// Code generated by main. DO NOT EDIT.

package caaa_013_001_07

import (
	"bytes"
	"encoding/base64"
	"encoding/xml"
	"time"
)

type AcceptorDiagnosticRequest7 struct {
	Envt CardPaymentEnvironment71 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 Envt"`
}

type AcceptorDiagnosticRequestV07 struct {
	Hdr       Header35                   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 Hdr"`
	DgnstcReq AcceptorDiagnosticRequest7 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 DgnstcReq"`
	SctyTrlr  ContentInformationType16   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 SctyTrlr,omitempty"`
}

type Acquirer4 struct {
	Id         GenericIdentification53 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 Id,omitempty"`
	ParamsVrsn Max256Text              `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 ParamsVrsn"`
}

// May be one of HS25, HS38, HS51, HS01
type Algorithm11Code string

// May be one of HS25, HS38, HS51, HS01, SH31, SH32, SH33, SH35, SHK1, SHK2
type Algorithm16Code string

// May be one of MACC, MCCS, CMA1, MCC1, CMA9, CMA5, CMA2, CM31, CM32, CM33, MCS3, CCA1, CCA2, CCA3
type Algorithm17Code string

// May be one of EA2C, E3DC, DKP9, UKPT, UKA1, EA9C, EA5C, DA12, DA19, DA25, N108, EA5R, EA9R, EA2R, E3DR, E36C, E36R, SD5C
type Algorithm18Code string

// May be one of ERS2, ERS1, RPSS, ECC5, ECC1, ECC4, ECC2, ECC3, ERS3, ECP2, ECP3, ECP5
type Algorithm19Code string

// May be one of ERSA, RSAO
type Algorithm7Code string

// May be one of MGF1
type Algorithm8Code string

type AlgorithmIdentification12 struct {
	Algo  Algorithm8Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 Algo"`
	Param Parameter5     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 Param,omitempty"`
}

type AlgorithmIdentification18 struct {
	Algo  Algorithm8Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 Algo"`
	Param Parameter9     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 Param,omitempty"`
}

type AlgorithmIdentification19 struct {
	Algo  Algorithm7Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 Algo"`
	Param Parameter10    `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 Param,omitempty"`
}

type AlgorithmIdentification20 struct {
	Algo  Algorithm19Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 Algo"`
	Param Parameter11     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 Param,omitempty"`
}

type AlgorithmIdentification21 struct {
	Algo Algorithm16Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 Algo"`
}

type AlgorithmIdentification22 struct {
	Algo  Algorithm17Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 Algo"`
	Param Parameter7      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 Param,omitempty"`
}

type AlgorithmIdentification23 struct {
	Algo  Algorithm18Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 Algo"`
	Param Parameter12     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 Param,omitempty"`
}

type AlgorithmIdentification24 struct {
	Algo  Algorithm18Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 Algo"`
	Param Parameter12     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 Param,omitempty"`
}

// May be one of CNAT, LATT, OATT, OUAT, CATT
type AttributeType1Code string

type AuthenticatedData5 struct {
	Vrsn        float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 Vrsn,omitempty"`
	Rcpt        []Recipient6Choice        `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 Rcpt"`
	MACAlgo     AlgorithmIdentification22 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 MACAlgo"`
	NcpsltdCntt EncapsulatedContent3      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 NcpsltdCntt"`
	MAC         Max140Binary              `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 MAC"`
}

// May be one of LNGT, NUL8, NULG, NULL, RAND
type BytePadding1Code string

type CardPaymentEnvironment71 struct {
	Acqrr           Acquirer4                      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 Acqrr"`
	AcqrrAvlbtyReqd bool                           `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 AcqrrAvlbtyReqd,omitempty"`
	MrchntId        GenericIdentification53        `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 MrchntId,omitempty"`
	POIId           GenericIdentification32        `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 POIId,omitempty"`
	POICmpnt        []PointOfInteractionComponent8 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 POICmpnt,omitempty"`
	AcqrrAvlbl      bool                           `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 AcqrrAvlbl,omitempty"`
}

type CertificateIssuer1 struct {
	RltvDstngshdNm []RelativeDistinguishedName1 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 RltvDstngshdNm"`
}

type CommunicationCharacteristics3 struct {
	ComTp   POICommunicationType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 ComTp"`
	RmotPty []PartyType7Code          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 RmotPty"`
	Actv    bool                      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 Actv"`
}

type ContentInformationType16 struct {
	CnttTp       ContentType2Code   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 CnttTp"`
	AuthntcdData AuthenticatedData5 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 AuthntcdData"`
}

type ContentInformationType19 struct {
	CnttTp       ContentType2Code   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 CnttTp"`
	EnvlpdData   EnvelopedData5     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 EnvlpdData,omitempty"`
	AuthntcdData AuthenticatedData5 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 AuthntcdData,omitempty"`
	SgndData     SignedData5        `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 SgndData,omitempty"`
	DgstdData    DigestedData5      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 DgstdData,omitempty"`
}

// May be one of DATA, SIGN, EVLP, DGST, AUTH
type ContentType2Code string

type CryptographicKey13 struct {
	Id           Max140Text                `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 Id"`
	AddtlId      Max35Binary               `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 AddtlId,omitempty"`
	Vrsn         Max256Text                `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 Vrsn"`
	Tp           CryptographicKeyType3Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 Tp,omitempty"`
	Fctn         []KeyUsage1Code           `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 Fctn,omitempty"`
	ActvtnDt     ISODateTime               `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 ActvtnDt,omitempty"`
	DeactvtnDt   ISODateTime               `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 DeactvtnDt,omitempty"`
	KeyVal       ContentInformationType19  `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 KeyVal,omitempty"`
	KeyChckVal   Max35Binary               `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 KeyChckVal,omitempty"`
	AddtlMgmtInf []GenericInformation1     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 AddtlMgmtInf,omitempty"`
}

// May be one of AES2, EDE3, DKP9, AES9, AES5, EDE4
type CryptographicKeyType3Code string

type DigestedData5 struct {
	Vrsn        float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 Vrsn,omitempty"`
	DgstAlgo    AlgorithmIdentification21 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 DgstAlgo"`
	NcpsltdCntt EncapsulatedContent3      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 NcpsltdCntt"`
	Dgst        Max140Binary              `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 Dgst"`
}

type Document struct {
	AccptrDgnstcReq AcceptorDiagnosticRequestV07 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 AccptrDgnstcReq"`
}

type EncapsulatedContent3 struct {
	CnttTp ContentType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 CnttTp"`
	Cntt   Max100KBinary    `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 Cntt,omitempty"`
}

type EncryptedContent4 struct {
	CnttTp         ContentType2Code          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 CnttTp"`
	CnttNcrptnAlgo AlgorithmIdentification24 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 CnttNcrptnAlgo,omitempty"`
	NcrptdData     Max100KBinary             `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 NcrptdData"`
}

// May be one of TR31, TR34, I238
type EncryptionFormat2Code string

type EnvelopedData5 struct {
	Vrsn       float64                `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 Vrsn,omitempty"`
	OrgtrInf   OriginatorInformation1 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 OrgtrInf,omitempty"`
	Rcpt       []Recipient6Choice     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 Rcpt"`
	NcrptdCntt EncryptedContent4      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 NcrptdCntt,omitempty"`
}

type GenericIdentification32 struct {
	Id     Max35Text      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 Id"`
	Tp     PartyType3Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 Tp,omitempty"`
	Issr   PartyType4Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 Issr,omitempty"`
	ShrtNm Max35Text      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 ShrtNm,omitempty"`
}

type GenericIdentification48 struct {
	Id   Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 Id"`
	Vrsn Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 Vrsn"`
	Issr Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 Issr"`
}

type GenericIdentification53 struct {
	Id     Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 Id"`
	Tp     PartyType3Code    `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 Tp,omitempty"`
	Issr   PartyType4Code    `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 Issr,omitempty"`
	Ctry   Min2Max3AlphaText `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 Ctry,omitempty"`
	ShrtNm Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 ShrtNm,omitempty"`
}

type GenericIdentification76 struct {
	Id     Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 Id"`
	Tp     PartyType3Code    `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 Tp"`
	Issr   PartyType4Code    `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 Issr,omitempty"`
	Ctry   Min2Max3AlphaText `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 Ctry,omitempty"`
	ShrtNm Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 ShrtNm,omitempty"`
}

type GenericIdentification94 struct {
	Id       Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 Id"`
	Tp       PartyType3Code     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 Tp,omitempty"`
	Issr     PartyType4Code     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 Issr,omitempty"`
	Ctry     Min2Max3AlphaText  `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 Ctry,omitempty"`
	ShrtNm   Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 ShrtNm,omitempty"`
	RmotAccs NetworkParameters5 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 RmotAccs,omitempty"`
}

type GenericInformation1 struct {
	Nm  Max70Text  `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 Nm"`
	Val Max140Text `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 Val,omitempty"`
}

type Header35 struct {
	MsgFctn    MessageFunction14Code   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 MsgFctn"`
	PrtcolVrsn Max6Text                `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 PrtcolVrsn"`
	XchgId     float64                 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 XchgId"`
	CreDtTm    ISODateTime             `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 CreDtTm"`
	InitgPty   GenericIdentification53 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 InitgPty"`
	RcptPty    GenericIdentification94 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 RcptPty,omitempty"`
	Tracblt    []Traceability5         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 Tracblt,omitempty"`
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

type IssuerAndSerialNumber1 struct {
	Issr  CertificateIssuer1 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 Issr"`
	SrlNb Max35Binary        `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 SrlNb"`
}

type KEK5 struct {
	Vrsn          float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 Vrsn,omitempty"`
	KEKId         KEKIdentifier2            `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 KEKId"`
	KeyNcrptnAlgo AlgorithmIdentification23 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 KeyNcrptnAlgo"`
	NcrptdKey     Max500Binary              `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 NcrptdKey"`
}

type KEKIdentifier2 struct {
	KeyId     Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 KeyId"`
	KeyVrsn   Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 KeyVrsn"`
	SeqNb     float64         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 SeqNb,omitempty"`
	DerivtnId Min5Max16Binary `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 DerivtnId,omitempty"`
}

type KeyTransport5 struct {
	Vrsn          float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 Vrsn,omitempty"`
	RcptId        Recipient5Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 RcptId"`
	KeyNcrptnAlgo AlgorithmIdentification19 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 KeyNcrptnAlgo"`
	NcrptdKey     Max5000Binary             `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 NcrptdKey"`
}

// May be one of ENCR, DCPT, DENC, DDEC, TRNI, TRNX, MACG, MACV, SIGG, SUGV, PINE, PIND, PINV, KEYG, KEYI, KEYX, KEYD
type KeyUsage1Code string

type Max100KBinary []byte

func (t *Max100KBinary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Max100KBinary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

type Max10KBinary []byte

func (t *Max10KBinary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Max10KBinary) MarshalText() ([]byte, error) {
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

// Must be at least 1 items long
type Max256Text string

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

type Max5000Binary []byte

func (t *Max5000Binary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Max5000Binary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

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

type MemoryCharacteristics1 struct {
	Id     Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 Id"`
	TtlSz  float64         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 TtlSz"`
	FreeSz float64         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 FreeSz"`
	Unit   MemoryUnit1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 Unit"`
}

// May be one of BYTE, EXAB, GIGA, KILO, MEGA, PETA, TERA
type MemoryUnit1Code string

// May be one of AUTQ, AUTP, CCAV, CCAK, CCAQ, CCAP, CMPV, CMPK, DCAV, DCRR, DCCQ, DCCP, DGNP, DGNQ, FAUQ, FAUP, FCMV, FCMK, FRVA, FRVR, RCLQ, RCLP, RVRA, RVRR, CDDQ, CDDK, CDDR, CDDP
type MessageFunction14Code string

// Must match the pattern [a-zA-Z]{2,3}
type Min2Max3AlphaText string

type Min5Max16Binary []byte

func (t *Min5Max16Binary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Min5Max16Binary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

type NetworkParameters4 struct {
	NtwkTp NetworkType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 NtwkTp"`
	AdrVal Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 AdrVal"`
}

type NetworkParameters5 struct {
	Adr        []NetworkParameters4 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 Adr"`
	UsrNm      Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 UsrNm,omitempty"`
	AccsCd     Max35Binary          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 AccsCd,omitempty"`
	SvrCert    []Max10KBinary       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 SvrCert,omitempty"`
	SvrCertIdr []Max140Binary       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 SvrCertIdr,omitempty"`
	ClntCert   []Max10KBinary       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 ClntCert,omitempty"`
	SctyPrfl   Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 SctyPrfl,omitempty"`
}

// May be one of IPNW, PSTN
type NetworkType1Code string

type OriginatorInformation1 struct {
	Cert []Max5000Binary `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 Cert,omitempty"`
}

// May be one of BLTH, ETHR, GPRS, GSMF, PSTN, RS23, USBD, USBH, WIFI, WT2G, WT3G, WT4G, WT5G
type POICommunicationType2Code string

// May be one of APPL, CERT, EVAL
type POIComponentAssessment1Code string

// May be one of WAIT, OUTD, OPER, DACT
type POIComponentStatus1Code string

// May be one of AQPP, APPR, TLPR, SCPR, SERV, TERM, DVCE, SECM, APLI, EMVK, EMVO, MDWR, DRVR, OPST, MRPR, CRTF, TMSP, SACP, SAPR
type POIComponentType5Code string

type Parameter10 struct {
	NcrptnFrmt   EncryptionFormat2Code     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 NcrptnFrmt,omitempty"`
	DgstAlgo     Algorithm16Code           `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 DgstAlgo,omitempty"`
	MskGnrtrAlgo AlgorithmIdentification18 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 MskGnrtrAlgo,omitempty"`
}

type Parameter11 struct {
	DgstAlgo     Algorithm16Code           `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 DgstAlgo"`
	MskGnrtrAlgo AlgorithmIdentification12 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 MskGnrtrAlgo"`
	SaltLngth    float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 SaltLngth"`
	TrlrFld      float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 TrlrFld,omitempty"`
}

type Parameter12 struct {
	NcrptnFrmt   EncryptionFormat2Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 NcrptnFrmt,omitempty"`
	InitlstnVctr Max500Binary          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 InitlstnVctr,omitempty"`
	BPddg        BytePadding1Code      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 BPddg,omitempty"`
}

type Parameter5 struct {
	DgstAlgo Algorithm11Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 DgstAlgo,omitempty"`
}

type Parameter7 struct {
	InitlstnVctr Max500Binary     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 InitlstnVctr,omitempty"`
	BPddg        BytePadding1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 BPddg,omitempty"`
}

type Parameter9 struct {
	DgstAlgo Algorithm16Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 DgstAlgo,omitempty"`
}

// May be one of OPOI, MERC, ACCP, ITAG, ACQR, CISS, DLIS
type PartyType3Code string

// May be one of MERC, ACCP, ITAG, ACQR, CISS, TAXH
type PartyType4Code string

// May be one of ACQR, ITAG, PCPT, TMGT, SALE
type PartyType7Code string

type PointOfInteractionComponent8 struct {
	Tp       POIComponentType5Code                       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 Tp"`
	Id       PointOfInteractionComponentIdentification1  `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 Id"`
	Sts      PointOfInteractionComponentStatus3          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 Sts,omitempty"`
	StdCmplc []GenericIdentification48                   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 StdCmplc,omitempty"`
	Chrtcs   PointOfInteractionComponentCharacteristics4 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 Chrtcs,omitempty"`
	Assmnt   []PointOfInteractionComponentAssessment1    `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 Assmnt,omitempty"`
}

type PointOfInteractionComponentAssessment1 struct {
	Tp      POIComponentAssessment1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 Tp"`
	Assgnr  []Max35Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 Assgnr"`
	DlvryDt ISODateTime                 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 DlvryDt,omitempty"`
	XprtnDt ISODateTime                 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 XprtnDt,omitempty"`
	Nb      Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 Nb"`
}

type PointOfInteractionComponentCharacteristics4 struct {
	Mmry           []MemoryCharacteristics1        `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 Mmry,omitempty"`
	Com            []CommunicationCharacteristics3 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 Com,omitempty"`
	SctyAccsMdls   float64                         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 SctyAccsMdls,omitempty"`
	SbcbrIdntyMdls float64                         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 SbcbrIdntyMdls,omitempty"`
	SctyElmt       []CryptographicKey13            `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 SctyElmt,omitempty"`
}

type PointOfInteractionComponentIdentification1 struct {
	ItmNb   Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 ItmNb,omitempty"`
	PrvdrId Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 PrvdrId,omitempty"`
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 Id,omitempty"`
	SrlNb   Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 SrlNb,omitempty"`
}

type PointOfInteractionComponentStatus3 struct {
	VrsnNb Max256Text              `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 VrsnNb,omitempty"`
	Sts    POIComponentStatus1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 Sts,omitempty"`
	XpryDt ISODate                 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 XpryDt,omitempty"`
}

type Recipient5Choice struct {
	IssrAndSrlNb IssuerAndSerialNumber1 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 IssrAndSrlNb"`
	KeyIdr       KEKIdentifier2         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 KeyIdr"`
}

type Recipient6Choice struct {
	KeyTrnsprt KeyTransport5  `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 KeyTrnsprt"`
	KEK        KEK5           `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 KEK"`
	KeyIdr     KEKIdentifier2 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 KeyIdr"`
}

type RelativeDistinguishedName1 struct {
	AttrTp  AttributeType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 AttrTp"`
	AttrVal Max140Text         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 AttrVal"`
}

type SignedData5 struct {
	Vrsn        float64                     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 Vrsn,omitempty"`
	DgstAlgo    []AlgorithmIdentification21 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 DgstAlgo,omitempty"`
	NcpsltdCntt EncapsulatedContent3        `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 NcpsltdCntt,omitempty"`
	Cert        []Max5000Binary             `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 Cert,omitempty"`
	Sgnr        []Signer4                   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 Sgnr,omitempty"`
}

type Signer4 struct {
	Vrsn        float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 Vrsn,omitempty"`
	SgnrId      Recipient5Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 SgnrId,omitempty"`
	DgstAlgo    AlgorithmIdentification21 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 DgstAlgo"`
	SgndAttrbts []GenericInformation1     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 SgndAttrbts,omitempty"`
	SgntrAlgo   AlgorithmIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 SgntrAlgo"`
	Sgntr       Max3000Binary             `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 Sgntr"`
}

type Traceability5 struct {
	RlayId      GenericIdentification76 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 RlayId"`
	PrtcolNm    Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 PrtcolNm,omitempty"`
	PrtcolVrsn  Max6Text                `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 PrtcolVrsn,omitempty"`
	TracDtTmIn  ISODateTime             `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 TracDtTmIn"`
	TracDtTmOut ISODateTime             `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.013.001.07 TracDtTmOut"`
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
