// Code generated by main. DO NOT EDIT.

package caad_005_001_02

import (
	"bytes"
	"encoding/base64"
	"encoding/xml"
	"time"
)

type AdditionalData1 struct {
	Tp  Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 Tp,omitempty"`
	Val Max2048Text `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 Val,omitempty"`
}

type AdditionalFee1 struct {
	Tp         TypeOfAmount10Code `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 Tp"`
	OthrTp     Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 OthrTp,omitempty"`
	FeePrgm    Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 FeePrgm,omitempty"`
	FeeDscrptr Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 FeeDscrptr,omitempty"`
	Amt        FeeAmount2         `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 Amt"`
	Labl       Max140Text         `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 Labl,omitempty"`
}

type AdditionalFeeReconciliation1 struct {
	RcncltnImpct ReconciliationImpact1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 RcncltnImpct"`
	Tp           TypeOfAmount10Code        `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 Tp"`
	OthrTp       Max35Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 OthrTp,omitempty"`
	Amt          float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 Amt"`
	Cnt          float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 Cnt"`
}

// May be one of EA2C, E3DC, DKP9, UKPT, UKA1, EA9C, EA5C
type Algorithm13Code string

// May be one of HS25, HS38, HS51
type Algorithm20Code string

// May be one of EA2C, E3DC, EA9C, EA5C, EA2R, EA9R, EA5R, E3DR, E36C, E36R, SD5C
type Algorithm23Code string

// May be one of HS25, HS38, HS51
type Algorithm5Code string

// May be one of ERSA, RSAO
type Algorithm7Code string

// May be one of MGF1
type Algorithm8Code string

type AlgorithmIdentification25 struct {
	Algo  Algorithm23Code `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 Algo"`
	Param Parameter7      `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 Param,omitempty"`
}

type AlgorithmIdentification26 struct {
	Algo  Algorithm8Code `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 Algo"`
	Param Algorithm5Code `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 Param,omitempty"`
}

type AlgorithmIdentification27 struct {
	Algo  Algorithm7Code `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 Algo"`
	Param Parameter13    `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 Param,omitempty"`
}

type AlgorithmIdentification28 struct {
	Algo  Algorithm13Code `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 Algo"`
	Param Parameter14     `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 Param,omitempty"`
}

type Amount5 struct {
	Amt float64 `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 Amt"`
	Sgn bool    `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 Sgn,omitempty"`
}

// May be one of CNAT, LATT, OATT, OUAT, CATT
type AttributeType1Code string

type BatchManagementInformation1 struct {
	ColltnId         Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 ColltnId,omitempty"`
	BtchId           Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 BtchId"`
	MsgSeqNb         Max15NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 MsgSeqNb,omitempty"`
	MsgChcksmInptVal Max140Binary     `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 MsgChcksmInptVal,omitempty"`
}

// May be one of LNGT, NUL8, NULG, NULL, RAND
type BytePadding1Code string

type CardProgrammeMode1 struct {
	Tp Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 Tp,omitempty"`
	Id Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 Id"`
}

// May be one of PART, FINL, OTHN, OTHP
type CardServiceType4Code string

type CertificateIssuer1 struct {
	RltvDstngshdNm []RelativeDistinguishedName1 `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 RltvDstngshdNm"`
}

type ContentInformationType20 struct {
	MACData MACData1          `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 MACData"`
	MAC     Max8HexBinaryText `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 MAC"`
}

// May be one of DATA, SIGN, EVLP, DGST, AUTH
type ContentType2Code string

// May be one of EVLP, IFSE
type ContentType3Code string

type Context9 struct {
	TxCntxt TransactionContext3 `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 TxCntxt,omitempty"`
}

type Document struct {
	RcncltnInitn ReconciliationInitiationV02 `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 RcncltnInitn"`
}

type EncryptedContent5 struct {
	CnttTp         ContentType2Code          `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 CnttTp"`
	CnttNcrptnAlgo AlgorithmIdentification25 `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 CnttNcrptnAlgo"`
	NcrptdDataElmt []EncryptedDataElement1   `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 NcrptdDataElmt"`
}

