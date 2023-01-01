package predict

import "back/protocol"

type Middleware func(protocol.PredictService) protocol.PredictService
