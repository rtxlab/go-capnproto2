package rpc

// AUTO GENERATED - DO NOT EDIT

import (
	strconv "strconv"
	capnp "zombiezen.com/go/capnproto2"
	text "zombiezen.com/go/capnproto2/encoding/text"
	schemas "zombiezen.com/go/capnproto2/schemas"
)

type Message struct{ capnp.Struct }
type Message_Which uint16

const (
	Message_Which_unimplemented  Message_Which = 0
	Message_Which_abort          Message_Which = 1
	Message_Which_bootstrap      Message_Which = 8
	Message_Which_call           Message_Which = 2
	Message_Which_return         Message_Which = 3
	Message_Which_finish         Message_Which = 4
	Message_Which_resolve        Message_Which = 5
	Message_Which_release        Message_Which = 6
	Message_Which_disembargo     Message_Which = 13
	Message_Which_obsoleteSave   Message_Which = 7
	Message_Which_obsoleteDelete Message_Which = 9
	Message_Which_provide        Message_Which = 10
	Message_Which_accept         Message_Which = 11
	Message_Which_join           Message_Which = 12
)

func (w Message_Which) String() string {
	const s = "unimplementedabortbootstrapcallreturnfinishresolvereleasedisembargoobsoleteSaveobsoleteDeleteprovideacceptjoin"
	switch w {
	case Message_Which_unimplemented:
		return s[0:13]
	case Message_Which_abort:
		return s[13:18]
	case Message_Which_bootstrap:
		return s[18:27]
	case Message_Which_call:
		return s[27:31]
	case Message_Which_return:
		return s[31:37]
	case Message_Which_finish:
		return s[37:43]
	case Message_Which_resolve:
		return s[43:50]
	case Message_Which_release:
		return s[50:57]
	case Message_Which_disembargo:
		return s[57:67]
	case Message_Which_obsoleteSave:
		return s[67:79]
	case Message_Which_obsoleteDelete:
		return s[79:93]
	case Message_Which_provide:
		return s[93:100]
	case Message_Which_accept:
		return s[100:106]
	case Message_Which_join:
		return s[106:110]

	}
	return "Message_Which(" + strconv.FormatUint(uint64(w), 10) + ")"
}

// Message_TypeID is the unique identifier for the type Message.
const Message_TypeID = 0x91b79f1f808db032

func NewMessage(s *capnp.Segment) (Message, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return Message{st}, err
}

func NewRootMessage(s *capnp.Segment) (Message, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return Message{st}, err
}

func ReadRootMessage(msg *capnp.Message) (Message, error) {
	root, err := msg.RootPtr()
	return Message{root.Struct()}, err
}

func (s Message) String() string {
	str, _ := text.Marshal(0x91b79f1f808db032, s.Struct)
	return str
}

func (s Message) Which() Message_Which {
	return Message_Which(s.Struct.Uint16(0))
}
func (s Message) Unimplemented() (Message, error) {
	p, err := s.Struct.Ptr(0)
	return Message{Struct: p.Struct()}, err
}

