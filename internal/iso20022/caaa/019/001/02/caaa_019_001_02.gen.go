// Code generated by main. DO NOT EDIT.

package caaa_019_001_02

import (
	"bytes"
	"encoding/base64"
	"encoding/xml"
	"time"
)

type AcceptorCancellationAdviceResponse7 struct {
	Envt     CardPaymentEnvironment69              `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 Envt"`
	Tx       CardPaymentTransactionAdviceResponse6 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 Tx"`
	TMSTrggr TMSTrigger1                           `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 TMSTrggr,omitempty"`
}

type AcceptorCurrencyConversionAdviceResponseV02 struct {
	Hdr              Header36                            `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 Hdr"`
	CcyConvsAdvcRspn AcceptorCancellationAdviceResponse7 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 CcyConvsAdvcRspn"`
	SctyTrlr         ContentInformationType16            `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 SctyTrlr,omitempty"`
}

// May be one of HS25, HS38, HS51, HS01, SH31, SH32, SH33, SH35, SHK1, SHK2
type Algorithm16Code string

// May be one of MACC, MCCS, CMA1, MCC1, CMA9, CMA5, CMA2, CM31, CM32, CM33, MCS3, CCA1, CCA2, CCA3
type Algorithm17Code string

// May be one of EA2C, E3DC, DKP9, UKPT, UKA1, EA9C, EA5C, DA12, DA19, DA25, N108, EA5R, EA9R, EA2R, E3DR, E36C, E36R, SD5C
type Algorithm18Code string

// May be one of ERSA, RSAO
type Algorithm7Code string

// May be one of MGF1
type Algorithm8Code string

type AlgorithmIdentification18 struct {
	Algo  Algorithm8Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 Algo"`
	Param Parameter9     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 Param,omitempty"`
}

type AlgorithmIdentification19 struct {
	Algo  Algorithm7Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 Algo"`
	Param Parameter10    `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 Param,omitempty"`
}

type AlgorithmIdentification22 struct {
	Algo  Algorithm17Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 Algo"`
	Param Parameter7      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 Param,omitempty"`
}

type AlgorithmIdentification23 struct {
	Algo  Algorithm18Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 Algo"`
	Param Parameter12     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 Param,omitempty"`
}

type AlgorithmIdentification24 struct {
	Algo  Algorithm18Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 Algo"`
	Param Parameter12     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 Param,omitempty"`
}

// May be one of CNAT, LATT, OATT, OUAT, CATT
type AttributeType1Code string

type AuthenticatedData5 struct {
	Vrsn        float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 Vrsn,omitempty"`
	Rcpt        []Recipient6Choice        `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 Rcpt"`
	MACAlgo     AlgorithmIdentification22 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 MACAlgo"`
	NcpsltdCntt EncapsulatedContent3      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 NcpsltdCntt"`
	MAC         Max140Binary              `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 MAC"`
}

// May be one of LNGT, NUL8, NULG, NULL, RAND
type BytePadding1Code string

type CardPaymentEnvironment69 struct {
	AcqrrId  GenericIdentification53 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 AcqrrId,omitempty"`
	MrchntId GenericIdentification32 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 MrchntId,omitempty"`
	POIId    GenericIdentification32 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 POIId,omitempty"`
	Card     PaymentCard28           `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 Card,omitempty"`
	PmtTkn   CardPaymentToken4       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 PmtTkn,omitempty"`
}

type CardPaymentToken4 struct {
	Tkn           Min8Max28NumericText     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 Tkn,omitempty"`
	CardSeqNb     Min2Max3NumericText      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 CardSeqNb,omitempty"`
	TknXpryDt     Max10Text                `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 TknXpryDt,omitempty"`
	TknChrtc      []Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 TknChrtc,omitempty"`
	TknRqstr      PaymentTokenIdentifiers1 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 TknRqstr,omitempty"`
	TknAssrncLvl  float64                  `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 TknAssrncLvl,omitempty"`
	TknAssrncData Max500Binary             `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 TknAssrncData,omitempty"`
}

type CardPaymentTransactionAdviceResponse6 struct {
	SaleRefId Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 SaleRefId,omitempty"`
	TxId      TransactionIdentifier1 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 TxId"`
	InitrTxId Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 InitrTxId,omitempty"`
	RcptTxId  Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 RcptTxId,omitempty"`
	RcncltnId Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 RcncltnId,omitempty"`
	Rspn      Response4Code          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 Rspn"`
}

