// Code generated by main. DO NOT EDIT.

package caaa_002_001_04

import (
	"bytes"
	"encoding/base64"
	"encoding/xml"
	"time"
)

type AcceptorAuthorisationResponse4 struct {
	Envt        CardPaymentEnvironment33 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 Envt"`
	Tx          CardPaymentTransaction38 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 Tx"`
	TxRspn      CardPaymentTransaction39 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 TxRspn"`
	SplmtryData []SupplementaryData1     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 SplmtryData,omitempty"`
}

type AcceptorAuthorisationResponseV04 struct {
	Hdr         Header10                       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 Hdr"`
	AuthstnRspn AcceptorAuthorisationResponse4 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 AuthstnRspn"`
	SctyTrlr    ContentInformationType11       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 SctyTrlr"`
}

type Action3 struct {
	ActnTp    ActionType3Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 ActnTp"`
	MsgToPres ActionMessage1  `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 MsgToPres,omitempty"`
}

type ActionMessage1 struct {
	MsgDstn      UserInterface1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 MsgDstn"`
	MsgCntt      Max256Text         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 MsgCntt"`
	MsgCnttSgntr Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 MsgCnttSgntr,omitempty"`
}

// May be one of BUSY, CPTR, DISP, NOVR, RQID, PINL, PINR, PRNT, RFRL, RQDT, DCCQ
type ActionType3Code string

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

type AlgorithmIdentification11 struct {
	Algo  Algorithm7Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 Algo"`
	Param Parameter4     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 Param,omitempty"`
}

type AlgorithmIdentification12 struct {
	Algo  Algorithm8Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 Algo"`
	Param Parameter5     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 Param,omitempty"`
}

type AlgorithmIdentification13 struct {
	Algo  Algorithm13Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 Algo"`
	Param Parameter6      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 Param,omitempty"`
}

type AlgorithmIdentification14 struct {
	Algo  Algorithm15Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 Algo"`
	Param Parameter6      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 Param,omitempty"`
}

type AlgorithmIdentification15 struct {
	Algo  Algorithm12Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 Algo"`
	Param Parameter7      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 Param,omitempty"`
}

type AmountAndDirection41 struct {
	Amt CurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 Amt"`
	Sgn bool              `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 Sgn,omitempty"`
}

// May be one of CNAT, LATT, OATT, OUAT, CATT
type AttributeType1Code string

type AuthenticatedData4 struct {
	Vrsn        float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 Vrsn,omitempty"`
	Rcpt        []Recipient4Choice        `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 Rcpt"`
	MACAlgo     AlgorithmIdentification15 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 MACAlgo"`
	NcpsltdCntt EncapsulatedContent3      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 NcpsltdCntt"`
	MAC         Max140Binary              `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 MAC"`
}

// May be one of ICCD, AGNT, MERC, ACQR, ISSR, TRML
type AuthenticationEntity2Code string

// May be one of NPIN, PPSG, PSWD, SCRT, SCNL, SNCT, CPSG, ADDB, BIOM, CDHI, CRYP, CSCV, PSVE, CSEC, ADDS, TOKN, MANU, FPIN
type AuthenticationMethod4Code string

type AuthorisationResult4 struct {
	AuthstnNtty   GenericIdentification70 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 AuthstnNtty,omitempty"`
	RspnToAuthstn ResponseType1           `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 RspnToAuthstn"`
	AuthstnCd     Min6Max8Text            `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 AuthstnCd,omitempty"`
	CmpltnReqrd   bool                    `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 CmpltnReqrd,omitempty"`
	TMSTrggr      TMSTrigger1             `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 TMSTrggr,omitempty"`
}

// May be one of LNGT, NUL8, NULG, NULL, RAND
type BytePadding1Code string

// May be one of CTDP, CHCK, CRDT, CURR, CDBT, DFLT, EPRS, HEQL, ISTL, INVS, LCDT, MBNW, MNMK, MNMC, MTGL, RTRM, RVLV, SVNG, STBD, UVRL
type CardAccountType2Code string

type CardPaymentEnvironment33 struct {
	AcqrrId  GenericIdentification53 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 AcqrrId,omitempty"`
	MrchntId GenericIdentification32 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 MrchntId,omitempty"`
	POIId    GenericIdentification32 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 POIId"`
	Card     PaymentCard10           `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 Card,omitempty"`
	PmtTkn   CardPaymentToken2       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 PmtTkn,omitempty"`
}

type CardPaymentToken2 struct {
	TknChrtc     []Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 TknChrtc,omitempty"`
	TknAssrncLvl float64     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 TknAssrncLvl,omitempty"`
}