func (s Message) HasUnimplemented() bool {
	if s.Struct.Uint16(0) != 0 {
		return false
	}
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Message) SetUnimplemented(v Message) error {
	s.Struct.SetUint16(0, 0)
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewUnimplemented sets the unimplemented field to a newly
// allocated Message struct, preferring placement in s's segment.
func (s Message) NewUnimplemented() (Message, error) {
	s.Struct.SetUint16(0, 0)
	ss, err := NewMessage(s.Struct.Segment())
	if err != nil {
		return Message{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s Message) Abort() (Exception, error) {
	p, err := s.Struct.Ptr(0)
	return Exception{Struct: p.Struct()}, err
}

func (s Message) HasAbort() bool {
	if s.Struct.Uint16(0) != 1 {
		return false
	}
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Message) SetAbort(v Exception) error {
	s.Struct.SetUint16(0, 1)
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewAbort sets the abort field to a newly
// allocated Exception struct, preferring placement in s's segment.
func (s Message) NewAbort() (Exception, error) {
	s.Struct.SetUint16(0, 1)
	ss, err := NewException(s.Struct.Segment())
	if err != nil {
		return Exception{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s Message) Bootstrap() (Bootstrap, error) {
	p, err := s.Struct.Ptr(0)
	return Bootstrap{Struct: p.Struct()}, err
}

func (s Message) HasBootstrap() bool {
	if s.Struct.Uint16(0) != 8 {
		return false
	}
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Message) SetBootstrap(v Bootstrap) error {
	s.Struct.SetUint16(0, 8)
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewBootstrap sets the bootstrap field to a newly
// allocated Bootstrap struct, preferring placement in s's segment.
func (s Message) NewBootstrap() (Bootstrap, error) {
	s.Struct.SetUint16(0, 8)
	ss, err := NewBootstrap(s.Struct.Segment())
	if err != nil {
		return Bootstrap{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s Message) Call() (Call, error) {
	p, err := s.Struct.Ptr(0)
	return Call{Struct: p.Struct()}, err
}

func (s Message) HasCall() bool {
	if s.Struct.Uint16(0) != 2 {
		return false
	}
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Message) SetCall(v Call) error {
	s.Struct.SetUint16(0, 2)
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewCall sets the call field to a newly
// allocated Call struct, preferring placement in s's segment.
func (s Message) NewCall() (Call, error) {
	s.Struct.SetUint16(0, 2)
	ss, err := NewCall(s.Struct.Segment())
	if err != nil {
		return Call{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s Message) Return() (Return, error) {
	p, err := s.Struct.Ptr(0)
	return Return{Struct: p.Struct()}, err
}

func (s Message) HasReturn() bool {
	if s.Struct.Uint16(0) != 3 {
		return false
	}
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Message) SetReturn(v Return) error {
	s.Struct.SetUint16(0, 3)
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewReturn sets the return field to a newly
// allocated Return struct, preferring placement in s's segment.
func (s Message) NewReturn() (Return, error) {
	s.Struct.SetUint16(0, 3)
	ss, err := NewReturn(s.Struct.Segment())
	if err != nil {
		return Return{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s Message) Finish() (Finish, error) {
	p, err := s.Struct.Ptr(0)
	return Finish{Struct: p.Struct()}, err
}

func (s Message) HasFinish() bool {
	if s.Struct.Uint16(0) != 4 {
		return false
	}
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Message) SetFinish(v Finish) error {
	s.Struct.SetUint16(0, 4)
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewFinish sets the finish field to a newly
// allocated Finish struct, preferring placement in s's segment.
func (s Message) NewFinish() (Finish, error) {
	s.Struct.SetUint16(0, 4)
	ss, err := NewFinish(s.Struct.Segment())
	if err != nil {
		return Finish{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s Message) Resolve() (Resolve, error) {
	p, err := s.Struct.Ptr(0)
	return Resolve{Struct: p.Struct()}, err
}

func (s Message) HasResolve() bool {
	if s.Struct.Uint16(0) != 5 {
		return false
	}
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Message) SetResolve(v Resolve) error {
	s.Struct.SetUint16(0, 5)
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewResolve sets the resolve field to a newly
// allocated Resolve struct, preferring placement in s's segment.
func (s Message) NewResolve() (Resolve, error) {
	s.Struct.SetUint16(0, 5)
	ss, err := NewResolve(s.Struct.Segment())
	if err != nil {
		return Resolve{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s Message) Release() (Release, error) {
	p, err := s.Struct.Ptr(0)
	return Release{Struct: p.Struct()}, err
}

func (s Message) HasRelease() bool {
	if s.Struct.Uint16(0) != 6 {
		return false
	}
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Message) SetRelease(v Release) error {
	s.Struct.SetUint16(0, 6)
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewRelease sets the release field to a newly
// allocated Release struct, preferring placement in s's segment.
func (s Message) NewRelease() (Release, error) {
	s.Struct.SetUint16(0, 6)
	ss, err := NewRelease(s.Struct.Segment())
	if err != nil {
		return Release{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s Message) Disembargo() (Disembargo, error) {
	p, err := s.Struct.Ptr(0)
	return Disembargo{Struct: p.Struct()}, err
}

func (s Message) HasDisembargo() bool {
	if s.Struct.Uint16(0) != 13 {
		return false
	}
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Message) SetDisembargo(v Disembargo) error {
	s.Struct.SetUint16(0, 13)
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewDisembargo sets the disembargo field to a newly
// allocated Disembargo struct, preferring placement in s's segment.
func (s Message) NewDisembargo() (Disembargo, error) {
	s.Struct.SetUint16(0, 13)
	ss, err := NewDisembargo(s.Struct.Segment())
	if err != nil {
		return Disembargo{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s Message) ObsoleteSave() (capnp.Pointer, error) {
	return s.Struct.Pointer(0)
}

func (s Message) HasObsoleteSave() bool {
	if s.Struct.Uint16(0) != 7 {
		return false
	}
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Message) ObsoleteSavePtr() (capnp.Ptr, error) {
	return s.Struct.Ptr(0)
}

func (s Message) SetObsoleteSave(v capnp.Pointer) error {
	s.Struct.SetUint16(0, 7)
	return s.Struct.SetPointer(0, v)
}

func (s Message) SetObsoleteSavePtr(v capnp.Ptr) error {
	s.Struct.SetUint16(0, 7)
	return s.Struct.SetPtr(0, v)
}

func (s Message) ObsoleteDelete() (capnp.Pointer, error) {
	return s.Struct.Pointer(0)
}

func (s Message) HasObsoleteDelete() bool {
	if s.Struct.Uint16(0) != 9 {
		return false
	}
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Message) ObsoleteDeletePtr() (capnp.Ptr, error) {
	return s.Struct.Ptr(0)
}

func (s Message) SetObsoleteDelete(v capnp.Pointer) error {
	s.Struct.SetUint16(0, 9)
	return s.Struct.SetPointer(0, v)
}

func (s Message) SetObsoleteDeletePtr(v capnp.Ptr) error {
	s.Struct.SetUint16(0, 9)
	return s.Struct.SetPtr(0, v)
}

func (s Message) Provide() (Provide, error) {
	p, err := s.Struct.Ptr(0)
	return Provide{Struct: p.Struct()}, err
}

func (s Message) HasProvide() bool {
	if s.Struct.Uint16(0) != 10 {
		return false
	}
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Message) SetProvide(v Provide) error {
	s.Struct.SetUint16(0, 10)
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewProvide sets the provide field to a newly
// allocated Provide struct, preferring placement in s's segment.
func (s Message) NewProvide() (Provide, error) {
	s.Struct.SetUint16(0, 10)
	ss, err := NewProvide(s.Struct.Segment())
	if err != nil {
		return Provide{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s Message) Accept() (Accept, error) {
	p, err := s.Struct.Ptr(0)
	return Accept{Struct: p.Struct()}, err
}

func (s Message) HasAccept() bool {
	if s.Struct.Uint16(0) != 11 {
		return false
	}
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Message) SetAccept(v Accept) error {
	s.Struct.SetUint16(0, 11)
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewAccept sets the accept field to a newly
// allocated Accept struct, preferring placement in s's segment.
func (s Message) NewAccept() (Accept, error) {
	s.Struct.SetUint16(0, 11)
	ss, err := NewAccept(s.Struct.Segment())
	if err != nil {
		return Accept{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s Message) Join() (Join, error) {
	p, err := s.Struct.Ptr(0)
	return Join{Struct: p.Struct()}, err
}

func (s Message) HasJoin() bool {
	if s.Struct.Uint16(0) != 12 {
		return false
	}
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Message) SetJoin(v Join) error {
	s.Struct.SetUint16(0, 12)
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewJoin sets the join field to a newly
// allocated Join struct, preferring placement in s's segment.
func (s Message) NewJoin() (Join, error) {
	s.Struct.SetUint16(0, 12)
	ss, err := NewJoin(s.Struct.Segment())
	if err != nil {
		return Join{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

// Message_List is a list of Message.
type Message_List struct{ capnp.List }

// NewMessage creates a new list of Message.
func NewMessage_List(s *capnp.Segment, sz int32) (Message_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1}, sz)
	return Message_List{l}, err
}

func (s Message_List) At(i int) Message { return Message{s.List.Struct(i)} }

func (s Message_List) Set(i int, v Message) error { return s.List.SetStruct(i, v.Struct) }

// Message_Promise is a wrapper for a Message promised by a client call.
type Message_Promise struct{ *capnp.Pipeline }

func (p Message_Promise) Struct() (Message, error) {
	s, err := p.Pipeline.Struct()
	return Message{s}, err
}

func (p Message_Promise) Unimplemented() Message_Promise {
	return Message_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p Message_Promise) Abort() Exception_Promise {
	return Exception_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p Message_Promise) Bootstrap() Bootstrap_Promise {
	return Bootstrap_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p Message_Promise) Call() Call_Promise {
	return Call_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p Message_Promise) Return() Return_Promise {
	return Return_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p Message_Promise) Finish() Finish_Promise {
	return Finish_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p Message_Promise) Resolve() Resolve_Promise {
	return Resolve_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p Message_Promise) Release() Release_Promise {
	return Release_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p Message_Promise) Disembargo() Disembargo_Promise {
	return Disembargo_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p Message_Promise) ObsoleteSave() *capnp.Pipeline {
	return p.Pipeline.GetPipeline(0)
}

func (p Message_Promise) ObsoleteDelete() *capnp.Pipeline {
	return p.Pipeline.GetPipeline(0)
}

func (p Message_Promise) Provide() Provide_Promise {
	return Provide_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p Message_Promise) Accept() Accept_Promise {
	return Accept_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p Message_Promise) Join() Join_Promise {
	return Join_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

type Bootstrap struct{ capnp.Struct }

// Bootstrap_TypeID is the unique identifier for the type Bootstrap.
const Bootstrap_TypeID = 0xe94ccf8031176ec4

func NewBootstrap(s *capnp.Segment) (Bootstrap, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return Bootstrap{st}, err
}

func NewRootBootstrap(s *capnp.Segment) (Bootstrap, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return Bootstrap{st}, err
}

func ReadRootBootstrap(msg *capnp.Message) (Bootstrap, error) {
	root, err := msg.RootPtr()
	return Bootstrap{root.Struct()}, err
}

func (s Bootstrap) String() string {
	str, _ := text.Marshal(0xe94ccf8031176ec4, s.Struct)
	return str
}

func (s Bootstrap) QuestionId() uint32 {
	return s.Struct.Uint32(0)
}

func (s Bootstrap) SetQuestionId(v uint32) {
	s.Struct.SetUint32(0, v)
}

func (s Bootstrap) DeprecatedObjectId() (capnp.Pointer, error) {
	return s.Struct.Pointer(0)
}

func (s Bootstrap) HasDeprecatedObjectId() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Bootstrap) DeprecatedObjectIdPtr() (capnp.Ptr, error) {
	return s.Struct.Ptr(0)
}

func (s Bootstrap) SetDeprecatedObjectId(v capnp.Pointer) error {
	return s.Struct.SetPointer(0, v)
}

func (s Bootstrap) SetDeprecatedObjectIdPtr(v capnp.Ptr) error {
	return s.Struct.SetPtr(0, v)
}

// Bootstrap_List is a list of Bootstrap.
type Bootstrap_List struct{ capnp.List }

// NewBootstrap creates a new list of Bootstrap.
func NewBootstrap_List(s *capnp.Segment, sz int32) (Bootstrap_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1}, sz)
	return Bootstrap_List{l}, err
}

func (s Bootstrap_List) At(i int) Bootstrap { return Bootstrap{s.List.Struct(i)} }

func (s Bootstrap_List) Set(i int, v Bootstrap) error { return s.List.SetStruct(i, v.Struct) }

// Bootstrap_Promise is a wrapper for a Bootstrap promised by a client call.
type Bootstrap_Promise struct{ *capnp.Pipeline }

func (p Bootstrap_Promise) Struct() (Bootstrap, error) {
	s, err := p.Pipeline.Struct()
	return Bootstrap{s}, err
}

func (p Bootstrap_Promise) DeprecatedObjectId() *capnp.Pipeline {
	return p.Pipeline.GetPipeline(0)
}

type Call struct{ capnp.Struct }
type Call_sendResultsTo Call
type Call_sendResultsTo_Which uint16

const (
	Call_sendResultsTo_Which_caller     Call_sendResultsTo_Which = 0
	Call_sendResultsTo_Which_yourself   Call_sendResultsTo_Which = 1
	Call_sendResultsTo_Which_thirdParty Call_sendResultsTo_Which = 2
)

func (w Call_sendResultsTo_Which) String() string {
	const s = "calleryourselfthirdParty"
	switch w {
	case Call_sendResultsTo_Which_caller:
		return s[0:6]
	case Call_sendResultsTo_Which_yourself:
		return s[6:14]
	case Call_sendResultsTo_Which_thirdParty:
		return s[14:24]

	}
	return "Call_sendResultsTo_Which(" + strconv.FormatUint(uint64(w), 10) + ")"
}

// Call_TypeID is the unique identifier for the type Call.
const Call_TypeID = 0x836a53ce789d4cd4

func NewCall(s *capnp.Segment) (Call, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 24, PointerCount: 3})
	return Call{st}, err
}

func NewRootCall(s *capnp.Segment) (Call, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 24, PointerCount: 3})
	return Call{st}, err
}

func ReadRootCall(msg *capnp.Message) (Call, error) {
	root, err := msg.RootPtr()
	return Call{root.Struct()}, err
}

func (s Call) String() string {
	str, _ := text.Marshal(0x836a53ce789d4cd4, s.Struct)
	return str
}

func (s Call) QuestionId() uint32 {
	return s.Struct.Uint32(0)
}

func (s Call) SetQuestionId(v uint32) {
	s.Struct.SetUint32(0, v)
}

func (s Call) Target() (MessageTarget, error) {
	p, err := s.Struct.Ptr(0)
	return MessageTarget{Struct: p.Struct()}, err
}

func (s Call) HasTarget() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Call) SetTarget(v MessageTarget) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewTarget sets the target field to a newly
// allocated MessageTarget struct, preferring placement in s's segment.
func (s Call) NewTarget() (MessageTarget, error) {
	ss, err := NewMessageTarget(s.Struct.Segment())
	if err != nil {
		return MessageTarget{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s Call) InterfaceId() uint64 {
	return s.Struct.Uint64(8)
}

func (s Call) SetInterfaceId(v uint64) {
	s.Struct.SetUint64(8, v)
}

func (s Call) MethodId() uint16 {
	return s.Struct.Uint16(4)
}

func (s Call) SetMethodId(v uint16) {
	s.Struct.SetUint16(4, v)
}

func (s Call) AllowThirdPartyTailCall() bool {
	return s.Struct.Bit(128)
}

func (s Call) SetAllowThirdPartyTailCall(v bool) {
	s.Struct.SetBit(128, v)
}

func (s Call) Params() (Payload, error) {
	p, err := s.Struct.Ptr(1)
	return Payload{Struct: p.Struct()}, err
}

func (s Call) HasParams() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Call) SetParams(v Payload) error {
	return s.Struct.SetPtr(1, v.Struct.ToPtr())
}

// NewParams sets the params field to a newly
// allocated Payload struct, preferring placement in s's segment.
func (s Call) NewParams() (Payload, error) {
	ss, err := NewPayload(s.Struct.Segment())
	if err != nil {
		return Payload{}, err
	}
	err = s.Struct.SetPtr(1, ss.Struct.ToPtr())
	return ss, err
}

func (s Call) SendResultsTo() Call_sendResultsTo { return Call_sendResultsTo(s) }

func (s Call_sendResultsTo) Which() Call_sendResultsTo_Which {
	return Call_sendResultsTo_Which(s.Struct.Uint16(6))
}
func (s Call_sendResultsTo) SetCaller() {
	s.Struct.SetUint16(6, 0)

}

func (s Call_sendResultsTo) SetYourself() {
	s.Struct.SetUint16(6, 1)

}

func (s Call_sendResultsTo) ThirdParty() (capnp.Pointer, error) {
	return s.Struct.Pointer(2)
}

func (s Call_sendResultsTo) HasThirdParty() bool {
	if s.Struct.Uint16(6) != 2 {
		return false
	}
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s Call_sendResultsTo) ThirdPartyPtr() (capnp.Ptr, error) {
	return s.Struct.Ptr(2)
}

func (s Call_sendResultsTo) SetThirdParty(v capnp.Pointer) error {
	s.Struct.SetUint16(6, 2)
	return s.Struct.SetPointer(2, v)
}

func (s Call_sendResultsTo) SetThirdPartyPtr(v capnp.Ptr) error {
	s.Struct.SetUint16(6, 2)
	return s.Struct.SetPtr(2, v)
}

// Call_List is a list of Call.
type Call_List struct{ capnp.List }

// NewCall creates a new list of Call.
func NewCall_List(s *capnp.Segment, sz int32) (Call_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 24, PointerCount: 3}, sz)
	return Call_List{l}, err
}

func (s Call_List) At(i int) Call { return Call{s.List.Struct(i)} }

func (s Call_List) Set(i int, v Call) error { return s.List.SetStruct(i, v.Struct) }

// Call_Promise is a wrapper for a Call promised by a client call.
type Call_Promise struct{ *capnp.Pipeline }

func (p Call_Promise) Struct() (Call, error) {
	s, err := p.Pipeline.Struct()
	return Call{s}, err
}

func (p Call_Promise) Target() MessageTarget_Promise {
	return MessageTarget_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p Call_Promise) Params() Payload_Promise {
	return Payload_Promise{Pipeline: p.Pipeline.GetPipeline(1)}
}

func (p Call_Promise) SendResultsTo() Call_sendResultsTo_Promise {
	return Call_sendResultsTo_Promise{p.Pipeline}
}

// Call_sendResultsTo_Promise is a wrapper for a Call_sendResultsTo promised by a client call.
type Call_sendResultsTo_Promise struct{ *capnp.Pipeline }

func (p Call_sendResultsTo_Promise) Struct() (Call_sendResultsTo, error) {
	s, err := p.Pipeline.Struct()
	return Call_sendResultsTo{s}, err
}

func (p Call_sendResultsTo_Promise) ThirdParty() *capnp.Pipeline {
	return p.Pipeline.GetPipeline(2)
}

type Return struct{ capnp.Struct }
type Return_Which uint16

const (
	Return_Which_results               Return_Which = 0
	Return_Which_exception             Return_Which = 1
	Return_Which_canceled              Return_Which = 2
	Return_Which_resultsSentElsewhere  Return_Which = 3
	Return_Which_takeFromOtherQuestion Return_Which = 4
	Return_Which_acceptFromThirdParty  Return_Which = 5
)

func (w Return_Which) String() string {
	const s = "resultsexceptioncanceledresultsSentElsewheretakeFromOtherQuestionacceptFromThirdParty"
	switch w {
	case Return_Which_results:
		return s[0:7]
	case Return_Which_exception:
		return s[7:16]
	case Return_Which_canceled:
		return s[16:24]
	case Return_Which_resultsSentElsewhere:
		return s[24:44]
	case Return_Which_takeFromOtherQuestion:
		return s[44:65]
	case Return_Which_acceptFromThirdParty:
		return s[65:85]

	}
	return "Return_Which(" + strconv.FormatUint(uint64(w), 10) + ")"
}

// Return_TypeID is the unique identifier for the type Return.
const Return_TypeID = 0x9e19b28d3db3573a

func NewReturn(s *capnp.Segment) (Return, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 1})
	return Return{st}, err
}

func NewRootReturn(s *capnp.Segment) (Return, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 1})
	return Return{st}, err
}

func ReadRootReturn(msg *capnp.Message) (Return, error) {
	root, err := msg.RootPtr()
	return Return{root.Struct()}, err
}

func (s Return) String() string {
	str, _ := text.Marshal(0x9e19b28d3db3573a, s.Struct)
	return str
}

func (s Return) Which() Return_Which {
	return Return_Which(s.Struct.Uint16(6))
}
func (s Return) AnswerId() uint32 {
	return s.Struct.Uint32(0)
}

func (s Return) SetAnswerId(v uint32) {
	s.Struct.SetUint32(0, v)
}

func (s Return) ReleaseParamCaps() bool {
	return !s.Struct.Bit(32)
}

func (s Return) SetReleaseParamCaps(v bool) {
	s.Struct.SetBit(32, !v)
}

func (s Return) Results() (Payload, error) {
	p, err := s.Struct.Ptr(0)
	return Payload{Struct: p.Struct()}, err
}

func (s Return) HasResults() bool {
	if s.Struct.Uint16(6) != 0 {
		return false
	}
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Return) SetResults(v Payload) error {
	s.Struct.SetUint16(6, 0)
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewResults sets the results field to a newly
// allocated Payload struct, preferring placement in s's segment.
func (s Return) NewResults() (Payload, error) {
	s.Struct.SetUint16(6, 0)
	ss, err := NewPayload(s.Struct.Segment())
	if err != nil {
		return Payload{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s Return) Exception() (Exception, error) {
	p, err := s.Struct.Ptr(0)
	return Exception{Struct: p.Struct()}, err
}

func (s Return) HasException() bool {
	if s.Struct.Uint16(6) != 1 {
		return false
	}
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Return) SetException(v Exception) error {
	s.Struct.SetUint16(6, 1)
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewException sets the exception field to a newly
// allocated Exception struct, preferring placement in s's segment.
func (s Return) NewException() (Exception, error) {
	s.Struct.SetUint16(6, 1)
	ss, err := NewException(s.Struct.Segment())
	if err != nil {
		return Exception{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s Return) SetCanceled() {
	s.Struct.SetUint16(6, 2)

}

func (s Return) SetResultsSentElsewhere() {
	s.Struct.SetUint16(6, 3)

}

func (s Return) TakeFromOtherQuestion() uint32 {
	return s.Struct.Uint32(8)
}

func (s Return) SetTakeFromOtherQuestion(v uint32) {
	s.Struct.SetUint16(6, 4)
	s.Struct.SetUint32(8, v)
}

func (s Return) AcceptFromThirdParty() (capnp.Pointer, error) {
	return s.Struct.Pointer(0)
}

func (s Return) HasAcceptFromThirdParty() bool {
	if s.Struct.Uint16(6) != 5 {
		return false
	}
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Return) AcceptFromThirdPartyPtr() (capnp.Ptr, error) {
	return s.Struct.Ptr(0)
}

func (s Return) SetAcceptFromThirdParty(v capnp.Pointer) error {
	s.Struct.SetUint16(6, 5)
	return s.Struct.SetPointer(0, v)
}

func (s Return) SetAcceptFromThirdPartyPtr(v capnp.Ptr) error {
	s.Struct.SetUint16(6, 5)
	return s.Struct.SetPtr(0, v)
}

// Return_List is a list of Return.
type Return_List struct{ capnp.List }

// NewReturn creates a new list of Return.
func NewReturn_List(s *capnp.Segment, sz int32) (Return_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 16, PointerCount: 1}, sz)
	return Return_List{l}, err
}

func (s Return_List) At(i int) Return { return Return{s.List.Struct(i)} }

func (s Return_List) Set(i int, v Return) error { return s.List.SetStruct(i, v.Struct) }

// Return_Promise is a wrapper for a Return promised by a client call.
type Return_Promise struct{ *capnp.Pipeline }

func (p Return_Promise) Struct() (Return, error) {
	s, err := p.Pipeline.Struct()
	return Return{s}, err
}

func (p Return_Promise) Results() Payload_Promise {
	return Payload_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p Return_Promise) Exception() Exception_Promise {
	return Exception_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p Return_Promise) AcceptFromThirdParty() *capnp.Pipeline {
	return p.Pipeline.GetPipeline(0)
}

type Finish struct{ capnp.Struct }

// Finish_TypeID is the unique identifier for the type Finish.
const Finish_TypeID = 0xd37d2eb2c2f80e63

func NewFinish(s *capnp.Segment) (Finish, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return Finish{st}, err
}

func NewRootFinish(s *capnp.Segment) (Finish, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return Finish{st}, err
}

func ReadRootFinish(msg *capnp.Message) (Finish, error) {
	root, err := msg.RootPtr()
	return Finish{root.Struct()}, err
}

func (s Finish) String() string {
	str, _ := text.Marshal(0xd37d2eb2c2f80e63, s.Struct)
	return str
}

func (s Finish) QuestionId() uint32 {
	return s.Struct.Uint32(0)
}

func (s Finish) SetQuestionId(v uint32) {
	s.Struct.SetUint32(0, v)
}

func (s Finish) ReleaseResultCaps() bool {
	return !s.Struct.Bit(32)
}

func (s Finish) SetReleaseResultCaps(v bool) {
	s.Struct.SetBit(32, !v)
}

// Finish_List is a list of Finish.
type Finish_List struct{ capnp.List }

// NewFinish creates a new list of Finish.
func NewFinish_List(s *capnp.Segment, sz int32) (Finish_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0}, sz)
	return Finish_List{l}, err
}

func (s Finish_List) At(i int) Finish { return Finish{s.List.Struct(i)} }

func (s Finish_List) Set(i int, v Finish) error { return s.List.SetStruct(i, v.Struct) }

// Finish_Promise is a wrapper for a Finish promised by a client call.
type Finish_Promise struct{ *capnp.Pipeline }

func (p Finish_Promise) Struct() (Finish, error) {
	s, err := p.Pipeline.Struct()
	return Finish{s}, err
}

type Resolve struct{ capnp.Struct }
type Resolve_Which uint16

const (
	Resolve_Which_cap       Resolve_Which = 0
	Resolve_Which_exception Resolve_Which = 1
)

func (w Resolve_Which) String() string {
	const s = "capexception"
	switch w {
	case Resolve_Which_cap:
		return s[0:3]
	case Resolve_Which_exception:
		return s[3:12]

	}
	return "Resolve_Which(" + strconv.FormatUint(uint64(w), 10) + ")"
}

// Resolve_TypeID is the unique identifier for the type Resolve.
const Resolve_TypeID = 0xbbc29655fa89086e

func NewResolve(s *capnp.Segment) (Resolve, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return Resolve{st}, err
}

func NewRootResolve(s *capnp.Segment) (Resolve, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return Resolve{st}, err
}

func ReadRootResolve(msg *capnp.Message) (Resolve, error) {
	root, err := msg.RootPtr()
	return Resolve{root.Struct()}, err
}

func (s Resolve) String() string {
	str, _ := text.Marshal(0xbbc29655fa89086e, s.Struct)
	return str
}

func (s Resolve) Which() Resolve_Which {
	return Resolve_Which(s.Struct.Uint16(4))
}
func (s Resolve) PromiseId() uint32 {
	return s.Struct.Uint32(0)
}

func (s Resolve) SetPromiseId(v uint32) {
	s.Struct.SetUint32(0, v)
}

func (s Resolve) Cap() (CapDescriptor, error) {
	p, err := s.Struct.Ptr(0)
	return CapDescriptor{Struct: p.Struct()}, err
}

func (s Resolve) HasCap() bool {
	if s.Struct.Uint16(4) != 0 {
		return false
	}
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Resolve) SetCap(v CapDescriptor) error {
	s.Struct.SetUint16(4, 0)
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewCap sets the cap field to a newly
// allocated CapDescriptor struct, preferring placement in s's segment.
func (s Resolve) NewCap() (CapDescriptor, error) {
	s.Struct.SetUint16(4, 0)
	ss, err := NewCapDescriptor(s.Struct.Segment())
	if err != nil {
		return CapDescriptor{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s Resolve) Exception() (Exception, error) {
	p, err := s.Struct.Ptr(0)
	return Exception{Struct: p.Struct()}, err
}

func (s Resolve) HasException() bool {
	if s.Struct.Uint16(4) != 1 {
		return false
	}
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Resolve) SetException(v Exception) error {
	s.Struct.SetUint16(4, 1)
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewException sets the exception field to a newly
// allocated Exception struct, preferring placement in s's segment.
func (s Resolve) NewException() (Exception, error) {
	s.Struct.SetUint16(4, 1)
	ss, err := NewException(s.Struct.Segment())
	if err != nil {
		return Exception{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

// Resolve_List is a list of Resolve.
type Resolve_List struct{ capnp.List }

// NewResolve creates a new list of Resolve.
func NewResolve_List(s *capnp.Segment, sz int32) (Resolve_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1}, sz)
	return Resolve_List{l}, err
}

func (s Resolve_List) At(i int) Resolve { return Resolve{s.List.Struct(i)} }

func (s Resolve_List) Set(i int, v Resolve) error { return s.List.SetStruct(i, v.Struct) }

// Resolve_Promise is a wrapper for a Resolve promised by a client call.
type Resolve_Promise struct{ *capnp.Pipeline }

func (p Resolve_Promise) Struct() (Resolve, error) {
	s, err := p.Pipeline.Struct()
	return Resolve{s}, err
}

func (p Resolve_Promise) Cap() CapDescriptor_Promise {
	return CapDescriptor_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p Resolve_Promise) Exception() Exception_Promise {
	return Exception_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

type Release struct{ capnp.Struct }

// Release_TypeID is the unique identifier for the type Release.
const Release_TypeID = 0xad1a6c0d7dd07497

func NewRelease(s *capnp.Segment) (Release, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return Release{st}, err
}

func NewRootRelease(s *capnp.Segment) (Release, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return Release{st}, err
}

func ReadRootRelease(msg *capnp.Message) (Release, error) {
	root, err := msg.RootPtr()
	return Release{root.Struct()}, err
}

func (s Release) String() string {
	str, _ := text.Marshal(0xad1a6c0d7dd07497, s.Struct)
	return str
}

func (s Release) Id() uint32 {
	return s.Struct.Uint32(0)
}

func (s Release) SetId(v uint32) {
	s.Struct.SetUint32(0, v)
}

func (s Release) ReferenceCount() uint32 {
	return s.Struct.Uint32(4)
}

func (s Release) SetReferenceCount(v uint32) {
	s.Struct.SetUint32(4, v)
}

// Release_List is a list of Release.
type Release_List struct{ capnp.List }

// NewRelease creates a new list of Release.
func NewRelease_List(s *capnp.Segment, sz int32) (Release_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0}, sz)
	return Release_List{l}, err
}

func (s Release_List) At(i int) Release { return Release{s.List.Struct(i)} }

func (s Release_List) Set(i int, v Release) error { return s.List.SetStruct(i, v.Struct) }

// Release_Promise is a wrapper for a Release promised by a client call.
type Release_Promise struct{ *capnp.Pipeline }

func (p Release_Promise) Struct() (Release, error) {
	s, err := p.Pipeline.Struct()
	return Release{s}, err
}

type Disembargo struct{ capnp.Struct }
type Disembargo_context Disembargo
type Disembargo_context_Which uint16

const (
	Disembargo_context_Which_senderLoopback   Disembargo_context_Which = 0
	Disembargo_context_Which_receiverLoopback Disembargo_context_Which = 1
	Disembargo_context_Which_accept           Disembargo_context_Which = 2
	Disembargo_context_Which_provide          Disembargo_context_Which = 3
)

func (w Disembargo_context_Which) String() string {
	const s = "senderLoopbackreceiverLoopbackacceptprovide"
	switch w {
	case Disembargo_context_Which_senderLoopback:
		return s[0:14]
	case Disembargo_context_Which_receiverLoopback:
		return s[14:30]
	case Disembargo_context_Which_accept:
		return s[30:36]
	case Disembargo_context_Which_provide:
		return s[36:43]

	}
	return "Disembargo_context_Which(" + strconv.FormatUint(uint64(w), 10) + ")"
}

// Disembargo_TypeID is the unique identifier for the type Disembargo.
const Disembargo_TypeID = 0xf964368b0fbd3711

func NewDisembargo(s *capnp.Segment) (Disembargo, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return Disembargo{st}, err
}

func NewRootDisembargo(s *capnp.Segment) (Disembargo, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return Disembargo{st}, err
}

func ReadRootDisembargo(msg *capnp.Message) (Disembargo, error) {
	root, err := msg.RootPtr()
	return Disembargo{root.Struct()}, err
}

func (s Disembargo) String() string {
	str, _ := text.Marshal(0xf964368b0fbd3711, s.Struct)
	return str
}

func (s Disembargo) Target() (MessageTarget, error) {
	p, err := s.Struct.Ptr(0)
	return MessageTarget{Struct: p.Struct()}, err
}

func (s Disembargo) HasTarget() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Disembargo) SetTarget(v MessageTarget) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewTarget sets the target field to a newly
// allocated MessageTarget struct, preferring placement in s's segment.
func (s Disembargo) NewTarget() (MessageTarget, error) {
	ss, err := NewMessageTarget(s.Struct.Segment())
	if err != nil {
		return MessageTarget{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s Disembargo) Context() Disembargo_context { return Disembargo_context(s) }

func (s Disembargo_context) Which() Disembargo_context_Which {
	return Disembargo_context_Which(s.Struct.Uint16(4))
}
func (s Disembargo_context) SenderLoopback() uint32 {
	return s.Struct.Uint32(0)
}

func (s Disembargo_context) SetSenderLoopback(v uint32) {
	s.Struct.SetUint16(4, 0)
	s.Struct.SetUint32(0, v)
}

func (s Disembargo_context) ReceiverLoopback() uint32 {
	return s.Struct.Uint32(0)
}

func (s Disembargo_context) SetReceiverLoopback(v uint32) {
	s.Struct.SetUint16(4, 1)
	s.Struct.SetUint32(0, v)
}

func (s Disembargo_context) SetAccept() {
	s.Struct.SetUint16(4, 2)

}

func (s Disembargo_context) Provide() uint32 {
	return s.Struct.Uint32(0)
}

func (s Disembargo_context) SetProvide(v uint32) {
	s.Struct.SetUint16(4, 3)
	s.Struct.SetUint32(0, v)
}

// Disembargo_List is a list of Disembargo.
type Disembargo_List struct{ capnp.List }

// NewDisembargo creates a new list of Disembargo.
func NewDisembargo_List(s *capnp.Segment, sz int32) (Disembargo_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1}, sz)
	return Disembargo_List{l}, err
}

func (s Disembargo_List) At(i int) Disembargo { return Disembargo{s.List.Struct(i)} }

func (s Disembargo_List) Set(i int, v Disembargo) error { return s.List.SetStruct(i, v.Struct) }

// Disembargo_Promise is a wrapper for a Disembargo promised by a client call.
type Disembargo_Promise struct{ *capnp.Pipeline }

func (p Disembargo_Promise) Struct() (Disembargo, error) {
	s, err := p.Pipeline.Struct()
	return Disembargo{s}, err
}

func (p Disembargo_Promise) Target() MessageTarget_Promise {
	return MessageTarget_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p Disembargo_Promise) Context() Disembargo_context_Promise {
	return Disembargo_context_Promise{p.Pipeline}
}

// Disembargo_context_Promise is a wrapper for a Disembargo_context promised by a client call.
type Disembargo_context_Promise struct{ *capnp.Pipeline }

func (p Disembargo_context_Promise) Struct() (Disembargo_context, error) {
	s, err := p.Pipeline.Struct()
	return Disembargo_context{s}, err
}

type Provide struct{ capnp.Struct }

// Provide_TypeID is the unique identifier for the type Provide.
const Provide_TypeID = 0x9c6a046bfbc1ac5a

func NewProvide(s *capnp.Segment) (Provide, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	return Provide{st}, err
}

func NewRootProvide(s *capnp.Segment) (Provide, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	return Provide{st}, err
}

func ReadRootProvide(msg *capnp.Message) (Provide, error) {
	root, err := msg.RootPtr()
	return Provide{root.Struct()}, err
}

func (s Provide) String() string {
	str, _ := text.Marshal(0x9c6a046bfbc1ac5a, s.Struct)
	return str
}

func (s Provide) QuestionId() uint32 {
	return s.Struct.Uint32(0)
}

func (s Provide) SetQuestionId(v uint32) {
	s.Struct.SetUint32(0, v)
}

func (s Provide) Target() (MessageTarget, error) {
	p, err := s.Struct.Ptr(0)
	return MessageTarget{Struct: p.Struct()}, err
}

func (s Provide) HasTarget() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Provide) SetTarget(v MessageTarget) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewTarget sets the target field to a newly
// allocated MessageTarget struct, preferring placement in s's segment.
func (s Provide) NewTarget() (MessageTarget, error) {
	ss, err := NewMessageTarget(s.Struct.Segment())
	if err != nil {
		return MessageTarget{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s Provide) Recipient() (capnp.Pointer, error) {
	return s.Struct.Pointer(1)
}

func (s Provide) HasRecipient() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Provide) RecipientPtr() (capnp.Ptr, error) {
	return s.Struct.Ptr(1)
}

func (s Provide) SetRecipient(v capnp.Pointer) error {
	return s.Struct.SetPointer(1, v)
}

func (s Provide) SetRecipientPtr(v capnp.Ptr) error {
	return s.Struct.SetPtr(1, v)
}

// Provide_List is a list of Provide.
type Provide_List struct{ capnp.List }

// NewProvide creates a new list of Provide.
func NewProvide_List(s *capnp.Segment, sz int32) (Provide_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2}, sz)
	return Provide_List{l}, err
}

func (s Provide_List) At(i int) Provide { return Provide{s.List.Struct(i)} }

func (s Provide_List) Set(i int, v Provide) error { return s.List.SetStruct(i, v.Struct) }

// Provide_Promise is a wrapper for a Provide promised by a client call.
type Provide_Promise struct{ *capnp.Pipeline }

func (p Provide_Promise) Struct() (Provide, error) {
	s, err := p.Pipeline.Struct()
	return Provide{s}, err
}

func (p Provide_Promise) Target() MessageTarget_Promise {
	return MessageTarget_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p Provide_Promise) Recipient() *capnp.Pipeline {
	return p.Pipeline.GetPipeline(1)
}

type Accept struct{ capnp.Struct }

// Accept_TypeID is the unique identifier for the type Accept.
const Accept_TypeID = 0xd4c9b56290554016

func NewAccept(s *capnp.Segment) (Accept, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return Accept{st}, err
}

func NewRootAccept(s *capnp.Segment) (Accept, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return Accept{st}, err
}

func ReadRootAccept(msg *capnp.Message) (Accept, error) {
	root, err := msg.RootPtr()
	return Accept{root.Struct()}, err
}

func (s Accept) String() string {
	str, _ := text.Marshal(0xd4c9b56290554016, s.Struct)
	return str
}

func (s Accept) QuestionId() uint32 {
	return s.Struct.Uint32(0)
}

func (s Accept) SetQuestionId(v uint32) {
	s.Struct.SetUint32(0, v)
}

func (s Accept) Provision() (capnp.Pointer, error) {
	return s.Struct.Pointer(0)
}

func (s Accept) HasProvision() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Accept) ProvisionPtr() (capnp.Ptr, error) {
	return s.Struct.Ptr(0)
}

func (s Accept) SetProvision(v capnp.Pointer) error {
	return s.Struct.SetPointer(0, v)
}

func (s Accept) SetProvisionPtr(v capnp.Ptr) error {
	return s.Struct.SetPtr(0, v)
}

func (s Accept) Embargo() bool {
	return s.Struct.Bit(32)
}

func (s Accept) SetEmbargo(v bool) {
	s.Struct.SetBit(32, v)
}

// Accept_List is a list of Accept.
type Accept_List struct{ capnp.List }

// NewAccept creates a new list of Accept.
func NewAccept_List(s *capnp.Segment, sz int32) (Accept_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1}, sz)
	return Accept_List{l}, err
}

func (s Accept_List) At(i int) Accept { return Accept{s.List.Struct(i)} }

func (s Accept_List) Set(i int, v Accept) error { return s.List.SetStruct(i, v.Struct) }

// Accept_Promise is a wrapper for a Accept promised by a client call.
type Accept_Promise struct{ *capnp.Pipeline }

func (p Accept_Promise) Struct() (Accept, error) {
	s, err := p.Pipeline.Struct()
	return Accept{s}, err
}

func (p Accept_Promise) Provision() *capnp.Pipeline {
	return p.Pipeline.GetPipeline(0)
}

type Join struct{ capnp.Struct }

// Join_TypeID is the unique identifier for the type Join.
const Join_TypeID = 0xfbe1980490e001af

func NewJoin(s *capnp.Segment) (Join, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	return Join{st}, err
}

func NewRootJoin(s *capnp.Segment) (Join, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	return Join{st}, err
}

func ReadRootJoin(msg *capnp.Message) (Join, error) {
	root, err := msg.RootPtr()
	return Join{root.Struct()}, err
}

func (s Join) String() string {
	str, _ := text.Marshal(0xfbe1980490e001af, s.Struct)
	return str
}

func (s Join) QuestionId() uint32 {
	return s.Struct.Uint32(0)
}

func (s Join) SetQuestionId(v uint32) {
	s.Struct.SetUint32(0, v)
}

func (s Join) Target() (MessageTarget, error) {
	p, err := s.Struct.Ptr(0)
	return MessageTarget{Struct: p.Struct()}, err
}

func (s Join) HasTarget() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Join) SetTarget(v MessageTarget) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewTarget sets the target field to a newly
// allocated MessageTarget struct, preferring placement in s's segment.
func (s Join) NewTarget() (MessageTarget, error) {
	ss, err := NewMessageTarget(s.Struct.Segment())
	if err != nil {
		return MessageTarget{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s Join) KeyPart() (capnp.Pointer, error) {
	return s.Struct.Pointer(1)
}

func (s Join) HasKeyPart() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Join) KeyPartPtr() (capnp.Ptr, error) {
	return s.Struct.Ptr(1)
}

func (s Join) SetKeyPart(v capnp.Pointer) error {
	return s.Struct.SetPointer(1, v)
}

func (s Join) SetKeyPartPtr(v capnp.Ptr) error {
	return s.Struct.SetPtr(1, v)
}

// Join_List is a list of Join.
type Join_List struct{ capnp.List }

// NewJoin creates a new list of Join.
func NewJoin_List(s *capnp.Segment, sz int32) (Join_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2}, sz)
	return Join_List{l}, err
}

func (s Join_List) At(i int) Join { return Join{s.List.Struct(i)} }

func (s Join_List) Set(i int, v Join) error { return s.List.SetStruct(i, v.Struct) }

// Join_Promise is a wrapper for a Join promised by a client call.
type Join_Promise struct{ *capnp.Pipeline }

func (p Join_Promise) Struct() (Join, error) {
	s, err := p.Pipeline.Struct()
	return Join{s}, err
}

func (p Join_Promise) Target() MessageTarget_Promise {
	return MessageTarget_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p Join_Promise) KeyPart() *capnp.Pipeline {
	return p.Pipeline.GetPipeline(1)
}

type MessageTarget struct{ capnp.Struct }
type MessageTarget_Which uint16

const (
	MessageTarget_Which_importedCap    MessageTarget_Which = 0
	MessageTarget_Which_promisedAnswer MessageTarget_Which = 1
)

func (w MessageTarget_Which) String() string {
	const s = "importedCappromisedAnswer"
	switch w {
	case MessageTarget_Which_importedCap:
		return s[0:11]
	case MessageTarget_Which_promisedAnswer:
		return s[11:25]

	}
	return "MessageTarget_Which(" + strconv.FormatUint(uint64(w), 10) + ")"
}

// MessageTarget_TypeID is the unique identifier for the type MessageTarget.
const MessageTarget_TypeID = 0x95bc14545813fbc1

func NewMessageTarget(s *capnp.Segment) (MessageTarget, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return MessageTarget{st}, err
}

func NewRootMessageTarget(s *capnp.Segment) (MessageTarget, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return MessageTarget{st}, err
}

func ReadRootMessageTarget(msg *capnp.Message) (MessageTarget, error) {
	root, err := msg.RootPtr()
	return MessageTarget{root.Struct()}, err
}

func (s MessageTarget) String() string {
	str, _ := text.Marshal(0x95bc14545813fbc1, s.Struct)
	return str
}

func (s MessageTarget) Which() MessageTarget_Which {
	return MessageTarget_Which(s.Struct.Uint16(4))
}
func (s MessageTarget) ImportedCap() uint32 {
	return s.Struct.Uint32(0)
}

func (s MessageTarget) SetImportedCap(v uint32) {
	s.Struct.SetUint16(4, 0)
	s.Struct.SetUint32(0, v)
}

func (s MessageTarget) PromisedAnswer() (PromisedAnswer, error) {
	p, err := s.Struct.Ptr(0)
	return PromisedAnswer{Struct: p.Struct()}, err
}

func (s MessageTarget) HasPromisedAnswer() bool {
	if s.Struct.Uint16(4) != 1 {
		return false
	}
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s MessageTarget) SetPromisedAnswer(v PromisedAnswer) error {
	s.Struct.SetUint16(4, 1)
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewPromisedAnswer sets the promisedAnswer field to a newly
// allocated PromisedAnswer struct, preferring placement in s's segment.
func (s MessageTarget) NewPromisedAnswer() (PromisedAnswer, error) {
	s.Struct.SetUint16(4, 1)
	ss, err := NewPromisedAnswer(s.Struct.Segment())
	if err != nil {
		return PromisedAnswer{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

// MessageTarget_List is a list of MessageTarget.
type MessageTarget_List struct{ capnp.List }

// NewMessageTarget creates a new list of MessageTarget.
func NewMessageTarget_List(s *capnp.Segment, sz int32) (MessageTarget_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1}, sz)
	return MessageTarget_List{l}, err
}

func (s MessageTarget_List) At(i int) MessageTarget { return MessageTarget{s.List.Struct(i)} }

func (s MessageTarget_List) Set(i int, v MessageTarget) error { return s.List.SetStruct(i, v.Struct) }

// MessageTarget_Promise is a wrapper for a MessageTarget promised by a client call.
type MessageTarget_Promise struct{ *capnp.Pipeline }

func (p MessageTarget_Promise) Struct() (MessageTarget, error) {
	s, err := p.Pipeline.Struct()
	return MessageTarget{s}, err
}

func (p MessageTarget_Promise) PromisedAnswer() PromisedAnswer_Promise {
	return PromisedAnswer_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

type Payload struct{ capnp.Struct }

// Payload_TypeID is the unique identifier for the type Payload.
const Payload_TypeID = 0x9a0e61223d96743b

func NewPayload(s *capnp.Segment) (Payload, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return Payload{st}, err
}

func NewRootPayload(s *capnp.Segment) (Payload, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return Payload{st}, err
}

func ReadRootPayload(msg *capnp.Message) (Payload, error) {
	root, err := msg.RootPtr()
	return Payload{root.Struct()}, err
}

func (s Payload) String() string {
	str, _ := text.Marshal(0x9a0e61223d96743b, s.Struct)
	return str
}

func (s Payload) Content() (capnp.Pointer, error) {
	return s.Struct.Pointer(0)
}

func (s Payload) HasContent() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Payload) ContentPtr() (capnp.Ptr, error) {
	return s.Struct.Ptr(0)
}

func (s Payload) SetContent(v capnp.Pointer) error {
	return s.Struct.SetPointer(0, v)
}

func (s Payload) SetContentPtr(v capnp.Ptr) error {
	return s.Struct.SetPtr(0, v)
}

func (s Payload) CapTable() (CapDescriptor_List, error) {
	p, err := s.Struct.Ptr(1)
	return CapDescriptor_List{List: p.List()}, err
}

func (s Payload) HasCapTable() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Payload) SetCapTable(v CapDescriptor_List) error {
	return s.Struct.SetPtr(1, v.List.ToPtr())
}

// NewCapTable sets the capTable field to a newly
// allocated CapDescriptor_List, preferring placement in s's segment.
func (s Payload) NewCapTable(n int32) (CapDescriptor_List, error) {
	l, err := NewCapDescriptor_List(s.Struct.Segment(), n)
	if err != nil {
		return CapDescriptor_List{}, err
	}
	err = s.Struct.SetPtr(1, l.List.ToPtr())
	return l, err
}

// Payload_List is a list of Payload.
type Payload_List struct{ capnp.List }

// NewPayload creates a new list of Payload.
func NewPayload_List(s *capnp.Segment, sz int32) (Payload_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	return Payload_List{l}, err
}

func (s Payload_List) At(i int) Payload { return Payload{s.List.Struct(i)} }

func (s Payload_List) Set(i int, v Payload) error { return s.List.SetStruct(i, v.Struct) }

// Payload_Promise is a wrapper for a Payload promised by a client call.
type Payload_Promise struct{ *capnp.Pipeline }

func (p Payload_Promise) Struct() (Payload, error) {
	s, err := p.Pipeline.Struct()
	return Payload{s}, err
}

func (p Payload_Promise) Content() *capnp.Pipeline {
	return p.Pipeline.GetPipeline(0)
}

type CapDescriptor struct{ capnp.Struct }
type CapDescriptor_Which uint16

const (
	CapDescriptor_Which_none             CapDescriptor_Which = 0
	CapDescriptor_Which_senderHosted     CapDescriptor_Which = 1
	CapDescriptor_Which_senderPromise    CapDescriptor_Which = 2
	CapDescriptor_Which_receiverHosted   CapDescriptor_Which = 3
	CapDescriptor_Which_receiverAnswer   CapDescriptor_Which = 4
	CapDescriptor_Which_thirdPartyHosted CapDescriptor_Which = 5
)

func (w CapDescriptor_Which) String() string {
	const s = "nonesenderHostedsenderPromisereceiverHostedreceiverAnswerthirdPartyHosted"
	switch w {
	case CapDescriptor_Which_none:
		return s[0:4]
	case CapDescriptor_Which_senderHosted:
		return s[4:16]
	case CapDescriptor_Which_senderPromise:
		return s[16:29]
	case CapDescriptor_Which_receiverHosted:
		return s[29:43]
	case CapDescriptor_Which_receiverAnswer:
		return s[43:57]
	case CapDescriptor_Which_thirdPartyHosted:
		return s[57:73]

	}
	return "CapDescriptor_Which(" + strconv.FormatUint(uint64(w), 10) + ")"
}

// CapDescriptor_TypeID is the unique identifier for the type CapDescriptor.
const CapDescriptor_TypeID = 0x8523ddc40b86b8b0

func NewCapDescriptor(s *capnp.Segment) (CapDescriptor, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return CapDescriptor{st}, err
}

func NewRootCapDescriptor(s *capnp.Segment) (CapDescriptor, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return CapDescriptor{st}, err
}

func ReadRootCapDescriptor(msg *capnp.Message) (CapDescriptor, error) {
	root, err := msg.RootPtr()
	return CapDescriptor{root.Struct()}, err
}

func (s CapDescriptor) String() string {
	str, _ := text.Marshal(0x8523ddc40b86b8b0, s.Struct)
	return str
}

func (s CapDescriptor) Which() CapDescriptor_Which {
	return CapDescriptor_Which(s.Struct.Uint16(0))
}
func (s CapDescriptor) SetNone() {
	s.Struct.SetUint16(0, 0)

}

func (s CapDescriptor) SenderHosted() uint32 {
	return s.Struct.Uint32(4)
}

func (s CapDescriptor) SetSenderHosted(v uint32) {
	s.Struct.SetUint16(0, 1)
	s.Struct.SetUint32(4, v)
}

func (s CapDescriptor) SenderPromise() uint32 {
	return s.Struct.Uint32(4)
}

func (s CapDescriptor) SetSenderPromise(v uint32) {
	s.Struct.SetUint16(0, 2)
	s.Struct.SetUint32(4, v)
}

func (s CapDescriptor) ReceiverHosted() uint32 {
	return s.Struct.Uint32(4)
}

func (s CapDescriptor) SetReceiverHosted(v uint32) {
	s.Struct.SetUint16(0, 3)
	s.Struct.SetUint32(4, v)
}

func (s CapDescriptor) ReceiverAnswer() (PromisedAnswer, error) {
	p, err := s.Struct.Ptr(0)
	return PromisedAnswer{Struct: p.Struct()}, err
}

func (s CapDescriptor) HasReceiverAnswer() bool {
	if s.Struct.Uint16(0) != 4 {
		return false
	}
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s CapDescriptor) SetReceiverAnswer(v PromisedAnswer) error {
	s.Struct.SetUint16(0, 4)
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewReceiverAnswer sets the receiverAnswer field to a newly
// allocated PromisedAnswer struct, preferring placement in s's segment.
func (s CapDescriptor) NewReceiverAnswer() (PromisedAnswer, error) {
	s.Struct.SetUint16(0, 4)
	ss, err := NewPromisedAnswer(s.Struct.Segment())
	if err != nil {
		return PromisedAnswer{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s CapDescriptor) ThirdPartyHosted() (ThirdPartyCapDescriptor, error) {
	p, err := s.Struct.Ptr(0)
	return ThirdPartyCapDescriptor{Struct: p.Struct()}, err
}

func (s CapDescriptor) HasThirdPartyHosted() bool {
	if s.Struct.Uint16(0) != 5 {
		return false
	}
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s CapDescriptor) SetThirdPartyHosted(v ThirdPartyCapDescriptor) error {
	s.Struct.SetUint16(0, 5)
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewThirdPartyHosted sets the thirdPartyHosted field to a newly
// allocated ThirdPartyCapDescriptor struct, preferring placement in s's segment.
func (s CapDescriptor) NewThirdPartyHosted() (ThirdPartyCapDescriptor, error) {
	s.Struct.SetUint16(0, 5)
	ss, err := NewThirdPartyCapDescriptor(s.Struct.Segment())
	if err != nil {
		return ThirdPartyCapDescriptor{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

// CapDescriptor_List is a list of CapDescriptor.
type CapDescriptor_List struct{ capnp.List }

// NewCapDescriptor creates a new list of CapDescriptor.
func NewCapDescriptor_List(s *capnp.Segment, sz int32) (CapDescriptor_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1}, sz)
	return CapDescriptor_List{l}, err
}

func (s CapDescriptor_List) At(i int) CapDescriptor { return CapDescriptor{s.List.Struct(i)} }

func (s CapDescriptor_List) Set(i int, v CapDescriptor) error { return s.List.SetStruct(i, v.Struct) }

// CapDescriptor_Promise is a wrapper for a CapDescriptor promised by a client call.
type CapDescriptor_Promise struct{ *capnp.Pipeline }

func (p CapDescriptor_Promise) Struct() (CapDescriptor, error) {
	s, err := p.Pipeline.Struct()
	return CapDescriptor{s}, err
}

func (p CapDescriptor_Promise) ReceiverAnswer() PromisedAnswer_Promise {
	return PromisedAnswer_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p CapDescriptor_Promise) ThirdPartyHosted() ThirdPartyCapDescriptor_Promise {
	return ThirdPartyCapDescriptor_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

type PromisedAnswer struct{ capnp.Struct }

// PromisedAnswer_TypeID is the unique identifier for the type PromisedAnswer.
const PromisedAnswer_TypeID = 0xd800b1d6cd6f1ca0

func NewPromisedAnswer(s *capnp.Segment) (PromisedAnswer, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return PromisedAnswer{st}, err
}

func NewRootPromisedAnswer(s *capnp.Segment) (PromisedAnswer, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return PromisedAnswer{st}, err
}

func ReadRootPromisedAnswer(msg *capnp.Message) (PromisedAnswer, error) {
	root, err := msg.RootPtr()
	return PromisedAnswer{root.Struct()}, err
}

func (s PromisedAnswer) String() string {
	str, _ := text.Marshal(0xd800b1d6cd6f1ca0, s.Struct)
	return str
}

func (s PromisedAnswer) QuestionId() uint32 {
	return s.Struct.Uint32(0)
}

func (s PromisedAnswer) SetQuestionId(v uint32) {
	s.Struct.SetUint32(0, v)
}

func (s PromisedAnswer) Transform() (PromisedAnswer_Op_List, error) {
	p, err := s.Struct.Ptr(0)
	return PromisedAnswer_Op_List{List: p.List()}, err
}

func (s PromisedAnswer) HasTransform() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s PromisedAnswer) SetTransform(v PromisedAnswer_Op_List) error {
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewTransform sets the transform field to a newly
// allocated PromisedAnswer_Op_List, preferring placement in s's segment.
func (s PromisedAnswer) NewTransform(n int32) (PromisedAnswer_Op_List, error) {
	l, err := NewPromisedAnswer_Op_List(s.Struct.Segment(), n)
	if err != nil {
		return PromisedAnswer_Op_List{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

// PromisedAnswer_List is a list of PromisedAnswer.
type PromisedAnswer_List struct{ capnp.List }

// NewPromisedAnswer creates a new list of PromisedAnswer.
func NewPromisedAnswer_List(s *capnp.Segment, sz int32) (PromisedAnswer_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1}, sz)
	return PromisedAnswer_List{l}, err
}

func (s PromisedAnswer_List) At(i int) PromisedAnswer { return PromisedAnswer{s.List.Struct(i)} }

func (s PromisedAnswer_List) Set(i int, v PromisedAnswer) error { return s.List.SetStruct(i, v.Struct) }

// PromisedAnswer_Promise is a wrapper for a PromisedAnswer promised by a client call.
type PromisedAnswer_Promise struct{ *capnp.Pipeline }

func (p PromisedAnswer_Promise) Struct() (PromisedAnswer, error) {
	s, err := p.Pipeline.Struct()
	return PromisedAnswer{s}, err
}

type PromisedAnswer_Op struct{ capnp.Struct }
type PromisedAnswer_Op_Which uint16

const (
	PromisedAnswer_Op_Which_noop            PromisedAnswer_Op_Which = 0
	PromisedAnswer_Op_Which_getPointerField PromisedAnswer_Op_Which = 1
)

func (w PromisedAnswer_Op_Which) String() string {
	const s = "noopgetPointerField"
	switch w {
	case PromisedAnswer_Op_Which_noop:
		return s[0:4]
	case PromisedAnswer_Op_Which_getPointerField:
		return s[4:19]

	}
	return "PromisedAnswer_Op_Which(" + strconv.FormatUint(uint64(w), 10) + ")"
}

// PromisedAnswer_Op_TypeID is the unique identifier for the type PromisedAnswer_Op.
const PromisedAnswer_Op_TypeID = 0xf316944415569081

func NewPromisedAnswer_Op(s *capnp.Segment) (PromisedAnswer_Op, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return PromisedAnswer_Op{st}, err
}

func NewRootPromisedAnswer_Op(s *capnp.Segment) (PromisedAnswer_Op, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return PromisedAnswer_Op{st}, err
}

func ReadRootPromisedAnswer_Op(msg *capnp.Message) (PromisedAnswer_Op, error) {
	root, err := msg.RootPtr()
	return PromisedAnswer_Op{root.Struct()}, err
}

func (s PromisedAnswer_Op) String() string {
	str, _ := text.Marshal(0xf316944415569081, s.Struct)
	return str
}

func (s PromisedAnswer_Op) Which() PromisedAnswer_Op_Which {
	return PromisedAnswer_Op_Which(s.Struct.Uint16(0))
}
func (s PromisedAnswer_Op) SetNoop() {
	s.Struct.SetUint16(0, 0)

}

func (s PromisedAnswer_Op) GetPointerField() uint16 {
	return s.Struct.Uint16(2)
}

func (s PromisedAnswer_Op) SetGetPointerField(v uint16) {
	s.Struct.SetUint16(0, 1)
	s.Struct.SetUint16(2, v)
}

// PromisedAnswer_Op_List is a list of PromisedAnswer_Op.
type PromisedAnswer_Op_List struct{ capnp.List }

// NewPromisedAnswer_Op creates a new list of PromisedAnswer_Op.
func NewPromisedAnswer_Op_List(s *capnp.Segment, sz int32) (PromisedAnswer_Op_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0}, sz)
	return PromisedAnswer_Op_List{l}, err
}

func (s PromisedAnswer_Op_List) At(i int) PromisedAnswer_Op {
	return PromisedAnswer_Op{s.List.Struct(i)}
}

func (s PromisedAnswer_Op_List) Set(i int, v PromisedAnswer_Op) error {
	return s.List.SetStruct(i, v.Struct)
}

// PromisedAnswer_Op_Promise is a wrapper for a PromisedAnswer_Op promised by a client call.
type PromisedAnswer_Op_Promise struct{ *capnp.Pipeline }

func (p PromisedAnswer_Op_Promise) Struct() (PromisedAnswer_Op, error) {
	s, err := p.Pipeline.Struct()
	return PromisedAnswer_Op{s}, err
}

type ThirdPartyCapDescriptor struct{ capnp.Struct }

// ThirdPartyCapDescriptor_TypeID is the unique identifier for the type ThirdPartyCapDescriptor.
const ThirdPartyCapDescriptor_TypeID = 0xd37007fde1f0027d

func NewThirdPartyCapDescriptor(s *capnp.Segment) (ThirdPartyCapDescriptor, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return ThirdPartyCapDescriptor{st}, err
}

func NewRootThirdPartyCapDescriptor(s *capnp.Segment) (ThirdPartyCapDescriptor, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return ThirdPartyCapDescriptor{st}, err
}

func ReadRootThirdPartyCapDescriptor(msg *capnp.Message) (ThirdPartyCapDescriptor, error) {
	root, err := msg.RootPtr()
	return ThirdPartyCapDescriptor{root.Struct()}, err
}

func (s ThirdPartyCapDescriptor) String() string {
	str, _ := text.Marshal(0xd37007fde1f0027d, s.Struct)
	return str
}

func (s ThirdPartyCapDescriptor) Id() (capnp.Pointer, error) {
	return s.Struct.Pointer(0)
}

func (s ThirdPartyCapDescriptor) HasId() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s ThirdPartyCapDescriptor) IdPtr() (capnp.Ptr, error) {
	return s.Struct.Ptr(0)
}

func (s ThirdPartyCapDescriptor) SetId(v capnp.Pointer) error {
	return s.Struct.SetPointer(0, v)
}

func (s ThirdPartyCapDescriptor) SetIdPtr(v capnp.Ptr) error {
	return s.Struct.SetPtr(0, v)
}

func (s ThirdPartyCapDescriptor) VineId() uint32 {
	return s.Struct.Uint32(0)
}

func (s ThirdPartyCapDescriptor) SetVineId(v uint32) {
	s.Struct.SetUint32(0, v)
}

// ThirdPartyCapDescriptor_List is a list of ThirdPartyCapDescriptor.
type ThirdPartyCapDescriptor_List struct{ capnp.List }

// NewThirdPartyCapDescriptor creates a new list of ThirdPartyCapDescriptor.
func NewThirdPartyCapDescriptor_List(s *capnp.Segment, sz int32) (ThirdPartyCapDescriptor_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1}, sz)
	return ThirdPartyCapDescriptor_List{l}, err
}

func (s ThirdPartyCapDescriptor_List) At(i int) ThirdPartyCapDescriptor {
	return ThirdPartyCapDescriptor{s.List.Struct(i)}
}

func (s ThirdPartyCapDescriptor_List) Set(i int, v ThirdPartyCapDescriptor) error {
	return s.List.SetStruct(i, v.Struct)
}

// ThirdPartyCapDescriptor_Promise is a wrapper for a ThirdPartyCapDescriptor promised by a client call.
type ThirdPartyCapDescriptor_Promise struct{ *capnp.Pipeline }

func (p ThirdPartyCapDescriptor_Promise) Struct() (ThirdPartyCapDescriptor, error) {
	s, err := p.Pipeline.Struct()
	return ThirdPartyCapDescriptor{s}, err
}

func (p ThirdPartyCapDescriptor_Promise) Id() *capnp.Pipeline {
	return p.Pipeline.GetPipeline(0)
}

type Exception struct{ capnp.Struct }

// Exception_TypeID is the unique identifier for the type Exception.
const Exception_TypeID = 0xd625b7063acf691a

func NewException(s *capnp.Segment) (Exception, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return Exception{st}, err
}

func NewRootException(s *capnp.Segment) (Exception, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return Exception{st}, err
}

func ReadRootException(msg *capnp.Message) (Exception, error) {
	root, err := msg.RootPtr()
	return Exception{root.Struct()}, err
}

func (s Exception) String() string {
	str, _ := text.Marshal(0xd625b7063acf691a, s.Struct)
	return str
}

func (s Exception) Reason() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s Exception) HasReason() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Exception) ReasonBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s Exception) SetReason(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s Exception) Type() Exception_Type {
	return Exception_Type(s.Struct.Uint16(4))
}

func (s Exception) SetType(v Exception_Type) {
	s.Struct.SetUint16(4, uint16(v))
}

func (s Exception) ObsoleteIsCallersFault() bool {
	return s.Struct.Bit(0)
}

func (s Exception) SetObsoleteIsCallersFault(v bool) {
	s.Struct.SetBit(0, v)
}

func (s Exception) ObsoleteDurability() uint16 {
	return s.Struct.Uint16(2)
}

func (s Exception) SetObsoleteDurability(v uint16) {
	s.Struct.SetUint16(2, v)
}

// Exception_List is a list of Exception.
type Exception_List struct{ capnp.List }

// NewException creates a new list of Exception.
func NewException_List(s *capnp.Segment, sz int32) (Exception_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1}, sz)
	return Exception_List{l}, err
}

func (s Exception_List) At(i int) Exception { return Exception{s.List.Struct(i)} }

func (s Exception_List) Set(i int, v Exception) error { return s.List.SetStruct(i, v.Struct) }

// Exception_Promise is a wrapper for a Exception promised by a client call.
type Exception_Promise struct{ *capnp.Pipeline }

func (p Exception_Promise) Struct() (Exception, error) {
	s, err := p.Pipeline.Struct()
	return Exception{s}, err
}

type Exception_Type uint16

// Values of Exception_Type.
const (
	Exception_Type_failed        Exception_Type = 0
	Exception_Type_overloaded    Exception_Type = 1
	Exception_Type_disconnected  Exception_Type = 2
	Exception_Type_unimplemented Exception_Type = 3
)

// String returns the enum's constant name.
func (c Exception_Type) String() string {
	switch c {
	case Exception_Type_failed:
		return "failed"
	case Exception_Type_overloaded:
		return "overloaded"
	case Exception_Type_disconnected:
		return "disconnected"
	case Exception_Type_unimplemented:
		return "unimplemented"

	default:
		return ""
	}
}

// Exception_TypeFromString returns the enum value with a name,
// or the zero value if there's no such value.
func Exception_TypeFromString(c string) Exception_Type {
	switch c {
	case "failed":
		return Exception_Type_failed
	case "overloaded":
		return Exception_Type_overloaded
	case "disconnected":
		return Exception_Type_disconnected
	case "unimplemented":
		return Exception_Type_unimplemented

	default:
		return 0
	}
}

type Exception_Type_List struct{ capnp.List }

func NewException_Type_List(s *capnp.Segment, sz int32) (Exception_Type_List, error) {
	l, err := capnp.NewUInt16List(s, sz)
	return Exception_Type_List{l.List}, err
}

func (l Exception_Type_List) At(i int) Exception_Type {
	ul := capnp.UInt16List{List: l.List}
	return Exception_Type(ul.At(i))
}

func (l Exception_Type_List) Set(i int, v Exception_Type) {
	ul := capnp.UInt16List{List: l.List}
	ul.Set(i, uint16(v))
}

const schema_b312981b2552a250 = "x\xda\x9cX\x7f\x8c\x15\xd5\x15>\xe7\xde\xb7\xef-\xec" +
	"\x8f\xf7f\xefC\x0b\x95@mM\x0a)D\xacm\xed" +
	"\xb6\xe6!\xec\x12\xd6,a\xef\xbe\xa5*5ig\xdf" +
	"\xbb\xec\xce2;3\xce\xcc\xc2.\x91\x00\xad6J%" +
	"E\xa2\x16\x8c\xb4jlb-FD\x88\xb4\x95T\x08" +
	"\xa6j\xa4j\x14\xa3\x8d\xa6j\xd2\xa8MM\xb4\xfe\xa8" +
	"\xca\x8fi\xce\xcc\xbc\x99\xc7\xee\xdb\x10\xfb\xd7\xbe\xccw" +
	"f\xee9\xdf=\xe7\xfb\xee\xddK\xaf\xcb.eK\x9a" +
	"V\xcf\x00\x90fS6x\xb9w\xef\xf8\xdf\xca#?" +
	"\x079\x13y\xd0w\x7f\xff%_\xdd\xdd\xf1\x184\xf1" +
	"\x1c\x80\xd8\x91\xd9$n\xcf\xe4\x00\xbe\xbd#\xf3+\x04" +
	"\x0c\xf6\x1f\xfeE\xcb\xf17\xbe~3Ec\x1a\xdd\x8d" +
	"\xb9,\x80\xd0\xb2\xc7\xc4\xec,\x85\xcf\xca^C\xe1\x97" +
	"\xed\xdf\xb1u\xdeo\x1f\xbf}jx;\x80\x98\xc8\xed" +
	"\x12\xdbr\x14\xbe9w!\x07\x0c\x8e\x9e\x12\xd7\x0e\x14" +
	"\x9f\xb8sj8C&\x8e\xcc<&\x9e\x9aIi\x1d" +
	"\x9d\xb9\x110\xf8\x81\x7f\xd7\x95\x17\xeb\xedw\x836\xb3" +
	".\xb8\x89Q\xc4%-\xbb\xc4\xa2\x16\xfa\xb5\xa0\x85b" +
	"\xd7\xee;zj}f\xe4\x9eI_\x8e\x82\xf7\xb4\xec" +
	"\x12\xf7\x85\xc1{[\x1e\x01\x0c:\xafy\xec\xca\x1d\x07" +
	"f\xff\x86\x82\xd9\xb9E\"\x17W\xb6n\x17\xdd\xad\x94" +
	"\xf5U\xad\x7f\xa5\"\x7f\xed\xbf\xb0\xb9\xcd\x9c\xf3\xf0\xa4" +
	"o\x13mbI\xfb.\xf1\xfdv\xfa\xf5\x9dv\xca\xe3" +
	"\xda#\xbd\xa5\xb7\xef\xba\xed\x00hE\x16\xcc1\x9e\xef" +
	"\xcc>~\xc9+\x00(\xeel\x7fV\xdc\x17\x06\xeem" +
	"\x1f\x02\x0c\xac\xe6[\xbfXs\xd7\xb1?7\xa6\xe2D" +
	"\xfb.q2\x8c~\xb1\x9d2\xde\xcc>x\xebL\xce" +
	"yiryHi\xde\x90\xef@\xb1-O\xd1\x9b\xf3" +
	"\x94D\xa5\xfd\xb3c\x07\x16o~\xa9Q\xc2\xaf\xe5\xb7" +
	"\x8b\xb7\xc2\xd87\xc2\xd8\x0b\x96\xae\xd99x\xe8\x99\x97" +
	"\x1b}Yt\x17\xb6\x8bU\x05\xfa\xd5S\xa04V\xbd" +
	"\xf1c\xf5\x8f\x83\x83'A\xceB\x0c\xb4\xef\x1d\xc9\xff" +
	"\xf2\xbb\xd5\xcfa\x0d\xe60\x83L\xbcW\xf8\x17\xa0x" +
	"\xbf\xf0\x0e`Z{\xa3\xef\x9e\xd0\xee\x17'\xb5\x0b)" +
	"\x09\xed\x1d\xc0\xbf\xdc{\x91}\xe2\x95G_m\x14z" +
	"\xb4\xe3Yq\xa2\x83BOvP\xbe{~\xf2\x879" +
	"\x9f\xee\x7f\xf7\xef \xf3\xc8\xd3\xe6^\xc3s\xc8\x91\x8b" +
	"\x1eA)\xac\x12\x94\xedq\xeb\xc2%[\x9f\xef}\xaf" +
	"a\x0a\xef\x8b\xfb\xc5'\x82~}(\xe8\xbb\xdbv\xfe" +
	"hV\xd7\x1d\x17|\x04r6&\x09u\xe5\x18\x80\xb8" +
	"\xae\xf8\xb6PE\x0a\xd5\x8b\x14\x9a\xd4\xdd0\xdf\xe2C" +
	"\xe2\x990\xf8\xa90\xf8\x11|sgf\xf7[\xa7\x1a" +
	"6f\xd3\xacMb\xc6\xac\xe8\xd7#\xb0<p\x9d\xca" +
	"\xe2\x8a\xeeXPr:\x97\xeb\xa6\xd9\x87(/\xe2\x19" +
	"\x80\x0c\x02h\x87\xd6\x02\xc8\x83\x1c\xe5\x93\x0c\x11\x8bH" +
	"\xcf\x8et\x02\xc8\xc3\x1c\xe5q\x86\x1a\xc3\"2\x00\xed" +
	"\xe8 \x80|\x92\xa3|\x8e\xa1\xc6Y\x119\x80\xf6\xcc" +
	"\xd5\x00\xf2i\x8e\xf2e\x86Z\x13\x161\x03\xa0\xbdH" +
	"\xaf?\xc7Q\xbe\xca\x10\xb3XG\xafv\xd2\x05\xa6e" +
	"\xb6\x16\xb1\x19Q;z\x0c@\x1e\xe7(_`\x18\xdc" +
	"0\xa6<\xdf\xb0-\xe0=Ul\x06\x86\xcd\x80%_" +
	"w\x87\x94\x8f\x85t\xc6\x01\xb1\x00\x18\x18\x96\xaf\xdcu" +
	"z\x05r\xaa\xa7\x8a3\x80\xe1\x0c\xc0`T\xf9\xc3v" +
	"\xb5\xa7\x0a\x00\x98\x03\x869\xc0\x92\xa3\xbb\xfa\xa8\x87\x85" +
	"t\xf0\xe3Ox\xca\xaa\xf6+o\x0c\xe6\x99\xbe7`" +
	"\x07\xbai\xda\x1b\x07\x86\x0d\xe6V\xfbt\xd7\x9f\x18\xd0" +
	"\x0d\x93\xe8\x02D`H#[#\x92\x11\x8fN\x97\xf2" +
	"*\xae\xe1\xf8\xb6\x0b\xc4\xe8Wx\xa65\x08BJ\xf7" +
	",\x04\x90wp\x94\xf72\x9c\x8bg\x83\x98\xd5\xbd#" +
	"\x00\xf2\x1e\x8e\xf2A\x86s\xd9\x99 \xe6\xf5w.\x80" +
	"|\x80\xa3\xdc\xcfp.?M\x8f\x89\xd9\x877\x01\xc8" +
	"}\x1c\xe5a\x86m\x99SAD\xed\xa1M\xe9n\xb5" +
	"5}\x11\x14\xb1\x89\xf6k{\xba5y\xcb\xb6\x14d" +
	"\xc3\xf2\x94\xbb\xd2\x86\xbc\xe7\xab\x84\xd1\xf8q\x9f\x0b\xf3" +
	"\xecQ\xc3S\xc9sWU\x94\xb1A\xb9PZi\x9f" +
	"\xf3B\x0a\\ey\x1b\x95\x8b\x85Z\x1f\xc7<\xfa\xc3" +
	"F\xc8\x18\xfa\x13\xd1\xab\x00XH\xb5%\x8e\xaaq\x87" +
	"N\xe7*\xe5y\xfa\x10*b\xed\x8a\x8451\x81." +
	"@y\x1c9\x96oB\x86mx6\x08y\x13\xdb\xf0" +
	"2\x80\xf2\x8d\x04\xdcB\x00?\x13\x84\xcc\x89\x9bq!" +
	"@y+\x01\xb7\x11\x909\x1d\x84\xdc\x89[\xb1\x13\xa0" +
	"|\x13\x01;\x09h\x8a\xe9\x13;B\xe0\x16\x02\xee " +
	" \x1b3(n\xc7e\x00\xe5\xdb\x08\xd8M@\xee\xf3" +
	"\xa0\x88dNw\x86\xc0N\x02\xee!`\xc6gA1" +
	"\x1c\xc9=8\x02P\xdeM\xc0\x03\x04\xb0\xff\x06El" +
	"\x06\x10\xf7a?@\xf9^\x02\xf6\x110\xf3\xd3\xa0\x88" +
	"3\x00\xc4\xefq\x13@\xf9A\x02\x0e\x12\xd0\xf2IP" +
	"\xc4\x99\x00\xe2\xd1p\x8d}\x04\x1c&\xa0\xf5\xe3\xa0\x88" +
	"-\x00\xe2P\x98\xee~\x02\x9e \xa0\xed\xa3\xa0\x88\xad" +
	"\x00\xe2\x8fa\xe5\x07\x09x\x92\x80\xe6\xff\x04El\x03" +
	"\x10Gp-@\xf9\x09\x02\x9eF\x86\xc1\x98e\x8c:" +
	"\xa6\x1a\x85y\xca\xa2]-\xa4\xe6\x1am\xcc<}\xd0" +
	"vi\xc2\xeal\x85\x9e\xe7+\xbaib!\xd5\xc2\xe8" +
	"q\xc9U\xfe\x98ka!\xb5\xbb\x18XgX\x867" +
	"\x8c\x85\xd4'\"`\x8b\xab<\xdb\xdc\xa0\xb0\x90\xbaS" +
	"\x82\x98J\xf7\x08I\xcc0\xee\x16{\xd0\xb3M\xe5+" +
	"\xc8\x97\xf5\x0d\x0a;\x80a\x07`0h\xdb\xbe\xe7\xbb" +
	":\xa0\x83\x85T\x89'\xbfT\xeaR\xf4\xb7\xf6\xda\x16" +
	"\xc7\xb57\x18UZ'1\xf48i\xbdRQ\x0eU" +
	"\x9f\x18V\\\xfd\x88mP\x91\x89\xce\xc6KT\x0dO" +
	"\x8d\x0e\xea.\xf0!\x1b\x0b\xa9fOjrVkr" +
	"5\x10\x0aX(\x10\xcd\xa9@, )\xfd&Gy" +
	"y]\x9fkKh\xb6/\xe5(\x7f\xc800F\x1d" +
	"\xdb\xa5a\xca-\xd7\x9dd\x18\x1d7\x9c\xda\xea\xb4\xc3" +
	"X7f}\xfa\x84i\xebX\x8d\xd7\x8e\xe5~\xc12" +
	"\x00\xf9\x0d\x8e\xf2R\x86ZM\xef\x17\x91\x8a\x7f\x8b\xa3" +
	"\\\xc9pK\xc5\xb6|e\xf9\x09\xe9\x15\xdd\x19\xd0\x07" +
	"ME\xa2\xda\x0e\xd8\xc7\x11\x0b\xe9\x89\x0e\x90\x1e\x9e\xb3" +
	"n\xc8v4\xde\xad\xc9\xba\xddd3]\x1ce_j" +
	"3\xab\xc8'Vr\x94\x03u6#\xfb\x01d\x1fG" +
	"y\xfd\x976\x05WU\x0c\xc7P\x16`\x9a}]b" +
	"\xfda\xebB\xb8\x19\xf3\x93\xc4^\xa4\xda_\xe0(_" +
	"'B\xe6\x17\x11\x11\xb5\xd7HP_\xe7(\xdf\xa5\xc9" +
	"\x0e\"\xbd\xd1\xfeI\xdc\xbd\xc9Q\xfe\x9bT\xe8l$" +
	"6\xda{\x94\xf0\xbb\x1c\xe5\xc7$Agb\xa1\xfe\x90" +
	">\xfb\x01Gy\x9a\xf4\xe7t,\xd4\x9f?\x04 O" +
	"s,7#\xc3\xb9\xd9S\x01\x8bT\xa6\x09\x0f\x00\x94" +
	"\x9bil\x8b\xa1\xfc|\x11\xab\x8c\x86\x0f\x01\x94\x8b\x04" +
	"\xcc\xa7y\xd6\xc3m\x8f\x1c.U\xe8p\x8c\xfa\x90\x9c" +
	"n\xb9\xeex\x10ZV\x13y\x16M\xdf\x98\xe97\xf2" +
	"?5N\xbdo\xd8\x80\xd6\xd4\xf1\x0f*\xbaUQ&" +
	"\x89y.\x88\xbfQFe\xf9\xdd\xa6\xa76\xe6\x87\x95" +
	"K\x1e\xe3\xeb\xeb\xd5\x0a\xd7\x1e\xc5\xd5\xfe\xb0r\xe5\x98" +
	"\x9a\x17\xeeV\x92Y4^+\\\xb4G\x07B\x97\xc8" +
	"\x93\xb16\xde\x1b\xaa!j\x9a\xbaf\x9d\xd3\xa8Y7" +
	"\xc5\xcdz\x05Cn\xd4\x1b\xd5:\xe5*\xab\x02%\xb5" +
	"\xdc\x1e\xb3\xfc\x14H\xa7\xb2;\xae\xd9Z<0\xe1\xa8" +
	"\xa8\x15\x0a\xe1\xde.\xe8\xa4\xca\xb5\xaf\xad\x05@\xa6\xcd" +
	"\x1d\x01@\xae\xcdv\x01J\xebt\xc3T\xd5\xc0\xde\xa0" +
	"\\\xd3\xd6\xab\xc0U\x95t\xa0b[\x96\x82|\xc5W" +
	"\xd5\xc9*{na\xa4~S\xa6\xa1?\x9d\x866\x0c" +
	"b\x01Xuq:\x0fm\xecl\xd0` b\x01\xe8" +
	"\x01L\x0a\xcfUtg\xd2D\x9e\x7f{k\x19r\xa7" +
	"s \xf6o\x7f\xa2\xfeP\x83\xee\xf4[\x91\xecDg" +
	"*c\xb4\x13\xf1\xbe\x966\x18\x96\xea\xa9N\xe1\x1f\x9d" +
	"\xce\x15\xa1I\x00L\xfa\xf6\xda\xf4;\xc9\x08.\xd9\x05" +
	" /\xe7(\x97N\xa3\x03\xb5\xbe\xef\xc7\xb0=I&" +
	"\xbd\xa4\xef\xeb\x17\xbd*\xec\xc2h\xd1\xf3\x08\x12Q\xdd" +
	"\xcbQ^K\x824?\xe2\x7f\xcd\xb2\xf3\x08R\x10\xfa" +
	"\x8b\x17Q]\xf3\x9c\xd0&\x86\xecFg\xc7\xae\xd8D" +
	"\x86\xec\xc5\x15\xdb\xca\xfbj\xdc\x97\x85\xd0\x1c\xa2,t" +
	"j\xf0\x9fr\x94f\xcd\x1d(\x0d\x83$\xc9\xe4(\xc7" +
	"\xa99\xce\xc4\xe23F[\xe0p\x947\x92$\x9d\x8e" +
	"\xc5g\x82R\xf69\xca\xad\xacv\xe2\xeb\xb5\xa1d;" +
	"\x83ze\xfd\x94\x93\x1d\xf6\xda\x11\x92jJl\x8c\x90" +
	"M\xbc\xb3\xc1fF\xc3\x943lKf\xb0\xfe\x92\x8a" +
	"\x0b\xf34^TTL\xb6Ni^\xcfQ\x0e3D" +
	"\x16\x95\xa9\xfe\x04 \x879J\x9f\xee\x13\xb1\xfa\xdfp" +
	"w\x9a\xb9\x86\xf1%c3\x9d\xa7\xc79\xca\x9b\x18\x1d" +
	"@t\xcf\xb6\xb0\x15\x18\xb6\xd6\x99>\xf6xtXW" +
	"n\xc9[\xa1\x8f\x99~B|\x12\xd05\xe6\xea\x83\x86" +
	"ip\x7f\xa2v9\xc8\xfb\x13\x8e\xc2|\x9a: \xe6" +
	"\xcf\xdd\xac\xbe\xd8q#\xbf\x05\x08KM\xaeu\x1a\xce" +
	"\xe1\xab\x9diZ\xb9\xd6UK\xfac_\xef\x9d\xae\x81" +
	"|W\xb7\xbcu\xb6\x0b8\x9aZl\xb2\xc8$\x8be" +
	"\xd1-nq\xed\xfeb\xe6\xe9\xfaB\x9d\x1dv\x10\xd9" +
	"L7\xd1\xbd4Z1\xea\xa0,\x80\xd6su*/" +
	"t\xff`\xa1\xc5hrm\xda\xdf\xa5J\xc8!d\x83" +
	"\x09{\xcc\xf5\x94\xb9\x8e\xf4\xbfv\xc2\x07\xdeX\xbc\x97" +
	"\x85\xc7\xb2\x9c\xab;\xd3\xcfuB\xc6\xdd\xe7\x1b\xeb\xaa" +
	"r\\U\xd1}T\xd5\xd5\x83#\xaa\xe2\x138y\xd5" +
	");\x93[\xbc\xda\x99|\xcaZ\x98JV\xdd5l" +
	"\xd1\xcfR\xff\xc8[\xb6\xed@6\x18R~\x9fmX" +
	">*w\x85\xa1\xccjr}\xac/3\x9a\xdb<\x0d" +
	"\xee\xa4:;\xeb\xb5\x11\xeb\xfe\xa3\xa1-Z\x06l\xda" +
	"\x03Kt\xd4\x1a\xf7\xcf\xb9\xa1_m\x1b\xd6\xff{t" +
	"Z\x96\xca\xd7\x97;:mY\xaf&\xc8\x02j<\xff" +
	"/\x00\x00\xff\xff<\xff\x0e\x96"

func init() {
	schemas.Register(schema_b312981b2552a250,
		0x836a53ce789d4cd4,
		0x8523ddc40b86b8b0,
		0x91b79f1f808db032,
		0x95bc14545813fbc1,
		0x9a0e61223d96743b,
		0x9c6a046bfbc1ac5a,
		0x9e19b28d3db3573a,
		0xad1a6c0d7dd07497,
		0xb28c96e23f4cbd58,
		0xbbc29655fa89086e,
		0xd37007fde1f0027d,
		0xd37d2eb2c2f80e63,
		0xd4c9b56290554016,
		0xd562b4df655bdd4d,
		0xd625b7063acf691a,
		0xd800b1d6cd6f1ca0,
		0xdae8b0f61aab5f99,
		0xe94ccf8031176ec4,
		0xf316944415569081,
		0xf964368b0fbd3711,
		0xfbe1980490e001af)
}
