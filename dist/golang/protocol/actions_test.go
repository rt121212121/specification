package protocol

import (
	"bytes"
	"reflect"
	"testing"
)

func TestAssetDefinition(t *testing.T) {
	// Create a randomized object
	initialMessage := AssetDefinition{}
	// AssetType (fixedchar)
	{
		text := make([]byte, 0, 3)
		for i := uint64(0); i < 3; i++ {
			text = append(text, byte(65+i+0))
		}
		initialMessage.AssetType = string(text)
	}

	// AssetCode (AssetCode)
	initialMessage.AssetCode = AssetCode{}

	// AssetAuthFlags (bin)
	// bin test not setup

	// TransfersPermitted (bool)
	// bool test not setup

	// TradeRestrictions (Polity)
	initialMessage.TradeRestrictions = Polity{}

	// EnforcementOrdersPermitted (bool)
	// bool test not setup

	// VoteMultiplier (uint)
	// uint test not setup

	// ReferendumProposal (bool)
	// bool test not setup

	// InitiativeProposal (bool)
	// bool test not setup

	// AssetModificationGovernance (bool)
	// bool test not setup

	// TokenQty (uint)
	// uint test not setup

	// AssetPayload (varbin)
	initialMessage.AssetPayload = make([]byte, 0, 16)
	for i := uint64(0); i < 16; i++ {
		initialMessage.AssetPayload = append(initialMessage.AssetPayload, byte(65+i+11))
	}

	// Encode message
	initialEncoding, err := initialMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Initial encoding : %d bytes", len(initialEncoding))

	// Decode message
	decodedMessage := AssetDefinition{}

	n, err := decodedMessage.write(initialEncoding)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Decoded : %d bytes", n)

	if n != len(initialEncoding) {
		t.Fatalf("got %v, want %v", n, len(initialEncoding))
	}

	// Serializing the message should give us the same bytes
	secondEncoding, err := decodedMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(initialEncoding, secondEncoding) {
		t.Errorf("got\n%+v\nwant\n%+v", initialEncoding, secondEncoding)
	}

	// if !reflect.DeepEqual(initialMessage, decodedMessage) {
	// 	t.Errorf("\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
	// }

	// Compare re-serialized values
	// AssetType (fixedchar)
	if initialMessage.AssetType != decodedMessage.AssetType {
		t.Errorf("AssetType doesn't match : %s != %s", initialMessage.AssetType, decodedMessage.AssetType)
	}

	// AssetCode (AssetCode)
	if initialMessage.AssetCode != decodedMessage.AssetCode {
		t.Errorf("AssetCode doesn't match : %v != %v", initialMessage.AssetCode, decodedMessage.AssetCode)
	}

	// AssetAuthFlags (bin)
	// bin test compare not setup

	// TransfersPermitted (bool)
	// bool test compare not setup

	// TradeRestrictions (Polity)
	if len(initialMessage.TradeRestrictions.Items) != len(decodedMessage.TradeRestrictions.Items) {
		t.Errorf("TradeRestrictions length doesn't match : %d != %d", initialMessage.TradeRestrictions, decodedMessage.TradeRestrictions)
	}
	for i, value := range initialMessage.TradeRestrictions.Items {
		if !bytes.Equal(value[:], decodedMessage.TradeRestrictions.Items[i][:]) {
			t.Errorf("TradeRestrictions.Items[%d] doesn't match : %s != %s", i, string(value[:]), string(decodedMessage.TradeRestrictions.Items[i][:]))
		}
	}

	// EnforcementOrdersPermitted (bool)
	// bool test compare not setup

	// VoteMultiplier (uint)
	// uint test compare not setup

	// ReferendumProposal (bool)
	// bool test compare not setup

	// InitiativeProposal (bool)
	// bool test compare not setup

	// AssetModificationGovernance (bool)
	// bool test compare not setup

	// TokenQty (uint)
	// uint test compare not setup

	// AssetPayload (varbin)
	if !bytes.Equal(initialMessage.AssetPayload, decodedMessage.AssetPayload) {
		t.Errorf("AssetPayload doesn't match : %x != %x", initialMessage.AssetPayload, decodedMessage.AssetPayload)
	}
}

func TestAssetCreation(t *testing.T) {
	// Create a randomized object
	initialMessage := AssetCreation{}
	// AssetType (fixedchar)
	{
		text := make([]byte, 0, 3)
		for i := uint64(0); i < 3; i++ {
			text = append(text, byte(65+i+0))
		}
		initialMessage.AssetType = string(text)
	}

	// AssetCode (AssetCode)
	initialMessage.AssetCode = AssetCode{}

	// AssetAuthFlags (bin)
	// bin test not setup

	// TransfersPermitted (bool)
	// bool test not setup

	// TradeRestrictions (Polity)
	initialMessage.TradeRestrictions = Polity{}

	// EnforcementOrdersPermitted (bool)
	// bool test not setup

	// VoteMultiplier (uint)
	// uint test not setup

	// ReferendumProposal (bool)
	// bool test not setup

	// InitiativeProposal (bool)
	// bool test not setup

	// AssetModificationGovernance (bool)
	// bool test not setup

	// TokenQty (uint)
	// uint test not setup

	// AssetPayload (varbin)
	initialMessage.AssetPayload = make([]byte, 0, 16)
	for i := uint64(0); i < 16; i++ {
		initialMessage.AssetPayload = append(initialMessage.AssetPayload, byte(65+i+11))
	}

	// AssetRevision (uint)
	// uint test not setup

	// Timestamp (Timestamp)
	initialMessage.Timestamp = Timestamp{}

	// Encode message
	initialEncoding, err := initialMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Initial encoding : %d bytes", len(initialEncoding))

	// Decode message
	decodedMessage := AssetCreation{}

	n, err := decodedMessage.write(initialEncoding)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Decoded : %d bytes", n)

	if n != len(initialEncoding) {
		t.Fatalf("got %v, want %v", n, len(initialEncoding))
	}

	// Serializing the message should give us the same bytes
	secondEncoding, err := decodedMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(initialEncoding, secondEncoding) {
		t.Errorf("got\n%+v\nwant\n%+v", initialEncoding, secondEncoding)
	}

	// if !reflect.DeepEqual(initialMessage, decodedMessage) {
	// 	t.Errorf("\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
	// }

	// Compare re-serialized values
	// AssetType (fixedchar)
	if initialMessage.AssetType != decodedMessage.AssetType {
		t.Errorf("AssetType doesn't match : %s != %s", initialMessage.AssetType, decodedMessage.AssetType)
	}

	// AssetCode (AssetCode)
	if initialMessage.AssetCode != decodedMessage.AssetCode {
		t.Errorf("AssetCode doesn't match : %v != %v", initialMessage.AssetCode, decodedMessage.AssetCode)
	}

	// AssetAuthFlags (bin)
	// bin test compare not setup

	// TransfersPermitted (bool)
	// bool test compare not setup

	// TradeRestrictions (Polity)
	if len(initialMessage.TradeRestrictions.Items) != len(decodedMessage.TradeRestrictions.Items) {
		t.Errorf("TradeRestrictions length doesn't match : %d != %d", initialMessage.TradeRestrictions, decodedMessage.TradeRestrictions)
	}
	for i, value := range initialMessage.TradeRestrictions.Items {
		if !bytes.Equal(value[:], decodedMessage.TradeRestrictions.Items[i][:]) {
			t.Errorf("TradeRestrictions.Items[%d] doesn't match : %s != %s", i, string(value[:]), string(decodedMessage.TradeRestrictions.Items[i][:]))
		}
	}

	// EnforcementOrdersPermitted (bool)
	// bool test compare not setup

	// VoteMultiplier (uint)
	// uint test compare not setup

	// ReferendumProposal (bool)
	// bool test compare not setup

	// InitiativeProposal (bool)
	// bool test compare not setup

	// AssetModificationGovernance (bool)
	// bool test compare not setup

	// TokenQty (uint)
	// uint test compare not setup

	// AssetPayload (varbin)
	if !bytes.Equal(initialMessage.AssetPayload, decodedMessage.AssetPayload) {
		t.Errorf("AssetPayload doesn't match : %x != %x", initialMessage.AssetPayload, decodedMessage.AssetPayload)
	}

	// AssetRevision (uint)
	// uint test compare not setup

	// Timestamp (Timestamp)
	if initialMessage.Timestamp != decodedMessage.Timestamp {
		t.Errorf("Timestamp doesn't match : %v != %v", initialMessage.Timestamp, decodedMessage.Timestamp)
	}
}

func TestAssetModification(t *testing.T) {
	// Create a randomized object
	initialMessage := AssetModification{}
	// AssetType (fixedchar)
	{
		text := make([]byte, 0, 3)
		for i := uint64(0); i < 3; i++ {
			text = append(text, byte(65+i+0))
		}
		initialMessage.AssetType = string(text)
	}

	// AssetCode (AssetCode)
	initialMessage.AssetCode = AssetCode{}

	// AssetRevision (uint)
	// uint test not setup

	// Modifications (Amendment[])
	for i := 0; i < 2; i++ {
		initialMessage.Modifications = append(initialMessage.Modifications, Amendment{})
	}

	// RefTxID (TxId)
	initialMessage.RefTxID = TxId{}

	// Encode message
	initialEncoding, err := initialMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Initial encoding : %d bytes", len(initialEncoding))

	// Decode message
	decodedMessage := AssetModification{}

	n, err := decodedMessage.write(initialEncoding)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Decoded : %d bytes", n)

	if n != len(initialEncoding) {
		t.Fatalf("got %v, want %v", n, len(initialEncoding))
	}

	// Serializing the message should give us the same bytes
	secondEncoding, err := decodedMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(initialEncoding, secondEncoding) {
		t.Errorf("got\n%+v\nwant\n%+v", initialEncoding, secondEncoding)
	}

	// if !reflect.DeepEqual(initialMessage, decodedMessage) {
	// 	t.Errorf("\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
	// }

	// Compare re-serialized values
	// AssetType (fixedchar)
	if initialMessage.AssetType != decodedMessage.AssetType {
		t.Errorf("AssetType doesn't match : %s != %s", initialMessage.AssetType, decodedMessage.AssetType)
	}

	// AssetCode (AssetCode)
	if initialMessage.AssetCode != decodedMessage.AssetCode {
		t.Errorf("AssetCode doesn't match : %v != %v", initialMessage.AssetCode, decodedMessage.AssetCode)
	}

	// AssetRevision (uint)
	// uint test compare not setup

	// Modifications (Amendment[])
	if len(initialMessage.Modifications) != len(decodedMessage.Modifications) {
		t.Errorf("Modifications lengths don't match : %d != %d", len(initialMessage.Modifications), len(decodedMessage.Modifications))
	}

	// RefTxID (TxId)
	if initialMessage.RefTxID != decodedMessage.RefTxID {
		t.Errorf("RefTxID doesn't match : %v != %v", initialMessage.RefTxID, decodedMessage.RefTxID)
	}
}