type CardPaymentTransaction38 struct {
	SaleRefId    Max35Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 SaleRefId,omitempty"`
	TxId         TransactionIdentifier1          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 TxId"`
	RcptTxId     Max35Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 RcptTxId,omitempty"`
	RcncltnId    Max35Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 RcncltnId,omitempty"`
	IntrchngData Max140Text                      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 IntrchngData,omitempty"`
	TxDtls       CardPaymentTransactionDetails20 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 TxDtls"`
}

type CardPaymentTransaction39 struct {
	AuthstnRslt  AuthorisationResult4             `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 AuthstnRslt"`
	TxVrfctnRslt []TransactionVerificationResult3 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 TxVrfctnRslt,omitempty"`
	DclndPdctCd  []Max70Text                      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 DclndPdctCd,omitempty"`
	Bal          AmountAndDirection41             `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 Bal,omitempty"`
	PrtctdBal    ContentInformationType10         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 PrtctdBal,omitempty"`
	Actn         []Action3                        `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 Actn,omitempty"`
	CcyConvs     CurrencyConversion3              `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 CcyConvs,omitempty"`
}

type CardPaymentTransactionDetails20 struct {
	Ccy         CurrencyCode         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 Ccy"`
	TtlAmt      float64              `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 TtlAmt"`
	DtldAmt     DetailedAmount7      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 DtldAmt,omitempty"`
	VldtyDt     ISODate              `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 VldtyDt,omitempty"`
	AcctTp      CardAccountType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 AcctTp,omitempty"`
	ICCRltdData Max10000Binary       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 ICCRltdData,omitempty"`
}

// May be one of COMM, CONS
type CardProductType1Code string

type CertificateIssuer1 struct {
	RltvDstngshdNm []RelativeDistinguishedName1 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 RltvDstngshdNm"`
}

type Commission18 struct {
	Rate     float64    `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 Rate"`
	AddtlInf Max350Text `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 AddtlInf,omitempty"`
}

type Commission19 struct {
	Amt      float64    `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 Amt"`
	AddtlInf Max350Text `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 AddtlInf,omitempty"`
}

type ContentInformationType10 struct {
	CnttTp     ContentType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 CnttTp"`
	EnvlpdData EnvelopedData4   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 EnvlpdData"`
}

type ContentInformationType11 struct {
	CnttTp       ContentType2Code     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 CnttTp"`
	AuthntcdData []AuthenticatedData4 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 AuthntcdData,omitempty"`
}

// May be one of DATA, SIGN, EVLP, DGST, AUTH
type ContentType2Code string

type CurrencyAndAmount struct {
	Value float64      `xml:",chardata"`
	Ccy   CurrencyCode `xml:"Ccy,attr"`
}

// Must match the pattern [A-Z]{3,3}
type CurrencyCode string

type CurrencyConversion2 struct {
	CcyConvsId    Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 CcyConvsId,omitempty"`
	TrgtCcy       CurrencyDetails1 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 TrgtCcy"`
	RsltgAmt      float64          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 RsltgAmt"`
	XchgRate      float64          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 XchgRate"`
	XchgRateDcml  float64          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 XchgRateDcml,omitempty"`
	NvrtdXchgRate float64          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 NvrtdXchgRate,omitempty"`
	QtnDt         ISODateTime      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 QtnDt,omitempty"`
	VldUntil      ISODateTime      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 VldUntil,omitempty"`
	SrcCcy        CurrencyDetails1 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 SrcCcy"`
	OrgnlAmt      float64          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 OrgnlAmt"`
	ComssnDtls    []Commission19   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 ComssnDtls,omitempty"`
	MrkUpDtls     []Commission18   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 MrkUpDtls,omitempty"`
	DclrtnDtls    Max2048Text      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 DclrtnDtls,omitempty"`
}

type CurrencyConversion3 struct {
	Rslt    CurrencyConversionResponse1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 Rslt"`
	RsltRsn Max35Text                       `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 RsltRsn,omitempty"`
	Convs   CurrencyConversion2             `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 Convs,omitempty"`
}

// May be one of ODCC, DCCA, ICRD, IMER, IPRD, IRAT, NDCC
type CurrencyConversionResponse1Code string

type CurrencyDetails1 struct {
	AlphaCd CurrencyCode      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 AlphaCd"`
	NmrcCd  Exact3NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 NmrcCd"`
	Dcml    float64           `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 Dcml"`
	Nm      Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 Nm,omitempty"`
}

type DetailedAmount4 struct {
	Amt  float64    `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 Amt"`
	Labl Max140Text `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 Labl,omitempty"`
}

