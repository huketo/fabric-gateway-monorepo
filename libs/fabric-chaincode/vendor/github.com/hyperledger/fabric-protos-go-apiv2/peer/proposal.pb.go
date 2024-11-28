// Copyright the Hyperledger Fabric contributors. All rights reserved.
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: peer/proposal.proto

package peer

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This structure is necessary to sign the proposal which contains the header
// and the payload. Without this structure, we would have to concatenate the
// header and the payload to verify the signature, which could be expensive
// with large payload
//
// When an endorser receives a SignedProposal message, it should verify the
// signature over the proposal bytes. This verification requires the following
// steps:
//  1. Verification of the validity of the certificate that was used to produce
//     the signature.  The certificate will be available once proposalBytes has
//     been unmarshalled to a Proposal message, and Proposal.header has been
//     unmarshalled to a Header message. While this unmarshalling-before-verifying
//     might not be ideal, it is unavoidable because i) the signature needs to also
//     protect the signing certificate; ii) it is desirable that Header is created
//     once by the client and never changed (for the sake of accountability and
//     non-repudiation). Note also that it is actually impossible to conclusively
//     verify the validity of the certificate included in a Proposal, because the
//     proposal needs to first be endorsed and ordered with respect to certificate
//     expiration transactions. Still, it is useful to pre-filter expired
//     certificates at this stage.
//  2. Verification that the certificate is trusted (signed by a trusted CA) and
//     that it is allowed to transact with us (with respect to some ACLs);
//  3. Verification that the signature on proposalBytes is valid;
//  4. Detect replay attacks;
type SignedProposal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The bytes of Proposal
	ProposalBytes []byte `protobuf:"bytes,1,opt,name=proposal_bytes,json=proposalBytes,proto3" json:"proposal_bytes,omitempty"`
	// Signaure over proposalBytes; this signature is to be verified against
	// the creator identity contained in the header of the Proposal message
	// marshaled as proposalBytes
	Signature []byte `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *SignedProposal) Reset() {
	*x = SignedProposal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_peer_proposal_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignedProposal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignedProposal) ProtoMessage() {}

func (x *SignedProposal) ProtoReflect() protoreflect.Message {
	mi := &file_peer_proposal_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignedProposal.ProtoReflect.Descriptor instead.
func (*SignedProposal) Descriptor() ([]byte, []int) {
	return file_peer_proposal_proto_rawDescGZIP(), []int{0}
}

func (x *SignedProposal) GetProposalBytes() []byte {
	if x != nil {
		return x.ProposalBytes
	}
	return nil
}

func (x *SignedProposal) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

// A Proposal is sent to an endorser for endorsement.  The proposal contains:
//  1. A header which should be unmarshaled to a Header message.  Note that
//     Header is both the header of a Proposal and of a Transaction, in that i)
//     both headers should be unmarshaled to this message; and ii) it is used to
//     compute cryptographic hashes and signatures.  The header has fields common
//     to all proposals/transactions.  In addition it has a type field for
//     additional customization. An example of this is the ChaincodeHeaderExtension
//     message used to extend the Header for type CHAINCODE.
//  2. A payload whose type depends on the header's type field.
//  3. An extension whose type depends on the header's type field.
//
// Let us see an example. For type CHAINCODE (see the Header message),
// we have the following:
//  1. The header is a Header message whose extensions field is a
//     ChaincodeHeaderExtension message.
//  2. The payload is a ChaincodeProposalPayload message.
//  3. The extension is a ChaincodeAction that might be used to ask the
//     endorsers to endorse a specific ChaincodeAction, thus emulating the
//     submitting peer model.
type Proposal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The header of the proposal. It is the bytes of the Header
	Header []byte `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// The payload of the proposal as defined by the type in the proposal
	// header.
	Payload []byte `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	// Optional extensions to the proposal. Its content depends on the Header's
	// type field.  For the type CHAINCODE, it might be the bytes of a
	// ChaincodeAction message.
	Extension []byte `protobuf:"bytes,3,opt,name=extension,proto3" json:"extension,omitempty"`
}

func (x *Proposal) Reset() {
	*x = Proposal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_peer_proposal_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Proposal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Proposal) ProtoMessage() {}

func (x *Proposal) ProtoReflect() protoreflect.Message {
	mi := &file_peer_proposal_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Proposal.ProtoReflect.Descriptor instead.
func (*Proposal) Descriptor() ([]byte, []int) {
	return file_peer_proposal_proto_rawDescGZIP(), []int{1}
}

func (x *Proposal) GetHeader() []byte {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *Proposal) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *Proposal) GetExtension() []byte {
	if x != nil {
		return x.Extension
	}
	return nil
}

// ChaincodeHeaderExtension is the Header's extentions message to be used when
// the Header's type is CHAINCODE.  This extensions is used to specify which
// chaincode to invoke and what should appear on the ledger.
type ChaincodeHeaderExtension struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the chaincode to target.
	ChaincodeId *ChaincodeID `protobuf:"bytes,2,opt,name=chaincode_id,json=chaincodeId,proto3" json:"chaincode_id,omitempty"`
}

func (x *ChaincodeHeaderExtension) Reset() {
	*x = ChaincodeHeaderExtension{}
	if protoimpl.UnsafeEnabled {
		mi := &file_peer_proposal_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChaincodeHeaderExtension) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChaincodeHeaderExtension) ProtoMessage() {}

func (x *ChaincodeHeaderExtension) ProtoReflect() protoreflect.Message {
	mi := &file_peer_proposal_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChaincodeHeaderExtension.ProtoReflect.Descriptor instead.
func (*ChaincodeHeaderExtension) Descriptor() ([]byte, []int) {
	return file_peer_proposal_proto_rawDescGZIP(), []int{2}
}

func (x *ChaincodeHeaderExtension) GetChaincodeId() *ChaincodeID {
	if x != nil {
		return x.ChaincodeId
	}
	return nil
}

// ChaincodeProposalPayload is the Proposal's payload message to be used when
// the Header's type is CHAINCODE.  It contains the arguments for this
// invocation.
type ChaincodeProposalPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Input contains the arguments for this invocation. If this invocation
	// deploys a new chaincode, ESCC/VSCC are part of this field.
	// This is usually a marshaled ChaincodeInvocationSpec
	Input []byte `protobuf:"bytes,1,opt,name=input,proto3" json:"input,omitempty"`
	// TransientMap contains data (e.g. cryptographic material) that might be used
	// to implement some form of application-level confidentiality. The contents
	// of this field are supposed to always be omitted from the transaction and
	// excluded from the ledger.
	TransientMap map[string][]byte `protobuf:"bytes,2,rep,name=TransientMap,proto3" json:"TransientMap,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ChaincodeProposalPayload) Reset() {
	*x = ChaincodeProposalPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_peer_proposal_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChaincodeProposalPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChaincodeProposalPayload) ProtoMessage() {}

