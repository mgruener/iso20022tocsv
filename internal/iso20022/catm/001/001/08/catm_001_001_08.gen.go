// Code generated by main. DO NOT EDIT.

package catm_001_001_08

import (
	"bytes"
	"encoding/base64"
	"encoding/xml"
	"time"
)

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
	Algo  Algorithm8Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 Algo"`
	Param Parameter5     `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 Param,omitempty"`
}

type AlgorithmIdentification18 struct {
	Algo  Algorithm8Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 Algo"`
	Param Parameter9     `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 Param,omitempty"`
}

type AlgorithmIdentification19 struct {
	Algo  Algorithm7Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 Algo"`
	Param Parameter10    `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 Param,omitempty"`
}

type AlgorithmIdentification20 struct {
	Algo  Algorithm19Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 Algo"`
	Param Parameter11     `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 Param,omitempty"`
}

type AlgorithmIdentification21 struct {
	Algo Algorithm16Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 Algo"`
}

type AlgorithmIdentification22 struct {
	Algo  Algorithm17Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 Algo"`
	Param Parameter7      `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 Param,omitempty"`
}

type AlgorithmIdentification23 struct {
	Algo  Algorithm18Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 Algo"`
	Param Parameter12     `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 Param,omitempty"`
}

type AlgorithmIdentification24 struct {
	Algo  Algorithm18Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 Algo"`
	Param Parameter12     `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 Param,omitempty"`
}

// May be one of ATTD, SATT, UATT
type AttendanceContext1Code string

// May be one of CNAT, LATT, OATT, OUAT, CATT
type AttributeType1Code string

type AuthenticatedData5 struct {
	Vrsn        float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 Vrsn,omitempty"`
	Rcpt        []Recipient6Choice        `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 Rcpt"`
	MACAlgo     AlgorithmIdentification22 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 MACAlgo"`
	NcpsltdCntt EncapsulatedContent3      `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 NcpsltdCntt"`
	MAC         Max140Binary              `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 MAC"`
}

// May be one of LNGT, NUL8, NULG, NULL, RAND
type BytePadding1Code string

// May be one of TAGC, PHYS, BRCD, MGST, CICC, DFLE, CTLS, ECTL, CDFL, SICC, UNKW
type CardDataReading6Code string

// May be one of APKI, CHDT, MNSG, MNVR, FBIG, FBIO, FDSG, FCPN, FEPN, NPIN, PKIS, SCEC, NBIO, NOVF, OTHR
type CardholderVerificationCapability4Code string

type CertificateIssuer1 struct {
	RltvDstngshdNm []RelativeDistinguishedName1 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 RltvDstngshdNm"`
}

type CommunicationCharacteristics4 struct {
	ComTp      POICommunicationType2Code   `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 ComTp"`
	RmotPty    []PartyType7Code            `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 RmotPty"`
	Actv       bool                        `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 Actv"`
	Params     NetworkParameters5          `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 Params,omitempty"`
	PhysIntrfc PhysicalInterfaceParameter1 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 PhysIntrfc,omitempty"`
}

type ContentInformationType18 struct {
	CnttTp       ContentType2Code   `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 CnttTp"`
	AuthntcdData AuthenticatedData5 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 AuthntcdData,omitempty"`
	SgndData     SignedData5        `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 SgndData,omitempty"`
}

type ContentInformationType19 struct {
	CnttTp       ContentType2Code   `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 CnttTp"`
	EnvlpdData   EnvelopedData5     `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 EnvlpdData,omitempty"`
	AuthntcdData AuthenticatedData5 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 AuthntcdData,omitempty"`
	SgndData     SignedData5        `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 SgndData,omitempty"`
	DgstdData    DigestedData5      `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 DgstdData,omitempty"`
}

// May be one of DATA, SIGN, EVLP, DGST, AUTH
type ContentType2Code string

type CryptographicKey13 struct {
	Id           Max140Text                `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 Id"`
	AddtlId      Max35Binary               `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 AddtlId,omitempty"`
	Vrsn         Max256Text                `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 Vrsn"`
	Tp           CryptographicKeyType3Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 Tp,omitempty"`
	Fctn         []KeyUsage1Code           `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 Fctn,omitempty"`
	ActvtnDt     ISODateTime               `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 ActvtnDt,omitempty"`
	DeactvtnDt   ISODateTime               `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 DeactvtnDt,omitempty"`
	KeyVal       ContentInformationType19  `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 KeyVal,omitempty"`
	KeyChckVal   Max35Binary               `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 KeyChckVal,omitempty"`
	AddtlMgmtInf []GenericInformation1     `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 AddtlMgmtInf,omitempty"`
}

