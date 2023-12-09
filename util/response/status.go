package response
// Status ...
type Status struct {
	Code        int    `json:"code"`
	Status      string `json:"status"`
	Description string `json:"description"`
}
var (
	OK = Status{
		Code:        200,
		Status:      "OK",
		Description: "The request has succeeded",
	}
	Created = Status{
		Code:        201,
		Status:      "CREATED",
		Description: "The request has been fulfilled and has resulted in one or more new resources being created",
	}
	NoContent = Status{
		Code:        204,
		Status:      "NO_CONTENT",
		Description: "There is no content to send for this request, but the headers may be useful",
	}
	BadEnvironment = Status{
		Code:        400,
		Status:      "BAD_ENVIRONMENT",
		Description: "The service has an invalid environment value",
	}
	BadRequest = Status{
		Code:        400,
		Status:      "BAD_REQUEST",
		Description: "The server could not understand the request due to invalid syntax",
	}
	InvalidArgument = Status{
		Code:        400,
		Status:      "INVALID_ARGUMENT",
		Description: "	Incorrect request parameters specified. Details are provided in the details field.",
	}
	FailedPrecondition = Status{
		Code:        400,
		Status:      "FAILED_PRECONDITION",
		Description: "	The operation was canceled because the conditions required for the operation were not met.",
	}
	OutOfRange = Status{
		Code:        400,
		Status:      "OUT_OF_RANGE",
		Description: "Out of range. For example, searching or reading outside of the file.",
	}
	Unauthorized = Status{
		Code:        401,
		Status:      "UNAUTHORIZED",
		Description: "The operation requires authentication.",
	}
	Forbidden = Status{
		Code:        403,
		Status:      "FORBIDDEN",
		Description: "...",
	}
	PermissionDenied = Status{
		Code:        403,
		Status:      "PERMISSION_DENIED",
		Description: "The user has no permissions required to perform the operation.",
	}
	NotFound = Status{
		Code:        404,
		Status:      "NOT_FOUND",
		Description: "The requested resource not found.",
	}
	AlreadyExists = Status{
		Code:        409,
		Status:      "ALREADY_EXISTS",
		Description: "arguments already exists",
	}
	Aborted = Status{
		Code:        409,
		Status:      "ABORTED",
		Description: "The operation failed due to a concurrent computing conflict, such as an invalid sequence of commands or an aborted transaction.",
	}
	FailedDependency = Status{
		Code:        424,
		Status:      "FAILED_DEPENDENCY",
		Description: "The operation failed due to a concurrent computing conflict, such as an invalid sequence of commands or an aborted transaction.",
	}
	TooManyRequests = Status{
		Code:        429,
		Status:      "TOO_MANY_REQUESTS",
		Description: "The user has sent too many requests in a given amount of time",
	}
	ResourceExhausted = Status{
		Code:        429,
		Status:      "RESOURCE_EXHAUSTED",
		Description: "	The request limit exceeded.",
	}
	Canceled = Status{
		Code:        499,
		Status:      "CANCELED",
		Description: "The operation was aborted on the client side.",
	}
	InternalServerError = Status{
		Code:        500,
		Status:      "INTERNAL_SERVER_ERROR",
		Description: "Internal server error. This error means that the operation cannot be performed due to a server-side technical problem. For example, due to insufficient computing resources.",
	}
	Internal = Status{
		Code:        500,
		Status:      "INTERNAL",
		Description: "Internal server error. This error means that the operation cannot be performed due to a server-side technical problem. For example, due to insufficient computing resources.",
	}
	GRPCError = Status{
		Code:        500,
		Status:      "GRPC_ERROR",
		Description: "The gRPC request failed",
	}
	Unknown = Status{
		Code:        500,
		Status:      "UNKNOWN",
		Description: "Unknown error.",
	}
	DataLoss = Status{
		Code:        500,
		Status:      "DATA_LOSS",
		Description: "Permanent data loss or damage.",
	}
	Unimplemented = Status{
		Code:        501,
		Status:      "UNIMPLEMENTED",
		Description: "The operation is not supported by the service.",
	}
	Unavailable = Status{
		Code:        503,
		Status:      "UNAVAILABLE",
		Description: "The service is currently unavailable. Try again in a few seconds.",
	}
	DeadlineExceeded = Status{
		Code:        504,
		Status:      "DEADLINE_EXCEEDED",
		Description: "Exceeded the server response timeout.",
	}
)
// Can be added as many as need like belows examples
// 400	BAD_CONTINUATION_TOKEN	Invalid continuation token passed.
// 400	BAD_PAGE	Page number does not exist or is an invalid format (e.g. negative).
// 400	BAD_REQUEST	The resource you’re creating already exists.
// 400	INVALID_ARGUMENT	Invalid argument value passed.
// 400	INVALID_AUTH	Authentication/OAuth token is invalid.
// 400	INVALID_AUTH_HEADER	Authentication header is invalid.
// 400	INVALID_BATCH	Batched request is missing or invalid.
// 400	INVALID_BODY	A request body that was not in JSON format was passed.
// 400	UNSUPPORTED_OPERATION	Requested operation not supported.
// 401	ACCESS_DENIED	Authentication unsuccessful.
// 401	NO_AUTH	Authentication not provided.
// 403	NOT_AUTHORIZED	User has not been authorized to perform that action.
// 404	NOT_FOUND	Invalid URL.
// 405	METHOD_NOT_ALLOWED	Method is not allowed for this endpoint.
// 409	REQUEST_CONFLICT	Requested operation resulted in conflict.
// 429	HIT_RATE_LIMIT	Hourly rate limit has been reached for this token. Default rate limits are 2,000 calls per hour.
// 500	EXPANSION_FAILED	Unhandled error occurred during expansion; the request is likely to succeed if you don’t ask for expansions, but contact Eventbrite support if this problem persists.
// 500	INTERNAL_ERROR	Unhandled error occurred in Eventbrite. contact Eventbrite support if this problem persists.