type DetailedAmount7 struct {
	CshBck      float64           `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 CshBck,omitempty"`
	Grtty       float64           `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 Grtty,omitempty"`
	Fees        []DetailedAmount4 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 Fees,omitempty"`
	Rbt         []DetailedAmount4 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 Rbt,omitempty"`
	ValAddedTax []DetailedAmount4 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 ValAddedTax,omitempty"`
	Srchrg      []DetailedAmount4 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 Srchrg,omitempty"`
}

type Document struct {
	AccptrAuthstnRspn AcceptorAuthorisationResponseV04 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 AccptrAuthstnRspn"`
}

type EncapsulatedContent3 struct {
	CnttTp ContentType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 CnttTp"`
	Cntt   Max100KBinary    `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 Cntt,omitempty"`
}

type EncryptedContent3 struct {
	CnttTp         ContentType2Code          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 CnttTp"`
	CnttNcrptnAlgo AlgorithmIdentification14 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 CnttNcrptnAlgo"`
	NcrptdData     Max100KBinary             `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 NcrptdData"`
}

// May be one of TR31, TR34
type EncryptionFormat1Code string

type EnvelopedData4 struct {
	Vrsn       float64            `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 Vrsn,omitempty"`
	Rcpt       []Recipient4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 Rcpt"`
	NcrptdCntt EncryptedContent3  `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 NcrptdCntt,omitempty"`
}

// Must match the pattern [0-9]{3}
type Exact3NumericText string

type GenericIdentification32 struct {
	Id     Max35Text      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 Id"`
	Tp     PartyType3Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 Tp,omitempty"`
	Issr   PartyType4Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 Issr,omitempty"`
	ShrtNm Max35Text      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 ShrtNm,omitempty"`
}

type GenericIdentification53 struct {
	Id     Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 Id"`
	Tp     PartyType3Code    `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 Tp,omitempty"`
	Issr   PartyType4Code    `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 Issr,omitempty"`
	Ctry   Min2Max3AlphaText `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 Ctry,omitempty"`
	ShrtNm Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 ShrtNm,omitempty"`
}

type GenericIdentification70 struct {
	Id     Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 Id,omitempty"`
	Tp     PartyType3Code    `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 Tp"`
	Issr   PartyType4Code    `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 Issr,omitempty"`
	Ctry   Min2Max3AlphaText `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 Ctry,omitempty"`
	ShrtNm Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 ShrtNm,omitempty"`
}

type GenericIdentification76 struct {
	Id     Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 Id"`
	Tp     PartyType3Code    `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 Tp"`
	Issr   PartyType4Code    `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 Issr,omitempty"`
	Ctry   Min2Max3AlphaText `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 Ctry,omitempty"`
	ShrtNm Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 ShrtNm,omitempty"`
}

type Header10 struct {
	MsgFctn    MessageFunction4Code    `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 MsgFctn"`
	PrtcolVrsn Max6Text                `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 PrtcolVrsn"`
	XchgId     Max3NumericText         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 XchgId"`
	CreDtTm    ISODateTime             `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 CreDtTm"`
	InitgPty   GenericIdentification53 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 InitgPty"`
	RcptPty    GenericIdentification53 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 RcptPty,omitempty"`
	Tracblt    []Traceability2         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 Tracblt,omitempty"`
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
	Issr  CertificateIssuer1 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 Issr"`
	SrlNb Max35Binary        `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 SrlNb"`
}

type KEK4 struct {
	Vrsn          float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 Vrsn,omitempty"`
	KEKId         KEKIdentifier2            `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 KEKId"`
	KeyNcrptnAlgo AlgorithmIdentification13 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 KeyNcrptnAlgo"`
	NcrptdKey     Max500Binary              `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 NcrptdKey"`
}

type KEKIdentifier2 struct {
	KeyId     Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 KeyId"`
	KeyVrsn   Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 KeyVrsn"`
	SeqNb     float64         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 SeqNb,omitempty"`
	DerivtnId Min5Max16Binary `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 DerivtnId,omitempty"`
}

type KeyTransport4 struct {
	Vrsn          float64                   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 Vrsn,omitempty"`
	RcptId        Recipient5Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 RcptId"`
	KeyNcrptnAlgo AlgorithmIdentification11 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 KeyNcrptnAlgo"`
	NcrptdKey     Max5000Binary             `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 NcrptdKey"`
}

type Max10000Binary []byte

func (t *Max10000Binary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Max10000Binary) MarshalText() ([]byte, error) {
	return xsdBase64Binary(t).MarshalText()
}

type Max100KBinary []byte