type EncryptedData1 struct {
	Ctrl           Exact1HexBinaryText     `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 Ctrl,omitempty"`
	KeySetIdr      Max8NumericText         `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 KeySetIdr,omitempty"`
	DrvdInf        Max32HexBinaryText      `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 DrvdInf,omitempty"`
	Algo           Max2NumericText         `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 Algo,omitempty"`
	KeyLngth       Max4NumericText         `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 KeyLngth,omitempty"`
	KeyPrtcn       Max2NumericText         `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 KeyPrtcn,omitempty"`
	KeyIndx        Max5NumericText         `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 KeyIndx,omitempty"`
	PddgMtd        Max2NumericText         `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 PddgMtd,omitempty"`
	NcrptdDataFrmt Max2NumericText         `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 NcrptdDataFrmt,omitempty"`
	NcrptdDataElmt []EncryptedDataElement1 `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 NcrptdDataElmt"`
}

type EncryptedData1Choice struct {
	BinryData   Max100KBinary `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 BinryData"`
	HexBinryVal string        `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 HexBinryVal"`
}

type EncryptedDataElement1 struct {
	Id                   ExternalEncryptedElementIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 Id,omitempty"`
	OthrId               Max35Text                                   `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 OthrId,omitempty"`
	NcrptdData           EncryptedData1Choice                        `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 NcrptdData"`
	ClearTxtDataFrmt     EncryptedDataFormat1Code                    `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 ClearTxtDataFrmt,omitempty"`
	OthrClearTxtDataFrmt Max35Text                                   `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 OthrClearTxtDataFrmt,omitempty"`
}

// May be one of ASCI, BINF, EBCD, HEXF, OTHN, OTHP
type EncryptedDataFormat1Code string

// May be one of TR34, TR31, CTCE, CBCE
type EncryptionFormat3Code string

type EnvelopedData6 struct {
	Vrsn       float64            `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 Vrsn,omitempty"`
	Rcpt       []Recipient7Choice `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 Rcpt"`
	NcrptdCntt EncryptedContent5  `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 NcrptdCntt,omitempty"`
}

type Environment3 struct {
	Orgtr PartyIdentification197 `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 Orgtr,omitempty"`
	Sndr  PartyIdentification197 `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 Sndr,omitempty"`
	Rcvr  PartyIdentification197 `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 Rcvr,omitempty"`
	Dstn  PartyIdentification197 `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 Dstn,omitempty"`
}

// May be no more than 12 items long
type Exact12Text string

// May be no more than 15 items long
type Exact15Text string

// Must match the pattern ([0-9A-F][0-9A-F]){1}
type Exact1HexBinaryText string

// Must match the pattern [0-9]{3}
type Exact3NumericText string

// Must match the pattern [0-9]{4}
type Exact4NumericText string

// Must match the pattern ([0-9A-F][0-9A-F]){1,3}
type ExternalEncryptedElementIdentification1Code string

type FeeAmount2 struct {
	Amt      float64                 `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 Amt"`
	Ccy      ISO3NumericCurrencyCode `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 Ccy,omitempty"`
	XchgRate float64                 `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 XchgRate,omitempty"`
	QtnDt    ISODateTime             `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 QtnDt,omitempty"`
	Sgn      bool                    `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 Sgn,omitempty"`
}

type FinancialReconcillation1 struct {
	RcncltnImpct ReconciliationImpact1Code   `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 RcncltnImpct"`
	Tp           ReconciliationCategory1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 Tp"`
	OthrTp       Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 OthrTp,omitempty"`
	Amt          float64                     `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 Amt"`
	Cnt          float64                     `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 Cnt"`
}

type GenericIdentification172 struct {
	Id     Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 Id"`
	Tp     PartyType17Code   `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 Tp,omitempty"`
	OthrTp Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 OthrTp,omitempty"`
	Assgnr PartyType18Code   `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 Assgnr,omitempty"`
	Ctry   Min2Max3AlphaText `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 Ctry,omitempty"`
	ShrtNm Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 ShrtNm,omitempty"`
}

type Header39 struct {
	MsgFctn        MessageFunction17Code       `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 MsgFctn"`
	PrtcolVrsn     Max2048Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 PrtcolVrsn"`
	XchgId         Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 XchgId,omitempty"`
	ReTrnsmssnCntr Max3NumericText             `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 ReTrnsmssnCntr,omitempty"`
	CreDtTm        ISODateTime                 `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 CreDtTm"`
	BtchMgmtInf    BatchManagementInformation1 `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 BtchMgmtInf,omitempty"`
	InitgPty       GenericIdentification172    `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 InitgPty"`
	RcptPty        GenericIdentification172    `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 RcptPty,omitempty"`
	TracData       []AdditionalData1           `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 TracData,omitempty"`
	Tracblt        []Traceability7             `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 Tracblt,omitempty"`
}