func TestContractOffer(t *testing.T) {
	// Create a randomized object
	initialMessage := ContractOffer{}
	// ContractName (varchar)
	initialMessage.ContractName = "Text 0"

	// ContractFileType (uint)
	// uint test not setup

	// ContractFile (varbin)
	initialMessage.ContractFile = make([]byte, 0, 32)
	for i := uint64(0); i < 32; i++ {
		initialMessage.ContractFile = append(initialMessage.ContractFile, byte(65+i+2))
	}

	// SupportingDocsFileType (uint)
	// uint test not setup

	// SupportingDocs (varchar)
	initialMessage.SupportingDocs = "Text 4"

	// GoverningLaw (fixedchar)
	{
		text := make([]byte, 0, 5)
		for i := uint64(0); i < 5; i++ {
			text = append(text, byte(65+i+5))
		}
		initialMessage.GoverningLaw = string(text)
	}

	// Jurisdiction (fixedchar)
	{
		text := make([]byte, 0, 5)
		for i := uint64(0); i < 5; i++ {
			text = append(text, byte(65+i+6))
		}
		initialMessage.Jurisdiction = string(text)
	}

	// ContractExpiration (Timestamp)
	initialMessage.ContractExpiration = Timestamp{}

	// ContractURI (varchar)
	initialMessage.ContractURI = "Text 8"

	// IssuerName (varchar)
	initialMessage.IssuerName = "Text 9"

	// IssuerType (fixedchar)
	// fixedchar test not setup

	// IssuerLEI (fixedchar)
	{
		text := make([]byte, 0, 20)
		for i := uint64(0); i < 20; i++ {
			text = append(text, byte(65+i+11))
		}
		initialMessage.IssuerLEI = string(text)
	}

	// IssuerLogoURL (varchar)
	initialMessage.IssuerLogoURL = "Text 12"

	// ContractOperatorIncluded (bool)
	// bool test not setup

	// ContractOperatorID (varchar)
	initialMessage.ContractOperatorID = "Text 14"

	// OperatorLEI (fixedchar)
	{
		text := make([]byte, 0, 20)
		for i := uint64(0); i < 20; i++ {
			text = append(text, byte(65+i+15))
		}
		initialMessage.OperatorLEI = string(text)
	}

	// ContractAuthFlags (bin)
	// bin test not setup

	// ActionFee (Fee[])
	for i := 0; i < 2; i++ {
		initialMessage.ActionFee = append(initialMessage.ActionFee, Fee{})
	}

	// VotingSystems (VotingSystem[])
	for i := 0; i < 2; i++ {
		initialMessage.VotingSystems = append(initialMessage.VotingSystems, VotingSystem{})
	}

	// RestrictedQtyAssets (uint)
	// uint test not setup

	// ReferendumProposal (bool)
	// bool test not setup

	// InitiativeProposal (bool)
	// bool test not setup

	// Registries (Registry[])
	for i := 0; i < 2; i++ {
		initialMessage.Registries = append(initialMessage.Registries, Registry{})
	}

	// IssuerAddress (bool)
	// bool test not setup

	// UnitNumber (varchar)
	initialMessage.UnitNumber = "Text 24"

	// BuildingNumber (varchar)
	initialMessage.BuildingNumber = "Text 25"

	// Street (varchar)
	initialMessage.Street = "Text 26"

	// SuburbCity (varchar)
	initialMessage.SuburbCity = "Text 27"

	// TerritoryStateProvinceCode (fixedchar)
	{
		text := make([]byte, 0, 5)
		for i := uint64(0); i < 5; i++ {
			text = append(text, byte(65+i+28))
		}
		initialMessage.TerritoryStateProvinceCode = string(text)
	}

	// CountryCode (fixedchar)
	{
		text := make([]byte, 0, 3)
		for i := uint64(0); i < 3; i++ {
			text = append(text, byte(65+i+29))
		}
		initialMessage.CountryCode = string(text)
	}

	// PostalZIPCode (varchar)
	initialMessage.PostalZIPCode = "Text 30"

	// EmailAddress (varchar)
	initialMessage.EmailAddress = "Text 31"

	// PhoneNumber (varchar)
	initialMessage.PhoneNumber = "Text 32"

	// KeyRoles (KeyRole[])
	for i := 0; i < 2; i++ {
		initialMessage.KeyRoles = append(initialMessage.KeyRoles, KeyRole{})
	}

	// NotableRoles (NotableRole[])
	for i := 0; i < 2; i++ {
		initialMessage.NotableRoles = append(initialMessage.NotableRoles, NotableRole{})
	}

	// Encode message
	initialEncoding, err := initialMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Initial encoding : %d bytes", len(initialEncoding))

	// Decode message
	decodedMessage := ContractOffer{}

	n, err := decodedMessage.write(initialEncoding)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Decoded : %d bytes", n)

	if n != len(initialEncoding) {
		t.Fatalf("got %v, want %v", n, len(initialEncoding))
	}

	// Serializing the message should give us the same bytes
	secondEncoding, err := decodedMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(initialEncoding, secondEncoding) {
		t.Errorf("got\n%+v\nwant\n%+v", initialEncoding, secondEncoding)
	}

	// if !reflect.DeepEqual(initialMessage, decodedMessage) {
	// 	t.Errorf("\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
	// }

	// Compare re-serialized values
	// ContractName (varchar)
	if initialMessage.ContractName != decodedMessage.ContractName {
		t.Errorf("ContractName doesn't match : %s != %s", initialMessage.ContractName, decodedMessage.ContractName)
	}

	// ContractFileType (uint)
	// uint test compare not setup

	// ContractFile (varbin)
	if !bytes.Equal(initialMessage.ContractFile, decodedMessage.ContractFile) {
		t.Errorf("ContractFile doesn't match : %x != %x", initialMessage.ContractFile, decodedMessage.ContractFile)
	}

	// SupportingDocsFileType (uint)
	// uint test compare not setup

	// SupportingDocs (varchar)
	if initialMessage.SupportingDocs != decodedMessage.SupportingDocs {
		t.Errorf("SupportingDocs doesn't match : %s != %s", initialMessage.SupportingDocs, decodedMessage.SupportingDocs)
	}

	// GoverningLaw (fixedchar)
	if initialMessage.GoverningLaw != decodedMessage.GoverningLaw {
		t.Errorf("GoverningLaw doesn't match : %s != %s", initialMessage.GoverningLaw, decodedMessage.GoverningLaw)
	}

	// Jurisdiction (fixedchar)
	if initialMessage.Jurisdiction != decodedMessage.Jurisdiction {
		t.Errorf("Jurisdiction doesn't match : %s != %s", initialMessage.Jurisdiction, decodedMessage.Jurisdiction)
	}

	// ContractExpiration (Timestamp)
	if initialMessage.ContractExpiration != decodedMessage.ContractExpiration {
		t.Errorf("ContractExpiration doesn't match : %v != %v", initialMessage.ContractExpiration, decodedMessage.ContractExpiration)
	}

	// ContractURI (varchar)
	if initialMessage.ContractURI != decodedMessage.ContractURI {
		t.Errorf("ContractURI doesn't match : %s != %s", initialMessage.ContractURI, decodedMessage.ContractURI)
	}

	// IssuerName (varchar)
	if initialMessage.IssuerName != decodedMessage.IssuerName {
		t.Errorf("IssuerName doesn't match : %s != %s", initialMessage.IssuerName, decodedMessage.IssuerName)
	}

	// IssuerType (fixedchar)
	// fixedchar test compare not setup

	// IssuerLEI (fixedchar)
	if initialMessage.IssuerLEI != decodedMessage.IssuerLEI {
		t.Errorf("IssuerLEI doesn't match : %s != %s", initialMessage.IssuerLEI, decodedMessage.IssuerLEI)
	}

	// IssuerLogoURL (varchar)
	if initialMessage.IssuerLogoURL != decodedMessage.IssuerLogoURL {
		t.Errorf("IssuerLogoURL doesn't match : %s != %s", initialMessage.IssuerLogoURL, decodedMessage.IssuerLogoURL)
	}

	// ContractOperatorIncluded (bool)
	// bool test compare not setup

	// ContractOperatorID (varchar)
	if initialMessage.ContractOperatorID != decodedMessage.ContractOperatorID {
		t.Errorf("ContractOperatorID doesn't match : %s != %s", initialMessage.ContractOperatorID, decodedMessage.ContractOperatorID)
	}

	// OperatorLEI (fixedchar)
	if initialMessage.OperatorLEI != decodedMessage.OperatorLEI {
		t.Errorf("OperatorLEI doesn't match : %s != %s", initialMessage.OperatorLEI, decodedMessage.OperatorLEI)
	}

	// ContractAuthFlags (bin)
	// bin test compare not setup

	// ActionFee (Fee[])
	if len(initialMessage.ActionFee) != len(decodedMessage.ActionFee) {
		t.Errorf("ActionFee lengths don't match : %d != %d", len(initialMessage.ActionFee), len(decodedMessage.ActionFee))
	}

	// VotingSystems (VotingSystem[])
	if len(initialMessage.VotingSystems) != len(decodedMessage.VotingSystems) {
		t.Errorf("VotingSystems lengths don't match : %d != %d", len(initialMessage.VotingSystems), len(decodedMessage.VotingSystems))
	}

	// RestrictedQtyAssets (uint)
	// uint test compare not setup

	// ReferendumProposal (bool)
	// bool test compare not setup

	// InitiativeProposal (bool)
	// bool test compare not setup

	// Registries (Registry[])
	if len(initialMessage.Registries) != len(decodedMessage.Registries) {
		t.Errorf("Registries lengths don't match : %d != %d", len(initialMessage.Registries), len(decodedMessage.Registries))
	}

	// IssuerAddress (bool)
	// bool test compare not setup

	// UnitNumber (varchar)
	if initialMessage.UnitNumber != decodedMessage.UnitNumber {
		t.Errorf("UnitNumber doesn't match : %s != %s", initialMessage.UnitNumber, decodedMessage.UnitNumber)
	}

	// BuildingNumber (varchar)
	if initialMessage.BuildingNumber != decodedMessage.BuildingNumber {
		t.Errorf("BuildingNumber doesn't match : %s != %s", initialMessage.BuildingNumber, decodedMessage.BuildingNumber)
	}

	// Street (varchar)
	if initialMessage.Street != decodedMessage.Street {
		t.Errorf("Street doesn't match : %s != %s", initialMessage.Street, decodedMessage.Street)
	}

	// SuburbCity (varchar)
	if initialMessage.SuburbCity != decodedMessage.SuburbCity {
		t.Errorf("SuburbCity doesn't match : %s != %s", initialMessage.SuburbCity, decodedMessage.SuburbCity)
	}

	// TerritoryStateProvinceCode (fixedchar)
	if initialMessage.TerritoryStateProvinceCode != decodedMessage.TerritoryStateProvinceCode {
		t.Errorf("TerritoryStateProvinceCode doesn't match : %s != %s", initialMessage.TerritoryStateProvinceCode, decodedMessage.TerritoryStateProvinceCode)
	}

	// CountryCode (fixedchar)
	if initialMessage.CountryCode != decodedMessage.CountryCode {
		t.Errorf("CountryCode doesn't match : %s != %s", initialMessage.CountryCode, decodedMessage.CountryCode)
	}

	// PostalZIPCode (varchar)
	if initialMessage.PostalZIPCode != decodedMessage.PostalZIPCode {
		t.Errorf("PostalZIPCode doesn't match : %s != %s", initialMessage.PostalZIPCode, decodedMessage.PostalZIPCode)
	}

	// EmailAddress (varchar)
	if initialMessage.EmailAddress != decodedMessage.EmailAddress {
		t.Errorf("EmailAddress doesn't match : %s != %s", initialMessage.EmailAddress, decodedMessage.EmailAddress)
	}

	// PhoneNumber (varchar)
	if initialMessage.PhoneNumber != decodedMessage.PhoneNumber {
		t.Errorf("PhoneNumber doesn't match : %s != %s", initialMessage.PhoneNumber, decodedMessage.PhoneNumber)
	}

	// KeyRoles (KeyRole[])
	if len(initialMessage.KeyRoles) != len(decodedMessage.KeyRoles) {
		t.Errorf("KeyRoles lengths don't match : %d != %d", len(initialMessage.KeyRoles), len(decodedMessage.KeyRoles))
	}

	// NotableRoles (NotableRole[])
	if len(initialMessage.NotableRoles) != len(decodedMessage.NotableRoles) {
		t.Errorf("NotableRoles lengths don't match : %d != %d", len(initialMessage.NotableRoles), len(decodedMessage.NotableRoles))
	}
}

