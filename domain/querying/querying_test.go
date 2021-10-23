package querying

import (
	"context"
	"errors"
	"testing"
)

type formatMock struct{}

func (mock *formatMock) ByID(ctx context.Context, id *string) ([]byte, error) {
	if id == nil {
		return nil, errors.New("id is required param")
	}

	templateQuery := `
		select attr1, attr2 from test where attr1 = {{ .id }};
	`

	templateFormat := `
		[
			{
				
			}
		]
	`

	data := map[string][]byte{
		"1": []byte(templateQuery),
		"2": []byte(templateFormat),
	}

	temp, ok := data[*id]
	if !ok {
		return nil, errors.New("invalid id")
	}

	return temp, nil
}

func Test_querying_Execute(t *testing.T) {
	//type fields struct {
	//	format    Format
	//	formatter Formatter
	//	querying  Querying
	//}
	//type args struct {
	//	ctx      context.Context
	//	queryID  *string
	//	formatID *string
	//	value    interface{}
	//}
	//tests := []struct {
	//	name    string
	//	fields  fields
	//	args    args
	//	want    []byte
	//	wantErr bool
	//}{
	//	// TODO: Add test cases.
	//}
	//for _, tt := range tests {
	//	t.Run(tt.name, func(t *testing.T) {
	//		mod := &querying{
	//			format:    tt.fields.format,
	//			formatter: tt.fields.formatter,
	//			querying:  tt.fields.querying,
	//		}
	//		got, err := mod.Execute(tt.args.ctx, tt.args.queryID, tt.args.formatID, tt.args.value)
	//		if (err != nil) != tt.wantErr {
	//			t.Errorf("Execute() error = %v, wantErr %v", err, tt.wantErr)
	//			return
	//		}
	//		if !reflect.DeepEqual(got, tt.want) {
	//			t.Errorf("Execute() got = %v, want %v", got, tt.want)
	//		}
	//	})
	//}
}