// May be one of AES2, EDE3, DKP9, AES9, AES5, EDE4
type CryptographicKeyType3Code string

// May be one of AQPR, APPR, TXCP, AKCP, DLGT, MGTP, MRPR, SCPR, SWPK, STRP, TRPR, VDPR, PARA, TMSP, CRTF, LOGF, CMRQ
type DataSetCategory12Code string

type DataSetIdentification7 struct {
	Nm      Max256Text            `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 Nm,omitempty"`
	Tp      DataSetCategory12Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 Tp"`
	Vrsn    Max256Text            `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 Vrsn,omitempty"`
	CreDtTm ISODateTime           `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 CreDtTm,omitempty"`
}

type DigestedData5 struct {
	Vrsn        float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 Vrsn,omitempty"`
	DgstAlgo    AlgorithmIdentification21 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 DgstAlgo"`
	NcpsltdCntt EncapsulatedContent3      `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 NcpsltdCntt"`
	Dgst        Max140Binary              `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 Dgst"`
}

type DisplayCapabilities4 struct {
	Dstn      []UserInterface4Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 Dstn"`
	AvlblFrmt []OutputFormat1Code  `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 AvlblFrmt,omitempty"`
	NbOfLines float64              `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 NbOfLines,omitempty"`
	LineWidth float64              `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 LineWidth,omitempty"`
	AvlblLang []string             `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 AvlblLang,omitempty"`
}

type Document struct {
	StsRpt StatusReportV08 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 StsRpt"`
}

type EncapsulatedContent3 struct {
	CnttTp ContentType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 CnttTp"`
	Cntt   Max100KBinary    `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 Cntt,omitempty"`
}

type EncryptedContent4 struct {
	CnttTp         ContentType2Code          `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 CnttTp"`
	CnttNcrptnAlgo AlgorithmIdentification24 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 CnttNcrptnAlgo,omitempty"`
	NcrptdData     Max100KBinary             `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 NcrptdData"`
}

// May be one of TR31, TR34, I238
type EncryptionFormat2Code string

type EnvelopedData5 struct {
	Vrsn       float64                `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 Vrsn,omitempty"`
	OrgtrInf   OriginatorInformation1 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 OrgtrInf,omitempty"`
	Rcpt       []Recipient6Choice     `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 Rcpt"`
	NcrptdCntt EncryptedContent4      `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 NcrptdCntt,omitempty"`
}

// May be one of ONDM, IMMD, ASAP, AGRP, NBLT, TTLT, CYCL, NONE
type ExchangePolicy1Code string

type GenericIdentification48 struct {
	Id   Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 Id"`
	Vrsn Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 Vrsn"`
	Issr Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 Issr"`
}

type GenericIdentification71 struct {
	Id     Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 Id"`
	Tp     PartyType5Code    `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 Tp,omitempty"`
	Issr   PartyType6Code    `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 Issr,omitempty"`
	Ctry   Min2Max3AlphaText `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 Ctry,omitempty"`
	ShrtNm Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 ShrtNm,omitempty"`
}

type GenericIdentification92 struct {
	Id       Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 Id"`
	Tp       PartyType5Code     `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 Tp,omitempty"`
	Issr     PartyType6Code     `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 Issr,omitempty"`
	Ctry     Min2Max3AlphaText  `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 Ctry,omitempty"`
	ShrtNm   Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 ShrtNm,omitempty"`
	RmotAccs NetworkParameters5 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 RmotAccs,omitempty"`
}

type GenericInformation1 struct {
	Nm  Max70Text  `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 Nm"`
	Val Max140Text `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 Val,omitempty"`
}

type Header27 struct {
	DwnldTrf bool                    `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 DwnldTrf"`
	FrmtVrsn Max6Text                `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 FrmtVrsn"`
	XchgId   float64                 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 XchgId"`
	CreDtTm  ISODateTime             `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 CreDtTm"`
	InitgPty GenericIdentification71 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 InitgPty"`
	RcptPty  GenericIdentification92 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 RcptPty,omitempty"`
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
	Issr  CertificateIssuer1 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 Issr"`
	SrlNb Max35Binary        `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 SrlNb"`
}