func TestContractFormation(t *testing.T) {
	// Create a randomized object
	initialMessage := ContractFormation{}
	// ContractName (varchar)
	initialMessage.ContractName = "Text 0"

	// ContractFileType (uint)
	// uint test not setup

	// ContractFile (varbin)
	initialMessage.ContractFile = make([]byte, 0, 32)
	for i := uint64(0); i < 32; i++ {
		initialMessage.ContractFile = append(initialMessage.ContractFile, byte(65+i+2))
	}

	// SupportingDocsFileType (uint)
	// uint test not setup

	// SupportingDocs (varchar)
	initialMessage.SupportingDocs = "Text 4"

	// GoverningLaw (fixedchar)
	{
		text := make([]byte, 0, 5)
		for i := uint64(0); i < 5; i++ {
			text = append(text, byte(65+i+5))
		}
		initialMessage.GoverningLaw = string(text)
	}

	// Jurisdiction (fixedchar)
	{
		text := make([]byte, 0, 5)
		for i := uint64(0); i < 5; i++ {
			text = append(text, byte(65+i+6))
		}
		initialMessage.Jurisdiction = string(text)
	}

	// ContractExpiration (Timestamp)
	initialMessage.ContractExpiration = Timestamp{}

	// ContractURI (varchar)
	initialMessage.ContractURI = "Text 8"

	// IssuerName (varchar)
	initialMessage.IssuerName = "Text 9"

	// IssuerType (fixedchar)
	// fixedchar test not setup

	// IssuerLEI (fixedchar)
	{
		text := make([]byte, 0, 20)
		for i := uint64(0); i < 20; i++ {
			text = append(text, byte(65+i+11))
		}
		initialMessage.IssuerLEI = string(text)
	}

	// IssuerLogoURL (varchar)
	initialMessage.IssuerLogoURL = "Text 12"

	// ContractOperatorID (varchar)
	initialMessage.ContractOperatorID = "Text 13"

	// OperatorLEI (fixedchar)
	{
		text := make([]byte, 0, 20)
		for i := uint64(0); i < 20; i++ {
			text = append(text, byte(65+i+14))
		}
		initialMessage.OperatorLEI = string(text)
	}

	// ContractAuthFlags (bin)
	// bin test not setup

	// ContractFees (Fee[])
	for i := 0; i < 2; i++ {
		initialMessage.ContractFees = append(initialMessage.ContractFees, Fee{})
	}

	// VotingSystems (VotingSystem[])
	for i := 0; i < 2; i++ {
		initialMessage.VotingSystems = append(initialMessage.VotingSystems, VotingSystem{})
	}

	// RestrictedQtyAssets (uint)
	// uint test not setup

	// ReferendumProposal (bool)
	// bool test not setup

	// InitiativeProposal (bool)
	// bool test not setup

	// Registries (Registry[])
	for i := 0; i < 2; i++ {
		initialMessage.Registries = append(initialMessage.Registries, Registry{})
	}

	// IssuerAddress (bool)
	// bool test not setup

	// UnitNumber (varchar)
	initialMessage.UnitNumber = "Text 23"

	// BuildingNumber (varchar)
	initialMessage.BuildingNumber = "Text 24"

	// Street (varchar)
	initialMessage.Street = "Text 25"

	// SuburbCity (varchar)
	initialMessage.SuburbCity = "Text 26"

	// TerritoryStateProvinceCode (fixedchar)
	{
		text := make([]byte, 0, 5)
		for i := uint64(0); i < 5; i++ {
			text = append(text, byte(65+i+27))
		}
		initialMessage.TerritoryStateProvinceCode = string(text)
	}

	// CountryCode (fixedchar)
	{
		text := make([]byte, 0, 3)
		for i := uint64(0); i < 3; i++ {
			text = append(text, byte(65+i+28))
		}
		initialMessage.CountryCode = string(text)
	}

	// PostalZIPCode (varchar)
	initialMessage.PostalZIPCode = "Text 29"

	// EmailAddress (varchar)
	initialMessage.EmailAddress = "Text 30"

	// PhoneNumber (varchar)
	initialMessage.PhoneNumber = "Text 31"

	// KeyRoles (KeyRole[])
	for i := 0; i < 2; i++ {
		initialMessage.KeyRoles = append(initialMessage.KeyRoles, KeyRole{})
	}

	// NotableRoles (NotableRole[])
	for i := 0; i < 2; i++ {
		initialMessage.NotableRoles = append(initialMessage.NotableRoles, NotableRole{})
	}

	// ContractRevision (uint)
	// uint test not setup

	// Timestamp (Timestamp)
	initialMessage.Timestamp = Timestamp{}

	// Encode message
	initialEncoding, err := initialMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Initial encoding : %d bytes", len(initialEncoding))

	// Decode message
	decodedMessage := ContractFormation{}

	n, err := decodedMessage.write(initialEncoding)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Decoded : %d bytes", n)

	if n != len(initialEncoding) {
		t.Fatalf("got %v, want %v", n, len(initialEncoding))
	}

	// Serializing the message should give us the same bytes
	secondEncoding, err := decodedMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(initialEncoding, secondEncoding) {
		t.Errorf("got\n%+v\nwant\n%+v", initialEncoding, secondEncoding)
	}

	// if !reflect.DeepEqual(initialMessage, decodedMessage) {
	// 	t.Errorf("\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
	// }

	// Compare re-serialized values
	// ContractName (varchar)
	if initialMessage.ContractName != decodedMessage.ContractName {
		t.Errorf("ContractName doesn't match : %s != %s", initialMessage.ContractName, decodedMessage.ContractName)
	}

	// ContractFileType (uint)
	// uint test compare not setup

	// ContractFile (varbin)
	if !bytes.Equal(initialMessage.ContractFile, decodedMessage.ContractFile) {
		t.Errorf("ContractFile doesn't match : %x != %x", initialMessage.ContractFile, decodedMessage.ContractFile)
	}

	// SupportingDocsFileType (uint)
	// uint test compare not setup

	// SupportingDocs (varchar)
	if initialMessage.SupportingDocs != decodedMessage.SupportingDocs {
		t.Errorf("SupportingDocs doesn't match : %s != %s", initialMessage.SupportingDocs, decodedMessage.SupportingDocs)
	}

	// GoverningLaw (fixedchar)
	if initialMessage.GoverningLaw != decodedMessage.GoverningLaw {
		t.Errorf("GoverningLaw doesn't match : %s != %s", initialMessage.GoverningLaw, decodedMessage.GoverningLaw)
	}

	// Jurisdiction (fixedchar)
	if initialMessage.Jurisdiction != decodedMessage.Jurisdiction {
		t.Errorf("Jurisdiction doesn't match : %s != %s", initialMessage.Jurisdiction, decodedMessage.Jurisdiction)
	}

	// ContractExpiration (Timestamp)
	if initialMessage.ContractExpiration != decodedMessage.ContractExpiration {
		t.Errorf("ContractExpiration doesn't match : %v != %v", initialMessage.ContractExpiration, decodedMessage.ContractExpiration)
	}

	// ContractURI (varchar)
	if initialMessage.ContractURI != decodedMessage.ContractURI {
		t.Errorf("ContractURI doesn't match : %s != %s", initialMessage.ContractURI, decodedMessage.ContractURI)
	}

	// IssuerName (varchar)
	if initialMessage.IssuerName != decodedMessage.IssuerName {
		t.Errorf("IssuerName doesn't match : %s != %s", initialMessage.IssuerName, decodedMessage.IssuerName)
	}

	// IssuerType (fixedchar)
	// fixedchar test compare not setup

	// IssuerLEI (fixedchar)
	if initialMessage.IssuerLEI != decodedMessage.IssuerLEI {
		t.Errorf("IssuerLEI doesn't match : %s != %s", initialMessage.IssuerLEI, decodedMessage.IssuerLEI)
	}

	// IssuerLogoURL (varchar)
	if initialMessage.IssuerLogoURL != decodedMessage.IssuerLogoURL {
		t.Errorf("IssuerLogoURL doesn't match : %s != %s", initialMessage.IssuerLogoURL, decodedMessage.IssuerLogoURL)
	}

	// ContractOperatorID (varchar)
	if initialMessage.ContractOperatorID != decodedMessage.ContractOperatorID {
		t.Errorf("ContractOperatorID doesn't match : %s != %s", initialMessage.ContractOperatorID, decodedMessage.ContractOperatorID)
	}

	// OperatorLEI (fixedchar)
	if initialMessage.OperatorLEI != decodedMessage.OperatorLEI {
		t.Errorf("OperatorLEI doesn't match : %s != %s", initialMessage.OperatorLEI, decodedMessage.OperatorLEI)
	}

	// ContractAuthFlags (bin)
	// bin test compare not setup

	// ContractFees (Fee[])
	if len(initialMessage.ContractFees) != len(decodedMessage.ContractFees) {
		t.Errorf("ContractFees lengths don't match : %d != %d", len(initialMessage.ContractFees), len(decodedMessage.ContractFees))
	}

	// VotingSystems (VotingSystem[])
	if len(initialMessage.VotingSystems) != len(decodedMessage.VotingSystems) {
		t.Errorf("VotingSystems lengths don't match : %d != %d", len(initialMessage.VotingSystems), len(decodedMessage.VotingSystems))
	}

	// RestrictedQtyAssets (uint)
	// uint test compare not setup

	// ReferendumProposal (bool)
	// bool test compare not setup

	// InitiativeProposal (bool)
	// bool test compare not setup

	// Registries (Registry[])
	if len(initialMessage.Registries) != len(decodedMessage.Registries) {
		t.Errorf("Registries lengths don't match : %d != %d", len(initialMessage.Registries), len(decodedMessage.Registries))
	}

	// IssuerAddress (bool)
	// bool test compare not setup

	// UnitNumber (varchar)
	if initialMessage.UnitNumber != decodedMessage.UnitNumber {
		t.Errorf("UnitNumber doesn't match : %s != %s", initialMessage.UnitNumber, decodedMessage.UnitNumber)
	}

	// BuildingNumber (varchar)
	if initialMessage.BuildingNumber != decodedMessage.BuildingNumber {
		t.Errorf("BuildingNumber doesn't match : %s != %s", initialMessage.BuildingNumber, decodedMessage.BuildingNumber)
	}

	// Street (varchar)
	if initialMessage.Street != decodedMessage.Street {
		t.Errorf("Street doesn't match : %s != %s", initialMessage.Street, decodedMessage.Street)
	}

	// SuburbCity (varchar)
	if initialMessage.SuburbCity != decodedMessage.SuburbCity {
		t.Errorf("SuburbCity doesn't match : %s != %s", initialMessage.SuburbCity, decodedMessage.SuburbCity)
	}

	// TerritoryStateProvinceCode (fixedchar)
	if initialMessage.TerritoryStateProvinceCode != decodedMessage.TerritoryStateProvinceCode {
		t.Errorf("TerritoryStateProvinceCode doesn't match : %s != %s", initialMessage.TerritoryStateProvinceCode, decodedMessage.TerritoryStateProvinceCode)
	}

	// CountryCode (fixedchar)
	if initialMessage.CountryCode != decodedMessage.CountryCode {
		t.Errorf("CountryCode doesn't match : %s != %s", initialMessage.CountryCode, decodedMessage.CountryCode)
	}

	// PostalZIPCode (varchar)
	if initialMessage.PostalZIPCode != decodedMessage.PostalZIPCode {
		t.Errorf("PostalZIPCode doesn't match : %s != %s", initialMessage.PostalZIPCode, decodedMessage.PostalZIPCode)
	}

	// EmailAddress (varchar)
	if initialMessage.EmailAddress != decodedMessage.EmailAddress {
		t.Errorf("EmailAddress doesn't match : %s != %s", initialMessage.EmailAddress, decodedMessage.EmailAddress)
	}

	// PhoneNumber (varchar)
	if initialMessage.PhoneNumber != decodedMessage.PhoneNumber {
		t.Errorf("PhoneNumber doesn't match : %s != %s", initialMessage.PhoneNumber, decodedMessage.PhoneNumber)
	}

	// KeyRoles (KeyRole[])
	if len(initialMessage.KeyRoles) != len(decodedMessage.KeyRoles) {
		t.Errorf("KeyRoles lengths don't match : %d != %d", len(initialMessage.KeyRoles), len(decodedMessage.KeyRoles))
	}

	// NotableRoles (NotableRole[])
	if len(initialMessage.NotableRoles) != len(decodedMessage.NotableRoles) {
		t.Errorf("NotableRoles lengths don't match : %d != %d", len(initialMessage.NotableRoles), len(decodedMessage.NotableRoles))
	}

	// ContractRevision (uint)
	// uint test compare not setup

	// Timestamp (Timestamp)
	if initialMessage.Timestamp != decodedMessage.Timestamp {
		t.Errorf("Timestamp doesn't match : %v != %v", initialMessage.Timestamp, decodedMessage.Timestamp)
	}
}