// May be one of COMM, CONS
type CardProductType1Code string

type CertificateIssuer1 struct {
	RltvDstngshdNm []RelativeDistinguishedName1 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 RltvDstngshdNm"`
}

type ContentInformationType16 struct {
	CnttTp       ContentType2Code   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 CnttTp"`
	AuthntcdData AuthenticatedData5 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 AuthntcdData"`
}

type ContentInformationType17 struct {
	CnttTp     ContentType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 CnttTp"`
	EnvlpdData EnvelopedData5   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 EnvlpdData"`
}

// May be one of DATA, SIGN, EVLP, DGST, AUTH
type ContentType2Code string

type Document struct {
	AccptrCcyConvsAdvcRspn AcceptorCurrencyConversionAdviceResponseV02 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 AccptrCcyConvsAdvcRspn"`
}

type EncapsulatedContent3 struct {
	CnttTp ContentType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 CnttTp"`
	Cntt   Max100KBinary    `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 Cntt,omitempty"`
}

type EncryptedContent4 struct {
	CnttTp         ContentType2Code          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 CnttTp"`
	CnttNcrptnAlgo AlgorithmIdentification24 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 CnttNcrptnAlgo,omitempty"`
	NcrptdData     Max100KBinary             `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 NcrptdData"`
}

// May be one of TR31, TR34, I238
type EncryptionFormat2Code string

type EnvelopedData5 struct {
	Vrsn       float64                `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 Vrsn,omitempty"`
	OrgtrInf   OriginatorInformation1 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 OrgtrInf,omitempty"`
	Rcpt       []Recipient6Choice     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 Rcpt"`
	NcrptdCntt EncryptedContent4      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 NcrptdCntt,omitempty"`
}

// Must match the pattern [a-zA-Z0-9]{3}
type Exact3AlphaNumericText string

// Must match the pattern [0-9]{3}
type Exact3NumericText string

type GenericIdentification32 struct {
	Id     Max35Text      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 Id"`
	Tp     PartyType3Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 Tp,omitempty"`
	Issr   PartyType4Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 Issr,omitempty"`
	ShrtNm Max35Text      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 ShrtNm,omitempty"`
}

type GenericIdentification53 struct {
	Id     Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 Id"`
	Tp     PartyType3Code    `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 Tp,omitempty"`
	Issr   PartyType4Code    `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 Issr,omitempty"`
	Ctry   Min2Max3AlphaText `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 Ctry,omitempty"`
	ShrtNm Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 ShrtNm,omitempty"`
}

type GenericIdentification76 struct {
	Id     Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 Id"`
	Tp     PartyType3Code    `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 Tp"`
	Issr   PartyType4Code    `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 Issr,omitempty"`
	Ctry   Min2Max3AlphaText `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 Ctry,omitempty"`
	ShrtNm Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 ShrtNm,omitempty"`
}

type GenericIdentification94 struct {
	Id       Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 Id"`
	Tp       PartyType3Code     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 Tp,omitempty"`
	Issr     PartyType4Code     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 Issr,omitempty"`
	Ctry     Min2Max3AlphaText  `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 Ctry,omitempty"`
	ShrtNm   Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 ShrtNm,omitempty"`
	RmotAccs NetworkParameters5 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 RmotAccs,omitempty"`
}

type Header36 struct {
	MsgFctn        MessageFunction14Code   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 MsgFctn"`
	PrtcolVrsn     Max6Text                `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 PrtcolVrsn"`
	XchgId         float64                 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 XchgId"`
	ReTrnsmssnCntr Max3NumericText         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 ReTrnsmssnCntr,omitempty"`
	CreDtTm        ISODateTime             `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 CreDtTm"`
	InitgPty       GenericIdentification53 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 InitgPty"`
	RcptPty        GenericIdentification94 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 RcptPty,omitempty"`
	Tracblt        []Traceability5         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 Tracblt,omitempty"`
}

type ISODateTime time.Time

func (t *ISODateTime) UnmarshalText(text []byte) error {
	return (*xsdDateTime)(t).UnmarshalText(text)
}
func (t ISODateTime) MarshalText() ([]byte, error) {
	return xsdDateTime(t).MarshalText()
}

