package containers

// States for messages
const (
	Vote       = iota
	Commit
	Empty
)

type NodeId int
type Stamp int
type StepId int

func (step *StepId) Inc() {
	*step += 1
}

func (step *StepId) Equal(rghStep *StepId) bool {
	return *step == *rghStep
}

func (step *StepId) NotEqual(rghStep *StepId) bool {
	return ! step.Equal(rghStep)
}

type Message struct {
	MessageType int
	Stamp       Stamp
	StepId      StepId
	VotedSet    Set // FIXME: remove this
	Carry       Carry
	NodeSet     Set
	IdFrom      NodeId
}

func NewMessageVote(carry Carry, votedSet Set, nodesSet Set) Message {
	return Message{
        MessageType: Vote,
        Carry:       carry,
        VotedSet:    votedSet,
        NodeSet:     nodesSet,
    }
}

func NewMessageCommit(carry Carry) *Message {
	return &Message{
        MessageType: Commit,
        Carry:       carry,
    }
}

func NewEmptyMessage() Message {
	return Message{MessageType: Empty}
}

func (msg *Message) GetType() int {
	return msg.MessageType
}

// For testing purposes
func (msg *Message) notEqual(otherMsg *Message) bool {
	return msg.VotedSet.NotEqual(otherMsg.VotedSet)
}

func (msg *Message) Equal(other Message) bool {
	// FIXME: додумать этот кусок
	if msg.MessageType != other.MessageType || msg.IdFrom != other.IdFrom {
		return false
	}

	if msg.MessageType == Empty {
		return true
	}

	if msg.MessageType == Vote {
		if msg.NodeSet.Equal(other.NodeSet) && msg.VotedSet.Equal(other.VotedSet) &&
		   msg.Stamp == other.Stamp && msg.StepId == other.StepId {
			return true
		}
	}

	if msg.MessageType == Commit {
		if msg.NodeSet.Equal(other.NodeSet) && msg.Carry.Equal(other.Carry) &&
			msg.Stamp == other.Stamp && msg.StepId == other.StepId {
			return true
		}
	}

	return false
}

func (msg *Message) NotEqual(other Message) bool {
	return !msg.Equal(other)
}
