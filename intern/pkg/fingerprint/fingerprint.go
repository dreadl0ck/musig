package fingerprint

import "github.com/dreadl0ck/musig/intern/pkg/model"

type Fingerprinter interface {
	Fingerprint(uint32, []model.ConstellationPoint) map[model.EncodedKey]model.TableValue
}