func TestContractAmendment(t *testing.T) {
	// Create a randomized object
	initialMessage := ContractAmendment{}
	// ChangeIssuerAddress (bool)
	// bool test not setup

	// ChangeOperatorAddress (bool)
	// bool test not setup

	// ContractRevision (uint)
	// uint test not setup

	// Amendments (Amendment[])
	for i := 0; i < 2; i++ {
		initialMessage.Amendments = append(initialMessage.Amendments, Amendment{})
	}

	// RefTxID (TxId)
	initialMessage.RefTxID = TxId{}

	// Encode message
	initialEncoding, err := initialMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Initial encoding : %d bytes", len(initialEncoding))

	// Decode message
	decodedMessage := ContractAmendment{}

	n, err := decodedMessage.write(initialEncoding)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Decoded : %d bytes", n)

	if n != len(initialEncoding) {
		t.Fatalf("got %v, want %v", n, len(initialEncoding))
	}

	// Serializing the message should give us the same bytes
	secondEncoding, err := decodedMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(initialEncoding, secondEncoding) {
		t.Errorf("got\n%+v\nwant\n%+v", initialEncoding, secondEncoding)
	}

	// if !reflect.DeepEqual(initialMessage, decodedMessage) {
	// 	t.Errorf("\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
	// }

	// Compare re-serialized values
	// ChangeIssuerAddress (bool)
	// bool test compare not setup

	// ChangeOperatorAddress (bool)
	// bool test compare not setup

	// ContractRevision (uint)
	// uint test compare not setup

	// Amendments (Amendment[])
	if len(initialMessage.Amendments) != len(decodedMessage.Amendments) {
		t.Errorf("Amendments lengths don't match : %d != %d", len(initialMessage.Amendments), len(decodedMessage.Amendments))
	}

	// RefTxID (TxId)
	if initialMessage.RefTxID != decodedMessage.RefTxID {
		t.Errorf("RefTxID doesn't match : %v != %v", initialMessage.RefTxID, decodedMessage.RefTxID)
	}
}

func TestStaticContractFormation(t *testing.T) {
	// Create a randomized object
	initialMessage := StaticContractFormation{}
	// ContractName (varchar)
	initialMessage.ContractName = "Text 0"

	// ContractType (varchar)
	initialMessage.ContractType = "Text 1"

	// ContractCode (ContractCode)
	initialMessage.ContractCode = ContractCode{}

	// ContractFileType (uint)
	// uint test not setup

	// ContractFile (varbin)
	initialMessage.ContractFile = make([]byte, 0, 32)
	for i := uint64(0); i < 32; i++ {
		initialMessage.ContractFile = append(initialMessage.ContractFile, byte(65+i+4))
	}

	// SupportingDocsFileType (uint)
	// uint test not setup

	// SupportingDocs (varchar)
	initialMessage.SupportingDocs = "Text 6"

	// ContractRevision (uint)
	// uint test not setup

	// GoverningLaw (fixedchar)
	{
		text := make([]byte, 0, 5)
		for i := uint64(0); i < 5; i++ {
			text = append(text, byte(65+i+8))
		}
		initialMessage.GoverningLaw = string(text)
	}

	// Jurisdiction (fixedchar)
	{
		text := make([]byte, 0, 5)
		for i := uint64(0); i < 5; i++ {
			text = append(text, byte(65+i+9))
		}
		initialMessage.Jurisdiction = string(text)
	}

	// EffectiveDate (Timestamp)
	initialMessage.EffectiveDate = Timestamp{}

	// ContractExpiration (Timestamp)
	initialMessage.ContractExpiration = Timestamp{}

	// ContractURI (varchar)
	initialMessage.ContractURI = "Text 12"

	// PrevRevTxID (TxId)
	initialMessage.PrevRevTxID = TxId{}

	// Entities (Entity[])
	for i := 0; i < 2; i++ {
		initialMessage.Entities = append(initialMessage.Entities, Entity{})
	}

	// Encode message
	initialEncoding, err := initialMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Initial encoding : %d bytes", len(initialEncoding))

	// Decode message
	decodedMessage := StaticContractFormation{}

	n, err := decodedMessage.write(initialEncoding)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Decoded : %d bytes", n)

	if n != len(initialEncoding) {
		t.Fatalf("got %v, want %v", n, len(initialEncoding))
	}

	// Serializing the message should give us the same bytes
	secondEncoding, err := decodedMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(initialEncoding, secondEncoding) {
		t.Errorf("got\n%+v\nwant\n%+v", initialEncoding, secondEncoding)
	}

	// if !reflect.DeepEqual(initialMessage, decodedMessage) {
	// 	t.Errorf("\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
	// }

	// Compare re-serialized values
	// ContractName (varchar)
	if initialMessage.ContractName != decodedMessage.ContractName {
		t.Errorf("ContractName doesn't match : %s != %s", initialMessage.ContractName, decodedMessage.ContractName)
	}

	// ContractType (varchar)
	if initialMessage.ContractType != decodedMessage.ContractType {
		t.Errorf("ContractType doesn't match : %s != %s", initialMessage.ContractType, decodedMessage.ContractType)
	}

	// ContractCode (ContractCode)
	if initialMessage.ContractCode != decodedMessage.ContractCode {
		t.Errorf("ContractCode doesn't match : %v != %v", initialMessage.ContractCode, decodedMessage.ContractCode)
	}

	// ContractFileType (uint)
	// uint test compare not setup

	// ContractFile (varbin)
	if !bytes.Equal(initialMessage.ContractFile, decodedMessage.ContractFile) {
		t.Errorf("ContractFile doesn't match : %x != %x", initialMessage.ContractFile, decodedMessage.ContractFile)
	}

	// SupportingDocsFileType (uint)
	// uint test compare not setup

	// SupportingDocs (varchar)
	if initialMessage.SupportingDocs != decodedMessage.SupportingDocs {
		t.Errorf("SupportingDocs doesn't match : %s != %s", initialMessage.SupportingDocs, decodedMessage.SupportingDocs)
	}

	// ContractRevision (uint)
	// uint test compare not setup

	// GoverningLaw (fixedchar)
	if initialMessage.GoverningLaw != decodedMessage.GoverningLaw {
		t.Errorf("GoverningLaw doesn't match : %s != %s", initialMessage.GoverningLaw, decodedMessage.GoverningLaw)
	}

	// Jurisdiction (fixedchar)
	if initialMessage.Jurisdiction != decodedMessage.Jurisdiction {
		t.Errorf("Jurisdiction doesn't match : %s != %s", initialMessage.Jurisdiction, decodedMessage.Jurisdiction)
	}

	// EffectiveDate (Timestamp)
	if initialMessage.EffectiveDate != decodedMessage.EffectiveDate {
		t.Errorf("EffectiveDate doesn't match : %v != %v", initialMessage.EffectiveDate, decodedMessage.EffectiveDate)
	}

	// ContractExpiration (Timestamp)
	if initialMessage.ContractExpiration != decodedMessage.ContractExpiration {
		t.Errorf("ContractExpiration doesn't match : %v != %v", initialMessage.ContractExpiration, decodedMessage.ContractExpiration)
	}

	// ContractURI (varchar)
	if initialMessage.ContractURI != decodedMessage.ContractURI {
		t.Errorf("ContractURI doesn't match : %s != %s", initialMessage.ContractURI, decodedMessage.ContractURI)
	}

	// PrevRevTxID (TxId)
	if initialMessage.PrevRevTxID != decodedMessage.PrevRevTxID {
		t.Errorf("PrevRevTxID doesn't match : %v != %v", initialMessage.PrevRevTxID, decodedMessage.PrevRevTxID)
	}

	// Entities (Entity[])
	if len(initialMessage.Entities) != len(decodedMessage.Entities) {
		t.Errorf("Entities lengths don't match : %d != %d", len(initialMessage.Entities), len(decodedMessage.Entities))
	}
}