// Must match the pattern [0-9]{3,3}
type ISO3NumericCountryCode string

// Must match the pattern [0-9]{3,3}
type ISO3NumericCurrencyCode string

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

type ISOTime time.Time

func (t *ISOTime) UnmarshalText(text []byte) error {
	return (*xsdTime)(t).UnmarshalText(text)
}
func (t ISOTime) MarshalText() ([]byte, error) {
	return xsdTime(t).MarshalText()
}

type IssuerAndSerialNumber1 struct {
	Issr  CertificateIssuer1 `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 Issr"`
	SrlNb Max35Binary        `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 SrlNb"`
}

type KEK6 struct {
	Vrsn          float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 Vrsn,omitempty"`
	KEKId         KEKIdentifier6            `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 KEKId"`
	KeyNcrptnAlgo AlgorithmIdentification28 `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 KeyNcrptnAlgo"`
	NcrptdKey     Max500Binary              `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 NcrptdKey,omitempty"`
}

type KEKIdentifier2 struct {
	KeyId     Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 KeyId"`
	KeyVrsn   Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 KeyVrsn"`
	SeqNb     float64         `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 SeqNb,omitempty"`
	DerivtnId Min5Max16Binary `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 DerivtnId,omitempty"`
}

type KEKIdentifier6 struct {
	KeyId     Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 KeyId"`
	KeyVrsn   Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 KeyVrsn,omitempty"`
	SeqNb     float64         `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 SeqNb,omitempty"`
	DerivtnId Min5Max16Binary `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 DerivtnId,omitempty"`
}

type KeyTransport6 struct {
	Vrsn          float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 Vrsn,omitempty"`
	RcptId        Recipient5Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 RcptId"`
	KeyNcrptnAlgo AlgorithmIdentification27 `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 KeyNcrptnAlgo"`
	NcrptdKey     Max5000Binary             `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 NcrptdKey"`
}

type MACData1 struct {
	Ctrl         Exact1HexBinaryText `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 Ctrl"`
	KeySetIdr    Max8NumericText     `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 KeySetIdr"`
	DrvdInf      Max32HexBinaryText  `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 DrvdInf,omitempty"`
	Algo         Max2NumericText     `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 Algo"`
	KeyLngth     Max4NumericText     `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 KeyLngth,omitempty"`
	KeyPrtcn     Max2NumericText     `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 KeyPrtcn,omitempty"`
	KeyIndx      Max5NumericText     `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 KeyIndx,omitempty"`
	PddgMtd      Max2NumericText     `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 PddgMtd,omitempty"`
	InitlstnVctr Max32HexBinaryText  `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 InitlstnVctr,omitempty"`
}

type Max100KBinary []byte

