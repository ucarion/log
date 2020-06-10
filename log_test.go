package log_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ucarion/log"
)

type debugLogger struct {
	msg    string
	fields map[string]interface{}
}

func (d *debugLogger) Log(msg string, fields map[string]interface{}) {
	d.msg = msg
	d.fields = fields
}

func TestLog(t *testing.T) {
	logger := debugLogger{}
	log.Logger = &logger

	ctx1 := log.NewContext(context.TODO())

	log.Set(ctx1, "foo", "bar")
	log.Log(ctx1, "msg1")

	assert.Equal(t, debugLogger{
		msg: "msg1",
		fields: map[string]interface{}{
			"foo": "bar",
		},
	}, logger)

	log.Set(ctx1, "foo", "bar2")
	log.Log(ctx1, "msg2")

	assert.Equal(t, debugLogger{
		msg: "msg2",
		fields: map[string]interface{}{
			"foo": "bar2",
		},
	}, logger)

	log.Set(ctx1, "baz", "quux")
	log.Log(ctx1, "msg3")

	assert.Equal(t, debugLogger{
		msg: "msg3",
		fields: map[string]interface{}{
			"foo": "bar2",
			"baz": "quux",
		},
	}, logger)

	ctx2 := log.WithPrefix(ctx1, "prefix1")

	log.Set(ctx2, "foo", "bar")
	log.Log(ctx1, "msg4")

	assert.Equal(t, debugLogger{
		msg: "msg4",
		fields: map[string]interface{}{
			"foo": "bar2",
			"baz": "quux",
			"prefix1": map[string]interface{}{
				"foo": "bar",
			},
		},
	}, logger)

	ctx3 := log.WithPrefix(ctx2, "prefix2")
	ctx4 := log.WithPrefix(ctx3, "prefix3")

	log.Set(ctx4, "foo", "bar")
	log.Log(ctx1, "msg5")

	assert.Equal(t, debugLogger{
		msg: "msg5",
		fields: map[string]interface{}{
			"foo": "bar2",
			"baz": "quux",
			"prefix1": map[string]interface{}{
				"foo": "bar",
				"prefix2": map[string]interface{}{
					"prefix3": map[string]interface{}{
						"foo": "bar",
					},
				},
			},
		},
	}, logger)

	log.Set(ctx3, "foo", "bar")
	log.Log(ctx1, "msg6")

	assert.Equal(t, debugLogger{
		msg: "msg6",
		fields: map[string]interface{}{
			"foo": "bar2",
			"baz": "quux",
			"prefix1": map[string]interface{}{
				"foo": "bar",
				"prefix2": map[string]interface{}{
					"foo": "bar",
					"prefix3": map[string]interface{}{
						"foo": "bar",
					},
				},
			},
		},
	}, logger)
}
