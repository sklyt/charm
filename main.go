package main

/*
#include <stdlib.h>
#include <stdint.h>
*/
import "C"
import (
	"unsafe"
)

//export InitEngine
func InitEngine() C.int {
	engine := GetEngineInstance()
	return C.int(engine.Initialize())
}

//export CreateEntity
func CreateEntity() C.uint32_t {
	engine := GetEngineInstance()
	return C.uint32_t(engine.CreateEntity())
}

//export DestroyEntity
func DestroyEntity(entityID C.uint32_t) C.int {
	engine := GetEngineInstance()
	return C.int(engine.DestroyEntity(EntityID(entityID)))
}

//export GetLastEngineError
func GetLastEngineError() *C.char {
	engine := GetEngineInstance()
	if engine.lastError == "" {
		return nil
	}
	return C.CString(engine.lastError)
}

//export FreeString
func FreeString(str *C.char) {
	C.free(unsafe.Pointer(str))
}

//export Shutdown
func Shutdown() {
	engine := GetEngineInstance()
	engine.Shutdown()
}

func main() {}