type KEK5 struct {
	Vrsn          float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 Vrsn,omitempty"`
	KEKId         KEKIdentifier2            `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 KEKId"`
	KeyNcrptnAlgo AlgorithmIdentification23 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 KeyNcrptnAlgo"`
	NcrptdKey     Max500Binary              `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 NcrptdKey"`
}

type KEKIdentifier2 struct {
	KeyId     Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 KeyId"`
	KeyVrsn   Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 KeyVrsn"`
	SeqNb     float64         `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 SeqNb,omitempty"`
	DerivtnId Min5Max16Binary `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 DerivtnId,omitempty"`
}

type KeyTransport5 struct {
	Vrsn          float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 Vrsn,omitempty"`
	RcptId        Recipient5Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 RcptId"`
	KeyNcrptnAlgo AlgorithmIdentification19 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 KeyNcrptnAlgo"`
	NcrptdKey     Max5000Binary             `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 NcrptdKey"`
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

type Max2KBinary []byte

func (t *Max2KBinary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Max2KBinary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

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

// Must match the pattern [0-9]{1,9}
type Max9NumericText string

type MemoryCharacteristics1 struct {
	Id     Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 Id"`
	TtlSz  float64         `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 TtlSz"`
	FreeSz float64         `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 FreeSz"`
	Unit   MemoryUnit1Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 Unit"`
}

// May be one of BYTE, EXAB, GIGA, KILO, MEGA, PETA, TERA
type MemoryUnit1Code string

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
	NtwkTp NetworkType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 NtwkTp"`
	AdrVal Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 AdrVal"`
}

type NetworkParameters5 struct {
	Adr        []NetworkParameters4 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 Adr"`
	UsrNm      Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 UsrNm,omitempty"`
	AccsCd     Max35Binary          `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 AccsCd,omitempty"`
	SvrCert    []Max10KBinary       `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 SvrCert,omitempty"`
	SvrCertIdr []Max140Binary       `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 SvrCertIdr,omitempty"`
	ClntCert   []Max10KBinary       `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 ClntCert,omitempty"`
	SctyPrfl   Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 SctyPrfl,omitempty"`
}

// May be one of IPNW, PSTN
type NetworkType1Code string

// May be one of OFLN, ONLN, SMON
type OnLineCapability1Code string

type OriginatorInformation1 struct {
	Cert []Max5000Binary `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 Cert,omitempty"`
}

// May be one of MREF, TEXT, HTML
type OutputFormat1Code string

// May be one of BLTH, ETHR, GPRS, GSMF, PSTN, RS23, USBD, USBH, WIFI, WT2G, WT3G, WT4G, WT5G
type POICommunicationType2Code string

// May be one of APPL, CERT, EVAL
type POIComponentAssessment1Code string

// May be one of WAIT, OUTD, OPER, DACT
type POIComponentStatus1Code string

// May be one of AQPP, APPR, TLPR, SCPR, SERV, TERM, DVCE, SECM, APLI, EMVK, EMVO, MDWR, DRVR, OPST, MRPR, CRTF, TMSP, SACP, SAPR
type POIComponentType5Code string

type Parameter10 struct {
	NcrptnFrmt   EncryptionFormat2Code     `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 NcrptnFrmt,omitempty"`
	DgstAlgo     Algorithm16Code           `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 DgstAlgo,omitempty"`
	MskGnrtrAlgo AlgorithmIdentification18 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 MskGnrtrAlgo,omitempty"`
}

type Parameter11 struct {
	DgstAlgo     Algorithm16Code           `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 DgstAlgo"`
	MskGnrtrAlgo AlgorithmIdentification12 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 MskGnrtrAlgo"`
	SaltLngth    float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 SaltLngth"`
	TrlrFld      float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 TrlrFld,omitempty"`
}

type Parameter12 struct {
	NcrptnFrmt   EncryptionFormat2Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 NcrptnFrmt,omitempty"`
	InitlstnVctr Max500Binary          `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 InitlstnVctr,omitempty"`
	BPddg        BytePadding1Code      `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 BPddg,omitempty"`
}

type Parameter5 struct {
	DgstAlgo Algorithm11Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 DgstAlgo,omitempty"`
}

type Parameter7 struct {
	InitlstnVctr Max500Binary     `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 InitlstnVctr,omitempty"`
	BPddg        BytePadding1Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 BPddg,omitempty"`
}

type Parameter9 struct {
	DgstAlgo Algorithm16Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 DgstAlgo,omitempty"`
}

// May be one of OPOI, ACCP, MERC, ACQR, ITAG, MTMG, TMGT
type PartyType5Code string

// May be one of ACCP, MERC, ACQR, ITAG, MTMG, TMGT
type PartyType6Code string

// May be one of ACQR, ITAG, PCPT, TMGT, SALE
type PartyType7Code string

type PhysicalInterfaceParameter1 struct {
	IntrfcNm    Max35Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 IntrfcNm"`
	IntrfcTp    POICommunicationType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 IntrfcTp,omitempty"`
	UsrNm       Max35Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 UsrNm,omitempty"`
	AccsCd      Max35Binary               `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 AccsCd,omitempty"`
	SctyPrfl    Max35Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 SctyPrfl,omitempty"`
	AddtlParams Max2KBinary               `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 AddtlParams,omitempty"`
}

type PointOfInteractionCapabilities8 struct {
	CardRdngCpblties      []CardDataReading6Code                  `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 CardRdngCpblties,omitempty"`
	CrdhldrVrfctnCpblties []CardholderVerificationCapability4Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 CrdhldrVrfctnCpblties,omitempty"`
	PINLngthCpblties      float64                                 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 PINLngthCpblties,omitempty"`
	ApprvlCdLngth         float64                                 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 ApprvlCdLngth,omitempty"`
	MxScrptLngth          float64                                 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 MxScrptLngth,omitempty"`
	CardCaptrCpbl         bool                                    `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 CardCaptrCpbl,omitempty"`
	OnLineCpblties        OnLineCapability1Code                   `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 OnLineCpblties,omitempty"`
	MsgCpblties           []DisplayCapabilities4                  `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 MsgCpblties,omitempty"`
}

type PointOfInteractionComponent9 struct {
	Tp       POIComponentType5Code                       `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 Tp"`
	SubTpInf Max70Text                                   `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 SubTpInf,omitempty"`
	Id       PointOfInteractionComponentIdentification1  `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 Id"`
	Sts      PointOfInteractionComponentStatus3          `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 Sts,omitempty"`
	StdCmplc []GenericIdentification48                   `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 StdCmplc,omitempty"`
	Chrtcs   PointOfInteractionComponentCharacteristics5 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 Chrtcs,omitempty"`
	Assmnt   []PointOfInteractionComponentAssessment1    `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 Assmnt,omitempty"`
}

type PointOfInteractionComponentAssessment1 struct {
	Tp      POIComponentAssessment1Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 Tp"`
	Assgnr  []Max35Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 Assgnr"`
	DlvryDt ISODateTime                 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 DlvryDt,omitempty"`
	XprtnDt ISODateTime                 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 XprtnDt,omitempty"`
	Nb      Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 Nb"`
}

type PointOfInteractionComponentCharacteristics5 struct {
	Mmry           []MemoryCharacteristics1        `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 Mmry,omitempty"`
	Com            []CommunicationCharacteristics4 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 Com,omitempty"`
	SctyAccsMdls   float64                         `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 SctyAccsMdls,omitempty"`
	SbcbrIdntyMdls float64                         `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 SbcbrIdntyMdls,omitempty"`
	SctyElmt       []CryptographicKey13            `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 SctyElmt,omitempty"`
}

type PointOfInteractionComponentIdentification1 struct {
	ItmNb   Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 ItmNb,omitempty"`
	PrvdrId Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 PrvdrId,omitempty"`
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 Id,omitempty"`
	SrlNb   Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 SrlNb,omitempty"`
}

type PointOfInteractionComponentStatus3 struct {
	VrsnNb Max256Text              `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 VrsnNb,omitempty"`
	Sts    POIComponentStatus1Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 Sts,omitempty"`
	XpryDt ISODate                 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 XpryDt,omitempty"`
}

type Recipient5Choice struct {
	IssrAndSrlNb IssuerAndSerialNumber1 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 IssrAndSrlNb"`
	KeyIdr       KEKIdentifier2         `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 KeyIdr"`
}

type Recipient6Choice struct {
	KeyTrnsprt KeyTransport5  `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 KeyTrnsprt"`
	KEK        KEK5           `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 KEK"`
	KeyIdr     KEKIdentifier2 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 KeyIdr"`
}

type RelativeDistinguishedName1 struct {
	AttrTp  AttributeType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 AttrTp"`
	AttrVal Max140Text         `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 AttrVal"`
}

type SignedData5 struct {
	Vrsn        float64                     `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 Vrsn,omitempty"`
	DgstAlgo    []AlgorithmIdentification21 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 DgstAlgo,omitempty"`
	NcpsltdCntt EncapsulatedContent3        `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 NcpsltdCntt,omitempty"`
	Cert        []Max5000Binary             `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 Cert,omitempty"`
	Sgnr        []Signer4                   `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 Sgnr,omitempty"`
}

type Signer4 struct {
	Vrsn        float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 Vrsn,omitempty"`
	SgnrId      Recipient5Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 SgnrId,omitempty"`
	DgstAlgo    AlgorithmIdentification21 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 DgstAlgo"`
	SgndAttrbts []GenericInformation1     `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 SgndAttrbts,omitempty"`
	SgntrAlgo   AlgorithmIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 SgntrAlgo"`
	Sgntr       Max3000Binary             `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 Sgntr"`
}

type StatusReport8 struct {
	POIId       GenericIdentification71     `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 POIId"`
	InitgTrggr  TriggerInformation1         `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 InitgTrggr,omitempty"`
	TermnlMgrId GenericIdentification71     `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 TermnlMgrId"`
	DataSet     TerminalManagementDataSet28 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 DataSet"`
}

type StatusReportContent8 struct {
	POICpblties  PointOfInteractionCapabilities8 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 POICpblties,omitempty"`
	POICmpnt     []PointOfInteractionComponent9  `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 POICmpnt,omitempty"`
	AttndncCntxt AttendanceContext1Code          `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 AttndncCntxt,omitempty"`
	POIDtTm      ISODateTime                     `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 POIDtTm"`
	DataSetReqrd []TerminalManagementDataSet25   `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 DataSetReqrd,omitempty"`
	Evt          []TMSEvent6                     `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 Evt,omitempty"`
	Errs         []Max140Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 Errs,omitempty"`
}

type StatusReportV08 struct {
	Hdr      Header27                 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 Hdr"`
	StsRpt   StatusReport8            `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 StsRpt"`
	SctyTrlr ContentInformationType18 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 SctyTrlr,omitempty"`
}

type TMSActionIdentification5 struct {
	ActnTp    TerminalManagementAction4Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 ActnTp"`
	DataSetId DataSetIdentification7        `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 DataSetId,omitempty"`
}

type TMSEvent6 struct {
	TmStmp      ISODateTime                         `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 TmStmp"`
	Rslt        TerminalManagementActionResult4Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 Rslt"`
	ActnId      TMSActionIdentification5            `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 ActnId"`
	AddtlErrInf Max70Text                           `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 AddtlErrInf,omitempty"`
	TermnlMgrId Max35Text                           `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 TermnlMgrId,omitempty"`
}

// May be one of DCTV, DELT, DWNL, INST, RSTR, UPLD, UPDT, BIND, RBND, UBND, ACTV
type TerminalManagementAction4Code string

// May be one of ACCD, CNTE, FMTE, INVC, LENE, OVER, MISS, NSUP, SIGE, SUCC, SYNE, TIMO, UKDT, UKRF, INDP, IDMP, DPRU, AERR, CMER, ULER
type TerminalManagementActionResult4Code string

type TerminalManagementDataSet25 struct {
	Id               DataSetIdentification7   `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 Id"`
	POIChllng        Max140Binary             `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 POIChllng,omitempty"`
	TMChllng         Max140Binary             `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 TMChllng,omitempty"`
	SsnKey           CryptographicKey13       `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 SsnKey,omitempty"`
	DlgtnProof       Max5000Binary            `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 DlgtnProof,omitempty"`
	PrtctdDlgtnProof ContentInformationType19 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 PrtctdDlgtnProof,omitempty"`
}

type TerminalManagementDataSet28 struct {
	Id      DataSetIdentification7 `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 Id"`
	SeqCntr Max9NumericText        `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 SeqCntr,omitempty"`
	Cntt    StatusReportContent8   `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 Cntt"`
}

type TriggerInformation1 struct {
	TrggrSrc PartyType5Code      `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 TrggrSrc"`
	SrcId    Max35Text           `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 SrcId"`
	TrggrTp  ExchangePolicy1Code `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 TrggrTp"`
	AddtlInf Max70Text           `xml:"urn:iso:std:iso:20022:tech:xsd:catm.001.001.08 AddtlInf,omitempty"`
}

// May be one of CDSP, CRCP, MDSP, MRCP, CRDO
type UserInterface4Code string

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