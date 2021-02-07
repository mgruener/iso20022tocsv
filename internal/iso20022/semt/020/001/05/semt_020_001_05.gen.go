// Code generated by main. DO NOT EDIT.

package semt_020_001_05

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type AnyBICIdentifier string

// May be one of FREE, APMT
type DeliveryReceiptType2Code string

type Document struct {
	SctiesMsgCxlAdvc SecuritiesMessageCancellationAdviceV05 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.05 SctiesMsgCxlAdvc"`
}

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

type GenericIdentification30 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.05 Id"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.05 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.05 SchmeNm,omitempty"`
}

type GenericIdentification36 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.05 Id"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.05 Issr"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.05 SchmeNm,omitempty"`
}

// Must match the pattern [A-Z0-9]{18,18}[0-9]{2,2}
type LEIIdentifier string

// Must be at least 1 items long
type Max350Text string

// Must be at least 1 items long
type Max35Text string

// Must be at least 1 items long
type Max70Text string

type PartyIdentification92Choice struct {
	AnyBIC  AnyBICIdentifier        `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.05 AnyBIC"`
	PrtryId GenericIdentification36 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.05 PrtryId"`
}

type PartyIdentification98 struct {
	Id  PartyIdentification92Choice `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.05 Id"`
	LEI LEIIdentifier               `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.05 LEI,omitempty"`
}

// May be one of DELI, RECE
type ReceiveDelivery1Code string

type References43Choice struct {
	SctiesSttlmTxConfId           SettlementTypeAndIdentification18 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.05 SctiesSttlmTxConfId"`
	IntraPosMvmntConfId           Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.05 IntraPosMvmntConfId"`
	SctiesBalAcctgRptId           Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.05 SctiesBalAcctgRptId"`
	SctiesBalCtdyRptId            Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.05 SctiesBalCtdyRptId"`
	IntraPosMvmntPstngRptId       Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.05 IntraPosMvmntPstngRptId"`
	SctiesFincgConfId             SettlementTypeAndIdentification18 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.05 SctiesFincgConfId"`
	SctiesTxPdgRptId              Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.05 SctiesTxPdgRptId"`
	SctiesTxPstngRptId            Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.05 SctiesTxPstngRptId"`
	SctiesSttlmTxAllgmtRptId      Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.05 SctiesSttlmTxAllgmtRptId"`
	SctiesSttlmTxAllgmtNtfctnTxId SettlementTypeAndIdentification18 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.05 SctiesSttlmTxAllgmtNtfctnTxId"`
	PrtflTrfNtfctnId              Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.05 PrtflTrfNtfctnId"`
	SctiesSttlmTxGnrtnNtfctnId    SettlementTypeAndIdentification18 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.05 SctiesSttlmTxGnrtnNtfctnId"`
	OthrMsgId                     Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.05 OthrMsgId"`
	TtlPrtflValtnRptId            Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.05 TtlPrtflValtnRptId"`
}

type SecuritiesAccount24 struct {
	Id Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.05 Id"`
	Tp GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.05 Tp,omitempty"`
	Nm Max70Text               `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.05 Nm,omitempty"`
}

type SecuritiesMessageCancellationAdviceV05 struct {
	Ref         References43Choice    `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.05 Ref"`
	AcctOwnr    PartyIdentification98 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.05 AcctOwnr,omitempty"`
	SfkpgAcct   SecuritiesAccount24   `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.05 SfkpgAcct"`
	SplmtryData []SupplementaryData1  `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.05 SplmtryData,omitempty"`
}

type SettlementTypeAndIdentification18 struct {
	TxId          Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.05 TxId"`
	SctiesMvmntTp ReceiveDelivery1Code     `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.05 SctiesMvmntTp"`
	Pmt           DeliveryReceiptType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.05 Pmt"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.05 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.05 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}
