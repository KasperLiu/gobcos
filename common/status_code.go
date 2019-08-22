package common

const (
	Success                    string = "0x0"
	Unknown                    string = "0x1"
	BadRLP                     string = "0x2"
	InvalidFormat              string = "0x3"
	OutOfGasIntrinsic          string = "0x4"
    InvalidSignature           string = "0x5"
    InvalidNonce               string = "0x6"
    NotEnoughCash              string = "0x7"
    OutOfGasBase               string = "0x8"
    BlockGasLimitReached       string = "0x9"
    BadInstruction             string = "0xa"
    BadJumpDestination         string = "0xb"
    OutOfGas                   string = "0xc"
    OutOfStack                 string = "0xd"
    StackUnderflow             string = "0xe"
    NonceCheckFail             string = "0xf"
    BlockLimitCheckFail        string = "0x10"
    FilterCheckFail            string = "0x11"
    NoDeployPermission         string = "0x12"
    NoCallPermission           string = "0x13"
    NoTxPermission             string = "0x14"
    PrecompiledError           string = "0x15"
    RevertInstruction          string = "0x16"
    InvalidZeroSignatureFormat string = "0x17"
    AddressAlreadyUsed         string = "0x18"
    PermissionDenied           string = "0x19"
    CallAddressError           string = "0x1a"
)

// GetStatusMessage returns the status message
func GetStatusMessage(status string) string {
	var message string 
	switch (status) {
	case Success:
		message = "success"
		break
	case Unknown:
		message = "unknown"
		break
	case BadRLP:
		message = "bad RLP"
		break
	case InvalidFormat:
		message = "invalid format"
		break
	case OutOfGasIntrinsic:
		message = "out of gas intrinsic"
		break
	case InvalidSignature:
		message = "invalid signature"
		break
	case InvalidNonce:
		message = "invalid nonce"
		break
	case NotEnoughCash:
		message = "not enough cash"
		break
	case OutOfGasBase:
		message = "out of gas base"
		break
	case BlockGasLimitReached:
		message = "block gas limit reached"
		break
	case BadInstruction:
		message = "bad instruction"
		break
	case BadJumpDestination:
		message = "bad jump destination"
		break
	case OutOfGas:
		message = "out of gas"
		break
	case OutOfStack:
		message = "out of stack"
		break
	case StackUnderflow:
		message = "stack underflow"
		break
	case NonceCheckFail:
		message = "nonce check fail"
		break
	case BlockLimitCheckFail:
		message = "block limit check fail"
		break
	case FilterCheckFail:
		message = "filter check fail"
		break
	case NoDeployPermission:
		message = "no deploy permission"
		break
	case NoCallPermission:
		message = "no call permission"
		break
	case NoTxPermission:
		message = "no tx permission"
		break
	case PrecompiledError:
		message = "precompiled error"
		break
	case RevertInstruction:
		message = "revert instruction"
		break
	case InvalidZeroSignatureFormat:
		message = "invalid zero signature format"
		break
	case AddressAlreadyUsed:
		message = "address already used"
		break
	case PermissionDenied:
		message = "permission denied"
		break
	case CallAddressError:
		message = "call address error"
		break
	default:
		message = status
		break
	}

	return message
}