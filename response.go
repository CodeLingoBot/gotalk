package gotalk
import "time"

type Response struct {
  MsgType
  Data []byte
  Wait time.Duration   // only valid when IsRetry()==true
}

// Error returns a string describing the error, when IsError()==true
func (r *Response) Error() string {
  return string(r.Data)
}

// IsError; True if this response is a requestor error (ErrorResult)
func (r *Response) IsError() bool {
  return r.MsgType == MsgTypeErrorRes
}

// IsRetry; True if response is a "server can't handle it right now, please retry" (RetryResult)
func (r *Response) IsRetry() bool {
  return r.MsgType == MsgTypeRetryRes
}

// IsStreaming; True if this is part of a streaming response (StreamResult)
func (r *Response) IsStreaming() bool {
  return r.MsgType == MsgTypeStreamRes
}