func (t *Max100KBinary) UnmarshalText(text []byte) error {
	return (*xsdBase64Binary)(t).UnmarshalText(text)
}
func (t Max100KBinary) MarshalText() ([]byte, error) {
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

// Must be at least 1 items long
type Max2048Text string

// Must be at least 1 items long
type Max256Text string

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
type Max500Text string

// Must be at least 1 items long
type Max6Text string

// Must be at least 1 items long
type Max70Text string

// May be one of AUTQ, AUTP, FAUQ, FAUP, CMPV, CMPK, FCMV, FCMK, RVRA, RVRR, FRVA, FRVR, CCAQ, CCAP, CCAV, CCAK, DGNP, DGNQ, RCLQ, RCLP, RJCT, DCCQ, DCCP
type MessageFunction4Code string

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

// Must be at least 6 items long
type Min6Max8Text string

// Must match the pattern [0-9]{8,28}
type Min8Max28NumericText string

type Parameter4 struct {
	NcrptnFrmt   EncryptionFormat1Code     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 NcrptnFrmt,omitempty"`
	DgstAlgo     Algorithm11Code           `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 DgstAlgo,omitempty"`
	MskGnrtrAlgo AlgorithmIdentification12 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 MskGnrtrAlgo,omitempty"`
}

type Parameter5 struct {
	DgstAlgo Algorithm11Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 DgstAlgo,omitempty"`
}

type Parameter6 struct {
	NcrptnFrmt   EncryptionFormat1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 NcrptnFrmt,omitempty"`
	InitlstnVctr Max500Binary          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 InitlstnVctr,omitempty"`
	BPddg        BytePadding1Code      `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 BPddg,omitempty"`
}

type Parameter7 struct {
	InitlstnVctr Max500Binary     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 InitlstnVctr,omitempty"`
	BPddg        BytePadding1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 BPddg,omitempty"`
}

// May be one of OPOI, MERC, ACCP, ITAG, ACQR, CISS, DLIS
type PartyType3Code string

// May be one of MERC, ACCP, ITAG, ACQR, CISS, TAXH
type PartyType4Code string

type PaymentCard10 struct {
	PrtctdCardData ContentInformationType10 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 PrtctdCardData,omitempty"`
	PlainCardData  PlainCardData8           `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 PlainCardData,omitempty"`
	MskdPAN        string                   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 MskdPAN,omitempty"`
	CardBrnd       Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 CardBrnd,omitempty"`
	CardPdctTp     CardProductType1Code     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 CardPdctTp,omitempty"`
}

type PlainCardData8 struct {
	PAN       Min8Max28NumericText `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 PAN"`
	CardSeqNb Min2Max3NumericText  `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 CardSeqNb,omitempty"`
	FctvDt    Max10Text            `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 FctvDt,omitempty"`
	XpryDt    Max10Text            `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 XpryDt"`
}

type Recipient4Choice struct {
	KeyTrnsprt KeyTransport4  `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 KeyTrnsprt"`
	KEK        KEK4           `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 KEK"`
	KeyIdr     KEKIdentifier2 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 KeyIdr"`
}

type Recipient5Choice struct {
	IssrAndSrlNb IssuerAndSerialNumber1 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 IssrAndSrlNb"`
	KeyIdr       KEKIdentifier2         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 KeyIdr"`
}

type RelativeDistinguishedName1 struct {
	AttrTp  AttributeType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 AttrTp"`
	AttrVal Max140Text         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 AttrVal"`
}

// May be one of DECL, APPR, PART, TECH
type Response1Code string

type ResponseType1 struct {
	Rspn    Response1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 Rspn"`
	RspnRsn Max35Text     `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 RspnRsn,omitempty"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

// May be one of CRIT, ASAP, DTIM
type TMSContactLevel1Code string

type TMSTrigger1 struct {
	TMSCtctLvl  TMSContactLevel1Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 TMSCtctLvl"`
	TMSId       Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 TMSId,omitempty"`
	TMSCtctDtTm ISODateTime          `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 TMSCtctDtTm,omitempty"`
}

type Traceability2 struct {
	RlayId      GenericIdentification76 `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 RlayId"`
	TracDtTmIn  ISODateTime             `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 TracDtTmIn"`
	TracDtTmOut ISODateTime             `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 TracDtTmOut"`
}

type TransactionIdentifier1 struct {
	TxDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 TxDtTm"`
	TxRef  Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 TxRef"`
}

type TransactionVerificationResult3 struct {
	Mtd        AuthenticationMethod4Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 Mtd"`
	VrfctnNtty AuthenticationEntity2Code `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 VrfctnNtty,omitempty"`
	Rslt       Verification1Code         `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 Rslt,omitempty"`
	AddtlRslt  Max500Text                `xml:"urn:iso:std:iso:20022:tech:xsd:caaa.002.001.04 AddtlRslt,omitempty"`
}

// May be one of CDSP, CRCP, MDSP, MRCP
type UserInterface1Code string

// May be one of FAIL, MISS, NOVF, PART, SUCC, ERRR
type Verification1Code string

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