func TestOrder(t *testing.T) {
	// Create a randomized object
	initialMessage := Order{}
	// AssetType (fixedchar)
	{
		text := make([]byte, 0, 3)
		for i := uint64(0); i < 3; i++ {
			text = append(text, byte(65+i+0))
		}
		initialMessage.AssetType = string(text)
	}

	// AssetCode (AssetCode)
	initialMessage.AssetCode = AssetCode{}

	// ComplianceAction (fixedchar)
	// fixedchar test not setup

	// TargetAddresses (TargetAddress[])
	for i := 0; i < 2; i++ {
		initialMessage.TargetAddresses = append(initialMessage.TargetAddresses, TargetAddress{})
	}

	// DepositAddress (PublicKeyHash)
	initialMessage.DepositAddress = PublicKeyHash{}

	// AuthorityName (varchar)
	initialMessage.AuthorityName = "Text 5"

	// SigAlgoAddressList (uint)
	// uint test not setup

	// AuthorityPublicKey (varchar)
	initialMessage.AuthorityPublicKey = "Text 7"

	// OrderSignature (varchar)
	initialMessage.OrderSignature = "Text 8"

	// SupportingEvidenceTxId (TxId)
	initialMessage.SupportingEvidenceTxId = TxId{}

	// RefTxnID (TxId)
	initialMessage.RefTxnID = TxId{}

	// FreezePeriod (Timestamp)
	initialMessage.FreezePeriod = Timestamp{}

	// Message (varchar)
	initialMessage.Message = "Text 12"

	// Encode message
	initialEncoding, err := initialMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Initial encoding : %d bytes", len(initialEncoding))

	// Decode message
	decodedMessage := Order{}

	n, err := decodedMessage.write(initialEncoding)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Decoded : %d bytes", n)

	if n != len(initialEncoding) {
		t.Fatalf("got %v, want %v", n, len(initialEncoding))
	}

	// Serializing the message should give us the same bytes
	secondEncoding, err := decodedMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(initialEncoding, secondEncoding) {
		t.Errorf("got\n%+v\nwant\n%+v", initialEncoding, secondEncoding)
	}

	// if !reflect.DeepEqual(initialMessage, decodedMessage) {
	// 	t.Errorf("\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
	// }

	// Compare re-serialized values
	// AssetType (fixedchar)
	if initialMessage.AssetType != decodedMessage.AssetType {
		t.Errorf("AssetType doesn't match : %s != %s", initialMessage.AssetType, decodedMessage.AssetType)
	}

	// AssetCode (AssetCode)
	if initialMessage.AssetCode != decodedMessage.AssetCode {
		t.Errorf("AssetCode doesn't match : %v != %v", initialMessage.AssetCode, decodedMessage.AssetCode)
	}

	// ComplianceAction (fixedchar)
	// fixedchar test compare not setup

	// TargetAddresses (TargetAddress[])
	if len(initialMessage.TargetAddresses) != len(decodedMessage.TargetAddresses) {
		t.Errorf("TargetAddresses lengths don't match : %d != %d", len(initialMessage.TargetAddresses), len(decodedMessage.TargetAddresses))
	}

	// DepositAddress (PublicKeyHash)
	if initialMessage.DepositAddress != decodedMessage.DepositAddress {
		t.Errorf("DepositAddress doesn't match : %v != %v", initialMessage.DepositAddress, decodedMessage.DepositAddress)
	}

	// AuthorityName (varchar)
	if initialMessage.AuthorityName != decodedMessage.AuthorityName {
		t.Errorf("AuthorityName doesn't match : %s != %s", initialMessage.AuthorityName, decodedMessage.AuthorityName)
	}

	// SigAlgoAddressList (uint)
	// uint test compare not setup

	// AuthorityPublicKey (varchar)
	if initialMessage.AuthorityPublicKey != decodedMessage.AuthorityPublicKey {
		t.Errorf("AuthorityPublicKey doesn't match : %s != %s", initialMessage.AuthorityPublicKey, decodedMessage.AuthorityPublicKey)
	}

	// OrderSignature (varchar)
	if initialMessage.OrderSignature != decodedMessage.OrderSignature {
		t.Errorf("OrderSignature doesn't match : %s != %s", initialMessage.OrderSignature, decodedMessage.OrderSignature)
	}

	// SupportingEvidenceTxId (TxId)
	if initialMessage.SupportingEvidenceTxId != decodedMessage.SupportingEvidenceTxId {
		t.Errorf("SupportingEvidenceTxId doesn't match : %v != %v", initialMessage.SupportingEvidenceTxId, decodedMessage.SupportingEvidenceTxId)
	}

	// RefTxnID (TxId)
	if initialMessage.RefTxnID != decodedMessage.RefTxnID {
		t.Errorf("RefTxnID doesn't match : %v != %v", initialMessage.RefTxnID, decodedMessage.RefTxnID)
	}

	// FreezePeriod (Timestamp)
	if initialMessage.FreezePeriod != decodedMessage.FreezePeriod {
		t.Errorf("FreezePeriod doesn't match : %v != %v", initialMessage.FreezePeriod, decodedMessage.FreezePeriod)
	}

	// Message (varchar)
	if initialMessage.Message != decodedMessage.Message {
		t.Errorf("Message doesn't match : %s != %s", initialMessage.Message, decodedMessage.Message)
	}
}

func TestFreeze(t *testing.T) {
	// Create a randomized object
	initialMessage := Freeze{}
	// Addresses (PublicKeyHash[])
	for i := 0; i < 2; i++ {
		initialMessage.Addresses = append(initialMessage.Addresses, PublicKeyHash{})
	}

	// Timestamp (Timestamp)
	initialMessage.Timestamp = Timestamp{}

	// Encode message
	initialEncoding, err := initialMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Initial encoding : %d bytes", len(initialEncoding))

	// Decode message
	decodedMessage := Freeze{}

	n, err := decodedMessage.write(initialEncoding)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Decoded : %d bytes", n)

	if n != len(initialEncoding) {
		t.Fatalf("got %v, want %v", n, len(initialEncoding))
	}

	// Serializing the message should give us the same bytes
	secondEncoding, err := decodedMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(initialEncoding, secondEncoding) {
		t.Errorf("got\n%+v\nwant\n%+v", initialEncoding, secondEncoding)
	}

	// if !reflect.DeepEqual(initialMessage, decodedMessage) {
	// 	t.Errorf("\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
	// }

	// Compare re-serialized values
	// Addresses (PublicKeyHash[])
	if len(initialMessage.Addresses) != len(decodedMessage.Addresses) {
		t.Errorf("Addresses lengths don't match : %d != %d", len(initialMessage.Addresses), len(decodedMessage.Addresses))
	}

	// Timestamp (Timestamp)
	if initialMessage.Timestamp != decodedMessage.Timestamp {
		t.Errorf("Timestamp doesn't match : %v != %v", initialMessage.Timestamp, decodedMessage.Timestamp)
	}
}

func TestThaw(t *testing.T) {
	// Create a randomized object
	initialMessage := Thaw{}
	// Addresses (PublicKeyHash[])
	for i := 0; i < 2; i++ {
		initialMessage.Addresses = append(initialMessage.Addresses, PublicKeyHash{})
	}

	// RefTxID (TxId)
	initialMessage.RefTxID = TxId{}

	// Timestamp (Timestamp)
	initialMessage.Timestamp = Timestamp{}

	// Encode message
	initialEncoding, err := initialMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Initial encoding : %d bytes", len(initialEncoding))

	// Decode message
	decodedMessage := Thaw{}

	n, err := decodedMessage.write(initialEncoding)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Decoded : %d bytes", n)

	if n != len(initialEncoding) {
		t.Fatalf("got %v, want %v", n, len(initialEncoding))
	}

	// Serializing the message should give us the same bytes
	secondEncoding, err := decodedMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(initialEncoding, secondEncoding) {
		t.Errorf("got\n%+v\nwant\n%+v", initialEncoding, secondEncoding)
	}

	// if !reflect.DeepEqual(initialMessage, decodedMessage) {
	// 	t.Errorf("\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
	// }

	// Compare re-serialized values
	// Addresses (PublicKeyHash[])
	if len(initialMessage.Addresses) != len(decodedMessage.Addresses) {
		t.Errorf("Addresses lengths don't match : %d != %d", len(initialMessage.Addresses), len(decodedMessage.Addresses))
	}

	// RefTxID (TxId)
	if initialMessage.RefTxID != decodedMessage.RefTxID {
		t.Errorf("RefTxID doesn't match : %v != %v", initialMessage.RefTxID, decodedMessage.RefTxID)
	}

	// Timestamp (Timestamp)
	if initialMessage.Timestamp != decodedMessage.Timestamp {
		t.Errorf("Timestamp doesn't match : %v != %v", initialMessage.Timestamp, decodedMessage.Timestamp)
	}
}

