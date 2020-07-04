package webdriver

import (
	"context"
	"github.com/MontFerret/ferret/pkg/drivers"
	"github.com/MontFerret/ferret/pkg/runtime/core"
	"github.com/MontFerret/ferret/pkg/runtime/values"
)

type HTMLPage struct {
}

func (p *HTMLPage) MarshalJSON() ([]byte, error) {
	return nil, UnimplementedError
}

func (p*HTMLPage) Type() core.Type  {
	return drivers.HTMLPageType
}

func (p *HTMLPage) String() string {
	return "UnimplementedError"
}

func (p *HTMLPage) Compare(other core.Value) int64 {
	return 1
}

func (p *HTMLPage) Unwrap() interface{} {
	return nil
}

func (p *HTMLPage) Hash() uint64 {
	return 121212
}

func (p *HTMLPage) Copy() core.Value {
	return values.None
}

func (p *HTMLPage) GetIn(ctx context.Context, path []core.Value) (core.Value, error) {
	return values.None, nil
}

func (p *HTMLPage) SetIn(ctx context.Context, path []core.Value, value core.Value) error {
	return nil
}

func (p *HTMLPage) Iterate(ctx context.Context) (core.Iterator, error) {
	return nil, nil
}

func (p *HTMLPage) Length() values.Int {
	return values.NewInt(0)
}

func (p *HTMLPage) Close() error {
	return UnimplementedError
}

func (p *HTMLPage) IsClosed() values.Boolean {
	return false
}

func (p *HTMLPage) GetURL() values.String {
	return ""
}

func (p *HTMLPage) GetMainFrame() drivers.HTMLDocument {
	return nil
}

func (p *HTMLPage) GetFrames(ctx context.Context) (*values.Array, error) {
	return values.NewArray(0), UnimplementedError
}

func (p *HTMLPage) GetFrame(ctx context.Context, idx values.Int) (core.Value, error) {
	return values.None, UnimplementedError
}

func (p *HTMLPage) GetCookies(ctx context.Context) (drivers.HTTPCookies, error) {
	return nil, UnimplementedError
}

func (p *HTMLPage) SetCookies(ctx context.Context, cookies drivers.HTTPCookies) error {
	return UnimplementedError
}

func (p *HTMLPage) DeleteCookies(ctx context.Context, cookies drivers.HTTPCookies) error {
	return UnimplementedError
}

func (p *HTMLPage) PrintToPDF(ctx context.Context, params drivers.PDFParams) (values.Binary, error) {
	return nil, UnimplementedError
}

func (p *HTMLPage) CaptureScreenshot(ctx context.Context, params drivers.ScreenshotParams) (values.Binary, error) {
	return nil, UnimplementedError
}

func (p *HTMLPage) WaitForNavigation(ctx context.Context, targetURL values.String) error {
	return UnimplementedError
}

func (p *HTMLPage) Navigate(ctx context.Context, url values.String) error{
	return UnimplementedError
}

func (p *HTMLPage) NavigateBack(ctx context.Context, skip values.Int) (values.Boolean, error) {
	return false, UnimplementedError
}

func (p *HTMLPage) NavigateForward(ctx context.Context, skip values.Int) (values.Boolean, error) {
	return false, UnimplementedError
}





