// Code generated; DO NOT EDIT.

//go:generate go run ctc_extra_gen.go
//go:generate go fmt ctc_extra.go

package ctc

const (

	// ForegroundBlack = ForegroundBlack
	// ForegroundRed = ForegroundRed
	// ForegroundGreen = ForegroundGreen
	// ForegroundYellow = ForegroundYellow
	// ForegroundBlue = ForegroundBlue
	// ForegroundMagenta = ForegroundMagenta
	// ForegroundCyan = ForegroundCyan
	// ForegroundWhite = ForegroundWhite
	ForegroundBrightBlack   = ForegroundBright | ForegroundBlack
	ForegroundBrightRed     = ForegroundBright | ForegroundRed
	ForegroundBrightGreen   = ForegroundBright | ForegroundGreen
	ForegroundBrightYellow  = ForegroundBright | ForegroundYellow
	ForegroundBrightBlue    = ForegroundBright | ForegroundBlue
	ForegroundBrightMagenta = ForegroundBright | ForegroundMagenta
	ForegroundBrightCyan    = ForegroundBright | ForegroundCyan
	ForegroundBrightWhite   = ForegroundBright | ForegroundWhite
	// BackgroundBlack = BackgroundBlack
	// BackgroundRed = BackgroundRed
	// BackgroundGreen = BackgroundGreen
	// BackgroundYellow = BackgroundYellow
	// BackgroundBlue = BackgroundBlue
	// BackgroundMagenta = BackgroundMagenta
	// BackgroundCyan = BackgroundCyan
	// BackgroundWhite = BackgroundWhite
	BackgroundBrightBlack                          = BackgroundBright | BackgroundBlack
	BackgroundBrightRed                            = BackgroundBright | BackgroundRed
	BackgroundBrightGreen                          = BackgroundBright | BackgroundGreen
	BackgroundBrightYellow                         = BackgroundBright | BackgroundYellow
	BackgroundBrightBlue                           = BackgroundBright | BackgroundBlue
	BackgroundBrightMagenta                        = BackgroundBright | BackgroundMagenta
	BackgroundBrightCyan                           = BackgroundBright | BackgroundCyan
	BackgroundBrightWhite                          = BackgroundBright | BackgroundWhite
	ForegroundBlackBackgroundBlack                 = ForegroundBlack | BackgroundBlack
	ForegroundRedBackgroundBlack                   = ForegroundRed | BackgroundBlack
	ForegroundGreenBackgroundBlack                 = ForegroundGreen | BackgroundBlack
	ForegroundYellowBackgroundBlack                = ForegroundYellow | BackgroundBlack
	ForegroundBlueBackgroundBlack                  = ForegroundBlue | BackgroundBlack
	ForegroundMagentaBackgroundBlack               = ForegroundMagenta | BackgroundBlack
	ForegroundCyanBackgroundBlack                  = ForegroundCyan | BackgroundBlack
	ForegroundWhiteBackgroundBlack                 = ForegroundWhite | BackgroundBlack
	ForegroundBrightBlackBackgroundBlack           = ForegroundBright | ForegroundBlack | BackgroundBlack
	ForegroundBrightRedBackgroundBlack             = ForegroundBright | ForegroundRed | BackgroundBlack
	ForegroundBrightGreenBackgroundBlack           = ForegroundBright | ForegroundGreen | BackgroundBlack
	ForegroundBrightYellowBackgroundBlack          = ForegroundBright | ForegroundYellow | BackgroundBlack
	ForegroundBrightBlueBackgroundBlack            = ForegroundBright | ForegroundBlue | BackgroundBlack
	ForegroundBrightMagentaBackgroundBlack         = ForegroundBright | ForegroundMagenta | BackgroundBlack
	ForegroundBrightCyanBackgroundBlack            = ForegroundBright | ForegroundCyan | BackgroundBlack
	ForegroundBrightWhiteBackgroundBlack           = ForegroundBright | ForegroundWhite | BackgroundBlack
	ForegroundBlackBackgroundRed                   = ForegroundBlack | BackgroundRed
	ForegroundRedBackgroundRed                     = ForegroundRed | BackgroundRed
	ForegroundGreenBackgroundRed                   = ForegroundGreen | BackgroundRed
	ForegroundYellowBackgroundRed                  = ForegroundYellow | BackgroundRed
	ForegroundBlueBackgroundRed                    = ForegroundBlue | BackgroundRed
	ForegroundMagentaBackgroundRed                 = ForegroundMagenta | BackgroundRed
	ForegroundCyanBackgroundRed                    = ForegroundCyan | BackgroundRed
	ForegroundWhiteBackgroundRed                   = ForegroundWhite | BackgroundRed
	ForegroundBrightBlackBackgroundRed             = ForegroundBright | ForegroundBlack | BackgroundRed
	ForegroundBrightRedBackgroundRed               = ForegroundBright | ForegroundRed | BackgroundRed
	ForegroundBrightGreenBackgroundRed             = ForegroundBright | ForegroundGreen | BackgroundRed
	ForegroundBrightYellowBackgroundRed            = ForegroundBright | ForegroundYellow | BackgroundRed
	ForegroundBrightBlueBackgroundRed              = ForegroundBright | ForegroundBlue | BackgroundRed
	ForegroundBrightMagentaBackgroundRed           = ForegroundBright | ForegroundMagenta | BackgroundRed
	ForegroundBrightCyanBackgroundRed              = ForegroundBright | ForegroundCyan | BackgroundRed
	ForegroundBrightWhiteBackgroundRed             = ForegroundBright | ForegroundWhite | BackgroundRed
	ForegroundBlackBackgroundGreen                 = ForegroundBlack | BackgroundGreen
	ForegroundRedBackgroundGreen                   = ForegroundRed | BackgroundGreen
	ForegroundGreenBackgroundGreen                 = ForegroundGreen | BackgroundGreen
	ForegroundYellowBackgroundGreen                = ForegroundYellow | BackgroundGreen
	ForegroundBlueBackgroundGreen                  = ForegroundBlue | BackgroundGreen
	ForegroundMagentaBackgroundGreen               = ForegroundMagenta | BackgroundGreen
	ForegroundCyanBackgroundGreen                  = ForegroundCyan | BackgroundGreen
	ForegroundWhiteBackgroundGreen                 = ForegroundWhite | BackgroundGreen
	ForegroundBrightBlackBackgroundGreen           = ForegroundBright | ForegroundBlack | BackgroundGreen
	ForegroundBrightRedBackgroundGreen             = ForegroundBright | ForegroundRed | BackgroundGreen
	ForegroundBrightGreenBackgroundGreen           = ForegroundBright | ForegroundGreen | BackgroundGreen
	ForegroundBrightYellowBackgroundGreen          = ForegroundBright | ForegroundYellow | BackgroundGreen
	ForegroundBrightBlueBackgroundGreen            = ForegroundBright | ForegroundBlue | BackgroundGreen
	ForegroundBrightMagentaBackgroundGreen         = ForegroundBright | ForegroundMagenta | BackgroundGreen
	ForegroundBrightCyanBackgroundGreen            = ForegroundBright | ForegroundCyan | BackgroundGreen
	ForegroundBrightWhiteBackgroundGreen           = ForegroundBright | ForegroundWhite | BackgroundGreen
	ForegroundBlackBackgroundYellow                = ForegroundBlack | BackgroundYellow
	ForegroundRedBackgroundYellow                  = ForegroundRed | BackgroundYellow
	ForegroundGreenBackgroundYellow                = ForegroundGreen | BackgroundYellow
	ForegroundYellowBackgroundYellow               = ForegroundYellow | BackgroundYellow
	ForegroundBlueBackgroundYellow                 = ForegroundBlue | BackgroundYellow
	ForegroundMagentaBackgroundYellow              = ForegroundMagenta | BackgroundYellow
	ForegroundCyanBackgroundYellow                 = ForegroundCyan | BackgroundYellow
	ForegroundWhiteBackgroundYellow                = ForegroundWhite | BackgroundYellow
	ForegroundBrightBlackBackgroundYellow          = ForegroundBright | ForegroundBlack | BackgroundYellow
	ForegroundBrightRedBackgroundYellow            = ForegroundBright | ForegroundRed | BackgroundYellow
	ForegroundBrightGreenBackgroundYellow          = ForegroundBright | ForegroundGreen | BackgroundYellow
	ForegroundBrightYellowBackgroundYellow         = ForegroundBright | ForegroundYellow | BackgroundYellow
	ForegroundBrightBlueBackgroundYellow           = ForegroundBright | ForegroundBlue | BackgroundYellow
	ForegroundBrightMagentaBackgroundYellow        = ForegroundBright | ForegroundMagenta | BackgroundYellow
	ForegroundBrightCyanBackgroundYellow           = ForegroundBright | ForegroundCyan | BackgroundYellow
	ForegroundBrightWhiteBackgroundYellow          = ForegroundBright | ForegroundWhite | BackgroundYellow
	ForegroundBlackBackgroundBlue                  = ForegroundBlack | BackgroundBlue
	ForegroundRedBackgroundBlue                    = ForegroundRed | BackgroundBlue
	ForegroundGreenBackgroundBlue                  = ForegroundGreen | BackgroundBlue
	ForegroundYellowBackgroundBlue                 = ForegroundYellow | BackgroundBlue
	ForegroundBlueBackgroundBlue                   = ForegroundBlue | BackgroundBlue
	ForegroundMagentaBackgroundBlue                = ForegroundMagenta | BackgroundBlue
	ForegroundCyanBackgroundBlue                   = ForegroundCyan | BackgroundBlue
	ForegroundWhiteBackgroundBlue                  = ForegroundWhite | BackgroundBlue
	ForegroundBrightBlackBackgroundBlue            = ForegroundBright | ForegroundBlack | BackgroundBlue
	ForegroundBrightRedBackgroundBlue              = ForegroundBright | ForegroundRed | BackgroundBlue
	ForegroundBrightGreenBackgroundBlue            = ForegroundBright | ForegroundGreen | BackgroundBlue
	ForegroundBrightYellowBackgroundBlue           = ForegroundBright | ForegroundYellow | BackgroundBlue
	ForegroundBrightBlueBackgroundBlue             = ForegroundBright | ForegroundBlue | BackgroundBlue
	ForegroundBrightMagentaBackgroundBlue          = ForegroundBright | ForegroundMagenta | BackgroundBlue
	ForegroundBrightCyanBackgroundBlue             = ForegroundBright | ForegroundCyan | BackgroundBlue
	ForegroundBrightWhiteBackgroundBlue            = ForegroundBright | ForegroundWhite | BackgroundBlue
	ForegroundBlackBackgroundMagenta               = ForegroundBlack | BackgroundMagenta
	ForegroundRedBackgroundMagenta                 = ForegroundRed | BackgroundMagenta
	ForegroundGreenBackgroundMagenta               = ForegroundGreen | BackgroundMagenta
	ForegroundYellowBackgroundMagenta              = ForegroundYellow | BackgroundMagenta
	ForegroundBlueBackgroundMagenta                = ForegroundBlue | BackgroundMagenta
	ForegroundMagentaBackgroundMagenta             = ForegroundMagenta | BackgroundMagenta
	ForegroundCyanBackgroundMagenta                = ForegroundCyan | BackgroundMagenta
	ForegroundWhiteBackgroundMagenta               = ForegroundWhite | BackgroundMagenta
	ForegroundBrightBlackBackgroundMagenta         = ForegroundBright | ForegroundBlack | BackgroundMagenta
	ForegroundBrightRedBackgroundMagenta           = ForegroundBright | ForegroundRed | BackgroundMagenta
	ForegroundBrightGreenBackgroundMagenta         = ForegroundBright | ForegroundGreen | BackgroundMagenta
	ForegroundBrightYellowBackgroundMagenta        = ForegroundBright | ForegroundYellow | BackgroundMagenta
	ForegroundBrightBlueBackgroundMagenta          = ForegroundBright | ForegroundBlue | BackgroundMagenta
	ForegroundBrightMagentaBackgroundMagenta       = ForegroundBright | ForegroundMagenta | BackgroundMagenta
	ForegroundBrightCyanBackgroundMagenta          = ForegroundBright | ForegroundCyan | BackgroundMagenta
	ForegroundBrightWhiteBackgroundMagenta         = ForegroundBright | ForegroundWhite | BackgroundMagenta
	ForegroundBlackBackgroundCyan                  = ForegroundBlack | BackgroundCyan
	ForegroundRedBackgroundCyan                    = ForegroundRed | BackgroundCyan
	ForegroundGreenBackgroundCyan                  = ForegroundGreen | BackgroundCyan
	ForegroundYellowBackgroundCyan                 = ForegroundYellow | BackgroundCyan
	ForegroundBlueBackgroundCyan                   = ForegroundBlue | BackgroundCyan
	ForegroundMagentaBackgroundCyan                = ForegroundMagenta | BackgroundCyan
	ForegroundCyanBackgroundCyan                   = ForegroundCyan | BackgroundCyan
	ForegroundWhiteBackgroundCyan                  = ForegroundWhite | BackgroundCyan
	ForegroundBrightBlackBackgroundCyan            = ForegroundBright | ForegroundBlack | BackgroundCyan
	ForegroundBrightRedBackgroundCyan              = ForegroundBright | ForegroundRed | BackgroundCyan
	ForegroundBrightGreenBackgroundCyan            = ForegroundBright | ForegroundGreen | BackgroundCyan
	ForegroundBrightYellowBackgroundCyan           = ForegroundBright | ForegroundYellow | BackgroundCyan
	ForegroundBrightBlueBackgroundCyan             = ForegroundBright | ForegroundBlue | BackgroundCyan
	ForegroundBrightMagentaBackgroundCyan          = ForegroundBright | ForegroundMagenta | BackgroundCyan
	ForegroundBrightCyanBackgroundCyan             = ForegroundBright | ForegroundCyan | BackgroundCyan
	ForegroundBrightWhiteBackgroundCyan            = ForegroundBright | ForegroundWhite | BackgroundCyan
	ForegroundBlackBackgroundWhite                 = ForegroundBlack | BackgroundWhite
	ForegroundRedBackgroundWhite                   = ForegroundRed | BackgroundWhite
	ForegroundGreenBackgroundWhite                 = ForegroundGreen | BackgroundWhite
	ForegroundYellowBackgroundWhite                = ForegroundYellow | BackgroundWhite
	ForegroundBlueBackgroundWhite                  = ForegroundBlue | BackgroundWhite
	ForegroundMagentaBackgroundWhite               = ForegroundMagenta | BackgroundWhite
	ForegroundCyanBackgroundWhite                  = ForegroundCyan | BackgroundWhite
	ForegroundWhiteBackgroundWhite                 = ForegroundWhite | BackgroundWhite
	ForegroundBrightBlackBackgroundWhite           = ForegroundBright | ForegroundBlack | BackgroundWhite
	ForegroundBrightRedBackgroundWhite             = ForegroundBright | ForegroundRed | BackgroundWhite
	ForegroundBrightGreenBackgroundWhite           = ForegroundBright | ForegroundGreen | BackgroundWhite
	ForegroundBrightYellowBackgroundWhite          = ForegroundBright | ForegroundYellow | BackgroundWhite
	ForegroundBrightBlueBackgroundWhite            = ForegroundBright | ForegroundBlue | BackgroundWhite
	ForegroundBrightMagentaBackgroundWhite         = ForegroundBright | ForegroundMagenta | BackgroundWhite
	ForegroundBrightCyanBackgroundWhite            = ForegroundBright | ForegroundCyan | BackgroundWhite
	ForegroundBrightWhiteBackgroundWhite           = ForegroundBright | ForegroundWhite | BackgroundWhite
	ForegroundBlackBackgroundBrightBlack           = ForegroundBlack | BackgroundBright | BackgroundBlack
	ForegroundRedBackgroundBrightBlack             = ForegroundRed | BackgroundBright | BackgroundBlack
	ForegroundGreenBackgroundBrightBlack           = ForegroundGreen | BackgroundBright | BackgroundBlack
	ForegroundYellowBackgroundBrightBlack          = ForegroundYellow | BackgroundBright | BackgroundBlack
	ForegroundBlueBackgroundBrightBlack            = ForegroundBlue | BackgroundBright | BackgroundBlack
	ForegroundMagentaBackgroundBrightBlack         = ForegroundMagenta | BackgroundBright | BackgroundBlack
	ForegroundCyanBackgroundBrightBlack            = ForegroundCyan | BackgroundBright | BackgroundBlack
	ForegroundWhiteBackgroundBrightBlack           = ForegroundWhite | BackgroundBright | BackgroundBlack
	ForegroundBrightBlackBackgroundBrightBlack     = ForegroundBright | ForegroundBlack | BackgroundBright | BackgroundBlack
	ForegroundBrightRedBackgroundBrightBlack       = ForegroundBright | ForegroundRed | BackgroundBright | BackgroundBlack
	ForegroundBrightGreenBackgroundBrightBlack     = ForegroundBright | ForegroundGreen | BackgroundBright | BackgroundBlack
	ForegroundBrightYellowBackgroundBrightBlack    = ForegroundBright | ForegroundYellow | BackgroundBright | BackgroundBlack
	ForegroundBrightBlueBackgroundBrightBlack      = ForegroundBright | ForegroundBlue | BackgroundBright | BackgroundBlack
	ForegroundBrightMagentaBackgroundBrightBlack   = ForegroundBright | ForegroundMagenta | BackgroundBright | BackgroundBlack
	ForegroundBrightCyanBackgroundBrightBlack      = ForegroundBright | ForegroundCyan | BackgroundBright | BackgroundBlack
	ForegroundBrightWhiteBackgroundBrightBlack     = ForegroundBright | ForegroundWhite | BackgroundBright | BackgroundBlack
	ForegroundBlackBackgroundBrightRed             = ForegroundBlack | BackgroundBright | BackgroundRed
	ForegroundRedBackgroundBrightRed               = ForegroundRed | BackgroundBright | BackgroundRed
	ForegroundGreenBackgroundBrightRed             = ForegroundGreen | BackgroundBright | BackgroundRed
	ForegroundYellowBackgroundBrightRed            = ForegroundYellow | BackgroundBright | BackgroundRed
	ForegroundBlueBackgroundBrightRed              = ForegroundBlue | BackgroundBright | BackgroundRed
	ForegroundMagentaBackgroundBrightRed           = ForegroundMagenta | BackgroundBright | BackgroundRed
	ForegroundCyanBackgroundBrightRed              = ForegroundCyan | BackgroundBright | BackgroundRed
	ForegroundWhiteBackgroundBrightRed             = ForegroundWhite | BackgroundBright | BackgroundRed
	ForegroundBrightBlackBackgroundBrightRed       = ForegroundBright | ForegroundBlack | BackgroundBright | BackgroundRed
	ForegroundBrightRedBackgroundBrightRed         = ForegroundBright | ForegroundRed | BackgroundBright | BackgroundRed
	ForegroundBrightGreenBackgroundBrightRed       = ForegroundBright | ForegroundGreen | BackgroundBright | BackgroundRed
	ForegroundBrightYellowBackgroundBrightRed      = ForegroundBright | ForegroundYellow | BackgroundBright | BackgroundRed
	ForegroundBrightBlueBackgroundBrightRed        = ForegroundBright | ForegroundBlue | BackgroundBright | BackgroundRed
	ForegroundBrightMagentaBackgroundBrightRed     = ForegroundBright | ForegroundMagenta | BackgroundBright | BackgroundRed
	ForegroundBrightCyanBackgroundBrightRed        = ForegroundBright | ForegroundCyan | BackgroundBright | BackgroundRed
	ForegroundBrightWhiteBackgroundBrightRed       = ForegroundBright | ForegroundWhite | BackgroundBright | BackgroundRed
	ForegroundBlackBackgroundBrightGreen           = ForegroundBlack | BackgroundBright | BackgroundGreen
	ForegroundRedBackgroundBrightGreen             = ForegroundRed | BackgroundBright | BackgroundGreen
	ForegroundGreenBackgroundBrightGreen           = ForegroundGreen | BackgroundBright | BackgroundGreen
	ForegroundYellowBackgroundBrightGreen          = ForegroundYellow | BackgroundBright | BackgroundGreen
	ForegroundBlueBackgroundBrightGreen            = ForegroundBlue | BackgroundBright | BackgroundGreen
	ForegroundMagentaBackgroundBrightGreen         = ForegroundMagenta | BackgroundBright | BackgroundGreen
	ForegroundCyanBackgroundBrightGreen            = ForegroundCyan | BackgroundBright | BackgroundGreen
	ForegroundWhiteBackgroundBrightGreen           = ForegroundWhite | BackgroundBright | BackgroundGreen
	ForegroundBrightBlackBackgroundBrightGreen     = ForegroundBright | ForegroundBlack | BackgroundBright | BackgroundGreen
	ForegroundBrightRedBackgroundBrightGreen       = ForegroundBright | ForegroundRed | BackgroundBright | BackgroundGreen
	ForegroundBrightGreenBackgroundBrightGreen     = ForegroundBright | ForegroundGreen | BackgroundBright | BackgroundGreen
	ForegroundBrightYellowBackgroundBrightGreen    = ForegroundBright | ForegroundYellow | BackgroundBright | BackgroundGreen
	ForegroundBrightBlueBackgroundBrightGreen      = ForegroundBright | ForegroundBlue | BackgroundBright | BackgroundGreen
	ForegroundBrightMagentaBackgroundBrightGreen   = ForegroundBright | ForegroundMagenta | BackgroundBright | BackgroundGreen
	ForegroundBrightCyanBackgroundBrightGreen      = ForegroundBright | ForegroundCyan | BackgroundBright | BackgroundGreen
	ForegroundBrightWhiteBackgroundBrightGreen     = ForegroundBright | ForegroundWhite | BackgroundBright | BackgroundGreen
	ForegroundBlackBackgroundBrightYellow          = ForegroundBlack | BackgroundBright | BackgroundYellow
	ForegroundRedBackgroundBrightYellow            = ForegroundRed | BackgroundBright | BackgroundYellow
	ForegroundGreenBackgroundBrightYellow          = ForegroundGreen | BackgroundBright | BackgroundYellow
	ForegroundYellowBackgroundBrightYellow         = ForegroundYellow | BackgroundBright | BackgroundYellow
	ForegroundBlueBackgroundBrightYellow           = ForegroundBlue | BackgroundBright | BackgroundYellow
	ForegroundMagentaBackgroundBrightYellow        = ForegroundMagenta | BackgroundBright | BackgroundYellow
	ForegroundCyanBackgroundBrightYellow           = ForegroundCyan | BackgroundBright | BackgroundYellow
	ForegroundWhiteBackgroundBrightYellow          = ForegroundWhite | BackgroundBright | BackgroundYellow
	ForegroundBrightBlackBackgroundBrightYellow    = ForegroundBright | ForegroundBlack | BackgroundBright | BackgroundYellow
	ForegroundBrightRedBackgroundBrightYellow      = ForegroundBright | ForegroundRed | BackgroundBright | BackgroundYellow
	ForegroundBrightGreenBackgroundBrightYellow    = ForegroundBright | ForegroundGreen | BackgroundBright | BackgroundYellow
	ForegroundBrightYellowBackgroundBrightYellow   = ForegroundBright | ForegroundYellow | BackgroundBright | BackgroundYellow
	ForegroundBrightBlueBackgroundBrightYellow     = ForegroundBright | ForegroundBlue | BackgroundBright | BackgroundYellow
	ForegroundBrightMagentaBackgroundBrightYellow  = ForegroundBright | ForegroundMagenta | BackgroundBright | BackgroundYellow
	ForegroundBrightCyanBackgroundBrightYellow     = ForegroundBright | ForegroundCyan | BackgroundBright | BackgroundYellow
	ForegroundBrightWhiteBackgroundBrightYellow    = ForegroundBright | ForegroundWhite | BackgroundBright | BackgroundYellow
	ForegroundBlackBackgroundBrightBlue            = ForegroundBlack | BackgroundBright | BackgroundBlue
	ForegroundRedBackgroundBrightBlue              = ForegroundRed | BackgroundBright | BackgroundBlue
	ForegroundGreenBackgroundBrightBlue            = ForegroundGreen | BackgroundBright | BackgroundBlue
	ForegroundYellowBackgroundBrightBlue           = ForegroundYellow | BackgroundBright | BackgroundBlue
	ForegroundBlueBackgroundBrightBlue             = ForegroundBlue | BackgroundBright | BackgroundBlue
	ForegroundMagentaBackgroundBrightBlue          = ForegroundMagenta | BackgroundBright | BackgroundBlue
	ForegroundCyanBackgroundBrightBlue             = ForegroundCyan | BackgroundBright | BackgroundBlue
	ForegroundWhiteBackgroundBrightBlue            = ForegroundWhite | BackgroundBright | BackgroundBlue
	ForegroundBrightBlackBackgroundBrightBlue      = ForegroundBright | ForegroundBlack | BackgroundBright | BackgroundBlue
	ForegroundBrightRedBackgroundBrightBlue        = ForegroundBright | ForegroundRed | BackgroundBright | BackgroundBlue
	ForegroundBrightGreenBackgroundBrightBlue      = ForegroundBright | ForegroundGreen | BackgroundBright | BackgroundBlue
	ForegroundBrightYellowBackgroundBrightBlue     = ForegroundBright | ForegroundYellow | BackgroundBright | BackgroundBlue
	ForegroundBrightBlueBackgroundBrightBlue       = ForegroundBright | ForegroundBlue | BackgroundBright | BackgroundBlue
	ForegroundBrightMagentaBackgroundBrightBlue    = ForegroundBright | ForegroundMagenta | BackgroundBright | BackgroundBlue
	ForegroundBrightCyanBackgroundBrightBlue       = ForegroundBright | ForegroundCyan | BackgroundBright | BackgroundBlue
	ForegroundBrightWhiteBackgroundBrightBlue      = ForegroundBright | ForegroundWhite | BackgroundBright | BackgroundBlue
	ForegroundBlackBackgroundBrightMagenta         = ForegroundBlack | BackgroundBright | BackgroundMagenta
	ForegroundRedBackgroundBrightMagenta           = ForegroundRed | BackgroundBright | BackgroundMagenta
	ForegroundGreenBackgroundBrightMagenta         = ForegroundGreen | BackgroundBright | BackgroundMagenta
	ForegroundYellowBackgroundBrightMagenta        = ForegroundYellow | BackgroundBright | BackgroundMagenta
	ForegroundBlueBackgroundBrightMagenta          = ForegroundBlue | BackgroundBright | BackgroundMagenta
	ForegroundMagentaBackgroundBrightMagenta       = ForegroundMagenta | BackgroundBright | BackgroundMagenta
	ForegroundCyanBackgroundBrightMagenta          = ForegroundCyan | BackgroundBright | BackgroundMagenta
	ForegroundWhiteBackgroundBrightMagenta         = ForegroundWhite | BackgroundBright | BackgroundMagenta
	ForegroundBrightBlackBackgroundBrightMagenta   = ForegroundBright | ForegroundBlack | BackgroundBright | BackgroundMagenta
	ForegroundBrightRedBackgroundBrightMagenta     = ForegroundBright | ForegroundRed | BackgroundBright | BackgroundMagenta
	ForegroundBrightGreenBackgroundBrightMagenta   = ForegroundBright | ForegroundGreen | BackgroundBright | BackgroundMagenta
	ForegroundBrightYellowBackgroundBrightMagenta  = ForegroundBright | ForegroundYellow | BackgroundBright | BackgroundMagenta
	ForegroundBrightBlueBackgroundBrightMagenta    = ForegroundBright | ForegroundBlue | BackgroundBright | BackgroundMagenta
	ForegroundBrightMagentaBackgroundBrightMagenta = ForegroundBright | ForegroundMagenta | BackgroundBright | BackgroundMagenta
	ForegroundBrightCyanBackgroundBrightMagenta    = ForegroundBright | ForegroundCyan | BackgroundBright | BackgroundMagenta
	ForegroundBrightWhiteBackgroundBrightMagenta   = ForegroundBright | ForegroundWhite | BackgroundBright | BackgroundMagenta
	ForegroundBlackBackgroundBrightCyan            = ForegroundBlack | BackgroundBright | BackgroundCyan
	ForegroundRedBackgroundBrightCyan              = ForegroundRed | BackgroundBright | BackgroundCyan
	ForegroundGreenBackgroundBrightCyan            = ForegroundGreen | BackgroundBright | BackgroundCyan
	ForegroundYellowBackgroundBrightCyan           = ForegroundYellow | BackgroundBright | BackgroundCyan
	ForegroundBlueBackgroundBrightCyan             = ForegroundBlue | BackgroundBright | BackgroundCyan
	ForegroundMagentaBackgroundBrightCyan          = ForegroundMagenta | BackgroundBright | BackgroundCyan
	ForegroundCyanBackgroundBrightCyan             = ForegroundCyan | BackgroundBright | BackgroundCyan
	ForegroundWhiteBackgroundBrightCyan            = ForegroundWhite | BackgroundBright | BackgroundCyan
	ForegroundBrightBlackBackgroundBrightCyan      = ForegroundBright | ForegroundBlack | BackgroundBright | BackgroundCyan
	ForegroundBrightRedBackgroundBrightCyan        = ForegroundBright | ForegroundRed | BackgroundBright | BackgroundCyan
	ForegroundBrightGreenBackgroundBrightCyan      = ForegroundBright | ForegroundGreen | BackgroundBright | BackgroundCyan
	ForegroundBrightYellowBackgroundBrightCyan     = ForegroundBright | ForegroundYellow | BackgroundBright | BackgroundCyan
	ForegroundBrightBlueBackgroundBrightCyan       = ForegroundBright | ForegroundBlue | BackgroundBright | BackgroundCyan
	ForegroundBrightMagentaBackgroundBrightCyan    = ForegroundBright | ForegroundMagenta | BackgroundBright | BackgroundCyan
	ForegroundBrightCyanBackgroundBrightCyan       = ForegroundBright | ForegroundCyan | BackgroundBright | BackgroundCyan
	ForegroundBrightWhiteBackgroundBrightCyan      = ForegroundBright | ForegroundWhite | BackgroundBright | BackgroundCyan
	ForegroundBlackBackgroundBrightWhite           = ForegroundBlack | BackgroundBright | BackgroundWhite
	ForegroundRedBackgroundBrightWhite             = ForegroundRed | BackgroundBright | BackgroundWhite
	ForegroundGreenBackgroundBrightWhite           = ForegroundGreen | BackgroundBright | BackgroundWhite
	ForegroundYellowBackgroundBrightWhite          = ForegroundYellow | BackgroundBright | BackgroundWhite
	ForegroundBlueBackgroundBrightWhite            = ForegroundBlue | BackgroundBright | BackgroundWhite
	ForegroundMagentaBackgroundBrightWhite         = ForegroundMagenta | BackgroundBright | BackgroundWhite
	ForegroundCyanBackgroundBrightWhite            = ForegroundCyan | BackgroundBright | BackgroundWhite
	ForegroundWhiteBackgroundBrightWhite           = ForegroundWhite | BackgroundBright | BackgroundWhite
	ForegroundBrightBlackBackgroundBrightWhite     = ForegroundBright | ForegroundBlack | BackgroundBright | BackgroundWhite
	ForegroundBrightRedBackgroundBrightWhite       = ForegroundBright | ForegroundRed | BackgroundBright | BackgroundWhite
	ForegroundBrightGreenBackgroundBrightWhite     = ForegroundBright | ForegroundGreen | BackgroundBright | BackgroundWhite
	ForegroundBrightYellowBackgroundBrightWhite    = ForegroundBright | ForegroundYellow | BackgroundBright | BackgroundWhite
	ForegroundBrightBlueBackgroundBrightWhite      = ForegroundBright | ForegroundBlue | BackgroundBright | BackgroundWhite
	ForegroundBrightMagentaBackgroundBrightWhite   = ForegroundBright | ForegroundMagenta | BackgroundBright | BackgroundWhite
	ForegroundBrightCyanBackgroundBrightWhite      = ForegroundBright | ForegroundCyan | BackgroundBright | BackgroundWhite
	ForegroundBrightWhiteBackgroundBrightWhite     = ForegroundBright | ForegroundWhite | BackgroundBright | BackgroundWhite
)