func TestConfiscation(t *testing.T) {
	// Create a randomized object
	initialMessage := Confiscation{}
	// Addresses (PublicKeyHash[])
	for i := 0; i < 2; i++ {
		initialMessage.Addresses = append(initialMessage.Addresses, PublicKeyHash{})
	}

	// DepositQty (uint)
	// uint test not setup

	// Timestamp (Timestamp)
	initialMessage.Timestamp = Timestamp{}

	// Encode message
	initialEncoding, err := initialMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Initial encoding : %d bytes", len(initialEncoding))

	// Decode message
	decodedMessage := Confiscation{}

	n, err := decodedMessage.write(initialEncoding)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Decoded : %d bytes", n)

	if n != len(initialEncoding) {
		t.Fatalf("got %v, want %v", n, len(initialEncoding))
	}

	// Serializing the message should give us the same bytes
	secondEncoding, err := decodedMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(initialEncoding, secondEncoding) {
		t.Errorf("got\n%+v\nwant\n%+v", initialEncoding, secondEncoding)
	}

	// if !reflect.DeepEqual(initialMessage, decodedMessage) {
	// 	t.Errorf("\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
	// }

	// Compare re-serialized values
	// Addresses (PublicKeyHash[])
	if len(initialMessage.Addresses) != len(decodedMessage.Addresses) {
		t.Errorf("Addresses lengths don't match : %d != %d", len(initialMessage.Addresses), len(decodedMessage.Addresses))
	}

	// DepositQty (uint)
	// uint test compare not setup

	// Timestamp (Timestamp)
	if initialMessage.Timestamp != decodedMessage.Timestamp {
		t.Errorf("Timestamp doesn't match : %v != %v", initialMessage.Timestamp, decodedMessage.Timestamp)
	}
}

func TestReconciliation(t *testing.T) {
	// Create a randomized object
	initialMessage := Reconciliation{}
	// Addresses (PublicKeyHash[])
	for i := 0; i < 2; i++ {
		initialMessage.Addresses = append(initialMessage.Addresses, PublicKeyHash{})
	}

	// Timestamp (Timestamp)
	initialMessage.Timestamp = Timestamp{}

	// Encode message
	initialEncoding, err := initialMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Initial encoding : %d bytes", len(initialEncoding))

	// Decode message
	decodedMessage := Reconciliation{}

	n, err := decodedMessage.write(initialEncoding)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Decoded : %d bytes", n)

	if n != len(initialEncoding) {
		t.Fatalf("got %v, want %v", n, len(initialEncoding))
	}

	// Serializing the message should give us the same bytes
	secondEncoding, err := decodedMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(initialEncoding, secondEncoding) {
		t.Errorf("got\n%+v\nwant\n%+v", initialEncoding, secondEncoding)
	}

	// if !reflect.DeepEqual(initialMessage, decodedMessage) {
	// 	t.Errorf("\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
	// }

	// Compare re-serialized values
	// Addresses (PublicKeyHash[])
	if len(initialMessage.Addresses) != len(decodedMessage.Addresses) {
		t.Errorf("Addresses lengths don't match : %d != %d", len(initialMessage.Addresses), len(decodedMessage.Addresses))
	}

	// Timestamp (Timestamp)
	if initialMessage.Timestamp != decodedMessage.Timestamp {
		t.Errorf("Timestamp doesn't match : %v != %v", initialMessage.Timestamp, decodedMessage.Timestamp)
	}
}

func TestInitiative(t *testing.T) {
	// Create a randomized object
	initialMessage := Initiative{}
	// AssetType (fixedchar)
	{
		text := make([]byte, 0, 3)
		for i := uint64(0); i < 3; i++ {
			text = append(text, byte(65+i+0))
		}
		initialMessage.AssetType = string(text)
	}

	// AssetCode (AssetCode)
	initialMessage.AssetCode = AssetCode{}

	// VoteSystem (uint)
	// uint test not setup

	// Proposal (bool)
	// bool test not setup

	// ProposedChanges (Amendment[])
	for i := 0; i < 2; i++ {
		initialMessage.ProposedChanges = append(initialMessage.ProposedChanges, Amendment{})
	}

	// VoteOptions (varchar)
	initialMessage.VoteOptions = "Text 5"

	// VoteMax (uint)
	// uint test not setup

	// ProposalDescription (varchar)
	initialMessage.ProposalDescription = "Text 7"

	// ProposalDocumentHash (bin)
	// bin test not setup

	// VoteCutOffTimestamp (Timestamp)
	initialMessage.VoteCutOffTimestamp = Timestamp{}

	// Encode message
	initialEncoding, err := initialMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Initial encoding : %d bytes", len(initialEncoding))

	// Decode message
	decodedMessage := Initiative{}

	n, err := decodedMessage.write(initialEncoding)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Decoded : %d bytes", n)

	if n != len(initialEncoding) {
		t.Fatalf("got %v, want %v", n, len(initialEncoding))
	}

	// Serializing the message should give us the same bytes
	secondEncoding, err := decodedMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(initialEncoding, secondEncoding) {
		t.Errorf("got\n%+v\nwant\n%+v", initialEncoding, secondEncoding)
	}

	// if !reflect.DeepEqual(initialMessage, decodedMessage) {
	// 	t.Errorf("\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
	// }

	// Compare re-serialized values
	// AssetType (fixedchar)
	if initialMessage.AssetType != decodedMessage.AssetType {
		t.Errorf("AssetType doesn't match : %s != %s", initialMessage.AssetType, decodedMessage.AssetType)
	}

	// AssetCode (AssetCode)
	if initialMessage.AssetCode != decodedMessage.AssetCode {
		t.Errorf("AssetCode doesn't match : %v != %v", initialMessage.AssetCode, decodedMessage.AssetCode)
	}

	// VoteSystem (uint)
	// uint test compare not setup

	// Proposal (bool)
	// bool test compare not setup

	// ProposedChanges (Amendment[])
	if len(initialMessage.ProposedChanges) != len(decodedMessage.ProposedChanges) {
		t.Errorf("ProposedChanges lengths don't match : %d != %d", len(initialMessage.ProposedChanges), len(decodedMessage.ProposedChanges))
	}

	// VoteOptions (varchar)
	if initialMessage.VoteOptions != decodedMessage.VoteOptions {
		t.Errorf("VoteOptions doesn't match : %s != %s", initialMessage.VoteOptions, decodedMessage.VoteOptions)
	}

	// VoteMax (uint)
	// uint test compare not setup

	// ProposalDescription (varchar)
	if initialMessage.ProposalDescription != decodedMessage.ProposalDescription {
		t.Errorf("ProposalDescription doesn't match : %s != %s", initialMessage.ProposalDescription, decodedMessage.ProposalDescription)
	}

	// ProposalDocumentHash (bin)
	// bin test compare not setup

	// VoteCutOffTimestamp (Timestamp)
	if initialMessage.VoteCutOffTimestamp != decodedMessage.VoteCutOffTimestamp {
		t.Errorf("VoteCutOffTimestamp doesn't match : %v != %v", initialMessage.VoteCutOffTimestamp, decodedMessage.VoteCutOffTimestamp)
	}
}

func TestReferendum(t *testing.T) {
	// Create a randomized object
	initialMessage := Referendum{}
	// AssetSpecificVote (bool)
	// bool test not setup

	// AssetType (fixedchar)
	{
		text := make([]byte, 0, 3)
		for i := uint64(0); i < 3; i++ {
			text = append(text, byte(65+i+1))
		}
		initialMessage.AssetType = string(text)
	}

	// AssetCode (AssetCode)
	initialMessage.AssetCode = AssetCode{}

	// VoteSystem (uint)
	// uint test not setup

	// Proposal (bool)
	// bool test not setup

	// ProposedChanges (Amendment[])
	for i := 0; i < 2; i++ {
		initialMessage.ProposedChanges = append(initialMessage.ProposedChanges, Amendment{})
	}

	// VoteOptions (varchar)
	initialMessage.VoteOptions = "Text 6"

	// VoteMax (uint)
	// uint test not setup

	// ProposalDescription (varchar)
	initialMessage.ProposalDescription = "Text 8"

	// ProposalDocumentHash (bin)
	// bin test not setup

	// VoteCutOffTimestamp (Timestamp)
	initialMessage.VoteCutOffTimestamp = Timestamp{}

	// Encode message
	initialEncoding, err := initialMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Initial encoding : %d bytes", len(initialEncoding))

	// Decode message
	decodedMessage := Referendum{}

	n, err := decodedMessage.write(initialEncoding)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Decoded : %d bytes", n)

	if n != len(initialEncoding) {
		t.Fatalf("got %v, want %v", n, len(initialEncoding))
	}

	// Serializing the message should give us the same bytes
	secondEncoding, err := decodedMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(initialEncoding, secondEncoding) {
		t.Errorf("got\n%+v\nwant\n%+v", initialEncoding, secondEncoding)
	}

	// if !reflect.DeepEqual(initialMessage, decodedMessage) {
	// 	t.Errorf("\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
	// }

	// Compare re-serialized values
	// AssetSpecificVote (bool)
	// bool test compare not setup

	// AssetType (fixedchar)
	if initialMessage.AssetType != decodedMessage.AssetType {
		t.Errorf("AssetType doesn't match : %s != %s", initialMessage.AssetType, decodedMessage.AssetType)
	}

	// AssetCode (AssetCode)
	if initialMessage.AssetCode != decodedMessage.AssetCode {
		t.Errorf("AssetCode doesn't match : %v != %v", initialMessage.AssetCode, decodedMessage.AssetCode)
	}

	// VoteSystem (uint)
	// uint test compare not setup

	// Proposal (bool)
	// bool test compare not setup

	// ProposedChanges (Amendment[])
	if len(initialMessage.ProposedChanges) != len(decodedMessage.ProposedChanges) {
		t.Errorf("ProposedChanges lengths don't match : %d != %d", len(initialMessage.ProposedChanges), len(decodedMessage.ProposedChanges))
	}

	// VoteOptions (varchar)
	if initialMessage.VoteOptions != decodedMessage.VoteOptions {
		t.Errorf("VoteOptions doesn't match : %s != %s", initialMessage.VoteOptions, decodedMessage.VoteOptions)
	}

	// VoteMax (uint)
	// uint test compare not setup

	// ProposalDescription (varchar)
	if initialMessage.ProposalDescription != decodedMessage.ProposalDescription {
		t.Errorf("ProposalDescription doesn't match : %s != %s", initialMessage.ProposalDescription, decodedMessage.ProposalDescription)
	}

	// ProposalDocumentHash (bin)
	// bin test compare not setup

	// VoteCutOffTimestamp (Timestamp)
	if initialMessage.VoteCutOffTimestamp != decodedMessage.VoteCutOffTimestamp {
		t.Errorf("VoteCutOffTimestamp doesn't match : %v != %v", initialMessage.VoteCutOffTimestamp, decodedMessage.VoteCutOffTimestamp)
	}
}