func (x *ChaincodeProposalPayload) ProtoReflect() protoreflect.Message {
	mi := &file_peer_proposal_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChaincodeProposalPayload.ProtoReflect.Descriptor instead.
func (*ChaincodeProposalPayload) Descriptor() ([]byte, []int) {
	return file_peer_proposal_proto_rawDescGZIP(), []int{3}
}

func (x *ChaincodeProposalPayload) GetInput() []byte {
	if x != nil {
		return x.Input
	}
	return nil
}

func (x *ChaincodeProposalPayload) GetTransientMap() map[string][]byte {
	if x != nil {
		return x.TransientMap
	}
	return nil
}

// ChaincodeAction contains the executed chaincode results, response, and event.
type ChaincodeAction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// This field contains the read set and the write set produced by the
	// chaincode executing this invocation.
	Results []byte `protobuf:"bytes,1,opt,name=results,proto3" json:"results,omitempty"`
	// This field contains the event generated by the chaincode.
	// Only a single marshaled ChaincodeEvent is included.
	Events []byte `protobuf:"bytes,2,opt,name=events,proto3" json:"events,omitempty"`
	// This field contains the result of executing this invocation.
	Response *Response `protobuf:"bytes,3,opt,name=response,proto3" json:"response,omitempty"`
	// This field contains the ChaincodeID of executing this invocation. Endorser
	// will set it with the ChaincodeID called by endorser while simulating proposal.
	// Committer will validate the version matching with latest chaincode version.
	// Adding ChaincodeID to keep version opens up the possibility of multiple
	// ChaincodeAction per transaction.
	ChaincodeId *ChaincodeID `protobuf:"bytes,4,opt,name=chaincode_id,json=chaincodeId,proto3" json:"chaincode_id,omitempty"`
}

func (x *ChaincodeAction) Reset() {
	*x = ChaincodeAction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_peer_proposal_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChaincodeAction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChaincodeAction) ProtoMessage() {}

