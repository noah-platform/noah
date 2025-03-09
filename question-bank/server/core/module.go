package core

type Module string

const (
	ModuleListening Module = "LISTENING"
	ModuleReading   Module = "READING"
	ModuleSpeaking  Module = "SPEAKING"
	ModuleWriting   Module = "WRITING"
)

func (m Module) IsValid() bool {
	switch m {
	case ModuleListening, ModuleReading, ModuleSpeaking, ModuleWriting:
		return true
	default:
		return false
	}
}
