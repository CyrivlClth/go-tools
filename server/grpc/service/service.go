// package service
package service

import (
	"context"

	"github.com/CyrivlClth/snowflake/idgen"
	snowflake "github.com/CyrivlClth/snowflake/server/grpc/proto"
)

var _ snowflake.SnowflakeServiceServer = (*Snowflake)(nil)

type Snowflake struct {
	generator idgen.IDGenerator
}

func NewSnowflake(generator idgen.IDGenerator) *Snowflake {
	return &Snowflake{generator: generator}
}

func (s *Snowflake) NextID(context.Context, *snowflake.NextIDReq) (*snowflake.NextIDResp, error) {
	id, err := s.generator.NextID()
	if err != nil {
		return nil, err
	}

	return &snowflake.NextIDResp{Id: id}, nil
}