func (x *ChaincodeAction) ProtoReflect() protoreflect.Message {
	mi := &file_peer_proposal_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChaincodeAction.ProtoReflect.Descriptor instead.
func (*ChaincodeAction) Descriptor() ([]byte, []int) {
	return file_peer_proposal_proto_rawDescGZIP(), []int{4}
}

func (x *ChaincodeAction) GetResults() []byte {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *ChaincodeAction) GetEvents() []byte {
	if x != nil {
		return x.Events
	}
	return nil
}

func (x *ChaincodeAction) GetResponse() *Response {
	if x != nil {
		return x.Response
	}
	return nil
}

func (x *ChaincodeAction) GetChaincodeId() *ChaincodeID {
	if x != nil {
		return x.ChaincodeId
	}
	return nil
}

var File_peer_proposal_proto protoreflect.FileDescriptor

var file_peer_proposal_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x65, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x1a, 0x14, 0x70,
	0x65, 0x65, 0x72, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x70, 0x65, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73,
	0x61, 0x6c, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x55, 0x0a, 0x0e, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x70, 0x6f,
	0x73, 0x61, 0x6c, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x5f,
	0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d, 0x70, 0x72, 0x6f,
	0x70, 0x6f, 0x73, 0x61, 0x6c, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73,
	0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x5a, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x70,
	0x6f, 0x73, 0x61, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07,
	0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x70,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x22, 0x6b, 0x0a, 0x18, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x63, 0x6f, 0x64,
	0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x36, 0x0a, 0x0c, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x43, 0x68, 0x61, 0x69, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x49, 0x44, 0x52, 0x0b, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x4a, 0x04, 0x08, 0x01, 0x10, 0x02, 0x52, 0x11,
	0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x76, 0x69, 0x73, 0x62, 0x69, 0x6c, 0x69, 0x74,
	0x79, 0x22, 0xc9, 0x01, 0x0a, 0x18, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x50,
	0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x69,
	0x6e, 0x70, 0x75, 0x74, 0x12, 0x56, 0x0a, 0x0c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x65, 0x6e,
	0x74, 0x4d, 0x61, 0x70, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2e, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x50, 0x72, 0x6f,
	0x70, 0x6f, 0x73, 0x61, 0x6c, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x61, 0x70, 0x1a, 0x3f, 0x0a, 0x11,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xc1, 0x01,
	0x0a, 0x0f, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x12, 0x2c, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x36, 0x0a, 0x0c, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x49, 0x44, 0x52, 0x0b, 0x63, 0x68,
	0x61, 0x69, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x4a, 0x04, 0x08, 0x05, 0x10, 0x06, 0x52,
	0x10, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x42, 0x9f, 0x01, 0x0a, 0x22, 0x6f, 0x72, 0x67, 0x2e, 0x68, 0x79, 0x70, 0x65, 0x72, 0x6c,
	0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x66, 0x61, 0x62, 0x72, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2e, 0x70, 0x65, 0x65, 0x72, 0x42, 0x0d, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73,
	0x61, 0x6c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x79, 0x70, 0x65, 0x72, 0x6c, 0x65, 0x64, 0x67, 0x65,
	0x72, 0x2f, 0x66, 0x61, 0x62, 0x72, 0x69, 0x63, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2d,
	0x67, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x70, 0x65, 0x65, 0x72, 0xa2, 0x02, 0x03,
	0x50, 0x58, 0x58, 0xaa, 0x02, 0x06, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0xca, 0x02, 0x06, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0xe2, 0x02, 0x12, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x06, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_peer_proposal_proto_rawDescOnce sync.Once
	file_peer_proposal_proto_rawDescData = file_peer_proposal_proto_rawDesc
)

func file_peer_proposal_proto_rawDescGZIP() []byte {
	file_peer_proposal_proto_rawDescOnce.Do(func() {
		file_peer_proposal_proto_rawDescData = protoimpl.X.CompressGZIP(file_peer_proposal_proto_rawDescData)
	})
	return file_peer_proposal_proto_rawDescData
}

var file_peer_proposal_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_peer_proposal_proto_goTypes = []interface{}{
	(*SignedProposal)(nil),           // 0: protos.SignedProposal
	(*Proposal)(nil),                 // 1: protos.Proposal
	(*ChaincodeHeaderExtension)(nil), // 2: protos.ChaincodeHeaderExtension
	(*ChaincodeProposalPayload)(nil), // 3: protos.ChaincodeProposalPayload
	(*ChaincodeAction)(nil),          // 4: protos.ChaincodeAction
	nil,                              // 5: protos.ChaincodeProposalPayload.TransientMapEntry
	(*ChaincodeID)(nil),              // 6: protos.ChaincodeID
	(*Response)(nil),                 // 7: protos.Response
}
var file_peer_proposal_proto_depIdxs = []int32{
	6, // 0: protos.ChaincodeHeaderExtension.chaincode_id:type_name -> protos.ChaincodeID
	5, // 1: protos.ChaincodeProposalPayload.TransientMap:type_name -> protos.ChaincodeProposalPayload.TransientMapEntry
	7, // 2: protos.ChaincodeAction.response:type_name -> protos.Response
	6, // 3: protos.ChaincodeAction.chaincode_id:type_name -> protos.ChaincodeID
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_peer_proposal_proto_init() }
func file_peer_proposal_proto_init() {
	if File_peer_proposal_proto != nil {
		return
	}
	file_peer_chaincode_proto_init()
	file_peer_proposal_response_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_peer_proposal_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignedProposal); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_peer_proposal_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Proposal); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_peer_proposal_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChaincodeHeaderExtension); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_peer_proposal_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChaincodeProposalPayload); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_peer_proposal_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChaincodeAction); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_peer_proposal_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_peer_proposal_proto_goTypes,
		DependencyIndexes: file_peer_proposal_proto_depIdxs,
		MessageInfos:      file_peer_proposal_proto_msgTypes,
	}.Build()
	File_peer_proposal_proto = out.File
	file_peer_proposal_proto_rawDesc = nil
	file_peer_proposal_proto_goTypes = nil
	file_peer_proposal_proto_depIdxs = nil
}