func (t *Max100KBinary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Max100KBinary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

// Must match the pattern [0-9]{1,12}
type Max12NumericText string

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

// Must be at least 1 items long
type Max2048Text string

// Must match the pattern [0-9]{1,2}
type Max2NumericText string

// Must match the pattern ([0-9A-F][0-9A-F]){1,32}
type Max32HexBinaryText string

// Must be at least 1 items long
type Max350Text string

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

// Must match the pattern [0-9]{1,4}
type Max4NumericText string

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

// Must match the pattern [0-9]{1,5}
type Max5NumericText string

// Must match the pattern ([0-9A-F][0-9A-F]){1,8}
type Max8HexBinaryText string

// Must match the pattern [0-9]{1,8}
type Max8NumericText string

// May be one of NOTI, REQU, ADVC
type MessageFunction17Code string

type MessageReconcillation1 struct {
	Tp     ReconciliationMessageType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 Tp"`
	OthrTp Max35Text                      `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 OthrTp,omitempty"`
	Cnt    float64                        `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 Cnt"`
}

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

type Parameter13 struct {
	DgstAlgo     Algorithm20Code           `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 DgstAlgo,omitempty"`
	MskGnrtrAlgo AlgorithmIdentification26 `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 MskGnrtrAlgo,omitempty"`
}

type Parameter14 struct {
	NcrptnFrmt   EncryptionFormat3Code `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 NcrptnFrmt,omitempty"`
	InitlstnVctr Max500Binary          `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 InitlstnVctr,omitempty"`
	BPddg        BytePadding1Code      `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 BPddg,omitempty"`
}

type Parameter7 struct {
	InitlstnVctr Max500Binary     `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 InitlstnVctr,omitempty"`
	BPddg        BytePadding1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 BPddg,omitempty"`
}

type PartyIdentification197 struct {
	Id      Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 Id"`
	Assgnr  Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 Assgnr,omitempty"`
	Ctry    ISO3NumericCountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 Ctry,omitempty"`
	ShrtNm  Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 ShrtNm,omitempty"`
	AddtlId Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 AddtlId,omitempty"`
}

// May be one of OTHN, OTHP, ACQR, ACQP, CISS, CISP, AGNT
type PartyType17Code string

// May be one of ACQR, CISS, CSCH, AGNT
type PartyType18Code string

// May be one of URGT, HIGH, NORM
type Priority3Code string

type ProtectedData1 struct {
	CnttTp     ContentType3Code `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 CnttTp"`
	EnvlpdData EnvelopedData6   `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 EnvlpdData,omitempty"`
	NcrptdData EncryptedData1   `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 NcrptdData,omitempty"`
}

type Recipient5Choice struct {
	IssrAndSrlNb IssuerAndSerialNumber1 `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 IssrAndSrlNb"`
	KeyIdr       KEKIdentifier2         `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 KeyIdr"`
}

type Recipient7Choice struct {
	KeyTrnsprt KeyTransport6  `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 KeyTrnsprt"`
	KEK        KEK6           `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 KEK"`
	KeyIdr     KEKIdentifier6 `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 KeyIdr"`
}

type Reconciliation3 struct {
	Id        Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 Id,omitempty"`
	Dt        ISODate   `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 Dt,omitempty"`
	ChckptRef Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 ChckptRef,omitempty"`
}

// May be one of RVSL, OTHP, OTHN, FNCL, CGBK
type ReconciliationCategory1Code string

// May be one of INQR, INCU, CNVY
type ReconciliationFunction1Code string

// May be one of DEBT, CRDT
type ReconciliationImpact1Code string

type ReconciliationInitiation1 struct {
	Envt        Environment3         `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 Envt"`
	Cntxt       Context9             `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 Cntxt,omitempty"`
	Tx          Transaction94        `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 Tx"`
	PrtctdData  []ProtectedData1     `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 PrtctdData,omitempty"`
	SplmtryData []SupplementaryData1 `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 SplmtryData,omitempty"`
}

type ReconciliationInitiationV02 struct {
	Hdr      Header39                  `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 Hdr"`
	Body     ReconciliationInitiation1 `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 Body"`
	SctyTrlr ContentInformationType20  `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 SctyTrlr,omitempty"`
}

// May be one of BATR, CAMI, CAMR, CGBI, CGBR, EROR, FECI, FECR, FIAI, FIAR, FINR, FINI, FRDI, FRDR, FRRI, FRRR, INQI, INQR, KYEI, KYER, NWMI, NWMR, RECI, RECR, RTFI, RTFR, RTRR, RTRI, REVI, REVR, SERI, SERR, VERI, VERR, AMDT, ATHI, ATHR, BAMI, BAMR, BATI, ADDI, ADDR
type ReconciliationMessageType1Code string

type RelativeDistinguishedName1 struct {
	AttrTp  AttributeType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 AttrTp"`
	AttrVal Max140Text         `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 AttrVal"`
}

type SettlementReportingEntity1 struct {
	Tp Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 Tp,omitempty"`
	Id Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 Id,omitempty"`
}

type SettlementService2 struct {
	SttlmSvcApld  SettlementServiceMode1     `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 SttlmSvcApld,omitempty"`
	SttlmSvcDts   SettlementServiceDate1     `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 SttlmSvcDts,omitempty"`
	SttlmRptgNtty SettlementReportingEntity1 `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 SttlmRptgNtty,omitempty"`
}

type SettlementServiceDate1 struct {
	ReqdSttlmDt   ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 ReqdSttlmDt,omitempty"`
	SttlmDt       ISODate     `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 SttlmDt,omitempty"`
	SttlmTm       ISOTime     `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 SttlmTm,omitempty"`
	SttlmPrd      Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 SttlmPrd,omitempty"`
	SttlmCutOffTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 SttlmCutOffTm,omitempty"`
}

type SettlementServiceMode1 struct {
	Tp        Max35Text     `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 Tp,omitempty"`
	Id        Max35Text     `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 Id,omitempty"`
	ShrtNm    Max35Text     `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 ShrtNm,omitempty"`
	SttlmPrty Priority3Code `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 SttlmPrty,omitempty"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

type Traceability7 struct {
	RlayId      GenericIdentification172 `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 RlayId"`
	TracDtTmIn  ISODateTime              `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 TracDtTmIn,omitempty"`
	TracDtTmOut ISODateTime              `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 TracDtTmOut,omitempty"`
}

type Transaction94 struct {
	RcncltnFctn   ReconciliationFunction1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 RcncltnFctn"`
	RcncltnTp     CardServiceType4Code        `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 RcncltnTp"`
	OthrRcncltnTp Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 OthrRcncltnTp,omitempty"`
	TxId          TransactionIdentification12 `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 TxId,omitempty"`
	ReqdCcy       []Exact3NumericText         `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 ReqdCcy,omitempty"`
	RcncltnTtls   []TransactionTotals11       `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 RcncltnTtls,omitempty"`
	AddtlFees     []AdditionalFee1            `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 AddtlFees,omitempty"`
	AddtlData     []AdditionalData1           `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 AddtlData,omitempty"`
}

type TransactionContext3 struct {
	MrchntCtgyCd   Exact4NumericText         `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 MrchntCtgyCd,omitempty"`
	TxInitr        TransactionInitiator1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 TxInitr,omitempty"`
	CardPrgrmmApld CardProgrammeMode1        `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 CardPrgrmmApld,omitempty"`
	SttlmSvc       SettlementService2        `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 SttlmSvc,omitempty"`
	Rcncltn        Reconciliation3           `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 Rcncltn,omitempty"`
}

type TransactionIdentification12 struct {
	SysTracAudtNb      Max12NumericText                    `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 SysTracAudtNb"`
	TrnsmssnDtTm       ISODateTime                         `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 TrnsmssnDtTm"`
	RtrvlRefNb         Exact12Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 RtrvlRefNb,omitempty"`
	LifeCyclTracIdData TransactionLifeCycleIdentification2 `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 LifeCyclTracIdData,omitempty"`
}

// May be one of MERC, CUST
type TransactionInitiator1Code string

type TransactionLifeCycleIdentification2 struct {
	Id Exact15Text `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 Id"`
}

type TransactionTotals11 struct {
	Rcncltn         Reconciliation3                `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 Rcncltn,omitempty"`
	FinRcncltn      []FinancialReconcillation1     `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 FinRcncltn,omitempty"`
	MsgRcncltn      []MessageReconcillation1       `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 MsgRcncltn,omitempty"`
	AddtlFeeRcncltn []AdditionalFeeReconciliation1 `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 AddtlFeeRcncltn,omitempty"`
	Ccy             Min2Max3NumericText            `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 Ccy"`
	NetAmtRcncltn   Amount5                        `xml:"urn:iso:std:iso:20022:tech:xsd:caad.005.001.02 NetAmtRcncltn,omitempty"`
}

// May be one of INTC, FEEP, OTHN, OTHP, FEEA
type TypeOfAmount10Code string

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

type xsdTime time.Time

func (t *xsdTime) UnmarshalText(text []byte) error {
	return _unmarshalTime(text, (*time.Time)(t), "15:04:05.999999999")
}
func (t xsdTime) MarshalText() ([]byte, error) {
	return _marshalTime((time.Time)(t), "15:04:05.999999999")
}
func (t xsdTime) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if (time.Time)(t).IsZero() {
		return nil
	}
	m, err := t.MarshalText()
	if err != nil {
		return err
	}
	return e.EncodeElement(m, start)
}
func (t xsdTime) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if (time.Time)(t).IsZero() {
		return xml.Attr{}, nil
	}
	m, err := t.MarshalText()
	return xml.Attr{Name: name, Value: string(m)}, err
}
