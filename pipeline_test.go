package gopipeline

import (
	"reflect"
	"strconv"
	"testing"
)

func Test_Pipeline(t *testing.T) {
	t.Run("case1", func(t *testing.T) {
		ctx := t.Context()
		pipeline := New2(
			ctx,
			From([]int{1, 2, 3, 4, 5}),
			Map[int, string](func(i int) (string, error) {
				return strconv.Itoa(i), nil
			}),
			Map[string, string](func(s string) (string, error) {
				return s + "_" + s, nil
			}),
		)
		result := Collect(ctx, pipeline)
		if !reflect.DeepEqual([]string{
			"1_1",
			"2_2",
			"3_3",
			"4_4",
			"5_5",
		}, result) {
			t.Errorf("unexpected result: %T", result)
		}
	})

}
