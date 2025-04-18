package answer

var StatusCode = map[ErrorCode]string{
	200: "OK",
	201: "Created",
	202: "Accepted",
	204: "No Content",
	301: "Moved Permanently",
	302: "Found",
	304: "Not Modified",
	400: "Bad Request",
	401: "Unauthorized",
	403: "Forbidden",
	404: "Not Found",
	500: "Internal Server Error",
	502: "Bad Gateway",
	503: "Service Unavailable",
	504: "Gateway Timeout",
}

type ErrorCode int

const (
	OK                  = 200
	Created             = 201
	Accepted            = 202
	NoContent           = 204
	MovedPermanently    = 301
	Found               = 302
	NotModified         = 304
	BadRequest          = 400
	Unauthorized        = 401
	Forbidden           = 403
	NotFound            = 404
	InternalServerError = 500
	BadGateway          = 502
	ServiceUnavailable  = 503
	GatewayTimeout      = 504
)
