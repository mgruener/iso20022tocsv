// Code generated by main. DO NOT EDIT.

package semt_020_001_03

// Must match the pattern [A-Z]{6,6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3,3}){0,1}
type AnyBICIdentifier string

// May be one of FREE, APMT
type DeliveryReceiptType2Code string

type Document struct {
	SctiesMsgCxlAdvc SecuritiesMessageCancellationAdviceV03 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.03 SctiesMsgCxlAdvc"`
}

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

type GenericIdentification19 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.03 Id"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.03 Issr"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.03 SchmeNm,omitempty"`
}

type GenericIdentification20 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.03 Id"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.03 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.03 SchmeNm,omitempty"`
}

// Must be at least 1 items long
type Max350Text string

// Must be at least 1 items long
type Max35Text string

// Must be at least 1 items long
type Max70Text string

type PartyIdentification36Choice struct {
	AnyBIC  AnyBICIdentifier        `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.03 AnyBIC"`
	PrtryId GenericIdentification19 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.03 PrtryId"`
}

// May be one of DELI, RECE
type ReceiveDelivery1Code string

type References18Choice struct {
	SctiesSttlmTxConfId           SettlementTypeAndIdentification13 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.03 SctiesSttlmTxConfId"`
	IntraPosMvmntConfId           Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.03 IntraPosMvmntConfId"`
	SctiesBalAcctgRptId           Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.03 SctiesBalAcctgRptId"`
	SctiesBalCtdyRptId            Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.03 SctiesBalCtdyRptId"`
	IntraPosMvmntPstngRptId       Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.03 IntraPosMvmntPstngRptId"`
	SctiesFincgConfId             SettlementTypeAndIdentification13 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.03 SctiesFincgConfId"`
	SctiesTxPdgRptId              Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.03 SctiesTxPdgRptId"`
	SctiesTxPstngRptId            Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.03 SctiesTxPstngRptId"`
	SctiesSttlmTxAllgmtRptId      Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.03 SctiesSttlmTxAllgmtRptId"`
	SctiesSttlmTxAllgmtNtfctnTxId SettlementTypeAndIdentification13 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.03 SctiesSttlmTxAllgmtNtfctnTxId"`
	PrtflTrfNtfctnId              Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.03 PrtflTrfNtfctnId"`
	SctiesSttlmTxGnrtnNtfctnId    SettlementTypeAndIdentification13 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.03 SctiesSttlmTxGnrtnNtfctnId"`
	OthrMsgId                     Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.03 OthrMsgId"`
}

type SecuritiesAccount13 struct {
	Id Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.03 Id"`
	Tp GenericIdentification20 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.03 Tp,omitempty"`
	Nm Max70Text               `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.03 Nm,omitempty"`
}

type SecuritiesMessageCancellationAdviceV03 struct {
	Ref         References18Choice          `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.03 Ref"`
	AcctOwnr    PartyIdentification36Choice `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.03 AcctOwnr,omitempty"`
	SfkpgAcct   SecuritiesAccount13         `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.03 SfkpgAcct"`
	SplmtryData []SupplementaryData1        `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.03 SplmtryData,omitempty"`
}

type SettlementTypeAndIdentification13 struct {
	TxId          Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.03 TxId"`
	SctiesMvmntTp ReceiveDelivery1Code     `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.03 SctiesMvmntTp"`
	Pmt           DeliveryReceiptType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.03 Pmt"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.03 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:semt.020.001.03 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}