package dsunit

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseColumnarData(t *testing.T) {
	data := "id,name,delays,active,income\n1,A\"bc,3,true,1232.3\n2,Bbc,8,FALSE,0.33\n3,\"X\\\"b,c\",8,,2.53"
	reader := bytes.NewReader([]byte(data))
	columns, rows := parseColumnarData(reader, ",")
	assert.Equal(t, 5, len(columns), "should have 5 columns")
	assert.Equal(t, 3, len(rows), "should have 3 rows")
	assert.Equal(t, 1, rows[0][0])
	assert.Equal(t, "A\"bc", rows[0][1])
	assert.Equal(t, 3, rows[0][2])
	assert.Equal(t, true, rows[0][3])
	assert.Equal(t, 1232.3, rows[0][4])

	assert.Equal(t, 3, rows[2][0])
	assert.Equal(t, "X\"b,c", rows[2][1])
	assert.Equal(t, 8, rows[2][2])
	assert.Equal(t, nil, rows[2][3])
	assert.Equal(t, 2.53, rows[2][4])
}