type IssuerAndSerialNumber1 struct {
	Issr  CertificateIssuer1 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 Issr"`
	SrlNb Max35Binary        `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 SrlNb"`
}

type KEK5 struct {
	Vrsn          float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 Vrsn,omitempty"`
	KEKId         KEKIdentifier2            `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 KEKId"`
	KeyNcrptnAlgo AlgorithmIdentification23 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 KeyNcrptnAlgo"`
	NcrptdKey     Max500Binary              `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 NcrptdKey"`
}

type KEKIdentifier2 struct {
	KeyId     Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 KeyId"`
	KeyVrsn   Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 KeyVrsn"`
	SeqNb     float64         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 SeqNb,omitempty"`
	DerivtnId Min5Max16Binary `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 DerivtnId,omitempty"`
}

type KeyTransport5 struct {
	Vrsn          float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 Vrsn,omitempty"`
	RcptId        Recipient5Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 RcptId"`
	KeyNcrptnAlgo AlgorithmIdentification19 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 KeyNcrptnAlgo"`
	NcrptdKey     Max5000Binary             `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 NcrptdKey"`
}

type Max100KBinary []byte

func (t *Max100KBinary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Max100KBinary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

// Must be at least 1 items long
type Max104Text string

type Max10KBinary []byte

func (t *Max10KBinary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Max10KBinary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

// Must be at least 1 items long
type Max10Text string

type Max140Binary []byte

func (t *Max140Binary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Max140Binary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

// Must be at least 1 items long
type Max140Text string

// Must match the pattern [0-9]{1,15}
type Max15NumericText string

type Max35Binary []byte

func (t *Max35Binary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Max35Binary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

// Must be at least 1 items long
type Max35Text string

// Must be at least 1 items long
type Max37Text string

// Must match the pattern [0-9]{1,3}
type Max3NumericText string

// Must be at least 1 items long
type Max3Text string

// Must be at least 1 items long
type Max45Text string

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

// Must be at least 1 items long
type Max76Text string

// May be one of AUTQ, AUTP, CCAV, CCAK, CCAQ, CCAP, CMPV, CMPK, DCAV, DCRR, DCCQ, DCCP, DGNP, DGNQ, FAUQ, FAUP, FCMV, FCMK, FRVA, FRVR, RCLQ, RCLP, RVRA, RVRR, CDDQ, CDDK, CDDR, CDDP
type MessageFunction14Code string

// Must match the pattern [a-zA-Z]{2,3}
type Min2Max3AlphaText string

// Must match the pattern [0-9]{2,3}
type Min2Max3NumericText string

type Min5Max16Binary []byte

func (t *Min5Max16Binary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Min5Max16Binary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

// Must match the pattern [0-9]{8,28}
type Min8Max28NumericText string

type NetworkParameters4 struct {
	NtwkTp NetworkType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 NtwkTp"`
	AdrVal Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 AdrVal"`
}

type NetworkParameters5 struct {
	Adr        []NetworkParameters4 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 Adr"`
	UsrNm      Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 UsrNm,omitempty"`
	AccsCd     Max35Binary          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 AccsCd,omitempty"`
	SvrCert    []Max10KBinary       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 SvrCert,omitempty"`
	SvrCertIdr []Max140Binary       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 SvrCertIdr,omitempty"`
	ClntCert   []Max10KBinary       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 ClntCert,omitempty"`
	SctyPrfl   Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 SctyPrfl,omitempty"`
}

// May be one of IPNW, PSTN
type NetworkType1Code string

type OriginatorInformation1 struct {
	Cert []Max5000Binary `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 Cert,omitempty"`
}

type Parameter10 struct {
	NcrptnFrmt   EncryptionFormat2Code     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 NcrptnFrmt,omitempty"`
	DgstAlgo     Algorithm16Code           `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 DgstAlgo,omitempty"`
	MskGnrtrAlgo AlgorithmIdentification18 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 MskGnrtrAlgo,omitempty"`
}

type Parameter12 struct {
	NcrptnFrmt   EncryptionFormat2Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 NcrptnFrmt,omitempty"`
	InitlstnVctr Max500Binary          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 InitlstnVctr,omitempty"`
	BPddg        BytePadding1Code      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 BPddg,omitempty"`
}

type Parameter7 struct {
	InitlstnVctr Max500Binary     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 InitlstnVctr,omitempty"`
	BPddg        BytePadding1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 BPddg,omitempty"`
}

type Parameter9 struct {
	DgstAlgo Algorithm16Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 DgstAlgo,omitempty"`
}

// May be one of OPOI, MERC, ACCP, ITAG, ACQR, CISS, DLIS
type PartyType3Code string

// May be one of MERC, ACCP, ITAG, ACQR, CISS, TAXH
type PartyType4Code string

type PaymentCard28 struct {
	PrtctdCardData ContentInformationType17 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 PrtctdCardData,omitempty"`
	PrvtCardData   Max100KBinary            `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 PrvtCardData,omitempty"`
	PlainCardData  PlainCardData15          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 PlainCardData,omitempty"`
	PmtAcctRef     Max70Text                `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 PmtAcctRef,omitempty"`
	MskdPAN        string                   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 MskdPAN,omitempty"`
	IssrBIN        Max15NumericText         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 IssrBIN,omitempty"`
	CardCtryCd     Max3Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 CardCtryCd,omitempty"`
	CardCcyCd      Exact3AlphaNumericText   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 CardCcyCd,omitempty"`
	CardPdctPrfl   Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 CardPdctPrfl,omitempty"`
	CardBrnd       Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 CardBrnd,omitempty"`
	CardPdctTp     CardProductType1Code     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 CardPdctTp,omitempty"`
	CardPdctSubTp  Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 CardPdctSubTp,omitempty"`
	IntrnlCard     bool                     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 IntrnlCard,omitempty"`
	AllwdPdct      []Max70Text              `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 AllwdPdct,omitempty"`
	SvcOptn        Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 SvcOptn,omitempty"`
	AddtlCardData  Max70Text                `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 AddtlCardData,omitempty"`
}

type PaymentTokenIdentifiers1 struct {
	PrvdrId Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 PrvdrId"`
	RqstrId Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 RqstrId"`
}

type PlainCardData15 struct {
	PAN       Min8Max28NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 PAN"`
	CardSeqNb Min2Max3NumericText  `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 CardSeqNb,omitempty"`
	FctvDt    Max10Text            `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 FctvDt,omitempty"`
	XpryDt    Max10Text            `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 XpryDt"`
	SvcCd     Exact3NumericText    `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 SvcCd,omitempty"`
	Trck1     Max76Text            `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 Trck1,omitempty"`
	Trck2     Max37Text            `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 Trck2,omitempty"`
	Trck3     Max104Text           `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 Trck3,omitempty"`
	CrdhldrNm Max45Text            `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 CrdhldrNm,omitempty"`
}

type Recipient5Choice struct {
	IssrAndSrlNb IssuerAndSerialNumber1 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 IssrAndSrlNb"`
	KeyIdr       KEKIdentifier2         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 KeyIdr"`
}

type Recipient6Choice struct {
	KeyTrnsprt KeyTransport5  `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 KeyTrnsprt"`
	KEK        KEK5           `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 KEK"`
	KeyIdr     KEKIdentifier2 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 KeyIdr"`
}

type RelativeDistinguishedName1 struct {
	AttrTp  AttributeType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 AttrTp"`
	AttrVal Max140Text         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 AttrVal"`
}

// May be one of APPR, DECL, PART
type Response4Code string

// May be one of CRIT, ASAP, DTIM
type TMSContactLevel1Code string

type TMSTrigger1 struct {
	TMSCtctLvl  TMSContactLevel1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 TMSCtctLvl"`
	TMSId       Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 TMSId,omitempty"`
	TMSCtctDtTm ISODateTime          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 TMSCtctDtTm,omitempty"`
}

type Traceability5 struct {
	RlayId      GenericIdentification76 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 RlayId"`
	PrtcolNm    Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 PrtcolNm,omitempty"`
	PrtcolVrsn  Max6Text                `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 PrtcolVrsn,omitempty"`
	TracDtTmIn  ISODateTime             `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 TracDtTmIn"`
	TracDtTmOut ISODateTime             `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 TracDtTmOut"`
}

type TransactionIdentifier1 struct {
	TxDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 TxDtTm"`
	TxRef  Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.019.001.02 TxRef"`
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
