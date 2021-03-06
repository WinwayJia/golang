// Autogenerated by Thrift Compiler (0.9.2)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package rpc

import (
	"bytes"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

type ProtoRpcService interface {
	// Parameters:
	//  - Msg
	DealTwowayMessage(msg *ProtoRequest) (r *ProtoReply, err error)
	// Parameters:
	//  - Msg
	DealOnewayMessage(msg *ProtoRequest) (err error)
}

type ProtoRpcServiceClient struct {
	Transport       thrift.TTransport
	ProtocolFactory thrift.TProtocolFactory
	InputProtocol   thrift.TProtocol
	OutputProtocol  thrift.TProtocol
	SeqId           int32
}

func NewProtoRpcServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *ProtoRpcServiceClient {
	return &ProtoRpcServiceClient{Transport: t,
		ProtocolFactory: f,
		InputProtocol:   f.GetProtocol(t),
		OutputProtocol:  f.GetProtocol(t),
		SeqId:           0,
	}
}

func NewProtoRpcServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *ProtoRpcServiceClient {
	return &ProtoRpcServiceClient{Transport: t,
		ProtocolFactory: nil,
		InputProtocol:   iprot,
		OutputProtocol:  oprot,
		SeqId:           0,
	}
}

// Parameters:
//  - Msg
func (p *ProtoRpcServiceClient) DealTwowayMessage(msg *ProtoRequest) (r *ProtoReply, err error) {
	if err = p.sendDealTwowayMessage(msg); err != nil {
		return
	}
	return p.recvDealTwowayMessage()
}

func (p *ProtoRpcServiceClient) sendDealTwowayMessage(msg *ProtoRequest) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("dealTwowayMessage", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := DealTwowayMessageArgs{
		Msg: msg,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *ProtoRpcServiceClient) recvDealTwowayMessage() (value *ProtoReply, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	_, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error0 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error1 error
		error1, err = error0.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error1
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "dealTwowayMessage failed: out of sequence response")
		return
	}
	result := DealTwowayMessageResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	value = result.GetSuccess()
	return
}

// Parameters:
//  - Msg
func (p *ProtoRpcServiceClient) DealOnewayMessage(msg *ProtoRequest) (err error) {
	if err = p.sendDealOnewayMessage(msg); err != nil {
		return
	}
	return p.recvDealOnewayMessage()
}

func (p *ProtoRpcServiceClient) sendDealOnewayMessage(msg *ProtoRequest) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("dealOnewayMessage", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := DealOnewayMessageArgs{
		Msg: msg,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *ProtoRpcServiceClient) recvDealOnewayMessage() (err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	_, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error2 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error3 error
		error3, err = error2.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error3
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "dealOnewayMessage failed: out of sequence response")
		return
	}
	result := DealOnewayMessageResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	return
}

type ProtoRpcServiceProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      ProtoRpcService
}

func (p *ProtoRpcServiceProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *ProtoRpcServiceProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *ProtoRpcServiceProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewProtoRpcServiceProcessor(handler ProtoRpcService) *ProtoRpcServiceProcessor {

	self4 := &ProtoRpcServiceProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self4.processorMap["dealTwowayMessage"] = &protoRpcServiceProcessorDealTwowayMessage{handler: handler}
	self4.processorMap["dealOnewayMessage"] = &protoRpcServiceProcessorDealOnewayMessage{handler: handler}
	return self4
}

func (p *ProtoRpcServiceProcessor) Process(iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return false, err
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(seqId, iprot, oprot)
	}
	iprot.Skip(thrift.STRUCT)
	iprot.ReadMessageEnd()
	x5 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
	x5.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush()
	return false, x5

}

type protoRpcServiceProcessorDealTwowayMessage struct {
	handler ProtoRpcService
}

