package log_toolkit_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	log_toolkit "github.com/validatecl/log-toolkit"
	as "gotest.tools/assert"
)

type example struct {
	message   string
	child     child
	childs    []child
	ptrChild  *child
	ptrChilds *[]child
	errors    errorMessage
}

type child struct {
	messagechild string
}

type errorMessage struct {
	messageerror string
	err          error
}

func TestMarshallSimpleFields(t *testing.T) {

	x := make(map[int]string)
	x[1] = "hola"
	x[2] = "chao"

	y := make(map[string]string)
	y["prueba"] = "prueba"

	marshaller := log_toolkit.NewFieldMarshaller()

	fields := marshaller.MarshalFields(x, y)

	as.Equal(t, len(fields), 3)

}

func TestMarshalFields(t *testing.T) {
	x := []child{child{"valuechild3"}, child{"valuechild4"}}
	y := make(map[string]string)
	y["map1"] = "map_1"
	y["map2"] = "map_2"

	t.Run("Verificando cantidad de registros.(Success)", func(t *testing.T) {
		marshaller := log_toolkit.NewFieldMarshaller()
		resultFields := marshaller.MarshalFields(example{"valueexample1", child{"valuechild"}, x, &child{"valueptrChild"}, &x, errorMessage{"message 1", errors.New("error 1")}}, y)
		assert.Equal(t, 11, len(resultFields))

		assert.Equal(t, resultFields[0].Key, "message")
		assert.Equal(t, "valueexample1", resultFields[0].String)

		assert.Equal(t, resultFields[1].Key, "messagechild")
		assert.Equal(t, "valuechild", resultFields[1].String)

		assert.Equal(t, resultFields[2].Key, "messagechild")
		assert.Equal(t, "valuechild3", resultFields[2].String)

		assert.Equal(t, resultFields[3].Key, "messagechild")
		assert.Equal(t, "valuechild4", resultFields[3].String)

		assert.Equal(t, resultFields[4].Key, "messagechild")
		assert.Equal(t, "valueptrChild", resultFields[4].String)

		assert.Equal(t, resultFields[5].Key, "messagechild")
		assert.Equal(t, "valuechild3", resultFields[5].String)

		assert.Equal(t, resultFields[6].Key, "messagechild")
		assert.Equal(t, "valuechild4", resultFields[6].String)

		assert.Equal(t, resultFields[7].Key, "messageerror")
		assert.Equal(t, "message 1", resultFields[7].String)

		assert.Equal(t, resultFields[8].Key, "err")
		assert.Equal(t, "Error: {error 1}", resultFields[8].String)

		assert.Equal(t, resultFields[9].Key, "map1")
		assert.Equal(t, "map_1", resultFields[9].String)

		assert.Equal(t, resultFields[10].Key, "map2")
		assert.Equal(t, "map_2", resultFields[10].String)

		fmt.Printf("Result: %v", resultFields)
	})

}