func TestVote(t *testing.T) {
	// Create a randomized object
	initialMessage := Vote{}
	// Timestamp (Timestamp)
	initialMessage.Timestamp = Timestamp{}

	// Encode message
	initialEncoding, err := initialMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Initial encoding : %d bytes", len(initialEncoding))

	// Decode message
	decodedMessage := Vote{}

	n, err := decodedMessage.write(initialEncoding)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Decoded : %d bytes", n)

	if n != len(initialEncoding) {
		t.Fatalf("got %v, want %v", n, len(initialEncoding))
	}

	// Serializing the message should give us the same bytes
	secondEncoding, err := decodedMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(initialEncoding, secondEncoding) {
		t.Errorf("got\n%+v\nwant\n%+v", initialEncoding, secondEncoding)
	}

	// if !reflect.DeepEqual(initialMessage, decodedMessage) {
	// 	t.Errorf("\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
	// }

	// Compare re-serialized values
	// Timestamp (Timestamp)
	if initialMessage.Timestamp != decodedMessage.Timestamp {
		t.Errorf("Timestamp doesn't match : %v != %v", initialMessage.Timestamp, decodedMessage.Timestamp)
	}
}

func TestBallotCast(t *testing.T) {
	// Create a randomized object
	initialMessage := BallotCast{}
	// AssetType (fixedchar)
	{
		text := make([]byte, 0, 3)
		for i := uint64(0); i < 3; i++ {
			text = append(text, byte(65+i+0))
		}
		initialMessage.AssetType = string(text)
	}

	// AssetCode (AssetCode)
	initialMessage.AssetCode = AssetCode{}

	// VoteTxID (TxId)
	initialMessage.VoteTxID = TxId{}

	// Vote (varchar)
	initialMessage.Vote = "Text 3"

	// Encode message
	initialEncoding, err := initialMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Initial encoding : %d bytes", len(initialEncoding))

	// Decode message
	decodedMessage := BallotCast{}

	n, err := decodedMessage.write(initialEncoding)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Decoded : %d bytes", n)

	if n != len(initialEncoding) {
		t.Fatalf("got %v, want %v", n, len(initialEncoding))
	}

	// Serializing the message should give us the same bytes
	secondEncoding, err := decodedMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(initialEncoding, secondEncoding) {
		t.Errorf("got\n%+v\nwant\n%+v", initialEncoding, secondEncoding)
	}

	// if !reflect.DeepEqual(initialMessage, decodedMessage) {
	// 	t.Errorf("\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
	// }

	// Compare re-serialized values
	// AssetType (fixedchar)
	if initialMessage.AssetType != decodedMessage.AssetType {
		t.Errorf("AssetType doesn't match : %s != %s", initialMessage.AssetType, decodedMessage.AssetType)
	}

	// AssetCode (AssetCode)
	if initialMessage.AssetCode != decodedMessage.AssetCode {
		t.Errorf("AssetCode doesn't match : %v != %v", initialMessage.AssetCode, decodedMessage.AssetCode)
	}

	// VoteTxID (TxId)
	if initialMessage.VoteTxID != decodedMessage.VoteTxID {
		t.Errorf("VoteTxID doesn't match : %v != %v", initialMessage.VoteTxID, decodedMessage.VoteTxID)
	}

	// Vote (varchar)
	if initialMessage.Vote != decodedMessage.Vote {
		t.Errorf("Vote doesn't match : %s != %s", initialMessage.Vote, decodedMessage.Vote)
	}
}

func TestBallotCounted(t *testing.T) {
	// Create a randomized object
	initialMessage := BallotCounted{}
	// Timestamp (Timestamp)
	initialMessage.Timestamp = Timestamp{}

	// Encode message
	initialEncoding, err := initialMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Initial encoding : %d bytes", len(initialEncoding))

	// Decode message
	decodedMessage := BallotCounted{}

	n, err := decodedMessage.write(initialEncoding)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Decoded : %d bytes", n)

	if n != len(initialEncoding) {
		t.Fatalf("got %v, want %v", n, len(initialEncoding))
	}

	// Serializing the message should give us the same bytes
	secondEncoding, err := decodedMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(initialEncoding, secondEncoding) {
		t.Errorf("got\n%+v\nwant\n%+v", initialEncoding, secondEncoding)
	}

	// if !reflect.DeepEqual(initialMessage, decodedMessage) {
	// 	t.Errorf("\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
	// }

	// Compare re-serialized values
	// Timestamp (Timestamp)
	if initialMessage.Timestamp != decodedMessage.Timestamp {
		t.Errorf("Timestamp doesn't match : %v != %v", initialMessage.Timestamp, decodedMessage.Timestamp)
	}
}

func TestResult(t *testing.T) {
	// Create a randomized object
	initialMessage := Result{}
	// AssetType (fixedchar)
	{
		text := make([]byte, 0, 3)
		for i := uint64(0); i < 3; i++ {
			text = append(text, byte(65+i+0))
		}
		initialMessage.AssetType = string(text)
	}

	// AssetCode (AssetCode)
	initialMessage.AssetCode = AssetCode{}

	// Proposal (bool)
	// bool test not setup

	// ProposedChanges (Amendment[])
	for i := 0; i < 2; i++ {
		initialMessage.ProposedChanges = append(initialMessage.ProposedChanges, Amendment{})
	}

	// VoteTxID (TxId)
	initialMessage.VoteTxID = TxId{}

	// VoteOptionsCount (uint)
	// uint test not setup

	// OptionXTally (uint)
	// uint test not setup

	// Result (varchar)
	initialMessage.Result = "Text 7"

	// Timestamp (Timestamp)
	initialMessage.Timestamp = Timestamp{}

	// Encode message
	initialEncoding, err := initialMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Initial encoding : %d bytes", len(initialEncoding))

	// Decode message
	decodedMessage := Result{}

	n, err := decodedMessage.write(initialEncoding)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Decoded : %d bytes", n)

	if n != len(initialEncoding) {
		t.Fatalf("got %v, want %v", n, len(initialEncoding))
	}

	// Serializing the message should give us the same bytes
	secondEncoding, err := decodedMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(initialEncoding, secondEncoding) {
		t.Errorf("got\n%+v\nwant\n%+v", initialEncoding, secondEncoding)
	}

	// if !reflect.DeepEqual(initialMessage, decodedMessage) {
	// 	t.Errorf("\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
	// }

	// Compare re-serialized values
	// AssetType (fixedchar)
	if initialMessage.AssetType != decodedMessage.AssetType {
		t.Errorf("AssetType doesn't match : %s != %s", initialMessage.AssetType, decodedMessage.AssetType)
	}

	// AssetCode (AssetCode)
	if initialMessage.AssetCode != decodedMessage.AssetCode {
		t.Errorf("AssetCode doesn't match : %v != %v", initialMessage.AssetCode, decodedMessage.AssetCode)
	}

	// Proposal (bool)
	// bool test compare not setup

	// ProposedChanges (Amendment[])
	if len(initialMessage.ProposedChanges) != len(decodedMessage.ProposedChanges) {
		t.Errorf("ProposedChanges lengths don't match : %d != %d", len(initialMessage.ProposedChanges), len(decodedMessage.ProposedChanges))
	}

	// VoteTxID (TxId)
	if initialMessage.VoteTxID != decodedMessage.VoteTxID {
		t.Errorf("VoteTxID doesn't match : %v != %v", initialMessage.VoteTxID, decodedMessage.VoteTxID)
	}

	// VoteOptionsCount (uint)
	// uint test compare not setup

	// OptionXTally (uint)
	// uint test compare not setup

	// Result (varchar)
	if initialMessage.Result != decodedMessage.Result {
		t.Errorf("Result doesn't match : %s != %s", initialMessage.Result, decodedMessage.Result)
	}

	// Timestamp (Timestamp)
	if initialMessage.Timestamp != decodedMessage.Timestamp {
		t.Errorf("Timestamp doesn't match : %v != %v", initialMessage.Timestamp, decodedMessage.Timestamp)
	}
}

func TestMessage(t *testing.T) {
	// Create a randomized object
	initialMessage := Message{}
	// AddressIndexes (uint16[])
	for i := 0; i < 5; i++ {
		var item uint16
		initialMessage.AddressIndexes = append(initialMessage.AddressIndexes, item)
	}

	// MessageType (fixedchar)
	{
		text := make([]byte, 0, 4)
		for i := uint64(0); i < 4; i++ {
			text = append(text, byte(65+i+1))
		}
		initialMessage.MessageType = string(text)
	}

	// MessagePayload (varbin)
	initialMessage.MessagePayload = make([]byte, 0, 32)
	for i := uint64(0); i < 32; i++ {
		initialMessage.MessagePayload = append(initialMessage.MessagePayload, byte(65+i+2))
	}

	// Encode message
	initialEncoding, err := initialMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Initial encoding : %d bytes", len(initialEncoding))

	// Decode message
	decodedMessage := Message{}

	n, err := decodedMessage.write(initialEncoding)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Decoded : %d bytes", n)

	if n != len(initialEncoding) {
		t.Fatalf("got %v, want %v", n, len(initialEncoding))
	}

	// Serializing the message should give us the same bytes
	secondEncoding, err := decodedMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(initialEncoding, secondEncoding) {
		t.Errorf("got\n%+v\nwant\n%+v", initialEncoding, secondEncoding)
	}

	// if !reflect.DeepEqual(initialMessage, decodedMessage) {
	// 	t.Errorf("\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
	// }

	// Compare re-serialized values
	// AddressIndexes (uint16[])
	if len(initialMessage.AddressIndexes) != len(decodedMessage.AddressIndexes) {
		t.Errorf("AddressIndexes lengths don't match : %d != %d", len(initialMessage.AddressIndexes), len(decodedMessage.AddressIndexes))
	}
	for i, value := range initialMessage.AddressIndexes {
		if value != decodedMessage.AddressIndexes[i] {
			t.Errorf("AddressIndexes value %d doesn't match : %v != %v", i, value, decodedMessage.AddressIndexes[i])
		}
	}

	// MessageType (fixedchar)
	if initialMessage.MessageType != decodedMessage.MessageType {
		t.Errorf("MessageType doesn't match : %s != %s", initialMessage.MessageType, decodedMessage.MessageType)
	}

	// MessagePayload (varbin)
	if !bytes.Equal(initialMessage.MessagePayload, decodedMessage.MessagePayload) {
		t.Errorf("MessagePayload doesn't match : %x != %x", initialMessage.MessagePayload, decodedMessage.MessagePayload)
	}
}