func (p *protoRpcServiceProcessorDealTwowayMessage) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := DealTwowayMessageArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("dealTwowayMessage", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := DealTwowayMessageResult{}
	var retval *ProtoReply
	var err2 error
	if retval, err2 = p.handler.DealTwowayMessage(args.Msg); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing dealTwowayMessage: "+err2.Error())
		oprot.WriteMessageBegin("dealTwowayMessage", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return true, err2
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("dealTwowayMessage", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type protoRpcServiceProcessorDealOnewayMessage struct {
	handler ProtoRpcService
}

func (p *protoRpcServiceProcessorDealOnewayMessage) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := DealOnewayMessageArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("dealOnewayMessage", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := DealOnewayMessageResult{}
	var err2 error
	if err2 = p.handler.DealOnewayMessage(args.Msg); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing dealOnewayMessage: "+err2.Error())
		oprot.WriteMessageBegin("dealOnewayMessage", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return true, err2
	}
	if err2 = oprot.WriteMessageBegin("dealOnewayMessage", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

// HELPER FUNCTIONS AND STRUCTURES

type DealTwowayMessageArgs struct {
	Msg *ProtoRequest `thrift:"msg,1" json:"msg"`
}

func NewDealTwowayMessageArgs() *DealTwowayMessageArgs {
	return &DealTwowayMessageArgs{}
}

var DealTwowayMessageArgs_Msg_DEFAULT *ProtoRequest

func (p *DealTwowayMessageArgs) GetMsg() *ProtoRequest {
	if !p.IsSetMsg() {
		return DealTwowayMessageArgs_Msg_DEFAULT
	}
	return p.Msg
}
func (p *DealTwowayMessageArgs) IsSetMsg() bool {
	return p.Msg != nil
}

func (p *DealTwowayMessageArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error: %s", p, err)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.ReadField1(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *DealTwowayMessageArgs) ReadField1(iprot thrift.TProtocol) error {
	p.Msg = &ProtoRequest{}
	if err := p.Msg.Read(iprot); err != nil {
		return fmt.Errorf("%T error reading struct: %s", p.Msg, err)
	}
	return nil
}

func (p *DealTwowayMessageArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("dealTwowayMessage_args"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("write struct stop error: %s", err)
	}
	return nil
}

func (p *DealTwowayMessageArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("msg", thrift.STRUCT, 1); err != nil {
		return fmt.Errorf("%T write field begin error 1:msg: %s", p, err)
	}
	if err := p.Msg.Write(oprot); err != nil {
		return fmt.Errorf("%T error writing struct: %s", p.Msg, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 1:msg: %s", p, err)
	}
	return err
}

func (p *DealTwowayMessageArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("DealTwowayMessageArgs(%+v)", *p)
}

type DealTwowayMessageResult struct {
	Success *ProtoReply `thrift:"success,0" json:"success"`
}

func NewDealTwowayMessageResult() *DealTwowayMessageResult {
	return &DealTwowayMessageResult{}
}

var DealTwowayMessageResult_Success_DEFAULT *ProtoReply

func (p *DealTwowayMessageResult) GetSuccess() *ProtoReply {
	if !p.IsSetSuccess() {
		return DealTwowayMessageResult_Success_DEFAULT
	}
	return p.Success
}
func (p *DealTwowayMessageResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *DealTwowayMessageResult) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error: %s", p, err)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if err := p.ReadField0(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *DealTwowayMessageResult) ReadField0(iprot thrift.TProtocol) error {
	p.Success = &ProtoReply{}
	if err := p.Success.Read(iprot); err != nil {
		return fmt.Errorf("%T error reading struct: %s", p.Success, err)
	}
	return nil
}

func (p *DealTwowayMessageResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("dealTwowayMessage_result"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField0(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("write struct stop error: %s", err)
	}
	return nil
}

func (p *DealTwowayMessageResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
			return fmt.Errorf("%T write field begin error 0:success: %s", p, err)
		}
		if err := p.Success.Write(oprot); err != nil {
			return fmt.Errorf("%T error writing struct: %s", p.Success, err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 0:success: %s", p, err)
		}
	}
	return err
}

func (p *DealTwowayMessageResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("DealTwowayMessageResult(%+v)", *p)
}

type DealOnewayMessageArgs struct {
	Msg *ProtoRequest `thrift:"msg,1" json:"msg"`
}

func NewDealOnewayMessageArgs() *DealOnewayMessageArgs {
	return &DealOnewayMessageArgs{}
}

var DealOnewayMessageArgs_Msg_DEFAULT *ProtoRequest

func (p *DealOnewayMessageArgs) GetMsg() *ProtoRequest {
	if !p.IsSetMsg() {
		return DealOnewayMessageArgs_Msg_DEFAULT
	}
	return p.Msg
}
func (p *DealOnewayMessageArgs) IsSetMsg() bool {
	return p.Msg != nil
}

func (p *DealOnewayMessageArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error: %s", p, err)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.ReadField1(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *DealOnewayMessageArgs) ReadField1(iprot thrift.TProtocol) error {
	p.Msg = &ProtoRequest{}
	if err := p.Msg.Read(iprot); err != nil {
		return fmt.Errorf("%T error reading struct: %s", p.Msg, err)
	}
	return nil
}

func (p *DealOnewayMessageArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("dealOnewayMessage_args"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("write struct stop error: %s", err)
	}
	return nil
}

func (p *DealOnewayMessageArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("msg", thrift.STRUCT, 1); err != nil {
		return fmt.Errorf("%T write field begin error 1:msg: %s", p, err)
	}
	if err := p.Msg.Write(oprot); err != nil {
		return fmt.Errorf("%T error writing struct: %s", p.Msg, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 1:msg: %s", p, err)
	}
	return err
}

func (p *DealOnewayMessageArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("DealOnewayMessageArgs(%+v)", *p)
}

type DealOnewayMessageResult struct {
}

func NewDealOnewayMessageResult() *DealOnewayMessageResult {
	return &DealOnewayMessageResult{}
}

func (p *DealOnewayMessageResult) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error: %s", p, err)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if err := iprot.Skip(fieldTypeId); err != nil {
			return err
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *DealOnewayMessageResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("dealOnewayMessage_result"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("write struct stop error: %s", err)
	}
	return nil
}

func (p *DealOnewayMessageResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("DealOnewayMessageResult(%+v)", *p)
}
