package vectorize

import "errors"

const imageKind = "image"
const textKind = "text"
const vectorEngineOpenAI = "openai"

func DoInit(kind string) error {
	switch kind {
	case textKind:
		return nil
	default:
		return errors.New("unknown kind")
	}
}

func DoSeed(kind string, inputDir string) error {
	switch kind {
	case textKind:
		return nil
	default:
		return errors.New("unknown kind")
	}
}

func DoFind(kind string, inputFile string) error {
	switch kind {
	case textKind:
		return nil
	default:
		return errors.New("unknown kind")
	}
}

func DoVectorize(kind string, vectorEngine string) error {
	switch kind {
	case textKind:
		return vectorizeText(vectorEngine)
	default:
		return errors.New("unknown kind")
	}
}

func vectorizeText(vectorEngine string) error {
	switch vectorEngine {
	case vectorEngineOpenAI:
		return nil
	default:
		return errors.New("unknown vectorizetion engine")
	}
}
