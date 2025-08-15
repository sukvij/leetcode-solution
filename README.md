"# leetcode-solution" 

// Levenshtein distance



import (
	"fmt"
)

// Define states
type State int16

const (
	StateNew State = iota
	StateReady
	StateRunning
	StateBlocked
	StateExit
)

// Define events
type Event int16

const (
	EventAdmit Event = iota
	EventDispatch
	EventTerminate
	EventTimeout
	EventWaitForEvent
	EventOccur
)

// Function type
type Function func() error

// State machine structure
type StateMachine struct {
	CurrentState State
	Transitions  map[StateEvent]*CmdState
}

// State event structure
type StateEvent struct {
	Src   State
	Event Event
}

// Command state structure
type CmdState struct {
	Cmd  Function
	Dest State
}

// Create a new state machine
func NewStateMachine(initialState State) *StateMachine {
	return &StateMachine{
		CurrentState: initialState,
		Transitions:  make(map[StateEvent]*CmdState),
	}
}

// Add a transition
func (sm *StateMachine) AddTransition(src State, ev Event, cmd Function, dest State) {
	se := StateEvent{src, ev}
	sm.Transitions[se] = &CmdState{cmd, dest}
}

// Run the state machine
func (sm *StateMachine) Run() error {
	const maxSteps = 100 // maximum steps allowed to avoid infinite loop
	stepCount := 0

	for sm.CurrentState != StateExit {
		if stepCount >= maxSteps {
			return fmt.Errorf("infinite loop detected after %d steps", stepCount)
		}

		var transition *CmdState
		for event := Event(0); event <= EventOccur; event++ { // Iterate through all events
			se := StateEvent{sm.CurrentState, event}
			if trans, ok := sm.Transitions[se]; ok {
				transition = trans
				break
			}
		}

		if transition == nil {
			return fmt.Errorf("no transition found for state %v", sm.CurrentState)
		}

		// Execute the command associated with the transition
		if err := transition.Cmd(); err != nil {
			return err
		}

		// Update the current state to the destination of the transition
		sm.CurrentState = transition.Dest
		stepCount++ // Increment step count
	}

	return nil
}

// Functions for state transitions
func AdmitProcess() error {
	fmt.Println("Admitting new process...")
	return nil
}

func DispatchProcess() error {
	fmt.Println("Dispatching process to CPU...")
	return nil
}

func TerminateProcess() error {
	fmt.Println("Terminating process...")
	return nil
}

func TimeoutProcess() error {
	fmt.Println("Handling timeout and moving process to ready...")
	return nil
}

func WaitForEvent() error {
	fmt.Println("Blocking process while waiting for event...")
	return nil
}

func EventOccurred() error {
	fmt.Println("Event occurred, moving process to ready...")
	return nil
}

func main() {
	// Create a new state machine with initial state
	sm := NewStateMachine(StateNew)

	// Define transitions
	sm.AddTransition(StateNew, EventAdmit, AdmitProcess, StateReady)
	sm.AddTransition(StateReady, EventDispatch, DispatchProcess, StateRunning)
	sm.AddTransition(StateRunning, EventTerminate, TerminateProcess, StateExit)
	sm.AddTransition(StateRunning, EventTimeout, TimeoutProcess, StateReady)
	sm.AddTransition(StateRunning, EventWaitForEvent, WaitForEvent, StateBlocked)
	sm.AddTransition(StateBlocked, EventOccur, EventOccurred, StateReady)

	if err := sm.Run(); err != nil {
		fmt.Println("Error:", err)
	}
}