func TestRejection(t *testing.T) {
	// Create a randomized object
	initialMessage := Rejection{}
	// QtyReceivingAddresses (uint8)
	// uint8 test not setup

	// AddressIndexes (uint16[])
	for i := 0; i < 5; i++ {
		var item uint16
		initialMessage.AddressIndexes = append(initialMessage.AddressIndexes, item)
	}

	// RejectionType (uint)
	// uint test not setup

	// MessagePayload (varchar)
	initialMessage.MessagePayload = "Text 3"

	// Timestamp (Timestamp)
	initialMessage.Timestamp = Timestamp{}

	// Encode message
	initialEncoding, err := initialMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Initial encoding : %d bytes", len(initialEncoding))

	// Decode message
	decodedMessage := Rejection{}

	n, err := decodedMessage.write(initialEncoding)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Decoded : %d bytes", n)

	if n != len(initialEncoding) {
		t.Fatalf("got %v, want %v", n, len(initialEncoding))
	}

	// Serializing the message should give us the same bytes
	secondEncoding, err := decodedMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(initialEncoding, secondEncoding) {
		t.Errorf("got\n%+v\nwant\n%+v", initialEncoding, secondEncoding)
	}

	// if !reflect.DeepEqual(initialMessage, decodedMessage) {
	// 	t.Errorf("\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
	// }

	// Compare re-serialized values
	// QtyReceivingAddresses (uint8)
	// uint8 test compare not setup

	// AddressIndexes (uint16[])
	if len(initialMessage.AddressIndexes) != len(decodedMessage.AddressIndexes) {
		t.Errorf("AddressIndexes lengths don't match : %d != %d", len(initialMessage.AddressIndexes), len(decodedMessage.AddressIndexes))
	}
	for i, value := range initialMessage.AddressIndexes {
		if value != decodedMessage.AddressIndexes[i] {
			t.Errorf("AddressIndexes value %d doesn't match : %v != %v", i, value, decodedMessage.AddressIndexes[i])
		}
	}

	// RejectionType (uint)
	// uint test compare not setup

	// MessagePayload (varchar)
	if initialMessage.MessagePayload != decodedMessage.MessagePayload {
		t.Errorf("MessagePayload doesn't match : %s != %s", initialMessage.MessagePayload, decodedMessage.MessagePayload)
	}

	// Timestamp (Timestamp)
	if initialMessage.Timestamp != decodedMessage.Timestamp {
		t.Errorf("Timestamp doesn't match : %v != %v", initialMessage.Timestamp, decodedMessage.Timestamp)
	}
}

func TestEstablishment(t *testing.T) {
	// Create a randomized object
	initialMessage := Establishment{}
	// Message (varchar)
	initialMessage.Message = "Text 0"

	// Encode message
	initialEncoding, err := initialMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Initial encoding : %d bytes", len(initialEncoding))

	// Decode message
	decodedMessage := Establishment{}

	n, err := decodedMessage.write(initialEncoding)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Decoded : %d bytes", n)

	if n != len(initialEncoding) {
		t.Fatalf("got %v, want %v", n, len(initialEncoding))
	}

	// Serializing the message should give us the same bytes
	secondEncoding, err := decodedMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(initialEncoding, secondEncoding) {
		t.Errorf("got\n%+v\nwant\n%+v", initialEncoding, secondEncoding)
	}

	// if !reflect.DeepEqual(initialMessage, decodedMessage) {
	// 	t.Errorf("\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
	// }

	// Compare re-serialized values
	// Message (varchar)
	if initialMessage.Message != decodedMessage.Message {
		t.Errorf("Message doesn't match : %s != %s", initialMessage.Message, decodedMessage.Message)
	}
}

func TestAddition(t *testing.T) {
	// Create a randomized object
	initialMessage := Addition{}
	// Message (varchar)
	initialMessage.Message = "Text 0"

	// Encode message
	initialEncoding, err := initialMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Initial encoding : %d bytes", len(initialEncoding))

	// Decode message
	decodedMessage := Addition{}

	n, err := decodedMessage.write(initialEncoding)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Decoded : %d bytes", n)

	if n != len(initialEncoding) {
		t.Fatalf("got %v, want %v", n, len(initialEncoding))
	}

	// Serializing the message should give us the same bytes
	secondEncoding, err := decodedMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(initialEncoding, secondEncoding) {
		t.Errorf("got\n%+v\nwant\n%+v", initialEncoding, secondEncoding)
	}

	// if !reflect.DeepEqual(initialMessage, decodedMessage) {
	// 	t.Errorf("\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
	// }

	// Compare re-serialized values
	// Message (varchar)
	if initialMessage.Message != decodedMessage.Message {
		t.Errorf("Message doesn't match : %s != %s", initialMessage.Message, decodedMessage.Message)
	}
}

func TestAlteration(t *testing.T) {
	// Create a randomized object
	initialMessage := Alteration{}
	// Message (varchar)
	initialMessage.Message = "Text 0"

	// Encode message
	initialEncoding, err := initialMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Initial encoding : %d bytes", len(initialEncoding))

	// Decode message
	decodedMessage := Alteration{}

	n, err := decodedMessage.write(initialEncoding)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Decoded : %d bytes", n)

	if n != len(initialEncoding) {
		t.Fatalf("got %v, want %v", n, len(initialEncoding))
	}

	// Serializing the message should give us the same bytes
	secondEncoding, err := decodedMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(initialEncoding, secondEncoding) {
		t.Errorf("got\n%+v\nwant\n%+v", initialEncoding, secondEncoding)
	}

	// if !reflect.DeepEqual(initialMessage, decodedMessage) {
	// 	t.Errorf("\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
	// }

	// Compare re-serialized values
	// Message (varchar)
	if initialMessage.Message != decodedMessage.Message {
		t.Errorf("Message doesn't match : %s != %s", initialMessage.Message, decodedMessage.Message)
	}
}

func TestRemoval(t *testing.T) {
	// Create a randomized object
	initialMessage := Removal{}
	// Message (varchar)
	initialMessage.Message = "Text 0"

	// Encode message
	initialEncoding, err := initialMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Initial encoding : %d bytes", len(initialEncoding))

	// Decode message
	decodedMessage := Removal{}

	n, err := decodedMessage.write(initialEncoding)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Decoded : %d bytes", n)

	if n != len(initialEncoding) {
		t.Fatalf("got %v, want %v", n, len(initialEncoding))
	}

	// Serializing the message should give us the same bytes
	secondEncoding, err := decodedMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(initialEncoding, secondEncoding) {
		t.Errorf("got\n%+v\nwant\n%+v", initialEncoding, secondEncoding)
	}

	// if !reflect.DeepEqual(initialMessage, decodedMessage) {
	// 	t.Errorf("\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
	// }

	// Compare re-serialized values
	// Message (varchar)
	if initialMessage.Message != decodedMessage.Message {
		t.Errorf("Message doesn't match : %s != %s", initialMessage.Message, decodedMessage.Message)
	}
}

func TestTransfer(t *testing.T) {
	// Create a randomized object
	initialMessage := Transfer{}
	// Assets (AssetTransfer[])
	for i := 0; i < 2; i++ {
		initialMessage.Assets = append(initialMessage.Assets, AssetTransfer{})
	}

	// OfferExpiry (Timestamp)
	initialMessage.OfferExpiry = Timestamp{}

	// ExchangeFeeCurrency (fixedchar)
	{
		text := make([]byte, 0, 3)
		for i := uint64(0); i < 3; i++ {
			text = append(text, byte(65+i+2))
		}
		initialMessage.ExchangeFeeCurrency = string(text)
	}

	// ExchangeFeeVar (float)
	// float test not setup

	// ExchangeFeeFixed (float)
	// float test not setup

	// ExchangeFeeAddress (PublicKeyHash)
	initialMessage.ExchangeFeeAddress = PublicKeyHash{}

	// Encode message
	initialEncoding, err := initialMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Initial encoding : %d bytes", len(initialEncoding))

	// Decode message
	decodedMessage := Transfer{}

	n, err := decodedMessage.write(initialEncoding)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Decoded : %d bytes", n)

	if n != len(initialEncoding) {
		t.Fatalf("got %v, want %v", n, len(initialEncoding))
	}

	// Serializing the message should give us the same bytes
	secondEncoding, err := decodedMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(initialEncoding, secondEncoding) {
		t.Errorf("got\n%+v\nwant\n%+v", initialEncoding, secondEncoding)
	}

	// if !reflect.DeepEqual(initialMessage, decodedMessage) {
	// 	t.Errorf("\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
	// }

	// Compare re-serialized values
	// Assets (AssetTransfer[])
	if len(initialMessage.Assets) != len(decodedMessage.Assets) {
		t.Errorf("Assets lengths don't match : %d != %d", len(initialMessage.Assets), len(decodedMessage.Assets))
	}

	// OfferExpiry (Timestamp)
	if initialMessage.OfferExpiry != decodedMessage.OfferExpiry {
		t.Errorf("OfferExpiry doesn't match : %v != %v", initialMessage.OfferExpiry, decodedMessage.OfferExpiry)
	}

	// ExchangeFeeCurrency (fixedchar)
	if initialMessage.ExchangeFeeCurrency != decodedMessage.ExchangeFeeCurrency {
		t.Errorf("ExchangeFeeCurrency doesn't match : %s != %s", initialMessage.ExchangeFeeCurrency, decodedMessage.ExchangeFeeCurrency)
	}

	// ExchangeFeeVar (float)
	// float test compare not setup

	// ExchangeFeeFixed (float)
	// float test compare not setup

	// ExchangeFeeAddress (PublicKeyHash)
	if initialMessage.ExchangeFeeAddress != decodedMessage.ExchangeFeeAddress {
		t.Errorf("ExchangeFeeAddress doesn't match : %v != %v", initialMessage.ExchangeFeeAddress, decodedMessage.ExchangeFeeAddress)
	}
}

func TestSettlement(t *testing.T) {
	// Create a randomized object
	initialMessage := Settlement{}
	// Assets (AssetSettlement[])
	for i := 0; i < 2; i++ {
		initialMessage.Assets = append(initialMessage.Assets, AssetSettlement{})
	}

	// Timestamp (Timestamp)
	initialMessage.Timestamp = Timestamp{}

	// Encode message
	initialEncoding, err := initialMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Initial encoding : %d bytes", len(initialEncoding))

	// Decode message
	decodedMessage := Settlement{}

	n, err := decodedMessage.write(initialEncoding)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Decoded : %d bytes", n)

	if n != len(initialEncoding) {
		t.Fatalf("got %v, want %v", n, len(initialEncoding))
	}

	// Serializing the message should give us the same bytes
	secondEncoding, err := decodedMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(initialEncoding, secondEncoding) {
		t.Errorf("got\n%+v\nwant\n%+v", initialEncoding, secondEncoding)
	}

	// if !reflect.DeepEqual(initialMessage, decodedMessage) {
	// 	t.Errorf("\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
	// }

	// Compare re-serialized values
	// Assets (AssetSettlement[])
	if len(initialMessage.Assets) != len(decodedMessage.Assets) {
		t.Errorf("Assets lengths don't match : %d != %d", len(initialMessage.Assets), len(decodedMessage.Assets))
	}

	// Timestamp (Timestamp)
	if initialMessage.Timestamp != decodedMessage.Timestamp {
		t.Errorf("Timestamp doesn't match : %v != %v", initialMessage.Timestamp, decodedMessage.Timestamp)
	}
}
